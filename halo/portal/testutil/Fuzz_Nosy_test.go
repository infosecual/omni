package testutil

import (
	"testing"

	xchain "github.com/omni-network/omni/lib/xchain"
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

func Fuzz_Nosy_MockPortal_EXPECT__(f *testing.F) {
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

		m := NewMockPortal(ctrl)
		m.EXPECT()
	})
}

func Fuzz_Nosy_MockPortal_EmitMsg__(f *testing.F) {
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
		var ctx types.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var typ types.MsgType
		fill_err = tp.Fill(&typ)
		if fill_err != nil {
			return
		}
		var msgTypeID uint64
		fill_err = tp.Fill(&msgTypeID)
		if fill_err != nil {
			return
		}
		var destChainID uint64
		fill_err = tp.Fill(&destChainID)
		if fill_err != nil {
			return
		}
		var shardID xchain.ShardID
		fill_err = tp.Fill(&shardID)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockPortal(ctrl)
		m.EmitMsg(ctx, typ, msgTypeID, destChainID, shardID)
	})
}

// skipping Fuzz_Nosy_MockPortalMockRecorder_EmitMsg__ because parameters include func, chan, or unsupported interface: any
