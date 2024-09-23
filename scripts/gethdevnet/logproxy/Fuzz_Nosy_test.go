package main

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

func Fuzz_Nosy_responseWriterWrapper_Header__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *responseWriterWrapper
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.Header()
	})
}

func Fuzz_Nosy_responseWriterWrapper_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *responseWriterWrapper
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.Write(b)
	})
}

func Fuzz_Nosy_responseWriterWrapper_WriteHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *responseWriterWrapper
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var status int
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.WriteHeader(status)
	})
}

func Fuzz_Nosy_gunzip__(f *testing.F) {
	f.Fuzz(func(t *testing.T, body []byte) {
		gunzip(body)
	})
}
