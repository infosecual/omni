package mock

import (
	context "context"
	big "math/big"
	"testing"

	ethereum "github.com/ethereum/go-ethereum"
	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	"github.com/omni-network/omni/lib/ethclient"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	gomock "go.uber.org/mock/gomock"
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

func Fuzz_Nosy_MockClient_Address__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.Address()
	})
}

func Fuzz_Nosy_MockClient_BalanceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if ctrl == nil || blockNumber == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.BalanceAt(ctx, account, blockNumber)
	})
}

func Fuzz_Nosy_MockClient_BlockByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.BlockByHash(ctx, hash)
	})
}

func Fuzz_Nosy_MockClient_BlockByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var number *big.Int
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if ctrl == nil || number == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.BlockByNumber(ctx, number)
	})
}

func Fuzz_Nosy_MockClient_BlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.BlockNumber(ctx)
	})
}

func Fuzz_Nosy_MockClient_CallContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var call ethereum.CallMsg
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if ctrl == nil || blockNumber == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.CallContract(ctx, call, blockNumber)
	})
}

func Fuzz_Nosy_MockClient_ChainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.ChainID(ctx)
	})
}

func Fuzz_Nosy_MockClient_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.Close()
	})
}

func Fuzz_Nosy_MockClient_CodeAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if ctrl == nil || blockNumber == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.CodeAt(ctx, account, blockNumber)
	})
}

func Fuzz_Nosy_MockClient_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.EXPECT()
	})
}

func Fuzz_Nosy_MockClient_EstimateGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var call ethereum.CallMsg
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.EstimateGas(ctx, call)
	})
}

func Fuzz_Nosy_MockClient_EtherBalanceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
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
		if ctrl == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.EtherBalanceAt(ctx, addr)
	})
}

func Fuzz_Nosy_MockClient_FilterLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var q ethereum.FilterQuery
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.FilterLogs(ctx, q)
	})
}

func Fuzz_Nosy_MockClient_HeaderByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.HeaderByHash(ctx, hash)
	})
}

func Fuzz_Nosy_MockClient_HeaderByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var number *big.Int
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if ctrl == nil || number == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.HeaderByNumber(ctx, number)
	})
}

func Fuzz_Nosy_MockClient_HeaderByType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var typ ethclient.HeadType
		fill_err = tp.Fill(&typ)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.HeaderByType(ctx, typ)
	})
}

func Fuzz_Nosy_MockClient_NonceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if ctrl == nil || blockNumber == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.NonceAt(ctx, account, blockNumber)
	})
}

func Fuzz_Nosy_MockClient_PeerCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.PeerCount(ctx)
	})
}

func Fuzz_Nosy_MockClient_PendingBalanceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.PendingBalanceAt(ctx, account)
	})
}

func Fuzz_Nosy_MockClient_PendingCodeAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.PendingCodeAt(ctx, account)
	})
}

func Fuzz_Nosy_MockClient_PendingNonceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.PendingNonceAt(ctx, account)
	})
}

func Fuzz_Nosy_MockClient_PendingStorageAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.PendingStorageAt(ctx, account, key)
	})
}

func Fuzz_Nosy_MockClient_PendingTransactionCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.PendingTransactionCount(ctx)
	})
}

func Fuzz_Nosy_MockClient_SendTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
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
		if ctrl == nil || tx == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.SendTransaction(ctx, tx)
	})
}

func Fuzz_Nosy_MockClient_SetHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
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
		if ctrl == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.SetHead(ctx, height)
	})
}

func Fuzz_Nosy_MockClient_StorageAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if ctrl == nil || blockNumber == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.StorageAt(ctx, account, key, blockNumber)
	})
}

// skipping Fuzz_Nosy_MockClient_SubscribeFilterLogs__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/core/types.Log

// skipping Fuzz_Nosy_MockClient_SubscribeNewHead__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum/go-ethereum/core/types.Header

func Fuzz_Nosy_MockClient_SuggestGasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.SuggestGasPrice(ctx)
	})
}

func Fuzz_Nosy_MockClient_SuggestGasTipCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.SuggestGasTipCap(ctx)
	})
}

func Fuzz_Nosy_MockClient_SyncProgress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.SyncProgress(ctx)
	})
}

func Fuzz_Nosy_MockClient_TransactionByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var txHash common.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.TransactionByHash(ctx, txHash)
	})
}

func Fuzz_Nosy_MockClient_TransactionCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
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
		if ctrl == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.TransactionCount(ctx, blockHash)
	})
}

func Fuzz_Nosy_MockClient_TransactionInBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
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
		var index uint
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.TransactionInBlock(ctx, blockHash, index)
	})
}

func Fuzz_Nosy_MockClient_TransactionReceipt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var txHash common.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClient(ctrl)
		m.TransactionReceipt(ctx, txHash)
	})
}

func Fuzz_Nosy_MockClientMockRecorder_Address__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mr *MockClientMockRecorder
		fill_err = tp.Fill(&mr)
		if fill_err != nil {
			return
		}
		if mr == nil {
			return
		}

		mr.Address()
	})
}

// skipping Fuzz_Nosy_MockClientMockRecorder_BalanceAt__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_BlockByHash__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_BlockByNumber__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_BlockNumber__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_CallContract__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_ChainID__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_MockClientMockRecorder_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mr *MockClientMockRecorder
		fill_err = tp.Fill(&mr)
		if fill_err != nil {
			return
		}
		if mr == nil {
			return
		}

		mr.Close()
	})
}

// skipping Fuzz_Nosy_MockClientMockRecorder_CodeAt__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_EstimateGas__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_EtherBalanceAt__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_FilterLogs__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_HeaderByHash__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_HeaderByNumber__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_HeaderByType__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_NonceAt__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_PeerCount__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_PendingBalanceAt__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_PendingCodeAt__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_PendingNonceAt__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_PendingStorageAt__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_PendingTransactionCount__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_SendTransaction__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_SetHead__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_StorageAt__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_SubscribeFilterLogs__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_SubscribeNewHead__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_SuggestGasPrice__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_SuggestGasTipCap__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_SyncProgress__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_TransactionByHash__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_TransactionCount__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_TransactionInBlock__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockClientMockRecorder_TransactionReceipt__ because parameters include func, chan, or unsupported interface: any
