package emitcache

import (
	context "context"
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

func Fuzz_Nosy_EmitCursor_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EmitCursor
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

func Fuzz_Nosy_EmitCursor_GetDstChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *EmitCursor
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetDstChainId()
	})
}

func Fuzz_Nosy_EmitCursor_GetHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *EmitCursor
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetHeight()
	})
}

func Fuzz_Nosy_EmitCursor_GetId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *EmitCursor
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

func Fuzz_Nosy_EmitCursor_GetMsgOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *EmitCursor
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetMsgOffset()
	})
}

func Fuzz_Nosy_EmitCursor_GetShardId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *EmitCursor
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

func Fuzz_Nosy_EmitCursor_GetSrcChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *EmitCursor
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetSrcChainId()
	})
}

func Fuzz_Nosy_EmitCursor_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EmitCursor
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

func Fuzz_Nosy_EmitCursor_ProtoReflect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *EmitCursor
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

func Fuzz_Nosy_EmitCursor_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *EmitCursor
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

func Fuzz_Nosy_EmitCursor_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *EmitCursor
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

func Fuzz_Nosy_emitCursorCache_AtOrBefore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *emitCursorCache
		fill_err = tp.Fill(&c)
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
		var stream xchain.StreamID
		fill_err = tp.Fill(&stream)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.AtOrBefore(ctx, height, stream)
	})
}

func Fuzz_Nosy_emitCursorCache_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *emitCursorCache
		fill_err = tp.Fill(&c)
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
		var stream xchain.StreamID
		fill_err = tp.Fill(&stream)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Get(ctx, height, stream)
	})
}

func Fuzz_Nosy_emitCursorCache_detectTrim__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *emitCursorCache
		fill_err = tp.Fill(&c)
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
		var retainHeights uint64
		fill_err = tp.Fill(&retainHeights)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.detectTrim(ctx, chainID, retainHeights)
	})
}

func Fuzz_Nosy_emitCursorCache_set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *emitCursorCache
		fill_err = tp.Fill(&c)
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
		var cursor xchain.EmitCursor
		fill_err = tp.Fill(&cursor)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.set(ctx, height, cursor)
	})
}

func Fuzz_Nosy_emitCursorCache_trimForever__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *emitCursorCache
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var network netconf.ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var chainID uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.trimForever(ctx, network, chainID)
	})
}

func Fuzz_Nosy_emitCursorCache_trimOnce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *emitCursorCache
		fill_err = tp.Fill(&c)
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
		var retain uint64
		fill_err = tp.Fill(&retain)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.trimOnce(ctx, chainID, retain)
	})
}

// skipping Fuzz_Nosy_Cache_AtOrBefore__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xmonitor/emitcache.Cache

// skipping Fuzz_Nosy_Cache_Get__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xmonitor/emitcache.Cache

func Fuzz_Nosy_EmitCursorIdIndexKey_WithId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this EmitCursorIdIndexKey
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

func Fuzz_Nosy_EmitCursorIdIndexKey_emitCursorIndexKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x EmitCursorIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.emitCursorIndexKey()
	})
}

func Fuzz_Nosy_EmitCursorIdIndexKey_id__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x EmitCursorIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.id()
	})
}

func Fuzz_Nosy_EmitCursorIdIndexKey_values__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x EmitCursorIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.values()
	})
}

// skipping Fuzz_Nosy_EmitCursorIndexKey_emitCursorIndexKey__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xmonitor/emitcache.EmitCursorIndexKey

// skipping Fuzz_Nosy_EmitCursorIndexKey_id__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xmonitor/emitcache.EmitCursorIndexKey

// skipping Fuzz_Nosy_EmitCursorIndexKey_values__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xmonitor/emitcache.EmitCursorIndexKey

func Fuzz_Nosy_EmitCursorIterator_Value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i EmitCursorIterator
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		i.Value()
	})
}

func Fuzz_Nosy_EmitCursorSrcChainIdDstChainIdShardIdHeightIndexKey_WithSrcChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this EmitCursorSrcChainIdDstChainIdShardIdHeightIndexKey
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var src_chain_id uint64
		fill_err = tp.Fill(&src_chain_id)
		if fill_err != nil {
			return
		}

		this.WithSrcChainId(src_chain_id)
	})
}

