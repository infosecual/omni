package routerecon

import (
	"net/url"
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

func Fuzz_Nosy_chainID_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *chainID
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var bz []byte
		fill_err = tp.Fill(&bz)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.UnmarshalJSON(bz)
	})
}

func Fuzz_Nosy_chainID_ID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c chainID
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.ID()
	})
}

func Fuzz_Nosy_crossTxJSON_ConfLevel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c crossTxJSON
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.ConfLevel()
	})
}

func Fuzz_Nosy_crossTxJSON_ExpectedID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c crossTxJSON
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.ExpectedID()
	})
}

func Fuzz_Nosy_crossTxJSON_IsCompleted__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c crossTxJSON
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.IsCompleted()
	})
}

func Fuzz_Nosy_crossTxJSON_IsPending__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c crossTxJSON
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.IsPending()
	})
}

func Fuzz_Nosy_crossTxJSON_MsgID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c crossTxJSON
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.MsgID()
	})
}

func Fuzz_Nosy_crossTxJSON_Success__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c crossTxJSON
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.Success()
	})
}

func Fuzz_Nosy_crossTxJSON_Verify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c crossTxJSON
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.Verify()
	})
}

// skipping Fuzz_Nosy_filter_Match__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/routerecon.filter

// skipping Fuzz_Nosy_filter_QueryParams__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/routerecon.filter

func Fuzz_Nosy_queryFilter_HasStream__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 queryFilter
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}

		f1.HasStream()
	})
}

func Fuzz_Nosy_queryFilter_Match__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 queryFilter
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var crossTx crossTxJSON
		fill_err = tp.Fill(&crossTx)
		if fill_err != nil {
			return
		}

		f1.Match(crossTx)
	})
}

func Fuzz_Nosy_queryFilter_QueryParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 queryFilter
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var q url.Values
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}

		f1.QueryParams(q)
	})
}

// skipping Fuzz_Nosy_ReconForever__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/xchain.Provider

// skipping Fuzz_Nosy_queryLatestCrossTx__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/monitor/routerecon.filter

func Fuzz_Nosy_routeScanChainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id uint64) {
		routeScanChainID(id)
	})
}
