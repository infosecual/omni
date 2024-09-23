package stream

import (
	"testing"
	"github.com/omni-network/omni/lib"
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
	
	// skipping Fuzz_Nosy_sortingBuffer[E any]_Add__ because parameters include func, chan, or unsupported interface: []E

func Fuzz_Nosy_sortingBuffer[E any]_Process__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *sortingBuffer[E]
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.Process(ctx)
	})
}

// skipping Fuzz_Nosy_sortingBuffer[E any]_retryLock__ because parameters include func, chan, or unsupported interface: func(ctx context.Context) (bool, error)

func Fuzz_Nosy_sortingBuffer[E any]_signal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *sortingBuffer[E]
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var signalID int
		fill_err = tp.Fill(&signalID)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.signal(signalID)
	})
}

// skipping Fuzz_Nosy_startFetchWorkers__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, height uint64) []E

