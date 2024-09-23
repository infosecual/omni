package app

import (
	"context"
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

func Fuzz_Nosy_voterLoader_GetAvailable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *voterLoader
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.GetAvailable()
	})
}

// skipping Fuzz_Nosy_voterLoader_LazyLoad__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/ethclient.Client

func Fuzz_Nosy_voterLoader_LocalAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *voterLoader
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.LocalAddress()
	})
}

func Fuzz_Nosy_voterLoader_SetCommitted__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *voterLoader
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var headers []*types.AttestHeader
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.SetCommitted(headers)
	})
}

func Fuzz_Nosy_voterLoader_SetProposed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *voterLoader
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var headers []*types.AttestHeader
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.SetProposed(headers)
	})
}

func Fuzz_Nosy_voterLoader_TrimBehind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *voterLoader
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var minsByChain map[xchain.ChainVersion]uint64
		fill_err = tp.Fill(&minsByChain)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.TrimBehind(minsByChain)
	})
}

func Fuzz_Nosy_voterLoader_UpdateValidatorSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *voterLoader
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var valset *types.ValidatorSetResponse
		fill_err = tp.Fill(&valset)
		if fill_err != nil {
			return
		}
		if l == nil || valset == nil {
			return
		}

		l.UpdateValidatorSet(valset)
	})
}

func Fuzz_Nosy_voterLoader_WaitDone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *voterLoader
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.WaitDone()
	})
}

func Fuzz_Nosy_voterLoader_getVoter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *voterLoader
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.getVoter()
	})
}

func Fuzz_Nosy_voterLoader_isValidator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *voterLoader
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.isValidator()
	})
}

func Fuzz_Nosy_App_ExportAppStateAndValidators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 App
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 bool
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 []string
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		var _x4 []string
		fill_err = tp.Fill(&_x4)
		if fill_err != nil {
			return
		}

		_x1.ExportAppStateAndValidators(_x2, _x3, _x4)
	})
}

func Fuzz_Nosy_App_LegacyAmino__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 App
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.LegacyAmino()
	})
}

// skipping Fuzz_Nosy_App_SetCometAPI__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/comet.API

func Fuzz_Nosy_App_SimulationManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 App
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.SimulationManager()
	})
}

func Fuzz_Nosy_Config_BackendType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c Config
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.BackendType()
	})
}

func Fuzz_Nosy_abciWrapper_ApplySnapshotChunk__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l abciWrapper
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chunk *types.RequestApplySnapshotChunk
		fill_err = tp.Fill(&chunk)
		if fill_err != nil {
			return
		}
		if chunk == nil {
			return
		}

		l.ApplySnapshotChunk(ctx, chunk)
	})
}

func Fuzz_Nosy_abciWrapper_CheckTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l abciWrapper
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var tx *types.RequestCheckTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		l.CheckTx(ctx, tx)
	})
}

func Fuzz_Nosy_abciWrapper_Commit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l abciWrapper
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var commit *types.RequestCommit
		fill_err = tp.Fill(&commit)
		if fill_err != nil {
			return
		}
		if commit == nil {
			return
		}

		l.Commit(ctx, commit)
	})
}

func Fuzz_Nosy_abciWrapper_ExtendVote__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l abciWrapper
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var vote *types.RequestExtendVote
		fill_err = tp.Fill(&vote)
		if fill_err != nil {
			return
		}
		if vote == nil {
			return
		}

		l.ExtendVote(ctx, vote)
	})
}

func Fuzz_Nosy_abciWrapper_FinalizeBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l abciWrapper
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *types.RequestFinalizeBlock
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		l.FinalizeBlock(ctx, req)
	})
}

func Fuzz_Nosy_abciWrapper_Info__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l abciWrapper
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var info *types.RequestInfo
		fill_err = tp.Fill(&info)
		if fill_err != nil {
			return
		}
		if info == nil {
			return
		}

		l.Info(ctx, info)
	})
}

func Fuzz_Nosy_abciWrapper_InitChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l abciWrapper
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chain *types.RequestInitChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		if chain == nil {
			return
		}

		l.InitChain(ctx, chain)
	})
}

func Fuzz_Nosy_abciWrapper_ListSnapshots__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l abciWrapper
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var listSnapshots *types.RequestListSnapshots
		fill_err = tp.Fill(&listSnapshots)
		if fill_err != nil {
			return
		}
		if listSnapshots == nil {
			return
		}

		l.ListSnapshots(ctx, listSnapshots)
	})
}

func Fuzz_Nosy_abciWrapper_LoadSnapshotChunk__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l abciWrapper
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chunk *types.RequestLoadSnapshotChunk
		fill_err = tp.Fill(&chunk)
		if fill_err != nil {
			return
		}
		if chunk == nil {
			return
		}

		l.LoadSnapshotChunk(ctx, chunk)
	})
}

