package tokenprice

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/tokens"
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

func Fuzz_Nosy_Buffer_Price__(f *testing.F) {
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
		var token tokens.Token
		fill_err = tp.Fill(&token)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Price(token)
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
		var token tokens.Token
		fill_err = tp.Fill(&token)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.price(token)
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
		var token tokens.Token
		fill_err = tp.Fill(&token)
		if fill_err != nil {
			return
		}
		var price float64
		fill_err = tp.Fill(&price)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.setPrice(token, price)
	})
}

func Fuzz_Nosy_Buffer_stream__(f *testing.F) {
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

		b.stream(ctx)
	})
}

func Fuzz_Nosy_WithThresholdPct__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pct float64) {
		WithThresholdPct(pct)
	})
}

// skipping Fuzz_Nosy_WithTicker__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xfeemngr/ticker.Ticker

func Fuzz_Nosy_guageBuffered__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var prices map[tokens.Token]float64
		fill_err = tp.Fill(&prices)
		if fill_err != nil {
			return
		}

		guageBuffered(prices)
	})
}

func Fuzz_Nosy_guageLive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var prices map[tokens.Token]float64
		fill_err = tp.Fill(&prices)
		if fill_err != nil {
			return
		}

		guageLive(prices)
	})
}

func Fuzz_Nosy_inThreshold__(f *testing.F) {
	f.Fuzz(func(t *testing.T, a float64, b float64, pct float64) {
		inThreshold(a, b, pct)
	})
}
