package keeper

import (
	"context"
	"testing"
	"time"

	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/beacon/engine"
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

func Fuzz_Nosy_ExecutionHead_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *ExecutionHead
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

func Fuzz_Nosy_ExecutionHead_GetBlockHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *ExecutionHead
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

func Fuzz_Nosy_ExecutionHead_GetBlockHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *ExecutionHead
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

func Fuzz_Nosy_ExecutionHead_GetBlockTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *ExecutionHead
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetBlockTime()
	})
}

func Fuzz_Nosy_ExecutionHead_GetCreatedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *ExecutionHead
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

func Fuzz_Nosy_ExecutionHead_GetId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *ExecutionHead
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

func Fuzz_Nosy_ExecutionHead_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *ExecutionHead
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Hash()
	})
}

func Fuzz_Nosy_ExecutionHead_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *ExecutionHead
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

func Fuzz_Nosy_ExecutionHead_ProtoReflect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *ExecutionHead
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

func Fuzz_Nosy_ExecutionHead_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *ExecutionHead
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

func Fuzz_Nosy_ExecutionHead_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *ExecutionHead
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

func Fuzz_Nosy_Keeper_InsertGenesisHead__(f *testing.F) {
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
		var executionBlockHash []byte
		fill_err = tp.Fill(&executionBlockHash)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		k.InsertGenesisHead(ctx, executionBlockHash)
	})
}

func Fuzz_Nosy_Keeper_PostFinalize__(f *testing.F) {
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
		if k == nil {
			return
		}

		k.PostFinalize(ctx)
	})
}

func Fuzz_Nosy_Keeper_PrepareProposal__(f *testing.F) {
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
		var req *types.RequestPrepareProposal
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if k == nil || req == nil {
			return
		}

		k.PrepareProposal(ctx, req)
	})
}

// skipping Fuzz_Nosy_Keeper_RegisterProposalService__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/grpc.Server

func Fuzz_Nosy_Keeper_SetBuildDelay__(f *testing.F) {
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
		var d time.Duration
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		k.SetBuildDelay(d)
	})
}

func Fuzz_Nosy_Keeper_SetBuildOptimistic__(f *testing.F) {
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
		var b bool
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		k.SetBuildOptimistic(b)
	})
}

// skipping Fuzz_Nosy_Keeper_SetCometAPI__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/comet.API

// skipping Fuzz_Nosy_Keeper_SetVoteProvider__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/octane/evmengine/types.VoteExtensionProvider

func Fuzz_Nosy_Keeper_evmEvents__(f *testing.F) {
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
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		k.evmEvents(ctx, blockHash)
	})
}

func Fuzz_Nosy_Keeper_getExecutionHead__(f *testing.F) {
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

		k.getExecutionHead(ctx)
	})
}

func Fuzz_Nosy_Keeper_getOptimisticPayload__(f *testing.F) {
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
		if k == nil {
			return
		}

		k.getOptimisticPayload()
	})
}

func Fuzz_Nosy_Keeper_isNextProposer__(f *testing.F) {
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
		var currentProposer []byte
		fill_err = tp.Fill(&currentProposer)
		if fill_err != nil {
			return
		}
		var currentHeight int64
		fill_err = tp.Fill(&currentHeight)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		k.isNextProposer(ctx, currentProposer, currentHeight)
	})
}

func Fuzz_Nosy_Keeper_parseAndVerifyProposedPayload__(f *testing.F) {
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
		var msg *types.MsgExecutionPayload
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if k == nil || msg == nil {
			return
		}

		k.parseAndVerifyProposedPayload(ctx, msg)
	})
}

func Fuzz_Nosy_Keeper_setOptimisticPayload__(f *testing.F) {
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
		var id *engine.PayloadID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if k == nil || id == nil {
			return
		}

		k.setOptimisticPayload(id, height)
	})
}

func Fuzz_Nosy_Keeper_startBuild__(f *testing.F) {
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
		var appHash common.Hash
		fill_err = tp.Fill(&appHash)
		if fill_err != nil {
			return
		}
		var timestamp time.Time
		fill_err = tp.Fill(&timestamp)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		k.startBuild(ctx, appHash, timestamp)
	})
}

func Fuzz_Nosy_Keeper_updateExecutionHead__(f *testing.F) {
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
		var payload engine.ExecutableData
		fill_err = tp.Fill(&payload)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		k.updateExecutionHead(ctx, payload)
	})
}

// skipping Fuzz_Nosy_EvmengineStore_ExecutionHeadTable__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/octane/evmengine/keeper.EvmengineStore

// skipping Fuzz_Nosy_EvmengineStore_doNotImplement__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/octane/evmengine/keeper.EvmengineStore

func Fuzz_Nosy_ExecutionHeadIdIndexKey_WithId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this ExecutionHeadIdIndexKey
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

func Fuzz_Nosy_ExecutionHeadIdIndexKey_executionHeadIndexKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x ExecutionHeadIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.executionHeadIndexKey()
	})
}

