package keeper

import (
	context "context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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

func Fuzz_Nosy_Attestation_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Attestation
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_Attestation_GetAttestOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Attestation
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetAttestOffset()
	})
}

func Fuzz_Nosy_Attestation_GetAttestationRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Attestation
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetAttestationRoot()
	})
}

func Fuzz_Nosy_Attestation_GetBlockHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Attestation
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetBlockHash()
	})
}

func Fuzz_Nosy_Attestation_GetBlockHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Attestation
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetBlockHeight()
	})
}

func Fuzz_Nosy_Attestation_GetChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Attestation
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetChainId()
	})
}

func Fuzz_Nosy_Attestation_GetConfLevel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Attestation
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetConfLevel()
	})
}

func Fuzz_Nosy_Attestation_GetCreatedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Attestation
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetCreatedHeight()
	})
}

func Fuzz_Nosy_Attestation_GetFinalizedAttId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Attestation
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetFinalizedAttId()
	})
}

func Fuzz_Nosy_Attestation_GetId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Attestation
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetId()
	})
}

func Fuzz_Nosy_Attestation_GetMsgRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Attestation
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetMsgRoot()
	})
}

func Fuzz_Nosy_Attestation_GetStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Attestation
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetStatus()
	})
}

func Fuzz_Nosy_Attestation_GetValidatorSetId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Attestation
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetValidatorSetId()
	})
}

func Fuzz_Nosy_Attestation_IsFuzzy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *Attestation
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.IsFuzzy()
	})
}

func Fuzz_Nosy_Attestation_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Attestation
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_Attestation_ProtoReflect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Attestation
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.ProtoReflect()
	})
}

func Fuzz_Nosy_Attestation_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Attestation
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.Reset()
	})
}

func Fuzz_Nosy_Attestation_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Attestation
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.String()
	})
}

func Fuzz_Nosy_Attestation_XChainVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *Attestation
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.XChainVersion()
	})
}

func Fuzz_Nosy_Keeper_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg *types.MsgAddVotes
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if k == nil || msg == nil {
			return
		}

		k.Add(ctx, msg)
	})
}

func Fuzz_Nosy_Keeper_Approve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var valset ValSet
		fill_err = tp.Fill(&valset)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		k.Approve(ctx, valset)
	})
}

func Fuzz_Nosy_Keeper_AttestationsFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *types.AttestationsFromRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if k == nil || req == nil {
			return
		}

		k.AttestationsFrom(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_BeginBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		k.BeginBlock(ctx)
	})
}

func Fuzz_Nosy_Keeper_EarliestAttestation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *types.EarliestAttestationRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if k == nil || req == nil {
			return
		}

		k.EarliestAttestation(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_EndBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		k.EndBlock(ctx)
	})
}

func Fuzz_Nosy_Keeper_ExtendVote__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx types.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var _x3 *types.RequestExtendVote
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if k == nil || _x3 == nil {
			return
		}

		k.ExtendVote(ctx, _x3)
	})
}

func Fuzz_Nosy_Keeper_LatestAttestation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *types.LatestAttestationRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if k == nil || req == nil {
			return
		}

		k.LatestAttestation(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_ListAllAttestations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *types.ListAllAttestationsRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if k == nil || req == nil {
			return
		}

		k.ListAllAttestations(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_ListAttestationsFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *Keeper
		fill_err = tp.Fill(&k)
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
		if k == nil {
			return
		}

		k.ListAttestationsFrom(ctx, chainID, confLevel, offset, max)
	})
}

func Fuzz_Nosy_Keeper_PrepareVotes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var commit types.ExtendedCommitInfo
		fill_err = tp.Fill(&commit)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		k.PrepareVotes(ctx, commit)
	})
}

// skipping Fuzz_Nosy_Keeper_RegisterProposalService__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/grpc.Server

// skipping Fuzz_Nosy_Keeper_SetPortalRegistry__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/registry/types.PortalRegistry

// skipping Fuzz_Nosy_Keeper_SetValidatorProvider__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/types.ValidatorProvider

