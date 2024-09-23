package p2putil

import (
	"context"
	"testing"

	"github.com/cometbft/cometbft/p2p"
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

// skipping Fuzz_Nosy_pexReactor_AddPeer__ because parameters include func, chan, or unsupported interface: github.com/cometbft/cometbft/p2p.Peer

func Fuzz_Nosy_pexReactor_GetChannels__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 pexReactor
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.GetChannels()
	})
}

// skipping Fuzz_Nosy_pexReactor_InitPeer__ because parameters include func, chan, or unsupported interface: github.com/cometbft/cometbft/p2p.Peer

func Fuzz_Nosy_pexReactor_OnStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 pexReactor
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.OnStart()
	})
}

func Fuzz_Nosy_pexReactor_OnStop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 pexReactor
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.OnStop()
	})
}

func Fuzz_Nosy_pexReactor_Receive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p pexReactor
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var e p2p.Envelope
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}

		p.Receive(e)
	})
}

// skipping Fuzz_Nosy_pexReactor_RemovePeer__ because parameters include func, chan, or unsupported interface: github.com/cometbft/cometbft/p2p.Peer

func Fuzz_Nosy_pexReactor_SetSwitch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 pexReactor
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 *p2p.Switch
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x2 == nil {
			return
		}

		_x1.SetSwitch(_x2)
	})
}

func Fuzz_Nosy_pexReactor_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 pexReactor
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Start()
	})
}

func Fuzz_Nosy_pexReactor_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 pexReactor
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Stop()
	})
}

func Fuzz_Nosy_FetchPexAddrs__(f *testing.F) {
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
		var network netconf.ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var peer *p2p.NetAddress
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		if peer == nil {
			return
		}

		FetchPexAddrs(ctx, network, peer)
	})
}
