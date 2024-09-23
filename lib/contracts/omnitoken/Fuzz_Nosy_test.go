package omnitoken

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

func Fuzz_Nosy_deploymentConfig_validate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var network netconf.ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}

		cfg, err := getConfig(network)
		if err != nil {
			return
		}
		cfg.validate()
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
		var backend *ethbackend.Backend
		fill_err = tp.Fill(&backend)
		if fill_err != nil {
			return
		}
		if backend == nil {
			return
		}

		Deploy(ctx, network, backend)
	})
}

func Fuzz_Nosy_DeployIfNeeded__(f *testing.F) {
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
		var backend *ethbackend.Backend
		fill_err = tp.Fill(&backend)
		if fill_err != nil {
			return
		}
		if backend == nil {
			return
		}

		DeployIfNeeded(ctx, network, backend)
	})
}

func Fuzz_Nosy_deploy__(f *testing.F) {
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
		var cfg deploymentConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var backend *ethbackend.Backend
		fill_err = tp.Fill(&backend)
		if fill_err != nil {
			return
		}
		if backend == nil {
			return
		}

		deploy(ctx, cfg, backend)
	})
}

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

func Fuzz_Nosy_isDeployed__(f *testing.F) {
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
		var backend *ethbackend.Backend
		fill_err = tp.Fill(&backend)
		if fill_err != nil {
			return
		}
		if backend == nil {
			return
		}

		isDeployed(ctx, network, backend)
	})
}

func Fuzz_Nosy_packInitCode__(f *testing.F) {
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

		packInitCode(cfg)
	})
}
