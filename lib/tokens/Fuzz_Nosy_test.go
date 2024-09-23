package tokens

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

func Fuzz_Nosy_CachedPricer_ClearCache__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *CachedPricer
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.ClearCache()
	})
}

func Fuzz_Nosy_CachedPricer_Price__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *CachedPricer
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var t3 []Token
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Price(ctx, t3...)
	})
}

func Fuzz_Nosy_MockPricer_Price__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var prices map[Token]float64
		fill_err = tp.Fill(&prices)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var tkns []Token
		fill_err = tp.Fill(&tkns)
		if fill_err != nil {
			return
		}

		m := NewMockPricer(prices)
		m.Price(_x2, tkns...)
	})
}

func Fuzz_Nosy_MockPricer_SetPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var prices map[Token]float64
		fill_err = tp.Fill(&prices)
		if fill_err != nil {
			return
		}
		var token Token
		fill_err = tp.Fill(&token)
		if fill_err != nil {
			return
		}
		var price float64
		fill_err = tp.Fill(&price)
		if fill_err != nil {
			return
		}

		m := NewMockPricer(prices)
		m.SetPrice(token, price)
	})
}

// skipping Fuzz_Nosy_Pricer_Price__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/tokens.Pricer

func Fuzz_Nosy_Token_CoingeckoID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id string) {
		t1 := MustFromCoingeckoID(id)
		t1.CoingeckoID()
	})
}

func Fuzz_Nosy_Token_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id string) {
		t1 := MustFromCoingeckoID(id)
		t1.String()
	})
}

func Fuzz_Nosy_FromCoingeckoID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id string) {
		FromCoingeckoID(id)
	})
}
