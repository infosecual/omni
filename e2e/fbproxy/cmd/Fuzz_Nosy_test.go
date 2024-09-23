package cmd

import (
	"testing"

	"github.com/omni-network/omni/e2e/fbproxy/app"
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

func Fuzz_Nosy_bindFlags__(f *testing.F) {
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
		var cfg *app.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if flags == nil || cfg == nil {
			return
		}

		bindFlags(flags, cfg)
	})
}