func Fuzz_Nosy_EmitCursorSrcChainIdDstChainIdShardIdHeightIndexKey_WithSrcChainIdDstChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this EmitCursorSrcChainIdDstChainIdShardIdHeightIndexKey
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var src_chain_id uint64
		fill_err = tp.Fill(&src_chain_id)
		if fill_err != nil {
			return
		}
		var dst_chain_id uint64
		fill_err = tp.Fill(&dst_chain_id)
		if fill_err != nil {
			return
		}

		this.WithSrcChainIdDstChainId(src_chain_id, dst_chain_id)
	})
}

func Fuzz_Nosy_EmitCursorSrcChainIdDstChainIdShardIdHeightIndexKey_WithSrcChainIdDstChainIdShardId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this EmitCursorSrcChainIdDstChainIdShardIdHeightIndexKey
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var src_chain_id uint64
		fill_err = tp.Fill(&src_chain_id)
		if fill_err != nil {
			return
		}
		var dst_chain_id uint64
		fill_err = tp.Fill(&dst_chain_id)
		if fill_err != nil {
			return
		}
		var shard_id uint64
		fill_err = tp.Fill(&shard_id)
		if fill_err != nil {
			return
		}

		this.WithSrcChainIdDstChainIdShardId(src_chain_id, dst_chain_id, shard_id)
	})
}

func Fuzz_Nosy_EmitCursorSrcChainIdDstChainIdShardIdHeightIndexKey_WithSrcChainIdDstChainIdShardIdHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this EmitCursorSrcChainIdDstChainIdShardIdHeightIndexKey
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var src_chain_id uint64
		fill_err = tp.Fill(&src_chain_id)
		if fill_err != nil {
			return
		}
		var dst_chain_id uint64
		fill_err = tp.Fill(&dst_chain_id)
		if fill_err != nil {
			return
		}
		var shard_id uint64
		fill_err = tp.Fill(&shard_id)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}

		this.WithSrcChainIdDstChainIdShardIdHeight(src_chain_id, dst_chain_id, shard_id, height)
	})
}

func Fuzz_Nosy_EmitCursorSrcChainIdDstChainIdShardIdHeightIndexKey_emitCursorIndexKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x EmitCursorSrcChainIdDstChainIdShardIdHeightIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.emitCursorIndexKey()
	})
}

func Fuzz_Nosy_EmitCursorSrcChainIdDstChainIdShardIdHeightIndexKey_id__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x EmitCursorSrcChainIdDstChainIdShardIdHeightIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.id()
	})
}

func Fuzz_Nosy_EmitCursorSrcChainIdDstChainIdShardIdHeightIndexKey_values__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x EmitCursorSrcChainIdDstChainIdShardIdHeightIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.values()
	})
}

// skipping Fuzz_Nosy_EmitCursorTable_Delete__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xmonitor/emitcache.EmitCursorTable

// skipping Fuzz_Nosy_EmitCursorTable_DeleteBy__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xmonitor/emitcache.EmitCursorTable

// skipping Fuzz_Nosy_EmitCursorTable_DeleteRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xmonitor/emitcache.EmitCursorTable

// skipping Fuzz_Nosy_EmitCursorTable_Get__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xmonitor/emitcache.EmitCursorTable

// skipping Fuzz_Nosy_EmitCursorTable_GetBySrcChainIdDstChainIdShardIdHeight__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xmonitor/emitcache.EmitCursorTable

// skipping Fuzz_Nosy_EmitCursorTable_Has__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xmonitor/emitcache.EmitCursorTable

// skipping Fuzz_Nosy_EmitCursorTable_HasBySrcChainIdDstChainIdShardIdHeight__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xmonitor/emitcache.EmitCursorTable

// skipping Fuzz_Nosy_EmitCursorTable_Insert__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xmonitor/emitcache.EmitCursorTable

// skipping Fuzz_Nosy_EmitCursorTable_InsertReturningId__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xmonitor/emitcache.EmitCursorTable

// skipping Fuzz_Nosy_EmitCursorTable_LastInsertedSequence__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xmonitor/emitcache.EmitCursorTable

// skipping Fuzz_Nosy_EmitCursorTable_List__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xmonitor/emitcache.EmitCursorTable

// skipping Fuzz_Nosy_EmitCursorTable_ListRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xmonitor/emitcache.EmitCursorTable

// skipping Fuzz_Nosy_EmitCursorTable_Save__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xmonitor/emitcache.EmitCursorTable

// skipping Fuzz_Nosy_EmitCursorTable_Update__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xmonitor/emitcache.EmitCursorTable

// skipping Fuzz_Nosy_EmitCursorTable_doNotImplement__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xmonitor/emitcache.EmitCursorTable

