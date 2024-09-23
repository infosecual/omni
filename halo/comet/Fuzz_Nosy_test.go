package comet

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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

// skipping Fuzz_Nosy_API_IsValidator__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/comet.API

// skipping Fuzz_Nosy_API_Validators__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/comet.API

func Fuzz_Nosy_adapter_IsValidator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a adapter
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var valAddress common.Address
		fill_err = tp.Fill(&valAddress)
		if fill_err != nil {
			return
		}

		a.IsValidator(ctx, valAddress)
	})
}

func Fuzz_Nosy_adapter_Validators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a adapter
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var height int64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}

		a.Validators(ctx, height)
	})
}