func Fuzz_Nosy_Keeper_VerifyVoteExtension__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx types.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *types.RequestVerifyVoteExtension
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if k == nil || req == nil {
			return
		}

		k.VerifyVoteExtension(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_WindowCompare__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *types.WindowCompareRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if k == nil || req == nil {
			return
		}

		k.WindowCompare(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_addOne__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var agg *types.AggVote
		fill_err = tp.Fill(&agg)
		if fill_err != nil {
			return
		}
		var valSetID uint64
		fill_err = tp.Fill(&valSetID)
		if fill_err != nil {
			return
		}
		if k == nil || agg == nil {
			return
		}

		k.addOne(ctx, agg, valSetID)
	})
}

func Fuzz_Nosy_Keeper_deleteBefore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *Keeper
		fill_err = tp.Fill(&k)
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
		var consensusID uint64
		fill_err = tp.Fill(&consensusID)
		if fill_err != nil {
			return
		}
		var cHeight uint64
		fill_err = tp.Fill(&cHeight)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		k.deleteBefore(ctx, height, consensusID, cHeight)
	})
}

func Fuzz_Nosy_Keeper_earliestAttestation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var version xchain.ChainVersion
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		k.earliestAttestation(ctx, version)
	})
}

func Fuzz_Nosy_Keeper_formAttestationResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var att *Attestation
		fill_err = tp.Fill(&att)
		if fill_err != nil {
			return
		}
		if k == nil || att == nil {
			return
		}

		k.formAttestationResponse(ctx, att)
	})
}

func Fuzz_Nosy_Keeper_getSigTuples__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var attID uint64
		fill_err = tp.Fill(&attID)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		k.getSigTuples(ctx, attID)
	})
}

func Fuzz_Nosy_Keeper_getSigs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var attID uint64
		fill_err = tp.Fill(&attID)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		k.getSigs(ctx, attID)
	})
}

func Fuzz_Nosy_Keeper_getValEthAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cmtAddr []byte
		fill_err = tp.Fill(&cmtAddr)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		k.getValEthAddr(ctx, cmtAddr)
	})
}

func Fuzz_Nosy_Keeper_isDoubleSign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var attID uint64
		fill_err = tp.Fill(&attID)
		if fill_err != nil {
			return
		}
		var agg *types.AggVote
		fill_err = tp.Fill(&agg)
		if fill_err != nil {
			return
		}
		var sig *types.SigTuple
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if k == nil || agg == nil || sig == nil {
			return
		}

		k.isDoubleSign(ctx, attID, agg, sig)
	})
}

func Fuzz_Nosy_Keeper_latestAttestation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var version xchain.ChainVersion
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		k.latestAttestation(ctx, version)
	})
}

func Fuzz_Nosy_Keeper_listAllAttestations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var version xchain.ChainVersion
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		var status Status
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var attestOfset uint64
		fill_err = tp.Fill(&attestOfset)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		k.listAllAttestations(ctx, version, status, attestOfset)
	})
}

func Fuzz_Nosy_Keeper_maybeOverrideFinalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var att *Attestation
		fill_err = tp.Fill(&att)
		if fill_err != nil {
			return
		}
		if k == nil || att == nil {
			return
		}

		k.maybeOverrideFinalized(ctx, att)
	})
}

func Fuzz_Nosy_Keeper_prevBlockValSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		k.prevBlockValSet(ctx)
	})
}

// skipping Fuzz_Nosy_Keeper_verifyAggVotes__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.windowCompareFunc

func Fuzz_Nosy_Keeper_windowCompare__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *Keeper
		fill_err = tp.Fill(&k)
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
		var offset uint64
		fill_err = tp.Fill(&offset)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		k.windowCompare(ctx, chainVer, offset)
	})
}

func Fuzz_Nosy_Signature_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Signature
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_Signature_GetAttId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Signature
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetAttId()
	})
}

func Fuzz_Nosy_Signature_GetAttestOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Signature
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetAttestOffset()
	})
}

func Fuzz_Nosy_Signature_GetChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Signature
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetChainId()
	})
}

func Fuzz_Nosy_Signature_GetConfLevel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Signature
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetConfLevel()
	})
}

func Fuzz_Nosy_Signature_GetId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Signature
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetId()
	})
}

func Fuzz_Nosy_Signature_GetSignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Signature
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetSignature()
	})
}

func Fuzz_Nosy_Signature_GetValidatorAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Signature
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetValidatorAddress()
	})
}

func Fuzz_Nosy_Signature_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Signature
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_Signature_ProtoReflect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Signature
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.ProtoReflect()
	})
}

func Fuzz_Nosy_Signature_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Signature
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.Reset()
	})
}

