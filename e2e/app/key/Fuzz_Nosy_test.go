package key

import (
	"context"
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

func Fuzz_Nosy_Key_Addr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var typ Type
		fill_err = tp.Fill(&typ)
		if fill_err != nil {
			return
		}

		k := Generate(typ)
		k.Addr()
	})
}

func Fuzz_Nosy_Key_ECDSA__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var typ Type
		fill_err = tp.Fill(&typ)
		if fill_err != nil {
			return
		}

		k := Generate(typ)
		k.ECDSA()
	})
}

func Fuzz_Nosy_Type_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 Type
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		t1.String()
	})
}

func Fuzz_Nosy_Type_Verify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 Type
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		t1.Verify()
	})
}

func Fuzz_Nosy_getGCPSecret__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}

		getGCPSecret(ctx, name)
	})
}

func Fuzz_Nosy_secretName__(f *testing.F) {
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
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var typ Type
		fill_err = tp.Fill(&typ)
		if fill_err != nil {
			return
		}
		var addr string
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}

		secretName(network, name, typ, addr)
	})
}
