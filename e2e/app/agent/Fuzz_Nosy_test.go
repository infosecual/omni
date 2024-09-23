package agent

import (
	"context"
	"testing"

	"github.com/omni-network/omni/e2e/types"
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

func Fuzz_Nosy_promScrapConfig_Targets__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c promScrapConfig
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.Targets()
	})
}

func Fuzz_Nosy_ConfigForHost__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bz []byte
		fill_err = tp.Fill(&bz)
		if fill_err != nil {
			return
		}
		var newHost string
		fill_err = tp.Fill(&newHost)
		if fill_err != nil {
			return
		}
		var halos []string
		fill_err = tp.Fill(&halos)
		if fill_err != nil {
			return
		}
		var geths []string
		fill_err = tp.Fill(&geths)
		if fill_err != nil {
			return
		}
		var services map[string]bool
		fill_err = tp.Fill(&services)
		if fill_err != nil {
			return
		}

		ConfigForHost(bz, newHost, halos, geths, services)
	})
}

func Fuzz_Nosy_genPromConfig__(f *testing.F) {
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
		var secrets Secrets
		fill_err = tp.Fill(&secrets)
		if fill_err != nil {
			return
		}
		var hostname string
		fill_err = tp.Fill(&hostname)
		if fill_err != nil {
			return
		}

		genPromConfig(ctx, testnet, secrets, hostname)
	})
}
