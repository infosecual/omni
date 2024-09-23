package txmgr

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

func Fuzz_Nosy_SendState_IsWaitingForConfirmation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonceTooLowCount uint64
		fill_err = tp.Fill(&nonceTooLowCount)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}

		s := NewSendState(nonceTooLowCount, timeout)
		s.IsWaitingForConfirmation()
	})
}

// skipping Fuzz_Nosy_SendState_ProcessSendError__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_SendState_ShouldAbortImmediately__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonceTooLowCount uint64
		fill_err = tp.Fill(&nonceTooLowCount)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}

		s := NewSendState(nonceTooLowCount, timeout)
		s.ShouldAbortImmediately()
	})
}

func Fuzz_Nosy_SendState_TxMined__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonceTooLowCount uint64
		fill_err = tp.Fill(&nonceTooLowCount)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}
		var txHash common.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}

		s := NewSendState(nonceTooLowCount, timeout)
		s.TxMined(txHash)
	})
}

func Fuzz_Nosy_SendState_TxNotMined__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonceTooLowCount uint64
		fill_err = tp.Fill(&nonceTooLowCount)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}
		var txHash common.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}

		s := NewSendState(nonceTooLowCount, timeout)
		s.TxNotMined(txHash)
	})
}

func Fuzz_Nosy_simple_From__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *simple
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.From()
	})
}

func Fuzz_Nosy_simple_ReserveNextNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *simple
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.ReserveNextNonce(ctx)
	})
}

func Fuzz_Nosy_simple_Send__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *simple
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var candidate TxCandidate
		fill_err = tp.Fill(&candidate)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Send(ctx, candidate)
	})
}

func Fuzz_Nosy_simple_checkLimits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *simple
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var tip *big.Int
		fill_err = tp.Fill(&tip)
		if fill_err != nil {
			return
		}
		var baseFee *big.Int
		fill_err = tp.Fill(&baseFee)
		if fill_err != nil {
			return
		}
		var bumpedTip *big.Int
		fill_err = tp.Fill(&bumpedTip)
		if fill_err != nil {
			return
		}
		var bumpedFee *big.Int
		fill_err = tp.Fill(&bumpedFee)
		if fill_err != nil {
			return
		}
		if m == nil || tip == nil || baseFee == nil || bumpedTip == nil || bumpedFee == nil {
			return
		}

		m.checkLimits(tip, baseFee, bumpedTip, bumpedFee)
	})
}

func Fuzz_Nosy_simple_craftTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *simple
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var candidate TxCandidate
		fill_err = tp.Fill(&candidate)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.craftTx(ctx, candidate)
	})
}

func Fuzz_Nosy_simple_doSend__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *simple
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var candidate TxCandidate
		fill_err = tp.Fill(&candidate)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.doSend(ctx, candidate)
	})
}

func Fuzz_Nosy_simple_increaseGasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *simple
		fill_err = tp.Fill(&m)
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
		if m == nil || tx == nil {
			return
		}

		m.increaseGasPrice(ctx, tx)
	})
}

func Fuzz_Nosy_simple_publishTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *simple
		fill_err = tp.Fill(&m)
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
		var sendState *SendState
		fill_err = tp.Fill(&sendState)
		if fill_err != nil {
			return
		}
		var bumpFeesImmediately bool
		fill_err = tp.Fill(&bumpFeesImmediately)
		if fill_err != nil {
			return
		}
		if m == nil || tx == nil || sendState == nil {
			return
		}

		m.publishTx(ctx, tx, sendState, bumpFeesImmediately)
	})
}

func Fuzz_Nosy_simple_queryReceipt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *simple
		fill_err = tp.Fill(&m)
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
		var sendState *SendState
		fill_err = tp.Fill(&sendState)
		if fill_err != nil {
			return
		}
		if m == nil || sendState == nil {
			return
		}

		m.queryReceipt(ctx, txHash, sendState)
	})
}

func Fuzz_Nosy_simple_resetNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *simple
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.resetNonce()
	})
}

func Fuzz_Nosy_simple_sendTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *simple
		fill_err = tp.Fill(&m)
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
		if m == nil || tx == nil {
			return
		}

		m.sendTx(ctx, tx)
	})
}

func Fuzz_Nosy_simple_suggestGasPriceCaps__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *simple
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.suggestGasPriceCaps(ctx)
	})
}

// skipping Fuzz_Nosy_simple_waitForTx__ because parameters include func, chan, or unsupported interface: chan github.com/omni-network/omni/lib/txmgr.minedTuple

func Fuzz_Nosy_simple_waitMined__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *simple
		fill_err = tp.Fill(&m)
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
		var sendState *SendState
		fill_err = tp.Fill(&sendState)
		if fill_err != nil {
			return
		}
		if m == nil || tx == nil || sendState == nil {
			return
		}

		m.waitMined(ctx, tx, sendState)
	})
}

func Fuzz_Nosy_CLIConfig_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainID uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var interval time.Duration
		fill_err = tp.Fill(&interval)
		if fill_err != nil {
			return
		}
		var defaults DefaultFlagValues
		fill_err = tp.Fill(&defaults)
		if fill_err != nil {
			return
		}

		m := NewCLIConfig(chainID, interval, defaults)
		m.Check()
	})
}

func Fuzz_Nosy_Config_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m Config
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}

		m.Check()
	})
}

func Fuzz_Nosy_TxManager_From__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var conf Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}

		_x1, err := NewSimple(chainName, conf)
		if err != nil {
			return
		}
		_x1.From()
	})
}

func Fuzz_Nosy_TxManager_ReserveNextNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var conf Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		_x1, err := NewSimple(chainName, conf)
		if err != nil {
			return
		}
		_x1.ReserveNextNonce(ctx)
	})
}

func Fuzz_Nosy_TxManager_Send__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var conf Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var candidate TxCandidate
		fill_err = tp.Fill(&candidate)
		if fill_err != nil {
			return
		}

		_x1, err := NewSimple(chainName, conf)
		if err != nil {
			return
		}
		_x1.Send(ctx, candidate)
	})
}

// skipping Fuzz_Nosy_errStringMatch__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_isNetworkError__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_maybeSetTimeout__(f *testing.F) {
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
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}

		maybeSetTimeout(ctx, timeout)
	})
}

func Fuzz_Nosy_txFields__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var logGas bool
		fill_err = tp.Fill(&logGas)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		txFields(tx, logGas)
	})
}

func Fuzz_Nosy_updateFees__(f *testing.F) {
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
		var oldTip *big.Int
		fill_err = tp.Fill(&oldTip)
		if fill_err != nil {
			return
		}
		var oldFeeCap *big.Int
		fill_err = tp.Fill(&oldFeeCap)
		if fill_err != nil {
			return
		}
		var newTip *big.Int
		fill_err = tp.Fill(&newTip)
		if fill_err != nil {
			return
		}
		var newBaseFee *big.Int
		fill_err = tp.Fill(&newBaseFee)
		if fill_err != nil {
			return
		}
		if oldTip == nil || oldFeeCap == nil || newTip == nil || newBaseFee == nil {
			return
		}

		updateFees(ctx, oldTip, oldFeeCap, newTip, newBaseFee)
	})
}
