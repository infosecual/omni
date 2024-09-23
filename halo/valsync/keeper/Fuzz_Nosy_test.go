package keeper

import (
	"context"
	"testing"

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

func Fuzz_Nosy_Keeper_ActiveSetByHeight__(f *testing.F) {
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
		if k == nil {
			return
		}

		k.ActiveSetByHeight(ctx, height)
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

func Fuzz_Nosy_Keeper_InsertGenesisSet__(f *testing.F) {
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

		k.InsertGenesisSet(ctx)
	})
}

func Fuzz_Nosy_Keeper_ValidatorSet__(f *testing.F) {
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
		var req *types.ValidatorSetRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if k == nil || req == nil {
			return
		}

		k.ValidatorSet(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_activeSetByHeight__(f *testing.F) {
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
		if k == nil {
			return
		}

		k.activeSetByHeight(ctx, height)
	})
}

func Fuzz_Nosy_Keeper_insertValidatorSet__(f *testing.F) {
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
		var vals []*Validator
		fill_err = tp.Fill(&vals)
		if fill_err != nil {
			return
		}
		var isGenesis bool
		fill_err = tp.Fill(&isGenesis)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		k.insertValidatorSet(ctx, vals, isGenesis)
	})
}

func Fuzz_Nosy_Keeper_maybeInitSubscriber__(f *testing.F) {
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

		k.maybeInitSubscriber(ctx)
	})
}

func Fuzz_Nosy_Keeper_maybeStoreValidatorUpdates__(f *testing.F) {
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
		var updates []types.ValidatorUpdate
		fill_err = tp.Fill(&updates)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		k.maybeStoreValidatorUpdates(ctx, updates)
	})
}

func Fuzz_Nosy_Keeper_nextUnattestedSet__(f *testing.F) {
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

		k.nextUnattestedSet(ctx)
	})
}

func Fuzz_Nosy_Keeper_processAttested__(f *testing.F) {
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

		k.processAttested(ctx)
	})
}

func Fuzz_Nosy_Validator_Address__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *Validator
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.Address()
	})
}

func Fuzz_Nosy_Validator_CometPubkey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *Validator
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.CometPubkey()
	})
}

func Fuzz_Nosy_Validator_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Validator
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

func Fuzz_Nosy_Validator_GetId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Validator
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

func Fuzz_Nosy_Validator_GetPower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Validator
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetPower()
	})
}

func Fuzz_Nosy_Validator_GetPubKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Validator
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetPubKey()
	})
}

func Fuzz_Nosy_Validator_GetUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Validator
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetUpdated()
	})
}

func Fuzz_Nosy_Validator_GetValsetId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Validator
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetValsetId()
	})
}

func Fuzz_Nosy_Validator_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Validator
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

func Fuzz_Nosy_Validator_ProtoReflect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Validator
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

func Fuzz_Nosy_Validator_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Validator
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

func Fuzz_Nosy_Validator_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Validator
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

func Fuzz_Nosy_Validator_ValidatorUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *Validator
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.ValidatorUpdate()
	})
}

func Fuzz_Nosy_ValidatorSet_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *ValidatorSet
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

func Fuzz_Nosy_ValidatorSet_GetActivatedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *ValidatorSet
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetActivatedHeight()
	})
}

func Fuzz_Nosy_ValidatorSet_GetAttestOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *ValidatorSet
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

func Fuzz_Nosy_ValidatorSet_GetAttested__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *ValidatorSet
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetAttested()
	})
}

func Fuzz_Nosy_ValidatorSet_GetCreatedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *ValidatorSet
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

func Fuzz_Nosy_ValidatorSet_GetId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *ValidatorSet
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

func Fuzz_Nosy_ValidatorSet_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *ValidatorSet
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

func Fuzz_Nosy_ValidatorSet_ProtoReflect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *ValidatorSet
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

func Fuzz_Nosy_ValidatorSet_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *ValidatorSet
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

func Fuzz_Nosy_ValidatorSet_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *ValidatorSet
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

func Fuzz_Nosy_ValidatorIdIndexKey_WithId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this ValidatorIdIndexKey
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

func Fuzz_Nosy_ValidatorIdIndexKey_id__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x ValidatorIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.id()
	})
}

