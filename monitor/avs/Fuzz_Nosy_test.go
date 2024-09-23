package avs

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

// skipping Fuzz_Nosy_monitorForever__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/avs.monitorOnce

func Fuzz_Nosy_startMonitoring__(f *testing.F) {
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
		var a2 *bindings.OmniAVS
		fill_err = tp.Fill(&a2)
		if fill_err != nil {
			return
		}
		if a2 == nil {
			return
		}

		startMonitoring(ctx, a2)
	})
}

func Fuzz_Nosy_weiToEth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var wei *big.Int
		fill_err = tp.Fill(&wei)
		if fill_err != nil {
			return
		}
		if wei == nil {
			return
		}

		weiToEth(wei)
	})
}
