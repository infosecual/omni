package tutil

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

func Fuzz_Nosy_RandomAvailablePort__(f *testing.F) {
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

		RandomAvailablePort(t1)
	})
}

func Fuzz_Nosy_RandomBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, l int) {
		RandomBytes(l)
	})
}

func Fuzz_Nosy_RandomListenAddress__(f *testing.F) {
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

		RandomListenAddress(t1)
	})
}

// skipping Fuzz_Nosy_RequireGoldenBytes__ because parameters include func, chan, or unsupported interface: []func(*string)

// skipping Fuzz_Nosy_RequireGoldenJSON__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_RequireNoError__ because parameters include func, chan, or unsupported interface: testing.TB

func Fuzz_Nosy_TempDir__(f *testing.F) {
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

		TempDir(t1)
	})
}

func Fuzz_Nosy_WithFilename__(f *testing.F) {
	f.Fuzz(func(t *testing.T, name string) {
		WithFilename(name)
	})
}

func Fuzz_Nosy_randStr__(f *testing.F) {
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

		randStr(t1)
	})
}
