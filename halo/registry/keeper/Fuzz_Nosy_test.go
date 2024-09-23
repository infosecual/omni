package keeper

import (
	"context"
	"testing"

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

func Fuzz_Nosy_Network_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Network
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

func Fuzz_Nosy_Network_GetCreatedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Network
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

func Fuzz_Nosy_Network_GetId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Network
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

func Fuzz_Nosy_Network_GetPortals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Network
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetPortals()
	})
}

func Fuzz_Nosy_Network_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Network
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

func Fuzz_Nosy_Network_ProtoReflect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Network
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

func Fuzz_Nosy_Network_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Network
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

func Fuzz_Nosy_Network_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Network
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

func Fuzz_Nosy_Portal_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Portal
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

func Fuzz_Nosy_Portal_GetAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Portal
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetAddress()
	})
}

func Fuzz_Nosy_Portal_GetAttestInterval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Portal
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetAttestInterval()
	})
}

func Fuzz_Nosy_Portal_GetBlockPeriodMs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Portal
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetBlockPeriodMs()
	})
}

func Fuzz_Nosy_Portal_GetChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Portal
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

func Fuzz_Nosy_Portal_GetDeployHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Portal
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetDeployHeight()
	})
}

func Fuzz_Nosy_Portal_GetName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Portal
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetName()
	})
}

func Fuzz_Nosy_Portal_GetShardIds__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Portal
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetShardIds()
	})
}

func Fuzz_Nosy_Portal_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Portal
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

func Fuzz_Nosy_Portal_ProtoReflect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Portal
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

func Fuzz_Nosy_Portal_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Portal
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

func Fuzz_Nosy_Portal_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Portal
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

func Fuzz_Nosy_Portal_Verify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Portal
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Verify()
	})
}

func Fuzz_Nosy_cache_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *cache
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Get()
	})
}

func Fuzz_Nosy_cache_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *cache
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var network *Network
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		if c == nil || network == nil {
			return
		}

		c.Set(network)
	})
}

func Fuzz_Nosy_Keeper_Addresses__(f *testing.F) {
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

		k.Addresses()
	})
}

func Fuzz_Nosy_Keeper_ConfLevels__(f *testing.F) {
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

		k.ConfLevels(ctx)
	})
}

func Fuzz_Nosy_Keeper_Deliver__(f *testing.F) {
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
		var _x3 common.Hash
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		var elog *types.EVMEvent
		fill_err = tp.Fill(&elog)
		if fill_err != nil {
			return
		}
		if elog == nil {
			return
		}

		k.Deliver(ctx, _x3, elog)
	})
}

func Fuzz_Nosy_Keeper_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 Keeper
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_Keeper_Network__(f *testing.F) {
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
		var req *types.NetworkRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.Network(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_Prepare__(f *testing.F) {
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
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}

		k.Prepare(ctx, blockHash)
	})
}

func Fuzz_Nosy_Keeper_SupportedChain__(f *testing.F) {
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
		var chainID uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}

		k.SupportedChain(ctx, chainID)
	})
}

func Fuzz_Nosy_Keeper_addPortal__(f *testing.F) {
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
		var portal *Portal
		fill_err = tp.Fill(&portal)
		if fill_err != nil {
			return
		}
		if portal == nil {
			return
		}

		k.addPortal(ctx, portal)
	})
}

func Fuzz_Nosy_Keeper_getLatestPortals__(f *testing.F) {
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

		k.getLatestPortals(ctx)
	})
}

func Fuzz_Nosy_Keeper_getOrCreateNetwork__(f *testing.F) {
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

		k.getOrCreateNetwork(ctx)
	})
}

func Fuzz_Nosy_Keeper_updateNetwork__(f *testing.F) {
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
		var network *Network
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		if network == nil {
			return
		}

		k.updateNetwork(ctx, network)
	})
}

func Fuzz_Nosy_NetworkIdIndexKey_WithId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this NetworkIdIndexKey
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

func Fuzz_Nosy_NetworkIdIndexKey_id__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x NetworkIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.id()
	})
}

func Fuzz_Nosy_NetworkIdIndexKey_networkIndexKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x NetworkIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.networkIndexKey()
	})
}

func Fuzz_Nosy_NetworkIdIndexKey_values__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x NetworkIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.values()
	})
}

// skipping Fuzz_Nosy_NetworkIndexKey_id__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/registry/keeper.NetworkIndexKey

// skipping Fuzz_Nosy_NetworkIndexKey_networkIndexKey__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/registry/keeper.NetworkIndexKey

