package testutil

import (
	context "context"
	"testing"

	"github.com/omni-network/omni/halo/valsync/types"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	gomock "go.uber.org/mock/gomock"
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

func Fuzz_Nosy_MockAttestKeeper_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockAttestKeeper(ctrl)
		m.EXPECT()
	})
}

func Fuzz_Nosy_MockAttestKeeper_ListAttestationsFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
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
		var confLevel uint32
		fill_err = tp.Fill(&confLevel)
		if fill_err != nil {
			return
		}
		var offset uint64
		fill_err = tp.Fill(&offset)
		if fill_err != nil {
			return
		}
		var max uint64
		fill_err = tp.Fill(&max)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockAttestKeeper(ctrl)
		m.ListAttestationsFrom(ctx, chainID, confLevel, offset, max)
	})
}

// skipping Fuzz_Nosy_MockAttestKeeperMockRecorder_ListAttestationsFrom__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_MockStakingKeeper_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockStakingKeeper(ctrl)
		m.EXPECT()
	})
}

func Fuzz_Nosy_MockStakingKeeper_EndBlocker__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockStakingKeeper(ctrl)
		m.EndBlocker(ctx)
	})
}

func Fuzz_Nosy_MockStakingKeeper_GetLastValidators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockStakingKeeper(ctrl)
		m.GetLastValidators(ctx)
	})
}

// skipping Fuzz_Nosy_MockStakingKeeperMockRecorder_EndBlocker__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockStakingKeeperMockRecorder_GetLastValidators__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_MockSubscriber_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockSubscriber(ctrl)
		m.EXPECT()
	})
}

func Fuzz_Nosy_MockSubscriber_UpdateValidatorSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var arg0 *types.ValidatorSetResponse
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if ctrl == nil || arg0 == nil {
			return
		}

		m := NewMockSubscriber(ctrl)
		m.UpdateValidatorSet(arg0)
	})
}

// skipping Fuzz_Nosy_MockSubscriberMockRecorder_UpdateValidatorSet__ because parameters include func, chan, or unsupported interface: any
