package provider

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/cchain"
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

func Fuzz_Nosy_Provider_AttestationsFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p Provider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainVer xchain.ChainVersion
		fill_err = tp.Fill(&chainVer)
		if fill_err != nil {
			return
		}
		var attestOffset uint64
		fill_err = tp.Fill(&attestOffset)
		if fill_err != nil {
			return
		}

		p.AttestationsFrom(ctx, chainVer, attestOffset)
	})
}

func Fuzz_Nosy_Provider_CometClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p Provider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		p.CometClient()
	})
}

func Fuzz_Nosy_Provider_CurrentUpgradePlan__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p Provider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		p.CurrentUpgradePlan(ctx)
	})
}

func Fuzz_Nosy_Provider_GenesisFiles__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p Provider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		p.GenesisFiles(ctx)
	})
}

func Fuzz_Nosy_Provider_LatestAttestation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p Provider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainVer xchain.ChainVersion
		fill_err = tp.Fill(&chainVer)
		if fill_err != nil {
			return
		}

		p.LatestAttestation(ctx, chainVer)
	})
}

func Fuzz_Nosy_Provider_Portals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p Provider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		p.Portals(ctx)
	})
}

// skipping Fuzz_Nosy_Provider_StreamAsync__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/cchain.ProviderCallback

// skipping Fuzz_Nosy_Provider_StreamAttestations__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/cchain.ProviderCallback

func Fuzz_Nosy_Provider_ValidatorSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p Provider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var valSetID uint64
		fill_err = tp.Fill(&valSetID)
		if fill_err != nil {
			return
		}

		p.ValidatorSet(ctx, valSetID)
	})
}

func Fuzz_Nosy_Provider_WindowCompare__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p Provider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainVer xchain.ChainVersion
		fill_err = tp.Fill(&chainVer)
		if fill_err != nil {
			return
		}
		var attestOffset uint64
		fill_err = tp.Fill(&attestOffset)
		if fill_err != nil {
			return
		}

		p.WindowCompare(ctx, chainVer, attestOffset)
	})
}

func Fuzz_Nosy_Provider_XBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p Provider
		fill_err = tp.Fill(&p)
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
		var latest bool
		fill_err = tp.Fill(&latest)
		if fill_err != nil {
			return
		}

		p.XBlock(ctx, height, latest)
	})
}

func Fuzz_Nosy_Provider_msgNetworkData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p Provider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg *types.Msg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		p.msgNetworkData(ctx, msg)
	})
}

func Fuzz_Nosy_Provider_msgValSetData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p Provider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg *types.Msg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		p.msgValSetData(ctx, msg)
	})
}

// skipping Fuzz_Nosy_Provider_stream__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/cchain.ProviderCallback

// skipping Fuzz_Nosy_rpcAdaptor_Invoke__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_IsErrHistoryPruned__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_attsFromAtHeight__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/types.QueryClient

func Fuzz_Nosy_fetchStepsMetrics__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainName string, lookbackSteps uint64, binarySearchSteps uint64) {
		fetchStepsMetrics(chainName, lookbackSteps, binarySearchSteps)
	})
}

// skipping Fuzz_Nosy_getEarliestStoreHeight__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/types.QueryClient

func Fuzz_Nosy_heightFromCtx__(f *testing.F) {
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

		heightFromCtx(ctx)
	})
}

func Fuzz_Nosy_incQueryErr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, endpoint string) {
		incQueryErr(endpoint)
	})
}

func Fuzz_Nosy_latency__(f *testing.F) {
	f.Fuzz(func(t *testing.T, endpoint string) {
		latency(endpoint)
	})
}

// skipping Fuzz_Nosy_queryEarliestAttestation__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/types.QueryClient

// skipping Fuzz_Nosy_queryLatestAttestation__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/halo/attest/types.QueryClient

// skipping Fuzz_Nosy_searchOffsetInHistory__ because parameters include func, chan, or unsupported interface: github.com/cometbft/cometbft/rpc/client.Client

func Fuzz_Nosy_spanName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, endpoint string) {
		spanName(endpoint)
	})
}

func Fuzz_Nosy_toPortalChains__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var portals []*types.Portal
		fill_err = tp.Fill(&portals)
		if fill_err != nil {
			return
		}

		toPortalChains(portals)
	})
}

func Fuzz_Nosy_toPortalVals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var vals []cchain.Validator
		fill_err = tp.Fill(&vals)
		if fill_err != nil {
			return
		}

		toPortalVals(vals)
	})
}