func Fuzz_Nosy_Signature_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Signature
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.String()
	})
}

func Fuzz_Nosy_valAddrCache_GetEthAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *valAddrCache
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var cmtAddr [20]byte
		fill_err = tp.Fill(&cmtAddr)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.GetEthAddress(cmtAddr)
	})
}

func Fuzz_Nosy_valAddrCache_SetAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *valAddrCache
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var vals []*types.Validator
		fill_err = tp.Fill(&vals)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.SetAll(vals)
	})
}

func Fuzz_Nosy_AttestationAttestationRootIndexKey_WithAttestationRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this AttestationAttestationRootIndexKey
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var attestation_root []byte
		fill_err = tp.Fill(&attestation_root)
		if fill_err != nil {
			return
		}

		this.WithAttestationRoot(attestation_root)
	})
}

func Fuzz_Nosy_AttestationAttestationRootIndexKey_attestationIndexKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x AttestationAttestationRootIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.attestationIndexKey()
	})
}

func Fuzz_Nosy_AttestationAttestationRootIndexKey_id__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x AttestationAttestationRootIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.id()
	})
}

func Fuzz_Nosy_AttestationAttestationRootIndexKey_values__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x AttestationAttestationRootIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.values()
	})
}

func Fuzz_Nosy_AttestationCreatedHeightIndexKey_WithCreatedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this AttestationCreatedHeightIndexKey
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var created_height uint64
		fill_err = tp.Fill(&created_height)
		if fill_err != nil {
			return
		}

		this.WithCreatedHeight(created_height)
	})
}

func Fuzz_Nosy_AttestationCreatedHeightIndexKey_attestationIndexKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x AttestationCreatedHeightIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.attestationIndexKey()
	})
}

func Fuzz_Nosy_AttestationCreatedHeightIndexKey_id__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x AttestationCreatedHeightIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.id()
	})
}

func Fuzz_Nosy_AttestationCreatedHeightIndexKey_values__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x AttestationCreatedHeightIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.values()
	})
}

func Fuzz_Nosy_AttestationIdIndexKey_WithId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this AttestationIdIndexKey
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		this.WithId(id)
	})
}

func Fuzz_Nosy_AttestationIdIndexKey_attestationIndexKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x AttestationIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.attestationIndexKey()
	})
}

func Fuzz_Nosy_AttestationIdIndexKey_id__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x AttestationIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.id()
	})
}

func Fuzz_Nosy_AttestationIdIndexKey_values__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x AttestationIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.values()
	})
}

// skipping Fuzz_Nosy_AttestationIndexKey_attestationIndexKey__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.AttestationIndexKey

// skipping Fuzz_Nosy_AttestationIndexKey_id__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.AttestationIndexKey

// skipping Fuzz_Nosy_AttestationIndexKey_values__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.AttestationIndexKey

func Fuzz_Nosy_AttestationIterator_Value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i AttestationIterator
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		i.Value()
	})
}

func Fuzz_Nosy_AttestationStatusChainIdConfLevelAttestOffsetIndexKey_WithStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this AttestationStatusChainIdConfLevelAttestOffsetIndexKey
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var status uint32
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}

		this.WithStatus(status)
	})
}

func Fuzz_Nosy_AttestationStatusChainIdConfLevelAttestOffsetIndexKey_WithStatusChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this AttestationStatusChainIdConfLevelAttestOffsetIndexKey
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var status uint32
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var chain_id uint64
		fill_err = tp.Fill(&chain_id)
		if fill_err != nil {
			return
		}

		this.WithStatusChainId(status, chain_id)
	})
}

func Fuzz_Nosy_AttestationStatusChainIdConfLevelAttestOffsetIndexKey_WithStatusChainIdConfLevel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this AttestationStatusChainIdConfLevelAttestOffsetIndexKey
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var status uint32
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var chain_id uint64
		fill_err = tp.Fill(&chain_id)
		if fill_err != nil {
			return
		}
		var conf_level uint32
		fill_err = tp.Fill(&conf_level)
		if fill_err != nil {
			return
		}

		this.WithStatusChainIdConfLevel(status, chain_id, conf_level)
	})
}

