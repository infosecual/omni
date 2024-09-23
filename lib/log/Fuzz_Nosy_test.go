package log

import (
	"context"
	"log/slog"
	"testing"

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

// skipping Fuzz_Nosy_omniErr_Attrs__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/log.omniErr

// skipping Fuzz_Nosy_stackTracer_StackTrace__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/log.stackTracer

func Fuzz_Nosy_stubHandler_Handle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 stubHandler
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var r slog.Record
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}

		t1.Handle(ctx, r)
	})
}

func Fuzz_Nosy_BindFlags__(f *testing.F) {
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
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if flags == nil || cfg == nil {
			return
		}

		BindFlags(flags, cfg)
	})
}

// skipping Fuzz_Nosy_Debug__ because parameters include func, chan, or unsupported interface: []any

// skipping Fuzz_Nosy_Error__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_Info__ because parameters include func, chan, or unsupported interface: []any

// skipping Fuzz_Nosy_Warn__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_errAttrs__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_log__ because parameters include func, chan, or unsupported interface: []any

// skipping Fuzz_Nosy_mergeAttrs__ because parameters include func, chan, or unsupported interface: []any

func Fuzz_Nosy_slogReplaceAtts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o options
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}

		slogReplaceAtts(o)
	})
}
