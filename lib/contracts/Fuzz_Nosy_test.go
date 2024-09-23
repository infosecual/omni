package contracts

import (
	"testing"

	"github.com/omni-network/omni/lib/netconf"
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

func Fuzz_Nosy_AVSSalt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var network netconf.ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}

		AVSSalt(network)
	})
}

func Fuzz_Nosy_L1BridgeSalt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var network netconf.ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}

		L1BridgeSalt(network)
	})
}

// skipping Fuzz_Nosy_PackInitCode__ because parameters include func, chan, or unsupported interface: []any

func Fuzz_Nosy_PortalSalt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var network netconf.ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}

		PortalSalt(network)
	})
}

func Fuzz_Nosy_TokenSalt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var network netconf.ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}

		TokenSalt(network)
	})
}

func Fuzz_Nosy_salt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var network netconf.ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var contract string
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}

		salt(network, contract)
	})
}
