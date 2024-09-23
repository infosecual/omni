package cmd

import (
	"testing"

	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/e2e/app/admin"
	"github.com/omni-network/omni/e2e/app/agent"
	"github.com/omni-network/omni/e2e/app/key"
	"github.com/omni-network/omni/e2e/types"
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

func Fuzz_Nosy_bindCreate3DeployFlags__(f *testing.F) {
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
		var cfg *app.Create3DeployConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if flags == nil || cfg == nil {
			return
		}

		bindCreate3DeployFlags(flags, cfg)
	})
}

func Fuzz_Nosy_bindDefFlags__(f *testing.F) {
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
		var cfg *app.DefinitionConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if flags == nil || cfg == nil {
			return
		}

		bindDefFlags(flags, cfg)
	})
}

func Fuzz_Nosy_bindDeployFlags__(f *testing.F) {
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
		var cfg *app.DeployConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if flags == nil || cfg == nil {
			return
		}

		bindDeployFlags(flags, cfg)
	})
}

func Fuzz_Nosy_bindE2EFlags__(f *testing.F) {
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
		var cfg *app.E2ETestConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if flags == nil || cfg == nil {
			return
		}

		bindE2EFlags(flags, cfg)
	})
}

func Fuzz_Nosy_bindERC20FaucetFlags__(f *testing.F) {
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
		var cfg *app.RunERC20FaucetConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if flags == nil || cfg == nil {
			return
		}

		bindERC20FaucetFlags(flags, cfg)
	})
}

func Fuzz_Nosy_bindKeyCreateFlags__(f *testing.F) {
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
		var cfg *key.UploadConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if c1 == nil || cfg == nil {
			return
		}

		bindKeyCreateFlags(c1, cfg)
	})
}

func Fuzz_Nosy_bindPortalAdminFlags__(f *testing.F) {
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
		var cfg *admin.PortalAdminConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if flags == nil || cfg == nil {
			return
		}

		bindPortalAdminFlags(flags, cfg)
	})
}

func Fuzz_Nosy_bindPromFlags__(f *testing.F) {
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
		var cfg *agent.Secrets
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if flags == nil || cfg == nil {
			return
		}

		bindPromFlags(flags, cfg)
	})
}

func Fuzz_Nosy_bindServiceFlags__(f *testing.F) {
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
		var cfg *types.ServiceConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if flags == nil || cfg == nil {
			return
		}

		bindServiceFlags(flags, cfg)
	})
}

func Fuzz_Nosy_matchAny__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var str string
		fill_err = tp.Fill(&str)
		if fill_err != nil {
			return
		}
		var patterns []string
		fill_err = tp.Fill(&patterns)
		if fill_err != nil {
			return
		}

		matchAny(str, patterns...)
	})
}