// skipping Fuzz_Nosy_NetworkIndexKey_values__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/registry/keeper.NetworkIndexKey

func Fuzz_Nosy_NetworkIterator_Value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i NetworkIterator
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		i.Value()
	})
}

// skipping Fuzz_Nosy_NetworkTable_Delete__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/registry/keeper.NetworkTable

// skipping Fuzz_Nosy_NetworkTable_DeleteBy__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/registry/keeper.NetworkTable

// skipping Fuzz_Nosy_NetworkTable_DeleteRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/registry/keeper.NetworkTable

// skipping Fuzz_Nosy_NetworkTable_Get__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/registry/keeper.NetworkTable

// skipping Fuzz_Nosy_NetworkTable_Has__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/registry/keeper.NetworkTable

// skipping Fuzz_Nosy_NetworkTable_Insert__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/registry/keeper.NetworkTable

// skipping Fuzz_Nosy_NetworkTable_InsertReturningId__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/registry/keeper.NetworkTable

// skipping Fuzz_Nosy_NetworkTable_LastInsertedSequence__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/registry/keeper.NetworkTable

// skipping Fuzz_Nosy_NetworkTable_List__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/registry/keeper.NetworkTable

// skipping Fuzz_Nosy_NetworkTable_ListRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/registry/keeper.NetworkTable

// skipping Fuzz_Nosy_NetworkTable_Save__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/registry/keeper.NetworkTable

// skipping Fuzz_Nosy_NetworkTable_Update__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/registry/keeper.NetworkTable

// skipping Fuzz_Nosy_NetworkTable_doNotImplement__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/registry/keeper.NetworkTable

// skipping Fuzz_Nosy_RegistryStore_NetworkTable__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/registry/keeper.RegistryStore

// skipping Fuzz_Nosy_RegistryStore_doNotImplement__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/registry/keeper.RegistryStore

func Fuzz_Nosy_networkTable_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this networkTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var network *Network
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		if network == nil {
			return
		}

		this.Delete(ctx, network)
	})
}

// skipping Fuzz_Nosy_networkTable_DeleteBy__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/registry/keeper.NetworkIndexKey

// skipping Fuzz_Nosy_networkTable_DeleteRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/registry/keeper.NetworkIndexKey

func Fuzz_Nosy_networkTable_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this networkTable
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

func Fuzz_Nosy_networkTable_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this networkTable
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

func Fuzz_Nosy_networkTable_Insert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this networkTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var network *Network
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		if network == nil {
			return
		}

		this.Insert(ctx, network)
	})
}

func Fuzz_Nosy_networkTable_InsertReturningId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this networkTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var network *Network
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		if network == nil {
			return
		}

		this.InsertReturningId(ctx, network)
	})
}

func Fuzz_Nosy_networkTable_LastInsertedSequence__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this networkTable
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

// skipping Fuzz_Nosy_networkTable_List__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/registry/keeper.NetworkIndexKey

// skipping Fuzz_Nosy_networkTable_ListRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/registry/keeper.NetworkIndexKey

func Fuzz_Nosy_networkTable_Save__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this networkTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var network *Network
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		if network == nil {
			return
		}

		this.Save(ctx, network)
	})
}

func Fuzz_Nosy_networkTable_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this networkTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var network *Network
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		if network == nil {
			return
		}

		this.Update(ctx, network)
	})
}

func Fuzz_Nosy_networkTable_doNotImplement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this networkTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}

		this.doNotImplement()
	})
}

func Fuzz_Nosy_registryStore_NetworkTable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x registryStore
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.NetworkTable()
	})
}

func Fuzz_Nosy_registryStore_doNotImplement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 registryStore
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.doNotImplement()
	})
}

func Fuzz_Nosy_mergePortal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var existing []*Portal
		fill_err = tp.Fill(&existing)
		if fill_err != nil {
			return
		}
		var portal *Portal
		fill_err = tp.Fill(&portal)
		if fill_err != nil {
			return
		}
		if portal == nil {
			return
		}

		mergePortal(existing, portal)
	})
}

func Fuzz_Nosy_newShards__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var existing []uint64
		fill_err = tp.Fill(&existing)
		if fill_err != nil {
			return
		}
		var shards []uint64
		fill_err = tp.Fill(&shards)
		if fill_err != nil {
			return
		}

		newShards(existing, shards)
	})
}

func Fuzz_Nosy_shardLabels__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var shardIDs []uint64
		fill_err = tp.Fill(&shardIDs)
		if fill_err != nil {
			return
		}

		shardLabels(shardIDs)
	})
}
