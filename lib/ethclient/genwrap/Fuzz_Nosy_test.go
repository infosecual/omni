package main

import (
	"testing"

	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"golang.org/x/tools/go/packages"
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

func Fuzz_Nosy_Method_Label__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m Method
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}

		m.Label()
	})
}

func Fuzz_Nosy_Method_NamedResults__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m Method
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}

		m.NamedResults()
	})
}

func Fuzz_Nosy_Method_ParamNames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m Method
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}

		m.ParamNames()
	})
}

func Fuzz_Nosy_Method_Params__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m Method
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}

		m.Params()
	})
}

func Fuzz_Nosy_Method_ResultNames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m Method
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}

		m.ResultNames()
	})
}

func Fuzz_Nosy_Method_ResultTypes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m Method
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}

		m.ResultTypes()
	})
}

func Fuzz_Nosy_Method_Results__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m Method
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}

		m.Results()
	})
}

func Fuzz_Nosy_parseEthMethods__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pkg *packages.Package
		fill_err = tp.Fill(&pkg)
		if fill_err != nil {
			return
		}
		if pkg == nil {
			return
		}

		parseEthMethods(pkg)
	})
}

func Fuzz_Nosy_parseImports__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pkg *packages.Package
		fill_err = tp.Fill(&pkg)
		if fill_err != nil {
			return
		}
		if pkg == nil {
			return
		}

		parseImports(pkg)
	})
}

func Fuzz_Nosy_toSnakeCase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, str string) {
		toSnakeCase(str)
	})
}
