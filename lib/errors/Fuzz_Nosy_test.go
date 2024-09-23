package errors

import (
	"testing"

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

// skipping Fuzz_Nosy_stackTracer_StackTrace__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/errors.stackTracer

func Fuzz_Nosy_structured_Attrs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s structured
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.Attrs()
	})
}

func Fuzz_Nosy_structured_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s structured
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.Error()
	})
}

// skipping Fuzz_Nosy_structured_Is__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_structured_StackTrace__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s structured
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.StackTrace()
	})
}

func Fuzz_Nosy_structured_Unwrap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s structured
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.Unwrap()
	})
}

// skipping Fuzz_Nosy_As__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_Is__ because parameters include func, chan, or unsupported interface: error
