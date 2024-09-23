package cmd

import (
	"context"
	"math/big"
	"testing"

	"github.com/omni-network/omni/lib/netconf"
	"github.com/spf13/cobra"
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

func Fuzz_Nosy_CliError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *CliError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Error()
	})
}

// skipping Fuzz_Nosy_CliError_Wrap__ because parameters include func, chan, or unsupported interface: []any

func Fuzz_Nosy_initConfig_Verify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c initConfig
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.Verify()
	})
}

func Fuzz_Nosy_avsFromChainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if chainID == nil {
			return
		}

		avsFromChainID(chainID)
	})
}

func Fuzz_Nosy_bindAVSAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 *cobra.Command
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var addr *string
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if c1 == nil || addr == nil {
			return
		}

		bindAVSAddress(c1, addr)
	})
}

func Fuzz_Nosy_bindDeveloperForgeProjectConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 *cobra.Command
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg *developerForgeProjectConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if c1 == nil || cfg == nil {
			return
		}

		bindDeveloperForgeProjectConfig(c1, cfg)
	})
}

func Fuzz_Nosy_bindDevnetAVSAllowConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 *cobra.Command
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg *devnetAllowConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if c1 == nil || cfg == nil {
			return
		}

		bindDevnetAVSAllowConfig(c1, cfg)
	})
}

func Fuzz_Nosy_bindDevnetFundConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 *cobra.Command
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var d *devnetFundConfig
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if c1 == nil || d == nil {
			return
		}

		bindDevnetFundConfig(c1, d)
	})
}

func Fuzz_Nosy_bindInitConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 *cobra.Command
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg *initConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if c1 == nil || cfg == nil {
			return
		}

		bindInitConfig(c1, cfg)
	})
}

func Fuzz_Nosy_bindRPCURL__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 *cobra.Command
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rpcURL *string
		fill_err = tp.Fill(&rpcURL)
		if fill_err != nil {
			return
		}
		if c1 == nil || rpcURL == nil {
			return
		}

		bindRPCURL(c1, rpcURL)
	})
}

func Fuzz_Nosy_bindRegConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 *cobra.Command
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg *RegConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if c1 == nil || cfg == nil {
			return
		}

		bindRegConfig(c1, cfg)
	})
}

func Fuzz_Nosy_chainNameFromID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id big.Int
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		chainNameFromID(id)
	})
}

func Fuzz_Nosy_devnetBackend__(f *testing.F) {
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
		var rpcURL string
		fill_err = tp.Fill(&rpcURL)
		if fill_err != nil {
			return
		}

		devnetBackend(ctx, rpcURL)
	})
}

func Fuzz_Nosy_homeDir__(f *testing.F) {
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

		homeDir(network)
	})
}

func Fuzz_Nosy_loadDevnetNetwork__(f *testing.F) {
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

		loadDevnetNetwork(ctx)
	})
}

func Fuzz_Nosy_writeTempFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, content []byte) {
		writeTempFile(content)
	})
}
