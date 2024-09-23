package main

import (
	"go/ast"
	"go/types"
	"testing"

	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"golang.org/x/tools/go/ast/inspector"
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

func Fuzz_Nosy_attrIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var selct *ast.SelectorExpr
		fill_err = tp.Fill(&selct)
		if fill_err != nil {
			return
		}
		var selectors map[string]map[string]int
		fill_err = tp.Fill(&selectors)
		if fill_err != nil {
			return
		}
		if selct == nil {
			return
		}

		attrIndex(selct, selectors)
	})
}

func Fuzz_Nosy_isSnakeCase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, value string) {
		isSnakeCase(value)
	})
}

// skipping Fuzz_Nosy_isVar__ because parameters include func, chan, or unsupported interface: go/ast.Expr

func Fuzz_Nosy_logAttrs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i *inspector.Inspector
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if i == nil {
			return
		}

		logAttrs(i)
	})
}

func Fuzz_Nosy_uintstract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i *inspector.Inspector
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var typeInfo *types.Info
		fill_err = tp.Fill(&typeInfo)
		if fill_err != nil {
			return
		}
		if i == nil || typeInfo == nil {
			return
		}

		uintstract(i, typeInfo)
	})
}
