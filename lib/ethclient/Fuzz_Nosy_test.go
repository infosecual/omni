package ethclient

import (
	"context"
	"math/big"
	"net/http"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/omni-network/omni/lib/netconf"
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

func Fuzz_Nosy_engineMock_BlockByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *engineMock
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var number *big.Int
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if m == nil || number == nil {
			return
		}

		m.BlockByNumber(ctx, number)
	})
}

func Fuzz_Nosy_engineMock_BlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *engineMock
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.BlockNumber(ctx)
	})
}

func Fuzz_Nosy_engineMock_FilterLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *engineMock
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var q ethereum.FilterQuery
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.FilterLogs(_x2, q)
	})
}

func Fuzz_Nosy_engineMock_ForkchoiceUpdatedV2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *engineMock
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 engine.ForkchoiceStateV1
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		var _x4 *engine.PayloadAttributes
		fill_err = tp.Fill(&_x4)
		if fill_err != nil {
			return
		}
		if _x1 == nil || _x4 == nil {
			return
		}

		_x1.ForkchoiceUpdatedV2(_x2, _x3, _x4)
	})
}

func Fuzz_Nosy_engineMock_ForkchoiceUpdatedV3__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *engineMock
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var update engine.ForkchoiceStateV1
		fill_err = tp.Fill(&update)
		if fill_err != nil {
			return
		}
		var attrs *engine.PayloadAttributes
		fill_err = tp.Fill(&attrs)
		if fill_err != nil {
			return
		}
		if m == nil || attrs == nil {
			return
		}

		m.ForkchoiceUpdatedV3(ctx, update, attrs)
	})
}

func Fuzz_Nosy_engineMock_GetPayloadV2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *engineMock
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 engine.PayloadID
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.GetPayloadV2(_x2, _x3)
	})
}

func Fuzz_Nosy_engineMock_GetPayloadV3__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *engineMock
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var payloadID engine.PayloadID
		fill_err = tp.Fill(&payloadID)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetPayloadV3(ctx, payloadID)
	})
}

func Fuzz_Nosy_engineMock_HeaderByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *engineMock
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.HeaderByHash(ctx, hash)
	})
}

func Fuzz_Nosy_engineMock_HeaderByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *engineMock
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var height *big.Int
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if m == nil || height == nil {
			return
		}

		m.HeaderByNumber(ctx, height)
	})
}

func Fuzz_Nosy_engineMock_HeaderByType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *engineMock
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var typ HeadType
		fill_err = tp.Fill(&typ)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.HeaderByType(ctx, typ)
	})
}

func Fuzz_Nosy_engineMock_NewPayloadV2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *engineMock
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 engine.ExecutableData
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.NewPayloadV2(_x2, _x3)
	})
}

func Fuzz_Nosy_engineMock_NewPayloadV3__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *engineMock
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var params engine.ExecutableData
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var _x4 []common.Hash
		fill_err = tp.Fill(&_x4)
		if fill_err != nil {
			return
		}
		var beaconRoot *common.Hash
		fill_err = tp.Fill(&beaconRoot)
		if fill_err != nil {
			return
		}
		if m == nil || beaconRoot == nil {
			return
		}

		m.NewPayloadV3(ctx, params, _x4, beaconRoot)
	})
}

func Fuzz_Nosy_engineMock_PeerCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *engineMock
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.PeerCount(_x2)
	})
}

func Fuzz_Nosy_engineMock_SyncProgress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *engineMock
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.SyncProgress(_x2)
	})
}

func Fuzz_Nosy_engineMock_maybeErr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *engineMock
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.maybeErr(ctx)
	})
}

func Fuzz_Nosy_jwtRoundTripper_RoundTrip__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *jwtRoundTripper
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var req *http.Request
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if t1 == nil || req == nil {
			return
		}

		t1.RoundTrip(req)
	})
}

// skipping Fuzz_Nosy_Client_Address__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/ethclient.Client

// skipping Fuzz_Nosy_Client_Close__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/ethclient.Client

// skipping Fuzz_Nosy_Client_EtherBalanceAt__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/ethclient.Client

// skipping Fuzz_Nosy_Client_HeaderByType__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/ethclient.Client