func Fuzz_Nosy_AttestationStatusChainIdConfLevelAttestOffsetIndexKey_WithStatusChainIdConfLevelAttestOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this AttestationStatusChainIdConfLevelAttestOffsetIndexKey
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var status uint32
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var chain_id uint64
		fill_err = tp.Fill(&chain_id)
		if fill_err != nil {
			return
		}
		var conf_level uint32
		fill_err = tp.Fill(&conf_level)
		if fill_err != nil {
			return
		}
		var attest_offset uint64
		fill_err = tp.Fill(&attest_offset)
		if fill_err != nil {
			return
		}

		this.WithStatusChainIdConfLevelAttestOffset(status, chain_id, conf_level, attest_offset)
	})
}

func Fuzz_Nosy_AttestationStatusChainIdConfLevelAttestOffsetIndexKey_attestationIndexKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x AttestationStatusChainIdConfLevelAttestOffsetIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.attestationIndexKey()
	})
}

func Fuzz_Nosy_AttestationStatusChainIdConfLevelAttestOffsetIndexKey_id__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x AttestationStatusChainIdConfLevelAttestOffsetIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.id()
	})
}

func Fuzz_Nosy_AttestationStatusChainIdConfLevelAttestOffsetIndexKey_values__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x AttestationStatusChainIdConfLevelAttestOffsetIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.values()
	})
}

// skipping Fuzz_Nosy_AttestationStore_AttestationTable__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.AttestationStore

// skipping Fuzz_Nosy_AttestationStore_SignatureTable__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.AttestationStore

// skipping Fuzz_Nosy_AttestationStore_doNotImplement__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.AttestationStore

// skipping Fuzz_Nosy_AttestationTable_Delete__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.AttestationTable

// skipping Fuzz_Nosy_AttestationTable_DeleteBy__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.AttestationTable

// skipping Fuzz_Nosy_AttestationTable_DeleteRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.AttestationTable

// skipping Fuzz_Nosy_AttestationTable_Get__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.AttestationTable

// skipping Fuzz_Nosy_AttestationTable_GetByAttestationRoot__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.AttestationTable

// skipping Fuzz_Nosy_AttestationTable_Has__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.AttestationTable

// skipping Fuzz_Nosy_AttestationTable_HasByAttestationRoot__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.AttestationTable

// skipping Fuzz_Nosy_AttestationTable_Insert__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.AttestationTable

// skipping Fuzz_Nosy_AttestationTable_InsertReturningId__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.AttestationTable

// skipping Fuzz_Nosy_AttestationTable_LastInsertedSequence__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.AttestationTable

// skipping Fuzz_Nosy_AttestationTable_List__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.AttestationTable

// skipping Fuzz_Nosy_AttestationTable_ListRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.AttestationTable

// skipping Fuzz_Nosy_AttestationTable_Save__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.AttestationTable

// skipping Fuzz_Nosy_AttestationTable_Update__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.AttestationTable

// skipping Fuzz_Nosy_AttestationTable_doNotImplement__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.AttestationTable

func Fuzz_Nosy_SignatureAttIdValidatorAddressIndexKey_WithAttId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this SignatureAttIdValidatorAddressIndexKey
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var att_id uint64
		fill_err = tp.Fill(&att_id)
		if fill_err != nil {
			return
		}

		this.WithAttId(att_id)
	})
}

func Fuzz_Nosy_SignatureAttIdValidatorAddressIndexKey_WithAttIdValidatorAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this SignatureAttIdValidatorAddressIndexKey
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var att_id uint64
		fill_err = tp.Fill(&att_id)
		if fill_err != nil {
			return
		}
		var validator_address []byte
		fill_err = tp.Fill(&validator_address)
		if fill_err != nil {
			return
		}

		this.WithAttIdValidatorAddress(att_id, validator_address)
	})
}

func Fuzz_Nosy_SignatureAttIdValidatorAddressIndexKey_id__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x SignatureAttIdValidatorAddressIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.id()
	})
}

func Fuzz_Nosy_SignatureAttIdValidatorAddressIndexKey_signatureIndexKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x SignatureAttIdValidatorAddressIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.signatureIndexKey()
	})
}

func Fuzz_Nosy_SignatureAttIdValidatorAddressIndexKey_values__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x SignatureAttIdValidatorAddressIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.values()
	})
}

func Fuzz_Nosy_SignatureChainIdConfLevelAttestOffsetValidatorAddressIndexKey_WithChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this SignatureChainIdConfLevelAttestOffsetValidatorAddressIndexKey
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var chain_id uint64
		fill_err = tp.Fill(&chain_id)
		if fill_err != nil {
			return
		}

		this.WithChainId(chain_id)
	})
}

