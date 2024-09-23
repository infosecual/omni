package xfeemngr

import (
	"context"
	"math/big"
	"testing"

	"github.com/omni-network/omni/lib/evmchain"
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

func Fuzz_Nosy_Manager_start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Manager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.start(ctx)
	})
}

func Fuzz_Nosy_feeOracle_correctPostsTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o feeOracle
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var dest evmchain.Metadata
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}

		o.correctPostsTo(ctx, dest)
	})
}

// skipping Fuzz_Nosy_feeOracle_syncForever__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xfeemngr/ticker.Ticker

func Fuzz_Nosy_feeOracle_syncGasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o feeOracle
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var dest evmchain.Metadata
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}

		o.syncGasPrice(ctx, dest)
	})
}

func Fuzz_Nosy_feeOracle_syncOnce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o feeOracle
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		o.syncOnce(ctx)
	})
}

func Fuzz_Nosy_feeOracle_syncToNativeRate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o feeOracle
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var dest evmchain.Metadata
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}

		o.syncToNativeRate(ctx, dest)
	})
}

func Fuzz_Nosy_chainsToSync__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var network netconf.Network
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}

		chainsToSync(network)
	})
}

func Fuzz_Nosy_guageGasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var src evmchain.Metadata
		fill_err = tp.Fill(&src)
		if fill_err != nil {
			return
		}
		var dest evmchain.Metadata
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var price uint64
		fill_err = tp.Fill(&price)
		if fill_err != nil {
			return
		}

		guageGasPrice(src, dest, price)
	})
}

func Fuzz_Nosy_guageRate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var src evmchain.Metadata
		fill_err = tp.Fill(&src)
		if fill_err != nil {
			return
		}
		var dest evmchain.Metadata
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var rate float64
		fill_err = tp.Fill(&rate)
		if fill_err != nil {
			return
		}

		guageRate(src, dest, rate)
	})
}

func Fuzz_Nosy_inEpsilon__(f *testing.F) {
	f.Fuzz(func(t *testing.T, a float64, b float64, epsilon float64) {
		inEpsilon(a, b, epsilon)
	})
}

func Fuzz_Nosy_makeEthClients__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chains []evmchain.Metadata
		fill_err = tp.Fill(&chains)
		if fill_err != nil {
			return
		}
		var rpcs xchain.RPCEndpoints
		fill_err = tp.Fill(&rpcs)
		if fill_err != nil {
			return
		}

		makeEthClients(chains, rpcs)
	})
}

// skipping Fuzz_Nosy_makeGasPricers__ because parameters include func, chan, or unsupported interface: map[uint64]github.com/omni-network/omni/lib/ethclient.Client

// skipping Fuzz_Nosy_makeOracles__ because parameters include func, chan, or unsupported interface: map[uint64]github.com/omni-network/omni/lib/ethclient.Client

func Fuzz_Nosy_numeratorToRate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *big.Int
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		numeratorToRate(n)
	})
}

// skipping Fuzz_Nosy_startMonitoring__ because parameters include func, chan, or unsupported interface: map[uint64]github.com/omni-network/omni/lib/ethclient.Client

func Fuzz_Nosy_withGasPriceShield__(f *testing.F) {
	f.Fuzz(func(t *testing.T, gasPrice uint64) {
		withGasPriceShield(gasPrice)
	})
}