func Fuzz_Nosy_ExecutionHeadIdIndexKey_id__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x ExecutionHeadIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.id()
	})
}

func Fuzz_Nosy_ExecutionHeadIdIndexKey_values__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x ExecutionHeadIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.values()
	})
}

// skipping Fuzz_Nosy_ExecutionHeadIndexKey_executionHeadIndexKey__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/octane/evmengine/keeper.ExecutionHeadIndexKey

// skipping Fuzz_Nosy_ExecutionHeadIndexKey_id__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/octane/evmengine/keeper.ExecutionHeadIndexKey

// skipping Fuzz_Nosy_ExecutionHeadIndexKey_values__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/octane/evmengine/keeper.ExecutionHeadIndexKey

func Fuzz_Nosy_ExecutionHeadIterator_Value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i ExecutionHeadIterator
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		i.Value()
	})
}

// skipping Fuzz_Nosy_ExecutionHeadTable_Delete__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/octane/evmengine/keeper.ExecutionHeadTable

// skipping Fuzz_Nosy_ExecutionHeadTable_DeleteBy__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/octane/evmengine/keeper.ExecutionHeadTable

// skipping Fuzz_Nosy_ExecutionHeadTable_DeleteRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/octane/evmengine/keeper.ExecutionHeadTable

// skipping Fuzz_Nosy_ExecutionHeadTable_Get__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/octane/evmengine/keeper.ExecutionHeadTable

// skipping Fuzz_Nosy_ExecutionHeadTable_Has__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/octane/evmengine/keeper.ExecutionHeadTable

// skipping Fuzz_Nosy_ExecutionHeadTable_Insert__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/octane/evmengine/keeper.ExecutionHeadTable

// skipping Fuzz_Nosy_ExecutionHeadTable_InsertReturningId__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/octane/evmengine/keeper.ExecutionHeadTable

// skipping Fuzz_Nosy_ExecutionHeadTable_LastInsertedSequence__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/octane/evmengine/keeper.ExecutionHeadTable

// skipping Fuzz_Nosy_ExecutionHeadTable_List__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/octane/evmengine/keeper.ExecutionHeadTable

// skipping Fuzz_Nosy_ExecutionHeadTable_ListRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/octane/evmengine/keeper.ExecutionHeadTable

// skipping Fuzz_Nosy_ExecutionHeadTable_Save__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/octane/evmengine/keeper.ExecutionHeadTable

// skipping Fuzz_Nosy_ExecutionHeadTable_Update__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/octane/evmengine/keeper.ExecutionHeadTable

// skipping Fuzz_Nosy_ExecutionHeadTable_doNotImplement__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/octane/evmengine/keeper.ExecutionHeadTable

func Fuzz_Nosy_Hooks_AfterDelegationModified__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 Hooks
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var accAddr types.AccAddress
		fill_err = tp.Fill(&accAddr)
		if fill_err != nil {
			return
		}
		var valAddr types.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		_x1.AfterDelegationModified(ctx, accAddr, valAddr)
	})
}

func Fuzz_Nosy_Hooks_AfterUnbondingInitiated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 Hooks
		fill_err = tp.Fill(&_x1)
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

		_x1.AfterUnbondingInitiated(ctx, id)
	})
}

func Fuzz_Nosy_Hooks_AfterValidatorBeginUnbonding__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 Hooks
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var consAddr types.ConsAddress
		fill_err = tp.Fill(&consAddr)
		if fill_err != nil {
			return
		}
		var valAddr types.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		_x1.AfterValidatorBeginUnbonding(ctx, consAddr, valAddr)
	})
}

func Fuzz_Nosy_Hooks_AfterValidatorBonded__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 Hooks
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var consAddr types.ConsAddress
		fill_err = tp.Fill(&consAddr)
		if fill_err != nil {
			return
		}
		var valAddr types.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		_x1.AfterValidatorBonded(ctx, consAddr, valAddr)
	})
}

func Fuzz_Nosy_Hooks_AfterValidatorCreated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 Hooks
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var valAddr types.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		_x1.AfterValidatorCreated(ctx, valAddr)
	})
}

func Fuzz_Nosy_Hooks_AfterValidatorRemoved__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 Hooks
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var consAddr types.ConsAddress
		fill_err = tp.Fill(&consAddr)
		if fill_err != nil {
			return
		}
		var valAddr types.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		_x1.AfterValidatorRemoved(ctx, consAddr, valAddr)
	})
}

func Fuzz_Nosy_Hooks_BeforeDelegationCreated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 Hooks
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var accAddr types.AccAddress
		fill_err = tp.Fill(&accAddr)
		if fill_err != nil {
			return
		}
		var valAddr types.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		_x1.BeforeDelegationCreated(ctx, accAddr, valAddr)
	})
}