func Fuzz_Nosy_SignatureChainIdConfLevelAttestOffsetValidatorAddressIndexKey_WithChainIdConfLevel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this SignatureChainIdConfLevelAttestOffsetValidatorAddressIndexKey
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var chain_id uint64
		fill_err = tp.Fill(&chain_id)
		if fill_err != nil {
			return
		}
		var conf_level uint32
		fill_err = tp.Fill(&conf_level)
		if fill_err != nil {
			return
		}

		this.WithChainIdConfLevel(chain_id, conf_level)
	})
}

func Fuzz_Nosy_SignatureChainIdConfLevelAttestOffsetValidatorAddressIndexKey_WithChainIdConfLevelAttestOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this SignatureChainIdConfLevelAttestOffsetValidatorAddressIndexKey
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var chain_id uint64
		fill_err = tp.Fill(&chain_id)
		if fill_err != nil {
			return
		}
		var conf_level uint32
		fill_err = tp.Fill(&conf_level)
		if fill_err != nil {
			return
		}
		var attest_offset uint64
		fill_err = tp.Fill(&attest_offset)
		if fill_err != nil {
			return
		}

		this.WithChainIdConfLevelAttestOffset(chain_id, conf_level, attest_offset)
	})
}

func Fuzz_Nosy_SignatureChainIdConfLevelAttestOffsetValidatorAddressIndexKey_WithChainIdConfLevelAttestOffsetValidatorAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this SignatureChainIdConfLevelAttestOffsetValidatorAddressIndexKey
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var chain_id uint64
		fill_err = tp.Fill(&chain_id)
		if fill_err != nil {
			return
		}
		var conf_level uint32
		fill_err = tp.Fill(&conf_level)
		if fill_err != nil {
			return
		}
		var attest_offset uint64
		fill_err = tp.Fill(&attest_offset)
		if fill_err != nil {
			return
		}
		var validator_address []byte
		fill_err = tp.Fill(&validator_address)
		if fill_err != nil {
			return
		}

		this.WithChainIdConfLevelAttestOffsetValidatorAddress(chain_id, conf_level, attest_offset, validator_address)
	})
}

func Fuzz_Nosy_SignatureChainIdConfLevelAttestOffsetValidatorAddressIndexKey_id__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x SignatureChainIdConfLevelAttestOffsetValidatorAddressIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.id()
	})
}

func Fuzz_Nosy_SignatureChainIdConfLevelAttestOffsetValidatorAddressIndexKey_signatureIndexKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x SignatureChainIdConfLevelAttestOffsetValidatorAddressIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.signatureIndexKey()
	})
}

func Fuzz_Nosy_SignatureChainIdConfLevelAttestOffsetValidatorAddressIndexKey_values__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x SignatureChainIdConfLevelAttestOffsetValidatorAddressIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.values()
	})
}

func Fuzz_Nosy_SignatureIdIndexKey_WithId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this SignatureIdIndexKey
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		this.WithId(id)
	})
}

func Fuzz_Nosy_SignatureIdIndexKey_id__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x SignatureIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.id()
	})
}

func Fuzz_Nosy_SignatureIdIndexKey_signatureIndexKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x SignatureIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.signatureIndexKey()
	})
}

func Fuzz_Nosy_SignatureIdIndexKey_values__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x SignatureIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.values()
	})
}

// skipping Fuzz_Nosy_SignatureIndexKey_id__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.SignatureIndexKey

// skipping Fuzz_Nosy_SignatureIndexKey_signatureIndexKey__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.SignatureIndexKey

// skipping Fuzz_Nosy_SignatureIndexKey_values__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.SignatureIndexKey

func Fuzz_Nosy_SignatureIterator_Value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i SignatureIterator
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		i.Value()
	})
}

// skipping Fuzz_Nosy_SignatureTable_Delete__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.SignatureTable

// skipping Fuzz_Nosy_SignatureTable_DeleteBy__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.SignatureTable

// skipping Fuzz_Nosy_SignatureTable_DeleteRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.SignatureTable

// skipping Fuzz_Nosy_SignatureTable_Get__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.SignatureTable

// skipping Fuzz_Nosy_SignatureTable_GetByAttIdValidatorAddress__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.SignatureTable

