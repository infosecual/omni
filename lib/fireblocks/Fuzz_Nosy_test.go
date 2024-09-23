package fireblocks

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

func Fuzz_Nosy_accountCache_Clone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var init map[common.Address]uint64
		fill_err = tp.Fill(&init)
		if fill_err != nil {
			return
		}

		c := newAccountCache(init)
		c.Clone()
	})
}

func Fuzz_Nosy_accountCache_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var init map[common.Address]uint64
		fill_err = tp.Fill(&init)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}

		c := newAccountCache(init)
		c.Get(addr)
	})
}

// skipping Fuzz_Nosy_accountCache_MaybePopulate__ because parameters include func, chan, or unsupported interface: func(context.Context) (map[github.com/ethereum/go-ethereum/common.Address]uint64, error)

func Fuzz_Nosy_accountCache_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var init map[common.Address]uint64
		fill_err = tp.Fill(&init)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		c := newAccountCache(init)
		c.Set(addr, id)
	})
}

func Fuzz_Nosy_Client_Accounts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		c.Accounts(ctx)
	})
}

func Fuzz_Nosy_Client_GetPublicKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account uint64
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}

		c.GetPublicKey(ctx, account)
	})
}

func Fuzz_Nosy_Client_GetSupportedAssets__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		c.GetSupportedAssets(ctx)
	})
}

func Fuzz_Nosy_Client_Sign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var digest common.Hash
		fill_err = tp.Fill(&digest)
		if fill_err != nil {
			return
		}
		var signer common.Address
		fill_err = tp.Fill(&signer)
		if fill_err != nil {
			return
		}

		c.Sign(ctx, digest, signer)
	})
}

// skipping Fuzz_Nosy_Client_authHeaders__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_Client_createRawSignTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account uint64
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var digest common.Hash
		fill_err = tp.Fill(&digest)
		if fill_err != nil {
			return
		}

		c.createRawSignTransaction(ctx, account, digest)
	})
}

func Fuzz_Nosy_Client_getAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}

		c.getAccount(ctx, addr)
	})
}

func Fuzz_Nosy_Client_getAssetID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.getAssetID()
	})
}

func Fuzz_Nosy_Client_getTransactionByID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var transactionID string
		fill_err = tp.Fill(&transactionID)
		if fill_err != nil {
			return
		}

		c.getTransactionByID(ctx, transactionID)
	})
}

func Fuzz_Nosy_Client_maybeGetSignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var txID string
		fill_err = tp.Fill(&txID)
		if fill_err != nil {
			return
		}
		var digest common.Hash
		fill_err = tp.Fill(&digest)
		if fill_err != nil {
			return
		}
		var signer common.Address
		fill_err = tp.Fill(&signer)
		if fill_err != nil {
			return
		}

		c.maybeGetSignature(ctx, txID, digest, signer)
	})
}

func Fuzz_Nosy_Client_newRawSignRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var account uint64
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var digest common.Hash
		fill_err = tp.Fill(&digest)
		if fill_err != nil {
			return
		}

		c.newRawSignRequest(account, digest)
	})
}

func Fuzz_Nosy_Client_pubkeyEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var account uint64
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}

		c.pubkeyEndpoint(account)
	})
}

func Fuzz_Nosy_Client_queryAccounts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		c.queryAccounts(ctx)
	})
}

// skipping Fuzz_Nosy_Client_token__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_Status_Completed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s Status
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.Completed()
	})
}

func Fuzz_Nosy_Status_Failed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s Status
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.Failed()
	})
}

// skipping Fuzz_Nosy_jsonHTTP_Send__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_transaction_Pubkey0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 transaction
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		t1.Pubkey0()
	})
}

func Fuzz_Nosy_transaction_Sig0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 transaction
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		t1.Sig0()
	})
}
