package keeper

import (
	"context"
	"testing"

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

func Fuzz_Nosy_Block_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Block
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

func Fuzz_Nosy_Block_GetCreatedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Block
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

func Fuzz_Nosy_Block_GetId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Block
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

func Fuzz_Nosy_Block_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Block
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

func Fuzz_Nosy_Block_ProtoReflect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Block
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

func Fuzz_Nosy_Block_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Block
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

func Fuzz_Nosy_Block_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Block
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

func Fuzz_Nosy_Msg_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Msg
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

func Fuzz_Nosy_Msg_GetBlockId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Msg
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetBlockId()
	})
}

func Fuzz_Nosy_Msg_GetDestChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Msg
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetDestChainId()
	})
}

func Fuzz_Nosy_Msg_GetId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Msg
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

func Fuzz_Nosy_Msg_GetMsgType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Msg
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetMsgType()
	})
}

func Fuzz_Nosy_Msg_GetMsgTypeId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Msg
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetMsgTypeId()
	})
}

func Fuzz_Nosy_Msg_GetShardId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Msg
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetShardId()
	})
}

func Fuzz_Nosy_Msg_GetStreamOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Msg
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetStreamOffset()
	})
}

func Fuzz_Nosy_Msg_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Msg
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

func Fuzz_Nosy_Msg_ProtoReflect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Msg
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

func Fuzz_Nosy_Msg_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Msg
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

func Fuzz_Nosy_Msg_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Msg
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

func Fuzz_Nosy_Offset_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Offset
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

func Fuzz_Nosy_Offset_GetDestChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Offset
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetDestChainId()
	})
}

func Fuzz_Nosy_Offset_GetId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Offset
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

func Fuzz_Nosy_Offset_GetOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Offset
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetOffset()
	})
}

func Fuzz_Nosy_Offset_GetShardId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Offset
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetShardId()
	})
}

func Fuzz_Nosy_Offset_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Offset
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

func Fuzz_Nosy_Offset_ProtoReflect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Offset
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

func Fuzz_Nosy_Offset_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Offset
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

func Fuzz_Nosy_Offset_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Offset
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

func Fuzz_Nosy_BlockCreatedHeightIndexKey_WithCreatedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this BlockCreatedHeightIndexKey
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

func Fuzz_Nosy_BlockCreatedHeightIndexKey_blockIndexKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x BlockCreatedHeightIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.blockIndexKey()
	})
}

func Fuzz_Nosy_BlockCreatedHeightIndexKey_id__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x BlockCreatedHeightIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.id()
	})
}

func Fuzz_Nosy_BlockCreatedHeightIndexKey_values__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x BlockCreatedHeightIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.values()
	})
}

func Fuzz_Nosy_BlockIdIndexKey_WithId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this BlockIdIndexKey
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

func Fuzz_Nosy_BlockIdIndexKey_blockIndexKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x BlockIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.blockIndexKey()
	})
}

func Fuzz_Nosy_BlockIdIndexKey_id__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x BlockIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.id()
	})
}

func Fuzz_Nosy_BlockIdIndexKey_values__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x BlockIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.values()
	})
}

// skipping Fuzz_Nosy_BlockIndexKey_blockIndexKey__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.BlockIndexKey

// skipping Fuzz_Nosy_BlockIndexKey_id__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.BlockIndexKey

// skipping Fuzz_Nosy_BlockIndexKey_values__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.BlockIndexKey

func Fuzz_Nosy_BlockIterator_Value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i BlockIterator
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		i.Value()
	})
}

// skipping Fuzz_Nosy_BlockTable_Delete__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.BlockTable

// skipping Fuzz_Nosy_BlockTable_DeleteBy__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.BlockTable

// skipping Fuzz_Nosy_BlockTable_DeleteRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.BlockTable

// skipping Fuzz_Nosy_BlockTable_Get__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.BlockTable

// skipping Fuzz_Nosy_BlockTable_GetByCreatedHeight__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.BlockTable

// skipping Fuzz_Nosy_BlockTable_Has__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.BlockTable

// skipping Fuzz_Nosy_BlockTable_HasByCreatedHeight__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.BlockTable

// skipping Fuzz_Nosy_BlockTable_Insert__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.BlockTable

// skipping Fuzz_Nosy_BlockTable_InsertReturningId__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.BlockTable

// skipping Fuzz_Nosy_BlockTable_LastInsertedSequence__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.BlockTable

// skipping Fuzz_Nosy_BlockTable_List__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.BlockTable

// skipping Fuzz_Nosy_BlockTable_ListRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.BlockTable

// skipping Fuzz_Nosy_BlockTable_Save__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.BlockTable

