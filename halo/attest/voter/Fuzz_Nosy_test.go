package voter

import (
	"context"
	"testing"
	"time"

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

func Fuzz_Nosy_Voter_AvailableCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *Voter
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.AvailableCount()
	})
}

func Fuzz_Nosy_Voter_GetAvailable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *Voter
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.GetAvailable()
	})
}

func Fuzz_Nosy_Voter_LocalAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *Voter
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.LocalAddress()
	})
}

func Fuzz_Nosy_Voter_SetCommitted__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *Voter
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var headers []*types.AttestHeader
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.SetCommitted(headers)
	})
}

func Fuzz_Nosy_Voter_SetProposed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *Voter
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var headers []*types.AttestHeader
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.SetProposed(headers)
	})
}

func Fuzz_Nosy_Voter_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *Voter
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.Start(ctx)
	})
}

func Fuzz_Nosy_Voter_TrimBehind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *Voter
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var minsByChain map[xchain.ChainVersion]uint64
		fill_err = tp.Fill(&minsByChain)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.TrimBehind(minsByChain)
	})
}

func Fuzz_Nosy_Voter_UpdateValidatorSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *Voter
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var valset *types.ValidatorSetResponse
		fill_err = tp.Fill(&valset)
		if fill_err != nil {
			return
		}
		if v == nil || valset == nil {
			return
		}

		v.UpdateValidatorSet(valset)
	})
}

func Fuzz_Nosy_Voter_Vote__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *Voter
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var attHeader xchain.AttestHeader
		fill_err = tp.Fill(&attHeader)
		if fill_err != nil {
			return
		}
		var block xchain.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var allowSkip bool
		fill_err = tp.Fill(&allowSkip)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.Vote(attHeader, block, allowSkip)
	})
}

func Fuzz_Nosy_Voter_WaitDone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *Voter
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.WaitDone()
	})
}

func Fuzz_Nosy_Voter_availableAndProposedUnsafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *Voter
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.availableAndProposedUnsafe()
	})
}

func Fuzz_Nosy_Voter_getFromHeightAndOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *Voter
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainVer xchain.ChainVersion
		fill_err = tp.Fill(&chainVer)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.getFromHeightAndOffset(ctx, chainVer)
	})
}

func Fuzz_Nosy_Voter_instrumentUnsafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *Voter
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.instrumentUnsafe()
	})
}

func Fuzz_Nosy_Voter_isValidator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *Voter
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.isValidator()
	})
}

func Fuzz_Nosy_Voter_latestByChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *Voter
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var chainVer xchain.ChainVersion
		fill_err = tp.Fill(&chainVer)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.latestByChain(chainVer)
	})
}

func Fuzz_Nosy_Voter_minWindow__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *Voter
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var chainVer xchain.ChainVersion
		fill_err = tp.Fill(&chainVer)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.minWindow(chainVer)
	})
}

func Fuzz_Nosy_Voter_runForever__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *Voter
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainVer xchain.ChainVersion
		fill_err = tp.Fill(&chainVer)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.runForever(ctx, chainVer)
	})
}

func Fuzz_Nosy_Voter_runOnce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *Voter
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainVer xchain.ChainVersion
		fill_err = tp.Fill(&chainVer)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.runOnce(ctx, chainVer)
	})
}

func Fuzz_Nosy_Voter_saveUnsafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *Voter
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.saveUnsafe()
	})
}

func Fuzz_Nosy_offsetTracker_NextAttestOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, nextAttestOffset uint64, blockHeight uint64) {
		c := newOffsetTracker(nextAttestOffset)
		c.NextAttestOffset(blockHeight)
	})
}

func Fuzz_Nosy_headerMap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var headers []*types.AttestHeader
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}

		headerMap(headers)
	})
}

func Fuzz_Nosy_latestFromJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var latest []*types.Vote
		fill_err = tp.Fill(&latest)
		if fill_err != nil {
			return
		}

		latestFromJSON(latest)
	})
}

func Fuzz_Nosy_latestMsgOffsets__(f *testing.F) {
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

		latestMsgOffsets(msgs)
	})
}

func Fuzz_Nosy_latestToJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var latest map[xchain.ChainVersion]*types.Vote
		fill_err = tp.Fill(&latest)
		if fill_err != nil {
			return
		}

		latestToJSON(latest)
	})
}

func Fuzz_Nosy_logVoteCreated__(f *testing.F) {
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
		var network netconf.Network
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var attHeader xchain.AttestHeader
		fill_err = tp.Fill(&attHeader)
		if fill_err != nil {
			return
		}
		var block xchain.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}

		logVoteCreated(ctx, network, attHeader, block)
	})
}

func Fuzz_Nosy_newDebugLogFilterer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var period time.Duration
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}

		newDebugLogFilterer(period)
	})
}

func Fuzz_Nosy_pruneLatestPerChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var atts []*types.Vote
		fill_err = tp.Fill(&atts)
		if fill_err != nil {
			return
		}

		pruneLatestPerChain(atts)
	})
}
