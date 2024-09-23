package expbackoff

import (
	"testing"
	"time"

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

// skipping Fuzz_Nosy_New__ because parameters include func, chan, or unsupported interface: []func(*github.com/omni-network/omni/lib/expbackoff.Config)

// skipping Fuzz_Nosy_NewWithAutoReset__ because parameters include func, chan, or unsupported interface: []func(*github.com/omni-network/omni/lib/expbackoff.Config)

// skipping Fuzz_Nosy_NewWithReset__ because parameters include func, chan, or unsupported interface: []func(*github.com/omni-network/omni/lib/expbackoff.Config)

// skipping Fuzz_Nosy_SetAfterForT__ because parameters include func, chan, or unsupported interface: func(d time.Duration) <-chan time.Time

// skipping Fuzz_Nosy_SetRandFloatForT__ because parameters include func, chan, or unsupported interface: func() float64

func Fuzz_Nosy_With__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c Config
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		With(c)
	})
}

func Fuzz_Nosy_WithPeriodicConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var period time.Duration
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}

		WithPeriodicConfig(period)
	})
}