func Fuzz_Nosy_ValidatorIdIndexKey_validatorIndexKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x ValidatorIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.validatorIndexKey()
	})
}

func Fuzz_Nosy_ValidatorIdIndexKey_values__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x ValidatorIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.values()
	})
}

// skipping Fuzz_Nosy_ValidatorIndexKey_id__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorIndexKey

// skipping Fuzz_Nosy_ValidatorIndexKey_validatorIndexKey__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorIndexKey

// skipping Fuzz_Nosy_ValidatorIndexKey_values__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorIndexKey

func Fuzz_Nosy_ValidatorIterator_Value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i ValidatorIterator
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		i.Value()
	})
}

func Fuzz_Nosy_ValidatorSetAttestedCreatedHeightIndexKey_WithAttested__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this ValidatorSetAttestedCreatedHeightIndexKey
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var attested bool
		fill_err = tp.Fill(&attested)
		if fill_err != nil {
			return
		}

		this.WithAttested(attested)
	})
}

func Fuzz_Nosy_ValidatorSetAttestedCreatedHeightIndexKey_WithAttestedCreatedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this ValidatorSetAttestedCreatedHeightIndexKey
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var attested bool
		fill_err = tp.Fill(&attested)
		if fill_err != nil {
			return
		}
		var created_height uint64
		fill_err = tp.Fill(&created_height)
		if fill_err != nil {
			return
		}

		this.WithAttestedCreatedHeight(attested, created_height)
	})
}

func Fuzz_Nosy_ValidatorSetAttestedCreatedHeightIndexKey_id__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x ValidatorSetAttestedCreatedHeightIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.id()
	})
}

func Fuzz_Nosy_ValidatorSetAttestedCreatedHeightIndexKey_validatorSetIndexKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x ValidatorSetAttestedCreatedHeightIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.validatorSetIndexKey()
	})
}

func Fuzz_Nosy_ValidatorSetAttestedCreatedHeightIndexKey_values__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x ValidatorSetAttestedCreatedHeightIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.values()
	})
}

func Fuzz_Nosy_ValidatorSetIdIndexKey_WithId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this ValidatorSetIdIndexKey
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

func Fuzz_Nosy_ValidatorSetIdIndexKey_id__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x ValidatorSetIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.id()
	})
}

func Fuzz_Nosy_ValidatorSetIdIndexKey_validatorSetIndexKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x ValidatorSetIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.validatorSetIndexKey()
	})
}

func Fuzz_Nosy_ValidatorSetIdIndexKey_values__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x ValidatorSetIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.values()
	})
}

// skipping Fuzz_Nosy_ValidatorSetIndexKey_id__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorSetIndexKey

// skipping Fuzz_Nosy_ValidatorSetIndexKey_validatorSetIndexKey__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorSetIndexKey

// skipping Fuzz_Nosy_ValidatorSetIndexKey_values__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorSetIndexKey

func Fuzz_Nosy_ValidatorSetIterator_Value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i ValidatorSetIterator
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		i.Value()
	})
}

// skipping Fuzz_Nosy_ValidatorSetTable_Delete__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorSetTable

// skipping Fuzz_Nosy_ValidatorSetTable_DeleteBy__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorSetTable

// skipping Fuzz_Nosy_ValidatorSetTable_DeleteRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorSetTable

// skipping Fuzz_Nosy_ValidatorSetTable_Get__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorSetTable

// skipping Fuzz_Nosy_ValidatorSetTable_GetByAttestedCreatedHeight__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorSetTable

// skipping Fuzz_Nosy_ValidatorSetTable_Has__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorSetTable

// skipping Fuzz_Nosy_ValidatorSetTable_HasByAttestedCreatedHeight__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorSetTable

// skipping Fuzz_Nosy_ValidatorSetTable_Insert__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorSetTable

// skipping Fuzz_Nosy_ValidatorSetTable_InsertReturningId__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorSetTable

// skipping Fuzz_Nosy_ValidatorSetTable_LastInsertedSequence__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorSetTable

// skipping Fuzz_Nosy_ValidatorSetTable_List__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorSetTable

// skipping Fuzz_Nosy_ValidatorSetTable_ListRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorSetTable

