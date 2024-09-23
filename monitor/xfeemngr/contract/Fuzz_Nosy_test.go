package contract

import (
	"context"
	"math/big"
	"testing"

	"github.com/omni-network/omni/contracts/bindings"
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

func Fuzz_Nosy_MockFeeOracleV1_BulkSetFeeParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 context.Context
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var params []bindings.IFeeOracleV1ChainFeeParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}

		m := NewMockFeeOracleV1()
		m.BulkSetFeeParams(_x1, params)
	})
}

func Fuzz_Nosy_MockFeeOracleV1_GasPriceOn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 context.Context
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var destChainID uint64
		fill_err = tp.Fill(&destChainID)
		if fill_err != nil {
			return
		}

		m := NewMockFeeOracleV1()
		m.GasPriceOn(_x1, destChainID)
	})
}

func Fuzz_Nosy_MockFeeOracleV1_PostsTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 context.Context
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var destChainID uint64
		fill_err = tp.Fill(&destChainID)
		if fill_err != nil {
			return
		}

		m := NewMockFeeOracleV1()
		m.PostsTo(_x1, destChainID)
	})
}

func Fuzz_Nosy_MockFeeOracleV1_SetGasPriceOn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 context.Context
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var destChainID uint64
		fill_err = tp.Fill(&destChainID)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		if gasPrice == nil {
			return
		}

		m := NewMockFeeOracleV1()
		m.SetGasPriceOn(_x1, destChainID, gasPrice)
	})
}

func Fuzz_Nosy_MockFeeOracleV1_SetToNativeRate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 context.Context
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var destChainID uint64
		fill_err = tp.Fill(&destChainID)
		if fill_err != nil {
			return
		}
		var rate *big.Int
		fill_err = tp.Fill(&rate)
		if fill_err != nil {
			return
		}
		if rate == nil {
			return
		}

		m := NewMockFeeOracleV1()
		m.SetToNativeRate(_x1, destChainID, rate)
	})
}

func Fuzz_Nosy_MockFeeOracleV1_ToNativeRate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 context.Context
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var destChainID uint64
		fill_err = tp.Fill(&destChainID)
		if fill_err != nil {
			return
		}

		m := NewMockFeeOracleV1()
		m.ToNativeRate(_x1, destChainID)
	})
}

func Fuzz_Nosy_BoundFeeOracleV1_BulkSetFeeParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c BoundFeeOracleV1
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var params []bindings.IFeeOracleV1ChainFeeParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}

		c.BulkSetFeeParams(ctx, params)
	})
}

func Fuzz_Nosy_BoundFeeOracleV1_GasPriceOn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c BoundFeeOracleV1
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var destChainID uint64
		fill_err = tp.Fill(&destChainID)
		if fill_err != nil {
			return
		}

		c.GasPriceOn(ctx, destChainID)
	})
}

func Fuzz_Nosy_BoundFeeOracleV1_PostsTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c BoundFeeOracleV1
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var destChainID uint64
		fill_err = tp.Fill(&destChainID)
		if fill_err != nil {
			return
		}

		c.PostsTo(ctx, destChainID)
	})
}

func Fuzz_Nosy_BoundFeeOracleV1_SetGasPriceOn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c BoundFeeOracleV1
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var destChainID uint64
		fill_err = tp.Fill(&destChainID)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		if gasPrice == nil {
			return
		}

		c.SetGasPriceOn(ctx, destChainID, gasPrice)
	})
}

func Fuzz_Nosy_BoundFeeOracleV1_SetToNativeRate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c BoundFeeOracleV1
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var destChainID uint64
		fill_err = tp.Fill(&destChainID)
		if fill_err != nil {
			return
		}
		var rate *big.Int
		fill_err = tp.Fill(&rate)
		if fill_err != nil {
			return
		}
		if rate == nil {
			return
		}

		c.SetToNativeRate(ctx, destChainID, rate)
	})
}

func Fuzz_Nosy_BoundFeeOracleV1_ToNativeRate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c BoundFeeOracleV1
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var destChainID uint64
		fill_err = tp.Fill(&destChainID)
		if fill_err != nil {
			return
		}

		c.ToNativeRate(ctx, destChainID)
	})
}

func Fuzz_Nosy_BoundFeeOracleV1_txOptsWithCtx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c BoundFeeOracleV1
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		c.txOptsWithCtx(ctx)
	})
}

// skipping Fuzz_Nosy_FeeOracleV1_BulkSetFeeParams__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xfeemngr/contract.FeeOracleV1

// skipping Fuzz_Nosy_FeeOracleV1_GasPriceOn__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xfeemngr/contract.FeeOracleV1

// skipping Fuzz_Nosy_FeeOracleV1_PostsTo__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xfeemngr/contract.FeeOracleV1

// skipping Fuzz_Nosy_FeeOracleV1_SetGasPriceOn__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xfeemngr/contract.FeeOracleV1

// skipping Fuzz_Nosy_FeeOracleV1_SetToNativeRate__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xfeemngr/contract.FeeOracleV1

// skipping Fuzz_Nosy_FeeOracleV1_ToNativeRate__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xfeemngr/contract.FeeOracleV1
