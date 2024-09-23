package cmd

import (
	"testing"

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

func Fuzz_Nosy_BindHomeFlag__(f *testing.F) {
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
		var homeDir *string
		fill_err = tp.Fill(&homeDir)
		if fill_err != nil {
			return
		}
		if flags == nil || homeDir == nil {
			return
		}

		BindHomeFlag(flags, homeDir)
	})
}

func Fuzz_Nosy_Main__(f *testing.F) {
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
		if c1 == nil {
			return
		}

		Main(c1)
	})
}

func Fuzz_Nosy_SilenceErrUsage__(f *testing.F) {
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
		if c1 == nil {
			return
		}

		SilenceErrUsage(c1)
	})
}

func Fuzz_Nosy_redact__(f *testing.F) {
	f.Fuzz(func(t *testing.T, flag string, val string) {
		redact(flag, val)
	})
}
