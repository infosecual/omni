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

func Fuzz_Nosy_commentTypes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var content string
		fill_err = tp.Fill(&content)
		if fill_err != nil {
			return
		}
		var types []string
		fill_err = tp.Fill(&types)
		if fill_err != nil {
			return
		}

		commentTypes(content, types)
	})
}

func Fuzz_Nosy_isTypeDefEnd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, line string) {
		isTypeDefEnd(line)
	})
}

func Fuzz_Nosy_isTypeDefStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var line string
		fill_err = tp.Fill(&line)
		if fill_err != nil {
			return
		}
		var types []string
		fill_err = tp.Fill(&types)
		if fill_err != nil {
			return
		}

		isTypeDefStart(line, types)
	})
}
