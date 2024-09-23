package testutil

import (
	context "context"
	"testing"

	"github.com/omni-network/omni/lib/xchain"
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

func Fuzz_Nosy_MockChainNamer_ChainName__(f *testing.F) {
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
		var chainVer xchain.ChainVersion
		fill_err = tp.Fill(&chainVer)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockChainNamer(ctrl)
		m.ChainName(chainVer)
	})
}

func Fuzz_Nosy_MockChainNamer_EXPECT__(f *testing.F) {
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

		m := NewMockChainNamer(ctrl)
		m.EXPECT()
	})
}

// skipping Fuzz_Nosy_MockChainNamerMockRecorder_ChainName__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_MockRegistry_ConfLevels__(f *testing.F) {
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

		m := NewMockRegistry(ctrl)
		m.ConfLevels(ctx)
	})
}

func Fuzz_Nosy_MockRegistry_EXPECT__(f *testing.F) {
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

		m := NewMockRegistry(ctrl)
		m.EXPECT()
	})
}

// skipping Fuzz_Nosy_MockRegistryMockRecorder_ConfLevels__ because parameters include func, chan, or unsupported interface: any

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

func Fuzz_Nosy_MockStakingKeeper_GetPubKeyByConsAddr__(f *testing.F) {
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
		var arg0 context.Context
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 types.ConsAddress
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockStakingKeeper(ctrl)
		m.GetPubKeyByConsAddr(arg0, arg1)
	})
}

// skipping Fuzz_Nosy_MockStakingKeeperMockRecorder_GetPubKeyByConsAddr__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_MockValProvider_ActiveSetByHeight__(f *testing.F) {
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
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockValProvider(ctrl)
		m.ActiveSetByHeight(ctx, height)
	})
}

func Fuzz_Nosy_MockValProvider_EXPECT__(f *testing.F) {
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

		m := NewMockValProvider(ctrl)
		m.EXPECT()
	})
}

// skipping Fuzz_Nosy_MockValProviderMockRecorder_ActiveSetByHeight__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_MockVoter_EXPECT__(f *testing.F) {
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

		m := NewMockVoter(ctrl)
		m.EXPECT()
	})
}

func Fuzz_Nosy_MockVoter_GetAvailable__(f *testing.F) {
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

		m := NewMockVoter(ctrl)
		m.GetAvailable()
	})
}

func Fuzz_Nosy_MockVoter_LocalAddress__(f *testing.F) {
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

		m := NewMockVoter(ctrl)
		m.LocalAddress()
	})
}

func Fuzz_Nosy_MockVoter_SetCommitted__(f *testing.F) {
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
		var headers []*types.AttestHeader
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockVoter(ctrl)
		m.SetCommitted(headers)
	})
}

func Fuzz_Nosy_MockVoter_SetProposed__(f *testing.F) {
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
		var headers []*types.AttestHeader
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockVoter(ctrl)
		m.SetProposed(headers)
	})
}

func Fuzz_Nosy_MockVoter_TrimBehind__(f *testing.F) {
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
		var minsByChain map[xchain.ChainVersion]uint64
		fill_err = tp.Fill(&minsByChain)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockVoter(ctrl)
		m.TrimBehind(minsByChain)
	})
}

func Fuzz_Nosy_MockVoter_UpdateValidatorSet__(f *testing.F) {
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
		var set *types.ValidatorSetResponse
		fill_err = tp.Fill(&set)
		if fill_err != nil {
			return
		}
		if ctrl == nil || set == nil {
			return
		}

		m := NewMockVoter(ctrl)
		m.UpdateValidatorSet(set)
	})
}

func Fuzz_Nosy_MockVoterMockRecorder_GetAvailable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mr *MockVoterMockRecorder
		fill_err = tp.Fill(&mr)
		if fill_err != nil {
			return
		}
		if mr == nil {
			return
		}

		mr.GetAvailable()
	})
}

func Fuzz_Nosy_MockVoterMockRecorder_LocalAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mr *MockVoterMockRecorder
		fill_err = tp.Fill(&mr)
		if fill_err != nil {
			return
		}
		if mr == nil {
			return
		}

		mr.LocalAddress()
	})
}

// skipping Fuzz_Nosy_MockVoterMockRecorder_SetCommitted__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockVoterMockRecorder_SetProposed__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockVoterMockRecorder_TrimBehind__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockVoterMockRecorder_UpdateValidatorSet__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_ChainNamer_ChainName__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/testutil.ChainNamer
