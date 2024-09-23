package evmchain

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

func Fuzz_Nosy_IsOmniEVM__(f *testing.F) {
	f.Fuzz(func(t *testing.T, name string) {
		IsOmniEVM(name)
	})
}

func Fuzz_Nosy_MetadataByID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64) {
		MetadataByID(chainID)
	})
}

func Fuzz_Nosy_MetadataByName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, name string) {
		MetadataByName(name)
	})
}