// skipping Fuzz_Nosy_Client_PeerCount__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/ethclient.Client

// skipping Fuzz_Nosy_Client_SetHead__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/lib/ethclient.Client

func Fuzz_Nosy_EngineClient_ForkchoiceUpdatedV2__(f *testing.F) {
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
		var urlAddr string
		fill_err = tp.Fill(&urlAddr)
		if fill_err != nil {
			return
		}
		var jwtSecret []byte
		fill_err = tp.Fill(&jwtSecret)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var update engine.ForkchoiceStateV1
		fill_err = tp.Fill(&update)
		if fill_err != nil {
			return
		}
		var payloadAttributes *engine.PayloadAttributes
		fill_err = tp.Fill(&payloadAttributes)
		if fill_err != nil {
			return
		}
		if payloadAttributes == nil {
			return
		}

		_x1, err := NewAuthClient(c1, urlAddr, jwtSecret)
		if err != nil {
			return
		}
		_x1.ForkchoiceUpdatedV2(c4, update, payloadAttributes)
	})
}

func Fuzz_Nosy_EngineClient_ForkchoiceUpdatedV3__(f *testing.F) {
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
		var urlAddr string
		fill_err = tp.Fill(&urlAddr)
		if fill_err != nil {
			return
		}
		var jwtSecret []byte
		fill_err = tp.Fill(&jwtSecret)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var update engine.ForkchoiceStateV1
		fill_err = tp.Fill(&update)
		if fill_err != nil {
			return
		}
		var payloadAttributes *engine.PayloadAttributes
		fill_err = tp.Fill(&payloadAttributes)
		if fill_err != nil {
			return
		}
		if payloadAttributes == nil {
			return
		}

		_x1, err := NewAuthClient(c1, urlAddr, jwtSecret)
		if err != nil {
			return
		}
		_x1.ForkchoiceUpdatedV3(c4, update, payloadAttributes)
	})
}

func Fuzz_Nosy_EngineClient_GetPayloadV2__(f *testing.F) {
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
		var urlAddr string
		fill_err = tp.Fill(&urlAddr)
		if fill_err != nil {
			return
		}
		var jwtSecret []byte
		fill_err = tp.Fill(&jwtSecret)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var payloadID engine.PayloadID
		fill_err = tp.Fill(&payloadID)
		if fill_err != nil {
			return
		}

		_x1, err := NewAuthClient(c1, urlAddr, jwtSecret)
		if err != nil {
			return
		}
		_x1.GetPayloadV2(c4, payloadID)
	})
}

func Fuzz_Nosy_EngineClient_GetPayloadV3__(f *testing.F) {
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
		var urlAddr string
		fill_err = tp.Fill(&urlAddr)
		if fill_err != nil {
			return
		}
		var jwtSecret []byte
		fill_err = tp.Fill(&jwtSecret)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var payloadID engine.PayloadID
		fill_err = tp.Fill(&payloadID)
		if fill_err != nil {
			return
		}

		_x1, err := NewAuthClient(c1, urlAddr, jwtSecret)
		if err != nil {
			return
		}
		_x1.GetPayloadV3(c4, payloadID)
	})
}

func Fuzz_Nosy_EngineClient_NewPayloadV2__(f *testing.F) {
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
		var urlAddr string
		fill_err = tp.Fill(&urlAddr)
		if fill_err != nil {
			return
		}
		var jwtSecret []byte
		fill_err = tp.Fill(&jwtSecret)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var params engine.ExecutableData
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}

		_x1, err := NewAuthClient(c1, urlAddr, jwtSecret)
		if err != nil {
			return
		}
		_x1.NewPayloadV2(c4, params)
	})
}

func Fuzz_Nosy_EngineClient_NewPayloadV3__(f *testing.F) {
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
		var urlAddr string
		fill_err = tp.Fill(&urlAddr)
		if fill_err != nil {
			return
		}
		var jwtSecret []byte
		fill_err = tp.Fill(&jwtSecret)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var params engine.ExecutableData
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var versionedHashes []common.Hash
		fill_err = tp.Fill(&versionedHashes)
		if fill_err != nil {
			return
		}
		var beaconRoot *common.Hash
		fill_err = tp.Fill(&beaconRoot)
		if fill_err != nil {
			return
		}
		if beaconRoot == nil {
			return
		}

		_x1, err := NewAuthClient(c1, urlAddr, jwtSecret)
		if err != nil {
			return
		}
		_x1.NewPayloadV3(c4, params, versionedHashes, beaconRoot)
	})
}