// skipping Fuzz_Nosy_EmitcursorStore_EmitCursorTable__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xmonitor/emitcache.EmitcursorStore

// skipping Fuzz_Nosy_EmitcursorStore_doNotImplement__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xmonitor/emitcache.EmitcursorStore

func Fuzz_Nosy_dbStoreService_OpenKVStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db dbStoreService
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}

		db.OpenKVStore(_x2)
	})
}

func Fuzz_Nosy_emitCursorTable_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this emitCursorTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var emitCursor *EmitCursor
		fill_err = tp.Fill(&emitCursor)
		if fill_err != nil {
			return
		}
		if emitCursor == nil {
			return
		}

		this.Delete(ctx, emitCursor)
	})
}

// skipping Fuzz_Nosy_emitCursorTable_DeleteBy__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xmonitor/emitcache.EmitCursorIndexKey

// skipping Fuzz_Nosy_emitCursorTable_DeleteRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xmonitor/emitcache.EmitCursorIndexKey

func Fuzz_Nosy_emitCursorTable_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this emitCursorTable
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

func Fuzz_Nosy_emitCursorTable_GetBySrcChainIdDstChainIdShardIdHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this emitCursorTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var src_chain_id uint64
		fill_err = tp.Fill(&src_chain_id)
		if fill_err != nil {
			return
		}
		var dst_chain_id uint64
		fill_err = tp.Fill(&dst_chain_id)
		if fill_err != nil {
			return
		}
		var shard_id uint64
		fill_err = tp.Fill(&shard_id)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}

		this.GetBySrcChainIdDstChainIdShardIdHeight(ctx, src_chain_id, dst_chain_id, shard_id, height)
	})
}

func Fuzz_Nosy_emitCursorTable_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this emitCursorTable
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

func Fuzz_Nosy_emitCursorTable_HasBySrcChainIdDstChainIdShardIdHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this emitCursorTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var src_chain_id uint64
		fill_err = tp.Fill(&src_chain_id)
		if fill_err != nil {
			return
		}
		var dst_chain_id uint64
		fill_err = tp.Fill(&dst_chain_id)
		if fill_err != nil {
			return
		}
		var shard_id uint64
		fill_err = tp.Fill(&shard_id)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}

		this.HasBySrcChainIdDstChainIdShardIdHeight(ctx, src_chain_id, dst_chain_id, shard_id, height)
	})
}

func Fuzz_Nosy_emitCursorTable_Insert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this emitCursorTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var emitCursor *EmitCursor
		fill_err = tp.Fill(&emitCursor)
		if fill_err != nil {
			return
		}
		if emitCursor == nil {
			return
		}

		this.Insert(ctx, emitCursor)
	})
}

func Fuzz_Nosy_emitCursorTable_InsertReturningId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this emitCursorTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var emitCursor *EmitCursor
		fill_err = tp.Fill(&emitCursor)
		if fill_err != nil {
			return
		}
		if emitCursor == nil {
			return
		}

		this.InsertReturningId(ctx, emitCursor)
	})
}

func Fuzz_Nosy_emitCursorTable_LastInsertedSequence__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this emitCursorTable
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

// skipping Fuzz_Nosy_emitCursorTable_List__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xmonitor/emitcache.EmitCursorIndexKey

// skipping Fuzz_Nosy_emitCursorTable_ListRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xmonitor/emitcache.EmitCursorIndexKey

func Fuzz_Nosy_emitCursorTable_Save__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this emitCursorTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var emitCursor *EmitCursor
		fill_err = tp.Fill(&emitCursor)
		if fill_err != nil {
			return
		}
		if emitCursor == nil {
			return
		}

		this.Save(ctx, emitCursor)
	})
}

func Fuzz_Nosy_emitCursorTable_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this emitCursorTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var emitCursor *EmitCursor
		fill_err = tp.Fill(&emitCursor)
		if fill_err != nil {
			return
		}
		if emitCursor == nil {
			return
		}

		this.Update(ctx, emitCursor)
	})
}

func Fuzz_Nosy_emitCursorTable_doNotImplement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this emitCursorTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}

		this.doNotImplement()
	})
}

func Fuzz_Nosy_emitcursorStore_EmitCursorTable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x emitcursorStore
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.EmitCursorTable()
	})
}

func Fuzz_Nosy_emitcursorStore_doNotImplement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 emitcursorStore
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.doNotImplement()
	})
}
