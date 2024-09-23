package app

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/omni-network/omni/e2e"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/netconf"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
)

func GetTypeProvider(data []byte) (*go_fuzz_utils.TypeProvider, error) {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsStringBounds(0, 1024)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsSliceBounds(0, 4096)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsBiases(0, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	return tp, nil
}

func Fuzz_Nosy_lazyNetwork_Init__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *lazyNetwork
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.Init()
	})
}

func Fuzz_Nosy_lazyNetwork_MustBackends__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *lazyNetwork
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.MustBackends()
	})
}

func Fuzz_Nosy_lazyNetwork_MustNetman__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *lazyNetwork
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.MustNetman()
	})
}

func Fuzz_Nosy_lazyNetwork_mustInit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *lazyNetwork
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.mustInit()
	})
}

func Fuzz_Nosy_Create3DeployConfig_Validate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg Create3DeployConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}

		cfg.Validate()
	})
}

func Fuzz_Nosy_Definition_Backends__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg DefinitionConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var commandName string
		fill_err = tp.Fill(&commandName)
		if fill_err != nil {
			return
		}

		d, err := MakeDefinition(ctx, cfg, commandName)
		if err != nil {
			return
		}
		d.Backends()
	})
}

func Fuzz_Nosy_Definition_DeployInfos__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg DefinitionConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var commandName string
		fill_err = tp.Fill(&commandName)
		if fill_err != nil {
			return
		}

		d, err := MakeDefinition(ctx, cfg, commandName)
		if err != nil {
			return
		}
		d.DeployInfos()
	})
}

func Fuzz_Nosy_Definition_InitLazyNetwork__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg DefinitionConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var commandName string
		fill_err = tp.Fill(&commandName)
		if fill_err != nil {
			return
		}

		d, err := MakeDefinition(ctx, cfg, commandName)
		if err != nil {
			return
		}
		d.InitLazyNetwork()
	})
}

func Fuzz_Nosy_Definition_Netman__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg DefinitionConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var commandName string
		fill_err = tp.Fill(&commandName)
		if fill_err != nil {
			return
		}

		d, err := MakeDefinition(ctx, cfg, commandName)
		if err != nil {
			return
		}
		d.Netman()
	})
}

func Fuzz_Nosy_registryMngr_registerPortals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var def Definition
		fill_err = tp.Fill(&def)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}

		m, err := newRegistryMngr(c1, def)
		if err != nil {
			return
		}
		m.registerPortals(c3)
	})
}

func Fuzz_Nosy_StartMonitoringReceipts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var def Definition
		fill_err = tp.Fill(&def)
		if fill_err != nil {
			return
		}

		StartMonitoringReceipts(ctx, def)
	})
}

// skipping Fuzz_Nosy_StartSendingXMsgs__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/e2e/netman.Manager

func Fuzz_Nosy_StartValidatorUpdates__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var def Definition
		fill_err = tp.Fill(&def)
		if fill_err != nil {
			return
		}

		StartValidatorUpdates(ctx, def)
	})
}

func Fuzz_Nosy_accountsToFund__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var network netconf.ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}

		accountsToFund(network)
	})
}

func Fuzz_Nosy_advertisedP2PAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var network netconf.ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var node *e2e.Node
		fill_err = tp.Fill(&node)
		if fill_err != nil {
			return
		}
		if node == nil {
			return
		}

		advertisedP2PAddr(network, node)
	})
}

// skipping Fuzz_Nosy_checkSupportedChains__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/e2e/netman.Manager

func Fuzz_Nosy_deployCreate3__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var def Definition
		fill_err = tp.Fill(&def)
		if fill_err != nil {
			return
		}
		var chainID uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}

		deployCreate3(ctx, def, chainID)
	})
}

func Fuzz_Nosy_etherStr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if amount == nil {
			return
		}

		etherStr(amount)
	})
}

func Fuzz_Nosy_evmChains__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var def Definition
		fill_err = tp.Fill(&def)
		if fill_err != nil {
			return
		}

		evmChains(def)
	})
}

func Fuzz_Nosy_getSortedNodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var testnet *e2e.Testnet
		fill_err = tp.Fill(&testnet)
		if fill_err != nil {
			return
		}
		if testnet == nil {
			return
		}

		getSortedNodes(testnet)
	})
}

func Fuzz_Nosy_isPublicNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var network netconf.ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var mode e2e.Mode
		fill_err = tp.Fill(&mode)
		if fill_err != nil {
			return
		}

		isPublicNode(network, mode)
	})
}

func Fuzz_Nosy_logRPCs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var def Definition
		fill_err = tp.Fill(&def)
		if fill_err != nil {
			return
		}

		logRPCs(ctx, def)
	})
}

func Fuzz_Nosy_makePortalDeps__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var def Definition
		fill_err = tp.Fill(&def)
		if fill_err != nil {
			return
		}

		makePortalDeps(def)
	})
}

func Fuzz_Nosy_noAnvilDev__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var accounts []common.Address
		fill_err = tp.Fill(&accounts)
		if fill_err != nil {
			return
		}

		noAnvilDev(accounts)
	})
}

func Fuzz_Nosy_privatePeerIDs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var testnet types.Testnet
		fill_err = tp.Fill(&testnet)
		if fill_err != nil {
			return
		}
		var self *e2e.Node
		fill_err = tp.Fill(&self)
		if fill_err != nil {
			return
		}
		if self == nil {
			return
		}

		privatePeerIDs(testnet, self)
	})
}

func Fuzz_Nosy_publicChains__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var manifest types.Manifest
		fill_err = tp.Fill(&manifest)
		if fill_err != nil {
			return
		}
		var cfg DefinitionConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}

		publicChains(manifest, cfg)
	})
}

// skipping Fuzz_Nosy_random__ because parameters include func, chan, or unsupported interface: []T

func Fuzz_Nosy_runsDir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, manifestFile string) {
		runsDir(manifestFile)
	})
}

func Fuzz_Nosy_startAddingMockPortals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var def Definition
		fill_err = tp.Fill(&def)
		if fill_err != nil {
			return
		}

		startAddingMockPortals(ctx, def)
	})
}

func Fuzz_Nosy_sum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batches []int
		fill_err = tp.Fill(&batches)
		if fill_err != nil {
			return
		}

		sum(batches)
	})
}

func Fuzz_Nosy_toPortalDepls__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var def Definition
		fill_err = tp.Fill(&def)
		if fill_err != nil {
			return
		}
		var chains []types.EVMChain
		fill_err = tp.Fill(&chains)
		if fill_err != nil {
			return
		}

		toPortalDepls(def, chains)
	})
}

func Fuzz_Nosy_toPortalValidators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var validators map[*e2e.Node]int64
		fill_err = tp.Fill(&validators)
		if fill_err != nil {
			return
		}

		toPortalValidators(validators)
	})
}

func Fuzz_Nosy_waitForAllNodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var testnet *e2e.Testnet
		fill_err = tp.Fill(&testnet)
		if fill_err != nil {
			return
		}
		var height int64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}
		if testnet == nil {
			return
		}

		waitForAllNodes(ctx, testnet, height, timeout)
	})
}

func Fuzz_Nosy_waitForHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var testnet *e2e.Testnet
		fill_err = tp.Fill(&testnet)
		if fill_err != nil {
			return
		}
		var height int64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if testnet == nil {
			return
		}

		waitForHeight(ctx, testnet, height)
	})
}
