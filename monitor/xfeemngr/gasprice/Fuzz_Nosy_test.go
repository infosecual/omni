package gasprice

import (
	"context"
	"testing"

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

func Fuzz_Nosy_Buffer_GasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Buffer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var chainID uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GasPrice(chainID)
	})
}

func Fuzz_Nosy_Buffer_Stream__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Buffer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Stream(ctx)
	})
}

func Fuzz_Nosy_Buffer_price__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Buffer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var chainID uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.price(chainID)
	})
}

func Fuzz_Nosy_Buffer_setPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Buffer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var chainID uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var price uint64
		fill_err = tp.Fill(&price)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.setPrice(chainID, price)
	})
}

func Fuzz_Nosy_Buffer_streamAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Buffer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.streamAll(ctx)
	})
}

func Fuzz_Nosy_Buffer_streamOne__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Buffer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.streamOne(ctx, chainID)
	})
}

func Fuzz_Nosy_MockPricer_Price__(f *testing.F) {
	f.Fuzz(func(t *testing.T, price uint64) {
		m := NewMockPricer(price)
		m.Price()
	})
}

func Fuzz_Nosy_MockPricer_SetPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, p1 uint64, p2 uint64) {
		m := NewMockPricer(p1)
		m.SetPrice(p2)
	})
}

func Fuzz_Nosy_MockPricer_SuggestGasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var price uint64
		fill_err = tp.Fill(&price)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}

		m := NewMockPricer(price)
		m.SuggestGasPrice(_x2)
	})
}

func Fuzz_Nosy_WithThresholdPct__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pct float64) {
		WithThresholdPct(pct)
	})
}

// skipping Fuzz_Nosy_WithTicker__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xfeemngr/ticker.Ticker

func Fuzz_Nosy_chainName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64) {
		chainName(chainID)
	})
}

func Fuzz_Nosy_guageBuffered__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64, price uint64) {
		guageBuffered(chainID, price)
	})
}

func Fuzz_Nosy_guageLive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64, price uint64) {
		guageLive(chainID, price)
	})
}

func Fuzz_Nosy_inThreshold__(f *testing.F) {
	f.Fuzz(func(t *testing.T, a uint64, b uint64, pct float64) {
		inThreshold(a, b, pct)
	})
}