// skipping Fuzz_Nosy_BlockTable_Update__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.BlockTable

// skipping Fuzz_Nosy_BlockTable_doNotImplement__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.BlockTable

func Fuzz_Nosy_Keeper_Block__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *types.BlockRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.Block(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_EmitMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
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

		k.EmitMsg(ctx, typ, msgTypeID, destChainID, shardID)
	})
}

func Fuzz_Nosy_Keeper_getBlockAndMsgs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockID uint64
		fill_err = tp.Fill(&blockID)
		if fill_err != nil {
			return
		}

		k.getBlockAndMsgs(ctx, blockID)
	})
}

func Fuzz_Nosy_Keeper_incAndGetOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
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

		k.incAndGetOffset(ctx, destChainID, shardID)
	})
}

func Fuzz_Nosy_MsgBlockIdIndexKey_WithBlockId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this MsgBlockIdIndexKey
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var block_id uint64
		fill_err = tp.Fill(&block_id)
		if fill_err != nil {
			return
		}

		this.WithBlockId(block_id)
	})
}

func Fuzz_Nosy_MsgBlockIdIndexKey_id__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x MsgBlockIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.id()
	})
}

func Fuzz_Nosy_MsgBlockIdIndexKey_msgIndexKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x MsgBlockIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.msgIndexKey()
	})
}

func Fuzz_Nosy_MsgBlockIdIndexKey_values__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x MsgBlockIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.values()
	})
}

func Fuzz_Nosy_MsgIdIndexKey_WithId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this MsgIdIndexKey
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

func Fuzz_Nosy_MsgIdIndexKey_id__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x MsgIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.id()
	})
}

func Fuzz_Nosy_MsgIdIndexKey_msgIndexKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x MsgIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.msgIndexKey()
	})
}

func Fuzz_Nosy_MsgIdIndexKey_values__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x MsgIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.values()
	})
}

// skipping Fuzz_Nosy_MsgIndexKey_id__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.MsgIndexKey

// skipping Fuzz_Nosy_MsgIndexKey_msgIndexKey__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.MsgIndexKey

// skipping Fuzz_Nosy_MsgIndexKey_values__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.MsgIndexKey

func Fuzz_Nosy_MsgIterator_Value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i MsgIterator
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		i.Value()
	})
}

// skipping Fuzz_Nosy_MsgTable_Delete__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.MsgTable

// skipping Fuzz_Nosy_MsgTable_DeleteBy__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.MsgTable

// skipping Fuzz_Nosy_MsgTable_DeleteRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.MsgTable

// skipping Fuzz_Nosy_MsgTable_Get__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.MsgTable

// skipping Fuzz_Nosy_MsgTable_Has__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.MsgTable

// skipping Fuzz_Nosy_MsgTable_Insert__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.MsgTable

// skipping Fuzz_Nosy_MsgTable_InsertReturningId__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.MsgTable

// skipping Fuzz_Nosy_MsgTable_LastInsertedSequence__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.MsgTable

// skipping Fuzz_Nosy_MsgTable_List__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.MsgTable

// skipping Fuzz_Nosy_MsgTable_ListRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.MsgTable

// skipping Fuzz_Nosy_MsgTable_Save__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.MsgTable

// skipping Fuzz_Nosy_MsgTable_Update__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.MsgTable

// skipping Fuzz_Nosy_MsgTable_doNotImplement__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.MsgTable

func Fuzz_Nosy_OffsetDestChainIdShardIdIndexKey_WithDestChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this OffsetDestChainIdShardIdIndexKey
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var dest_chain_id uint64
		fill_err = tp.Fill(&dest_chain_id)
		if fill_err != nil {
			return
		}

		this.WithDestChainId(dest_chain_id)
	})
}

func Fuzz_Nosy_OffsetDestChainIdShardIdIndexKey_WithDestChainIdShardId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this OffsetDestChainIdShardIdIndexKey
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var dest_chain_id uint64
		fill_err = tp.Fill(&dest_chain_id)
		if fill_err != nil {
			return
		}
		var shard_id uint64
		fill_err = tp.Fill(&shard_id)
		if fill_err != nil {
			return
		}

		this.WithDestChainIdShardId(dest_chain_id, shard_id)
	})
}

func Fuzz_Nosy_OffsetDestChainIdShardIdIndexKey_id__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x OffsetDestChainIdShardIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.id()
	})
}

func Fuzz_Nosy_OffsetDestChainIdShardIdIndexKey_offsetIndexKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x OffsetDestChainIdShardIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.offsetIndexKey()
	})
}

func Fuzz_Nosy_OffsetDestChainIdShardIdIndexKey_values__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x OffsetDestChainIdShardIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.values()
	})
}

