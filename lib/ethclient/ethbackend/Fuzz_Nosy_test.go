package ethbackend

import (
	"context"
	"crypto/ecdsa"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/omni-network/omni/lib/fireblocks"
	"github.com/omni-network/omni/lib/txmgr"
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

func Fuzz_Nosy_Backend_AddAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Backend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var privkey *ecdsa.PrivateKey
		fill_err = tp.Fill(&privkey)
		if fill_err != nil {
			return
		}
		if b == nil || privkey == nil {
			return
		}

		b.AddAccount(privkey)
	})
}

func Fuzz_Nosy_Backend_BindOpts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Backend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var from common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.BindOpts(ctx, from)
	})
}

func Fuzz_Nosy_Backend_Chain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Backend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Chain()
	})
}

func Fuzz_Nosy_Backend_EnsureSynced__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Backend
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

		b.EnsureSynced(ctx)
	})
}

func Fuzz_Nosy_Backend_PublicKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Backend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var from common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.PublicKey(from)
	})
}

func Fuzz_Nosy_Backend_Send__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Backend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var from common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var candidate txmgr.TxCandidate
		fill_err = tp.Fill(&candidate)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Send(ctx, from, candidate)
	})
}

func Fuzz_Nosy_Backend_SendTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Backend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var in *types.Transaction
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		if b == nil || in == nil {
			return
		}

		b.SendTransaction(ctx, in)
	})
}

func Fuzz_Nosy_Backend_Sign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Backend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var from common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var input [32]byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Sign(ctx, from, input)
	})
}

func Fuzz_Nosy_Backend_WaitMined__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Backend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if b == nil || tx == nil {
			return
		}

		b.WaitMined(ctx, tx)
	})
}

func Fuzz_Nosy_Waiter_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *Waiter
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var chainID uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if w == nil || tx == nil {
			return
		}

		w.Add(chainID, tx)
	})
}

func Fuzz_Nosy_Waiter_Wait__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *Waiter
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

		w.Wait(ctx)
	})
}

func Fuzz_Nosy_Backends_All__(f *testing.F) {
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
		var testnet types.Testnet
		fill_err = tp.Fill(&testnet)
		if fill_err != nil {
			return
		}
		var fireCl fireblocks.Client
		fill_err = tp.Fill(&fireCl)
		if fill_err != nil {
			return
		}

		b, err := NewFireBackends(ctx, testnet, fireCl)
		if err != nil {
			return
		}
		b.All()
	})
}

func Fuzz_Nosy_Backends_Backend__(f *testing.F) {
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
		var testnet types.Testnet
		fill_err = tp.Fill(&testnet)
		if fill_err != nil {
			return
		}
		var fireCl fireblocks.Client
		fill_err = tp.Fill(&fireCl)
		if fill_err != nil {
			return
		}
		var sourceChainID uint64
		fill_err = tp.Fill(&sourceChainID)
		if fill_err != nil {
			return
		}

		b, err := NewFireBackends(ctx, testnet, fireCl)
		if err != nil {
			return
		}
		b.Backend(sourceChainID)
	})
}

func Fuzz_Nosy_Backends_BindOpts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var testnet types.Testnet
		fill_err = tp.Fill(&testnet)
		if fill_err != nil {
			return
		}
		var fireCl fireblocks.Client
		fill_err = tp.Fill(&fireCl)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var sourceChainID uint64
		fill_err = tp.Fill(&sourceChainID)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}

		b, err := NewFireBackends(c1, testnet, fireCl)
		if err != nil {
			return
		}
		b.BindOpts(c4, sourceChainID, addr)
	})
}

func Fuzz_Nosy_Backends_NewWaiter__(f *testing.F) {
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
		var testnet types.Testnet
		fill_err = tp.Fill(&testnet)
		if fill_err != nil {
			return
		}
		var fireCl fireblocks.Client
		fill_err = tp.Fill(&fireCl)
		if fill_err != nil {
			return
		}

		b, err := NewFireBackends(ctx, testnet, fireCl)
		if err != nil {
			return
		}
		b.NewWaiter()
	})
}

func Fuzz_Nosy_Backends_RPCClients__(f *testing.F) {
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
		var testnet types.Testnet
		fill_err = tp.Fill(&testnet)
		if fill_err != nil {
			return
		}
		var fireCl fireblocks.Client
		fill_err = tp.Fill(&fireCl)
		if fill_err != nil {
			return
		}

		b, err := NewFireBackends(ctx, testnet, fireCl)
		if err != nil {
			return
		}
		b.RPCClients()
	})
}

func Fuzz_Nosy_backendStubSigner_ChainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 backendStubSigner
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.ChainID()
	})
}

func Fuzz_Nosy_backendStubSigner_Sender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 backendStubSigner
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		_x1.Sender(tx)
	})
}

func Fuzz_Nosy_backendStubSigner_SignatureValues__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 backendStubSigner
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 *types.Transaction
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var sig []byte
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if _x2 == nil {
			return
		}

		_x1.SignatureValues(_x2, sig)
	})
}
