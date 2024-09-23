package cmd

import (
	"context"
	"testing"

	"github.com/cometbft/cometbft/rpc/client/http"
	"github.com/omni-network/omni/halo/app"
	"github.com/omni-network/omni/halo/config"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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

func Fuzz_Nosy_InitConfig_CometCfg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c InitConfig
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c.CometCfg(cfg)
	})
}

func Fuzz_Nosy_InitConfig_HaloCfg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c InitConfig
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c.HaloCfg(cfg)
	})
}

func Fuzz_Nosy_InitConfig_LogCfg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c InitConfig
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.LogCfg()
	})
}

func Fuzz_Nosy_InitConfig_Verify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c InitConfig
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.Verify()
	})
}

func Fuzz_Nosy_bindInitFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var flags *pflag.FlagSet
		fill_err = tp.Fill(&flags)
		if fill_err != nil {
			return
		}
		var cfg *InitConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if flags == nil || cfg == nil {
			return
		}

		bindInitFlags(flags, cfg)
	})
}

func Fuzz_Nosy_bindRollbackFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var flags *pflag.FlagSet
		fill_err = tp.Fill(&flags)
		if fill_err != nil {
			return
		}
		var cfg *app.RollbackConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if flags == nil || cfg == nil {
			return
		}

		bindRollbackFlags(flags, cfg)
	})
}

func Fuzz_Nosy_bindRunFlags__(f *testing.F) {
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
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if c1 == nil || cfg == nil {
			return
		}

		bindRunFlags(c1, cfg)
	})
}

func Fuzz_Nosy_getPeers__(f *testing.F) {
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
		var cl *http.HTTP
		fill_err = tp.Fill(&cl)
		if fill_err != nil {
			return
		}
		var max int
		fill_err = tp.Fill(&max)
		if fill_err != nil {
			return
		}
		if cl == nil {
			return
		}

		getPeers(ctx, cl, max)
	})
}

func Fuzz_Nosy_getTrustHeightAndHash__(f *testing.F) {
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
		var cl *http.HTTP
		fill_err = tp.Fill(&cl)
		if fill_err != nil {
			return
		}
		if cl == nil {
			return
		}

		getTrustHeightAndHash(ctx, cl)
	})
}

func Fuzz_Nosy_setMonikerForT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		setMonikerForT(t1)
	})
}