// skipping Fuzz_Nosy_ValidatorSetTable_Save__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorSetTable

// skipping Fuzz_Nosy_ValidatorSetTable_Update__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorSetTable

// skipping Fuzz_Nosy_ValidatorSetTable_doNotImplement__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorSetTable

// skipping Fuzz_Nosy_ValidatorTable_Delete__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorTable

// skipping Fuzz_Nosy_ValidatorTable_DeleteBy__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorTable

// skipping Fuzz_Nosy_ValidatorTable_DeleteRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorTable

// skipping Fuzz_Nosy_ValidatorTable_Get__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorTable

// skipping Fuzz_Nosy_ValidatorTable_Has__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorTable

// skipping Fuzz_Nosy_ValidatorTable_Insert__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorTable

// skipping Fuzz_Nosy_ValidatorTable_InsertReturningId__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorTable

// skipping Fuzz_Nosy_ValidatorTable_LastInsertedSequence__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorTable

// skipping Fuzz_Nosy_ValidatorTable_List__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorTable

// skipping Fuzz_Nosy_ValidatorTable_ListRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorTable

// skipping Fuzz_Nosy_ValidatorTable_Save__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorTable

// skipping Fuzz_Nosy_ValidatorTable_Update__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorTable

// skipping Fuzz_Nosy_ValidatorTable_doNotImplement__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorTable

func Fuzz_Nosy_ValidatorValsetIdIndexKey_WithValsetId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this ValidatorValsetIdIndexKey
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var valset_id uint64
		fill_err = tp.Fill(&valset_id)
		if fill_err != nil {
			return
		}

		this.WithValsetId(valset_id)
	})
}

func Fuzz_Nosy_ValidatorValsetIdIndexKey_id__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x ValidatorValsetIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.id()
	})
}

func Fuzz_Nosy_ValidatorValsetIdIndexKey_validatorIndexKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x ValidatorValsetIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.validatorIndexKey()
	})
}

func Fuzz_Nosy_ValidatorValsetIdIndexKey_values__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x ValidatorValsetIdIndexKey
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.values()
	})
}

// skipping Fuzz_Nosy_ValsyncStore_ValidatorSetTable__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValsyncStore

// skipping Fuzz_Nosy_ValsyncStore_ValidatorTable__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValsyncStore

// skipping Fuzz_Nosy_ValsyncStore_doNotImplement__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValsyncStore

func Fuzz_Nosy_validatorSetTable_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this validatorSetTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var validatorSet *ValidatorSet
		fill_err = tp.Fill(&validatorSet)
		if fill_err != nil {
			return
		}
		if validatorSet == nil {
			return
		}

		this.Delete(ctx, validatorSet)
	})
}

// skipping Fuzz_Nosy_validatorSetTable_DeleteBy__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorSetIndexKey

// skipping Fuzz_Nosy_validatorSetTable_DeleteRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorSetIndexKey

func Fuzz_Nosy_validatorSetTable_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this validatorSetTable
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

func Fuzz_Nosy_validatorSetTable_GetByAttestedCreatedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this validatorSetTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var attested bool
		fill_err = tp.Fill(&attested)
		if fill_err != nil {
			return
		}
		var created_height uint64
		fill_err = tp.Fill(&created_height)
		if fill_err != nil {
			return
		}

		this.GetByAttestedCreatedHeight(ctx, attested, created_height)
	})
}

func Fuzz_Nosy_validatorSetTable_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this validatorSetTable
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

func Fuzz_Nosy_validatorSetTable_HasByAttestedCreatedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this validatorSetTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var attested bool
		fill_err = tp.Fill(&attested)
		if fill_err != nil {
			return
		}
		var created_height uint64
		fill_err = tp.Fill(&created_height)
		if fill_err != nil {
			return
		}

		this.HasByAttestedCreatedHeight(ctx, attested, created_height)
	})
}

func Fuzz_Nosy_validatorSetTable_Insert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this validatorSetTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var validatorSet *ValidatorSet
		fill_err = tp.Fill(&validatorSet)
		if fill_err != nil {
			return
		}
		if validatorSet == nil {
			return
		}

		this.Insert(ctx, validatorSet)
	})
}

