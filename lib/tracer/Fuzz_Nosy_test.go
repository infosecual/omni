package tracer

import (
	"context"
	"io"
	"testing"

	"github.com/spf13/pflag"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"go.opentelemetry.io/otel/attribute"
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

func Fuzz_Nosy_AddEvent__(f *testing.F) {
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
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var attrs []attribute.KeyValue
		fill_err = tp.Fill(&attrs)
		if fill_err != nil {
			return
		}

		AddEvent(ctx, name, attrs...)
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

// skipping Fuzz_Nosy_Init__ because parameters include func, chan, or unsupported interface: []func(*github.com/omni-network/omni/lib/tracer.options)

// skipping Fuzz_Nosy_Start__ because parameters include func, chan, or unsupported interface: []go.opentelemetry.io/otel/trace.SpanStartOption

// skipping Fuzz_Nosy_StartChainHeight__ because parameters include func, chan, or unsupported interface: []go.opentelemetry.io/otel/trace.SpanStartOption

func Fuzz_Nosy_WithOTLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var endpoint string
		fill_err = tp.Fill(&endpoint)
		if fill_err != nil {
			return
		}
		var headers map[string]string
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}

		WithOTLP(endpoint, headers)
	})
}

func Fuzz_Nosy_WithStdOut__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}

		WithStdOut(w)
	})
}

func Fuzz_Nosy_stringToHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, value string) {
		stringToHeader(value)
	})
}
