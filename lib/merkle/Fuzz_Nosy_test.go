package merkle

import (
	"testing"

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

func Fuzz_Nosy_MakeTree__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var leaves [][32]byte
		fill_err = tp.Fill(&leaves)
		if fill_err != nil {
			return
		}

		MakeTree(leaves)
	})
}

func Fuzz_Nosy_StdLeafHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dst DomainSeparationTag
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}

		StdLeafHash(dst, d2)
	})
}

func Fuzz_Nosy_concatBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var buf [][]byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}

		concatBytes(buf...)
	})
}

func Fuzz_Nosy_hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, buf []byte) {
		hash(buf)
	})
}

func Fuzz_Nosy_hashPair__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a [32]byte
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b [32]byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		hashPair(a, b)
	})
}

func Fuzz_Nosy_isInternalNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tree [][32]byte
		fill_err = tp.Fill(&tree)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		isInternalNode(tree, i)
	})
}

func Fuzz_Nosy_isLeafNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tree [][32]byte
		fill_err = tp.Fill(&tree)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		isLeafNode(tree, i)
	})
}

func Fuzz_Nosy_isTreeNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tree [][32]byte
		fill_err = tp.Fill(&tree)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		isTreeNode(tree, i)
	})
}

func Fuzz_Nosy_leftChildIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i int) {
		leftChildIndex(i)
	})
}

func Fuzz_Nosy_parentIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i int) {
		parentIndex(i)
	})
}

func Fuzz_Nosy_rightChildIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i int) {
		rightChildIndex(i)
	})
}

func Fuzz_Nosy_siblingIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i int) {
		siblingIndex(i)
	})
}

func Fuzz_Nosy_sortBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var buf [][]byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}

		sortBytes(buf...)
	})
}
