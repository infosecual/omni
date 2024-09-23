package module

import (
	"context"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
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

func Fuzz_Nosy_Module_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Module
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_Module_GetAuthority__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Module
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetAuthority()
	})
}

func Fuzz_Nosy_Module_GetConsensusTrimLag__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Module
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetConsensusTrimLag()
	})
}

func Fuzz_Nosy_Module_GetTrimLag__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Module
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetTrimLag()
	})
}

func Fuzz_Nosy_Module_GetVoteExtensionLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Module
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetVoteExtensionLimit()
	})
}

func Fuzz_Nosy_Module_GetVoteWindow__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Module
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetVoteWindow()
	})
}

func Fuzz_Nosy_Module_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Module
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_Module_ProtoReflect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Module
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.ProtoReflect()
	})
}

func Fuzz_Nosy_Module_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Module
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.Reset()
	})
}

func Fuzz_Nosy_Module_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Module
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.String()
	})
}

func Fuzz_Nosy_Module_slowProtoReflect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *Module
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.slowProtoReflect()
	})
}

// skipping Fuzz_Nosy_fastReflection_Module_Clear__ because parameters include func, chan, or unsupported interface: google.golang.org/protobuf/reflect/protoreflect.FieldDescriptor

func Fuzz_Nosy_fastReflection_Module_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *fastReflection_Module
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.Descriptor()
	})
}

// skipping Fuzz_Nosy_fastReflection_Module_Get__ because parameters include func, chan, or unsupported interface: google.golang.org/protobuf/reflect/protoreflect.FieldDescriptor

func Fuzz_Nosy_fastReflection_Module_GetUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *fastReflection_Module
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetUnknown()
	})
}

// skipping Fuzz_Nosy_fastReflection_Module_Has__ because parameters include func, chan, or unsupported interface: google.golang.org/protobuf/reflect/protoreflect.FieldDescriptor

func Fuzz_Nosy_fastReflection_Module_Interface__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *fastReflection_Module
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.Interface()
	})
}

func Fuzz_Nosy_fastReflection_Module_IsValid__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *fastReflection_Module
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.IsValid()
	})
}

// skipping Fuzz_Nosy_fastReflection_Module_Mutable__ because parameters include func, chan, or unsupported interface: google.golang.org/protobuf/reflect/protoreflect.FieldDescriptor

func Fuzz_Nosy_fastReflection_Module_New__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *fastReflection_Module
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.New()
	})
}

// skipping Fuzz_Nosy_fastReflection_Module_NewField__ because parameters include func, chan, or unsupported interface: google.golang.org/protobuf/reflect/protoreflect.FieldDescriptor

func Fuzz_Nosy_fastReflection_Module_ProtoMethods__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *fastReflection_Module
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.ProtoMethods()
	})
}

// skipping Fuzz_Nosy_fastReflection_Module_Range__ because parameters include func, chan, or unsupported interface: func(google.golang.org/protobuf/reflect/protoreflect.FieldDescriptor, google.golang.org/protobuf/reflect/protoreflect.Value) bool

// skipping Fuzz_Nosy_fastReflection_Module_Set__ because parameters include func, chan, or unsupported interface: google.golang.org/protobuf/reflect/protoreflect.FieldDescriptor

func Fuzz_Nosy_fastReflection_Module_SetUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *fastReflection_Module
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var fields protoreflect.RawFields
		fill_err = tp.Fill(&fields)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.SetUnknown(fields)
	})
}

func Fuzz_Nosy_fastReflection_Module_Type__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *fastReflection_Module
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.Type()
	})
}

// skipping Fuzz_Nosy_fastReflection_Module_WhichOneof__ because parameters include func, chan, or unsupported interface: google.golang.org/protobuf/reflect/protoreflect.OneofDescriptor

func Fuzz_Nosy_AppModule_BeginBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m AppModule
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		m.BeginBlock(ctx)
	})
}

func Fuzz_Nosy_AppModule_EndBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m AppModule
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		m.EndBlock(ctx)
	})
}

func Fuzz_Nosy_AppModule_IsAppModule__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 AppModule
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.IsAppModule()
	})
}

func Fuzz_Nosy_AppModule_IsOnePerModuleType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 AppModule
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.IsOnePerModuleType()
	})
}

// skipping Fuzz_Nosy_AppModule_RegisterServices__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/types/module.Configurator

func Fuzz_Nosy_AppModuleBasic_ConsensusVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 AppModuleBasic
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.ConsensusVersion()
	})
}

func Fuzz_Nosy_AppModuleBasic_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 AppModuleBasic
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_AppModuleBasic_RegisterGRPCGatewayRoutes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 AppModuleBasic
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 client.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 *runtime.ServeMux
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if _x3 == nil {
			return
		}

		_x1.RegisterGRPCGatewayRoutes(_x2, _x3)
	})
}

// skipping Fuzz_Nosy_AppModuleBasic_RegisterInterfaces__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/codec/types.InterfaceRegistry

func Fuzz_Nosy_AppModuleBasic_RegisterLegacyAminoCodec__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 AppModuleBasic
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 *codec.LegacyAmino
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x2 == nil {
			return
		}

		_x1.RegisterLegacyAminoCodec(_x2)
	})
}

func Fuzz_Nosy_fastReflection_Module_messageType_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x fastReflection_Module_messageType
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.Descriptor()
	})
}

func Fuzz_Nosy_fastReflection_Module_messageType_New__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x fastReflection_Module_messageType
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.New()
	})
}

func Fuzz_Nosy_fastReflection_Module_messageType_Zero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x fastReflection_Module_messageType
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.Zero()
	})
}