func Fuzz_Nosy_HeadType_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h HeadType
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.String()
	})
}

func Fuzz_Nosy_HeadType_Verify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h HeadType
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.Verify()
	})
}

func Fuzz_Nosy_Wrapper_Address__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainName string, url string) {
		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.Address()
	})
}

func Fuzz_Nosy_Wrapper_BalanceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if blockNumber == nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.BalanceAt(ctx, account, blockNumber)
	})
}

func Fuzz_Nosy_Wrapper_BlockByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.BlockByHash(ctx, hash)
	})
}

func Fuzz_Nosy_Wrapper_BlockByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var number *big.Int
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if number == nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.BlockByNumber(ctx, number)
	})
}

func Fuzz_Nosy_Wrapper_BlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.BlockNumber(ctx)
	})
}

func Fuzz_Nosy_Wrapper_CallContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var call ethereum.CallMsg
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if blockNumber == nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.CallContract(ctx, call, blockNumber)
	})
}

func Fuzz_Nosy_Wrapper_ChainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.ChainID(ctx)
	})
}

func Fuzz_Nosy_Wrapper_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainName string, url string) {
		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.Close()
	})
}

func Fuzz_Nosy_Wrapper_CodeAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if blockNumber == nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.CodeAt(ctx, account, blockNumber)
	})
}

func Fuzz_Nosy_Wrapper_EstimateGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var call ethereum.CallMsg
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.EstimateGas(ctx, call)
	})
}

func Fuzz_Nosy_Wrapper_EtherBalanceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.EtherBalanceAt(ctx, addr)
	})
}

func Fuzz_Nosy_Wrapper_FilterLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var q ethereum.FilterQuery
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.FilterLogs(ctx, q)
	})
}

func Fuzz_Nosy_Wrapper_HeaderByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.HeaderByHash(ctx, hash)
	})
}

func Fuzz_Nosy_Wrapper_HeaderByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var number *big.Int
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if number == nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.HeaderByNumber(ctx, number)
	})
}

func Fuzz_Nosy_Wrapper_HeaderByType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var typ HeadType
		fill_err = tp.Fill(&typ)
		if fill_err != nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.HeaderByType(ctx, typ)
	})
}

func Fuzz_Nosy_Wrapper_NonceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if blockNumber == nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.NonceAt(ctx, account, blockNumber)
	})
}

func Fuzz_Nosy_Wrapper_PeerCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.PeerCount(ctx)
	})
}

func Fuzz_Nosy_Wrapper_PendingBalanceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.PendingBalanceAt(ctx, account)
	})
}

func Fuzz_Nosy_Wrapper_PendingCodeAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.PendingCodeAt(ctx, account)
	})
}

func Fuzz_Nosy_Wrapper_PendingNonceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.PendingNonceAt(ctx, account)
	})
}

func Fuzz_Nosy_Wrapper_PendingStorageAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.PendingStorageAt(ctx, account, key)
	})
}

func Fuzz_Nosy_Wrapper_PendingTransactionCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.PendingTransactionCount(ctx)
	})
}

func Fuzz_Nosy_Wrapper_SendTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.SendTransaction(ctx, tx)
	})
}

func Fuzz_Nosy_Wrapper_SetHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.SetHead(ctx, height)
	})
}

func Fuzz_Nosy_Wrapper_StorageAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if blockNumber == nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.StorageAt(ctx, account, key, blockNumber)
	})
}

// skipping Fuzz_Nosy_Wrapper_SubscribeFilterLogs__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/core/types.Log

// skipping Fuzz_Nosy_Wrapper_SubscribeNewHead__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum/go-ethereum/core/types.Header

func Fuzz_Nosy_Wrapper_SuggestGasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.SuggestGasPrice(ctx)
	})
}

