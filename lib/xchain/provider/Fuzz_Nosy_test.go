package provider

import (
	"context"
	"io"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/omni-network/omni/lib/xchain"
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

func Fuzz_Nosy_Mock_ChainVersionHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Mock
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 xchain.ChainVersion
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ChainVersionHeight(_x2, _x3)
	})
}

func Fuzz_Nosy_Mock_GetBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Mock
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var req xchain.ProviderRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetBlock(_x2, req)
	})
}

func Fuzz_Nosy_Mock_GetEmittedCursor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Mock
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 xchain.EmitRef
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		var stream xchain.StreamID
		fill_err = tp.Fill(&stream)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.GetEmittedCursor(_x2, _x3, stream)
	})
}

func Fuzz_Nosy_Mock_GetSubmission__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Mock
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 uint64
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		var _x4 common.Hash
		fill_err = tp.Fill(&_x4)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.GetSubmission(_x2, _x3, _x4)
	})
}

func Fuzz_Nosy_Mock_GetSubmittedCursor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Mock
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var stream xchain.StreamID
		fill_err = tp.Fill(&stream)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.GetSubmittedCursor(_x2, stream)
	})
}

// skipping Fuzz_Nosy_Mock_StreamAsync__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/xchain.ProviderCallback

// skipping Fuzz_Nosy_Mock_StreamBlocks__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/xchain.ProviderCallback

func Fuzz_Nosy_Mock_addBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Mock
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var block xchain.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var confLevel xchain.ConfLevel
		fill_err = tp.Fill(&confLevel)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.addBlock(block, confLevel)
	})
}

// skipping Fuzz_Nosy_Mock_nextBlock__ because parameters include func, chan, or unsupported interface: func(github.com/omni-network/omni/lib/xchain.StreamID) uint64

func Fuzz_Nosy_Mock_parentBlockHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Mock
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var chainVer xchain.ChainVersion
		fill_err = tp.Fill(&chainVer)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.parentBlockHash(chainVer, height)
	})
}

// skipping Fuzz_Nosy_Mock_stream__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/xchain.ProviderCallback

func Fuzz_Nosy_Provider_ChainVersionHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Provider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainVer xchain.ChainVersion
		fill_err = tp.Fill(&chainVer)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.ChainVersionHeight(ctx, chainVer)
	})
}

func Fuzz_Nosy_Provider_GetBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Provider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req xchain.ProviderRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetBlock(ctx, req)
	})
}

func Fuzz_Nosy_Provider_GetEmittedCursor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Provider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var ref xchain.EmitRef
		fill_err = tp.Fill(&ref)
		if fill_err != nil {
			return
		}
		var stream xchain.StreamID
		fill_err = tp.Fill(&stream)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetEmittedCursor(ctx, ref, stream)
	})
}

func Fuzz_Nosy_Provider_GetSubmission__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Provider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var txHash common.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetSubmission(ctx, chainID, txHash)
	})
}

func Fuzz_Nosy_Provider_GetSubmittedCursor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Provider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stream xchain.StreamID
		fill_err = tp.Fill(&stream)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetSubmittedCursor(ctx, stream)
	})
}

// skipping Fuzz_Nosy_Provider_StreamAsync__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/xchain.ProviderCallback

// skipping Fuzz_Nosy_Provider_StreamBlocks__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/xchain.ProviderCallback

func Fuzz_Nosy_Provider_confirmedCache__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Provider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var chainVer xchain.ChainVersion
		fill_err = tp.Fill(&chainVer)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.confirmedCache(chainVer, height)
	})
}

func Fuzz_Nosy_Provider_getEVMChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Provider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var chainID uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.getEVMChain(chainID)
	})
}

func Fuzz_Nosy_Provider_getXMsgLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Provider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.getXMsgLogs(ctx, chainID, blockHash)
	})
}

func Fuzz_Nosy_Provider_getXReceiptLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Provider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.getXReceiptLogs(ctx, chainID, blockHash)
	})
}

func Fuzz_Nosy_Provider_headerByChainVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Provider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainVer xchain.ChainVersion
		fill_err = tp.Fill(&chainVer)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.headerByChainVersion(ctx, chainVer)
	})
}

// skipping Fuzz_Nosy_Provider_stream__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/xchain.ProviderCallback

func Fuzz_Nosy_streamOffseter_offset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o streamOffseter
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var id xchain.StreamID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		o.offset(id)
	})
}

// skipping Fuzz_Nosy_getLogs__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/ethclient.Client

func Fuzz_Nosy_headTypeFromConfLevel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf xchain.ConfLevel
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}

		headTypeFromConfLevel(conf)
	})
}

func Fuzz_Nosy_random20__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}

		random20(r)
	})
}

func Fuzz_Nosy_random32__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}

		random32(r)
	})
}

func Fuzz_Nosy_spanName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, method string) {
		spanName(method)
	})
}