func Fuzz_Nosy_abciWrapper_OfferSnapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l abciWrapper
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var snapshot *types.RequestOfferSnapshot
		fill_err = tp.Fill(&snapshot)
		if fill_err != nil {
			return
		}
		if snapshot == nil {
			return
		}

		l.OfferSnapshot(ctx, snapshot)
	})
}

func Fuzz_Nosy_abciWrapper_PrepareProposal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l abciWrapper
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var proposal *types.RequestPrepareProposal
		fill_err = tp.Fill(&proposal)
		if fill_err != nil {
			return
		}
		if proposal == nil {
			return
		}

		l.PrepareProposal(ctx, proposal)
	})
}

func Fuzz_Nosy_abciWrapper_ProcessProposal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l abciWrapper
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var proposal *types.RequestProcessProposal
		fill_err = tp.Fill(&proposal)
		if fill_err != nil {
			return
		}
		if proposal == nil {
			return
		}

		l.ProcessProposal(ctx, proposal)
	})
}

func Fuzz_Nosy_abciWrapper_Query__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l abciWrapper
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var query *types.RequestQuery
		fill_err = tp.Fill(&query)
		if fill_err != nil {
			return
		}
		if query == nil {
			return
		}

		l.Query(ctx, query)
	})
}

func Fuzz_Nosy_abciWrapper_VerifyVoteExtension__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l abciWrapper
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var extension *types.RequestVerifyVoteExtension
		fill_err = tp.Fill(&extension)
		if fill_err != nil {
			return
		}
		if extension == nil {
			return
		}

		l.VerifyVoteExtension(ctx, extension)
	})
}

func Fuzz_Nosy_burnEVMFees_LocalFeeRecipient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 burnEVMFees
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.LocalFeeRecipient()
	})
}

func Fuzz_Nosy_burnEVMFees_VerifyFeeRecipient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 burnEVMFees
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}

		_x1.VerifyFeeRecipient(address)
	})
}

// skipping Fuzz_Nosy_cmtLogger_Debug__ because parameters include func, chan, or unsupported interface: []any

// skipping Fuzz_Nosy_cmtLogger_Error__ because parameters include func, chan, or unsupported interface: []any

// skipping Fuzz_Nosy_cmtLogger_Info__ because parameters include func, chan, or unsupported interface: []any

// skipping Fuzz_Nosy_cmtLogger_With__ because parameters include func, chan, or unsupported interface: []any

// skipping Fuzz_Nosy_sdkLogger_Debug__ because parameters include func, chan, or unsupported interface: []any

// skipping Fuzz_Nosy_sdkLogger_Error__ because parameters include func, chan, or unsupported interface: []any

func Fuzz_Nosy_sdkLogger_Impl__(f *testing.F) {
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

		c := newSDKLogger(ctx)
		c.Impl()
	})
}

// skipping Fuzz_Nosy_sdkLogger_Info__ because parameters include func, chan, or unsupported interface: []any

// skipping Fuzz_Nosy_sdkLogger_Warn__ because parameters include func, chan, or unsupported interface: []any

// skipping Fuzz_Nosy_sdkLogger_With__ because parameters include func, chan, or unsupported interface: []any

func Fuzz_Nosy_serverAppOpts_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}

		o := serverAppOptsFromCfg(cfg)
		o.Get(key)
	})
}

func Fuzz_Nosy_Start__(f *testing.F) {
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
		var cfg Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}

		Start(ctx, cfg)
	})
}

func Fuzz_Nosy_chainIDFromGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}

		chainIDFromGenesis(cfg)
	})
}

func Fuzz_Nosy_dirSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		dirSize(path)
	})
}

func Fuzz_Nosy_exists__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string) {
		exists(file)
	})
}

func Fuzz_Nosy_makeBaseAppOpts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}

		makeBaseAppOpts(cfg)
	})
}

// skipping Fuzz_Nosy_monitorCometForever__ because parameters include func, chan, or unsupported interface: github.com/cometbft/cometbft/rpc/client.Client

// skipping Fuzz_Nosy_monitorCometOnce__ because parameters include func, chan, or unsupported interface: github.com/cometbft/cometbft/rpc/client.Client

// skipping Fuzz_Nosy_monitorEVMForever__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/ethclient.Client

// skipping Fuzz_Nosy_setConstantGauge__ because parameters include func, chan, or unsupported interface: github.com/prometheus/client_golang/prometheus.Gauge

// skipping Fuzz_Nosy_splitOutError__ because parameters include func, chan, or unsupported interface: []any