// skipping Fuzz_Nosy_SignatureTable_GetByChainIdConfLevelAttestOffsetValidatorAddress__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.SignatureTable

// skipping Fuzz_Nosy_SignatureTable_Has__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.SignatureTable

// skipping Fuzz_Nosy_SignatureTable_HasByAttIdValidatorAddress__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.SignatureTable

// skipping Fuzz_Nosy_SignatureTable_HasByChainIdConfLevelAttestOffsetValidatorAddress__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.SignatureTable

// skipping Fuzz_Nosy_SignatureTable_Insert__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.SignatureTable

// skipping Fuzz_Nosy_SignatureTable_InsertReturningId__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.SignatureTable

// skipping Fuzz_Nosy_SignatureTable_LastInsertedSequence__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.SignatureTable

// skipping Fuzz_Nosy_SignatureTable_List__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.SignatureTable

// skipping Fuzz_Nosy_SignatureTable_ListRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.SignatureTable

// skipping Fuzz_Nosy_SignatureTable_Save__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.SignatureTable

// skipping Fuzz_Nosy_SignatureTable_Update__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.SignatureTable

// skipping Fuzz_Nosy_SignatureTable_doNotImplement__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.SignatureTable

func Fuzz_Nosy_Status_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 Status
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_Status_Enum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x Status
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.Enum()
	})
}

func Fuzz_Nosy_Status_EnumDescriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 Status
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.EnumDescriptor()
	})
}

func Fuzz_Nosy_Status_Number__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x Status
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.Number()
	})
}

func Fuzz_Nosy_Status_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x Status
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.String()
	})
}

func Fuzz_Nosy_Status_Type__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 Status
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Type()
	})
}

func Fuzz_Nosy_ValSet_Contains__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s ValSet
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}

		s.Contains(addr)
	})
}

func Fuzz_Nosy_ValSet_TotalPower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s ValSet
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.TotalPower()
	})
}

func Fuzz_Nosy_attestationStore_AttestationTable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x attestationStore
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.AttestationTable()
	})
}

func Fuzz_Nosy_attestationStore_SignatureTable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x attestationStore
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.SignatureTable()
	})
}

func Fuzz_Nosy_attestationStore_doNotImplement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 attestationStore
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.doNotImplement()
	})
}

func Fuzz_Nosy_attestationTable_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this attestationTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var attestation *Attestation
		fill_err = tp.Fill(&attestation)
		if fill_err != nil {
			return
		}
		if attestation == nil {
			return
		}

		this.Delete(ctx, attestation)
	})
}

// skipping Fuzz_Nosy_attestationTable_DeleteBy__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.AttestationIndexKey

// skipping Fuzz_Nosy_attestationTable_DeleteRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.AttestationIndexKey

func Fuzz_Nosy_attestationTable_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this attestationTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		this.Get(ctx, id)
	})
}

func Fuzz_Nosy_attestationTable_GetByAttestationRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this attestationTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var attestation_root []byte
		fill_err = tp.Fill(&attestation_root)
		if fill_err != nil {
			return
		}

		this.GetByAttestationRoot(ctx, attestation_root)
	})
}

func Fuzz_Nosy_attestationTable_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this attestationTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		this.Has(ctx, id)
	})
}

func Fuzz_Nosy_attestationTable_HasByAttestationRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this attestationTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var attestation_root []byte
		fill_err = tp.Fill(&attestation_root)
		if fill_err != nil {
			return
		}

		this.HasByAttestationRoot(ctx, attestation_root)
	})
}

func Fuzz_Nosy_attestationTable_Insert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this attestationTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var attestation *Attestation
		fill_err = tp.Fill(&attestation)
		if fill_err != nil {
			return
		}
		if attestation == nil {
			return
		}

		this.Insert(ctx, attestation)
	})
}

func Fuzz_Nosy_attestationTable_InsertReturningId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this attestationTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var attestation *Attestation
		fill_err = tp.Fill(&attestation)
		if fill_err != nil {
			return
		}
		if attestation == nil {
			return
		}

		this.InsertReturningId(ctx, attestation)
	})
}

func Fuzz_Nosy_attestationTable_LastInsertedSequence__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this attestationTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		this.LastInsertedSequence(ctx)
	})
}

// skipping Fuzz_Nosy_attestationTable_List__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.AttestationIndexKey

// skipping Fuzz_Nosy_attestationTable_ListRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.AttestationIndexKey

