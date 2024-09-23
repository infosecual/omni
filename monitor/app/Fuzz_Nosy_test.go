package monitor

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
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

func Fuzz_Nosy_calcMaxDiff__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var heights []int64
		fill_err = tp.Fill(&heights)
		if fill_err != nil {
			return
		}

		calcMaxDiff(heights)
	})
}

// skipping Fuzz_Nosy_concurrentLatestHeights__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/ethclient.Client

func Fuzz_Nosy_initializeEthClients__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chains []netconf.Chain
		fill_err = tp.Fill(&chains)
		if fill_err != nil {
			return
		}
		var endpoints xchain.RPCEndpoints
		fill_err = tp.Fill(&endpoints)
		if fill_err != nil {
			return
		}

		initializeEthClients(chains, endpoints)
	})
}

// skipping Fuzz_Nosy_keys__ because parameters include func, chan, or unsupported interface: map[K]V

// skipping Fuzz_Nosy_pick__ because parameters include func, chan, or unsupported interface: []V

// skipping Fuzz_Nosy_runHistoricalBaselineForever__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/cchain.Provider

// skipping Fuzz_Nosy_runHistoricalBaselineOnce__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/cchain.Provider

func Fuzz_Nosy_serveMonitoring__(f *testing.F) {
	f.Fuzz(func(t *testing.T, address string) {
		serveMonitoring(address)
	})
}

// skipping Fuzz_Nosy_startMonitoringSyncDiff__ because parameters include func, chan, or unsupported interface: map[uint64]github.com/omni-network/omni/lib/ethclient.Client

func Fuzz_Nosy_syncAVSForever__(f *testing.F) {
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
		var omniAVS *bindings.OmniAVS
		fill_err = tp.Fill(&omniAVS)
		if fill_err != nil {
			return
		}
		var backend *ethbackend.Backend
		fill_err = tp.Fill(&backend)
		if fill_err != nil {
			return
		}
		var from common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		if omniAVS == nil || backend == nil {
			return
		}

		syncAVSForever(ctx, omniAVS, backend, from)
	})
}
