package main

import (
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

func Fuzz_Nosy_ChainName_RPCURL__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c ChainName
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.RPCURL()
	})
}

func Fuzz_Nosy_ChainName_Validate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c ChainName
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.Validate()
	})
}

func Fuzz_Nosy_ChainName_VerifierURL__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c ChainName
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.VerifierURL()
	})
}

func Fuzz_Nosy_bindFlags__(f *testing.F) {
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
		var apikeys *APIKeys
		fill_err = tp.Fill(&apikeys)
		if fill_err != nil {
			return
		}
		if flags == nil || apikeys == nil {
			return
		}

		bindFlags(flags, apikeys)
	})
}

func Fuzz_Nosy_mustDecodeHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hex string) {
		mustDecodeHex(hex)
	})
}
