package relayer

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
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

func Fuzz_Nosy_Worker_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *Worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.Run(ctx)
	})
}

// skipping Fuzz_Nosy_Worker_newCallback__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/relayer/app.SendFunc

func Fuzz_Nosy_Worker_runOnce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *Worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.runOnce(ctx)
	})
}

func Fuzz_Nosy_activeBuffer_AddInput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *activeBuffer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var submission xchain.Submission
		fill_err = tp.Fill(&submission)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.AddInput(ctx, submission)
	})
}

func Fuzz_Nosy_activeBuffer_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *activeBuffer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Run(ctx)
	})
}

// skipping Fuzz_Nosy_activeBuffer_submitErr__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_msgCursorFilter_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cursors []xchain.SubmitCursor
		fill_err = tp.Fill(&cursors)
		if fill_err != nil {
			return
		}
		var stream xchain.StreamID
		fill_err = tp.Fill(&stream)
		if fill_err != nil {
			return
		}
		var msgOffset uint64
		fill_err = tp.Fill(&msgOffset)
		if fill_err != nil {
			return
		}

		f1, err := newMsgOffsetFilter(cursors)
		if err != nil {
			return
		}
		f1.Check(stream, msgOffset)
	})
}

func Fuzz_Nosy_Sender_SendTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s Sender
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var sub xchain.Submission
		fill_err = tp.Fill(&sub)
		if fill_err != nil {
			return
		}

		s.SendTransaction(ctx, sub)
	})
}

func Fuzz_Nosy_CreateSubmissions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var up StreamUpdate
		fill_err = tp.Fill(&up)
		if fill_err != nil {
			return
		}

		CreateSubmissions(up)
	})
}

func Fuzz_Nosy_attestationForShard__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var att xchain.Attestation
		fill_err = tp.Fill(&att)
		if fill_err != nil {
			return
		}
		var shard xchain.ShardID
		fill_err = tp.Fill(&shard)
		if fill_err != nil {
			return
		}

		attestationForShard(att, shard)
	})
}

// skipping Fuzz_Nosy_fetchXBlock__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/xchain.Provider

// skipping Fuzz_Nosy_filterMsgs__ because parameters include func, chan, or unsupported interface: func(github.com/omni-network/omni/lib/xchain.StreamID) string

func Fuzz_Nosy_fromChainVersionOffsets__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cursors []xchain.SubmitCursor
		fill_err = tp.Fill(&cursors)
		if fill_err != nil {
			return
		}
		var chainVers []xchain.ChainVersion
		fill_err = tp.Fill(&chainVers)
		if fill_err != nil {
			return
		}

		fromChainVersionOffsets(cursors, chainVers)
	})
}

// skipping Fuzz_Nosy_getSubmittedCursors__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/xchain.Provider

func Fuzz_Nosy_groupMsgsByCost__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msgs []xchain.Msg
		fill_err = tp.Fill(&msgs)
		if fill_err != nil {
			return
		}

		groupMsgsByCost(msgs)
	})
}

func Fuzz_Nosy_initializeRPCClients__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chains []netconf.Chain
		fill_err = tp.Fill(&chains)
		if fill_err != nil {
			return
		}
		var endpoints xchain.RPCEndpoints
		fill_err = tp.Fill(&endpoints)
		if fill_err != nil {
			return
		}

		initializeRPCClients(chains, endpoints)
	})
}

func Fuzz_Nosy_naiveSubmissionGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msgs []xchain.Msg
		fill_err = tp.Fill(&msgs)
		if fill_err != nil {
			return
		}

		naiveSubmissionGas(msgs)
	})
}

func Fuzz_Nosy_serveMonitoring__(f *testing.F) {
	f.Fuzz(func(t *testing.T, address string) {
		serveMonitoring(address)
	})
}