func Fuzz_Nosy_validatorSetTable_InsertReturningId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this validatorSetTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var validatorSet *ValidatorSet
		fill_err = tp.Fill(&validatorSet)
		if fill_err != nil {
			return
		}
		if validatorSet == nil {
			return
		}

		this.InsertReturningId(ctx, validatorSet)
	})
}

func Fuzz_Nosy_validatorSetTable_LastInsertedSequence__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this validatorSetTable
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

// skipping Fuzz_Nosy_validatorSetTable_List__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorSetIndexKey

// skipping Fuzz_Nosy_validatorSetTable_ListRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorSetIndexKey

func Fuzz_Nosy_validatorSetTable_Save__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this validatorSetTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var validatorSet *ValidatorSet
		fill_err = tp.Fill(&validatorSet)
		if fill_err != nil {
			return
		}
		if validatorSet == nil {
			return
		}

		this.Save(ctx, validatorSet)
	})
}

func Fuzz_Nosy_validatorSetTable_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this validatorSetTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var validatorSet *ValidatorSet
		fill_err = tp.Fill(&validatorSet)
		if fill_err != nil {
			return
		}
		if validatorSet == nil {
			return
		}

		this.Update(ctx, validatorSet)
	})
}

func Fuzz_Nosy_validatorSetTable_doNotImplement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this validatorSetTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}

		this.doNotImplement()
	})
}

func Fuzz_Nosy_validatorTable_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this validatorTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var validator *Validator
		fill_err = tp.Fill(&validator)
		if fill_err != nil {
			return
		}
		if validator == nil {
			return
		}

		this.Delete(ctx, validator)
	})
}

// skipping Fuzz_Nosy_validatorTable_DeleteBy__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorIndexKey

// skipping Fuzz_Nosy_validatorTable_DeleteRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorIndexKey

func Fuzz_Nosy_validatorTable_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this validatorTable
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

func Fuzz_Nosy_validatorTable_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this validatorTable
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

func Fuzz_Nosy_validatorTable_Insert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this validatorTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var validator *Validator
		fill_err = tp.Fill(&validator)
		if fill_err != nil {
			return
		}
		if validator == nil {
			return
		}

		this.Insert(ctx, validator)
	})
}

func Fuzz_Nosy_validatorTable_InsertReturningId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this validatorTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var validator *Validator
		fill_err = tp.Fill(&validator)
		if fill_err != nil {
			return
		}
		if validator == nil {
			return
		}

		this.InsertReturningId(ctx, validator)
	})
}

func Fuzz_Nosy_validatorTable_LastInsertedSequence__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this validatorTable
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

// skipping Fuzz_Nosy_validatorTable_List__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorIndexKey

// skipping Fuzz_Nosy_validatorTable_ListRange__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/valsync/keeper.ValidatorIndexKey

func Fuzz_Nosy_validatorTable_Save__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this validatorTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var validator *Validator
		fill_err = tp.Fill(&validator)
		if fill_err != nil {
			return
		}
		if validator == nil {
			return
		}

		this.Save(ctx, validator)
	})
}

func Fuzz_Nosy_validatorTable_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this validatorTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var validator *Validator
		fill_err = tp.Fill(&validator)
		if fill_err != nil {
			return
		}
		if validator == nil {
			return
		}

		this.Update(ctx, validator)
	})
}

func Fuzz_Nosy_validatorTable_doNotImplement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var this validatorTable
		fill_err = tp.Fill(&this)
		if fill_err != nil {
			return
		}

		this.doNotImplement()
	})
}

func Fuzz_Nosy_valsyncStore_ValidatorSetTable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x valsyncStore
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.ValidatorSetTable()
	})
}

func Fuzz_Nosy_valsyncStore_ValidatorTable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x valsyncStore
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.ValidatorTable()
	})
}

func Fuzz_Nosy_valsyncStore_doNotImplement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 valsyncStore
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.doNotImplement()
	})
}

func Fuzz_Nosy_mergeValidatorSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var valset []types.Validator
		fill_err = tp.Fill(&valset)
		if fill_err != nil {
			return
		}
		var updates []types.ValidatorUpdate
		fill_err = tp.Fill(&updates)
		if fill_err != nil {
			return
		}

		mergeValidatorSet(valset, updates)
	})
}
