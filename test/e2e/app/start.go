package app

import (
	"context"
	"sort"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	"github.com/cometbft/cometbft/test/e2e/pkg/infra"
)

func Start(ctx context.Context, testnet *e2e.Testnet, p infra.Provider) error {
	if len(testnet.Nodes) == 0 {
		return errors.New("no nodes in testnet")
	}

	// Nodes are already sorted by name. Sort them by name then startAt,
	// which gives the overall order startAt, mode, name.
	nodeQueue := testnet.Nodes
	sort.SliceStable(nodeQueue, func(i, j int) bool {
		a, b := nodeQueue[i], nodeQueue[j]
		switch {
		case a.Mode == b.Mode:
			return false
		case a.Mode == e2e.ModeSeed:
			return true
		case a.Mode == e2e.ModeValidator && b.Mode == e2e.ModeFull:
			return true
		}

		return false
	})

	sort.SliceStable(nodeQueue, func(i, j int) bool {
		return nodeQueue[i].StartAt < nodeQueue[j].StartAt
	})

	if nodeQueue[0].StartAt > 0 {
		return errors.New("no initial nodes in testnet")
	}

	// Start initial nodes (StartAt: 0)
	log.Info(ctx, "Starting initial network nodes...")
	nodesAtZero := make([]*e2e.Node, 0)
	for len(nodeQueue) > 0 && nodeQueue[0].StartAt == 0 {
		nodesAtZero = append(nodesAtZero, nodeQueue[0])
		nodeQueue = nodeQueue[1:]
	}
	err := p.StartNodes(ctx, nodesAtZero...)
	if err != nil {
		return errors.Wrap(err, "starting initial nodes")
	}
	for _, node := range nodesAtZero {
		if _, err := waitForNode(ctx, node, 0, 15*time.Second); err != nil {
			return err
		}
		log.Info(ctx, "Starting node",
			"name", node.Name,
			"external_ip", node.ExternalIP,
			"proxy_port", node.ProxyPort,
			"prom", node.PrometheusProxyPort,
		)
	}

	networkHeight := testnet.InitialHeight

	// Wait for initial height
	log.Info(ctx, "Waiting for initial height",
		"height", networkHeight,
		"nodes", len(testnet.Nodes)-len(nodeQueue),
		"pending", len(nodeQueue))

	block, blockID, err := waitForHeight(ctx, testnet, networkHeight)
	if err != nil {
		return err
	}

	// Update any state sync nodes with a trusted height and hash
	for _, node := range nodeQueue {
		if node.StateSync || node.Mode == e2e.ModeLight {
			err = UpdateConfigStateSync(node, block.Height, blockID.Hash.Bytes())
			if err != nil {
				return err
			}
		}
	}

	for _, node := range nodeQueue {
		if node.StartAt > networkHeight {
			// if we're starting a node that's ahead of
			// the last known height of the network, then
			// we should make sure that the rest of the
			// network has reached at least the height
			// that this node will start at before we
			// start the node.

			networkHeight = node.StartAt

			log.Info(ctx, "Waiting for network to advance before starting catch up node",
				"node", node.Name,
				"height", networkHeight)

			if _, _, err := waitForHeight(ctx, testnet, networkHeight); err != nil {
				return err
			}
		}

		log.Info(ctx, "Starting catch up node", "node", node.Name, "height", node.StartAt)

		err := p.StartNodes(ctx, node)
		if err != nil {
			return err
		}
		status, err := waitForNode(ctx, node, node.StartAt, 3*time.Minute)
		if err != nil {
			return err
		}
		log.Info(ctx, "Started node", "name", node.Name, "hiehgt", status.SyncInfo.LatestBlockHeight)
	}

	return nil
}