func Fuzz_Nosy_Hooks_BeforeDelegationRemoved__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 Hooks
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var accAddr types.AccAddress
		fill_err = tp.Fill(&accAddr)
		if fill_err != nil {
			return
		}
		var valAddr types.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		_x1.BeforeDelegationRemoved(ctx, accAddr, valAddr)
	})
}

func Fuzz_Nosy_Hooks_BeforeDelegationSharesModified__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 Hooks
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var accAddr types.AccAddress
		fill_err = tp.Fill(&accAddr)
		if fill_err != nil {
			return
		}
		var valAddr types.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		_x1.BeforeDelegationSharesModified(ctx, accAddr, valAddr)
	})
}

func Fuzz_Nosy_Hooks_BeforeValidatorModified__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 Hooks
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var valAddr types.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		_x1.BeforeValidatorModified(ctx, valAddr)
	})
}

func Fuzz_Nosy_Hooks_BeforeValidatorSlashed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 Hooks
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var valAddr types.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}
		var amount math.LegacyDec
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}

		_x1.BeforeValidatorSlashed(ctx, valAddr, amount)
	})
}

func Fuzz_Nosy_evmengineStore_ExecutionHeadTable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x evmengineStore
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.ExecutionHeadTable()
	})
}

func Fuzz_Nosy_evmengineStore_doNotImplement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 evmengineStore
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.doNotImplement()
	})
}

func Fuzz_Nosy_executionHeadTable_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this executionHeadTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var executionHead *ExecutionHead
		fill_err = tp.Fill(&executionHead)
		if fill_err != nil {
			return
		}
		if executionHead == nil {
			return
		}

		this.Delete(ctx, executionHead)
	})
}

// skipping Fuzz_Nosy_executionHeadTable_DeleteBy__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/octane/evmengine/keeper.ExecutionHeadIndexKey

// skipping Fuzz_Nosy_executionHeadTable_DeleteRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/octane/evmengine/keeper.ExecutionHeadIndexKey

func Fuzz_Nosy_executionHeadTable_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this executionHeadTable
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

func Fuzz_Nosy_executionHeadTable_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this executionHeadTable
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

func Fuzz_Nosy_executionHeadTable_Insert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this executionHeadTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var executionHead *ExecutionHead
		fill_err = tp.Fill(&executionHead)
		if fill_err != nil {
			return
		}
		if executionHead == nil {
			return
		}

		this.Insert(ctx, executionHead)
	})
}

func Fuzz_Nosy_executionHeadTable_InsertReturningId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this executionHeadTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var executionHead *ExecutionHead
		fill_err = tp.Fill(&executionHead)
		if fill_err != nil {
			return
		}
		if executionHead == nil {
			return
		}

		this.InsertReturningId(ctx, executionHead)
	})
}

func Fuzz_Nosy_executionHeadTable_LastInsertedSequence__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this executionHeadTable
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

// skipping Fuzz_Nosy_executionHeadTable_List__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/octane/evmengine/keeper.ExecutionHeadIndexKey

// skipping Fuzz_Nosy_executionHeadTable_ListRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/octane/evmengine/keeper.ExecutionHeadIndexKey

func Fuzz_Nosy_executionHeadTable_Save__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this executionHeadTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var executionHead *ExecutionHead
		fill_err = tp.Fill(&executionHead)
		if fill_err != nil {
			return
		}
		if executionHead == nil {
			return
		}

		this.Save(ctx, executionHead)
	})
}

func Fuzz_Nosy_executionHeadTable_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this executionHeadTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var executionHead *ExecutionHead
		fill_err = tp.Fill(&executionHead)
		if fill_err != nil {
			return
		}
		if executionHead == nil {
			return
		}

		this.Update(ctx, executionHead)
	})
}

func Fuzz_Nosy_executionHeadTable_doNotImplement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this executionHeadTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}

		this.doNotImplement()
	})
}

func Fuzz_Nosy_msgServer_ExecutionPayload__(f *testing.F) {
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
		var msg *types.MsgExecutionPayload
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		s.ExecutionPayload(ctx, msg)
	})
}

func Fuzz_Nosy_msgServer_deliverEvents__(f *testing.F) {
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
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		var logs []*types.EVMEvent
		fill_err = tp.Fill(&logs)
		if fill_err != nil {
			return
		}

		s.deliverEvents(ctx, height, blockHash, logs)
	})
}

func Fuzz_Nosy_proposalServer_ExecutionPayload__(f *testing.F) {
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
		var msg *types.MsgExecutionPayload
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		s.ExecutionPayload(ctx, msg)
	})
}

func Fuzz_Nosy_isInvalid__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var status engine.PayloadStatusV1
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}

		isInvalid(status)
	})
}

func Fuzz_Nosy_isSyncing__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var status engine.PayloadStatusV1
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}

		isSyncing(status)
	})
}

func Fuzz_Nosy_isUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var status engine.PayloadStatusV1
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}

		isUnknown(status)
	})
}

// skipping Fuzz_Nosy_isUnknownPayload__ because parameters include func, chan, or unsupported interface: error
