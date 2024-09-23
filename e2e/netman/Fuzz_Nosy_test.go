package netman

import (
	"context"
	"testing"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
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

func Fuzz_Nosy_manager_DeployInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *manager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.DeployInfo()
	})
}

func Fuzz_Nosy_manager_DeployPrivatePortals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *manager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var valSetID uint64
		fill_err = tp.Fill(&valSetID)
		if fill_err != nil {
			return
		}
		var validators []bindings.XTypesValidator
		fill_err = tp.Fill(&validators)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.DeployPrivatePortals(ctx, valSetID, validators)
	})
}

func Fuzz_Nosy_manager_DeployPublicPortals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *manager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var valSetID uint64
		fill_err = tp.Fill(&valSetID)
		if fill_err != nil {
			return
		}
		var validators []bindings.XTypesValidator
		fill_err = tp.Fill(&validators)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.DeployPublicPortals(ctx, valSetID, validators)
	})
}

func Fuzz_Nosy_manager_Portals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *manager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Portals()
	})
}

func Fuzz_Nosy_manager_deployIfNeeded__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *manager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chain types.EVMChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var backend *ethbackend.Backend
		fill_err = tp.Fill(&backend)
		if fill_err != nil {
			return
		}
		var valSetID uint64
		fill_err = tp.Fill(&valSetID)
		if fill_err != nil {
			return
		}
		var validators []bindings.XTypesValidator
		fill_err = tp.Fill(&validators)
		if fill_err != nil {
			return
		}
		if m == nil || backend == nil {
			return
		}

		m.deployIfNeeded(ctx, chain, backend, valSetID, validators)
	})
}

func Fuzz_Nosy_manager_destChainIDs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *manager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var srcChainID uint64
		fill_err = tp.Fill(&srcChainID)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.destChainIDs(srcChainID)
	})
}

func Fuzz_Nosy_Manager_DeployInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var testnet types.Testnet
		fill_err = tp.Fill(&testnet)
		if fill_err != nil {
			return
		}
		var backends ethbackend.Backends
		fill_err = tp.Fill(&backends)
		if fill_err != nil {
			return
		}

		_x1, err := NewManager(testnet, backends)
		if err != nil {
			return
		}
		_x1.DeployInfo()
	})
}

func Fuzz_Nosy_Manager_DeployPrivatePortals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var testnet types.Testnet
		fill_err = tp.Fill(&testnet)
		if fill_err != nil {
			return
		}
		var backends ethbackend.Backends
		fill_err = tp.Fill(&backends)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var valSetID uint64
		fill_err = tp.Fill(&valSetID)
		if fill_err != nil {
			return
		}
		var validators []bindings.XTypesValidator
		fill_err = tp.Fill(&validators)
		if fill_err != nil {
			return
		}

		_x1, err := NewManager(testnet, backends)
		if err != nil {
			return
		}
		_x1.DeployPrivatePortals(ctx, valSetID, validators)
	})
}

func Fuzz_Nosy_Manager_DeployPublicPortals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var testnet types.Testnet
		fill_err = tp.Fill(&testnet)
		if fill_err != nil {
			return
		}
		var backends ethbackend.Backends
		fill_err = tp.Fill(&backends)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var valSetID uint64
		fill_err = tp.Fill(&valSetID)
		if fill_err != nil {
			return
		}
		var validators []bindings.XTypesValidator
		fill_err = tp.Fill(&validators)
		if fill_err != nil {
			return
		}

		_x1, err := NewManager(testnet, backends)
		if err != nil {
			return
		}
		_x1.DeployPublicPortals(ctx, valSetID, validators)
	})
}

func Fuzz_Nosy_Manager_Portals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var testnet types.Testnet
		fill_err = tp.Fill(&testnet)
		if fill_err != nil {
			return
		}
		var backends ethbackend.Backends
		fill_err = tp.Fill(&backends)
		if fill_err != nil {
			return
		}

		_x1, err := NewManager(testnet, backends)
		if err != nil {
			return
		}
		_x1.Portals()
	})
}
