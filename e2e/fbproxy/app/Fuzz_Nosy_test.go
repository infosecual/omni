package app

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

func Fuzz_Nosy_TransactionArgs_ToTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var args *TransactionArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		if args == nil {
			return
		}

		args.ToTransaction()
	})
}

func Fuzz_Nosy_TransactionArgs_data__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var args *TransactionArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		if args == nil {
			return
		}

		args.data()
	})
}

// skipping Fuzz_Nosy_proxy_Proxy__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Nosy_proxy_proxy__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

func Fuzz_Nosy_closeReader_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 closeReader
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Close()
	})
}

// skipping Fuzz_Nosy_txSigner_Sign__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/e2e/fbproxy/app.txSigner

func Fuzz_Nosy_Start__(f *testing.F) {
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
		var cfg Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}

		Start(ctx, cfg)
	})
}

func Fuzz_Nosy_getChainID__(f *testing.F) {
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
		var rpc string
		fill_err = tp.Fill(&rpc)
		if fill_err != nil {
			return
		}

		getChainID(ctx, rpc)
	})
}