func Fuzz_Nosy_OffsetIdIndexKey_WithId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this OffsetIdIndexKey
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

func Fuzz_Nosy_OffsetIdIndexKey_id__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x OffsetIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.id()
	})
}

func Fuzz_Nosy_OffsetIdIndexKey_offsetIndexKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x OffsetIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.offsetIndexKey()
	})
}

func Fuzz_Nosy_OffsetIdIndexKey_values__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x OffsetIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.values()
	})
}

// skipping Fuzz_Nosy_OffsetIndexKey_id__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.OffsetIndexKey

// skipping Fuzz_Nosy_OffsetIndexKey_offsetIndexKey__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.OffsetIndexKey

// skipping Fuzz_Nosy_OffsetIndexKey_values__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.OffsetIndexKey

func Fuzz_Nosy_OffsetIterator_Value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i OffsetIterator
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		i.Value()
	})
}

// skipping Fuzz_Nosy_OffsetTable_Delete__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.OffsetTable

// skipping Fuzz_Nosy_OffsetTable_DeleteBy__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.OffsetTable

// skipping Fuzz_Nosy_OffsetTable_DeleteRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.OffsetTable

// skipping Fuzz_Nosy_OffsetTable_Get__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.OffsetTable

// skipping Fuzz_Nosy_OffsetTable_GetByDestChainIdShardId__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.OffsetTable

// skipping Fuzz_Nosy_OffsetTable_Has__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.OffsetTable

// skipping Fuzz_Nosy_OffsetTable_HasByDestChainIdShardId__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.OffsetTable

// skipping Fuzz_Nosy_OffsetTable_Insert__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.OffsetTable

// skipping Fuzz_Nosy_OffsetTable_InsertReturningId__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.OffsetTable

// skipping Fuzz_Nosy_OffsetTable_LastInsertedSequence__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.OffsetTable

// skipping Fuzz_Nosy_OffsetTable_List__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.OffsetTable

// skipping Fuzz_Nosy_OffsetTable_ListRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.OffsetTable

// skipping Fuzz_Nosy_OffsetTable_Save__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.OffsetTable

// skipping Fuzz_Nosy_OffsetTable_Update__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.OffsetTable

// skipping Fuzz_Nosy_OffsetTable_doNotImplement__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.OffsetTable

// skipping Fuzz_Nosy_PortalStore_BlockTable__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.PortalStore

// skipping Fuzz_Nosy_PortalStore_MsgTable__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.PortalStore

// skipping Fuzz_Nosy_PortalStore_OffsetTable__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.PortalStore

// skipping Fuzz_Nosy_PortalStore_doNotImplement__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.PortalStore

func Fuzz_Nosy_blockTable_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this blockTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block *Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if block == nil {
			return
		}

		this.Delete(ctx, block)
	})
}

// skipping Fuzz_Nosy_blockTable_DeleteBy__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.BlockIndexKey

// skipping Fuzz_Nosy_blockTable_DeleteRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.BlockIndexKey

func Fuzz_Nosy_blockTable_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this blockTable
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

func Fuzz_Nosy_blockTable_GetByCreatedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this blockTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var created_height uint64
		fill_err = tp.Fill(&created_height)
		if fill_err != nil {
			return
		}

		this.GetByCreatedHeight(ctx, created_height)
	})
}

func Fuzz_Nosy_blockTable_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this blockTable
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

func Fuzz_Nosy_blockTable_HasByCreatedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this blockTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var created_height uint64
		fill_err = tp.Fill(&created_height)
		if fill_err != nil {
			return
		}

		this.HasByCreatedHeight(ctx, created_height)
	})
}

func Fuzz_Nosy_blockTable_Insert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this blockTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block *Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if block == nil {
			return
		}

		this.Insert(ctx, block)
	})
}

func Fuzz_Nosy_blockTable_InsertReturningId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this blockTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block *Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if block == nil {
			return
		}

		this.InsertReturningId(ctx, block)
	})
}

func Fuzz_Nosy_blockTable_LastInsertedSequence__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this blockTable
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

// skipping Fuzz_Nosy_blockTable_List__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.BlockIndexKey

// skipping Fuzz_Nosy_blockTable_ListRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.BlockIndexKey

func Fuzz_Nosy_blockTable_Save__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this blockTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block *Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if block == nil {
			return
		}

		this.Save(ctx, block)
	})
}

func Fuzz_Nosy_blockTable_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this blockTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block *Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if block == nil {
			return
		}

		this.Update(ctx, block)
	})
}

func Fuzz_Nosy_blockTable_doNotImplement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this blockTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}

		this.doNotImplement()
	})
}

func Fuzz_Nosy_msgTable_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this msgTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg *Msg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		this.Delete(ctx, msg)
	})
}

