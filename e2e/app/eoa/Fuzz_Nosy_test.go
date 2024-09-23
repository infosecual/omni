package eoa

import (
	"crypto/ecdsa"
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

func Fuzz_Nosy_Account_privKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a Account
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}

		a.privKey()
	})
}

func Fuzz_Nosy_FundThresholds_MinBalance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 FundThresholds
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		t1.MinBalance()
	})
}

func Fuzz_Nosy_FundThresholds_TargetBalance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 FundThresholds
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		t1.TargetBalance()
	})
}

func Fuzz_Nosy_Role_Verify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r Role
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}

		r.Verify()
	})
}

func Fuzz_Nosy_AccountForRole__(f *testing.F) {
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
		var role Role
		fill_err = tp.Fill(&role)
		if fill_err != nil {
			return
		}

		AccountForRole(network, role)
	})
}

func Fuzz_Nosy_Address__(f *testing.F) {
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
		var role Role
		fill_err = tp.Fill(&role)
		if fill_err != nil {
			return
		}

		Address(network, role)
	})
}

func Fuzz_Nosy_AllAccounts__(f *testing.F) {
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

		AllAccounts(network)
	})
}

func Fuzz_Nosy_GetFundThresholds__(f *testing.F) {
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
		var role Role
		fill_err = tp.Fill(&role)
		if fill_err != nil {
			return
		}

		GetFundThresholds(network, role)
	})
}

func Fuzz_Nosy_MustAddresses__(f *testing.F) {
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
		var roles []Role
		fill_err = tp.Fill(&roles)
		if fill_err != nil {
			return
		}

		MustAddresses(network, roles...)
	})
}

func Fuzz_Nosy_dummy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var roles []Role
		fill_err = tp.Fill(&roles)
		if fill_err != nil {
			return
		}

		dummy(roles...)
	})
}

func Fuzz_Nosy_flatten__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slices [][]T
		fill_err = tp.Fill(&slices)
		if fill_err != nil {
			return
		}

		flatten(slices...)
	})
}

func Fuzz_Nosy_remote__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hex string
		fill_err = tp.Fill(&hex)
		if fill_err != nil {
			return
		}
		var roles []Role
		fill_err = tp.Fill(&roles)
		if fill_err != nil {
			return
		}

		remote(hex, roles...)
	})
}

func Fuzz_Nosy_secret__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hex string
		fill_err = tp.Fill(&hex)
		if fill_err != nil {
			return
		}
		var roles []Role
		fill_err = tp.Fill(&roles)
		if fill_err != nil {
			return
		}

		secret(hex, roles...)
	})
}

func Fuzz_Nosy_wellKnown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pk *ecdsa.PrivateKey
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}
		var roles []Role
		fill_err = tp.Fill(&roles)
		if fill_err != nil {
			return
		}
		if pk == nil {
			return
		}

		wellKnown(pk, roles...)
	})
}
