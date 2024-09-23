package admin

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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

func Fuzz_Nosy_forgeRunner_run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r forgeRunner
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		var senders []common.Address
		fill_err = tp.Fill(&senders)
		if fill_err != nil {
			return
		}

		r.run(ctx, input, senders...)
	})
}

// skipping Fuzz_Nosy_runner_run__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/e2e/app/admin.runner

func Fuzz_Nosy_dedup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var strs []string
		fill_err = tp.Fill(&strs)
		if fill_err != nil {
			return
		}

		dedup(strs)
	})
}

func Fuzz_Nosy_execCmd__(f *testing.F) {
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
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var cmd string
		fill_err = tp.Fill(&cmd)
		if fill_err != nil {
			return
		}
		var args []string
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}

		execCmd(ctx, dir, cmd, args...)
	})
}

func Fuzz_Nosy_getChains__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var network netconf.Network
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var chain string
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}

		getChains(network, chain)
	})
}

// skipping Fuzz_Nosy_pausePortal__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/e2e/app/admin.runner

func Fuzz_Nosy_runForge__(f *testing.F) {
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
		var script string
		fill_err = tp.Fill(&script)
		if fill_err != nil {
			return
		}
		var rpc string
		fill_err = tp.Fill(&rpc)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		var senders []common.Address
		fill_err = tp.Fill(&senders)
		if fill_err != nil {
			return
		}

		runForge(ctx, script, rpc, input, senders...)
	})
}

func Fuzz_Nosy_startFBProxy__(f *testing.F) {
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
		var netID netconf.ID
		fill_err = tp.Fill(&netID)
		if fill_err != nil {
			return
		}
		var baseRPC string
		fill_err = tp.Fill(&baseRPC)
		if fill_err != nil {
			return
		}
		var fireAPIKey string
		fill_err = tp.Fill(&fireAPIKey)
		if fill_err != nil {
			return
		}
		var fireKeyPath string
		fill_err = tp.Fill(&fireKeyPath)
		if fill_err != nil {
			return
		}

		startFBProxy(ctx, netID, baseRPC, fireAPIKey, fireKeyPath)
	})
}

// skipping Fuzz_Nosy_unpausePortal__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/e2e/app/admin.runner

// skipping Fuzz_Nosy_upgradePortal__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/e2e/app/admin.runner
