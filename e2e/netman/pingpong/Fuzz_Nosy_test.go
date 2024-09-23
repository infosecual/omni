package pingpong

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/ethclient/ethbackend"
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

func Fuzz_Nosy_XDapp_LogBalances__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var network netconf.Network
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var backends ethbackend.Backends
		fill_err = tp.Fill(&backends)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}

		d, err := Deploy(c1, network, backends)
		if err != nil {
			return
		}
		d.LogBalances(c4)
	})
}

func Fuzz_Nosy_XDapp_StartAllEdges__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var network netconf.Network
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var backends ethbackend.Backends
		fill_err = tp.Fill(&backends)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var latest uint64
		fill_err = tp.Fill(&latest)
		if fill_err != nil {
			return
		}
		var parallel uint64
		fill_err = tp.Fill(&parallel)
		if fill_err != nil {
			return
		}
		var count uint64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}

		d, err := Deploy(c1, network, backends)
		if err != nil {
			return
		}
		d.StartAllEdges(c4, latest, parallel, count)
	})
}

func Fuzz_Nosy_XDapp_WaitDone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var network netconf.Network
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var backends ethbackend.Backends
		fill_err = tp.Fill(&backends)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}

		d, err := Deploy(c1, network, backends)
		if err != nil {
			return
		}
		d.WaitDone(c4)
	})
}

func Fuzz_Nosy_XDapp_Watch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var network netconf.Network
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var backends ethbackend.Backends
		fill_err = tp.Fill(&backends)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}

		d, err := Deploy(c1, network, backends)
		if err != nil {
			return
		}
		d.Watch(c4)
	})
}

func Fuzz_Nosy_XDapp_fund__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var network netconf.Network
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var backends ethbackend.Backends
		fill_err = tp.Fill(&backends)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}

		d, err := Deploy(c1, network, backends)
		if err != nil {
			return
		}
		d.fund(c4)
	})
}

func Fuzz_Nosy_edges__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var contracts map[uint64]contract
		fill_err = tp.Fill(&contracts)
		if fill_err != nil {
			return
		}

		edges(contracts)
	})
}
