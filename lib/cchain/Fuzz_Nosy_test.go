package cchain

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

// skipping Fuzz_Nosy_Provider_AttestationsFrom__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/cchain.Provider

// skipping Fuzz_Nosy_Provider_CometClient__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/cchain.Provider

// skipping Fuzz_Nosy_Provider_CurrentUpgradePlan__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/cchain.Provider

// skipping Fuzz_Nosy_Provider_GenesisFiles__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/cchain.Provider

// skipping Fuzz_Nosy_Provider_LatestAttestation__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/cchain.Provider

// skipping Fuzz_Nosy_Provider_Portals__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/cchain.Provider

// skipping Fuzz_Nosy_Provider_StreamAsync__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/cchain.Provider

// skipping Fuzz_Nosy_Provider_StreamAttestations__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/cchain.Provider

// skipping Fuzz_Nosy_Provider_ValidatorSet__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/cchain.Provider

// skipping Fuzz_Nosy_Provider_WindowCompare__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/cchain.Provider

// skipping Fuzz_Nosy_Provider_XBlock__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/cchain.Provider

func Fuzz_Nosy_Validator_Verify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v Validator
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		v.Verify()
	})
}
