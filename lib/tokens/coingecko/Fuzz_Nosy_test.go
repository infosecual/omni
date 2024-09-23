package coingecko

import (
	"context"
	"net/url"
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

func Fuzz_Nosy_Client_Price__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var tkns []tokens.Token
		fill_err = tp.Fill(&tkns)
		if fill_err != nil {
			return
		}

		c.Price(ctx, tkns...)
	})
}

// skipping Fuzz_Nosy_Client_doReq__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_Client_getPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var currency string
		fill_err = tp.Fill(&currency)
		if fill_err != nil {
			return
		}
		var tkns []tokens.Token
		fill_err = tp.Fill(&tkns)
		if fill_err != nil {
			return
		}

		c.getPrice(ctx, currency, tkns...)
	})
}

func Fuzz_Nosy_Client_uri__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var params url.Values
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}

		c.uri(path, params)
	})
}

func Fuzz_Nosy_WithHost__(f *testing.F) {
	f.Fuzz(func(t *testing.T, host string) {
		WithHost(host)
	})
}
