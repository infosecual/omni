package ticker

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

// skipping Fuzz_Nosy_MockTicker_Go__ because parameters include func, chan, or unsupported interface: func(context.Context)

// skipping Fuzz_Nosy_Ticker_Go__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/xfeemngr/ticker.Ticker

// skipping Fuzz_Nosy_TimeTicker_Go__ because parameters include func, chan, or unsupported interface: func(context.Context)

func Fuzz_Nosy_WithInterval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var interval time.Duration
		fill_err = tp.Fill(&interval)
		if fill_err != nil {
			return
		}

		WithInterval(interval)
	})
}
