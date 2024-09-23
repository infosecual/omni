package k1util

import (
	"crypto/ecdsa"
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

func Fuzz_Nosy_PubKeyToBytes64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey *ecdsa.PublicKey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		if pubkey == nil {
			return
		}

		PubKeyToBytes64(pubkey)
	})
}

// skipping Fuzz_Nosy_Sign__ because parameters include func, chan, or unsupported interface: github.com/cometbft/cometbft/crypto.PrivKey

func Fuzz_Nosy_Verify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var hash [32]byte
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var sig [65]byte
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}

		Verify(address, hash, sig)
	})
}
