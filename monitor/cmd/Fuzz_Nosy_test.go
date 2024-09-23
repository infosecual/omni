package cmd

import (
	"testing"

	"github.com/omni-network/omni/monitor"
	"github.com/omni-network/omni/monitor/loadgen"
	"github.com/omni-network/omni/monitor/xfeemngr"
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

func Fuzz_Nosy_bindLoadGenFlags__(f *testing.F) {
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
		var cfg *loadgen.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if flags == nil || cfg == nil {
			return
		}

		bindLoadGenFlags(flags, cfg)
	})
}

func Fuzz_Nosy_bindRunFlags__(f *testing.F) {
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
		var cfg *monitor.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if flags == nil || cfg == nil {
			return
		}

		bindRunFlags(flags, cfg)
	})
}

func Fuzz_Nosy_bindXFeeMngrFlags__(f *testing.F) {
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
		var cfg *xfeemngr.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if flags == nil || cfg == nil {
			return
		}

		bindXFeeMngrFlags(flags, cfg)
	})
}