// skipping Fuzz_Nosy_msgTable_DeleteBy__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.MsgIndexKey

// skipping Fuzz_Nosy_msgTable_DeleteRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.MsgIndexKey

func Fuzz_Nosy_msgTable_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this msgTable
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

func Fuzz_Nosy_msgTable_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this msgTable
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

func Fuzz_Nosy_msgTable_Insert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this msgTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg *Msg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		this.Insert(ctx, msg)
	})
}

func Fuzz_Nosy_msgTable_InsertReturningId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this msgTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg *Msg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		this.InsertReturningId(ctx, msg)
	})
}

func Fuzz_Nosy_msgTable_LastInsertedSequence__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this msgTable
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

// skipping Fuzz_Nosy_msgTable_List__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.MsgIndexKey

// skipping Fuzz_Nosy_msgTable_ListRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.MsgIndexKey

func Fuzz_Nosy_msgTable_Save__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this msgTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg *Msg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		this.Save(ctx, msg)
	})
}

func Fuzz_Nosy_msgTable_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this msgTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg *Msg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		this.Update(ctx, msg)
	})
}

func Fuzz_Nosy_msgTable_doNotImplement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this msgTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}

		this.doNotImplement()
	})
}

func Fuzz_Nosy_offsetTable_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this offsetTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var offset *Offset
		fill_err = tp.Fill(&offset)
		if fill_err != nil {
			return
		}
		if offset == nil {
			return
		}

		this.Delete(ctx, offset)
	})
}

// skipping Fuzz_Nosy_offsetTable_DeleteBy__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.OffsetIndexKey

// skipping Fuzz_Nosy_offsetTable_DeleteRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.OffsetIndexKey

func Fuzz_Nosy_offsetTable_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this offsetTable
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

func Fuzz_Nosy_offsetTable_GetByDestChainIdShardId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this offsetTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var dest_chain_id uint64
		fill_err = tp.Fill(&dest_chain_id)
		if fill_err != nil {
			return
		}
		var shard_id uint64
		fill_err = tp.Fill(&shard_id)
		if fill_err != nil {
			return
		}

		this.GetByDestChainIdShardId(ctx, dest_chain_id, shard_id)
	})
}

func Fuzz_Nosy_offsetTable_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this offsetTable
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

func Fuzz_Nosy_offsetTable_HasByDestChainIdShardId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this offsetTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var dest_chain_id uint64
		fill_err = tp.Fill(&dest_chain_id)
		if fill_err != nil {
			return
		}
		var shard_id uint64
		fill_err = tp.Fill(&shard_id)
		if fill_err != nil {
			return
		}

		this.HasByDestChainIdShardId(ctx, dest_chain_id, shard_id)
	})
}

func Fuzz_Nosy_offsetTable_Insert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this offsetTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var offset *Offset
		fill_err = tp.Fill(&offset)
		if fill_err != nil {
			return
		}
		if offset == nil {
			return
		}

		this.Insert(ctx, offset)
	})
}

func Fuzz_Nosy_offsetTable_InsertReturningId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this offsetTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var offset *Offset
		fill_err = tp.Fill(&offset)
		if fill_err != nil {
			return
		}
		if offset == nil {
			return
		}

		this.InsertReturningId(ctx, offset)
	})
}

func Fuzz_Nosy_offsetTable_LastInsertedSequence__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this offsetTable
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

// skipping Fuzz_Nosy_offsetTable_List__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.OffsetIndexKey

// skipping Fuzz_Nosy_offsetTable_ListRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/portal/keeper.OffsetIndexKey

func Fuzz_Nosy_offsetTable_Save__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this offsetTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var offset *Offset
		fill_err = tp.Fill(&offset)
		if fill_err != nil {
			return
		}
		if offset == nil {
			return
		}

		this.Save(ctx, offset)
	})
}

func Fuzz_Nosy_offsetTable_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this offsetTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var offset *Offset
		fill_err = tp.Fill(&offset)
		if fill_err != nil {
			return
		}
		if offset == nil {
			return
		}

		this.Update(ctx, offset)
	})
}

func Fuzz_Nosy_offsetTable_doNotImplement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this offsetTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}

		this.doNotImplement()
	})
}

func Fuzz_Nosy_portalStore_BlockTable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x portalStore
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.BlockTable()
	})
}

func Fuzz_Nosy_portalStore_MsgTable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x portalStore
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.MsgTable()
	})
}

func Fuzz_Nosy_portalStore_OffsetTable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x portalStore
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.OffsetTable()
	})
}

func Fuzz_Nosy_portalStore_doNotImplement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 portalStore
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.doNotImplement()
	})
}
