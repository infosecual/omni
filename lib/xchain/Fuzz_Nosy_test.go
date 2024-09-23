package xchain

import (
	"testing"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/spf13/pflag"
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

func Fuzz_Nosy_Attestation_AttestationRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a Attestation
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}

		a.AttestationRoot()
	})
}

func Fuzz_Nosy_Block_ShouldAttest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Block
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var attestInterval uint64
		fill_err = tp.Fill(&attestInterval)
		if fill_err != nil {
			return
		}

		b.ShouldAttest(attestInterval)
	})
}

func Fuzz_Nosy_ConfLevel_IsFuzzy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c ConfLevel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.IsFuzzy()
	})
}

func Fuzz_Nosy_ConfLevel_Label__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c ConfLevel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.Label()
	})
}

func Fuzz_Nosy_ConfLevel_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i ConfLevel
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		i.String()
	})
}

func Fuzz_Nosy_ConfLevel_Valid__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c ConfLevel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.Valid()
	})
}

func Fuzz_Nosy_EmitRef_Valid__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var level ConfLevel
		fill_err = tp.Fill(&level)
		if fill_err != nil {
			return
		}

		r := ConfEmitRef(level)
		r.Valid()
	})
}

func Fuzz_Nosy_MsgTree_MsgRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msgs []Msg
		fill_err = tp.Fill(&msgs)
		if fill_err != nil {
			return
		}

		t1, err := NewMsgTree(msgs)
		if err != nil {
			return
		}
		t1.MsgRoot()
	})
}

func Fuzz_Nosy_MsgTree_Proof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 []Msg
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var m2 []Msg
		fill_err = tp.Fill(&m2)
		if fill_err != nil {
			return
		}

		t1, err := NewMsgTree(m1)
		if err != nil {
			return
		}
		t1.Proof(m2)
	})
}

func Fuzz_Nosy_MsgTree_leafIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msgs []Msg
		fill_err = tp.Fill(&msgs)
		if fill_err != nil {
			return
		}
		var leaf [32]byte
		fill_err = tp.Fill(&leaf)
		if fill_err != nil {
			return
		}

		t1, err := NewMsgTree(msgs)
		if err != nil {
			return
		}
		t1.leafIndex(leaf)
	})
}

// skipping Fuzz_Nosy_Provider_ChainVersionHeight__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/xchain.Provider

// skipping Fuzz_Nosy_Provider_GetBlock__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/xchain.Provider

// skipping Fuzz_Nosy_Provider_GetEmittedCursor__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/xchain.Provider

// skipping Fuzz_Nosy_Provider_GetSubmission__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/xchain.Provider

// skipping Fuzz_Nosy_Provider_GetSubmittedCursor__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/xchain.Provider

// skipping Fuzz_Nosy_Provider_StreamAsync__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/xchain.Provider

// skipping Fuzz_Nosy_Provider_StreamBlocks__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/xchain.Provider

func Fuzz_Nosy_ProviderRequest_ChainVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r ProviderRequest
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}

		r.ChainVersion()
	})
}

func Fuzz_Nosy_RPCEndpoints_ByNameOrID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e RPCEndpoints
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var chainID uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}

		e.ByNameOrID(name, chainID)
	})
}

func Fuzz_Nosy_RPCEndpoints_Keys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e RPCEndpoints
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}

		e.Keys()
	})
}

func Fuzz_Nosy_ShardID_Broadcast__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s ShardID
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.Broadcast()
	})
}

func Fuzz_Nosy_ShardID_ConfLevel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s ShardID
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.ConfLevel()
	})
}

func Fuzz_Nosy_ShardID_Flags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s ShardID
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.Flags()
	})
}

func Fuzz_Nosy_ShardID_Label__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s ShardID
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.Label()
	})
}

func Fuzz_Nosy_StreamID_ChainVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s StreamID
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.ChainVersion()
	})
}

func Fuzz_Nosy_StreamID_ConfLevel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s StreamID
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.ConfLevel()
	})
}

func Fuzz_Nosy_BindFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var flags *pflag.FlagSet
		fill_err = tp.Fill(&flags)
		if fill_err != nil {
			return
		}
		var endpoints *RPCEndpoints
		fill_err = tp.Fill(&endpoints)
		if fill_err != nil {
			return
		}
		if flags == nil || endpoints == nil {
			return
		}

		BindFlags(flags, endpoints)
	})
}

func Fuzz_Nosy_EncodeXSubmit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sub bindings.XTypesSubmission
		fill_err = tp.Fill(&sub)
		if fill_err != nil {
			return
		}

		EncodeXSubmit(sub)
	})
}

func Fuzz_Nosy_encodeMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg Msg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		encodeMsg(msg)
	})
}

func Fuzz_Nosy_encodeSubmissionHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var attHeader AttestHeader
		fill_err = tp.Fill(&attHeader)
		if fill_err != nil {
			return
		}
		var blockHeader BlockHeader
		fill_err = tp.Fill(&blockHeader)
		if fill_err != nil {
			return
		}

		encodeSubmissionHeader(attHeader, blockHeader)
	})
}

func Fuzz_Nosy_msgLeaf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg Msg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msgLeaf(msg)
	})
}

func Fuzz_Nosy_submissionHeaderLeaf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var attHeader AttestHeader
		fill_err = tp.Fill(&attHeader)
		if fill_err != nil {
			return
		}
		var blockHeader BlockHeader
		fill_err = tp.Fill(&blockHeader)
		if fill_err != nil {
			return
		}

		submissionHeaderLeaf(attHeader, blockHeader)
	})
}