func Fuzz_Nosy_Wrapper_SuggestGasTipCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.SuggestGasTipCap(ctx)
	})
}

func Fuzz_Nosy_Wrapper_SyncProgress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.SyncProgress(ctx)
	})
}

func Fuzz_Nosy_Wrapper_TransactionByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var txHash common.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.TransactionByHash(ctx, txHash)
	})
}

func Fuzz_Nosy_Wrapper_TransactionCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.TransactionCount(ctx, blockHash)
	})
}

func Fuzz_Nosy_Wrapper_TransactionInBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		var index uint
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.TransactionInBlock(ctx, blockHash, index)
	})
}

func Fuzz_Nosy_Wrapper_TransactionReceipt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var txHash common.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}

		w, err := Dial(chainName, url)
		if err != nil {
			return
		}
		w.TransactionReceipt(ctx, txHash)
	})
}

func Fuzz_Nosy_engineClient_ForkchoiceUpdatedV2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c engineClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var update engine.ForkchoiceStateV1
		fill_err = tp.Fill(&update)
		if fill_err != nil {
			return
		}
		var payloadAttributes *engine.PayloadAttributes
		fill_err = tp.Fill(&payloadAttributes)
		if fill_err != nil {
			return
		}
		if payloadAttributes == nil {
			return
		}

		c.ForkchoiceUpdatedV2(ctx, update, payloadAttributes)
	})
}

func Fuzz_Nosy_engineClient_ForkchoiceUpdatedV3__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c engineClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var update engine.ForkchoiceStateV1
		fill_err = tp.Fill(&update)
		if fill_err != nil {
			return
		}
		var payloadAttributes *engine.PayloadAttributes
		fill_err = tp.Fill(&payloadAttributes)
		if fill_err != nil {
			return
		}
		if payloadAttributes == nil {
			return
		}

		c.ForkchoiceUpdatedV3(ctx, update, payloadAttributes)
	})
}

func Fuzz_Nosy_engineClient_GetPayloadV2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c engineClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var payloadID engine.PayloadID
		fill_err = tp.Fill(&payloadID)
		if fill_err != nil {
			return
		}

		c.GetPayloadV2(ctx, payloadID)
	})
}

func Fuzz_Nosy_engineClient_GetPayloadV3__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c engineClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var payloadID engine.PayloadID
		fill_err = tp.Fill(&payloadID)
		if fill_err != nil {
			return
		}

		c.GetPayloadV3(ctx, payloadID)
	})
}

func Fuzz_Nosy_engineClient_NewPayloadV2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c engineClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var params engine.ExecutableData
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}

		c.NewPayloadV2(ctx, params)
	})
}

func Fuzz_Nosy_engineClient_NewPayloadV3__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c engineClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var params engine.ExecutableData
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var versionedHashes []common.Hash
		fill_err = tp.Fill(&versionedHashes)
		if fill_err != nil {
			return
		}
		var beaconRoot *common.Hash
		fill_err = tp.Fill(&beaconRoot)
		if fill_err != nil {
			return
		}
		if beaconRoot == nil {
			return
		}

		c.NewPayloadV3(ctx, params, versionedHashes, beaconRoot)
	})
}

// skipping Fuzz_Nosy_IsErrMethodNotAvailable__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_LoadJWTHexFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string) {
		LoadJWTHexFile(file)
	})
}

// skipping Fuzz_Nosy_WithMockSelfDelegation__ because parameters include func, chan, or unsupported interface: github.com/cometbft/cometbft/crypto.PubKey

func Fuzz_Nosy_WithPortalRegister__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var network netconf.Network
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}

		WithPortalRegister(network)
	})
}

func Fuzz_Nosy_hasRandomErr__(f *testing.F) {
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

		hasRandomErr(ctx)
	})
}

func Fuzz_Nosy_incError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chain string, endpoint string) {
		incError(chain, endpoint)
	})
}

func Fuzz_Nosy_latency__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chain string, endpoint string) {
		latency(chain, endpoint)
	})
}

func Fuzz_Nosy_spanName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, endpoint string) {
		spanName(endpoint)
	})
}
