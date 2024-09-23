package evmslashing

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/octane/evmengine/types"
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

func Fuzz_Nosy_EventProcessor_Addresses__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p EventProcessor
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		p.Addresses()
	})
}

func Fuzz_Nosy_EventProcessor_Deliver__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p EventProcessor
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var _x3 common.Hash
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		var elog *types.EVMEvent
		fill_err = tp.Fill(&elog)
		if fill_err != nil {
			return
		}
		if elog == nil {
			return
		}

		p.Deliver(ctx, _x3, elog)
	})
}

func Fuzz_Nosy_EventProcessor_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 EventProcessor
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_EventProcessor_Prepare__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p EventProcessor
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}

		p.Prepare(ctx, blockHash)
	})
}

func Fuzz_Nosy_EventProcessor_deliverUnjail__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p EventProcessor
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var ev *bindings.SlashingUnjail
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		if ev == nil {
			return
		}

		p.deliverUnjail(ctx, ev)
	})
}
