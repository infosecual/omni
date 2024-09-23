package docker

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
		var infd types.InfrastructureData
		fill_err = tp.Fill(&infd)
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

		_x1 := NewProvider(testnet, infd, imgTag)
		_x1.Clean(ctx)
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
		var infd types.InfrastructureData
		fill_err = tp.Fill(&infd)
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
		var _x5 types.ServiceConfig
		fill_err = tp.Fill(&_x5)
		if fill_err != nil {
			return
		}

		_x1 := NewProvider(testnet, infd, imgTag)
		_x1.Restart(_x4, _x5)
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
		var infd types.InfrastructureData
		fill_err = tp.Fill(&infd)
		if fill_err != nil {
			return
		}
		var imgTag string
		fill_err = tp.Fill(&imgTag)
		if fill_err != nil {
			return
		}

		p := NewProvider(testnet, infd, imgTag)
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
		var infd types.InfrastructureData
		fill_err = tp.Fill(&infd)
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
		var nodes []*e2e.Node
		fill_err = tp.Fill(&nodes)
		if fill_err != nil {
			return
		}

		p := NewProvider(testnet, infd, imgTag)
		p.StartNodes(ctx, nodes...)
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
		var infd types.InfrastructureData
		fill_err = tp.Fill(&infd)
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
		var _x5 types.ServiceConfig
		fill_err = tp.Fill(&_x5)
		if fill_err != nil {
			return
		}

		_x1 := NewProvider(testnet, infd, imgTag)
		_x1.Upgrade(_x4, _x5)
	})
}

func Fuzz_Nosy_ComposeDef_GethTag__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var def ComposeDef
		fill_err = tp.Fill(&def)
		if fill_err != nil {
			return
		}
		var manifest types.Manifest
		fill_err = tp.Fill(&manifest)
		if fill_err != nil {
			return
		}
		var omniImgTag string
		fill_err = tp.Fill(&omniImgTag)
		if fill_err != nil {
			return
		}

		_x1 := SetImageTags(def, manifest, omniImgTag)
		_x1.GethTag()
	})
}

func Fuzz_Nosy_ComposeDef_NodeOmniEVMs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var def ComposeDef
		fill_err = tp.Fill(&def)
		if fill_err != nil {
			return
		}
		var manifest types.Manifest
		fill_err = tp.Fill(&manifest)
		if fill_err != nil {
			return
		}
		var omniImgTag string
		fill_err = tp.Fill(&omniImgTag)
		if fill_err != nil {
			return
		}

		c := SetImageTags(def, manifest, omniImgTag)
		c.NodeOmniEVMs()
	})
}

func Fuzz_Nosy_CleanCmds__(f *testing.F) {
	f.Fuzz(func(t *testing.T, sudo bool, isLinux bool) {
		CleanCmds(sudo, isLinux)
	})
}

func Fuzz_Nosy_GenerateComposeFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var def ComposeDef
		fill_err = tp.Fill(&def)
		if fill_err != nil {
			return
		}

		GenerateComposeFile(def)
	})
}

func Fuzz_Nosy_GenerateOmniEVMInitFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var def ComposeDef
		fill_err = tp.Fill(&def)
		if fill_err != nil {
			return
		}

		GenerateOmniEVMInitFile(def)
	})
}

func Fuzz_Nosy_additionalServices__(f *testing.F) {
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

		additionalServices(testnet)
	})
}