func Fuzz_Nosy_attestationTable_Save__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this attestationTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var attestation *Attestation
		fill_err = tp.Fill(&attestation)
		if fill_err != nil {
			return
		}
		if attestation == nil {
			return
		}

		this.Save(ctx, attestation)
	})
}

func Fuzz_Nosy_attestationTable_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this attestationTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var attestation *Attestation
		fill_err = tp.Fill(&attestation)
		if fill_err != nil {
			return
		}
		if attestation == nil {
			return
		}

		this.Update(ctx, attestation)
	})
}

func Fuzz_Nosy_attestationTable_doNotImplement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this attestationTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}

		this.doNotImplement()
	})
}

func Fuzz_Nosy_msgServer_AddVotes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s msgServer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg *types.MsgAddVotes
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		s.AddVotes(ctx, msg)
	})
}

func Fuzz_Nosy_proposalServer_AddVotes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s proposalServer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg *types.MsgAddVotes
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		s.AddVotes(ctx, msg)
	})
}

func Fuzz_Nosy_signatureTable_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this signatureTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var signature *Signature
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		if signature == nil {
			return
		}

		this.Delete(ctx, signature)
	})
}

// skipping Fuzz_Nosy_signatureTable_DeleteBy__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.SignatureIndexKey

// skipping Fuzz_Nosy_signatureTable_DeleteRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.SignatureIndexKey

func Fuzz_Nosy_signatureTable_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this signatureTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		this.Get(ctx, id)
	})
}

func Fuzz_Nosy_signatureTable_GetByAttIdValidatorAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this signatureTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var att_id uint64
		fill_err = tp.Fill(&att_id)
		if fill_err != nil {
			return
		}
		var validator_address []byte
		fill_err = tp.Fill(&validator_address)
		if fill_err != nil {
			return
		}

		this.GetByAttIdValidatorAddress(ctx, att_id, validator_address)
	})
}

func Fuzz_Nosy_signatureTable_GetByChainIdConfLevelAttestOffsetValidatorAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this signatureTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chain_id uint64
		fill_err = tp.Fill(&chain_id)
		if fill_err != nil {
			return
		}
		var conf_level uint32
		fill_err = tp.Fill(&conf_level)
		if fill_err != nil {
			return
		}
		var attest_offset uint64
		fill_err = tp.Fill(&attest_offset)
		if fill_err != nil {
			return
		}
		var validator_address []byte
		fill_err = tp.Fill(&validator_address)
		if fill_err != nil {
			return
		}

		this.GetByChainIdConfLevelAttestOffsetValidatorAddress(ctx, chain_id, conf_level, attest_offset, validator_address)
	})
}

func Fuzz_Nosy_signatureTable_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this signatureTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		this.Has(ctx, id)
	})
}

func Fuzz_Nosy_signatureTable_HasByAttIdValidatorAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this signatureTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var att_id uint64
		fill_err = tp.Fill(&att_id)
		if fill_err != nil {
			return
		}
		var validator_address []byte
		fill_err = tp.Fill(&validator_address)
		if fill_err != nil {
			return
		}

		this.HasByAttIdValidatorAddress(ctx, att_id, validator_address)
	})
}

func Fuzz_Nosy_signatureTable_HasByChainIdConfLevelAttestOffsetValidatorAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this signatureTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chain_id uint64
		fill_err = tp.Fill(&chain_id)
		if fill_err != nil {
			return
		}
		var conf_level uint32
		fill_err = tp.Fill(&conf_level)
		if fill_err != nil {
			return
		}
		var attest_offset uint64
		fill_err = tp.Fill(&attest_offset)
		if fill_err != nil {
			return
		}
		var validator_address []byte
		fill_err = tp.Fill(&validator_address)
		if fill_err != nil {
			return
		}

		this.HasByChainIdConfLevelAttestOffsetValidatorAddress(ctx, chain_id, conf_level, attest_offset, validator_address)
	})
}

func Fuzz_Nosy_signatureTable_Insert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this signatureTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var signature *Signature
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		if signature == nil {
			return
		}

		this.Insert(ctx, signature)
	})
}

func Fuzz_Nosy_signatureTable_InsertReturningId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this signatureTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var signature *Signature
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		if signature == nil {
			return
		}

		this.InsertReturningId(ctx, signature)
	})
}

