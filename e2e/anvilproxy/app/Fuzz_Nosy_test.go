package app

import (
	"context"
	"testing"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
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

// skipping Fuzz_Nosy_proxy_DisableFuzzyHead__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Nosy_proxy_EnableFuzzyHead__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

func Fuzz_Nosy_proxy_GetTarget__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var instance anvilInstance
		fill_err = tp.Fill(&instance)
		if fill_err != nil {
			return
		}

		p, err := newProxy(instance)
		if err != nil {
			return
		}
		p.GetTarget()
	})
}

func Fuzz_Nosy_proxy_IsFuzzyEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var instance anvilInstance
		fill_err = tp.Fill(&instance)
		if fill_err != nil {
			return
		}

		p, err := newProxy(instance)
		if err != nil {
			return
		}
		p.IsFuzzyEnabled()
	})
}

func Fuzz_Nosy_proxy_Perturb__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var instance anvilInstance
		fill_err = tp.Fill(&instance)
		if fill_err != nil {
			return
		}

		p, err := newProxy(instance)
		if err != nil {
			return
		}
		p.Perturb()
	})
}

// skipping Fuzz_Nosy_proxy_Proxy__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

func Fuzz_Nosy_proxy_getInstance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var instance anvilInstance
		fill_err = tp.Fill(&instance)
		if fill_err != nil {
			return
		}

		p, err := newProxy(instance)
		if err != nil {
			return
		}
		p.getInstance()
	})
}

// skipping Fuzz_Nosy_proxy_proxy__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

func Fuzz_Nosy_proxy_setPerturb__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var instance anvilInstance
		fill_err = tp.Fill(&instance)
		if fill_err != nil {
			return
		}
		var perturb types.Perturb
		fill_err = tp.Fill(&perturb)
		if fill_err != nil {
			return
		}

		p, err := newProxy(instance)
		if err != nil {
			return
		}
		p.setPerturb(perturb)
	})
}

func Fuzz_Nosy_proxy_setTarget__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var instance anvilInstance
		fill_err = tp.Fill(&instance)
		if fill_err != nil {
			return
		}
		var t2 anvilInstance
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}

		p, err := newProxy(instance)
		if err != nil {
			return
		}
		p.setTarget(t2)
	})
}

func Fuzz_Nosy_anvilInstance_AwaitHeight__(f *testing.F) {
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
		var cfg anvilConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var in context.Context
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		var min uint64
		fill_err = tp.Fill(&min)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}

		i, err := startAnvil(ctx, cfg)
		if err != nil {
			return
		}
		i.AwaitHeight(in, min, timeout)
	})
}

func Fuzz_Nosy_anvilInstance_Height__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg anvilConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}

		i, err := startAnvil(c1, cfg)
		if err != nil {
			return
		}
		i.Height(c3)
	})
}

func Fuzz_Nosy_anvilInstance_URL__(f *testing.F) {
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
		var cfg anvilConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}

		i, err := startAnvil(ctx, cfg)
		if err != nil {
			return
		}
		i.URL()
	})
}

func Fuzz_Nosy_closeReader_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 closeReader
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Close()
	})
}

func Fuzz_Nosy_fuzzXMsgs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var perturb types.Perturb
		fill_err = tp.Fill(&perturb)
		if fill_err != nil {
			return
		}
		var msgs []*bindings.OmniPortalXMsg
		fill_err = tp.Fill(&msgs)
		if fill_err != nil {
			return
		}

		fuzzXMsgs(perturb, msgs)
	})
}

func Fuzz_Nosy_isFuzzyXMsgLogFilter__(f *testing.F) {
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
		var perturb types.Perturb
		fill_err = tp.Fill(&perturb)
		if fill_err != nil {
			return
		}
		var t3 string
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var reqMsg jsonRPCMessage
		fill_err = tp.Fill(&reqMsg)
		if fill_err != nil {
			return
		}

		isFuzzyXMsgLogFilter(ctx, perturb, t3, reqMsg)
	})
}

func Fuzz_Nosy_parseAndFuzzXMsgs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var perturb types.Perturb
		fill_err = tp.Fill(&perturb)
		if fill_err != nil {
			return
		}
		var respBody []byte
		fill_err = tp.Fill(&respBody)
		if fill_err != nil {
			return
		}

		parseAndFuzzXMsgs(perturb, respBody)
	})
}
