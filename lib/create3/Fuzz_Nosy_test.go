package create3

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
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

func Fuzz_Nosy_HashSalt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		HashSalt(s)
	})
}

func Fuzz_Nosy_namespacedSalt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var deployer common.Address
		fill_err = tp.Fill(&deployer)
		if fill_err != nil {
			return
		}
		var salt string
		fill_err = tp.Fill(&salt)
		if fill_err != nil {
			return
		}

		namespacedSalt(deployer, salt)
	})
}