func Fuzz_Nosy_signatureTable_LastInsertedSequence__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this signatureTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		this.LastInsertedSequence(ctx)
	})
}

// skipping Fuzz_Nosy_signatureTable_List__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.SignatureIndexKey

// skipping Fuzz_Nosy_signatureTable_ListRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/keeper.SignatureIndexKey

func Fuzz_Nosy_signatureTable_Save__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this signatureTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var signature *Signature
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		if signature == nil {
			return
		}

		this.Save(ctx, signature)
	})
}

func Fuzz_Nosy_signatureTable_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this signatureTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var signature *Signature
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		if signature == nil {
			return
		}

		this.Update(ctx, signature)
	})
}

func Fuzz_Nosy_signatureTable_doNotImplement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this signatureTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}

		this.doNotImplement()
	})
}

func Fuzz_Nosy_stubPortalRegistry_ConfLevels__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 stubPortalRegistry
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}

		_x1.ConfLevels(_x2)
	})
}

func Fuzz_Nosy_aggregateVotes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var votes []*types.Vote
		fill_err = tp.Fill(&votes)
		if fill_err != nil {
			return
		}

		aggregateVotes(votes)
	})
}

func Fuzz_Nosy_flattenAggs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var aggsByRoot map[common.Hash]*types.AggVote
		fill_err = tp.Fill(&aggsByRoot)
		if fill_err != nil {
			return
		}

		flattenAggs(aggsByRoot)
	})
}

func Fuzz_Nosy_fuzzyDependents__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainVer xchain.ChainVersion
		fill_err = tp.Fill(&chainVer)
		if fill_err != nil {
			return
		}
		var confLevels map[uint64][]xchain.ConfLevel
		fill_err = tp.Fill(&confLevels)
		if fill_err != nil {
			return
		}

		fuzzyDependents(chainVer, confLevels)
	})
}

func Fuzz_Nosy_getConsensusChainID__(f *testing.F) {
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

		getConsensusChainID(ctx)
	})
}

func Fuzz_Nosy_headersByAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var aggregates []*types.AggVote
		fill_err = tp.Fill(&aggregates)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}

		headersByAddress(aggregates, address)
	})
}

func Fuzz_Nosy_isApproved__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sigs []*Signature
		fill_err = tp.Fill(&sigs)
		if fill_err != nil {
			return
		}
		var valset ValSet
		fill_err = tp.Fill(&valset)
		if fill_err != nil {
			return
		}

		isApproved(sigs, valset)
	})
}

func Fuzz_Nosy_isApprovedByDifferentSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var att *Attestation
		fill_err = tp.Fill(&att)
		if fill_err != nil {
			return
		}
		var valSetID uint64
		fill_err = tp.Fill(&valSetID)
		if fill_err != nil {
			return
		}
		if att == nil {
			return
		}

		isApprovedByDifferentSet(att, valSetID)
	})
}

func Fuzz_Nosy_latency__(f *testing.F) {
	f.Fuzz(func(t *testing.T, method string) {
		latency(method)
	})
}

func Fuzz_Nosy_logLocalVotes__(f *testing.F) {
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
		var headers []*types.AttestHeader
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		var typ string
		fill_err = tp.Fill(&typ)
		if fill_err != nil {
			return
		}

		logLocalVotes(ctx, headers, typ)
	})
}

func Fuzz_Nosy_sigsFromDB__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sigs []*Signature
		fill_err = tp.Fill(&sigs)
		if fill_err != nil {
			return
		}

		sigsFromDB(sigs)
	})
}

func Fuzz_Nosy_sortAggregates__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var aggs []*types.AggVote
		fill_err = tp.Fill(&aggs)
		if fill_err != nil {
			return
		}

		sortAggregates(aggs)
	})
}

func Fuzz_Nosy_uintSub__(f *testing.F) {
	f.Fuzz(func(t *testing.T, a uint64, b uint64) {
		uintSub(a, b)
	})
}

func Fuzz_Nosy_votesFromExtension__(f *testing.F) {
	f.Fuzz(func(t *testing.T, voteExtension []byte) {
		votesFromExtension(voteExtension)
	})
}

func Fuzz_Nosy_windowCompare__(f *testing.F) {
	f.Fuzz(func(t *testing.T, voteWindow uint64, mid uint64, x uint64) {
		windowCompare(voteWindow, mid, x)
	})
}
