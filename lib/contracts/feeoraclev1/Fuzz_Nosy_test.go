package feeoraclev1

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"
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

func Fuzz_Nosy_deploymentConfig_Validate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg deploymentConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}

		cfg.Validate()
	})
}

func Fuzz_Nosy_Deploy__(f *testing.F) {
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
		var destChainIDs []uint64
		fill_err = tp.Fill(&destChainIDs)
		if fill_err != nil {
			return
		}
		var backends ethbackend.Backends
		fill_err = tp.Fill(&backends)
		if fill_err != nil {
			return
		}

		Deploy(ctx, network, chainID, destChainIDs, backends)
	})
}

// skipping Fuzz_Nosy_conversionRate__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/tokens.Pricer

// skipping Fuzz_Nosy_feeParams__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/tokens.Pricer

func Fuzz_Nosy_isDeadOrEmpty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}

		isDeadOrEmpty(addr)
	})
}
