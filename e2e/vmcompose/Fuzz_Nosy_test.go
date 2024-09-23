package vmcompose

import (
	"context"
	"testing"

	"github.com/omni-network/omni/e2e"
	"github.com/omni-network/omni/e2e/types"
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

func Fuzz_Nosy_Provider_Clean__(f *testing.F) {
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
		var d2 types.InfrastructureData
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		var imgTag string
		fill_err = tp.Fill(&imgTag)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		p := NewProvider(testnet, d2, imgTag)
		p.Clean(ctx)
	})
}

func Fuzz_Nosy_Provider_GetInfrastructureData__(f *testing.F) {
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
		var d2 types.InfrastructureData
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		var imgTag string
		fill_err = tp.Fill(&imgTag)
		if fill_err != nil {
			return
		}

		p := NewProvider(testnet, d2, imgTag)
		p.GetInfrastructureData()
	})
}

func Fuzz_Nosy_Provider_Restart__(f *testing.F) {
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
		var d2 types.InfrastructureData
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		var imgTag string
		fill_err = tp.Fill(&imgTag)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg types.ServiceConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}

		p := NewProvider(testnet, d2, imgTag)
		p.Restart(ctx, cfg)
	})
}

func Fuzz_Nosy_Provider_Setup__(f *testing.F) {
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
		var d2 types.InfrastructureData
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		var imgTag string
		fill_err = tp.Fill(&imgTag)
		if fill_err != nil {
			return
		}

		p := NewProvider(testnet, d2, imgTag)
		p.Setup()
	})
}

func Fuzz_Nosy_Provider_StartNodes__(f *testing.F) {
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
		var d2 types.InfrastructureData
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		var imgTag string
		fill_err = tp.Fill(&imgTag)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var _x5 []*e2e.Node
		fill_err = tp.Fill(&_x5)
		if fill_err != nil {
			return
		}

		p := NewProvider(testnet, d2, imgTag)
		p.StartNodes(ctx, _x5...)
	})
}

func Fuzz_Nosy_Provider_StopTestnet__(f *testing.F) {
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
		var d2 types.InfrastructureData
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		var imgTag string
		fill_err = tp.Fill(&imgTag)
		if fill_err != nil {
			return
		}
		var _x4 context.Context
		fill_err = tp.Fill(&_x4)
		if fill_err != nil {
			return
		}

		_x1 := NewProvider(testnet, d2, imgTag)
		_x1.StopTestnet(_x4)
	})
}

func Fuzz_Nosy_Provider_Upgrade__(f *testing.F) {
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
		var d2 types.InfrastructureData
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		var imgTag string
		fill_err = tp.Fill(&imgTag)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg types.ServiceConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}

		p := NewProvider(testnet, d2, imgTag)
		p.Upgrade(ctx, cfg)
	})
}

func Fuzz_Nosy_groupByVM__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var instances map[string]e2e.InstanceData
		fill_err = tp.Fill(&instances)
		if fill_err != nil {
			return
		}

		groupByVM(instances)
	})
}

func Fuzz_Nosy_matchAny__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg types.ServiceConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var services map[string]bool
		fill_err = tp.Fill(&services)
		if fill_err != nil {
			return
		}

		matchAny(cfg, services)
	})
}

func Fuzz_Nosy_vmAgentFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, internalIP string) {
		vmAgentFile(internalIP)
	})
}

func Fuzz_Nosy_vmComposeFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, internalIP string) {
		vmComposeFile(internalIP)
	})
}

func Fuzz_Nosy_vmInitFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, internalIP string) {
		vmInitFile(internalIP)
	})
}
