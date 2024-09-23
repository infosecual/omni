package bindings

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

func Fuzz_Nosy_AVSDirectoryAVSMetadataURIUpdatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *AVSDirectoryAVSMetadataURIUpdatedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_AVSDirectoryAVSMetadataURIUpdatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *AVSDirectoryAVSMetadataURIUpdatedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_AVSDirectoryAVSMetadataURIUpdatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *AVSDirectoryAVSMetadataURIUpdatedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_AVSDirectoryCaller_AvsOperatorStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryCaller
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 common.Address
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil {
			return
		}

		_AVSDirectory.AvsOperatorStatus(opts, arg0, arg1)
	})
}

func Fuzz_Nosy_AVSDirectoryCaller_CalculateOperatorAVSRegistrationDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryCaller
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var avs common.Address
		fill_err = tp.Fill(&avs)
		if fill_err != nil {
			return
		}
		var salt [32]byte
		fill_err = tp.Fill(&salt)
		if fill_err != nil {
			return
		}
		var expiry *big.Int
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil || expiry == nil {
			return
		}

		_AVSDirectory.CalculateOperatorAVSRegistrationDigestHash(opts, operator, avs, salt, expiry)
	})
}

func Fuzz_Nosy_AVSDirectoryCaller_DOMAINTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryCaller
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil {
			return
		}

		_AVSDirectory.DOMAINTYPEHASH(opts)
	})
}

func Fuzz_Nosy_AVSDirectoryCaller_Delegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryCaller
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil {
			return
		}

		_AVSDirectory.Delegation(opts)
	})
}

func Fuzz_Nosy_AVSDirectoryCaller_DomainSeparator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryCaller
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil {
			return
		}

		_AVSDirectory.DomainSeparator(opts)
	})
}

func Fuzz_Nosy_AVSDirectoryCaller_OPERATORAVSREGISTRATIONTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryCaller
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil {
			return
		}

		_AVSDirectory.OPERATORAVSREGISTRATIONTYPEHASH(opts)
	})
}

func Fuzz_Nosy_AVSDirectoryCaller_OperatorSaltIsSpent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryCaller
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 [32]byte
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil {
			return
		}

		_AVSDirectory.OperatorSaltIsSpent(opts, arg0, arg1)
	})
}

func Fuzz_Nosy_AVSDirectoryCaller_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryCaller
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil {
			return
		}

		_AVSDirectory.Owner(opts)
	})
}

func Fuzz_Nosy_AVSDirectoryCaller_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryCaller
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var index uint8
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil {
			return
		}

		_AVSDirectory.Paused(opts, index)
	})
}

func Fuzz_Nosy_AVSDirectoryCaller_Paused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryCaller
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil {
			return
		}

		_AVSDirectory.Paused0(opts)
	})
}

func Fuzz_Nosy_AVSDirectoryCaller_PauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryCaller
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil {
			return
		}

		_AVSDirectory.PauserRegistry(opts)
	})
}

// skipping Fuzz_Nosy_AVSDirectoryCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_AVSDirectoryCallerSession_AvsOperatorStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryCallerSession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 common.Address
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.AvsOperatorStatus(arg0, arg1)
	})
}

func Fuzz_Nosy_AVSDirectoryCallerSession_CalculateOperatorAVSRegistrationDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryCallerSession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var avs common.Address
		fill_err = tp.Fill(&avs)
		if fill_err != nil {
			return
		}
		var salt [32]byte
		fill_err = tp.Fill(&salt)
		if fill_err != nil {
			return
		}
		var expiry *big.Int
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || expiry == nil {
			return
		}

		_AVSDirectory.CalculateOperatorAVSRegistrationDigestHash(operator, avs, salt, expiry)
	})
}

func Fuzz_Nosy_AVSDirectoryCallerSession_DOMAINTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryCallerSession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.DOMAINTYPEHASH()
	})
}

func Fuzz_Nosy_AVSDirectoryCallerSession_Delegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryCallerSession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.Delegation()
	})
}

func Fuzz_Nosy_AVSDirectoryCallerSession_DomainSeparator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryCallerSession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.DomainSeparator()
	})
}

func Fuzz_Nosy_AVSDirectoryCallerSession_OPERATORAVSREGISTRATIONTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryCallerSession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.OPERATORAVSREGISTRATIONTYPEHASH()
	})
}

func Fuzz_Nosy_AVSDirectoryCallerSession_OperatorSaltIsSpent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryCallerSession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 [32]byte
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.OperatorSaltIsSpent(arg0, arg1)
	})
}

func Fuzz_Nosy_AVSDirectoryCallerSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryCallerSession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.Owner()
	})
}

func Fuzz_Nosy_AVSDirectoryCallerSession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryCallerSession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var index uint8
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.Paused(index)
	})
}

func Fuzz_Nosy_AVSDirectoryCallerSession_Paused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryCallerSession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.Paused0()
	})
}

func Fuzz_Nosy_AVSDirectoryCallerSession_PauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryCallerSession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.PauserRegistry()
	})
}

func Fuzz_Nosy_AVSDirectoryFilterer_FilterAVSMetadataURIUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryFilterer
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var avs []common.Address
		fill_err = tp.Fill(&avs)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil {
			return
		}

		_AVSDirectory.FilterAVSMetadataURIUpdated(opts, avs)
	})
}

func Fuzz_Nosy_AVSDirectoryFilterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryFilterer
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil {
			return
		}

		_AVSDirectory.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_AVSDirectoryFilterer_FilterOperatorAVSRegistrationStatusUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryFilterer
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator []common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var avs []common.Address
		fill_err = tp.Fill(&avs)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil {
			return
		}

		_AVSDirectory.FilterOperatorAVSRegistrationStatusUpdated(opts, operator, avs)
	})
}

func Fuzz_Nosy_AVSDirectoryFilterer_FilterOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryFilterer
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var previousOwner []common.Address
		fill_err = tp.Fill(&previousOwner)
		if fill_err != nil {
			return
		}
		var newOwner []common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil {
			return
		}

		_AVSDirectory.FilterOwnershipTransferred(opts, previousOwner, newOwner)
	})
}

func Fuzz_Nosy_AVSDirectoryFilterer_FilterPaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryFilterer
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var account []common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil {
			return
		}

		_AVSDirectory.FilterPaused(opts, account)
	})
}

func Fuzz_Nosy_AVSDirectoryFilterer_FilterPauserRegistrySet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryFilterer
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil {
			return
		}

		_AVSDirectory.FilterPauserRegistrySet(opts)
	})
}

func Fuzz_Nosy_AVSDirectoryFilterer_FilterUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryFilterer
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var account []common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil {
			return
		}

		_AVSDirectory.FilterUnpaused(opts, account)
	})
}

func Fuzz_Nosy_AVSDirectoryFilterer_ParseAVSMetadataURIUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryFilterer
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.ParseAVSMetadataURIUpdated(log)
	})
}

func Fuzz_Nosy_AVSDirectoryFilterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryFilterer
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.ParseInitialized(log)
	})
}

func Fuzz_Nosy_AVSDirectoryFilterer_ParseOperatorAVSRegistrationStatusUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryFilterer
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.ParseOperatorAVSRegistrationStatusUpdated(log)
	})
}

func Fuzz_Nosy_AVSDirectoryFilterer_ParseOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryFilterer
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.ParseOwnershipTransferred(log)
	})
}

func Fuzz_Nosy_AVSDirectoryFilterer_ParsePaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryFilterer
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.ParsePaused(log)
	})
}

func Fuzz_Nosy_AVSDirectoryFilterer_ParsePauserRegistrySet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryFilterer
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.ParsePauserRegistrySet(log)
	})
}

func Fuzz_Nosy_AVSDirectoryFilterer_ParseUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryFilterer
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.ParseUnpaused(log)
	})
}

// skipping Fuzz_Nosy_AVSDirectoryFilterer_WatchAVSMetadataURIUpdated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.AVSDirectoryAVSMetadataURIUpdated

// skipping Fuzz_Nosy_AVSDirectoryFilterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.AVSDirectoryInitialized

// skipping Fuzz_Nosy_AVSDirectoryFilterer_WatchOperatorAVSRegistrationStatusUpdated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.AVSDirectoryOperatorAVSRegistrationStatusUpdated

// skipping Fuzz_Nosy_AVSDirectoryFilterer_WatchOwnershipTransferred__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.AVSDirectoryOwnershipTransferred

// skipping Fuzz_Nosy_AVSDirectoryFilterer_WatchPaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.AVSDirectoryPaused

// skipping Fuzz_Nosy_AVSDirectoryFilterer_WatchPauserRegistrySet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.AVSDirectoryPauserRegistrySet

// skipping Fuzz_Nosy_AVSDirectoryFilterer_WatchUnpaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.AVSDirectoryUnpaused

func Fuzz_Nosy_AVSDirectoryInitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *AVSDirectoryInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_AVSDirectoryInitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *AVSDirectoryInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_AVSDirectoryInitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *AVSDirectoryInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_AVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *AVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_AVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *AVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_AVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *AVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_AVSDirectoryOwnershipTransferredIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *AVSDirectoryOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_AVSDirectoryOwnershipTransferredIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *AVSDirectoryOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_AVSDirectoryOwnershipTransferredIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *AVSDirectoryOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_AVSDirectoryPausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *AVSDirectoryPausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_AVSDirectoryPausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *AVSDirectoryPausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_AVSDirectoryPausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *AVSDirectoryPausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_AVSDirectoryPauserRegistrySetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *AVSDirectoryPauserRegistrySetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_AVSDirectoryPauserRegistrySetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *AVSDirectoryPauserRegistrySetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_AVSDirectoryPauserRegistrySetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *AVSDirectoryPauserRegistrySetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

// skipping Fuzz_Nosy_AVSDirectoryRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_AVSDirectoryRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_AVSDirectoryRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryRaw
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil {
			return
		}

		_AVSDirectory.Transfer(opts)
	})
}

func Fuzz_Nosy_AVSDirectorySession_AvsOperatorStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectorySession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 common.Address
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.AvsOperatorStatus(arg0, arg1)
	})
}

func Fuzz_Nosy_AVSDirectorySession_CalculateOperatorAVSRegistrationDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectorySession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var avs common.Address
		fill_err = tp.Fill(&avs)
		if fill_err != nil {
			return
		}
		var salt [32]byte
		fill_err = tp.Fill(&salt)
		if fill_err != nil {
			return
		}
		var expiry *big.Int
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || expiry == nil {
			return
		}

		_AVSDirectory.CalculateOperatorAVSRegistrationDigestHash(operator, avs, salt, expiry)
	})
}

func Fuzz_Nosy_AVSDirectorySession_CancelSalt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectorySession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var salt [32]byte
		fill_err = tp.Fill(&salt)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.CancelSalt(salt)
	})
}

func Fuzz_Nosy_AVSDirectorySession_DOMAINTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectorySession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.DOMAINTYPEHASH()
	})
}

func Fuzz_Nosy_AVSDirectorySession_Delegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectorySession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.Delegation()
	})
}

func Fuzz_Nosy_AVSDirectorySession_DeregisterOperatorFromAVS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectorySession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.DeregisterOperatorFromAVS(operator)
	})
}

func Fuzz_Nosy_AVSDirectorySession_DomainSeparator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectorySession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.DomainSeparator()
	})
}

func Fuzz_Nosy_AVSDirectorySession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectorySession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var initialOwner common.Address
		fill_err = tp.Fill(&initialOwner)
		if fill_err != nil {
			return
		}
		var _pauserRegistry common.Address
		fill_err = tp.Fill(&_pauserRegistry)
		if fill_err != nil {
			return
		}
		var initialPausedStatus *big.Int
		fill_err = tp.Fill(&initialPausedStatus)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || initialPausedStatus == nil {
			return
		}

		_AVSDirectory.Initialize(initialOwner, _pauserRegistry, initialPausedStatus)
	})
}

func Fuzz_Nosy_AVSDirectorySession_OPERATORAVSREGISTRATIONTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectorySession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.OPERATORAVSREGISTRATIONTYPEHASH()
	})
}

func Fuzz_Nosy_AVSDirectorySession_OperatorSaltIsSpent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectorySession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 [32]byte
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.OperatorSaltIsSpent(arg0, arg1)
	})
}

func Fuzz_Nosy_AVSDirectorySession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectorySession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.Owner()
	})
}

func Fuzz_Nosy_AVSDirectorySession_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectorySession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || newPausedStatus == nil {
			return
		}

		_AVSDirectory.Pause(newPausedStatus)
	})
}

func Fuzz_Nosy_AVSDirectorySession_PauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectorySession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.PauseAll()
	})
}

func Fuzz_Nosy_AVSDirectorySession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectorySession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var index uint8
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.Paused(index)
	})
}

func Fuzz_Nosy_AVSDirectorySession_Paused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectorySession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.Paused0()
	})
}

func Fuzz_Nosy_AVSDirectorySession_PauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectorySession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.PauserRegistry()
	})
}

func Fuzz_Nosy_AVSDirectorySession_RegisterOperatorToAVS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectorySession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry
		fill_err = tp.Fill(&operatorSignature)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.RegisterOperatorToAVS(operator, operatorSignature)
	})
}

func Fuzz_Nosy_AVSDirectorySession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectorySession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.RenounceOwnership()
	})
}

func Fuzz_Nosy_AVSDirectorySession_SetPauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectorySession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var newPauserRegistry common.Address
		fill_err = tp.Fill(&newPauserRegistry)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.SetPauserRegistry(newPauserRegistry)
	})
}

func Fuzz_Nosy_AVSDirectorySession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectorySession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_AVSDirectorySession_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectorySession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || newPausedStatus == nil {
			return
		}

		_AVSDirectory.Unpause(newPausedStatus)
	})
}

func Fuzz_Nosy_AVSDirectorySession_UpdateAVSMetadataURI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectorySession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var metadataURI string
		fill_err = tp.Fill(&metadataURI)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.UpdateAVSMetadataURI(metadataURI)
	})
}

func Fuzz_Nosy_AVSDirectoryTransactor_CancelSalt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryTransactor
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var salt [32]byte
		fill_err = tp.Fill(&salt)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil {
			return
		}

		_AVSDirectory.CancelSalt(opts, salt)
	})
}

func Fuzz_Nosy_AVSDirectoryTransactor_DeregisterOperatorFromAVS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryTransactor
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil {
			return
		}

		_AVSDirectory.DeregisterOperatorFromAVS(opts, operator)
	})
}

func Fuzz_Nosy_AVSDirectoryTransactor_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryTransactor
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var initialOwner common.Address
		fill_err = tp.Fill(&initialOwner)
		if fill_err != nil {
			return
		}
		var _pauserRegistry common.Address
		fill_err = tp.Fill(&_pauserRegistry)
		if fill_err != nil {
			return
		}
		var initialPausedStatus *big.Int
		fill_err = tp.Fill(&initialPausedStatus)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil || initialPausedStatus == nil {
			return
		}

		_AVSDirectory.Initialize(opts, initialOwner, _pauserRegistry, initialPausedStatus)
	})
}

func Fuzz_Nosy_AVSDirectoryTransactor_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryTransactor
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil || newPausedStatus == nil {
			return
		}

		_AVSDirectory.Pause(opts, newPausedStatus)
	})
}

func Fuzz_Nosy_AVSDirectoryTransactor_PauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryTransactor
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil {
			return
		}

		_AVSDirectory.PauseAll(opts)
	})
}

func Fuzz_Nosy_AVSDirectoryTransactor_RegisterOperatorToAVS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryTransactor
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry
		fill_err = tp.Fill(&operatorSignature)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil {
			return
		}

		_AVSDirectory.RegisterOperatorToAVS(opts, operator, operatorSignature)
	})
}

func Fuzz_Nosy_AVSDirectoryTransactor_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryTransactor
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil {
			return
		}

		_AVSDirectory.RenounceOwnership(opts)
	})
}

func Fuzz_Nosy_AVSDirectoryTransactor_SetPauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryTransactor
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newPauserRegistry common.Address
		fill_err = tp.Fill(&newPauserRegistry)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil {
			return
		}

		_AVSDirectory.SetPauserRegistry(opts, newPauserRegistry)
	})
}

func Fuzz_Nosy_AVSDirectoryTransactor_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryTransactor
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil {
			return
		}

		_AVSDirectory.TransferOwnership(opts, newOwner)
	})
}

func Fuzz_Nosy_AVSDirectoryTransactor_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryTransactor
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil || newPausedStatus == nil {
			return
		}

		_AVSDirectory.Unpause(opts, newPausedStatus)
	})
}

func Fuzz_Nosy_AVSDirectoryTransactor_UpdateAVSMetadataURI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryTransactor
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var metadataURI string
		fill_err = tp.Fill(&metadataURI)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil {
			return
		}

		_AVSDirectory.UpdateAVSMetadataURI(opts, metadataURI)
	})
}

// skipping Fuzz_Nosy_AVSDirectoryTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_AVSDirectoryTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryTransactorRaw
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || opts == nil {
			return
		}

		_AVSDirectory.Transfer(opts)
	})
}

func Fuzz_Nosy_AVSDirectoryTransactorSession_CancelSalt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryTransactorSession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var salt [32]byte
		fill_err = tp.Fill(&salt)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.CancelSalt(salt)
	})
}

func Fuzz_Nosy_AVSDirectoryTransactorSession_DeregisterOperatorFromAVS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryTransactorSession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.DeregisterOperatorFromAVS(operator)
	})
}

func Fuzz_Nosy_AVSDirectoryTransactorSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryTransactorSession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var initialOwner common.Address
		fill_err = tp.Fill(&initialOwner)
		if fill_err != nil {
			return
		}
		var _pauserRegistry common.Address
		fill_err = tp.Fill(&_pauserRegistry)
		if fill_err != nil {
			return
		}
		var initialPausedStatus *big.Int
		fill_err = tp.Fill(&initialPausedStatus)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || initialPausedStatus == nil {
			return
		}

		_AVSDirectory.Initialize(initialOwner, _pauserRegistry, initialPausedStatus)
	})
}

func Fuzz_Nosy_AVSDirectoryTransactorSession_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryTransactorSession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || newPausedStatus == nil {
			return
		}

		_AVSDirectory.Pause(newPausedStatus)
	})
}

func Fuzz_Nosy_AVSDirectoryTransactorSession_PauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryTransactorSession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.PauseAll()
	})
}

func Fuzz_Nosy_AVSDirectoryTransactorSession_RegisterOperatorToAVS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryTransactorSession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry
		fill_err = tp.Fill(&operatorSignature)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.RegisterOperatorToAVS(operator, operatorSignature)
	})
}

func Fuzz_Nosy_AVSDirectoryTransactorSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryTransactorSession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.RenounceOwnership()
	})
}

func Fuzz_Nosy_AVSDirectoryTransactorSession_SetPauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryTransactorSession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var newPauserRegistry common.Address
		fill_err = tp.Fill(&newPauserRegistry)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.SetPauserRegistry(newPauserRegistry)
	})
}

func Fuzz_Nosy_AVSDirectoryTransactorSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryTransactorSession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_AVSDirectoryTransactorSession_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryTransactorSession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil || newPausedStatus == nil {
			return
		}

		_AVSDirectory.Unpause(newPausedStatus)
	})
}

func Fuzz_Nosy_AVSDirectoryTransactorSession_UpdateAVSMetadataURI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AVSDirectory *AVSDirectoryTransactorSession
		fill_err = tp.Fill(&_AVSDirectory)
		if fill_err != nil {
			return
		}
		var metadataURI string
		fill_err = tp.Fill(&metadataURI)
		if fill_err != nil {
			return
		}
		if _AVSDirectory == nil {
			return
		}

		_AVSDirectory.UpdateAVSMetadataURI(metadataURI)
	})
}

func Fuzz_Nosy_AVSDirectoryUnpausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *AVSDirectoryUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_AVSDirectoryUnpausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *AVSDirectoryUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_AVSDirectoryUnpausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *AVSDirectoryUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_AdminCaller_ISSCRIPT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Admin *AdminCaller
		fill_err = tp.Fill(&_Admin)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Admin == nil || opts == nil {
			return
		}

		_Admin.ISSCRIPT(opts)
	})
}

// skipping Fuzz_Nosy_AdminCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_AdminCallerSession_ISSCRIPT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Admin *AdminCallerSession
		fill_err = tp.Fill(&_Admin)
		if fill_err != nil {
			return
		}
		if _Admin == nil {
			return
		}

		_Admin.ISSCRIPT()
	})
}

// skipping Fuzz_Nosy_AdminRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_AdminRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_AdminRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Admin *AdminRaw
		fill_err = tp.Fill(&_Admin)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Admin == nil || opts == nil {
			return
		}

		_Admin.Transfer(opts)
	})
}

func Fuzz_Nosy_AdminSession_ISSCRIPT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Admin *AdminSession
		fill_err = tp.Fill(&_Admin)
		if fill_err != nil {
			return
		}
		if _Admin == nil {
			return
		}

		_Admin.ISSCRIPT()
	})
}

func Fuzz_Nosy_AdminSession_PausePortal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Admin *AdminSession
		fill_err = tp.Fill(&_Admin)
		if fill_err != nil {
			return
		}
		var admin common.Address
		fill_err = tp.Fill(&admin)
		if fill_err != nil {
			return
		}
		var portal common.Address
		fill_err = tp.Fill(&portal)
		if fill_err != nil {
			return
		}
		if _Admin == nil {
			return
		}

		_Admin.PausePortal(admin, portal)
	})
}

func Fuzz_Nosy_AdminSession_UnpausePortal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Admin *AdminSession
		fill_err = tp.Fill(&_Admin)
		if fill_err != nil {
			return
		}
		var admin common.Address
		fill_err = tp.Fill(&admin)
		if fill_err != nil {
			return
		}
		var portal common.Address
		fill_err = tp.Fill(&portal)
		if fill_err != nil {
			return
		}
		if _Admin == nil {
			return
		}

		_Admin.UnpausePortal(admin, portal)
	})
}

func Fuzz_Nosy_AdminSession_UpgradePortal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Admin *AdminSession
		fill_err = tp.Fill(&_Admin)
		if fill_err != nil {
			return
		}
		var admin common.Address
		fill_err = tp.Fill(&admin)
		if fill_err != nil {
			return
		}
		var deployer common.Address
		fill_err = tp.Fill(&deployer)
		if fill_err != nil {
			return
		}
		var portal common.Address
		fill_err = tp.Fill(&portal)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if _Admin == nil {
			return
		}

		_Admin.UpgradePortal(admin, deployer, portal, d5)
	})
}

func Fuzz_Nosy_AdminTransactor_PausePortal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Admin *AdminTransactor
		fill_err = tp.Fill(&_Admin)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var admin common.Address
		fill_err = tp.Fill(&admin)
		if fill_err != nil {
			return
		}
		var portal common.Address
		fill_err = tp.Fill(&portal)
		if fill_err != nil {
			return
		}
		if _Admin == nil || opts == nil {
			return
		}

		_Admin.PausePortal(opts, admin, portal)
	})
}

func Fuzz_Nosy_AdminTransactor_UnpausePortal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Admin *AdminTransactor
		fill_err = tp.Fill(&_Admin)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var admin common.Address
		fill_err = tp.Fill(&admin)
		if fill_err != nil {
			return
		}
		var portal common.Address
		fill_err = tp.Fill(&portal)
		if fill_err != nil {
			return
		}
		if _Admin == nil || opts == nil {
			return
		}

		_Admin.UnpausePortal(opts, admin, portal)
	})
}

func Fuzz_Nosy_AdminTransactor_UpgradePortal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Admin *AdminTransactor
		fill_err = tp.Fill(&_Admin)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var admin common.Address
		fill_err = tp.Fill(&admin)
		if fill_err != nil {
			return
		}
		var deployer common.Address
		fill_err = tp.Fill(&deployer)
		if fill_err != nil {
			return
		}
		var portal common.Address
		fill_err = tp.Fill(&portal)
		if fill_err != nil {
			return
		}
		var d6 []byte
		fill_err = tp.Fill(&d6)
		if fill_err != nil {
			return
		}
		if _Admin == nil || opts == nil {
			return
		}

		_Admin.UpgradePortal(opts, admin, deployer, portal, d6)
	})
}

// skipping Fuzz_Nosy_AdminTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_AdminTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Admin *AdminTransactorRaw
		fill_err = tp.Fill(&_Admin)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Admin == nil || opts == nil {
			return
		}

		_Admin.Transfer(opts)
	})
}

func Fuzz_Nosy_AdminTransactorSession_PausePortal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Admin *AdminTransactorSession
		fill_err = tp.Fill(&_Admin)
		if fill_err != nil {
			return
		}
		var admin common.Address
		fill_err = tp.Fill(&admin)
		if fill_err != nil {
			return
		}
		var portal common.Address
		fill_err = tp.Fill(&portal)
		if fill_err != nil {
			return
		}
		if _Admin == nil {
			return
		}

		_Admin.PausePortal(admin, portal)
	})
}

func Fuzz_Nosy_AdminTransactorSession_UnpausePortal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Admin *AdminTransactorSession
		fill_err = tp.Fill(&_Admin)
		if fill_err != nil {
			return
		}
		var admin common.Address
		fill_err = tp.Fill(&admin)
		if fill_err != nil {
			return
		}
		var portal common.Address
		fill_err = tp.Fill(&portal)
		if fill_err != nil {
			return
		}
		if _Admin == nil {
			return
		}

		_Admin.UnpausePortal(admin, portal)
	})
}

func Fuzz_Nosy_AdminTransactorSession_UpgradePortal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Admin *AdminTransactorSession
		fill_err = tp.Fill(&_Admin)
		if fill_err != nil {
			return
		}
		var admin common.Address
		fill_err = tp.Fill(&admin)
		if fill_err != nil {
			return
		}
		var deployer common.Address
		fill_err = tp.Fill(&deployer)
		if fill_err != nil {
			return
		}
		var portal common.Address
		fill_err = tp.Fill(&portal)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if _Admin == nil {
			return
		}

		_Admin.UpgradePortal(admin, deployer, portal, d5)
	})
}

func Fuzz_Nosy_AllocPredeploysCaller_ISSCRIPT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AllocPredeploys *AllocPredeploysCaller
		fill_err = tp.Fill(&_AllocPredeploys)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _AllocPredeploys == nil || opts == nil {
			return
		}

		_AllocPredeploys.ISSCRIPT(opts)
	})
}

// skipping Fuzz_Nosy_AllocPredeploysCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_AllocPredeploysCallerSession_ISSCRIPT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AllocPredeploys *AllocPredeploysCallerSession
		fill_err = tp.Fill(&_AllocPredeploys)
		if fill_err != nil {
			return
		}
		if _AllocPredeploys == nil {
			return
		}

		_AllocPredeploys.ISSCRIPT()
	})
}

// skipping Fuzz_Nosy_AllocPredeploysRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_AllocPredeploysRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_AllocPredeploysRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AllocPredeploys *AllocPredeploysRaw
		fill_err = tp.Fill(&_AllocPredeploys)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _AllocPredeploys == nil || opts == nil {
			return
		}

		_AllocPredeploys.Transfer(opts)
	})
}

func Fuzz_Nosy_AllocPredeploysSession_ISSCRIPT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AllocPredeploys *AllocPredeploysSession
		fill_err = tp.Fill(&_AllocPredeploys)
		if fill_err != nil {
			return
		}
		if _AllocPredeploys == nil {
			return
		}

		_AllocPredeploys.ISSCRIPT()
	})
}

func Fuzz_Nosy_AllocPredeploysSession_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AllocPredeploys *AllocPredeploysSession
		fill_err = tp.Fill(&_AllocPredeploys)
		if fill_err != nil {
			return
		}
		var config AllocPredeploysConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if _AllocPredeploys == nil {
			return
		}

		_AllocPredeploys.Run(config)
	})
}

func Fuzz_Nosy_AllocPredeploysSession_RunNoStateDump__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AllocPredeploys *AllocPredeploysSession
		fill_err = tp.Fill(&_AllocPredeploys)
		if fill_err != nil {
			return
		}
		var config AllocPredeploysConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if _AllocPredeploys == nil {
			return
		}

		_AllocPredeploys.RunNoStateDump(config)
	})
}

func Fuzz_Nosy_AllocPredeploysTransactor_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AllocPredeploys *AllocPredeploysTransactor
		fill_err = tp.Fill(&_AllocPredeploys)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var config AllocPredeploysConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if _AllocPredeploys == nil || opts == nil {
			return
		}

		_AllocPredeploys.Run(opts, config)
	})
}

func Fuzz_Nosy_AllocPredeploysTransactor_RunNoStateDump__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AllocPredeploys *AllocPredeploysTransactor
		fill_err = tp.Fill(&_AllocPredeploys)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var config AllocPredeploysConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if _AllocPredeploys == nil || opts == nil {
			return
		}

		_AllocPredeploys.RunNoStateDump(opts, config)
	})
}

// skipping Fuzz_Nosy_AllocPredeploysTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_AllocPredeploysTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AllocPredeploys *AllocPredeploysTransactorRaw
		fill_err = tp.Fill(&_AllocPredeploys)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _AllocPredeploys == nil || opts == nil {
			return
		}

		_AllocPredeploys.Transfer(opts)
	})
}

func Fuzz_Nosy_AllocPredeploysTransactorSession_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AllocPredeploys *AllocPredeploysTransactorSession
		fill_err = tp.Fill(&_AllocPredeploys)
		if fill_err != nil {
			return
		}
		var config AllocPredeploysConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if _AllocPredeploys == nil {
			return
		}

		_AllocPredeploys.Run(config)
	})
}

func Fuzz_Nosy_AllocPredeploysTransactorSession_RunNoStateDump__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _AllocPredeploys *AllocPredeploysTransactorSession
		fill_err = tp.Fill(&_AllocPredeploys)
		if fill_err != nil {
			return
		}
		var config AllocPredeploysConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if _AllocPredeploys == nil {
			return
		}

		_AllocPredeploys.RunNoStateDump(config)
	})
}

func Fuzz_Nosy_Create3Caller_GetDeployed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Create3 *Create3Caller
		fill_err = tp.Fill(&_Create3)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var deployer common.Address
		fill_err = tp.Fill(&deployer)
		if fill_err != nil {
			return
		}
		var salt [32]byte
		fill_err = tp.Fill(&salt)
		if fill_err != nil {
			return
		}
		if _Create3 == nil || opts == nil {
			return
		}

		_Create3.GetDeployed(opts, deployer, salt)
	})
}

// skipping Fuzz_Nosy_Create3CallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_Create3CallerSession_GetDeployed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Create3 *Create3CallerSession
		fill_err = tp.Fill(&_Create3)
		if fill_err != nil {
			return
		}
		var deployer common.Address
		fill_err = tp.Fill(&deployer)
		if fill_err != nil {
			return
		}
		var salt [32]byte
		fill_err = tp.Fill(&salt)
		if fill_err != nil {
			return
		}
		if _Create3 == nil {
			return
		}

		_Create3.GetDeployed(deployer, salt)
	})
}

// skipping Fuzz_Nosy_Create3Raw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_Create3Raw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_Create3Raw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Create3 *Create3Raw
		fill_err = tp.Fill(&_Create3)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Create3 == nil || opts == nil {
			return
		}

		_Create3.Transfer(opts)
	})
}

func Fuzz_Nosy_Create3Session_Deploy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Create3 *Create3Session
		fill_err = tp.Fill(&_Create3)
		if fill_err != nil {
			return
		}
		var salt [32]byte
		fill_err = tp.Fill(&salt)
		if fill_err != nil {
			return
		}
		var creationCode []byte
		fill_err = tp.Fill(&creationCode)
		if fill_err != nil {
			return
		}
		if _Create3 == nil {
			return
		}

		_Create3.Deploy(salt, creationCode)
	})
}

func Fuzz_Nosy_Create3Session_GetDeployed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Create3 *Create3Session
		fill_err = tp.Fill(&_Create3)
		if fill_err != nil {
			return
		}
		var deployer common.Address
		fill_err = tp.Fill(&deployer)
		if fill_err != nil {
			return
		}
		var salt [32]byte
		fill_err = tp.Fill(&salt)
		if fill_err != nil {
			return
		}
		if _Create3 == nil {
			return
		}

		_Create3.GetDeployed(deployer, salt)
	})
}

func Fuzz_Nosy_Create3Transactor_Deploy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Create3 *Create3Transactor
		fill_err = tp.Fill(&_Create3)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var salt [32]byte
		fill_err = tp.Fill(&salt)
		if fill_err != nil {
			return
		}
		var creationCode []byte
		fill_err = tp.Fill(&creationCode)
		if fill_err != nil {
			return
		}
		if _Create3 == nil || opts == nil {
			return
		}

		_Create3.Deploy(opts, salt, creationCode)
	})
}

// skipping Fuzz_Nosy_Create3TransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_Create3TransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Create3 *Create3TransactorRaw
		fill_err = tp.Fill(&_Create3)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Create3 == nil || opts == nil {
			return
		}

		_Create3.Transfer(opts)
	})
}

func Fuzz_Nosy_Create3TransactorSession_Deploy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Create3 *Create3TransactorSession
		fill_err = tp.Fill(&_Create3)
		if fill_err != nil {
			return
		}
		var salt [32]byte
		fill_err = tp.Fill(&salt)
		if fill_err != nil {
			return
		}
		var creationCode []byte
		fill_err = tp.Fill(&creationCode)
		if fill_err != nil {
			return
		}
		if _Create3 == nil {
			return
		}

		_Create3.Deploy(salt, creationCode)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_BeaconChainETHStrategy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.BeaconChainETHStrategy(opts)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_CalculateCurrentStakerDelegationDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var expiry *big.Int
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil || expiry == nil {
			return
		}

		_DelegationManager.CalculateCurrentStakerDelegationDigestHash(opts, staker, operator, expiry)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_CalculateDelegationApprovalDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var _delegationApprover common.Address
		fill_err = tp.Fill(&_delegationApprover)
		if fill_err != nil {
			return
		}
		var approverSalt [32]byte
		fill_err = tp.Fill(&approverSalt)
		if fill_err != nil {
			return
		}
		var expiry *big.Int
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil || expiry == nil {
			return
		}

		_DelegationManager.CalculateDelegationApprovalDigestHash(opts, staker, operator, _delegationApprover, approverSalt, expiry)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_CalculateStakerDelegationDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var _stakerNonce *big.Int
		fill_err = tp.Fill(&_stakerNonce)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var expiry *big.Int
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil || _stakerNonce == nil || expiry == nil {
			return
		}

		_DelegationManager.CalculateStakerDelegationDigestHash(opts, staker, _stakerNonce, operator, expiry)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_CalculateWithdrawalRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var withdrawal IDelegationManagerWithdrawal
		fill_err = tp.Fill(&withdrawal)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.CalculateWithdrawalRoot(opts, withdrawal)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_CumulativeWithdrawalsQueued__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.CumulativeWithdrawalsQueued(opts, arg0)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_DELEGATIONAPPROVALTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.DELEGATIONAPPROVALTYPEHASH(opts)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_DOMAINTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.DOMAINTYPEHASH(opts)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_DelegatedTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.DelegatedTo(opts, arg0)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_DelegationApprover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.DelegationApprover(opts, operator)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_DelegationApproverSaltIsSpent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 [32]byte
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.DelegationApproverSaltIsSpent(opts, arg0, arg1)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_DomainSeparator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.DomainSeparator(opts)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_EarningsReceiver__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.EarningsReceiver(opts, operator)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_EigenPodManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.EigenPodManager(opts)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_GetDelegatableShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.GetDelegatableShares(opts, staker)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_GetOperatorShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var strategies []common.Address
		fill_err = tp.Fill(&strategies)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.GetOperatorShares(opts, operator, strategies)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_GetWithdrawalDelay__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var strategies []common.Address
		fill_err = tp.Fill(&strategies)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.GetWithdrawalDelay(opts, strategies)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_IsDelegated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.IsDelegated(opts, staker)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_IsOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.IsOperator(opts, operator)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_MAXSTAKEROPTOUTWINDOWBLOCKS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.MAXSTAKEROPTOUTWINDOWBLOCKS(opts)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_MAXWITHDRAWALDELAYBLOCKS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.MAXWITHDRAWALDELAYBLOCKS(opts)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_MinWithdrawalDelayBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.MinWithdrawalDelayBlocks(opts)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_OperatorDetails__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.OperatorDetails(opts, operator)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_OperatorShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 common.Address
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.OperatorShares(opts, arg0, arg1)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.Owner(opts)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var index uint8
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.Paused(opts, index)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_Paused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.Paused0(opts)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_PauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.PauserRegistry(opts)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_PendingWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.PendingWithdrawals(opts, arg0)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_STAKERDELEGATIONTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.STAKERDELEGATIONTYPEHASH(opts)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_Slasher__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.Slasher(opts)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_StakerNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.StakerNonce(opts, arg0)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_StakerOptOutWindowBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.StakerOptOutWindowBlocks(opts, operator)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_StrategyManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.StrategyManager(opts)
	})
}

func Fuzz_Nosy_DelegationManagerCaller_StrategyWithdrawalDelayBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCaller
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.StrategyWithdrawalDelayBlocks(opts, arg0)
	})
}

// skipping Fuzz_Nosy_DelegationManagerCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_DelegationManagerCallerSession_BeaconChainETHStrategy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.BeaconChainETHStrategy()
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_CalculateCurrentStakerDelegationDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var expiry *big.Int
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || expiry == nil {
			return
		}

		_DelegationManager.CalculateCurrentStakerDelegationDigestHash(staker, operator, expiry)
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_CalculateDelegationApprovalDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var _delegationApprover common.Address
		fill_err = tp.Fill(&_delegationApprover)
		if fill_err != nil {
			return
		}
		var approverSalt [32]byte
		fill_err = tp.Fill(&approverSalt)
		if fill_err != nil {
			return
		}
		var expiry *big.Int
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || expiry == nil {
			return
		}

		_DelegationManager.CalculateDelegationApprovalDigestHash(staker, operator, _delegationApprover, approverSalt, expiry)
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_CalculateStakerDelegationDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var _stakerNonce *big.Int
		fill_err = tp.Fill(&_stakerNonce)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var expiry *big.Int
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || _stakerNonce == nil || expiry == nil {
			return
		}

		_DelegationManager.CalculateStakerDelegationDigestHash(staker, _stakerNonce, operator, expiry)
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_CalculateWithdrawalRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var withdrawal IDelegationManagerWithdrawal
		fill_err = tp.Fill(&withdrawal)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.CalculateWithdrawalRoot(withdrawal)
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_CumulativeWithdrawalsQueued__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.CumulativeWithdrawalsQueued(arg0)
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_DELEGATIONAPPROVALTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.DELEGATIONAPPROVALTYPEHASH()
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_DOMAINTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.DOMAINTYPEHASH()
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_DelegatedTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.DelegatedTo(arg0)
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_DelegationApprover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.DelegationApprover(operator)
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_DelegationApproverSaltIsSpent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 [32]byte
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.DelegationApproverSaltIsSpent(arg0, arg1)
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_DomainSeparator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.DomainSeparator()
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_EarningsReceiver__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.EarningsReceiver(operator)
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_EigenPodManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.EigenPodManager()
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_GetDelegatableShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.GetDelegatableShares(staker)
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_GetOperatorShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var strategies []common.Address
		fill_err = tp.Fill(&strategies)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.GetOperatorShares(operator, strategies)
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_GetWithdrawalDelay__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var strategies []common.Address
		fill_err = tp.Fill(&strategies)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.GetWithdrawalDelay(strategies)
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_IsDelegated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.IsDelegated(staker)
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_IsOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.IsOperator(operator)
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_MAXSTAKEROPTOUTWINDOWBLOCKS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.MAXSTAKEROPTOUTWINDOWBLOCKS()
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_MAXWITHDRAWALDELAYBLOCKS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.MAXWITHDRAWALDELAYBLOCKS()
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_MinWithdrawalDelayBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.MinWithdrawalDelayBlocks()
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_OperatorDetails__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.OperatorDetails(operator)
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_OperatorShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 common.Address
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.OperatorShares(arg0, arg1)
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.Owner()
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var index uint8
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.Paused(index)
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_Paused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.Paused0()
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_PauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.PauserRegistry()
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_PendingWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.PendingWithdrawals(arg0)
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_STAKERDELEGATIONTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.STAKERDELEGATIONTYPEHASH()
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_Slasher__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.Slasher()
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_StakerNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.StakerNonce(arg0)
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_StakerOptOutWindowBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.StakerOptOutWindowBlocks(operator)
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_StrategyManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.StrategyManager()
	})
}

func Fuzz_Nosy_DelegationManagerCallerSession_StrategyWithdrawalDelayBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerCallerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.StrategyWithdrawalDelayBlocks(arg0)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_FilterMinWithdrawalDelayBlocksSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.FilterMinWithdrawalDelayBlocksSet(opts)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_FilterOperatorDetailsModified__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator []common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.FilterOperatorDetailsModified(opts, operator)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_FilterOperatorMetadataURIUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator []common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.FilterOperatorMetadataURIUpdated(opts, operator)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_FilterOperatorRegistered__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator []common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.FilterOperatorRegistered(opts, operator)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_FilterOperatorSharesDecreased__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator []common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.FilterOperatorSharesDecreased(opts, operator)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_FilterOperatorSharesIncreased__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator []common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.FilterOperatorSharesIncreased(opts, operator)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_FilterOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var previousOwner []common.Address
		fill_err = tp.Fill(&previousOwner)
		if fill_err != nil {
			return
		}
		var newOwner []common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.FilterOwnershipTransferred(opts, previousOwner, newOwner)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_FilterPaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var account []common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.FilterPaused(opts, account)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_FilterPauserRegistrySet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.FilterPauserRegistrySet(opts)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_FilterStakerDelegated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker []common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var operator []common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.FilterStakerDelegated(opts, staker, operator)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_FilterStakerForceUndelegated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker []common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var operator []common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.FilterStakerForceUndelegated(opts, staker, operator)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_FilterStakerUndelegated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker []common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var operator []common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.FilterStakerUndelegated(opts, staker, operator)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_FilterStrategyWithdrawalDelayBlocksSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.FilterStrategyWithdrawalDelayBlocksSet(opts)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_FilterUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var account []common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.FilterUnpaused(opts, account)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_FilterWithdrawalCompleted__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.FilterWithdrawalCompleted(opts)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_FilterWithdrawalMigrated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.FilterWithdrawalMigrated(opts)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_FilterWithdrawalQueued__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.FilterWithdrawalQueued(opts)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.ParseInitialized(log)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_ParseMinWithdrawalDelayBlocksSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.ParseMinWithdrawalDelayBlocksSet(log)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_ParseOperatorDetailsModified__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.ParseOperatorDetailsModified(log)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_ParseOperatorMetadataURIUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.ParseOperatorMetadataURIUpdated(log)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_ParseOperatorRegistered__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.ParseOperatorRegistered(log)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_ParseOperatorSharesDecreased__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.ParseOperatorSharesDecreased(log)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_ParseOperatorSharesIncreased__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.ParseOperatorSharesIncreased(log)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_ParseOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.ParseOwnershipTransferred(log)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_ParsePaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.ParsePaused(log)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_ParsePauserRegistrySet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.ParsePauserRegistrySet(log)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_ParseStakerDelegated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.ParseStakerDelegated(log)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_ParseStakerForceUndelegated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.ParseStakerForceUndelegated(log)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_ParseStakerUndelegated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.ParseStakerUndelegated(log)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_ParseStrategyWithdrawalDelayBlocksSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.ParseStrategyWithdrawalDelayBlocksSet(log)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_ParseUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.ParseUnpaused(log)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_ParseWithdrawalCompleted__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.ParseWithdrawalCompleted(log)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_ParseWithdrawalMigrated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.ParseWithdrawalMigrated(log)
	})
}

func Fuzz_Nosy_DelegationManagerFilterer_ParseWithdrawalQueued__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerFilterer
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.ParseWithdrawalQueued(log)
	})
}

// skipping Fuzz_Nosy_DelegationManagerFilterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.DelegationManagerInitialized

// skipping Fuzz_Nosy_DelegationManagerFilterer_WatchMinWithdrawalDelayBlocksSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.DelegationManagerMinWithdrawalDelayBlocksSet

// skipping Fuzz_Nosy_DelegationManagerFilterer_WatchOperatorDetailsModified__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.DelegationManagerOperatorDetailsModified

// skipping Fuzz_Nosy_DelegationManagerFilterer_WatchOperatorMetadataURIUpdated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.DelegationManagerOperatorMetadataURIUpdated

// skipping Fuzz_Nosy_DelegationManagerFilterer_WatchOperatorRegistered__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.DelegationManagerOperatorRegistered

// skipping Fuzz_Nosy_DelegationManagerFilterer_WatchOperatorSharesDecreased__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.DelegationManagerOperatorSharesDecreased

// skipping Fuzz_Nosy_DelegationManagerFilterer_WatchOperatorSharesIncreased__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.DelegationManagerOperatorSharesIncreased

// skipping Fuzz_Nosy_DelegationManagerFilterer_WatchOwnershipTransferred__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.DelegationManagerOwnershipTransferred

// skipping Fuzz_Nosy_DelegationManagerFilterer_WatchPaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.DelegationManagerPaused

// skipping Fuzz_Nosy_DelegationManagerFilterer_WatchPauserRegistrySet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.DelegationManagerPauserRegistrySet

// skipping Fuzz_Nosy_DelegationManagerFilterer_WatchStakerDelegated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.DelegationManagerStakerDelegated

// skipping Fuzz_Nosy_DelegationManagerFilterer_WatchStakerForceUndelegated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.DelegationManagerStakerForceUndelegated

// skipping Fuzz_Nosy_DelegationManagerFilterer_WatchStakerUndelegated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.DelegationManagerStakerUndelegated

// skipping Fuzz_Nosy_DelegationManagerFilterer_WatchStrategyWithdrawalDelayBlocksSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.DelegationManagerStrategyWithdrawalDelayBlocksSet

// skipping Fuzz_Nosy_DelegationManagerFilterer_WatchUnpaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.DelegationManagerUnpaused

// skipping Fuzz_Nosy_DelegationManagerFilterer_WatchWithdrawalCompleted__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.DelegationManagerWithdrawalCompleted

// skipping Fuzz_Nosy_DelegationManagerFilterer_WatchWithdrawalMigrated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.DelegationManagerWithdrawalMigrated

// skipping Fuzz_Nosy_DelegationManagerFilterer_WatchWithdrawalQueued__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.DelegationManagerWithdrawalQueued

func Fuzz_Nosy_DelegationManagerInitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_DelegationManagerInitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_DelegationManagerInitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_DelegationManagerMinWithdrawalDelayBlocksSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerMinWithdrawalDelayBlocksSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_DelegationManagerMinWithdrawalDelayBlocksSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerMinWithdrawalDelayBlocksSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_DelegationManagerMinWithdrawalDelayBlocksSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerMinWithdrawalDelayBlocksSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_DelegationManagerOperatorDetailsModifiedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerOperatorDetailsModifiedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_DelegationManagerOperatorDetailsModifiedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerOperatorDetailsModifiedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_DelegationManagerOperatorDetailsModifiedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerOperatorDetailsModifiedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_DelegationManagerOperatorMetadataURIUpdatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerOperatorMetadataURIUpdatedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_DelegationManagerOperatorMetadataURIUpdatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerOperatorMetadataURIUpdatedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_DelegationManagerOperatorMetadataURIUpdatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerOperatorMetadataURIUpdatedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_DelegationManagerOperatorRegisteredIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerOperatorRegisteredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_DelegationManagerOperatorRegisteredIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerOperatorRegisteredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_DelegationManagerOperatorRegisteredIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerOperatorRegisteredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_DelegationManagerOperatorSharesDecreasedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerOperatorSharesDecreasedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_DelegationManagerOperatorSharesDecreasedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerOperatorSharesDecreasedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_DelegationManagerOperatorSharesDecreasedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerOperatorSharesDecreasedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_DelegationManagerOperatorSharesIncreasedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerOperatorSharesIncreasedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_DelegationManagerOperatorSharesIncreasedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerOperatorSharesIncreasedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_DelegationManagerOperatorSharesIncreasedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerOperatorSharesIncreasedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_DelegationManagerOwnershipTransferredIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_DelegationManagerOwnershipTransferredIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_DelegationManagerOwnershipTransferredIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_DelegationManagerPausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerPausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_DelegationManagerPausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerPausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_DelegationManagerPausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerPausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_DelegationManagerPauserRegistrySetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerPauserRegistrySetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_DelegationManagerPauserRegistrySetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerPauserRegistrySetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_DelegationManagerPauserRegistrySetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerPauserRegistrySetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

// skipping Fuzz_Nosy_DelegationManagerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_DelegationManagerRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_DelegationManagerRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerRaw
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.Transfer(opts)
	})
}

func Fuzz_Nosy_DelegationManagerSession_BeaconChainETHStrategy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.BeaconChainETHStrategy()
	})
}

func Fuzz_Nosy_DelegationManagerSession_CalculateCurrentStakerDelegationDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var expiry *big.Int
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || expiry == nil {
			return
		}

		_DelegationManager.CalculateCurrentStakerDelegationDigestHash(staker, operator, expiry)
	})
}

func Fuzz_Nosy_DelegationManagerSession_CalculateDelegationApprovalDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var _delegationApprover common.Address
		fill_err = tp.Fill(&_delegationApprover)
		if fill_err != nil {
			return
		}
		var approverSalt [32]byte
		fill_err = tp.Fill(&approverSalt)
		if fill_err != nil {
			return
		}
		var expiry *big.Int
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || expiry == nil {
			return
		}

		_DelegationManager.CalculateDelegationApprovalDigestHash(staker, operator, _delegationApprover, approverSalt, expiry)
	})
}

func Fuzz_Nosy_DelegationManagerSession_CalculateStakerDelegationDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var _stakerNonce *big.Int
		fill_err = tp.Fill(&_stakerNonce)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var expiry *big.Int
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || _stakerNonce == nil || expiry == nil {
			return
		}

		_DelegationManager.CalculateStakerDelegationDigestHash(staker, _stakerNonce, operator, expiry)
	})
}

func Fuzz_Nosy_DelegationManagerSession_CalculateWithdrawalRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var withdrawal IDelegationManagerWithdrawal
		fill_err = tp.Fill(&withdrawal)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.CalculateWithdrawalRoot(withdrawal)
	})
}

func Fuzz_Nosy_DelegationManagerSession_CompleteQueuedWithdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var withdrawal IDelegationManagerWithdrawal
		fill_err = tp.Fill(&withdrawal)
		if fill_err != nil {
			return
		}
		var tokens []common.Address
		fill_err = tp.Fill(&tokens)
		if fill_err != nil {
			return
		}
		var middlewareTimesIndex *big.Int
		fill_err = tp.Fill(&middlewareTimesIndex)
		if fill_err != nil {
			return
		}
		var receiveAsTokens bool
		fill_err = tp.Fill(&receiveAsTokens)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || middlewareTimesIndex == nil {
			return
		}

		_DelegationManager.CompleteQueuedWithdrawal(withdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
	})
}

func Fuzz_Nosy_DelegationManagerSession_CompleteQueuedWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var withdrawals []IDelegationManagerWithdrawal
		fill_err = tp.Fill(&withdrawals)
		if fill_err != nil {
			return
		}
		var tokens [][]common.Address
		fill_err = tp.Fill(&tokens)
		if fill_err != nil {
			return
		}
		var middlewareTimesIndexes []*big.Int
		fill_err = tp.Fill(&middlewareTimesIndexes)
		if fill_err != nil {
			return
		}
		var receiveAsTokens []bool
		fill_err = tp.Fill(&receiveAsTokens)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.CompleteQueuedWithdrawals(withdrawals, tokens, middlewareTimesIndexes, receiveAsTokens)
	})
}

func Fuzz_Nosy_DelegationManagerSession_CumulativeWithdrawalsQueued__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.CumulativeWithdrawalsQueued(arg0)
	})
}

func Fuzz_Nosy_DelegationManagerSession_DELEGATIONAPPROVALTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.DELEGATIONAPPROVALTYPEHASH()
	})
}

func Fuzz_Nosy_DelegationManagerSession_DOMAINTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.DOMAINTYPEHASH()
	})
}

func Fuzz_Nosy_DelegationManagerSession_DecreaseDelegatedShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var shares *big.Int
		fill_err = tp.Fill(&shares)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || shares == nil {
			return
		}

		_DelegationManager.DecreaseDelegatedShares(staker, strategy, shares)
	})
}

func Fuzz_Nosy_DelegationManagerSession_DelegateTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry
		fill_err = tp.Fill(&approverSignatureAndExpiry)
		if fill_err != nil {
			return
		}
		var approverSalt [32]byte
		fill_err = tp.Fill(&approverSalt)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.DelegateTo(operator, approverSignatureAndExpiry, approverSalt)
	})
}

func Fuzz_Nosy_DelegationManagerSession_DelegateToBySignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var stakerSignatureAndExpiry ISignatureUtilsSignatureWithExpiry
		fill_err = tp.Fill(&stakerSignatureAndExpiry)
		if fill_err != nil {
			return
		}
		var approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry
		fill_err = tp.Fill(&approverSignatureAndExpiry)
		if fill_err != nil {
			return
		}
		var approverSalt [32]byte
		fill_err = tp.Fill(&approverSalt)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.DelegateToBySignature(staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, approverSalt)
	})
}

func Fuzz_Nosy_DelegationManagerSession_DelegatedTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.DelegatedTo(arg0)
	})
}

func Fuzz_Nosy_DelegationManagerSession_DelegationApprover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.DelegationApprover(operator)
	})
}

func Fuzz_Nosy_DelegationManagerSession_DelegationApproverSaltIsSpent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 [32]byte
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.DelegationApproverSaltIsSpent(arg0, arg1)
	})
}

func Fuzz_Nosy_DelegationManagerSession_DomainSeparator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.DomainSeparator()
	})
}

func Fuzz_Nosy_DelegationManagerSession_EarningsReceiver__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.EarningsReceiver(operator)
	})
}

func Fuzz_Nosy_DelegationManagerSession_EigenPodManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.EigenPodManager()
	})
}

func Fuzz_Nosy_DelegationManagerSession_GetDelegatableShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.GetDelegatableShares(staker)
	})
}

func Fuzz_Nosy_DelegationManagerSession_GetOperatorShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var strategies []common.Address
		fill_err = tp.Fill(&strategies)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.GetOperatorShares(operator, strategies)
	})
}

func Fuzz_Nosy_DelegationManagerSession_GetWithdrawalDelay__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var strategies []common.Address
		fill_err = tp.Fill(&strategies)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.GetWithdrawalDelay(strategies)
	})
}

func Fuzz_Nosy_DelegationManagerSession_IncreaseDelegatedShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var shares *big.Int
		fill_err = tp.Fill(&shares)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || shares == nil {
			return
		}

		_DelegationManager.IncreaseDelegatedShares(staker, strategy, shares)
	})
}

func Fuzz_Nosy_DelegationManagerSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var initialOwner common.Address
		fill_err = tp.Fill(&initialOwner)
		if fill_err != nil {
			return
		}
		var _pauserRegistry common.Address
		fill_err = tp.Fill(&_pauserRegistry)
		if fill_err != nil {
			return
		}
		var initialPausedStatus *big.Int
		fill_err = tp.Fill(&initialPausedStatus)
		if fill_err != nil {
			return
		}
		var _minWithdrawalDelayBlocks *big.Int
		fill_err = tp.Fill(&_minWithdrawalDelayBlocks)
		if fill_err != nil {
			return
		}
		var _strategies []common.Address
		fill_err = tp.Fill(&_strategies)
		if fill_err != nil {
			return
		}
		var _withdrawalDelayBlocks []*big.Int
		fill_err = tp.Fill(&_withdrawalDelayBlocks)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || initialPausedStatus == nil || _minWithdrawalDelayBlocks == nil {
			return
		}

		_DelegationManager.Initialize(initialOwner, _pauserRegistry, initialPausedStatus, _minWithdrawalDelayBlocks, _strategies, _withdrawalDelayBlocks)
	})
}

func Fuzz_Nosy_DelegationManagerSession_IsDelegated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.IsDelegated(staker)
	})
}

func Fuzz_Nosy_DelegationManagerSession_IsOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.IsOperator(operator)
	})
}

func Fuzz_Nosy_DelegationManagerSession_MAXSTAKEROPTOUTWINDOWBLOCKS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.MAXSTAKEROPTOUTWINDOWBLOCKS()
	})
}

func Fuzz_Nosy_DelegationManagerSession_MAXWITHDRAWALDELAYBLOCKS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.MAXWITHDRAWALDELAYBLOCKS()
	})
}

func Fuzz_Nosy_DelegationManagerSession_MigrateQueuedWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var withdrawalsToMigrate []IStrategyManagerDeprecatedStructQueuedWithdrawal
		fill_err = tp.Fill(&withdrawalsToMigrate)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.MigrateQueuedWithdrawals(withdrawalsToMigrate)
	})
}

func Fuzz_Nosy_DelegationManagerSession_MinWithdrawalDelayBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.MinWithdrawalDelayBlocks()
	})
}

func Fuzz_Nosy_DelegationManagerSession_ModifyOperatorDetails__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var newOperatorDetails IDelegationManagerOperatorDetails
		fill_err = tp.Fill(&newOperatorDetails)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.ModifyOperatorDetails(newOperatorDetails)
	})
}

func Fuzz_Nosy_DelegationManagerSession_OperatorDetails__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.OperatorDetails(operator)
	})
}

func Fuzz_Nosy_DelegationManagerSession_OperatorShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 common.Address
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.OperatorShares(arg0, arg1)
	})
}

func Fuzz_Nosy_DelegationManagerSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.Owner()
	})
}

func Fuzz_Nosy_DelegationManagerSession_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || newPausedStatus == nil {
			return
		}

		_DelegationManager.Pause(newPausedStatus)
	})
}

func Fuzz_Nosy_DelegationManagerSession_PauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.PauseAll()
	})
}

func Fuzz_Nosy_DelegationManagerSession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var index uint8
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.Paused(index)
	})
}

func Fuzz_Nosy_DelegationManagerSession_Paused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.Paused0()
	})
}

func Fuzz_Nosy_DelegationManagerSession_PauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.PauserRegistry()
	})
}

func Fuzz_Nosy_DelegationManagerSession_PendingWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.PendingWithdrawals(arg0)
	})
}

func Fuzz_Nosy_DelegationManagerSession_QueueWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var queuedWithdrawalParams []IDelegationManagerQueuedWithdrawalParams
		fill_err = tp.Fill(&queuedWithdrawalParams)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.QueueWithdrawals(queuedWithdrawalParams)
	})
}

func Fuzz_Nosy_DelegationManagerSession_RegisterAsOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var registeringOperatorDetails IDelegationManagerOperatorDetails
		fill_err = tp.Fill(&registeringOperatorDetails)
		if fill_err != nil {
			return
		}
		var metadataURI string
		fill_err = tp.Fill(&metadataURI)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.RegisterAsOperator(registeringOperatorDetails, metadataURI)
	})
}

func Fuzz_Nosy_DelegationManagerSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.RenounceOwnership()
	})
}

func Fuzz_Nosy_DelegationManagerSession_STAKERDELEGATIONTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.STAKERDELEGATIONTYPEHASH()
	})
}

func Fuzz_Nosy_DelegationManagerSession_SetMinWithdrawalDelayBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var newMinWithdrawalDelayBlocks *big.Int
		fill_err = tp.Fill(&newMinWithdrawalDelayBlocks)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || newMinWithdrawalDelayBlocks == nil {
			return
		}

		_DelegationManager.SetMinWithdrawalDelayBlocks(newMinWithdrawalDelayBlocks)
	})
}

func Fuzz_Nosy_DelegationManagerSession_SetPauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var newPauserRegistry common.Address
		fill_err = tp.Fill(&newPauserRegistry)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.SetPauserRegistry(newPauserRegistry)
	})
}

func Fuzz_Nosy_DelegationManagerSession_SetStrategyWithdrawalDelayBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var strategies []common.Address
		fill_err = tp.Fill(&strategies)
		if fill_err != nil {
			return
		}
		var withdrawalDelayBlocks []*big.Int
		fill_err = tp.Fill(&withdrawalDelayBlocks)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.SetStrategyWithdrawalDelayBlocks(strategies, withdrawalDelayBlocks)
	})
}

func Fuzz_Nosy_DelegationManagerSession_Slasher__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.Slasher()
	})
}

func Fuzz_Nosy_DelegationManagerSession_StakerNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.StakerNonce(arg0)
	})
}

func Fuzz_Nosy_DelegationManagerSession_StakerOptOutWindowBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.StakerOptOutWindowBlocks(operator)
	})
}

func Fuzz_Nosy_DelegationManagerSession_StrategyManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.StrategyManager()
	})
}

func Fuzz_Nosy_DelegationManagerSession_StrategyWithdrawalDelayBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.StrategyWithdrawalDelayBlocks(arg0)
	})
}

func Fuzz_Nosy_DelegationManagerSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_DelegationManagerSession_Undelegate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.Undelegate(staker)
	})
}

func Fuzz_Nosy_DelegationManagerSession_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || newPausedStatus == nil {
			return
		}

		_DelegationManager.Unpause(newPausedStatus)
	})
}

func Fuzz_Nosy_DelegationManagerSession_UpdateOperatorMetadataURI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var metadataURI string
		fill_err = tp.Fill(&metadataURI)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.UpdateOperatorMetadataURI(metadataURI)
	})
}

func Fuzz_Nosy_DelegationManagerStakerDelegatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerStakerDelegatedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_DelegationManagerStakerDelegatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerStakerDelegatedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_DelegationManagerStakerDelegatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerStakerDelegatedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_DelegationManagerStakerForceUndelegatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerStakerForceUndelegatedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_DelegationManagerStakerForceUndelegatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerStakerForceUndelegatedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_DelegationManagerStakerForceUndelegatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerStakerForceUndelegatedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_DelegationManagerStakerUndelegatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerStakerUndelegatedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_DelegationManagerStakerUndelegatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerStakerUndelegatedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_DelegationManagerStakerUndelegatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerStakerUndelegatedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_DelegationManagerStrategyWithdrawalDelayBlocksSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerStrategyWithdrawalDelayBlocksSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_DelegationManagerStrategyWithdrawalDelayBlocksSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerStrategyWithdrawalDelayBlocksSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_DelegationManagerStrategyWithdrawalDelayBlocksSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerStrategyWithdrawalDelayBlocksSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_DelegationManagerTransactor_CompleteQueuedWithdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactor
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var withdrawal IDelegationManagerWithdrawal
		fill_err = tp.Fill(&withdrawal)
		if fill_err != nil {
			return
		}
		var tokens []common.Address
		fill_err = tp.Fill(&tokens)
		if fill_err != nil {
			return
		}
		var middlewareTimesIndex *big.Int
		fill_err = tp.Fill(&middlewareTimesIndex)
		if fill_err != nil {
			return
		}
		var receiveAsTokens bool
		fill_err = tp.Fill(&receiveAsTokens)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil || middlewareTimesIndex == nil {
			return
		}

		_DelegationManager.CompleteQueuedWithdrawal(opts, withdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
	})
}

func Fuzz_Nosy_DelegationManagerTransactor_CompleteQueuedWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactor
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var withdrawals []IDelegationManagerWithdrawal
		fill_err = tp.Fill(&withdrawals)
		if fill_err != nil {
			return
		}
		var tokens [][]common.Address
		fill_err = tp.Fill(&tokens)
		if fill_err != nil {
			return
		}
		var middlewareTimesIndexes []*big.Int
		fill_err = tp.Fill(&middlewareTimesIndexes)
		if fill_err != nil {
			return
		}
		var receiveAsTokens []bool
		fill_err = tp.Fill(&receiveAsTokens)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.CompleteQueuedWithdrawals(opts, withdrawals, tokens, middlewareTimesIndexes, receiveAsTokens)
	})
}

func Fuzz_Nosy_DelegationManagerTransactor_DecreaseDelegatedShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactor
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var shares *big.Int
		fill_err = tp.Fill(&shares)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil || shares == nil {
			return
		}

		_DelegationManager.DecreaseDelegatedShares(opts, staker, strategy, shares)
	})
}

func Fuzz_Nosy_DelegationManagerTransactor_DelegateTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactor
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry
		fill_err = tp.Fill(&approverSignatureAndExpiry)
		if fill_err != nil {
			return
		}
		var approverSalt [32]byte
		fill_err = tp.Fill(&approverSalt)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.DelegateTo(opts, operator, approverSignatureAndExpiry, approverSalt)
	})
}

func Fuzz_Nosy_DelegationManagerTransactor_DelegateToBySignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactor
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var stakerSignatureAndExpiry ISignatureUtilsSignatureWithExpiry
		fill_err = tp.Fill(&stakerSignatureAndExpiry)
		if fill_err != nil {
			return
		}
		var approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry
		fill_err = tp.Fill(&approverSignatureAndExpiry)
		if fill_err != nil {
			return
		}
		var approverSalt [32]byte
		fill_err = tp.Fill(&approverSalt)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.DelegateToBySignature(opts, staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, approverSalt)
	})
}

func Fuzz_Nosy_DelegationManagerTransactor_IncreaseDelegatedShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactor
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var shares *big.Int
		fill_err = tp.Fill(&shares)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil || shares == nil {
			return
		}

		_DelegationManager.IncreaseDelegatedShares(opts, staker, strategy, shares)
	})
}

func Fuzz_Nosy_DelegationManagerTransactor_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactor
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var initialOwner common.Address
		fill_err = tp.Fill(&initialOwner)
		if fill_err != nil {
			return
		}
		var _pauserRegistry common.Address
		fill_err = tp.Fill(&_pauserRegistry)
		if fill_err != nil {
			return
		}
		var initialPausedStatus *big.Int
		fill_err = tp.Fill(&initialPausedStatus)
		if fill_err != nil {
			return
		}
		var _minWithdrawalDelayBlocks *big.Int
		fill_err = tp.Fill(&_minWithdrawalDelayBlocks)
		if fill_err != nil {
			return
		}
		var _strategies []common.Address
		fill_err = tp.Fill(&_strategies)
		if fill_err != nil {
			return
		}
		var _withdrawalDelayBlocks []*big.Int
		fill_err = tp.Fill(&_withdrawalDelayBlocks)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil || initialPausedStatus == nil || _minWithdrawalDelayBlocks == nil {
			return
		}

		_DelegationManager.Initialize(opts, initialOwner, _pauserRegistry, initialPausedStatus, _minWithdrawalDelayBlocks, _strategies, _withdrawalDelayBlocks)
	})
}

func Fuzz_Nosy_DelegationManagerTransactor_MigrateQueuedWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactor
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var withdrawalsToMigrate []IStrategyManagerDeprecatedStructQueuedWithdrawal
		fill_err = tp.Fill(&withdrawalsToMigrate)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.MigrateQueuedWithdrawals(opts, withdrawalsToMigrate)
	})
}

func Fuzz_Nosy_DelegationManagerTransactor_ModifyOperatorDetails__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactor
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newOperatorDetails IDelegationManagerOperatorDetails
		fill_err = tp.Fill(&newOperatorDetails)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.ModifyOperatorDetails(opts, newOperatorDetails)
	})
}

func Fuzz_Nosy_DelegationManagerTransactor_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactor
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil || newPausedStatus == nil {
			return
		}

		_DelegationManager.Pause(opts, newPausedStatus)
	})
}

func Fuzz_Nosy_DelegationManagerTransactor_PauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactor
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.PauseAll(opts)
	})
}

func Fuzz_Nosy_DelegationManagerTransactor_QueueWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactor
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var queuedWithdrawalParams []IDelegationManagerQueuedWithdrawalParams
		fill_err = tp.Fill(&queuedWithdrawalParams)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.QueueWithdrawals(opts, queuedWithdrawalParams)
	})
}

func Fuzz_Nosy_DelegationManagerTransactor_RegisterAsOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactor
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var registeringOperatorDetails IDelegationManagerOperatorDetails
		fill_err = tp.Fill(&registeringOperatorDetails)
		if fill_err != nil {
			return
		}
		var metadataURI string
		fill_err = tp.Fill(&metadataURI)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.RegisterAsOperator(opts, registeringOperatorDetails, metadataURI)
	})
}

func Fuzz_Nosy_DelegationManagerTransactor_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactor
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.RenounceOwnership(opts)
	})
}

func Fuzz_Nosy_DelegationManagerTransactor_SetMinWithdrawalDelayBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactor
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newMinWithdrawalDelayBlocks *big.Int
		fill_err = tp.Fill(&newMinWithdrawalDelayBlocks)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil || newMinWithdrawalDelayBlocks == nil {
			return
		}

		_DelegationManager.SetMinWithdrawalDelayBlocks(opts, newMinWithdrawalDelayBlocks)
	})
}

func Fuzz_Nosy_DelegationManagerTransactor_SetPauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactor
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newPauserRegistry common.Address
		fill_err = tp.Fill(&newPauserRegistry)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.SetPauserRegistry(opts, newPauserRegistry)
	})
}

func Fuzz_Nosy_DelegationManagerTransactor_SetStrategyWithdrawalDelayBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactor
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var strategies []common.Address
		fill_err = tp.Fill(&strategies)
		if fill_err != nil {
			return
		}
		var withdrawalDelayBlocks []*big.Int
		fill_err = tp.Fill(&withdrawalDelayBlocks)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.SetStrategyWithdrawalDelayBlocks(opts, strategies, withdrawalDelayBlocks)
	})
}

func Fuzz_Nosy_DelegationManagerTransactor_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactor
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.TransferOwnership(opts, newOwner)
	})
}

func Fuzz_Nosy_DelegationManagerTransactor_Undelegate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactor
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.Undelegate(opts, staker)
	})
}

func Fuzz_Nosy_DelegationManagerTransactor_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactor
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil || newPausedStatus == nil {
			return
		}

		_DelegationManager.Unpause(opts, newPausedStatus)
	})
}

func Fuzz_Nosy_DelegationManagerTransactor_UpdateOperatorMetadataURI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactor
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var metadataURI string
		fill_err = tp.Fill(&metadataURI)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.UpdateOperatorMetadataURI(opts, metadataURI)
	})
}

// skipping Fuzz_Nosy_DelegationManagerTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_DelegationManagerTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactorRaw
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || opts == nil {
			return
		}

		_DelegationManager.Transfer(opts)
	})
}

func Fuzz_Nosy_DelegationManagerTransactorSession_CompleteQueuedWithdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactorSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var withdrawal IDelegationManagerWithdrawal
		fill_err = tp.Fill(&withdrawal)
		if fill_err != nil {
			return
		}
		var tokens []common.Address
		fill_err = tp.Fill(&tokens)
		if fill_err != nil {
			return
		}
		var middlewareTimesIndex *big.Int
		fill_err = tp.Fill(&middlewareTimesIndex)
		if fill_err != nil {
			return
		}
		var receiveAsTokens bool
		fill_err = tp.Fill(&receiveAsTokens)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || middlewareTimesIndex == nil {
			return
		}

		_DelegationManager.CompleteQueuedWithdrawal(withdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
	})
}

func Fuzz_Nosy_DelegationManagerTransactorSession_CompleteQueuedWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactorSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var withdrawals []IDelegationManagerWithdrawal
		fill_err = tp.Fill(&withdrawals)
		if fill_err != nil {
			return
		}
		var tokens [][]common.Address
		fill_err = tp.Fill(&tokens)
		if fill_err != nil {
			return
		}
		var middlewareTimesIndexes []*big.Int
		fill_err = tp.Fill(&middlewareTimesIndexes)
		if fill_err != nil {
			return
		}
		var receiveAsTokens []bool
		fill_err = tp.Fill(&receiveAsTokens)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.CompleteQueuedWithdrawals(withdrawals, tokens, middlewareTimesIndexes, receiveAsTokens)
	})
}

func Fuzz_Nosy_DelegationManagerTransactorSession_DecreaseDelegatedShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactorSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var shares *big.Int
		fill_err = tp.Fill(&shares)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || shares == nil {
			return
		}

		_DelegationManager.DecreaseDelegatedShares(staker, strategy, shares)
	})
}

func Fuzz_Nosy_DelegationManagerTransactorSession_DelegateTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactorSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry
		fill_err = tp.Fill(&approverSignatureAndExpiry)
		if fill_err != nil {
			return
		}
		var approverSalt [32]byte
		fill_err = tp.Fill(&approverSalt)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.DelegateTo(operator, approverSignatureAndExpiry, approverSalt)
	})
}

func Fuzz_Nosy_DelegationManagerTransactorSession_DelegateToBySignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactorSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var stakerSignatureAndExpiry ISignatureUtilsSignatureWithExpiry
		fill_err = tp.Fill(&stakerSignatureAndExpiry)
		if fill_err != nil {
			return
		}
		var approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry
		fill_err = tp.Fill(&approverSignatureAndExpiry)
		if fill_err != nil {
			return
		}
		var approverSalt [32]byte
		fill_err = tp.Fill(&approverSalt)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.DelegateToBySignature(staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, approverSalt)
	})
}

func Fuzz_Nosy_DelegationManagerTransactorSession_IncreaseDelegatedShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactorSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var shares *big.Int
		fill_err = tp.Fill(&shares)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || shares == nil {
			return
		}

		_DelegationManager.IncreaseDelegatedShares(staker, strategy, shares)
	})
}

func Fuzz_Nosy_DelegationManagerTransactorSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactorSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var initialOwner common.Address
		fill_err = tp.Fill(&initialOwner)
		if fill_err != nil {
			return
		}
		var _pauserRegistry common.Address
		fill_err = tp.Fill(&_pauserRegistry)
		if fill_err != nil {
			return
		}
		var initialPausedStatus *big.Int
		fill_err = tp.Fill(&initialPausedStatus)
		if fill_err != nil {
			return
		}
		var _minWithdrawalDelayBlocks *big.Int
		fill_err = tp.Fill(&_minWithdrawalDelayBlocks)
		if fill_err != nil {
			return
		}
		var _strategies []common.Address
		fill_err = tp.Fill(&_strategies)
		if fill_err != nil {
			return
		}
		var _withdrawalDelayBlocks []*big.Int
		fill_err = tp.Fill(&_withdrawalDelayBlocks)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || initialPausedStatus == nil || _minWithdrawalDelayBlocks == nil {
			return
		}

		_DelegationManager.Initialize(initialOwner, _pauserRegistry, initialPausedStatus, _minWithdrawalDelayBlocks, _strategies, _withdrawalDelayBlocks)
	})
}

func Fuzz_Nosy_DelegationManagerTransactorSession_MigrateQueuedWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactorSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var withdrawalsToMigrate []IStrategyManagerDeprecatedStructQueuedWithdrawal
		fill_err = tp.Fill(&withdrawalsToMigrate)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.MigrateQueuedWithdrawals(withdrawalsToMigrate)
	})
}

func Fuzz_Nosy_DelegationManagerTransactorSession_ModifyOperatorDetails__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactorSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var newOperatorDetails IDelegationManagerOperatorDetails
		fill_err = tp.Fill(&newOperatorDetails)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.ModifyOperatorDetails(newOperatorDetails)
	})
}

func Fuzz_Nosy_DelegationManagerTransactorSession_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactorSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || newPausedStatus == nil {
			return
		}

		_DelegationManager.Pause(newPausedStatus)
	})
}

func Fuzz_Nosy_DelegationManagerTransactorSession_PauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactorSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.PauseAll()
	})
}

func Fuzz_Nosy_DelegationManagerTransactorSession_QueueWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactorSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var queuedWithdrawalParams []IDelegationManagerQueuedWithdrawalParams
		fill_err = tp.Fill(&queuedWithdrawalParams)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.QueueWithdrawals(queuedWithdrawalParams)
	})
}

func Fuzz_Nosy_DelegationManagerTransactorSession_RegisterAsOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactorSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var registeringOperatorDetails IDelegationManagerOperatorDetails
		fill_err = tp.Fill(&registeringOperatorDetails)
		if fill_err != nil {
			return
		}
		var metadataURI string
		fill_err = tp.Fill(&metadataURI)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.RegisterAsOperator(registeringOperatorDetails, metadataURI)
	})
}

func Fuzz_Nosy_DelegationManagerTransactorSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactorSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.RenounceOwnership()
	})
}

func Fuzz_Nosy_DelegationManagerTransactorSession_SetMinWithdrawalDelayBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactorSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var newMinWithdrawalDelayBlocks *big.Int
		fill_err = tp.Fill(&newMinWithdrawalDelayBlocks)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || newMinWithdrawalDelayBlocks == nil {
			return
		}

		_DelegationManager.SetMinWithdrawalDelayBlocks(newMinWithdrawalDelayBlocks)
	})
}

func Fuzz_Nosy_DelegationManagerTransactorSession_SetPauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactorSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var newPauserRegistry common.Address
		fill_err = tp.Fill(&newPauserRegistry)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.SetPauserRegistry(newPauserRegistry)
	})
}

func Fuzz_Nosy_DelegationManagerTransactorSession_SetStrategyWithdrawalDelayBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactorSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var strategies []common.Address
		fill_err = tp.Fill(&strategies)
		if fill_err != nil {
			return
		}
		var withdrawalDelayBlocks []*big.Int
		fill_err = tp.Fill(&withdrawalDelayBlocks)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.SetStrategyWithdrawalDelayBlocks(strategies, withdrawalDelayBlocks)
	})
}

func Fuzz_Nosy_DelegationManagerTransactorSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactorSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_DelegationManagerTransactorSession_Undelegate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactorSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.Undelegate(staker)
	})
}

func Fuzz_Nosy_DelegationManagerTransactorSession_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactorSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil || newPausedStatus == nil {
			return
		}

		_DelegationManager.Unpause(newPausedStatus)
	})
}

func Fuzz_Nosy_DelegationManagerTransactorSession_UpdateOperatorMetadataURI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DelegationManager *DelegationManagerTransactorSession
		fill_err = tp.Fill(&_DelegationManager)
		if fill_err != nil {
			return
		}
		var metadataURI string
		fill_err = tp.Fill(&metadataURI)
		if fill_err != nil {
			return
		}
		if _DelegationManager == nil {
			return
		}

		_DelegationManager.UpdateOperatorMetadataURI(metadataURI)
	})
}

func Fuzz_Nosy_DelegationManagerUnpausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_DelegationManagerUnpausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_DelegationManagerUnpausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_DelegationManagerWithdrawalCompletedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerWithdrawalCompletedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_DelegationManagerWithdrawalCompletedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerWithdrawalCompletedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_DelegationManagerWithdrawalCompletedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerWithdrawalCompletedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_DelegationManagerWithdrawalMigratedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerWithdrawalMigratedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_DelegationManagerWithdrawalMigratedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerWithdrawalMigratedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_DelegationManagerWithdrawalMigratedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerWithdrawalMigratedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_DelegationManagerWithdrawalQueuedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerWithdrawalQueuedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_DelegationManagerWithdrawalQueuedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerWithdrawalQueuedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_DelegationManagerWithdrawalQueuedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DelegationManagerWithdrawalQueuedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_FeeOracleV1BaseGasLimitSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *FeeOracleV1BaseGasLimitSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_FeeOracleV1BaseGasLimitSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *FeeOracleV1BaseGasLimitSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_FeeOracleV1BaseGasLimitSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *FeeOracleV1BaseGasLimitSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_FeeOracleV1Caller_BaseGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Caller
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil {
			return
		}

		_FeeOracleV1.BaseGasLimit(opts)
	})
}

func Fuzz_Nosy_FeeOracleV1Caller_CONVERSIONRATEDENOM__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Caller
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil {
			return
		}

		_FeeOracleV1.CONVERSIONRATEDENOM(opts)
	})
}

func Fuzz_Nosy_FeeOracleV1Caller_FeeFor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Caller
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var destChainId uint64
		fill_err = tp.Fill(&destChainId)
		if fill_err != nil {
			return
		}
		var d4 []byte
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil {
			return
		}

		_FeeOracleV1.FeeFor(opts, destChainId, d4, gasLimit)
	})
}

func Fuzz_Nosy_FeeOracleV1Caller_FeeParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Caller
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var chainId uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil {
			return
		}

		_FeeOracleV1.FeeParams(opts, chainId)
	})
}

func Fuzz_Nosy_FeeOracleV1Caller_GasPriceOn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Caller
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var chainId uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil {
			return
		}

		_FeeOracleV1.GasPriceOn(opts, chainId)
	})
}

func Fuzz_Nosy_FeeOracleV1Caller_Manager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Caller
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil {
			return
		}

		_FeeOracleV1.Manager(opts)
	})
}

func Fuzz_Nosy_FeeOracleV1Caller_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Caller
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil {
			return
		}

		_FeeOracleV1.Owner(opts)
	})
}

func Fuzz_Nosy_FeeOracleV1Caller_PostsTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Caller
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var chainId uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil {
			return
		}

		_FeeOracleV1.PostsTo(opts, chainId)
	})
}

func Fuzz_Nosy_FeeOracleV1Caller_ProtocolFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Caller
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil {
			return
		}

		_FeeOracleV1.ProtocolFee(opts)
	})
}

func Fuzz_Nosy_FeeOracleV1Caller_ToNativeRate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Caller
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var chainId uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil {
			return
		}

		_FeeOracleV1.ToNativeRate(opts, chainId)
	})
}

// skipping Fuzz_Nosy_FeeOracleV1CallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_FeeOracleV1CallerSession_BaseGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1CallerSession
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.BaseGasLimit()
	})
}

func Fuzz_Nosy_FeeOracleV1CallerSession_CONVERSIONRATEDENOM__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1CallerSession
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.CONVERSIONRATEDENOM()
	})
}

func Fuzz_Nosy_FeeOracleV1CallerSession_FeeFor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1CallerSession
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var destChainId uint64
		fill_err = tp.Fill(&destChainId)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.FeeFor(destChainId, d3, gasLimit)
	})
}

func Fuzz_Nosy_FeeOracleV1CallerSession_FeeParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1CallerSession
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var chainId uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.FeeParams(chainId)
	})
}

func Fuzz_Nosy_FeeOracleV1CallerSession_GasPriceOn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1CallerSession
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var chainId uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.GasPriceOn(chainId)
	})
}

func Fuzz_Nosy_FeeOracleV1CallerSession_Manager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1CallerSession
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.Manager()
	})
}

func Fuzz_Nosy_FeeOracleV1CallerSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1CallerSession
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.Owner()
	})
}

func Fuzz_Nosy_FeeOracleV1CallerSession_PostsTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1CallerSession
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var chainId uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.PostsTo(chainId)
	})
}

func Fuzz_Nosy_FeeOracleV1CallerSession_ProtocolFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1CallerSession
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.ProtocolFee()
	})
}

func Fuzz_Nosy_FeeOracleV1CallerSession_ToNativeRate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1CallerSession
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var chainId uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.ToNativeRate(chainId)
	})
}

func Fuzz_Nosy_FeeOracleV1FeeParamsSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *FeeOracleV1FeeParamsSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_FeeOracleV1FeeParamsSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *FeeOracleV1FeeParamsSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_FeeOracleV1FeeParamsSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *FeeOracleV1FeeParamsSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_FeeOracleV1Filterer_FilterBaseGasLimitSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Filterer
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil {
			return
		}

		_FeeOracleV1.FilterBaseGasLimitSet(opts)
	})
}

func Fuzz_Nosy_FeeOracleV1Filterer_FilterFeeParamsSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Filterer
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil {
			return
		}

		_FeeOracleV1.FilterFeeParamsSet(opts)
	})
}

func Fuzz_Nosy_FeeOracleV1Filterer_FilterGasPriceSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Filterer
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil {
			return
		}

		_FeeOracleV1.FilterGasPriceSet(opts)
	})
}

func Fuzz_Nosy_FeeOracleV1Filterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Filterer
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil {
			return
		}

		_FeeOracleV1.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_FeeOracleV1Filterer_FilterManagerSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Filterer
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil {
			return
		}

		_FeeOracleV1.FilterManagerSet(opts)
	})
}

func Fuzz_Nosy_FeeOracleV1Filterer_FilterOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Filterer
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var previousOwner []common.Address
		fill_err = tp.Fill(&previousOwner)
		if fill_err != nil {
			return
		}
		var newOwner []common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil {
			return
		}

		_FeeOracleV1.FilterOwnershipTransferred(opts, previousOwner, newOwner)
	})
}

func Fuzz_Nosy_FeeOracleV1Filterer_FilterProtocolFeeSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Filterer
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil {
			return
		}

		_FeeOracleV1.FilterProtocolFeeSet(opts)
	})
}

func Fuzz_Nosy_FeeOracleV1Filterer_FilterToNativeRateSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Filterer
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil {
			return
		}

		_FeeOracleV1.FilterToNativeRateSet(opts)
	})
}

func Fuzz_Nosy_FeeOracleV1Filterer_ParseBaseGasLimitSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Filterer
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.ParseBaseGasLimitSet(log)
	})
}

func Fuzz_Nosy_FeeOracleV1Filterer_ParseFeeParamsSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Filterer
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.ParseFeeParamsSet(log)
	})
}

func Fuzz_Nosy_FeeOracleV1Filterer_ParseGasPriceSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Filterer
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.ParseGasPriceSet(log)
	})
}

func Fuzz_Nosy_FeeOracleV1Filterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Filterer
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.ParseInitialized(log)
	})
}

func Fuzz_Nosy_FeeOracleV1Filterer_ParseManagerSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Filterer
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.ParseManagerSet(log)
	})
}

func Fuzz_Nosy_FeeOracleV1Filterer_ParseOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Filterer
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.ParseOwnershipTransferred(log)
	})
}

func Fuzz_Nosy_FeeOracleV1Filterer_ParseProtocolFeeSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Filterer
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.ParseProtocolFeeSet(log)
	})
}

func Fuzz_Nosy_FeeOracleV1Filterer_ParseToNativeRateSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Filterer
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.ParseToNativeRateSet(log)
	})
}

// skipping Fuzz_Nosy_FeeOracleV1Filterer_WatchBaseGasLimitSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.FeeOracleV1BaseGasLimitSet

// skipping Fuzz_Nosy_FeeOracleV1Filterer_WatchFeeParamsSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.FeeOracleV1FeeParamsSet

// skipping Fuzz_Nosy_FeeOracleV1Filterer_WatchGasPriceSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.FeeOracleV1GasPriceSet

// skipping Fuzz_Nosy_FeeOracleV1Filterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.FeeOracleV1Initialized

// skipping Fuzz_Nosy_FeeOracleV1Filterer_WatchManagerSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.FeeOracleV1ManagerSet

// skipping Fuzz_Nosy_FeeOracleV1Filterer_WatchOwnershipTransferred__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.FeeOracleV1OwnershipTransferred

// skipping Fuzz_Nosy_FeeOracleV1Filterer_WatchProtocolFeeSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.FeeOracleV1ProtocolFeeSet

// skipping Fuzz_Nosy_FeeOracleV1Filterer_WatchToNativeRateSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.FeeOracleV1ToNativeRateSet

func Fuzz_Nosy_FeeOracleV1GasPriceSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *FeeOracleV1GasPriceSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_FeeOracleV1GasPriceSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *FeeOracleV1GasPriceSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_FeeOracleV1GasPriceSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *FeeOracleV1GasPriceSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_FeeOracleV1InitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *FeeOracleV1InitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_FeeOracleV1InitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *FeeOracleV1InitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_FeeOracleV1InitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *FeeOracleV1InitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_FeeOracleV1ManagerSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *FeeOracleV1ManagerSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_FeeOracleV1ManagerSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *FeeOracleV1ManagerSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_FeeOracleV1ManagerSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *FeeOracleV1ManagerSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_FeeOracleV1OwnershipTransferredIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *FeeOracleV1OwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_FeeOracleV1OwnershipTransferredIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *FeeOracleV1OwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_FeeOracleV1OwnershipTransferredIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *FeeOracleV1OwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_FeeOracleV1ProtocolFeeSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *FeeOracleV1ProtocolFeeSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_FeeOracleV1ProtocolFeeSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *FeeOracleV1ProtocolFeeSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_FeeOracleV1ProtocolFeeSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *FeeOracleV1ProtocolFeeSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

// skipping Fuzz_Nosy_FeeOracleV1Raw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_FeeOracleV1Raw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_FeeOracleV1Raw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Raw
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil {
			return
		}

		_FeeOracleV1.Transfer(opts)
	})
}

func Fuzz_Nosy_FeeOracleV1Session_BaseGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Session
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.BaseGasLimit()
	})
}

func Fuzz_Nosy_FeeOracleV1Session_BulkSetFeeParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Session
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var params []IFeeOracleV1ChainFeeParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.BulkSetFeeParams(params)
	})
}

func Fuzz_Nosy_FeeOracleV1Session_CONVERSIONRATEDENOM__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Session
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.CONVERSIONRATEDENOM()
	})
}

func Fuzz_Nosy_FeeOracleV1Session_FeeFor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Session
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var destChainId uint64
		fill_err = tp.Fill(&destChainId)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.FeeFor(destChainId, d3, gasLimit)
	})
}

func Fuzz_Nosy_FeeOracleV1Session_FeeParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Session
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var chainId uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.FeeParams(chainId)
	})
}

func Fuzz_Nosy_FeeOracleV1Session_GasPriceOn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Session
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var chainId uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.GasPriceOn(chainId)
	})
}

func Fuzz_Nosy_FeeOracleV1Session_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Session
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var owner_ common.Address
		fill_err = tp.Fill(&owner_)
		if fill_err != nil {
			return
		}
		var manager_ common.Address
		fill_err = tp.Fill(&manager_)
		if fill_err != nil {
			return
		}
		var baseGasLimit_ uint64
		fill_err = tp.Fill(&baseGasLimit_)
		if fill_err != nil {
			return
		}
		var protocolFee_ *big.Int
		fill_err = tp.Fill(&protocolFee_)
		if fill_err != nil {
			return
		}
		var params []IFeeOracleV1ChainFeeParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || protocolFee_ == nil {
			return
		}

		_FeeOracleV1.Initialize(owner_, manager_, baseGasLimit_, protocolFee_, params)
	})
}

func Fuzz_Nosy_FeeOracleV1Session_Manager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Session
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.Manager()
	})
}

func Fuzz_Nosy_FeeOracleV1Session_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Session
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.Owner()
	})
}

func Fuzz_Nosy_FeeOracleV1Session_PostsTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Session
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var chainId uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.PostsTo(chainId)
	})
}

func Fuzz_Nosy_FeeOracleV1Session_ProtocolFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Session
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.ProtocolFee()
	})
}

func Fuzz_Nosy_FeeOracleV1Session_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Session
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.RenounceOwnership()
	})
}

func Fuzz_Nosy_FeeOracleV1Session_SetBaseGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Session
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.SetBaseGasLimit(gasLimit)
	})
}

func Fuzz_Nosy_FeeOracleV1Session_SetGasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Session
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var chainId uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || gasPrice == nil {
			return
		}

		_FeeOracleV1.SetGasPrice(chainId, gasPrice)
	})
}

func Fuzz_Nosy_FeeOracleV1Session_SetManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Session
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var manager_ common.Address
		fill_err = tp.Fill(&manager_)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.SetManager(manager_)
	})
}

func Fuzz_Nosy_FeeOracleV1Session_SetProtocolFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Session
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var fee *big.Int
		fill_err = tp.Fill(&fee)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || fee == nil {
			return
		}

		_FeeOracleV1.SetProtocolFee(fee)
	})
}

func Fuzz_Nosy_FeeOracleV1Session_SetToNativeRate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Session
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var chainId uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var rate *big.Int
		fill_err = tp.Fill(&rate)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || rate == nil {
			return
		}

		_FeeOracleV1.SetToNativeRate(chainId, rate)
	})
}

func Fuzz_Nosy_FeeOracleV1Session_ToNativeRate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Session
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var chainId uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.ToNativeRate(chainId)
	})
}

func Fuzz_Nosy_FeeOracleV1Session_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Session
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_FeeOracleV1ToNativeRateSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *FeeOracleV1ToNativeRateSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_FeeOracleV1ToNativeRateSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *FeeOracleV1ToNativeRateSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_FeeOracleV1ToNativeRateSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *FeeOracleV1ToNativeRateSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_FeeOracleV1Transactor_BulkSetFeeParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Transactor
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var params []IFeeOracleV1ChainFeeParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil {
			return
		}

		_FeeOracleV1.BulkSetFeeParams(opts, params)
	})
}

func Fuzz_Nosy_FeeOracleV1Transactor_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Transactor
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var owner_ common.Address
		fill_err = tp.Fill(&owner_)
		if fill_err != nil {
			return
		}
		var manager_ common.Address
		fill_err = tp.Fill(&manager_)
		if fill_err != nil {
			return
		}
		var baseGasLimit_ uint64
		fill_err = tp.Fill(&baseGasLimit_)
		if fill_err != nil {
			return
		}
		var protocolFee_ *big.Int
		fill_err = tp.Fill(&protocolFee_)
		if fill_err != nil {
			return
		}
		var params []IFeeOracleV1ChainFeeParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil || protocolFee_ == nil {
			return
		}

		_FeeOracleV1.Initialize(opts, owner_, manager_, baseGasLimit_, protocolFee_, params)
	})
}

func Fuzz_Nosy_FeeOracleV1Transactor_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Transactor
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil {
			return
		}

		_FeeOracleV1.RenounceOwnership(opts)
	})
}

func Fuzz_Nosy_FeeOracleV1Transactor_SetBaseGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Transactor
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil {
			return
		}

		_FeeOracleV1.SetBaseGasLimit(opts, gasLimit)
	})
}

func Fuzz_Nosy_FeeOracleV1Transactor_SetGasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Transactor
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var chainId uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil || gasPrice == nil {
			return
		}

		_FeeOracleV1.SetGasPrice(opts, chainId, gasPrice)
	})
}

func Fuzz_Nosy_FeeOracleV1Transactor_SetManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Transactor
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var manager_ common.Address
		fill_err = tp.Fill(&manager_)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil {
			return
		}

		_FeeOracleV1.SetManager(opts, manager_)
	})
}

func Fuzz_Nosy_FeeOracleV1Transactor_SetProtocolFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Transactor
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var fee *big.Int
		fill_err = tp.Fill(&fee)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil || fee == nil {
			return
		}

		_FeeOracleV1.SetProtocolFee(opts, fee)
	})
}

func Fuzz_Nosy_FeeOracleV1Transactor_SetToNativeRate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Transactor
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var chainId uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var rate *big.Int
		fill_err = tp.Fill(&rate)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil || rate == nil {
			return
		}

		_FeeOracleV1.SetToNativeRate(opts, chainId, rate)
	})
}

func Fuzz_Nosy_FeeOracleV1Transactor_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1Transactor
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil {
			return
		}

		_FeeOracleV1.TransferOwnership(opts, newOwner)
	})
}

// skipping Fuzz_Nosy_FeeOracleV1TransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_FeeOracleV1TransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1TransactorRaw
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || opts == nil {
			return
		}

		_FeeOracleV1.Transfer(opts)
	})
}

func Fuzz_Nosy_FeeOracleV1TransactorSession_BulkSetFeeParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1TransactorSession
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var params []IFeeOracleV1ChainFeeParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.BulkSetFeeParams(params)
	})
}

func Fuzz_Nosy_FeeOracleV1TransactorSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1TransactorSession
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var owner_ common.Address
		fill_err = tp.Fill(&owner_)
		if fill_err != nil {
			return
		}
		var manager_ common.Address
		fill_err = tp.Fill(&manager_)
		if fill_err != nil {
			return
		}
		var baseGasLimit_ uint64
		fill_err = tp.Fill(&baseGasLimit_)
		if fill_err != nil {
			return
		}
		var protocolFee_ *big.Int
		fill_err = tp.Fill(&protocolFee_)
		if fill_err != nil {
			return
		}
		var params []IFeeOracleV1ChainFeeParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || protocolFee_ == nil {
			return
		}

		_FeeOracleV1.Initialize(owner_, manager_, baseGasLimit_, protocolFee_, params)
	})
}

func Fuzz_Nosy_FeeOracleV1TransactorSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1TransactorSession
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.RenounceOwnership()
	})
}

func Fuzz_Nosy_FeeOracleV1TransactorSession_SetBaseGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1TransactorSession
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.SetBaseGasLimit(gasLimit)
	})
}

func Fuzz_Nosy_FeeOracleV1TransactorSession_SetGasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1TransactorSession
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var chainId uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || gasPrice == nil {
			return
		}

		_FeeOracleV1.SetGasPrice(chainId, gasPrice)
	})
}

func Fuzz_Nosy_FeeOracleV1TransactorSession_SetManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1TransactorSession
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var manager_ common.Address
		fill_err = tp.Fill(&manager_)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.SetManager(manager_)
	})
}

func Fuzz_Nosy_FeeOracleV1TransactorSession_SetProtocolFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1TransactorSession
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var fee *big.Int
		fill_err = tp.Fill(&fee)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || fee == nil {
			return
		}

		_FeeOracleV1.SetProtocolFee(fee)
	})
}

func Fuzz_Nosy_FeeOracleV1TransactorSession_SetToNativeRate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1TransactorSession
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var chainId uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var rate *big.Int
		fill_err = tp.Fill(&rate)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil || rate == nil {
			return
		}

		_FeeOracleV1.SetToNativeRate(chainId, rate)
	})
}

func Fuzz_Nosy_FeeOracleV1TransactorSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _FeeOracleV1 *FeeOracleV1TransactorSession
		fill_err = tp.Fill(&_FeeOracleV1)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _FeeOracleV1 == nil {
			return
		}

		_FeeOracleV1.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_MockERC20ApprovalIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *MockERC20ApprovalIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_MockERC20ApprovalIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *MockERC20ApprovalIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_MockERC20ApprovalIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *MockERC20ApprovalIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_MockERC20Caller_Allowance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Caller
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var owner common.Address
		fill_err = tp.Fill(&owner)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || opts == nil {
			return
		}

		_MockERC20.Allowance(opts, owner, spender)
	})
}

func Fuzz_Nosy_MockERC20Caller_BalanceOf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Caller
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || opts == nil {
			return
		}

		_MockERC20.BalanceOf(opts, account)
	})
}

func Fuzz_Nosy_MockERC20Caller_Decimals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Caller
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || opts == nil {
			return
		}

		_MockERC20.Decimals(opts)
	})
}

func Fuzz_Nosy_MockERC20Caller_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Caller
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || opts == nil {
			return
		}

		_MockERC20.Name(opts)
	})
}

func Fuzz_Nosy_MockERC20Caller_Symbol__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Caller
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || opts == nil {
			return
		}

		_MockERC20.Symbol(opts)
	})
}

func Fuzz_Nosy_MockERC20Caller_TotalSupply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Caller
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || opts == nil {
			return
		}

		_MockERC20.TotalSupply(opts)
	})
}

// skipping Fuzz_Nosy_MockERC20CallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_MockERC20CallerSession_Allowance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20CallerSession
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var owner common.Address
		fill_err = tp.Fill(&owner)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil {
			return
		}

		_MockERC20.Allowance(owner, spender)
	})
}

func Fuzz_Nosy_MockERC20CallerSession_BalanceOf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20CallerSession
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil {
			return
		}

		_MockERC20.BalanceOf(account)
	})
}

func Fuzz_Nosy_MockERC20CallerSession_Decimals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20CallerSession
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil {
			return
		}

		_MockERC20.Decimals()
	})
}

func Fuzz_Nosy_MockERC20CallerSession_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20CallerSession
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil {
			return
		}

		_MockERC20.Name()
	})
}

func Fuzz_Nosy_MockERC20CallerSession_Symbol__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20CallerSession
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil {
			return
		}

		_MockERC20.Symbol()
	})
}

func Fuzz_Nosy_MockERC20CallerSession_TotalSupply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20CallerSession
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil {
			return
		}

		_MockERC20.TotalSupply()
	})
}

func Fuzz_Nosy_MockERC20Filterer_FilterApproval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Filterer
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var owner []common.Address
		fill_err = tp.Fill(&owner)
		if fill_err != nil {
			return
		}
		var spender []common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || opts == nil {
			return
		}

		_MockERC20.FilterApproval(opts, owner, spender)
	})
}

func Fuzz_Nosy_MockERC20Filterer_FilterTransfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Filterer
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var from []common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var to []common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || opts == nil {
			return
		}

		_MockERC20.FilterTransfer(opts, from, to)
	})
}

func Fuzz_Nosy_MockERC20Filterer_ParseApproval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Filterer
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil {
			return
		}

		_MockERC20.ParseApproval(log)
	})
}

func Fuzz_Nosy_MockERC20Filterer_ParseTransfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Filterer
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil {
			return
		}

		_MockERC20.ParseTransfer(log)
	})
}

// skipping Fuzz_Nosy_MockERC20Filterer_WatchApproval__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.MockERC20Approval

// skipping Fuzz_Nosy_MockERC20Filterer_WatchTransfer__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.MockERC20Transfer

// skipping Fuzz_Nosy_MockERC20Raw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_MockERC20Raw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_MockERC20Raw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Raw
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || opts == nil {
			return
		}

		_MockERC20.Transfer(opts)
	})
}

func Fuzz_Nosy_MockERC20Session_Allowance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Session
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var owner common.Address
		fill_err = tp.Fill(&owner)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil {
			return
		}

		_MockERC20.Allowance(owner, spender)
	})
}

func Fuzz_Nosy_MockERC20Session_Approve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Session
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || amount == nil {
			return
		}

		_MockERC20.Approve(spender, amount)
	})
}

func Fuzz_Nosy_MockERC20Session_BalanceOf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Session
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil {
			return
		}

		_MockERC20.BalanceOf(account)
	})
}

func Fuzz_Nosy_MockERC20Session_Decimals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Session
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil {
			return
		}

		_MockERC20.Decimals()
	})
}

func Fuzz_Nosy_MockERC20Session_DecreaseAllowance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Session
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		var subtractedValue *big.Int
		fill_err = tp.Fill(&subtractedValue)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || subtractedValue == nil {
			return
		}

		_MockERC20.DecreaseAllowance(spender, subtractedValue)
	})
}

func Fuzz_Nosy_MockERC20Session_IncreaseAllowance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Session
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		var addedValue *big.Int
		fill_err = tp.Fill(&addedValue)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || addedValue == nil {
			return
		}

		_MockERC20.IncreaseAllowance(spender, addedValue)
	})
}

func Fuzz_Nosy_MockERC20Session_Mint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Session
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || amount == nil {
			return
		}

		_MockERC20.Mint(to, amount)
	})
}

func Fuzz_Nosy_MockERC20Session_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Session
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil {
			return
		}

		_MockERC20.Name()
	})
}

func Fuzz_Nosy_MockERC20Session_Symbol__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Session
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil {
			return
		}

		_MockERC20.Symbol()
	})
}

func Fuzz_Nosy_MockERC20Session_TotalSupply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Session
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil {
			return
		}

		_MockERC20.TotalSupply()
	})
}

func Fuzz_Nosy_MockERC20Session_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Session
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || amount == nil {
			return
		}

		_MockERC20.Transfer(to, amount)
	})
}

func Fuzz_Nosy_MockERC20Session_TransferFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Session
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var from common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || amount == nil {
			return
		}

		_MockERC20.TransferFrom(from, to, amount)
	})
}

func Fuzz_Nosy_MockERC20Transactor_Approve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Transactor
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || opts == nil || amount == nil {
			return
		}

		_MockERC20.Approve(opts, spender, amount)
	})
}

func Fuzz_Nosy_MockERC20Transactor_DecreaseAllowance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Transactor
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		var subtractedValue *big.Int
		fill_err = tp.Fill(&subtractedValue)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || opts == nil || subtractedValue == nil {
			return
		}

		_MockERC20.DecreaseAllowance(opts, spender, subtractedValue)
	})
}

func Fuzz_Nosy_MockERC20Transactor_IncreaseAllowance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Transactor
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		var addedValue *big.Int
		fill_err = tp.Fill(&addedValue)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || opts == nil || addedValue == nil {
			return
		}

		_MockERC20.IncreaseAllowance(opts, spender, addedValue)
	})
}

func Fuzz_Nosy_MockERC20Transactor_Mint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Transactor
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || opts == nil || amount == nil {
			return
		}

		_MockERC20.Mint(opts, to, amount)
	})
}

func Fuzz_Nosy_MockERC20Transactor_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Transactor
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || opts == nil || amount == nil {
			return
		}

		_MockERC20.Transfer(opts, to, amount)
	})
}

func Fuzz_Nosy_MockERC20Transactor_TransferFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20Transactor
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var from common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || opts == nil || amount == nil {
			return
		}

		_MockERC20.TransferFrom(opts, from, to, amount)
	})
}

// skipping Fuzz_Nosy_MockERC20TransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_MockERC20TransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20TransactorRaw
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || opts == nil {
			return
		}

		_MockERC20.Transfer(opts)
	})
}

func Fuzz_Nosy_MockERC20TransactorSession_Approve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20TransactorSession
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || amount == nil {
			return
		}

		_MockERC20.Approve(spender, amount)
	})
}

func Fuzz_Nosy_MockERC20TransactorSession_DecreaseAllowance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20TransactorSession
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		var subtractedValue *big.Int
		fill_err = tp.Fill(&subtractedValue)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || subtractedValue == nil {
			return
		}

		_MockERC20.DecreaseAllowance(spender, subtractedValue)
	})
}

func Fuzz_Nosy_MockERC20TransactorSession_IncreaseAllowance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20TransactorSession
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		var addedValue *big.Int
		fill_err = tp.Fill(&addedValue)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || addedValue == nil {
			return
		}

		_MockERC20.IncreaseAllowance(spender, addedValue)
	})
}

func Fuzz_Nosy_MockERC20TransactorSession_Mint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20TransactorSession
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || amount == nil {
			return
		}

		_MockERC20.Mint(to, amount)
	})
}

func Fuzz_Nosy_MockERC20TransactorSession_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20TransactorSession
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || amount == nil {
			return
		}

		_MockERC20.Transfer(to, amount)
	})
}

func Fuzz_Nosy_MockERC20TransactorSession_TransferFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _MockERC20 *MockERC20TransactorSession
		fill_err = tp.Fill(&_MockERC20)
		if fill_err != nil {
			return
		}
		var from common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _MockERC20 == nil || amount == nil {
			return
		}

		_MockERC20.TransferFrom(from, to, amount)
	})
}

func Fuzz_Nosy_MockERC20TransferIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *MockERC20TransferIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_MockERC20TransferIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *MockERC20TransferIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_MockERC20TransferIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *MockERC20TransferIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniAVSAllowlistDisabledIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSAllowlistDisabledIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniAVSAllowlistDisabledIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSAllowlistDisabledIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniAVSAllowlistDisabledIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSAllowlistDisabledIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniAVSAllowlistEnabledIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSAllowlistEnabledIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniAVSAllowlistEnabledIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSAllowlistEnabledIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniAVSAllowlistEnabledIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSAllowlistEnabledIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniAVSCaller_AllowlistEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCaller
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.AllowlistEnabled(opts)
	})
}

func Fuzz_Nosy_OmniAVSCaller_AvsDirectory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCaller
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.AvsDirectory(opts)
	})
}

func Fuzz_Nosy_OmniAVSCaller_CanRegister__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCaller
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.CanRegister(opts, operator)
	})
}

func Fuzz_Nosy_OmniAVSCaller_EthStakeInbox__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCaller
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.EthStakeInbox(opts)
	})
}

func Fuzz_Nosy_OmniAVSCaller_FeeForSync__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCaller
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.FeeForSync(opts)
	})
}

func Fuzz_Nosy_OmniAVSCaller_GetOperatorRestakedStrategies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCaller
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.GetOperatorRestakedStrategies(opts, operator)
	})
}

func Fuzz_Nosy_OmniAVSCaller_GetRestakeableStrategies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCaller
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.GetRestakeableStrategies(opts)
	})
}

func Fuzz_Nosy_OmniAVSCaller_IsInAllowlist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCaller
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.IsInAllowlist(opts, operator)
	})
}

func Fuzz_Nosy_OmniAVSCaller_MaxOperatorCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCaller
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.MaxOperatorCount(opts)
	})
}

func Fuzz_Nosy_OmniAVSCaller_MinOperatorStake__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCaller
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.MinOperatorStake(opts)
	})
}

func Fuzz_Nosy_OmniAVSCaller_Omni__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCaller
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.Omni(opts)
	})
}

func Fuzz_Nosy_OmniAVSCaller_OmniChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCaller
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.OmniChainId(opts)
	})
}

func Fuzz_Nosy_OmniAVSCaller_Operators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCaller
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.Operators(opts)
	})
}

func Fuzz_Nosy_OmniAVSCaller_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCaller
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.Owner(opts)
	})
}

func Fuzz_Nosy_OmniAVSCaller_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCaller
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.Paused(opts)
	})
}

func Fuzz_Nosy_OmniAVSCaller_StrategyParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCaller
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.StrategyParams(opts)
	})
}

func Fuzz_Nosy_OmniAVSCaller_XcallBaseGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCaller
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.XcallBaseGasLimit(opts)
	})
}

func Fuzz_Nosy_OmniAVSCaller_XcallGasLimitPerOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCaller
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.XcallGasLimitPerOperator(opts)
	})
}

// skipping Fuzz_Nosy_OmniAVSCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_OmniAVSCallerSession_AllowlistEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCallerSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.AllowlistEnabled()
	})
}

func Fuzz_Nosy_OmniAVSCallerSession_AvsDirectory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCallerSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.AvsDirectory()
	})
}

func Fuzz_Nosy_OmniAVSCallerSession_CanRegister__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCallerSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.CanRegister(operator)
	})
}

func Fuzz_Nosy_OmniAVSCallerSession_EthStakeInbox__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCallerSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.EthStakeInbox()
	})
}

func Fuzz_Nosy_OmniAVSCallerSession_FeeForSync__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCallerSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.FeeForSync()
	})
}

func Fuzz_Nosy_OmniAVSCallerSession_GetOperatorRestakedStrategies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCallerSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.GetOperatorRestakedStrategies(operator)
	})
}

func Fuzz_Nosy_OmniAVSCallerSession_GetRestakeableStrategies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCallerSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.GetRestakeableStrategies()
	})
}

func Fuzz_Nosy_OmniAVSCallerSession_IsInAllowlist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCallerSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.IsInAllowlist(operator)
	})
}

func Fuzz_Nosy_OmniAVSCallerSession_MaxOperatorCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCallerSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.MaxOperatorCount()
	})
}

func Fuzz_Nosy_OmniAVSCallerSession_MinOperatorStake__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCallerSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.MinOperatorStake()
	})
}

func Fuzz_Nosy_OmniAVSCallerSession_Omni__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCallerSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.Omni()
	})
}

func Fuzz_Nosy_OmniAVSCallerSession_OmniChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCallerSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.OmniChainId()
	})
}

func Fuzz_Nosy_OmniAVSCallerSession_Operators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCallerSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.Operators()
	})
}

func Fuzz_Nosy_OmniAVSCallerSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCallerSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.Owner()
	})
}

func Fuzz_Nosy_OmniAVSCallerSession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCallerSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.Paused()
	})
}

func Fuzz_Nosy_OmniAVSCallerSession_StrategyParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCallerSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.StrategyParams()
	})
}

func Fuzz_Nosy_OmniAVSCallerSession_XcallBaseGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCallerSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.XcallBaseGasLimit()
	})
}

func Fuzz_Nosy_OmniAVSCallerSession_XcallGasLimitPerOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSCallerSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.XcallGasLimitPerOperator()
	})
}

func Fuzz_Nosy_OmniAVSEthStakeInboxSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSEthStakeInboxSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniAVSEthStakeInboxSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSEthStakeInboxSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniAVSEthStakeInboxSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSEthStakeInboxSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniAVSFilterer_FilterAllowlistDisabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.FilterAllowlistDisabled(opts)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_FilterAllowlistEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.FilterAllowlistEnabled(opts)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_FilterEthStakeInboxSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var inbox []common.Address
		fill_err = tp.Fill(&inbox)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.FilterEthStakeInboxSet(opts, inbox)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_FilterMaxOperatorCountSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.FilterMaxOperatorCountSet(opts)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_FilterMinOperatorStakeSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.FilterMinOperatorStakeSet(opts)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_FilterOmniChainIdSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var chainID []uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.FilterOmniChainIdSet(opts, chainID)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_FilterOmniPortalSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var portal []common.Address
		fill_err = tp.Fill(&portal)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.FilterOmniPortalSet(opts, portal)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_FilterOperatorAdded__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator []common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.FilterOperatorAdded(opts, operator)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_FilterOperatorAllowed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator []common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.FilterOperatorAllowed(opts, operator)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_FilterOperatorDisallowed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator []common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.FilterOperatorDisallowed(opts, operator)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_FilterOperatorRemoved__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator []common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.FilterOperatorRemoved(opts, operator)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_FilterOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var previousOwner []common.Address
		fill_err = tp.Fill(&previousOwner)
		if fill_err != nil {
			return
		}
		var newOwner []common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.FilterOwnershipTransferred(opts, previousOwner, newOwner)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_FilterPaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.FilterPaused(opts)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_FilterStrategyParamsSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.FilterStrategyParamsSet(opts)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_FilterUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.FilterUnpaused(opts)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_FilterXCallGasLimitsSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.FilterXCallGasLimitsSet(opts)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_ParseAllowlistDisabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.ParseAllowlistDisabled(log)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_ParseAllowlistEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.ParseAllowlistEnabled(log)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_ParseEthStakeInboxSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.ParseEthStakeInboxSet(log)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.ParseInitialized(log)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_ParseMaxOperatorCountSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.ParseMaxOperatorCountSet(log)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_ParseMinOperatorStakeSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.ParseMinOperatorStakeSet(log)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_ParseOmniChainIdSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.ParseOmniChainIdSet(log)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_ParseOmniPortalSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.ParseOmniPortalSet(log)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_ParseOperatorAdded__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.ParseOperatorAdded(log)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_ParseOperatorAllowed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.ParseOperatorAllowed(log)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_ParseOperatorDisallowed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.ParseOperatorDisallowed(log)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_ParseOperatorRemoved__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.ParseOperatorRemoved(log)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_ParseOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.ParseOwnershipTransferred(log)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_ParsePaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.ParsePaused(log)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_ParseStrategyParamsSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.ParseStrategyParamsSet(log)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_ParseUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.ParseUnpaused(log)
	})
}

func Fuzz_Nosy_OmniAVSFilterer_ParseXCallGasLimitsSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSFilterer
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.ParseXCallGasLimitsSet(log)
	})
}

// skipping Fuzz_Nosy_OmniAVSFilterer_WatchAllowlistDisabled__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniAVSAllowlistDisabled

// skipping Fuzz_Nosy_OmniAVSFilterer_WatchAllowlistEnabled__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniAVSAllowlistEnabled

// skipping Fuzz_Nosy_OmniAVSFilterer_WatchEthStakeInboxSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniAVSEthStakeInboxSet

// skipping Fuzz_Nosy_OmniAVSFilterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniAVSInitialized

// skipping Fuzz_Nosy_OmniAVSFilterer_WatchMaxOperatorCountSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniAVSMaxOperatorCountSet

// skipping Fuzz_Nosy_OmniAVSFilterer_WatchMinOperatorStakeSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniAVSMinOperatorStakeSet

// skipping Fuzz_Nosy_OmniAVSFilterer_WatchOmniChainIdSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniAVSOmniChainIdSet

// skipping Fuzz_Nosy_OmniAVSFilterer_WatchOmniPortalSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniAVSOmniPortalSet

// skipping Fuzz_Nosy_OmniAVSFilterer_WatchOperatorAdded__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniAVSOperatorAdded

// skipping Fuzz_Nosy_OmniAVSFilterer_WatchOperatorAllowed__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniAVSOperatorAllowed

// skipping Fuzz_Nosy_OmniAVSFilterer_WatchOperatorDisallowed__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniAVSOperatorDisallowed

// skipping Fuzz_Nosy_OmniAVSFilterer_WatchOperatorRemoved__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniAVSOperatorRemoved

// skipping Fuzz_Nosy_OmniAVSFilterer_WatchOwnershipTransferred__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniAVSOwnershipTransferred

// skipping Fuzz_Nosy_OmniAVSFilterer_WatchPaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniAVSPaused

// skipping Fuzz_Nosy_OmniAVSFilterer_WatchStrategyParamsSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniAVSStrategyParamsSet

// skipping Fuzz_Nosy_OmniAVSFilterer_WatchUnpaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniAVSUnpaused

// skipping Fuzz_Nosy_OmniAVSFilterer_WatchXCallGasLimitsSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniAVSXCallGasLimitsSet

func Fuzz_Nosy_OmniAVSInitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniAVSInitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniAVSInitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniAVSMaxOperatorCountSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSMaxOperatorCountSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniAVSMaxOperatorCountSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSMaxOperatorCountSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniAVSMaxOperatorCountSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSMaxOperatorCountSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniAVSMinOperatorStakeSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSMinOperatorStakeSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniAVSMinOperatorStakeSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSMinOperatorStakeSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniAVSMinOperatorStakeSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSMinOperatorStakeSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniAVSOmniChainIdSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSOmniChainIdSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniAVSOmniChainIdSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSOmniChainIdSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniAVSOmniChainIdSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSOmniChainIdSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniAVSOmniPortalSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSOmniPortalSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniAVSOmniPortalSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSOmniPortalSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniAVSOmniPortalSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSOmniPortalSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniAVSOperatorAddedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSOperatorAddedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniAVSOperatorAddedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSOperatorAddedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniAVSOperatorAddedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSOperatorAddedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniAVSOperatorAllowedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSOperatorAllowedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniAVSOperatorAllowedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSOperatorAllowedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniAVSOperatorAllowedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSOperatorAllowedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniAVSOperatorDisallowedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSOperatorDisallowedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniAVSOperatorDisallowedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSOperatorDisallowedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniAVSOperatorDisallowedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSOperatorDisallowedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniAVSOperatorRemovedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSOperatorRemovedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniAVSOperatorRemovedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSOperatorRemovedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniAVSOperatorRemovedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSOperatorRemovedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniAVSOwnershipTransferredIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniAVSOwnershipTransferredIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniAVSOwnershipTransferredIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniAVSPausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSPausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniAVSPausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSPausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniAVSPausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSPausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

// skipping Fuzz_Nosy_OmniAVSRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_OmniAVSRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_OmniAVSRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSRaw
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.Transfer(opts)
	})
}

func Fuzz_Nosy_OmniAVSSession_AddToAllowlist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.AddToAllowlist(operator)
	})
}

func Fuzz_Nosy_OmniAVSSession_AllowlistEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.AllowlistEnabled()
	})
}

func Fuzz_Nosy_OmniAVSSession_AvsDirectory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.AvsDirectory()
	})
}

func Fuzz_Nosy_OmniAVSSession_CanRegister__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.CanRegister(operator)
	})
}

func Fuzz_Nosy_OmniAVSSession_DisableAllowlist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.DisableAllowlist()
	})
}

func Fuzz_Nosy_OmniAVSSession_EjectOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.EjectOperator(operator)
	})
}

func Fuzz_Nosy_OmniAVSSession_EnableAllowlist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.EnableAllowlist()
	})
}

func Fuzz_Nosy_OmniAVSSession_EthStakeInbox__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.EthStakeInbox()
	})
}

func Fuzz_Nosy_OmniAVSSession_FeeForSync__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.FeeForSync()
	})
}

func Fuzz_Nosy_OmniAVSSession_GetOperatorRestakedStrategies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.GetOperatorRestakedStrategies(operator)
	})
}

func Fuzz_Nosy_OmniAVSSession_GetRestakeableStrategies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.GetRestakeableStrategies()
	})
}

func Fuzz_Nosy_OmniAVSSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var owner_ common.Address
		fill_err = tp.Fill(&owner_)
		if fill_err != nil {
			return
		}
		var omni_ common.Address
		fill_err = tp.Fill(&omni_)
		if fill_err != nil {
			return
		}
		var omniChainId_ uint64
		fill_err = tp.Fill(&omniChainId_)
		if fill_err != nil {
			return
		}
		var ethStakeInbox_ common.Address
		fill_err = tp.Fill(&ethStakeInbox_)
		if fill_err != nil {
			return
		}
		var minOperatorStake_ *big.Int
		fill_err = tp.Fill(&minOperatorStake_)
		if fill_err != nil {
			return
		}
		var maxOperatorCount_ uint32
		fill_err = tp.Fill(&maxOperatorCount_)
		if fill_err != nil {
			return
		}
		var strategyParams_ []IOmniAVSStrategyParam
		fill_err = tp.Fill(&strategyParams_)
		if fill_err != nil {
			return
		}
		var metadataURI_ string
		fill_err = tp.Fill(&metadataURI_)
		if fill_err != nil {
			return
		}
		var allowlistEnabled_ bool
		fill_err = tp.Fill(&allowlistEnabled_)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || minOperatorStake_ == nil {
			return
		}

		_OmniAVS.Initialize(owner_, omni_, omniChainId_, ethStakeInbox_, minOperatorStake_, maxOperatorCount_, strategyParams_, metadataURI_, allowlistEnabled_)
	})
}

func Fuzz_Nosy_OmniAVSSession_IsInAllowlist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.IsInAllowlist(operator)
	})
}

func Fuzz_Nosy_OmniAVSSession_MaxOperatorCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.MaxOperatorCount()
	})
}

func Fuzz_Nosy_OmniAVSSession_MinOperatorStake__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.MinOperatorStake()
	})
}

func Fuzz_Nosy_OmniAVSSession_Omni__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.Omni()
	})
}

func Fuzz_Nosy_OmniAVSSession_OmniChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.OmniChainId()
	})
}

func Fuzz_Nosy_OmniAVSSession_Operators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.Operators()
	})
}

func Fuzz_Nosy_OmniAVSSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.Owner()
	})
}

func Fuzz_Nosy_OmniAVSSession_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.Pause()
	})
}

func Fuzz_Nosy_OmniAVSSession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.Paused()
	})
}

func Fuzz_Nosy_OmniAVSSession_RegisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var pubkey []byte
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry
		fill_err = tp.Fill(&operatorSignature)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.RegisterOperator(pubkey, operatorSignature)
	})
}

func Fuzz_Nosy_OmniAVSSession_RemoveFromAllowlist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.RemoveFromAllowlist(operator)
	})
}

func Fuzz_Nosy_OmniAVSSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.RenounceOwnership()
	})
}

func Fuzz_Nosy_OmniAVSSession_SetEthStakeInbox__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var inbox common.Address
		fill_err = tp.Fill(&inbox)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.SetEthStakeInbox(inbox)
	})
}

func Fuzz_Nosy_OmniAVSSession_SetMaxOperatorCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var count uint32
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.SetMaxOperatorCount(count)
	})
}

func Fuzz_Nosy_OmniAVSSession_SetMetadataURI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var metadataURI string
		fill_err = tp.Fill(&metadataURI)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.SetMetadataURI(metadataURI)
	})
}

func Fuzz_Nosy_OmniAVSSession_SetMinOperatorStake__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var stake *big.Int
		fill_err = tp.Fill(&stake)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || stake == nil {
			return
		}

		_OmniAVS.SetMinOperatorStake(stake)
	})
}

func Fuzz_Nosy_OmniAVSSession_SetOmniChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var chainId uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.SetOmniChainId(chainId)
	})
}

func Fuzz_Nosy_OmniAVSSession_SetOmniPortal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var portal common.Address
		fill_err = tp.Fill(&portal)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.SetOmniPortal(portal)
	})
}

func Fuzz_Nosy_OmniAVSSession_SetStrategyParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var params []IOmniAVSStrategyParam
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.SetStrategyParams(params)
	})
}

func Fuzz_Nosy_OmniAVSSession_SetXCallGasLimits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var base uint64
		fill_err = tp.Fill(&base)
		if fill_err != nil {
			return
		}
		var perOperator uint64
		fill_err = tp.Fill(&perOperator)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.SetXCallGasLimits(base, perOperator)
	})
}

func Fuzz_Nosy_OmniAVSSession_StrategyParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.StrategyParams()
	})
}

func Fuzz_Nosy_OmniAVSSession_SyncWithOmni__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.SyncWithOmni()
	})
}

func Fuzz_Nosy_OmniAVSSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_OmniAVSSession_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.Unpause()
	})
}

func Fuzz_Nosy_OmniAVSSession_XcallBaseGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.XcallBaseGasLimit()
	})
}

func Fuzz_Nosy_OmniAVSSession_XcallGasLimitPerOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.XcallGasLimitPerOperator()
	})
}

func Fuzz_Nosy_OmniAVSStrategyParamsSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSStrategyParamsSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniAVSStrategyParamsSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSStrategyParamsSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniAVSStrategyParamsSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSStrategyParamsSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniAVSTransactor_AddToAllowlist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactor
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.AddToAllowlist(opts, operator)
	})
}

func Fuzz_Nosy_OmniAVSTransactor_DisableAllowlist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactor
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.DisableAllowlist(opts)
	})
}

func Fuzz_Nosy_OmniAVSTransactor_EjectOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactor
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.EjectOperator(opts, operator)
	})
}

func Fuzz_Nosy_OmniAVSTransactor_EnableAllowlist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactor
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.EnableAllowlist(opts)
	})
}

func Fuzz_Nosy_OmniAVSTransactor_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactor
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var owner_ common.Address
		fill_err = tp.Fill(&owner_)
		if fill_err != nil {
			return
		}
		var omni_ common.Address
		fill_err = tp.Fill(&omni_)
		if fill_err != nil {
			return
		}
		var omniChainId_ uint64
		fill_err = tp.Fill(&omniChainId_)
		if fill_err != nil {
			return
		}
		var ethStakeInbox_ common.Address
		fill_err = tp.Fill(&ethStakeInbox_)
		if fill_err != nil {
			return
		}
		var minOperatorStake_ *big.Int
		fill_err = tp.Fill(&minOperatorStake_)
		if fill_err != nil {
			return
		}
		var maxOperatorCount_ uint32
		fill_err = tp.Fill(&maxOperatorCount_)
		if fill_err != nil {
			return
		}
		var strategyParams_ []IOmniAVSStrategyParam
		fill_err = tp.Fill(&strategyParams_)
		if fill_err != nil {
			return
		}
		var metadataURI_ string
		fill_err = tp.Fill(&metadataURI_)
		if fill_err != nil {
			return
		}
		var allowlistEnabled_ bool
		fill_err = tp.Fill(&allowlistEnabled_)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil || minOperatorStake_ == nil {
			return
		}

		_OmniAVS.Initialize(opts, owner_, omni_, omniChainId_, ethStakeInbox_, minOperatorStake_, maxOperatorCount_, strategyParams_, metadataURI_, allowlistEnabled_)
	})
}

func Fuzz_Nosy_OmniAVSTransactor_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactor
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.Pause(opts)
	})
}

func Fuzz_Nosy_OmniAVSTransactor_RegisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactor
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var pubkey []byte
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry
		fill_err = tp.Fill(&operatorSignature)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.RegisterOperator(opts, pubkey, operatorSignature)
	})
}

func Fuzz_Nosy_OmniAVSTransactor_RemoveFromAllowlist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactor
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.RemoveFromAllowlist(opts, operator)
	})
}

func Fuzz_Nosy_OmniAVSTransactor_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactor
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.RenounceOwnership(opts)
	})
}

func Fuzz_Nosy_OmniAVSTransactor_SetEthStakeInbox__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactor
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var inbox common.Address
		fill_err = tp.Fill(&inbox)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.SetEthStakeInbox(opts, inbox)
	})
}

func Fuzz_Nosy_OmniAVSTransactor_SetMaxOperatorCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactor
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var count uint32
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.SetMaxOperatorCount(opts, count)
	})
}

func Fuzz_Nosy_OmniAVSTransactor_SetMetadataURI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactor
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var metadataURI string
		fill_err = tp.Fill(&metadataURI)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.SetMetadataURI(opts, metadataURI)
	})
}

func Fuzz_Nosy_OmniAVSTransactor_SetMinOperatorStake__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactor
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var stake *big.Int
		fill_err = tp.Fill(&stake)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil || stake == nil {
			return
		}

		_OmniAVS.SetMinOperatorStake(opts, stake)
	})
}

func Fuzz_Nosy_OmniAVSTransactor_SetOmniChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactor
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var chainId uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.SetOmniChainId(opts, chainId)
	})
}

func Fuzz_Nosy_OmniAVSTransactor_SetOmniPortal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactor
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var portal common.Address
		fill_err = tp.Fill(&portal)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.SetOmniPortal(opts, portal)
	})
}

func Fuzz_Nosy_OmniAVSTransactor_SetStrategyParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactor
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var params []IOmniAVSStrategyParam
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.SetStrategyParams(opts, params)
	})
}

func Fuzz_Nosy_OmniAVSTransactor_SetXCallGasLimits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactor
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var base uint64
		fill_err = tp.Fill(&base)
		if fill_err != nil {
			return
		}
		var perOperator uint64
		fill_err = tp.Fill(&perOperator)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.SetXCallGasLimits(opts, base, perOperator)
	})
}

func Fuzz_Nosy_OmniAVSTransactor_SyncWithOmni__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactor
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.SyncWithOmni(opts)
	})
}

func Fuzz_Nosy_OmniAVSTransactor_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactor
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.TransferOwnership(opts, newOwner)
	})
}

func Fuzz_Nosy_OmniAVSTransactor_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactor
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.Unpause(opts)
	})
}

// skipping Fuzz_Nosy_OmniAVSTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_OmniAVSTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactorRaw
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || opts == nil {
			return
		}

		_OmniAVS.Transfer(opts)
	})
}

func Fuzz_Nosy_OmniAVSTransactorSession_AddToAllowlist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactorSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.AddToAllowlist(operator)
	})
}

func Fuzz_Nosy_OmniAVSTransactorSession_DisableAllowlist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactorSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.DisableAllowlist()
	})
}

func Fuzz_Nosy_OmniAVSTransactorSession_EjectOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactorSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.EjectOperator(operator)
	})
}

func Fuzz_Nosy_OmniAVSTransactorSession_EnableAllowlist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactorSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.EnableAllowlist()
	})
}

func Fuzz_Nosy_OmniAVSTransactorSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactorSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var owner_ common.Address
		fill_err = tp.Fill(&owner_)
		if fill_err != nil {
			return
		}
		var omni_ common.Address
		fill_err = tp.Fill(&omni_)
		if fill_err != nil {
			return
		}
		var omniChainId_ uint64
		fill_err = tp.Fill(&omniChainId_)
		if fill_err != nil {
			return
		}
		var ethStakeInbox_ common.Address
		fill_err = tp.Fill(&ethStakeInbox_)
		if fill_err != nil {
			return
		}
		var minOperatorStake_ *big.Int
		fill_err = tp.Fill(&minOperatorStake_)
		if fill_err != nil {
			return
		}
		var maxOperatorCount_ uint32
		fill_err = tp.Fill(&maxOperatorCount_)
		if fill_err != nil {
			return
		}
		var strategyParams_ []IOmniAVSStrategyParam
		fill_err = tp.Fill(&strategyParams_)
		if fill_err != nil {
			return
		}
		var metadataURI_ string
		fill_err = tp.Fill(&metadataURI_)
		if fill_err != nil {
			return
		}
		var allowlistEnabled_ bool
		fill_err = tp.Fill(&allowlistEnabled_)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || minOperatorStake_ == nil {
			return
		}

		_OmniAVS.Initialize(owner_, omni_, omniChainId_, ethStakeInbox_, minOperatorStake_, maxOperatorCount_, strategyParams_, metadataURI_, allowlistEnabled_)
	})
}

func Fuzz_Nosy_OmniAVSTransactorSession_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactorSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.Pause()
	})
}

func Fuzz_Nosy_OmniAVSTransactorSession_RegisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactorSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var pubkey []byte
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry
		fill_err = tp.Fill(&operatorSignature)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.RegisterOperator(pubkey, operatorSignature)
	})
}

func Fuzz_Nosy_OmniAVSTransactorSession_RemoveFromAllowlist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactorSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.RemoveFromAllowlist(operator)
	})
}

func Fuzz_Nosy_OmniAVSTransactorSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactorSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.RenounceOwnership()
	})
}

func Fuzz_Nosy_OmniAVSTransactorSession_SetEthStakeInbox__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactorSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var inbox common.Address
		fill_err = tp.Fill(&inbox)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.SetEthStakeInbox(inbox)
	})
}

func Fuzz_Nosy_OmniAVSTransactorSession_SetMaxOperatorCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactorSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var count uint32
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.SetMaxOperatorCount(count)
	})
}

func Fuzz_Nosy_OmniAVSTransactorSession_SetMetadataURI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactorSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var metadataURI string
		fill_err = tp.Fill(&metadataURI)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.SetMetadataURI(metadataURI)
	})
}

func Fuzz_Nosy_OmniAVSTransactorSession_SetMinOperatorStake__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactorSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var stake *big.Int
		fill_err = tp.Fill(&stake)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil || stake == nil {
			return
		}

		_OmniAVS.SetMinOperatorStake(stake)
	})
}

func Fuzz_Nosy_OmniAVSTransactorSession_SetOmniChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactorSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var chainId uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.SetOmniChainId(chainId)
	})
}

func Fuzz_Nosy_OmniAVSTransactorSession_SetOmniPortal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactorSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var portal common.Address
		fill_err = tp.Fill(&portal)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.SetOmniPortal(portal)
	})
}

func Fuzz_Nosy_OmniAVSTransactorSession_SetStrategyParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactorSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var params []IOmniAVSStrategyParam
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.SetStrategyParams(params)
	})
}

func Fuzz_Nosy_OmniAVSTransactorSession_SetXCallGasLimits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactorSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var base uint64
		fill_err = tp.Fill(&base)
		if fill_err != nil {
			return
		}
		var perOperator uint64
		fill_err = tp.Fill(&perOperator)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.SetXCallGasLimits(base, perOperator)
	})
}

func Fuzz_Nosy_OmniAVSTransactorSession_SyncWithOmni__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactorSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.SyncWithOmni()
	})
}

func Fuzz_Nosy_OmniAVSTransactorSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactorSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_OmniAVSTransactorSession_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniAVS *OmniAVSTransactorSession
		fill_err = tp.Fill(&_OmniAVS)
		if fill_err != nil {
			return
		}
		if _OmniAVS == nil {
			return
		}

		_OmniAVS.Unpause()
	})
}

func Fuzz_Nosy_OmniAVSUnpausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniAVSUnpausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniAVSUnpausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniAVSXCallGasLimitsSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSXCallGasLimitsSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniAVSXCallGasLimitsSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSXCallGasLimitsSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniAVSXCallGasLimitsSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniAVSXCallGasLimitsSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniApprovalIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniApprovalIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniApprovalIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniApprovalIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniApprovalIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniApprovalIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniBridgeL1BridgeIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniBridgeL1BridgeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniBridgeL1BridgeIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniBridgeL1BridgeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniBridgeL1BridgeIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniBridgeL1BridgeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniBridgeL1Caller_BridgeFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Caller
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var payor common.Address
		fill_err = tp.Fill(&payor)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil || opts == nil || amount == nil {
			return
		}

		_OmniBridgeL1.BridgeFee(opts, payor, to, amount)
	})
}

func Fuzz_Nosy_OmniBridgeL1Caller_Omni__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Caller
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil || opts == nil {
			return
		}

		_OmniBridgeL1.Omni(opts)
	})
}

func Fuzz_Nosy_OmniBridgeL1Caller_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Caller
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil || opts == nil {
			return
		}

		_OmniBridgeL1.Owner(opts)
	})
}

func Fuzz_Nosy_OmniBridgeL1Caller_Token__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Caller
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil || opts == nil {
			return
		}

		_OmniBridgeL1.Token(opts)
	})
}

func Fuzz_Nosy_OmniBridgeL1Caller_TotalL1Supply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Caller
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil || opts == nil {
			return
		}

		_OmniBridgeL1.TotalL1Supply(opts)
	})
}

func Fuzz_Nosy_OmniBridgeL1Caller_XCALLWITHDRAWGASLIMIT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Caller
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil || opts == nil {
			return
		}

		_OmniBridgeL1.XCALLWITHDRAWGASLIMIT(opts)
	})
}

// skipping Fuzz_Nosy_OmniBridgeL1CallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_OmniBridgeL1CallerSession_BridgeFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1CallerSession
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var payor common.Address
		fill_err = tp.Fill(&payor)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil || amount == nil {
			return
		}

		_OmniBridgeL1.BridgeFee(payor, to, amount)
	})
}

func Fuzz_Nosy_OmniBridgeL1CallerSession_Omni__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1CallerSession
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil {
			return
		}

		_OmniBridgeL1.Omni()
	})
}

func Fuzz_Nosy_OmniBridgeL1CallerSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1CallerSession
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil {
			return
		}

		_OmniBridgeL1.Owner()
	})
}

func Fuzz_Nosy_OmniBridgeL1CallerSession_Token__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1CallerSession
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil {
			return
		}

		_OmniBridgeL1.Token()
	})
}

func Fuzz_Nosy_OmniBridgeL1CallerSession_TotalL1Supply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1CallerSession
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil {
			return
		}

		_OmniBridgeL1.TotalL1Supply()
	})
}

func Fuzz_Nosy_OmniBridgeL1CallerSession_XCALLWITHDRAWGASLIMIT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1CallerSession
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil {
			return
		}

		_OmniBridgeL1.XCALLWITHDRAWGASLIMIT()
	})
}

func Fuzz_Nosy_OmniBridgeL1Filterer_FilterBridge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Filterer
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var payor []common.Address
		fill_err = tp.Fill(&payor)
		if fill_err != nil {
			return
		}
		var to []common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil || opts == nil {
			return
		}

		_OmniBridgeL1.FilterBridge(opts, payor, to)
	})
}

func Fuzz_Nosy_OmniBridgeL1Filterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Filterer
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil || opts == nil {
			return
		}

		_OmniBridgeL1.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_OmniBridgeL1Filterer_FilterOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Filterer
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var previousOwner []common.Address
		fill_err = tp.Fill(&previousOwner)
		if fill_err != nil {
			return
		}
		var newOwner []common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil || opts == nil {
			return
		}

		_OmniBridgeL1.FilterOwnershipTransferred(opts, previousOwner, newOwner)
	})
}

func Fuzz_Nosy_OmniBridgeL1Filterer_FilterWithdraw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Filterer
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var to []common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil || opts == nil {
			return
		}

		_OmniBridgeL1.FilterWithdraw(opts, to)
	})
}

func Fuzz_Nosy_OmniBridgeL1Filterer_ParseBridge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Filterer
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil {
			return
		}

		_OmniBridgeL1.ParseBridge(log)
	})
}

func Fuzz_Nosy_OmniBridgeL1Filterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Filterer
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil {
			return
		}

		_OmniBridgeL1.ParseInitialized(log)
	})
}

func Fuzz_Nosy_OmniBridgeL1Filterer_ParseOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Filterer
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil {
			return
		}

		_OmniBridgeL1.ParseOwnershipTransferred(log)
	})
}

func Fuzz_Nosy_OmniBridgeL1Filterer_ParseWithdraw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Filterer
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil {
			return
		}

		_OmniBridgeL1.ParseWithdraw(log)
	})
}

// skipping Fuzz_Nosy_OmniBridgeL1Filterer_WatchBridge__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniBridgeL1Bridge

// skipping Fuzz_Nosy_OmniBridgeL1Filterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniBridgeL1Initialized

// skipping Fuzz_Nosy_OmniBridgeL1Filterer_WatchOwnershipTransferred__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniBridgeL1OwnershipTransferred

// skipping Fuzz_Nosy_OmniBridgeL1Filterer_WatchWithdraw__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniBridgeL1Withdraw

func Fuzz_Nosy_OmniBridgeL1InitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniBridgeL1InitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniBridgeL1InitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniBridgeL1InitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniBridgeL1InitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniBridgeL1InitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniBridgeL1OwnershipTransferredIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniBridgeL1OwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniBridgeL1OwnershipTransferredIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniBridgeL1OwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniBridgeL1OwnershipTransferredIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniBridgeL1OwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

// skipping Fuzz_Nosy_OmniBridgeL1Raw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_OmniBridgeL1Raw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_OmniBridgeL1Raw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Raw
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil || opts == nil {
			return
		}

		_OmniBridgeL1.Transfer(opts)
	})
}

func Fuzz_Nosy_OmniBridgeL1Session_Bridge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Session
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil || amount == nil {
			return
		}

		_OmniBridgeL1.Bridge(to, amount)
	})
}

func Fuzz_Nosy_OmniBridgeL1Session_BridgeFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Session
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var payor common.Address
		fill_err = tp.Fill(&payor)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil || amount == nil {
			return
		}

		_OmniBridgeL1.BridgeFee(payor, to, amount)
	})
}

func Fuzz_Nosy_OmniBridgeL1Session_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Session
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var owner_ common.Address
		fill_err = tp.Fill(&owner_)
		if fill_err != nil {
			return
		}
		var omni_ common.Address
		fill_err = tp.Fill(&omni_)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil {
			return
		}

		_OmniBridgeL1.Initialize(owner_, omni_)
	})
}

func Fuzz_Nosy_OmniBridgeL1Session_Omni__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Session
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil {
			return
		}

		_OmniBridgeL1.Omni()
	})
}

func Fuzz_Nosy_OmniBridgeL1Session_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Session
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil {
			return
		}

		_OmniBridgeL1.Owner()
	})
}

func Fuzz_Nosy_OmniBridgeL1Session_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Session
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil {
			return
		}

		_OmniBridgeL1.RenounceOwnership()
	})
}

func Fuzz_Nosy_OmniBridgeL1Session_Token__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Session
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil {
			return
		}

		_OmniBridgeL1.Token()
	})
}

func Fuzz_Nosy_OmniBridgeL1Session_TotalL1Supply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Session
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil {
			return
		}

		_OmniBridgeL1.TotalL1Supply()
	})
}

func Fuzz_Nosy_OmniBridgeL1Session_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Session
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil {
			return
		}

		_OmniBridgeL1.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_OmniBridgeL1Session_Withdraw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Session
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil || amount == nil {
			return
		}

		_OmniBridgeL1.Withdraw(to, amount)
	})
}

func Fuzz_Nosy_OmniBridgeL1Session_XCALLWITHDRAWGASLIMIT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Session
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil {
			return
		}

		_OmniBridgeL1.XCALLWITHDRAWGASLIMIT()
	})
}

func Fuzz_Nosy_OmniBridgeL1Transactor_Bridge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Transactor
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil || opts == nil || amount == nil {
			return
		}

		_OmniBridgeL1.Bridge(opts, to, amount)
	})
}

func Fuzz_Nosy_OmniBridgeL1Transactor_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Transactor
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var owner_ common.Address
		fill_err = tp.Fill(&owner_)
		if fill_err != nil {
			return
		}
		var omni_ common.Address
		fill_err = tp.Fill(&omni_)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil || opts == nil {
			return
		}

		_OmniBridgeL1.Initialize(opts, owner_, omni_)
	})
}

func Fuzz_Nosy_OmniBridgeL1Transactor_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Transactor
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil || opts == nil {
			return
		}

		_OmniBridgeL1.RenounceOwnership(opts)
	})
}

func Fuzz_Nosy_OmniBridgeL1Transactor_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Transactor
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil || opts == nil {
			return
		}

		_OmniBridgeL1.TransferOwnership(opts, newOwner)
	})
}

func Fuzz_Nosy_OmniBridgeL1Transactor_Withdraw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1Transactor
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil || opts == nil || amount == nil {
			return
		}

		_OmniBridgeL1.Withdraw(opts, to, amount)
	})
}

// skipping Fuzz_Nosy_OmniBridgeL1TransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_OmniBridgeL1TransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1TransactorRaw
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil || opts == nil {
			return
		}

		_OmniBridgeL1.Transfer(opts)
	})
}

func Fuzz_Nosy_OmniBridgeL1TransactorSession_Bridge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1TransactorSession
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil || amount == nil {
			return
		}

		_OmniBridgeL1.Bridge(to, amount)
	})
}

func Fuzz_Nosy_OmniBridgeL1TransactorSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1TransactorSession
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var owner_ common.Address
		fill_err = tp.Fill(&owner_)
		if fill_err != nil {
			return
		}
		var omni_ common.Address
		fill_err = tp.Fill(&omni_)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil {
			return
		}

		_OmniBridgeL1.Initialize(owner_, omni_)
	})
}

func Fuzz_Nosy_OmniBridgeL1TransactorSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1TransactorSession
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil {
			return
		}

		_OmniBridgeL1.RenounceOwnership()
	})
}

func Fuzz_Nosy_OmniBridgeL1TransactorSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1TransactorSession
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil {
			return
		}

		_OmniBridgeL1.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_OmniBridgeL1TransactorSession_Withdraw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeL1 *OmniBridgeL1TransactorSession
		fill_err = tp.Fill(&_OmniBridgeL1)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _OmniBridgeL1 == nil || amount == nil {
			return
		}

		_OmniBridgeL1.Withdraw(to, amount)
	})
}

func Fuzz_Nosy_OmniBridgeL1WithdrawIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniBridgeL1WithdrawIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniBridgeL1WithdrawIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniBridgeL1WithdrawIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniBridgeL1WithdrawIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniBridgeL1WithdrawIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniBridgeNativeBridgeIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniBridgeNativeBridgeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniBridgeNativeBridgeIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniBridgeNativeBridgeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniBridgeNativeBridgeIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniBridgeNativeBridgeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniBridgeNativeCaller_BridgeFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeCaller
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || opts == nil || amount == nil {
			return
		}

		_OmniBridgeNative.BridgeFee(opts, to, amount)
	})
}

func Fuzz_Nosy_OmniBridgeNativeCaller_Claimable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeCaller
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || opts == nil {
			return
		}

		_OmniBridgeNative.Claimable(opts, arg0)
	})
}

func Fuzz_Nosy_OmniBridgeNativeCaller_L1Bridge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeCaller
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || opts == nil {
			return
		}

		_OmniBridgeNative.L1Bridge(opts)
	})
}

func Fuzz_Nosy_OmniBridgeNativeCaller_L1BridgeBalance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeCaller
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || opts == nil {
			return
		}

		_OmniBridgeNative.L1BridgeBalance(opts)
	})
}

func Fuzz_Nosy_OmniBridgeNativeCaller_L1ChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeCaller
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || opts == nil {
			return
		}

		_OmniBridgeNative.L1ChainId(opts)
	})
}

func Fuzz_Nosy_OmniBridgeNativeCaller_Omni__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeCaller
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || opts == nil {
			return
		}

		_OmniBridgeNative.Omni(opts)
	})
}

func Fuzz_Nosy_OmniBridgeNativeCaller_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeCaller
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || opts == nil {
			return
		}

		_OmniBridgeNative.Owner(opts)
	})
}

func Fuzz_Nosy_OmniBridgeNativeCaller_TotalL1Supply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeCaller
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || opts == nil {
			return
		}

		_OmniBridgeNative.TotalL1Supply(opts)
	})
}

func Fuzz_Nosy_OmniBridgeNativeCaller_XCALLWITHDRAWGASLIMIT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeCaller
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || opts == nil {
			return
		}

		_OmniBridgeNative.XCALLWITHDRAWGASLIMIT(opts)
	})
}

// skipping Fuzz_Nosy_OmniBridgeNativeCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_OmniBridgeNativeCallerSession_BridgeFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeCallerSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || amount == nil {
			return
		}

		_OmniBridgeNative.BridgeFee(to, amount)
	})
}

func Fuzz_Nosy_OmniBridgeNativeCallerSession_Claimable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeCallerSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.Claimable(arg0)
	})
}

func Fuzz_Nosy_OmniBridgeNativeCallerSession_L1Bridge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeCallerSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.L1Bridge()
	})
}

func Fuzz_Nosy_OmniBridgeNativeCallerSession_L1BridgeBalance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeCallerSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.L1BridgeBalance()
	})
}

func Fuzz_Nosy_OmniBridgeNativeCallerSession_L1ChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeCallerSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.L1ChainId()
	})
}

func Fuzz_Nosy_OmniBridgeNativeCallerSession_Omni__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeCallerSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.Omni()
	})
}

func Fuzz_Nosy_OmniBridgeNativeCallerSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeCallerSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.Owner()
	})
}

func Fuzz_Nosy_OmniBridgeNativeCallerSession_TotalL1Supply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeCallerSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.TotalL1Supply()
	})
}

func Fuzz_Nosy_OmniBridgeNativeCallerSession_XCALLWITHDRAWGASLIMIT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeCallerSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.XCALLWITHDRAWGASLIMIT()
	})
}

func Fuzz_Nosy_OmniBridgeNativeClaimedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniBridgeNativeClaimedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniBridgeNativeClaimedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniBridgeNativeClaimedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniBridgeNativeClaimedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniBridgeNativeClaimedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniBridgeNativeFilterer_FilterBridge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeFilterer
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var payor []common.Address
		fill_err = tp.Fill(&payor)
		if fill_err != nil {
			return
		}
		var to []common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || opts == nil {
			return
		}

		_OmniBridgeNative.FilterBridge(opts, payor, to)
	})
}

func Fuzz_Nosy_OmniBridgeNativeFilterer_FilterClaimed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeFilterer
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var claimant []common.Address
		fill_err = tp.Fill(&claimant)
		if fill_err != nil {
			return
		}
		var to []common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || opts == nil {
			return
		}

		_OmniBridgeNative.FilterClaimed(opts, claimant, to)
	})
}

func Fuzz_Nosy_OmniBridgeNativeFilterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeFilterer
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || opts == nil {
			return
		}

		_OmniBridgeNative.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_OmniBridgeNativeFilterer_FilterOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeFilterer
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var previousOwner []common.Address
		fill_err = tp.Fill(&previousOwner)
		if fill_err != nil {
			return
		}
		var newOwner []common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || opts == nil {
			return
		}

		_OmniBridgeNative.FilterOwnershipTransferred(opts, previousOwner, newOwner)
	})
}

func Fuzz_Nosy_OmniBridgeNativeFilterer_FilterWithdraw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeFilterer
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var payor []common.Address
		fill_err = tp.Fill(&payor)
		if fill_err != nil {
			return
		}
		var to []common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || opts == nil {
			return
		}

		_OmniBridgeNative.FilterWithdraw(opts, payor, to)
	})
}

func Fuzz_Nosy_OmniBridgeNativeFilterer_ParseBridge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeFilterer
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.ParseBridge(log)
	})
}

func Fuzz_Nosy_OmniBridgeNativeFilterer_ParseClaimed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeFilterer
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.ParseClaimed(log)
	})
}

func Fuzz_Nosy_OmniBridgeNativeFilterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeFilterer
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.ParseInitialized(log)
	})
}

func Fuzz_Nosy_OmniBridgeNativeFilterer_ParseOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeFilterer
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.ParseOwnershipTransferred(log)
	})
}

func Fuzz_Nosy_OmniBridgeNativeFilterer_ParseWithdraw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeFilterer
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.ParseWithdraw(log)
	})
}

// skipping Fuzz_Nosy_OmniBridgeNativeFilterer_WatchBridge__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniBridgeNativeBridge

// skipping Fuzz_Nosy_OmniBridgeNativeFilterer_WatchClaimed__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniBridgeNativeClaimed

// skipping Fuzz_Nosy_OmniBridgeNativeFilterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniBridgeNativeInitialized

// skipping Fuzz_Nosy_OmniBridgeNativeFilterer_WatchOwnershipTransferred__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniBridgeNativeOwnershipTransferred

// skipping Fuzz_Nosy_OmniBridgeNativeFilterer_WatchWithdraw__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniBridgeNativeWithdraw

func Fuzz_Nosy_OmniBridgeNativeInitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniBridgeNativeInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniBridgeNativeInitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniBridgeNativeInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniBridgeNativeInitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniBridgeNativeInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniBridgeNativeOwnershipTransferredIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniBridgeNativeOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniBridgeNativeOwnershipTransferredIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniBridgeNativeOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniBridgeNativeOwnershipTransferredIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniBridgeNativeOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

// skipping Fuzz_Nosy_OmniBridgeNativeRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_OmniBridgeNativeRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_OmniBridgeNativeRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeRaw
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || opts == nil {
			return
		}

		_OmniBridgeNative.Transfer(opts)
	})
}

func Fuzz_Nosy_OmniBridgeNativeSession_Bridge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || amount == nil {
			return
		}

		_OmniBridgeNative.Bridge(to, amount)
	})
}

func Fuzz_Nosy_OmniBridgeNativeSession_BridgeFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || amount == nil {
			return
		}

		_OmniBridgeNative.BridgeFee(to, amount)
	})
}

func Fuzz_Nosy_OmniBridgeNativeSession_Claim__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.Claim(to)
	})
}

func Fuzz_Nosy_OmniBridgeNativeSession_Claimable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.Claimable(arg0)
	})
}

func Fuzz_Nosy_OmniBridgeNativeSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var owner_ common.Address
		fill_err = tp.Fill(&owner_)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.Initialize(owner_)
	})
}

func Fuzz_Nosy_OmniBridgeNativeSession_L1Bridge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.L1Bridge()
	})
}

func Fuzz_Nosy_OmniBridgeNativeSession_L1BridgeBalance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.L1BridgeBalance()
	})
}

func Fuzz_Nosy_OmniBridgeNativeSession_L1ChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.L1ChainId()
	})
}

func Fuzz_Nosy_OmniBridgeNativeSession_Omni__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.Omni()
	})
}

func Fuzz_Nosy_OmniBridgeNativeSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.Owner()
	})
}

func Fuzz_Nosy_OmniBridgeNativeSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.RenounceOwnership()
	})
}

func Fuzz_Nosy_OmniBridgeNativeSession_Setup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var l1ChainId_ uint64
		fill_err = tp.Fill(&l1ChainId_)
		if fill_err != nil {
			return
		}
		var omni_ common.Address
		fill_err = tp.Fill(&omni_)
		if fill_err != nil {
			return
		}
		var l1Bridge_ common.Address
		fill_err = tp.Fill(&l1Bridge_)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.Setup(l1ChainId_, omni_, l1Bridge_)
	})
}

func Fuzz_Nosy_OmniBridgeNativeSession_TotalL1Supply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.TotalL1Supply()
	})
}

func Fuzz_Nosy_OmniBridgeNativeSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_OmniBridgeNativeSession_Withdraw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var payor common.Address
		fill_err = tp.Fill(&payor)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || amount == nil {
			return
		}

		_OmniBridgeNative.Withdraw(payor, to, amount)
	})
}

func Fuzz_Nosy_OmniBridgeNativeSession_XCALLWITHDRAWGASLIMIT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.XCALLWITHDRAWGASLIMIT()
	})
}

func Fuzz_Nosy_OmniBridgeNativeTransactor_Bridge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeTransactor
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || opts == nil || amount == nil {
			return
		}

		_OmniBridgeNative.Bridge(opts, to, amount)
	})
}

func Fuzz_Nosy_OmniBridgeNativeTransactor_Claim__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeTransactor
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || opts == nil {
			return
		}

		_OmniBridgeNative.Claim(opts, to)
	})
}

func Fuzz_Nosy_OmniBridgeNativeTransactor_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeTransactor
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var owner_ common.Address
		fill_err = tp.Fill(&owner_)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || opts == nil {
			return
		}

		_OmniBridgeNative.Initialize(opts, owner_)
	})
}

func Fuzz_Nosy_OmniBridgeNativeTransactor_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeTransactor
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || opts == nil {
			return
		}

		_OmniBridgeNative.RenounceOwnership(opts)
	})
}

func Fuzz_Nosy_OmniBridgeNativeTransactor_Setup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeTransactor
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var l1ChainId_ uint64
		fill_err = tp.Fill(&l1ChainId_)
		if fill_err != nil {
			return
		}
		var omni_ common.Address
		fill_err = tp.Fill(&omni_)
		if fill_err != nil {
			return
		}
		var l1Bridge_ common.Address
		fill_err = tp.Fill(&l1Bridge_)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || opts == nil {
			return
		}

		_OmniBridgeNative.Setup(opts, l1ChainId_, omni_, l1Bridge_)
	})
}

func Fuzz_Nosy_OmniBridgeNativeTransactor_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeTransactor
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || opts == nil {
			return
		}

		_OmniBridgeNative.TransferOwnership(opts, newOwner)
	})
}

func Fuzz_Nosy_OmniBridgeNativeTransactor_Withdraw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeTransactor
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var payor common.Address
		fill_err = tp.Fill(&payor)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || opts == nil || amount == nil {
			return
		}

		_OmniBridgeNative.Withdraw(opts, payor, to, amount)
	})
}

// skipping Fuzz_Nosy_OmniBridgeNativeTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_OmniBridgeNativeTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeTransactorRaw
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || opts == nil {
			return
		}

		_OmniBridgeNative.Transfer(opts)
	})
}

func Fuzz_Nosy_OmniBridgeNativeTransactorSession_Bridge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeTransactorSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || amount == nil {
			return
		}

		_OmniBridgeNative.Bridge(to, amount)
	})
}

func Fuzz_Nosy_OmniBridgeNativeTransactorSession_Claim__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeTransactorSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.Claim(to)
	})
}

func Fuzz_Nosy_OmniBridgeNativeTransactorSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeTransactorSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var owner_ common.Address
		fill_err = tp.Fill(&owner_)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.Initialize(owner_)
	})
}

func Fuzz_Nosy_OmniBridgeNativeTransactorSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeTransactorSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.RenounceOwnership()
	})
}

func Fuzz_Nosy_OmniBridgeNativeTransactorSession_Setup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeTransactorSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var l1ChainId_ uint64
		fill_err = tp.Fill(&l1ChainId_)
		if fill_err != nil {
			return
		}
		var omni_ common.Address
		fill_err = tp.Fill(&omni_)
		if fill_err != nil {
			return
		}
		var l1Bridge_ common.Address
		fill_err = tp.Fill(&l1Bridge_)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.Setup(l1ChainId_, omni_, l1Bridge_)
	})
}

func Fuzz_Nosy_OmniBridgeNativeTransactorSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeTransactorSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil {
			return
		}

		_OmniBridgeNative.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_OmniBridgeNativeTransactorSession_Withdraw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniBridgeNative *OmniBridgeNativeTransactorSession
		fill_err = tp.Fill(&_OmniBridgeNative)
		if fill_err != nil {
			return
		}
		var payor common.Address
		fill_err = tp.Fill(&payor)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _OmniBridgeNative == nil || amount == nil {
			return
		}

		_OmniBridgeNative.Withdraw(payor, to, amount)
	})
}

func Fuzz_Nosy_OmniBridgeNativeWithdrawIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniBridgeNativeWithdrawIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniBridgeNativeWithdrawIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniBridgeNativeWithdrawIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniBridgeNativeWithdrawIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniBridgeNativeWithdrawIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniCaller_Allowance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniCaller
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var owner common.Address
		fill_err = tp.Fill(&owner)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		if _Omni == nil || opts == nil {
			return
		}

		_Omni.Allowance(opts, owner, spender)
	})
}

func Fuzz_Nosy_OmniCaller_BalanceOf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniCaller
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if _Omni == nil || opts == nil {
			return
		}

		_Omni.BalanceOf(opts, account)
	})
}

func Fuzz_Nosy_OmniCaller_DOMAINSEPARATOR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniCaller
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Omni == nil || opts == nil {
			return
		}

		_Omni.DOMAINSEPARATOR(opts)
	})
}

func Fuzz_Nosy_OmniCaller_Decimals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniCaller
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Omni == nil || opts == nil {
			return
		}

		_Omni.Decimals(opts)
	})
}

func Fuzz_Nosy_OmniCaller_Eip712Domain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniCaller
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Omni == nil || opts == nil {
			return
		}

		_Omni.Eip712Domain(opts)
	})
}

func Fuzz_Nosy_OmniCaller_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniCaller
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Omni == nil || opts == nil {
			return
		}

		_Omni.Name(opts)
	})
}

func Fuzz_Nosy_OmniCaller_Nonces__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniCaller
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var owner common.Address
		fill_err = tp.Fill(&owner)
		if fill_err != nil {
			return
		}
		if _Omni == nil || opts == nil {
			return
		}

		_Omni.Nonces(opts, owner)
	})
}

func Fuzz_Nosy_OmniCaller_Symbol__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniCaller
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Omni == nil || opts == nil {
			return
		}

		_Omni.Symbol(opts)
	})
}

func Fuzz_Nosy_OmniCaller_TotalSupply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniCaller
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Omni == nil || opts == nil {
			return
		}

		_Omni.TotalSupply(opts)
	})
}

// skipping Fuzz_Nosy_OmniCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_OmniCallerSession_Allowance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniCallerSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var owner common.Address
		fill_err = tp.Fill(&owner)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		if _Omni == nil {
			return
		}

		_Omni.Allowance(owner, spender)
	})
}

func Fuzz_Nosy_OmniCallerSession_BalanceOf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniCallerSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if _Omni == nil {
			return
		}

		_Omni.BalanceOf(account)
	})
}

func Fuzz_Nosy_OmniCallerSession_DOMAINSEPARATOR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniCallerSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		if _Omni == nil {
			return
		}

		_Omni.DOMAINSEPARATOR()
	})
}

func Fuzz_Nosy_OmniCallerSession_Decimals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniCallerSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		if _Omni == nil {
			return
		}

		_Omni.Decimals()
	})
}

func Fuzz_Nosy_OmniCallerSession_Eip712Domain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniCallerSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		if _Omni == nil {
			return
		}

		_Omni.Eip712Domain()
	})
}

func Fuzz_Nosy_OmniCallerSession_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniCallerSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		if _Omni == nil {
			return
		}

		_Omni.Name()
	})
}

func Fuzz_Nosy_OmniCallerSession_Nonces__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniCallerSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var owner common.Address
		fill_err = tp.Fill(&owner)
		if fill_err != nil {
			return
		}
		if _Omni == nil {
			return
		}

		_Omni.Nonces(owner)
	})
}

func Fuzz_Nosy_OmniCallerSession_Symbol__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniCallerSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		if _Omni == nil {
			return
		}

		_Omni.Symbol()
	})
}

func Fuzz_Nosy_OmniCallerSession_TotalSupply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniCallerSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		if _Omni == nil {
			return
		}

		_Omni.TotalSupply()
	})
}

func Fuzz_Nosy_OmniEIP712DomainChangedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniEIP712DomainChangedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniEIP712DomainChangedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniEIP712DomainChangedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniEIP712DomainChangedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniEIP712DomainChangedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniFilterer_FilterApproval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniFilterer
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var owner []common.Address
		fill_err = tp.Fill(&owner)
		if fill_err != nil {
			return
		}
		var spender []common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		if _Omni == nil || opts == nil {
			return
		}

		_Omni.FilterApproval(opts, owner, spender)
	})
}

func Fuzz_Nosy_OmniFilterer_FilterEIP712DomainChanged__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniFilterer
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Omni == nil || opts == nil {
			return
		}

		_Omni.FilterEIP712DomainChanged(opts)
	})
}

func Fuzz_Nosy_OmniFilterer_FilterTransfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniFilterer
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var from []common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var to []common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		if _Omni == nil || opts == nil {
			return
		}

		_Omni.FilterTransfer(opts, from, to)
	})
}

func Fuzz_Nosy_OmniFilterer_ParseApproval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniFilterer
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _Omni == nil {
			return
		}

		_Omni.ParseApproval(log)
	})
}

func Fuzz_Nosy_OmniFilterer_ParseEIP712DomainChanged__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniFilterer
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _Omni == nil {
			return
		}

		_Omni.ParseEIP712DomainChanged(log)
	})
}

func Fuzz_Nosy_OmniFilterer_ParseTransfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniFilterer
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _Omni == nil {
			return
		}

		_Omni.ParseTransfer(log)
	})
}

// skipping Fuzz_Nosy_OmniFilterer_WatchApproval__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniApproval

// skipping Fuzz_Nosy_OmniFilterer_WatchEIP712DomainChanged__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniEIP712DomainChanged

// skipping Fuzz_Nosy_OmniFilterer_WatchTransfer__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniTransfer

func Fuzz_Nosy_OmniPortalCaller_ActionXCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.ActionXCall(opts)
	})
}

func Fuzz_Nosy_OmniPortalCaller_ActionXSubmit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.ActionXSubmit(opts)
	})
}

func Fuzz_Nosy_OmniPortalCaller_ChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.ChainId(opts)
	})
}

func Fuzz_Nosy_OmniPortalCaller_FeeFor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var destChainId uint64
		fill_err = tp.Fill(&destChainId)
		if fill_err != nil {
			return
		}
		var d4 []byte
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.FeeFor(opts, destChainId, d4, gasLimit)
	})
}

func Fuzz_Nosy_OmniPortalCaller_FeeOracle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.FeeOracle(opts)
	})
}

func Fuzz_Nosy_OmniPortalCaller_InXBlockOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 uint64
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 uint64
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.InXBlockOffset(opts, arg0, arg1)
	})
}

func Fuzz_Nosy_OmniPortalCaller_InXMsgOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 uint64
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 uint64
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.InXMsgOffset(opts, arg0, arg1)
	})
}

func Fuzz_Nosy_OmniPortalCaller_IsPaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var actionId [32]byte
		fill_err = tp.Fill(&actionId)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.IsPaused(opts, actionId)
	})
}

func Fuzz_Nosy_OmniPortalCaller_IsPaused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var actionId [32]byte
		fill_err = tp.Fill(&actionId)
		if fill_err != nil {
			return
		}
		var chainId_ uint64
		fill_err = tp.Fill(&chainId_)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.IsPaused0(opts, actionId, chainId_)
	})
}

func Fuzz_Nosy_OmniPortalCaller_IsPaused1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.IsPaused1(opts)
	})
}

func Fuzz_Nosy_OmniPortalCaller_IsSupportedDest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 uint64
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.IsSupportedDest(opts, arg0)
	})
}

func Fuzz_Nosy_OmniPortalCaller_IsSupportedShard__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 uint64
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.IsSupportedShard(opts, arg0)
	})
}

func Fuzz_Nosy_OmniPortalCaller_IsXCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.IsXCall(opts)
	})
}

func Fuzz_Nosy_OmniPortalCaller_KeyPauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.KeyPauseAll(opts)
	})
}

func Fuzz_Nosy_OmniPortalCaller_LatestValSetId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.LatestValSetId(opts)
	})
}

func Fuzz_Nosy_OmniPortalCaller_Network__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 *big.Int
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil || arg0 == nil {
			return
		}

		_OmniPortal.Network(opts, arg0)
	})
}

func Fuzz_Nosy_OmniPortalCaller_OmniCChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.OmniCChainId(opts)
	})
}

func Fuzz_Nosy_OmniPortalCaller_OmniChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.OmniChainId(opts)
	})
}

func Fuzz_Nosy_OmniPortalCaller_OutXMsgOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 uint64
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 uint64
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.OutXMsgOffset(opts, arg0, arg1)
	})
}

func Fuzz_Nosy_OmniPortalCaller_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.Owner(opts)
	})
}

func Fuzz_Nosy_OmniPortalCaller_ValSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 uint64
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 common.Address
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.ValSet(opts, arg0, arg1)
	})
}

func Fuzz_Nosy_OmniPortalCaller_ValSetTotalPower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 uint64
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.ValSetTotalPower(opts, arg0)
	})
}

func Fuzz_Nosy_OmniPortalCaller_XSubQuorumDenominator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.XSubQuorumDenominator(opts)
	})
}

func Fuzz_Nosy_OmniPortalCaller_XSubQuorumNumerator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.XSubQuorumNumerator(opts)
	})
}

func Fuzz_Nosy_OmniPortalCaller_Xmsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.Xmsg(opts)
	})
}

func Fuzz_Nosy_OmniPortalCaller_XmsgMaxDataSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.XmsgMaxDataSize(opts)
	})
}

func Fuzz_Nosy_OmniPortalCaller_XmsgMaxGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.XmsgMaxGasLimit(opts)
	})
}

func Fuzz_Nosy_OmniPortalCaller_XmsgMinGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.XmsgMinGasLimit(opts)
	})
}

func Fuzz_Nosy_OmniPortalCaller_XreceiptMaxErrorSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.XreceiptMaxErrorSize(opts)
	})
}

func Fuzz_Nosy_OmniPortalCaller_XsubValsetCutoff__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCaller
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.XsubValsetCutoff(opts)
	})
}

// skipping Fuzz_Nosy_OmniPortalCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_OmniPortalCallerSession_ActionXCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ActionXCall()
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_ActionXSubmit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ActionXSubmit()
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_ChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ChainId()
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_FeeFor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var destChainId uint64
		fill_err = tp.Fill(&destChainId)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.FeeFor(destChainId, d3, gasLimit)
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_FeeOracle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.FeeOracle()
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_InXBlockOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var arg0 uint64
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 uint64
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.InXBlockOffset(arg0, arg1)
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_InXMsgOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var arg0 uint64
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 uint64
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.InXMsgOffset(arg0, arg1)
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_IsPaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var actionId [32]byte
		fill_err = tp.Fill(&actionId)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.IsPaused(actionId)
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_IsPaused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var actionId [32]byte
		fill_err = tp.Fill(&actionId)
		if fill_err != nil {
			return
		}
		var chainId_ uint64
		fill_err = tp.Fill(&chainId_)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.IsPaused0(actionId, chainId_)
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_IsPaused1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.IsPaused1()
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_IsSupportedDest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var arg0 uint64
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.IsSupportedDest(arg0)
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_IsSupportedShard__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var arg0 uint64
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.IsSupportedShard(arg0)
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_IsXCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.IsXCall()
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_KeyPauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.KeyPauseAll()
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_LatestValSetId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.LatestValSetId()
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_Network__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var arg0 *big.Int
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || arg0 == nil {
			return
		}

		_OmniPortal.Network(arg0)
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_OmniCChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.OmniCChainId()
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_OmniChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.OmniChainId()
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_OutXMsgOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var arg0 uint64
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 uint64
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.OutXMsgOffset(arg0, arg1)
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.Owner()
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_ValSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var arg0 uint64
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 common.Address
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ValSet(arg0, arg1)
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_ValSetTotalPower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var arg0 uint64
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ValSetTotalPower(arg0)
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_XSubQuorumDenominator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.XSubQuorumDenominator()
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_XSubQuorumNumerator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.XSubQuorumNumerator()
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_Xmsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.Xmsg()
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_XmsgMaxDataSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.XmsgMaxDataSize()
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_XmsgMaxGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.XmsgMaxGasLimit()
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_XmsgMinGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.XmsgMinGasLimit()
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_XreceiptMaxErrorSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.XreceiptMaxErrorSize()
	})
}

func Fuzz_Nosy_OmniPortalCallerSession_XsubValsetCutoff__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalCallerSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.XsubValsetCutoff()
	})
}

func Fuzz_Nosy_OmniPortalFeeOracleSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalFeeOracleSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniPortalFeeOracleSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalFeeOracleSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniPortalFeeOracleSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalFeeOracleSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniPortalFeesCollectedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalFeesCollectedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniPortalFeesCollectedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalFeesCollectedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniPortalFeesCollectedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalFeesCollectedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniPortalFilterer_FilterFeeOracleSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.FilterFeeOracleSet(opts)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_FilterFeesCollected__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var to []common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.FilterFeesCollected(opts, to)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_FilterInXBlockOffsetSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var srcChainId []uint64
		fill_err = tp.Fill(&srcChainId)
		if fill_err != nil {
			return
		}
		var shardId []uint64
		fill_err = tp.Fill(&shardId)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.FilterInXBlockOffsetSet(opts, srcChainId, shardId)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_FilterInXMsgOffsetSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var srcChainId []uint64
		fill_err = tp.Fill(&srcChainId)
		if fill_err != nil {
			return
		}
		var shardId []uint64
		fill_err = tp.Fill(&shardId)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.FilterInXMsgOffsetSet(opts, srcChainId, shardId)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_FilterOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var previousOwner []common.Address
		fill_err = tp.Fill(&previousOwner)
		if fill_err != nil {
			return
		}
		var newOwner []common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.FilterOwnershipTransferred(opts, previousOwner, newOwner)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_FilterPaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.FilterPaused(opts)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_FilterUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.FilterUnpaused(opts)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_FilterValidatorSetAdded__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var setId []uint64
		fill_err = tp.Fill(&setId)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.FilterValidatorSetAdded(opts, setId)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_FilterXCallPaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.FilterXCallPaused(opts)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_FilterXCallToPaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var chainId []uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.FilterXCallToPaused(opts, chainId)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_FilterXCallToUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var chainId []uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.FilterXCallToUnpaused(opts, chainId)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_FilterXCallUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.FilterXCallUnpaused(opts)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_FilterXMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var destChainId []uint64
		fill_err = tp.Fill(&destChainId)
		if fill_err != nil {
			return
		}
		var shardId []uint64
		fill_err = tp.Fill(&shardId)
		if fill_err != nil {
			return
		}
		var offset []uint64
		fill_err = tp.Fill(&offset)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.FilterXMsg(opts, destChainId, shardId, offset)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_FilterXMsgMaxDataSizeSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.FilterXMsgMaxDataSizeSet(opts)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_FilterXMsgMaxGasLimitSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.FilterXMsgMaxGasLimitSet(opts)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_FilterXMsgMinGasLimitSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.FilterXMsgMinGasLimitSet(opts)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_FilterXReceipt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var sourceChainId []uint64
		fill_err = tp.Fill(&sourceChainId)
		if fill_err != nil {
			return
		}
		var shardId []uint64
		fill_err = tp.Fill(&shardId)
		if fill_err != nil {
			return
		}
		var offset []uint64
		fill_err = tp.Fill(&offset)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.FilterXReceipt(opts, sourceChainId, shardId, offset)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_FilterXReceiptMaxErrorSizeSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.FilterXReceiptMaxErrorSizeSet(opts)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_FilterXSubValsetCutoffSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.FilterXSubValsetCutoffSet(opts)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_FilterXSubmitFromPaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var chainId []uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.FilterXSubmitFromPaused(opts, chainId)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_FilterXSubmitFromUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var chainId []uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.FilterXSubmitFromUnpaused(opts, chainId)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_FilterXSubmitPaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.FilterXSubmitPaused(opts)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_FilterXSubmitUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.FilterXSubmitUnpaused(opts)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_ParseFeeOracleSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ParseFeeOracleSet(log)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_ParseFeesCollected__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ParseFeesCollected(log)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_ParseInXBlockOffsetSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ParseInXBlockOffsetSet(log)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_ParseInXMsgOffsetSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ParseInXMsgOffsetSet(log)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ParseInitialized(log)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_ParseOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ParseOwnershipTransferred(log)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_ParsePaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ParsePaused(log)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_ParseUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ParseUnpaused(log)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_ParseValidatorSetAdded__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ParseValidatorSetAdded(log)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_ParseXCallPaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ParseXCallPaused(log)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_ParseXCallToPaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ParseXCallToPaused(log)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_ParseXCallToUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ParseXCallToUnpaused(log)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_ParseXCallUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ParseXCallUnpaused(log)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_ParseXMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ParseXMsg(log)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_ParseXMsgMaxDataSizeSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ParseXMsgMaxDataSizeSet(log)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_ParseXMsgMaxGasLimitSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ParseXMsgMaxGasLimitSet(log)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_ParseXMsgMinGasLimitSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ParseXMsgMinGasLimitSet(log)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_ParseXReceipt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ParseXReceipt(log)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_ParseXReceiptMaxErrorSizeSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ParseXReceiptMaxErrorSizeSet(log)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_ParseXSubValsetCutoffSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ParseXSubValsetCutoffSet(log)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_ParseXSubmitFromPaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ParseXSubmitFromPaused(log)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_ParseXSubmitFromUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ParseXSubmitFromUnpaused(log)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_ParseXSubmitPaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ParseXSubmitPaused(log)
	})
}

func Fuzz_Nosy_OmniPortalFilterer_ParseXSubmitUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalFilterer
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ParseXSubmitUnpaused(log)
	})
}

// skipping Fuzz_Nosy_OmniPortalFilterer_WatchFeeOracleSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniPortalFeeOracleSet

// skipping Fuzz_Nosy_OmniPortalFilterer_WatchFeesCollected__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniPortalFeesCollected

// skipping Fuzz_Nosy_OmniPortalFilterer_WatchInXBlockOffsetSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniPortalInXBlockOffsetSet

// skipping Fuzz_Nosy_OmniPortalFilterer_WatchInXMsgOffsetSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniPortalInXMsgOffsetSet

// skipping Fuzz_Nosy_OmniPortalFilterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniPortalInitialized

// skipping Fuzz_Nosy_OmniPortalFilterer_WatchOwnershipTransferred__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniPortalOwnershipTransferred

// skipping Fuzz_Nosy_OmniPortalFilterer_WatchPaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniPortalPaused

// skipping Fuzz_Nosy_OmniPortalFilterer_WatchUnpaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniPortalUnpaused

// skipping Fuzz_Nosy_OmniPortalFilterer_WatchValidatorSetAdded__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniPortalValidatorSetAdded

// skipping Fuzz_Nosy_OmniPortalFilterer_WatchXCallPaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniPortalXCallPaused

// skipping Fuzz_Nosy_OmniPortalFilterer_WatchXCallToPaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniPortalXCallToPaused

// skipping Fuzz_Nosy_OmniPortalFilterer_WatchXCallToUnpaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniPortalXCallToUnpaused

// skipping Fuzz_Nosy_OmniPortalFilterer_WatchXCallUnpaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniPortalXCallUnpaused

// skipping Fuzz_Nosy_OmniPortalFilterer_WatchXMsg__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniPortalXMsg

// skipping Fuzz_Nosy_OmniPortalFilterer_WatchXMsgMaxDataSizeSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniPortalXMsgMaxDataSizeSet

// skipping Fuzz_Nosy_OmniPortalFilterer_WatchXMsgMaxGasLimitSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniPortalXMsgMaxGasLimitSet

// skipping Fuzz_Nosy_OmniPortalFilterer_WatchXMsgMinGasLimitSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniPortalXMsgMinGasLimitSet

// skipping Fuzz_Nosy_OmniPortalFilterer_WatchXReceipt__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniPortalXReceipt

// skipping Fuzz_Nosy_OmniPortalFilterer_WatchXReceiptMaxErrorSizeSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniPortalXReceiptMaxErrorSizeSet

// skipping Fuzz_Nosy_OmniPortalFilterer_WatchXSubValsetCutoffSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniPortalXSubValsetCutoffSet

// skipping Fuzz_Nosy_OmniPortalFilterer_WatchXSubmitFromPaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniPortalXSubmitFromPaused

// skipping Fuzz_Nosy_OmniPortalFilterer_WatchXSubmitFromUnpaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniPortalXSubmitFromUnpaused

// skipping Fuzz_Nosy_OmniPortalFilterer_WatchXSubmitPaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniPortalXSubmitPaused

// skipping Fuzz_Nosy_OmniPortalFilterer_WatchXSubmitUnpaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.OmniPortalXSubmitUnpaused

func Fuzz_Nosy_OmniPortalInXBlockOffsetSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalInXBlockOffsetSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniPortalInXBlockOffsetSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalInXBlockOffsetSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniPortalInXBlockOffsetSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalInXBlockOffsetSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniPortalInXMsgOffsetSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalInXMsgOffsetSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniPortalInXMsgOffsetSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalInXMsgOffsetSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniPortalInXMsgOffsetSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalInXMsgOffsetSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniPortalInitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniPortalInitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniPortalInitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniPortalOwnershipTransferredIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniPortalOwnershipTransferredIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniPortalOwnershipTransferredIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniPortalPausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalPausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniPortalPausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalPausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniPortalPausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalPausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

// skipping Fuzz_Nosy_OmniPortalRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_OmniPortalRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_OmniPortalRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalRaw
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.Transfer(opts)
	})
}

func Fuzz_Nosy_OmniPortalSession_ActionXCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ActionXCall()
	})
}

func Fuzz_Nosy_OmniPortalSession_ActionXSubmit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ActionXSubmit()
	})
}

func Fuzz_Nosy_OmniPortalSession_AddValidatorSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var valSetId uint64
		fill_err = tp.Fill(&valSetId)
		if fill_err != nil {
			return
		}
		var validators []XTypesValidator
		fill_err = tp.Fill(&validators)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.AddValidatorSet(valSetId, validators)
	})
}

func Fuzz_Nosy_OmniPortalSession_ChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ChainId()
	})
}

func Fuzz_Nosy_OmniPortalSession_CollectFees__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.CollectFees(to)
	})
}

func Fuzz_Nosy_OmniPortalSession_FeeFor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var destChainId uint64
		fill_err = tp.Fill(&destChainId)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.FeeFor(destChainId, d3, gasLimit)
	})
}

func Fuzz_Nosy_OmniPortalSession_FeeOracle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.FeeOracle()
	})
}

func Fuzz_Nosy_OmniPortalSession_InXBlockOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var arg0 uint64
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 uint64
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.InXBlockOffset(arg0, arg1)
	})
}

func Fuzz_Nosy_OmniPortalSession_InXMsgOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var arg0 uint64
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 uint64
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.InXMsgOffset(arg0, arg1)
	})
}

func Fuzz_Nosy_OmniPortalSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var p OmniPortalInitParams
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.Initialize(p)
	})
}

func Fuzz_Nosy_OmniPortalSession_IsPaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var actionId [32]byte
		fill_err = tp.Fill(&actionId)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.IsPaused(actionId)
	})
}

func Fuzz_Nosy_OmniPortalSession_IsPaused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var actionId [32]byte
		fill_err = tp.Fill(&actionId)
		if fill_err != nil {
			return
		}
		var chainId_ uint64
		fill_err = tp.Fill(&chainId_)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.IsPaused0(actionId, chainId_)
	})
}

func Fuzz_Nosy_OmniPortalSession_IsPaused1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.IsPaused1()
	})
}

func Fuzz_Nosy_OmniPortalSession_IsSupportedDest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var arg0 uint64
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.IsSupportedDest(arg0)
	})
}

func Fuzz_Nosy_OmniPortalSession_IsSupportedShard__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var arg0 uint64
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.IsSupportedShard(arg0)
	})
}

func Fuzz_Nosy_OmniPortalSession_IsXCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.IsXCall()
	})
}

func Fuzz_Nosy_OmniPortalSession_KeyPauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.KeyPauseAll()
	})
}

func Fuzz_Nosy_OmniPortalSession_LatestValSetId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.LatestValSetId()
	})
}

func Fuzz_Nosy_OmniPortalSession_Network__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var arg0 *big.Int
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || arg0 == nil {
			return
		}

		_OmniPortal.Network(arg0)
	})
}

func Fuzz_Nosy_OmniPortalSession_OmniCChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.OmniCChainId()
	})
}

func Fuzz_Nosy_OmniPortalSession_OmniChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.OmniChainId()
	})
}

func Fuzz_Nosy_OmniPortalSession_OutXMsgOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var arg0 uint64
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 uint64
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.OutXMsgOffset(arg0, arg1)
	})
}

func Fuzz_Nosy_OmniPortalSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.Owner()
	})
}

func Fuzz_Nosy_OmniPortalSession_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.Pause()
	})
}

func Fuzz_Nosy_OmniPortalSession_PauseXCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.PauseXCall()
	})
}

func Fuzz_Nosy_OmniPortalSession_PauseXCallTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var chainId_ uint64
		fill_err = tp.Fill(&chainId_)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.PauseXCallTo(chainId_)
	})
}

func Fuzz_Nosy_OmniPortalSession_PauseXSubmit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.PauseXSubmit()
	})
}

func Fuzz_Nosy_OmniPortalSession_PauseXSubmitFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var chainId_ uint64
		fill_err = tp.Fill(&chainId_)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.PauseXSubmitFrom(chainId_)
	})
}

func Fuzz_Nosy_OmniPortalSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.RenounceOwnership()
	})
}

func Fuzz_Nosy_OmniPortalSession_SetFeeOracle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var feeOracle_ common.Address
		fill_err = tp.Fill(&feeOracle_)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.SetFeeOracle(feeOracle_)
	})
}

func Fuzz_Nosy_OmniPortalSession_SetInXBlockOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var sourceChainId uint64
		fill_err = tp.Fill(&sourceChainId)
		if fill_err != nil {
			return
		}
		var shardId uint64
		fill_err = tp.Fill(&shardId)
		if fill_err != nil {
			return
		}
		var offset uint64
		fill_err = tp.Fill(&offset)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.SetInXBlockOffset(sourceChainId, shardId, offset)
	})
}

func Fuzz_Nosy_OmniPortalSession_SetInXMsgOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var sourceChainId uint64
		fill_err = tp.Fill(&sourceChainId)
		if fill_err != nil {
			return
		}
		var shardId uint64
		fill_err = tp.Fill(&shardId)
		if fill_err != nil {
			return
		}
		var offset uint64
		fill_err = tp.Fill(&offset)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.SetInXMsgOffset(sourceChainId, shardId, offset)
	})
}

func Fuzz_Nosy_OmniPortalSession_SetNetwork__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var network_ []XTypesChain
		fill_err = tp.Fill(&network_)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.SetNetwork(network_)
	})
}

func Fuzz_Nosy_OmniPortalSession_SetXMsgMaxDataSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var numBytes uint16
		fill_err = tp.Fill(&numBytes)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.SetXMsgMaxDataSize(numBytes)
	})
}

func Fuzz_Nosy_OmniPortalSession_SetXMsgMaxGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.SetXMsgMaxGasLimit(gasLimit)
	})
}

func Fuzz_Nosy_OmniPortalSession_SetXMsgMinGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.SetXMsgMinGasLimit(gasLimit)
	})
}

func Fuzz_Nosy_OmniPortalSession_SetXReceiptMaxErrorSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var numBytes uint16
		fill_err = tp.Fill(&numBytes)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.SetXReceiptMaxErrorSize(numBytes)
	})
}

func Fuzz_Nosy_OmniPortalSession_SetXSubValsetCutoff__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var xsubValsetCutoff_ uint8
		fill_err = tp.Fill(&xsubValsetCutoff_)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.SetXSubValsetCutoff(xsubValsetCutoff_)
	})
}

func Fuzz_Nosy_OmniPortalSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_OmniPortalSession_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.Unpause()
	})
}

func Fuzz_Nosy_OmniPortalSession_UnpauseXCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.UnpauseXCall()
	})
}

func Fuzz_Nosy_OmniPortalSession_UnpauseXCallTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var chainId_ uint64
		fill_err = tp.Fill(&chainId_)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.UnpauseXCallTo(chainId_)
	})
}

func Fuzz_Nosy_OmniPortalSession_UnpauseXSubmit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.UnpauseXSubmit()
	})
}

func Fuzz_Nosy_OmniPortalSession_UnpauseXSubmitFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var chainId_ uint64
		fill_err = tp.Fill(&chainId_)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.UnpauseXSubmitFrom(chainId_)
	})
}

func Fuzz_Nosy_OmniPortalSession_ValSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var arg0 uint64
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 common.Address
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ValSet(arg0, arg1)
	})
}

func Fuzz_Nosy_OmniPortalSession_ValSetTotalPower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var arg0 uint64
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.ValSetTotalPower(arg0)
	})
}

func Fuzz_Nosy_OmniPortalSession_XSubQuorumDenominator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.XSubQuorumDenominator()
	})
}

func Fuzz_Nosy_OmniPortalSession_XSubQuorumNumerator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.XSubQuorumNumerator()
	})
}

func Fuzz_Nosy_OmniPortalSession_Xcall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var destChainId uint64
		fill_err = tp.Fill(&destChainId)
		if fill_err != nil {
			return
		}
		var conf uint8
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.Xcall(destChainId, conf, to, d5, gasLimit)
	})
}

func Fuzz_Nosy_OmniPortalSession_Xmsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.Xmsg()
	})
}

func Fuzz_Nosy_OmniPortalSession_XmsgMaxDataSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.XmsgMaxDataSize()
	})
}

func Fuzz_Nosy_OmniPortalSession_XmsgMaxGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.XmsgMaxGasLimit()
	})
}

func Fuzz_Nosy_OmniPortalSession_XmsgMinGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.XmsgMinGasLimit()
	})
}

func Fuzz_Nosy_OmniPortalSession_XreceiptMaxErrorSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.XreceiptMaxErrorSize()
	})
}

func Fuzz_Nosy_OmniPortalSession_XsubValsetCutoff__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.XsubValsetCutoff()
	})
}

func Fuzz_Nosy_OmniPortalSession_Xsubmit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var xsub XTypesSubmission
		fill_err = tp.Fill(&xsub)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.Xsubmit(xsub)
	})
}

func Fuzz_Nosy_OmniPortalTransactor_AddValidatorSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactor
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var valSetId uint64
		fill_err = tp.Fill(&valSetId)
		if fill_err != nil {
			return
		}
		var validators []XTypesValidator
		fill_err = tp.Fill(&validators)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.AddValidatorSet(opts, valSetId, validators)
	})
}

func Fuzz_Nosy_OmniPortalTransactor_CollectFees__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactor
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.CollectFees(opts, to)
	})
}

func Fuzz_Nosy_OmniPortalTransactor_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactor
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var p OmniPortalInitParams
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.Initialize(opts, p)
	})
}

func Fuzz_Nosy_OmniPortalTransactor_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactor
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.Pause(opts)
	})
}

func Fuzz_Nosy_OmniPortalTransactor_PauseXCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactor
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.PauseXCall(opts)
	})
}

func Fuzz_Nosy_OmniPortalTransactor_PauseXCallTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactor
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var chainId_ uint64
		fill_err = tp.Fill(&chainId_)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.PauseXCallTo(opts, chainId_)
	})
}

func Fuzz_Nosy_OmniPortalTransactor_PauseXSubmit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactor
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.PauseXSubmit(opts)
	})
}

func Fuzz_Nosy_OmniPortalTransactor_PauseXSubmitFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactor
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var chainId_ uint64
		fill_err = tp.Fill(&chainId_)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.PauseXSubmitFrom(opts, chainId_)
	})
}

func Fuzz_Nosy_OmniPortalTransactor_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactor
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.RenounceOwnership(opts)
	})
}

func Fuzz_Nosy_OmniPortalTransactor_SetFeeOracle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactor
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var feeOracle_ common.Address
		fill_err = tp.Fill(&feeOracle_)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.SetFeeOracle(opts, feeOracle_)
	})
}

func Fuzz_Nosy_OmniPortalTransactor_SetInXBlockOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactor
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var sourceChainId uint64
		fill_err = tp.Fill(&sourceChainId)
		if fill_err != nil {
			return
		}
		var shardId uint64
		fill_err = tp.Fill(&shardId)
		if fill_err != nil {
			return
		}
		var offset uint64
		fill_err = tp.Fill(&offset)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.SetInXBlockOffset(opts, sourceChainId, shardId, offset)
	})
}

func Fuzz_Nosy_OmniPortalTransactor_SetInXMsgOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactor
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var sourceChainId uint64
		fill_err = tp.Fill(&sourceChainId)
		if fill_err != nil {
			return
		}
		var shardId uint64
		fill_err = tp.Fill(&shardId)
		if fill_err != nil {
			return
		}
		var offset uint64
		fill_err = tp.Fill(&offset)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.SetInXMsgOffset(opts, sourceChainId, shardId, offset)
	})
}

func Fuzz_Nosy_OmniPortalTransactor_SetNetwork__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactor
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var network_ []XTypesChain
		fill_err = tp.Fill(&network_)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.SetNetwork(opts, network_)
	})
}

func Fuzz_Nosy_OmniPortalTransactor_SetXMsgMaxDataSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactor
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var numBytes uint16
		fill_err = tp.Fill(&numBytes)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.SetXMsgMaxDataSize(opts, numBytes)
	})
}

func Fuzz_Nosy_OmniPortalTransactor_SetXMsgMaxGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactor
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.SetXMsgMaxGasLimit(opts, gasLimit)
	})
}

func Fuzz_Nosy_OmniPortalTransactor_SetXMsgMinGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactor
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.SetXMsgMinGasLimit(opts, gasLimit)
	})
}

func Fuzz_Nosy_OmniPortalTransactor_SetXReceiptMaxErrorSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactor
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var numBytes uint16
		fill_err = tp.Fill(&numBytes)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.SetXReceiptMaxErrorSize(opts, numBytes)
	})
}

func Fuzz_Nosy_OmniPortalTransactor_SetXSubValsetCutoff__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactor
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var xsubValsetCutoff_ uint8
		fill_err = tp.Fill(&xsubValsetCutoff_)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.SetXSubValsetCutoff(opts, xsubValsetCutoff_)
	})
}

func Fuzz_Nosy_OmniPortalTransactor_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactor
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.TransferOwnership(opts, newOwner)
	})
}

func Fuzz_Nosy_OmniPortalTransactor_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactor
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.Unpause(opts)
	})
}

func Fuzz_Nosy_OmniPortalTransactor_UnpauseXCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactor
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.UnpauseXCall(opts)
	})
}

func Fuzz_Nosy_OmniPortalTransactor_UnpauseXCallTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactor
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var chainId_ uint64
		fill_err = tp.Fill(&chainId_)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.UnpauseXCallTo(opts, chainId_)
	})
}

func Fuzz_Nosy_OmniPortalTransactor_UnpauseXSubmit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactor
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.UnpauseXSubmit(opts)
	})
}

func Fuzz_Nosy_OmniPortalTransactor_UnpauseXSubmitFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactor
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var chainId_ uint64
		fill_err = tp.Fill(&chainId_)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.UnpauseXSubmitFrom(opts, chainId_)
	})
}

func Fuzz_Nosy_OmniPortalTransactor_Xcall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactor
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var destChainId uint64
		fill_err = tp.Fill(&destChainId)
		if fill_err != nil {
			return
		}
		var conf uint8
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var d6 []byte
		fill_err = tp.Fill(&d6)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.Xcall(opts, destChainId, conf, to, d6, gasLimit)
	})
}

func Fuzz_Nosy_OmniPortalTransactor_Xsubmit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactor
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var xsub XTypesSubmission
		fill_err = tp.Fill(&xsub)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.Xsubmit(opts, xsub)
	})
}

// skipping Fuzz_Nosy_OmniPortalTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_OmniPortalTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactorRaw
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil || opts == nil {
			return
		}

		_OmniPortal.Transfer(opts)
	})
}

func Fuzz_Nosy_OmniPortalTransactorSession_AddValidatorSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactorSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var valSetId uint64
		fill_err = tp.Fill(&valSetId)
		if fill_err != nil {
			return
		}
		var validators []XTypesValidator
		fill_err = tp.Fill(&validators)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.AddValidatorSet(valSetId, validators)
	})
}

func Fuzz_Nosy_OmniPortalTransactorSession_CollectFees__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactorSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.CollectFees(to)
	})
}

func Fuzz_Nosy_OmniPortalTransactorSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactorSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var p OmniPortalInitParams
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.Initialize(p)
	})
}

func Fuzz_Nosy_OmniPortalTransactorSession_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactorSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.Pause()
	})
}

func Fuzz_Nosy_OmniPortalTransactorSession_PauseXCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactorSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.PauseXCall()
	})
}

func Fuzz_Nosy_OmniPortalTransactorSession_PauseXCallTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactorSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var chainId_ uint64
		fill_err = tp.Fill(&chainId_)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.PauseXCallTo(chainId_)
	})
}

func Fuzz_Nosy_OmniPortalTransactorSession_PauseXSubmit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactorSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.PauseXSubmit()
	})
}

func Fuzz_Nosy_OmniPortalTransactorSession_PauseXSubmitFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactorSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var chainId_ uint64
		fill_err = tp.Fill(&chainId_)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.PauseXSubmitFrom(chainId_)
	})
}

func Fuzz_Nosy_OmniPortalTransactorSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactorSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.RenounceOwnership()
	})
}

func Fuzz_Nosy_OmniPortalTransactorSession_SetFeeOracle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactorSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var feeOracle_ common.Address
		fill_err = tp.Fill(&feeOracle_)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.SetFeeOracle(feeOracle_)
	})
}

func Fuzz_Nosy_OmniPortalTransactorSession_SetInXBlockOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactorSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var sourceChainId uint64
		fill_err = tp.Fill(&sourceChainId)
		if fill_err != nil {
			return
		}
		var shardId uint64
		fill_err = tp.Fill(&shardId)
		if fill_err != nil {
			return
		}
		var offset uint64
		fill_err = tp.Fill(&offset)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.SetInXBlockOffset(sourceChainId, shardId, offset)
	})
}

func Fuzz_Nosy_OmniPortalTransactorSession_SetInXMsgOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactorSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var sourceChainId uint64
		fill_err = tp.Fill(&sourceChainId)
		if fill_err != nil {
			return
		}
		var shardId uint64
		fill_err = tp.Fill(&shardId)
		if fill_err != nil {
			return
		}
		var offset uint64
		fill_err = tp.Fill(&offset)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.SetInXMsgOffset(sourceChainId, shardId, offset)
	})
}

func Fuzz_Nosy_OmniPortalTransactorSession_SetNetwork__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactorSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var network_ []XTypesChain
		fill_err = tp.Fill(&network_)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.SetNetwork(network_)
	})
}

func Fuzz_Nosy_OmniPortalTransactorSession_SetXMsgMaxDataSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactorSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var numBytes uint16
		fill_err = tp.Fill(&numBytes)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.SetXMsgMaxDataSize(numBytes)
	})
}

func Fuzz_Nosy_OmniPortalTransactorSession_SetXMsgMaxGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactorSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.SetXMsgMaxGasLimit(gasLimit)
	})
}

func Fuzz_Nosy_OmniPortalTransactorSession_SetXMsgMinGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactorSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.SetXMsgMinGasLimit(gasLimit)
	})
}

func Fuzz_Nosy_OmniPortalTransactorSession_SetXReceiptMaxErrorSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactorSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var numBytes uint16
		fill_err = tp.Fill(&numBytes)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.SetXReceiptMaxErrorSize(numBytes)
	})
}

func Fuzz_Nosy_OmniPortalTransactorSession_SetXSubValsetCutoff__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactorSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var xsubValsetCutoff_ uint8
		fill_err = tp.Fill(&xsubValsetCutoff_)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.SetXSubValsetCutoff(xsubValsetCutoff_)
	})
}

func Fuzz_Nosy_OmniPortalTransactorSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactorSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_OmniPortalTransactorSession_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactorSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.Unpause()
	})
}

func Fuzz_Nosy_OmniPortalTransactorSession_UnpauseXCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactorSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.UnpauseXCall()
	})
}

func Fuzz_Nosy_OmniPortalTransactorSession_UnpauseXCallTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactorSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var chainId_ uint64
		fill_err = tp.Fill(&chainId_)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.UnpauseXCallTo(chainId_)
	})
}

func Fuzz_Nosy_OmniPortalTransactorSession_UnpauseXSubmit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactorSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.UnpauseXSubmit()
	})
}

func Fuzz_Nosy_OmniPortalTransactorSession_UnpauseXSubmitFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactorSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var chainId_ uint64
		fill_err = tp.Fill(&chainId_)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.UnpauseXSubmitFrom(chainId_)
	})
}

func Fuzz_Nosy_OmniPortalTransactorSession_Xcall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactorSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var destChainId uint64
		fill_err = tp.Fill(&destChainId)
		if fill_err != nil {
			return
		}
		var conf uint8
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.Xcall(destChainId, conf, to, d5, gasLimit)
	})
}

func Fuzz_Nosy_OmniPortalTransactorSession_Xsubmit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OmniPortal *OmniPortalTransactorSession
		fill_err = tp.Fill(&_OmniPortal)
		if fill_err != nil {
			return
		}
		var xsub XTypesSubmission
		fill_err = tp.Fill(&xsub)
		if fill_err != nil {
			return
		}
		if _OmniPortal == nil {
			return
		}

		_OmniPortal.Xsubmit(xsub)
	})
}

func Fuzz_Nosy_OmniPortalUnpausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniPortalUnpausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniPortalUnpausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniPortalValidatorSetAddedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalValidatorSetAddedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniPortalValidatorSetAddedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalValidatorSetAddedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniPortalValidatorSetAddedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalValidatorSetAddedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniPortalXCallPausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXCallPausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniPortalXCallPausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXCallPausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniPortalXCallPausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXCallPausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniPortalXCallToPausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXCallToPausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniPortalXCallToPausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXCallToPausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniPortalXCallToPausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXCallToPausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniPortalXCallToUnpausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXCallToUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniPortalXCallToUnpausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXCallToUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniPortalXCallToUnpausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXCallToUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniPortalXCallUnpausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXCallUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniPortalXCallUnpausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXCallUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniPortalXCallUnpausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXCallUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniPortalXMsgIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXMsgIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniPortalXMsgIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXMsgIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniPortalXMsgIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXMsgIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniPortalXMsgMaxDataSizeSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXMsgMaxDataSizeSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniPortalXMsgMaxDataSizeSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXMsgMaxDataSizeSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniPortalXMsgMaxDataSizeSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXMsgMaxDataSizeSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniPortalXMsgMaxGasLimitSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXMsgMaxGasLimitSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniPortalXMsgMaxGasLimitSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXMsgMaxGasLimitSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniPortalXMsgMaxGasLimitSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXMsgMaxGasLimitSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniPortalXMsgMinGasLimitSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXMsgMinGasLimitSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniPortalXMsgMinGasLimitSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXMsgMinGasLimitSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniPortalXMsgMinGasLimitSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXMsgMinGasLimitSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniPortalXReceiptIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXReceiptIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniPortalXReceiptIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXReceiptIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniPortalXReceiptIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXReceiptIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniPortalXReceiptMaxErrorSizeSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXReceiptMaxErrorSizeSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniPortalXReceiptMaxErrorSizeSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXReceiptMaxErrorSizeSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniPortalXReceiptMaxErrorSizeSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXReceiptMaxErrorSizeSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniPortalXSubValsetCutoffSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXSubValsetCutoffSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniPortalXSubValsetCutoffSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXSubValsetCutoffSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniPortalXSubValsetCutoffSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXSubValsetCutoffSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniPortalXSubmitFromPausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXSubmitFromPausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniPortalXSubmitFromPausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXSubmitFromPausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniPortalXSubmitFromPausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXSubmitFromPausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniPortalXSubmitFromUnpausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXSubmitFromUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniPortalXSubmitFromUnpausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXSubmitFromUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniPortalXSubmitFromUnpausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXSubmitFromUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniPortalXSubmitPausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXSubmitPausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniPortalXSubmitPausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXSubmitPausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniPortalXSubmitPausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXSubmitPausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_OmniPortalXSubmitUnpausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXSubmitUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniPortalXSubmitUnpausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXSubmitUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniPortalXSubmitUnpausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniPortalXSubmitUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

// skipping Fuzz_Nosy_OmniRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_OmniRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_OmniRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniRaw
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Omni == nil || opts == nil {
			return
		}

		_Omni.Transfer(opts)
	})
}

func Fuzz_Nosy_OmniSession_Allowance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var owner common.Address
		fill_err = tp.Fill(&owner)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		if _Omni == nil {
			return
		}

		_Omni.Allowance(owner, spender)
	})
}

func Fuzz_Nosy_OmniSession_Approve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _Omni == nil || amount == nil {
			return
		}

		_Omni.Approve(spender, amount)
	})
}

func Fuzz_Nosy_OmniSession_BalanceOf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if _Omni == nil {
			return
		}

		_Omni.BalanceOf(account)
	})
}

func Fuzz_Nosy_OmniSession_DOMAINSEPARATOR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		if _Omni == nil {
			return
		}

		_Omni.DOMAINSEPARATOR()
	})
}

func Fuzz_Nosy_OmniSession_Decimals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		if _Omni == nil {
			return
		}

		_Omni.Decimals()
	})
}

func Fuzz_Nosy_OmniSession_DecreaseAllowance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		var subtractedValue *big.Int
		fill_err = tp.Fill(&subtractedValue)
		if fill_err != nil {
			return
		}
		if _Omni == nil || subtractedValue == nil {
			return
		}

		_Omni.DecreaseAllowance(spender, subtractedValue)
	})
}

func Fuzz_Nosy_OmniSession_Eip712Domain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		if _Omni == nil {
			return
		}

		_Omni.Eip712Domain()
	})
}

func Fuzz_Nosy_OmniSession_IncreaseAllowance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		var addedValue *big.Int
		fill_err = tp.Fill(&addedValue)
		if fill_err != nil {
			return
		}
		if _Omni == nil || addedValue == nil {
			return
		}

		_Omni.IncreaseAllowance(spender, addedValue)
	})
}

func Fuzz_Nosy_OmniSession_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		if _Omni == nil {
			return
		}

		_Omni.Name()
	})
}

func Fuzz_Nosy_OmniSession_Nonces__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var owner common.Address
		fill_err = tp.Fill(&owner)
		if fill_err != nil {
			return
		}
		if _Omni == nil {
			return
		}

		_Omni.Nonces(owner)
	})
}

func Fuzz_Nosy_OmniSession_Permit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var owner common.Address
		fill_err = tp.Fill(&owner)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		var value *big.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		var deadline *big.Int
		fill_err = tp.Fill(&deadline)
		if fill_err != nil {
			return
		}
		var v uint8
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var r [32]byte
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var s [32]byte
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if _Omni == nil || value == nil || deadline == nil {
			return
		}

		_Omni.Permit(owner, spender, value, deadline, v, r, s)
	})
}

func Fuzz_Nosy_OmniSession_Symbol__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		if _Omni == nil {
			return
		}

		_Omni.Symbol()
	})
}

func Fuzz_Nosy_OmniSession_TotalSupply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		if _Omni == nil {
			return
		}

		_Omni.TotalSupply()
	})
}

func Fuzz_Nosy_OmniSession_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _Omni == nil || amount == nil {
			return
		}

		_Omni.Transfer(to, amount)
	})
}

func Fuzz_Nosy_OmniSession_TransferFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var from common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _Omni == nil || amount == nil {
			return
		}

		_Omni.TransferFrom(from, to, amount)
	})
}

func Fuzz_Nosy_OmniTransactor_Approve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniTransactor
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _Omni == nil || opts == nil || amount == nil {
			return
		}

		_Omni.Approve(opts, spender, amount)
	})
}

func Fuzz_Nosy_OmniTransactor_DecreaseAllowance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniTransactor
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		var subtractedValue *big.Int
		fill_err = tp.Fill(&subtractedValue)
		if fill_err != nil {
			return
		}
		if _Omni == nil || opts == nil || subtractedValue == nil {
			return
		}

		_Omni.DecreaseAllowance(opts, spender, subtractedValue)
	})
}

func Fuzz_Nosy_OmniTransactor_IncreaseAllowance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniTransactor
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		var addedValue *big.Int
		fill_err = tp.Fill(&addedValue)
		if fill_err != nil {
			return
		}
		if _Omni == nil || opts == nil || addedValue == nil {
			return
		}

		_Omni.IncreaseAllowance(opts, spender, addedValue)
	})
}

func Fuzz_Nosy_OmniTransactor_Permit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniTransactor
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var owner common.Address
		fill_err = tp.Fill(&owner)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		var value *big.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		var deadline *big.Int
		fill_err = tp.Fill(&deadline)
		if fill_err != nil {
			return
		}
		var v uint8
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var r [32]byte
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var s [32]byte
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if _Omni == nil || opts == nil || value == nil || deadline == nil {
			return
		}

		_Omni.Permit(opts, owner, spender, value, deadline, v, r, s)
	})
}

func Fuzz_Nosy_OmniTransactor_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniTransactor
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _Omni == nil || opts == nil || amount == nil {
			return
		}

		_Omni.Transfer(opts, to, amount)
	})
}

func Fuzz_Nosy_OmniTransactor_TransferFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniTransactor
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var from common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _Omni == nil || opts == nil || amount == nil {
			return
		}

		_Omni.TransferFrom(opts, from, to, amount)
	})
}

// skipping Fuzz_Nosy_OmniTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_OmniTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniTransactorRaw
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Omni == nil || opts == nil {
			return
		}

		_Omni.Transfer(opts)
	})
}

func Fuzz_Nosy_OmniTransactorSession_Approve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniTransactorSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _Omni == nil || amount == nil {
			return
		}

		_Omni.Approve(spender, amount)
	})
}

func Fuzz_Nosy_OmniTransactorSession_DecreaseAllowance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniTransactorSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		var subtractedValue *big.Int
		fill_err = tp.Fill(&subtractedValue)
		if fill_err != nil {
			return
		}
		if _Omni == nil || subtractedValue == nil {
			return
		}

		_Omni.DecreaseAllowance(spender, subtractedValue)
	})
}

func Fuzz_Nosy_OmniTransactorSession_IncreaseAllowance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniTransactorSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		var addedValue *big.Int
		fill_err = tp.Fill(&addedValue)
		if fill_err != nil {
			return
		}
		if _Omni == nil || addedValue == nil {
			return
		}

		_Omni.IncreaseAllowance(spender, addedValue)
	})
}

func Fuzz_Nosy_OmniTransactorSession_Permit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniTransactorSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var owner common.Address
		fill_err = tp.Fill(&owner)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		var value *big.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		var deadline *big.Int
		fill_err = tp.Fill(&deadline)
		if fill_err != nil {
			return
		}
		var v uint8
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var r [32]byte
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var s [32]byte
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if _Omni == nil || value == nil || deadline == nil {
			return
		}

		_Omni.Permit(owner, spender, value, deadline, v, r, s)
	})
}

func Fuzz_Nosy_OmniTransactorSession_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniTransactorSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _Omni == nil || amount == nil {
			return
		}

		_Omni.Transfer(to, amount)
	})
}

func Fuzz_Nosy_OmniTransactorSession_TransferFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Omni *OmniTransactorSession
		fill_err = tp.Fill(&_Omni)
		if fill_err != nil {
			return
		}
		var from common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _Omni == nil || amount == nil {
			return
		}

		_Omni.TransferFrom(from, to, amount)
	})
}

func Fuzz_Nosy_OmniTransferIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniTransferIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_OmniTransferIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniTransferIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_OmniTransferIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OmniTransferIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_PingPongCaller_DefaultConfLevel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongCaller
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _PingPong == nil || opts == nil {
			return
		}

		_PingPong.DefaultConfLevel(opts)
	})
}

func Fuzz_Nosy_PingPongCaller_GASLIMIT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongCaller
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _PingPong == nil || opts == nil {
			return
		}

		_PingPong.GASLIMIT(opts)
	})
}

func Fuzz_Nosy_PingPongCaller_Omni__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongCaller
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _PingPong == nil || opts == nil {
			return
		}

		_PingPong.Omni(opts)
	})
}

// skipping Fuzz_Nosy_PingPongCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_PingPongCallerSession_DefaultConfLevel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongCallerSession
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		if _PingPong == nil {
			return
		}

		_PingPong.DefaultConfLevel()
	})
}

func Fuzz_Nosy_PingPongCallerSession_GASLIMIT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongCallerSession
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		if _PingPong == nil {
			return
		}

		_PingPong.GASLIMIT()
	})
}

func Fuzz_Nosy_PingPongCallerSession_Omni__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongCallerSession
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		if _PingPong == nil {
			return
		}

		_PingPong.Omni()
	})
}

func Fuzz_Nosy_PingPongDefaultConfLevelSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *PingPongDefaultConfLevelSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_PingPongDefaultConfLevelSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *PingPongDefaultConfLevelSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_PingPongDefaultConfLevelSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *PingPongDefaultConfLevelSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_PingPongDoneIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *PingPongDoneIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_PingPongDoneIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *PingPongDoneIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_PingPongDoneIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *PingPongDoneIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_PingPongFilterer_FilterDefaultConfLevelSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongFilterer
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _PingPong == nil || opts == nil {
			return
		}

		_PingPong.FilterDefaultConfLevelSet(opts)
	})
}

func Fuzz_Nosy_PingPongFilterer_FilterDone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongFilterer
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _PingPong == nil || opts == nil {
			return
		}

		_PingPong.FilterDone(opts)
	})
}

func Fuzz_Nosy_PingPongFilterer_FilterOmniPortalSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongFilterer
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _PingPong == nil || opts == nil {
			return
		}

		_PingPong.FilterOmniPortalSet(opts)
	})
}

func Fuzz_Nosy_PingPongFilterer_FilterPing__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongFilterer
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _PingPong == nil || opts == nil {
			return
		}

		_PingPong.FilterPing(opts)
	})
}

func Fuzz_Nosy_PingPongFilterer_ParseDefaultConfLevelSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongFilterer
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _PingPong == nil {
			return
		}

		_PingPong.ParseDefaultConfLevelSet(log)
	})
}

func Fuzz_Nosy_PingPongFilterer_ParseDone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongFilterer
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _PingPong == nil {
			return
		}

		_PingPong.ParseDone(log)
	})
}

func Fuzz_Nosy_PingPongFilterer_ParseOmniPortalSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongFilterer
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _PingPong == nil {
			return
		}

		_PingPong.ParseOmniPortalSet(log)
	})
}

func Fuzz_Nosy_PingPongFilterer_ParsePing__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongFilterer
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _PingPong == nil {
			return
		}

		_PingPong.ParsePing(log)
	})
}

// skipping Fuzz_Nosy_PingPongFilterer_WatchDefaultConfLevelSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.PingPongDefaultConfLevelSet

// skipping Fuzz_Nosy_PingPongFilterer_WatchDone__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.PingPongDone

// skipping Fuzz_Nosy_PingPongFilterer_WatchOmniPortalSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.PingPongOmniPortalSet

// skipping Fuzz_Nosy_PingPongFilterer_WatchPing__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.PingPongPing

func Fuzz_Nosy_PingPongOmniPortalSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *PingPongOmniPortalSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_PingPongOmniPortalSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *PingPongOmniPortalSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_PingPongOmniPortalSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *PingPongOmniPortalSetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_PingPongPingIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *PingPongPingIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_PingPongPingIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *PingPongPingIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_PingPongPingIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *PingPongPingIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

// skipping Fuzz_Nosy_PingPongRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_PingPongRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_PingPongRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongRaw
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _PingPong == nil || opts == nil {
			return
		}

		_PingPong.Transfer(opts)
	})
}

func Fuzz_Nosy_PingPongSession_DefaultConfLevel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongSession
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		if _PingPong == nil {
			return
		}

		_PingPong.DefaultConfLevel()
	})
}

func Fuzz_Nosy_PingPongSession_GASLIMIT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongSession
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		if _PingPong == nil {
			return
		}

		_PingPong.GASLIMIT()
	})
}

func Fuzz_Nosy_PingPongSession_Omni__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongSession
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		if _PingPong == nil {
			return
		}

		_PingPong.Omni()
	})
}

func Fuzz_Nosy_PingPongSession_Pingpong__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongSession
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var conf uint8
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		var times uint64
		fill_err = tp.Fill(&times)
		if fill_err != nil {
			return
		}
		var n uint64
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if _PingPong == nil {
			return
		}

		_PingPong.Pingpong(id, conf, times, n)
	})
}

func Fuzz_Nosy_PingPongSession_Receive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongSession
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		if _PingPong == nil {
			return
		}

		_PingPong.Receive()
	})
}

func Fuzz_Nosy_PingPongSession_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongSession
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var destChainID uint64
		fill_err = tp.Fill(&destChainID)
		if fill_err != nil {
			return
		}
		var conf uint8
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var times uint64
		fill_err = tp.Fill(&times)
		if fill_err != nil {
			return
		}
		if _PingPong == nil {
			return
		}

		_PingPong.Start(id, destChainID, conf, to, times)
	})
}

func Fuzz_Nosy_PingPongTransactor_Pingpong__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongTransactor
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var conf uint8
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		var times uint64
		fill_err = tp.Fill(&times)
		if fill_err != nil {
			return
		}
		var n uint64
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if _PingPong == nil || opts == nil {
			return
		}

		_PingPong.Pingpong(opts, id, conf, times, n)
	})
}

func Fuzz_Nosy_PingPongTransactor_Receive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongTransactor
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _PingPong == nil || opts == nil {
			return
		}

		_PingPong.Receive(opts)
	})
}

func Fuzz_Nosy_PingPongTransactor_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongTransactor
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var destChainID uint64
		fill_err = tp.Fill(&destChainID)
		if fill_err != nil {
			return
		}
		var conf uint8
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var times uint64
		fill_err = tp.Fill(&times)
		if fill_err != nil {
			return
		}
		if _PingPong == nil || opts == nil {
			return
		}

		_PingPong.Start(opts, id, destChainID, conf, to, times)
	})
}

// skipping Fuzz_Nosy_PingPongTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_PingPongTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongTransactorRaw
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _PingPong == nil || opts == nil {
			return
		}

		_PingPong.Transfer(opts)
	})
}

func Fuzz_Nosy_PingPongTransactorSession_Pingpong__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongTransactorSession
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var conf uint8
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		var times uint64
		fill_err = tp.Fill(&times)
		if fill_err != nil {
			return
		}
		var n uint64
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if _PingPong == nil {
			return
		}

		_PingPong.Pingpong(id, conf, times, n)
	})
}

func Fuzz_Nosy_PingPongTransactorSession_Receive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongTransactorSession
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		if _PingPong == nil {
			return
		}

		_PingPong.Receive()
	})
}

func Fuzz_Nosy_PingPongTransactorSession_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PingPong *PingPongTransactorSession
		fill_err = tp.Fill(&_PingPong)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var destChainID uint64
		fill_err = tp.Fill(&destChainID)
		if fill_err != nil {
			return
		}
		var conf uint8
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var times uint64
		fill_err = tp.Fill(&times)
		if fill_err != nil {
			return
		}
		if _PingPong == nil {
			return
		}

		_PingPong.Start(id, destChainID, conf, to, times)
	})
}

func Fuzz_Nosy_PortalRegistryCaller_ChainIds__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryCaller
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 *big.Int
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil || opts == nil || arg0 == nil {
			return
		}

		_PortalRegistry.ChainIds(opts, arg0)
	})
}

func Fuzz_Nosy_PortalRegistryCaller_Deployments__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryCaller
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 uint64
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil || opts == nil {
			return
		}

		_PortalRegistry.Deployments(opts, arg0)
	})
}

func Fuzz_Nosy_PortalRegistryCaller_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryCaller
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var chainId uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil || opts == nil {
			return
		}

		_PortalRegistry.Get(opts, chainId)
	})
}

func Fuzz_Nosy_PortalRegistryCaller_List__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryCaller
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil || opts == nil {
			return
		}

		_PortalRegistry.List(opts)
	})
}

func Fuzz_Nosy_PortalRegistryCaller_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryCaller
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil || opts == nil {
			return
		}

		_PortalRegistry.Owner(opts)
	})
}

// skipping Fuzz_Nosy_PortalRegistryCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_PortalRegistryCallerSession_ChainIds__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryCallerSession
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var arg0 *big.Int
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil || arg0 == nil {
			return
		}

		_PortalRegistry.ChainIds(arg0)
	})
}

func Fuzz_Nosy_PortalRegistryCallerSession_Deployments__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryCallerSession
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var arg0 uint64
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil {
			return
		}

		_PortalRegistry.Deployments(arg0)
	})
}

func Fuzz_Nosy_PortalRegistryCallerSession_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryCallerSession
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var chainId uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil {
			return
		}

		_PortalRegistry.Get(chainId)
	})
}

func Fuzz_Nosy_PortalRegistryCallerSession_List__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryCallerSession
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil {
			return
		}

		_PortalRegistry.List()
	})
}

func Fuzz_Nosy_PortalRegistryCallerSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryCallerSession
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil {
			return
		}

		_PortalRegistry.Owner()
	})
}

func Fuzz_Nosy_PortalRegistryFilterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryFilterer
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil || opts == nil {
			return
		}

		_PortalRegistry.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_PortalRegistryFilterer_FilterOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryFilterer
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var previousOwner []common.Address
		fill_err = tp.Fill(&previousOwner)
		if fill_err != nil {
			return
		}
		var newOwner []common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil || opts == nil {
			return
		}

		_PortalRegistry.FilterOwnershipTransferred(opts, previousOwner, newOwner)
	})
}

func Fuzz_Nosy_PortalRegistryFilterer_FilterPortalRegistered__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryFilterer
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var chainId []uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var addr []common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil || opts == nil {
			return
		}

		_PortalRegistry.FilterPortalRegistered(opts, chainId, addr)
	})
}

func Fuzz_Nosy_PortalRegistryFilterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryFilterer
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil {
			return
		}

		_PortalRegistry.ParseInitialized(log)
	})
}

func Fuzz_Nosy_PortalRegistryFilterer_ParseOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryFilterer
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil {
			return
		}

		_PortalRegistry.ParseOwnershipTransferred(log)
	})
}

func Fuzz_Nosy_PortalRegistryFilterer_ParsePortalRegistered__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryFilterer
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil {
			return
		}

		_PortalRegistry.ParsePortalRegistered(log)
	})
}

// skipping Fuzz_Nosy_PortalRegistryFilterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.PortalRegistryInitialized

// skipping Fuzz_Nosy_PortalRegistryFilterer_WatchOwnershipTransferred__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.PortalRegistryOwnershipTransferred

// skipping Fuzz_Nosy_PortalRegistryFilterer_WatchPortalRegistered__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.PortalRegistryPortalRegistered

func Fuzz_Nosy_PortalRegistryInitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *PortalRegistryInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_PortalRegistryInitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *PortalRegistryInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_PortalRegistryInitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *PortalRegistryInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_PortalRegistryOwnershipTransferredIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *PortalRegistryOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_PortalRegistryOwnershipTransferredIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *PortalRegistryOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_PortalRegistryOwnershipTransferredIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *PortalRegistryOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_PortalRegistryPortalRegisteredIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *PortalRegistryPortalRegisteredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_PortalRegistryPortalRegisteredIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *PortalRegistryPortalRegisteredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_PortalRegistryPortalRegisteredIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *PortalRegistryPortalRegisteredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

// skipping Fuzz_Nosy_PortalRegistryRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_PortalRegistryRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_PortalRegistryRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryRaw
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil || opts == nil {
			return
		}

		_PortalRegistry.Transfer(opts)
	})
}

func Fuzz_Nosy_PortalRegistrySession_BulkRegister__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistrySession
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var deps []PortalRegistryDeployment
		fill_err = tp.Fill(&deps)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil {
			return
		}

		_PortalRegistry.BulkRegister(deps)
	})
}

func Fuzz_Nosy_PortalRegistrySession_ChainIds__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistrySession
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var arg0 *big.Int
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil || arg0 == nil {
			return
		}

		_PortalRegistry.ChainIds(arg0)
	})
}

func Fuzz_Nosy_PortalRegistrySession_Deployments__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistrySession
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var arg0 uint64
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil {
			return
		}

		_PortalRegistry.Deployments(arg0)
	})
}

func Fuzz_Nosy_PortalRegistrySession_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistrySession
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var chainId uint64
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil {
			return
		}

		_PortalRegistry.Get(chainId)
	})
}

func Fuzz_Nosy_PortalRegistrySession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistrySession
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var owner_ common.Address
		fill_err = tp.Fill(&owner_)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil {
			return
		}

		_PortalRegistry.Initialize(owner_)
	})
}

func Fuzz_Nosy_PortalRegistrySession_List__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistrySession
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil {
			return
		}

		_PortalRegistry.List()
	})
}

func Fuzz_Nosy_PortalRegistrySession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistrySession
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil {
			return
		}

		_PortalRegistry.Owner()
	})
}

func Fuzz_Nosy_PortalRegistrySession_Register__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistrySession
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var dep PortalRegistryDeployment
		fill_err = tp.Fill(&dep)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil {
			return
		}

		_PortalRegistry.Register(dep)
	})
}

func Fuzz_Nosy_PortalRegistrySession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistrySession
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil {
			return
		}

		_PortalRegistry.RenounceOwnership()
	})
}

func Fuzz_Nosy_PortalRegistrySession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistrySession
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil {
			return
		}

		_PortalRegistry.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_PortalRegistryTransactor_BulkRegister__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryTransactor
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var deps []PortalRegistryDeployment
		fill_err = tp.Fill(&deps)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil || opts == nil {
			return
		}

		_PortalRegistry.BulkRegister(opts, deps)
	})
}

func Fuzz_Nosy_PortalRegistryTransactor_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryTransactor
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var owner_ common.Address
		fill_err = tp.Fill(&owner_)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil || opts == nil {
			return
		}

		_PortalRegistry.Initialize(opts, owner_)
	})
}

func Fuzz_Nosy_PortalRegistryTransactor_Register__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryTransactor
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var dep PortalRegistryDeployment
		fill_err = tp.Fill(&dep)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil || opts == nil {
			return
		}

		_PortalRegistry.Register(opts, dep)
	})
}

func Fuzz_Nosy_PortalRegistryTransactor_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryTransactor
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil || opts == nil {
			return
		}

		_PortalRegistry.RenounceOwnership(opts)
	})
}

func Fuzz_Nosy_PortalRegistryTransactor_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryTransactor
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil || opts == nil {
			return
		}

		_PortalRegistry.TransferOwnership(opts, newOwner)
	})
}

// skipping Fuzz_Nosy_PortalRegistryTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_PortalRegistryTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryTransactorRaw
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil || opts == nil {
			return
		}

		_PortalRegistry.Transfer(opts)
	})
}

func Fuzz_Nosy_PortalRegistryTransactorSession_BulkRegister__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryTransactorSession
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var deps []PortalRegistryDeployment
		fill_err = tp.Fill(&deps)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil {
			return
		}

		_PortalRegistry.BulkRegister(deps)
	})
}

func Fuzz_Nosy_PortalRegistryTransactorSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryTransactorSession
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var owner_ common.Address
		fill_err = tp.Fill(&owner_)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil {
			return
		}

		_PortalRegistry.Initialize(owner_)
	})
}

func Fuzz_Nosy_PortalRegistryTransactorSession_Register__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryTransactorSession
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var dep PortalRegistryDeployment
		fill_err = tp.Fill(&dep)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil {
			return
		}

		_PortalRegistry.Register(dep)
	})
}

func Fuzz_Nosy_PortalRegistryTransactorSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryTransactorSession
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil {
			return
		}

		_PortalRegistry.RenounceOwnership()
	})
}

func Fuzz_Nosy_PortalRegistryTransactorSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _PortalRegistry *PortalRegistryTransactorSession
		fill_err = tp.Fill(&_PortalRegistry)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _PortalRegistry == nil {
			return
		}

		_PortalRegistry.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_ProxyAdminCaller_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ProxyAdmin *ProxyAdminCaller
		fill_err = tp.Fill(&_ProxyAdmin)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ProxyAdmin == nil || opts == nil {
			return
		}

		_ProxyAdmin.Owner(opts)
	})
}

func Fuzz_Nosy_ProxyAdminCaller_UPGRADEINTERFACEVERSION__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ProxyAdmin *ProxyAdminCaller
		fill_err = tp.Fill(&_ProxyAdmin)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ProxyAdmin == nil || opts == nil {
			return
		}

		_ProxyAdmin.UPGRADEINTERFACEVERSION(opts)
	})
}

// skipping Fuzz_Nosy_ProxyAdminCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_ProxyAdminCallerSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ProxyAdmin *ProxyAdminCallerSession
		fill_err = tp.Fill(&_ProxyAdmin)
		if fill_err != nil {
			return
		}
		if _ProxyAdmin == nil {
			return
		}

		_ProxyAdmin.Owner()
	})
}

func Fuzz_Nosy_ProxyAdminCallerSession_UPGRADEINTERFACEVERSION__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ProxyAdmin *ProxyAdminCallerSession
		fill_err = tp.Fill(&_ProxyAdmin)
		if fill_err != nil {
			return
		}
		if _ProxyAdmin == nil {
			return
		}

		_ProxyAdmin.UPGRADEINTERFACEVERSION()
	})
}

func Fuzz_Nosy_ProxyAdminFilterer_FilterOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ProxyAdmin *ProxyAdminFilterer
		fill_err = tp.Fill(&_ProxyAdmin)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var previousOwner []common.Address
		fill_err = tp.Fill(&previousOwner)
		if fill_err != nil {
			return
		}
		var newOwner []common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _ProxyAdmin == nil || opts == nil {
			return
		}

		_ProxyAdmin.FilterOwnershipTransferred(opts, previousOwner, newOwner)
	})
}

func Fuzz_Nosy_ProxyAdminFilterer_ParseOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ProxyAdmin *ProxyAdminFilterer
		fill_err = tp.Fill(&_ProxyAdmin)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ProxyAdmin == nil {
			return
		}

		_ProxyAdmin.ParseOwnershipTransferred(log)
	})
}

// skipping Fuzz_Nosy_ProxyAdminFilterer_WatchOwnershipTransferred__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.ProxyAdminOwnershipTransferred

func Fuzz_Nosy_ProxyAdminOwnershipTransferredIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ProxyAdminOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_ProxyAdminOwnershipTransferredIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ProxyAdminOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_ProxyAdminOwnershipTransferredIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ProxyAdminOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

// skipping Fuzz_Nosy_ProxyAdminRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_ProxyAdminRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ProxyAdminRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ProxyAdmin *ProxyAdminRaw
		fill_err = tp.Fill(&_ProxyAdmin)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ProxyAdmin == nil || opts == nil {
			return
		}

		_ProxyAdmin.Transfer(opts)
	})
}

func Fuzz_Nosy_ProxyAdminSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ProxyAdmin *ProxyAdminSession
		fill_err = tp.Fill(&_ProxyAdmin)
		if fill_err != nil {
			return
		}
		if _ProxyAdmin == nil {
			return
		}

		_ProxyAdmin.Owner()
	})
}

func Fuzz_Nosy_ProxyAdminSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ProxyAdmin *ProxyAdminSession
		fill_err = tp.Fill(&_ProxyAdmin)
		if fill_err != nil {
			return
		}
		if _ProxyAdmin == nil {
			return
		}

		_ProxyAdmin.RenounceOwnership()
	})
}

func Fuzz_Nosy_ProxyAdminSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ProxyAdmin *ProxyAdminSession
		fill_err = tp.Fill(&_ProxyAdmin)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _ProxyAdmin == nil {
			return
		}

		_ProxyAdmin.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_ProxyAdminSession_UPGRADEINTERFACEVERSION__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ProxyAdmin *ProxyAdminSession
		fill_err = tp.Fill(&_ProxyAdmin)
		if fill_err != nil {
			return
		}
		if _ProxyAdmin == nil {
			return
		}

		_ProxyAdmin.UPGRADEINTERFACEVERSION()
	})
}

func Fuzz_Nosy_ProxyAdminSession_UpgradeAndCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ProxyAdmin *ProxyAdminSession
		fill_err = tp.Fill(&_ProxyAdmin)
		if fill_err != nil {
			return
		}
		var proxy common.Address
		fill_err = tp.Fill(&proxy)
		if fill_err != nil {
			return
		}
		var implementation common.Address
		fill_err = tp.Fill(&implementation)
		if fill_err != nil {
			return
		}
		var d4 []byte
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}
		if _ProxyAdmin == nil {
			return
		}

		_ProxyAdmin.UpgradeAndCall(proxy, implementation, d4)
	})
}

func Fuzz_Nosy_ProxyAdminTransactor_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ProxyAdmin *ProxyAdminTransactor
		fill_err = tp.Fill(&_ProxyAdmin)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ProxyAdmin == nil || opts == nil {
			return
		}

		_ProxyAdmin.RenounceOwnership(opts)
	})
}

func Fuzz_Nosy_ProxyAdminTransactor_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ProxyAdmin *ProxyAdminTransactor
		fill_err = tp.Fill(&_ProxyAdmin)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _ProxyAdmin == nil || opts == nil {
			return
		}

		_ProxyAdmin.TransferOwnership(opts, newOwner)
	})
}

func Fuzz_Nosy_ProxyAdminTransactor_UpgradeAndCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ProxyAdmin *ProxyAdminTransactor
		fill_err = tp.Fill(&_ProxyAdmin)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var proxy common.Address
		fill_err = tp.Fill(&proxy)
		if fill_err != nil {
			return
		}
		var implementation common.Address
		fill_err = tp.Fill(&implementation)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if _ProxyAdmin == nil || opts == nil {
			return
		}

		_ProxyAdmin.UpgradeAndCall(opts, proxy, implementation, d5)
	})
}

// skipping Fuzz_Nosy_ProxyAdminTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ProxyAdminTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ProxyAdmin *ProxyAdminTransactorRaw
		fill_err = tp.Fill(&_ProxyAdmin)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ProxyAdmin == nil || opts == nil {
			return
		}

		_ProxyAdmin.Transfer(opts)
	})
}

func Fuzz_Nosy_ProxyAdminTransactorSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ProxyAdmin *ProxyAdminTransactorSession
		fill_err = tp.Fill(&_ProxyAdmin)
		if fill_err != nil {
			return
		}
		if _ProxyAdmin == nil {
			return
		}

		_ProxyAdmin.RenounceOwnership()
	})
}

func Fuzz_Nosy_ProxyAdminTransactorSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ProxyAdmin *ProxyAdminTransactorSession
		fill_err = tp.Fill(&_ProxyAdmin)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _ProxyAdmin == nil {
			return
		}

		_ProxyAdmin.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_ProxyAdminTransactorSession_UpgradeAndCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ProxyAdmin *ProxyAdminTransactorSession
		fill_err = tp.Fill(&_ProxyAdmin)
		if fill_err != nil {
			return
		}
		var proxy common.Address
		fill_err = tp.Fill(&proxy)
		if fill_err != nil {
			return
		}
		var implementation common.Address
		fill_err = tp.Fill(&implementation)
		if fill_err != nil {
			return
		}
		var d4 []byte
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}
		if _ProxyAdmin == nil {
			return
		}

		_ProxyAdmin.UpgradeAndCall(proxy, implementation, d4)
	})
}

func Fuzz_Nosy_SlashingCaller_Fee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Slashing *SlashingCaller
		fill_err = tp.Fill(&_Slashing)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Slashing == nil || opts == nil {
			return
		}

		_Slashing.Fee(opts)
	})
}

// skipping Fuzz_Nosy_SlashingCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_SlashingCallerSession_Fee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Slashing *SlashingCallerSession
		fill_err = tp.Fill(&_Slashing)
		if fill_err != nil {
			return
		}
		if _Slashing == nil {
			return
		}

		_Slashing.Fee()
	})
}

func Fuzz_Nosy_SlashingFilterer_FilterUnjail__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Slashing *SlashingFilterer
		fill_err = tp.Fill(&_Slashing)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var validator []common.Address
		fill_err = tp.Fill(&validator)
		if fill_err != nil {
			return
		}
		if _Slashing == nil || opts == nil {
			return
		}

		_Slashing.FilterUnjail(opts, validator)
	})
}

func Fuzz_Nosy_SlashingFilterer_ParseUnjail__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Slashing *SlashingFilterer
		fill_err = tp.Fill(&_Slashing)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _Slashing == nil {
			return
		}

		_Slashing.ParseUnjail(log)
	})
}

// skipping Fuzz_Nosy_SlashingFilterer_WatchUnjail__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.SlashingUnjail

// skipping Fuzz_Nosy_SlashingRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_SlashingRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_SlashingRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Slashing *SlashingRaw
		fill_err = tp.Fill(&_Slashing)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Slashing == nil || opts == nil {
			return
		}

		_Slashing.Transfer(opts)
	})
}

func Fuzz_Nosy_SlashingSession_Fee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Slashing *SlashingSession
		fill_err = tp.Fill(&_Slashing)
		if fill_err != nil {
			return
		}
		if _Slashing == nil {
			return
		}

		_Slashing.Fee()
	})
}

func Fuzz_Nosy_SlashingSession_Unjail__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Slashing *SlashingSession
		fill_err = tp.Fill(&_Slashing)
		if fill_err != nil {
			return
		}
		if _Slashing == nil {
			return
		}

		_Slashing.Unjail()
	})
}

func Fuzz_Nosy_SlashingTransactor_Unjail__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Slashing *SlashingTransactor
		fill_err = tp.Fill(&_Slashing)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Slashing == nil || opts == nil {
			return
		}

		_Slashing.Unjail(opts)
	})
}

// skipping Fuzz_Nosy_SlashingTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_SlashingTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Slashing *SlashingTransactorRaw
		fill_err = tp.Fill(&_Slashing)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Slashing == nil || opts == nil {
			return
		}

		_Slashing.Transfer(opts)
	})
}

func Fuzz_Nosy_SlashingTransactorSession_Unjail__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Slashing *SlashingTransactorSession
		fill_err = tp.Fill(&_Slashing)
		if fill_err != nil {
			return
		}
		if _Slashing == nil {
			return
		}

		_Slashing.Unjail()
	})
}

func Fuzz_Nosy_SlashingUnjailIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *SlashingUnjailIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_SlashingUnjailIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *SlashingUnjailIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_SlashingUnjailIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *SlashingUnjailIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_StakingCaller_IsAllowedValidator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingCaller
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _Staking == nil || opts == nil {
			return
		}

		_Staking.IsAllowedValidator(opts, arg0)
	})
}

func Fuzz_Nosy_StakingCaller_IsAllowlistEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingCaller
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Staking == nil || opts == nil {
			return
		}

		_Staking.IsAllowlistEnabled(opts)
	})
}

func Fuzz_Nosy_StakingCaller_MinDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingCaller
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Staking == nil || opts == nil {
			return
		}

		_Staking.MinDelegation(opts)
	})
}

func Fuzz_Nosy_StakingCaller_MinDeposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingCaller
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Staking == nil || opts == nil {
			return
		}

		_Staking.MinDeposit(opts)
	})
}

func Fuzz_Nosy_StakingCaller_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingCaller
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Staking == nil || opts == nil {
			return
		}

		_Staking.Owner(opts)
	})
}

// skipping Fuzz_Nosy_StakingCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_StakingCallerSession_IsAllowedValidator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingCallerSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.IsAllowedValidator(arg0)
	})
}

func Fuzz_Nosy_StakingCallerSession_IsAllowlistEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingCallerSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.IsAllowlistEnabled()
	})
}

func Fuzz_Nosy_StakingCallerSession_MinDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingCallerSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.MinDelegation()
	})
}

func Fuzz_Nosy_StakingCallerSession_MinDeposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingCallerSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.MinDeposit()
	})
}

func Fuzz_Nosy_StakingCallerSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingCallerSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.Owner()
	})
}

func Fuzz_Nosy_StakingCreateValidatorIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StakingCreateValidatorIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_StakingCreateValidatorIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StakingCreateValidatorIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_StakingCreateValidatorIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StakingCreateValidatorIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_StakingDelegateIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StakingDelegateIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_StakingDelegateIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StakingDelegateIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_StakingDelegateIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StakingDelegateIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_StakingFilterer_FilterCreateValidator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingFilterer
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var validator []common.Address
		fill_err = tp.Fill(&validator)
		if fill_err != nil {
			return
		}
		if _Staking == nil || opts == nil {
			return
		}

		_Staking.FilterCreateValidator(opts, validator)
	})
}

func Fuzz_Nosy_StakingFilterer_FilterDelegate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingFilterer
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var delegator []common.Address
		fill_err = tp.Fill(&delegator)
		if fill_err != nil {
			return
		}
		var validator []common.Address
		fill_err = tp.Fill(&validator)
		if fill_err != nil {
			return
		}
		if _Staking == nil || opts == nil {
			return
		}

		_Staking.FilterDelegate(opts, delegator, validator)
	})
}

func Fuzz_Nosy_StakingFilterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingFilterer
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Staking == nil || opts == nil {
			return
		}

		_Staking.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_StakingFilterer_FilterOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingFilterer
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var previousOwner []common.Address
		fill_err = tp.Fill(&previousOwner)
		if fill_err != nil {
			return
		}
		var newOwner []common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _Staking == nil || opts == nil {
			return
		}

		_Staking.FilterOwnershipTransferred(opts, previousOwner, newOwner)
	})
}

func Fuzz_Nosy_StakingFilterer_ParseCreateValidator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingFilterer
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.ParseCreateValidator(log)
	})
}

func Fuzz_Nosy_StakingFilterer_ParseDelegate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingFilterer
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.ParseDelegate(log)
	})
}

func Fuzz_Nosy_StakingFilterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingFilterer
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.ParseInitialized(log)
	})
}

func Fuzz_Nosy_StakingFilterer_ParseOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingFilterer
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.ParseOwnershipTransferred(log)
	})
}

// skipping Fuzz_Nosy_StakingFilterer_WatchCreateValidator__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.StakingCreateValidator

// skipping Fuzz_Nosy_StakingFilterer_WatchDelegate__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.StakingDelegate

// skipping Fuzz_Nosy_StakingFilterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.StakingInitialized

// skipping Fuzz_Nosy_StakingFilterer_WatchOwnershipTransferred__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.StakingOwnershipTransferred

func Fuzz_Nosy_StakingInitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StakingInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_StakingInitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StakingInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_StakingInitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StakingInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_StakingOwnershipTransferredIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StakingOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_StakingOwnershipTransferredIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StakingOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_StakingOwnershipTransferredIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StakingOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

// skipping Fuzz_Nosy_StakingRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_StakingRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_StakingRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingRaw
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Staking == nil || opts == nil {
			return
		}

		_Staking.Transfer(opts)
	})
}

func Fuzz_Nosy_StakingSession_AllowValidators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var validators []common.Address
		fill_err = tp.Fill(&validators)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.AllowValidators(validators)
	})
}

func Fuzz_Nosy_StakingSession_CreateValidator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var pubkey []byte
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.CreateValidator(pubkey)
	})
}

func Fuzz_Nosy_StakingSession_Delegate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var validator common.Address
		fill_err = tp.Fill(&validator)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.Delegate(validator)
	})
}

func Fuzz_Nosy_StakingSession_DisableAllowlist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.DisableAllowlist()
	})
}

func Fuzz_Nosy_StakingSession_DisallowValidators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var validators []common.Address
		fill_err = tp.Fill(&validators)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.DisallowValidators(validators)
	})
}

func Fuzz_Nosy_StakingSession_EnableAllowlist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.EnableAllowlist()
	})
}

func Fuzz_Nosy_StakingSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var owner_ common.Address
		fill_err = tp.Fill(&owner_)
		if fill_err != nil {
			return
		}
		var isAllowlistEnabled_ bool
		fill_err = tp.Fill(&isAllowlistEnabled_)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.Initialize(owner_, isAllowlistEnabled_)
	})
}

func Fuzz_Nosy_StakingSession_IsAllowedValidator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.IsAllowedValidator(arg0)
	})
}

func Fuzz_Nosy_StakingSession_IsAllowlistEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.IsAllowlistEnabled()
	})
}

func Fuzz_Nosy_StakingSession_MinDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.MinDelegation()
	})
}

func Fuzz_Nosy_StakingSession_MinDeposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.MinDeposit()
	})
}

func Fuzz_Nosy_StakingSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.Owner()
	})
}

func Fuzz_Nosy_StakingSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.RenounceOwnership()
	})
}

func Fuzz_Nosy_StakingSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_StakingTransactor_AllowValidators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingTransactor
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var validators []common.Address
		fill_err = tp.Fill(&validators)
		if fill_err != nil {
			return
		}
		if _Staking == nil || opts == nil {
			return
		}

		_Staking.AllowValidators(opts, validators)
	})
}

func Fuzz_Nosy_StakingTransactor_CreateValidator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingTransactor
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var pubkey []byte
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		if _Staking == nil || opts == nil {
			return
		}

		_Staking.CreateValidator(opts, pubkey)
	})
}

func Fuzz_Nosy_StakingTransactor_Delegate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingTransactor
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var validator common.Address
		fill_err = tp.Fill(&validator)
		if fill_err != nil {
			return
		}
		if _Staking == nil || opts == nil {
			return
		}

		_Staking.Delegate(opts, validator)
	})
}

func Fuzz_Nosy_StakingTransactor_DisableAllowlist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingTransactor
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Staking == nil || opts == nil {
			return
		}

		_Staking.DisableAllowlist(opts)
	})
}

func Fuzz_Nosy_StakingTransactor_DisallowValidators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingTransactor
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var validators []common.Address
		fill_err = tp.Fill(&validators)
		if fill_err != nil {
			return
		}
		if _Staking == nil || opts == nil {
			return
		}

		_Staking.DisallowValidators(opts, validators)
	})
}

func Fuzz_Nosy_StakingTransactor_EnableAllowlist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingTransactor
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Staking == nil || opts == nil {
			return
		}

		_Staking.EnableAllowlist(opts)
	})
}

func Fuzz_Nosy_StakingTransactor_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingTransactor
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var owner_ common.Address
		fill_err = tp.Fill(&owner_)
		if fill_err != nil {
			return
		}
		var isAllowlistEnabled_ bool
		fill_err = tp.Fill(&isAllowlistEnabled_)
		if fill_err != nil {
			return
		}
		if _Staking == nil || opts == nil {
			return
		}

		_Staking.Initialize(opts, owner_, isAllowlistEnabled_)
	})
}

func Fuzz_Nosy_StakingTransactor_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingTransactor
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Staking == nil || opts == nil {
			return
		}

		_Staking.RenounceOwnership(opts)
	})
}

func Fuzz_Nosy_StakingTransactor_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingTransactor
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _Staking == nil || opts == nil {
			return
		}

		_Staking.TransferOwnership(opts, newOwner)
	})
}

// skipping Fuzz_Nosy_StakingTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_StakingTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingTransactorRaw
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Staking == nil || opts == nil {
			return
		}

		_Staking.Transfer(opts)
	})
}

func Fuzz_Nosy_StakingTransactorSession_AllowValidators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingTransactorSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var validators []common.Address
		fill_err = tp.Fill(&validators)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.AllowValidators(validators)
	})
}

func Fuzz_Nosy_StakingTransactorSession_CreateValidator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingTransactorSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var pubkey []byte
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.CreateValidator(pubkey)
	})
}

func Fuzz_Nosy_StakingTransactorSession_Delegate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingTransactorSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var validator common.Address
		fill_err = tp.Fill(&validator)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.Delegate(validator)
	})
}

func Fuzz_Nosy_StakingTransactorSession_DisableAllowlist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingTransactorSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.DisableAllowlist()
	})
}

func Fuzz_Nosy_StakingTransactorSession_DisallowValidators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingTransactorSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var validators []common.Address
		fill_err = tp.Fill(&validators)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.DisallowValidators(validators)
	})
}

func Fuzz_Nosy_StakingTransactorSession_EnableAllowlist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingTransactorSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.EnableAllowlist()
	})
}

func Fuzz_Nosy_StakingTransactorSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingTransactorSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var owner_ common.Address
		fill_err = tp.Fill(&owner_)
		if fill_err != nil {
			return
		}
		var isAllowlistEnabled_ bool
		fill_err = tp.Fill(&isAllowlistEnabled_)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.Initialize(owner_, isAllowlistEnabled_)
	})
}

func Fuzz_Nosy_StakingTransactorSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingTransactorSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.RenounceOwnership()
	})
}

func Fuzz_Nosy_StakingTransactorSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Staking *StakingTransactorSession
		fill_err = tp.Fill(&_Staking)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _Staking == nil {
			return
		}

		_Staking.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_StrategyBaseCaller_Explanation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseCaller
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || opts == nil {
			return
		}

		_StrategyBase.Explanation(opts)
	})
}

func Fuzz_Nosy_StrategyBaseCaller_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseCaller
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var index uint8
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || opts == nil {
			return
		}

		_StrategyBase.Paused(opts, index)
	})
}

func Fuzz_Nosy_StrategyBaseCaller_Paused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseCaller
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || opts == nil {
			return
		}

		_StrategyBase.Paused0(opts)
	})
}

func Fuzz_Nosy_StrategyBaseCaller_PauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseCaller
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || opts == nil {
			return
		}

		_StrategyBase.PauserRegistry(opts)
	})
}

func Fuzz_Nosy_StrategyBaseCaller_Shares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseCaller
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var user common.Address
		fill_err = tp.Fill(&user)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || opts == nil {
			return
		}

		_StrategyBase.Shares(opts, user)
	})
}

func Fuzz_Nosy_StrategyBaseCaller_SharesToUnderlying__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseCaller
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var amountShares *big.Int
		fill_err = tp.Fill(&amountShares)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || opts == nil || amountShares == nil {
			return
		}

		_StrategyBase.SharesToUnderlying(opts, amountShares)
	})
}

func Fuzz_Nosy_StrategyBaseCaller_SharesToUnderlyingView__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseCaller
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var amountShares *big.Int
		fill_err = tp.Fill(&amountShares)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || opts == nil || amountShares == nil {
			return
		}

		_StrategyBase.SharesToUnderlyingView(opts, amountShares)
	})
}

func Fuzz_Nosy_StrategyBaseCaller_StrategyManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseCaller
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || opts == nil {
			return
		}

		_StrategyBase.StrategyManager(opts)
	})
}

func Fuzz_Nosy_StrategyBaseCaller_TotalShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseCaller
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || opts == nil {
			return
		}

		_StrategyBase.TotalShares(opts)
	})
}

func Fuzz_Nosy_StrategyBaseCaller_UnderlyingToShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseCaller
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var amountUnderlying *big.Int
		fill_err = tp.Fill(&amountUnderlying)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || opts == nil || amountUnderlying == nil {
			return
		}

		_StrategyBase.UnderlyingToShares(opts, amountUnderlying)
	})
}

func Fuzz_Nosy_StrategyBaseCaller_UnderlyingToSharesView__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseCaller
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var amountUnderlying *big.Int
		fill_err = tp.Fill(&amountUnderlying)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || opts == nil || amountUnderlying == nil {
			return
		}

		_StrategyBase.UnderlyingToSharesView(opts, amountUnderlying)
	})
}

func Fuzz_Nosy_StrategyBaseCaller_UnderlyingToken__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseCaller
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || opts == nil {
			return
		}

		_StrategyBase.UnderlyingToken(opts)
	})
}

func Fuzz_Nosy_StrategyBaseCaller_UserUnderlyingView__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseCaller
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var user common.Address
		fill_err = tp.Fill(&user)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || opts == nil {
			return
		}

		_StrategyBase.UserUnderlyingView(opts, user)
	})
}

// skipping Fuzz_Nosy_StrategyBaseCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_StrategyBaseCallerSession_Explanation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseCallerSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.Explanation()
	})
}

func Fuzz_Nosy_StrategyBaseCallerSession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseCallerSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var index uint8
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.Paused(index)
	})
}

func Fuzz_Nosy_StrategyBaseCallerSession_Paused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseCallerSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.Paused0()
	})
}

func Fuzz_Nosy_StrategyBaseCallerSession_PauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseCallerSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.PauserRegistry()
	})
}

func Fuzz_Nosy_StrategyBaseCallerSession_Shares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseCallerSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var user common.Address
		fill_err = tp.Fill(&user)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.Shares(user)
	})
}

func Fuzz_Nosy_StrategyBaseCallerSession_SharesToUnderlying__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseCallerSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var amountShares *big.Int
		fill_err = tp.Fill(&amountShares)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || amountShares == nil {
			return
		}

		_StrategyBase.SharesToUnderlying(amountShares)
	})
}

func Fuzz_Nosy_StrategyBaseCallerSession_SharesToUnderlyingView__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseCallerSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var amountShares *big.Int
		fill_err = tp.Fill(&amountShares)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || amountShares == nil {
			return
		}

		_StrategyBase.SharesToUnderlyingView(amountShares)
	})
}

func Fuzz_Nosy_StrategyBaseCallerSession_StrategyManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseCallerSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.StrategyManager()
	})
}

func Fuzz_Nosy_StrategyBaseCallerSession_TotalShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseCallerSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.TotalShares()
	})
}

func Fuzz_Nosy_StrategyBaseCallerSession_UnderlyingToShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseCallerSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var amountUnderlying *big.Int
		fill_err = tp.Fill(&amountUnderlying)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || amountUnderlying == nil {
			return
		}

		_StrategyBase.UnderlyingToShares(amountUnderlying)
	})
}

func Fuzz_Nosy_StrategyBaseCallerSession_UnderlyingToSharesView__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseCallerSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var amountUnderlying *big.Int
		fill_err = tp.Fill(&amountUnderlying)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || amountUnderlying == nil {
			return
		}

		_StrategyBase.UnderlyingToSharesView(amountUnderlying)
	})
}

func Fuzz_Nosy_StrategyBaseCallerSession_UnderlyingToken__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseCallerSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.UnderlyingToken()
	})
}

func Fuzz_Nosy_StrategyBaseCallerSession_UserUnderlyingView__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseCallerSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var user common.Address
		fill_err = tp.Fill(&user)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.UserUnderlyingView(user)
	})
}

func Fuzz_Nosy_StrategyBaseFilterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseFilterer
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || opts == nil {
			return
		}

		_StrategyBase.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_StrategyBaseFilterer_FilterPaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseFilterer
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var account []common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || opts == nil {
			return
		}

		_StrategyBase.FilterPaused(opts, account)
	})
}

func Fuzz_Nosy_StrategyBaseFilterer_FilterPauserRegistrySet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseFilterer
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || opts == nil {
			return
		}

		_StrategyBase.FilterPauserRegistrySet(opts)
	})
}

func Fuzz_Nosy_StrategyBaseFilterer_FilterUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseFilterer
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var account []common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || opts == nil {
			return
		}

		_StrategyBase.FilterUnpaused(opts, account)
	})
}

func Fuzz_Nosy_StrategyBaseFilterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseFilterer
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.ParseInitialized(log)
	})
}

func Fuzz_Nosy_StrategyBaseFilterer_ParsePaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseFilterer
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.ParsePaused(log)
	})
}

func Fuzz_Nosy_StrategyBaseFilterer_ParsePauserRegistrySet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseFilterer
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.ParsePauserRegistrySet(log)
	})
}

func Fuzz_Nosy_StrategyBaseFilterer_ParseUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseFilterer
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.ParseUnpaused(log)
	})
}

// skipping Fuzz_Nosy_StrategyBaseFilterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.StrategyBaseInitialized

// skipping Fuzz_Nosy_StrategyBaseFilterer_WatchPaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.StrategyBasePaused

// skipping Fuzz_Nosy_StrategyBaseFilterer_WatchPauserRegistrySet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.StrategyBasePauserRegistrySet

// skipping Fuzz_Nosy_StrategyBaseFilterer_WatchUnpaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.StrategyBaseUnpaused

func Fuzz_Nosy_StrategyBaseInitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyBaseInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_StrategyBaseInitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyBaseInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_StrategyBaseInitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyBaseInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_StrategyBasePausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyBasePausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_StrategyBasePausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyBasePausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_StrategyBasePausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyBasePausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_StrategyBasePauserRegistrySetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyBasePauserRegistrySetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_StrategyBasePauserRegistrySetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyBasePauserRegistrySetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_StrategyBasePauserRegistrySetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyBasePauserRegistrySetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

// skipping Fuzz_Nosy_StrategyBaseRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_StrategyBaseRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_StrategyBaseRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseRaw
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || opts == nil {
			return
		}

		_StrategyBase.Transfer(opts)
	})
}

func Fuzz_Nosy_StrategyBaseSession_Deposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var token common.Address
		fill_err = tp.Fill(&token)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || amount == nil {
			return
		}

		_StrategyBase.Deposit(token, amount)
	})
}

func Fuzz_Nosy_StrategyBaseSession_Explanation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.Explanation()
	})
}

func Fuzz_Nosy_StrategyBaseSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var _underlyingToken common.Address
		fill_err = tp.Fill(&_underlyingToken)
		if fill_err != nil {
			return
		}
		var _pauserRegistry common.Address
		fill_err = tp.Fill(&_pauserRegistry)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.Initialize(_underlyingToken, _pauserRegistry)
	})
}

func Fuzz_Nosy_StrategyBaseSession_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || newPausedStatus == nil {
			return
		}

		_StrategyBase.Pause(newPausedStatus)
	})
}

func Fuzz_Nosy_StrategyBaseSession_PauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.PauseAll()
	})
}

func Fuzz_Nosy_StrategyBaseSession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var index uint8
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.Paused(index)
	})
}

func Fuzz_Nosy_StrategyBaseSession_Paused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.Paused0()
	})
}

func Fuzz_Nosy_StrategyBaseSession_PauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.PauserRegistry()
	})
}

func Fuzz_Nosy_StrategyBaseSession_SetPauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var newPauserRegistry common.Address
		fill_err = tp.Fill(&newPauserRegistry)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.SetPauserRegistry(newPauserRegistry)
	})
}

func Fuzz_Nosy_StrategyBaseSession_Shares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var user common.Address
		fill_err = tp.Fill(&user)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.Shares(user)
	})
}

func Fuzz_Nosy_StrategyBaseSession_SharesToUnderlying__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var amountShares *big.Int
		fill_err = tp.Fill(&amountShares)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || amountShares == nil {
			return
		}

		_StrategyBase.SharesToUnderlying(amountShares)
	})
}

func Fuzz_Nosy_StrategyBaseSession_SharesToUnderlyingView__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var amountShares *big.Int
		fill_err = tp.Fill(&amountShares)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || amountShares == nil {
			return
		}

		_StrategyBase.SharesToUnderlyingView(amountShares)
	})
}

func Fuzz_Nosy_StrategyBaseSession_StrategyManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.StrategyManager()
	})
}

func Fuzz_Nosy_StrategyBaseSession_TotalShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.TotalShares()
	})
}

func Fuzz_Nosy_StrategyBaseSession_UnderlyingToShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var amountUnderlying *big.Int
		fill_err = tp.Fill(&amountUnderlying)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || amountUnderlying == nil {
			return
		}

		_StrategyBase.UnderlyingToShares(amountUnderlying)
	})
}

func Fuzz_Nosy_StrategyBaseSession_UnderlyingToSharesView__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var amountUnderlying *big.Int
		fill_err = tp.Fill(&amountUnderlying)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || amountUnderlying == nil {
			return
		}

		_StrategyBase.UnderlyingToSharesView(amountUnderlying)
	})
}

func Fuzz_Nosy_StrategyBaseSession_UnderlyingToken__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.UnderlyingToken()
	})
}

func Fuzz_Nosy_StrategyBaseSession_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || newPausedStatus == nil {
			return
		}

		_StrategyBase.Unpause(newPausedStatus)
	})
}

func Fuzz_Nosy_StrategyBaseSession_UserUnderlying__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var user common.Address
		fill_err = tp.Fill(&user)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.UserUnderlying(user)
	})
}

func Fuzz_Nosy_StrategyBaseSession_UserUnderlyingView__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var user common.Address
		fill_err = tp.Fill(&user)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.UserUnderlyingView(user)
	})
}

func Fuzz_Nosy_StrategyBaseSession_Withdraw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var recipient common.Address
		fill_err = tp.Fill(&recipient)
		if fill_err != nil {
			return
		}
		var token common.Address
		fill_err = tp.Fill(&token)
		if fill_err != nil {
			return
		}
		var amountShares *big.Int
		fill_err = tp.Fill(&amountShares)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || amountShares == nil {
			return
		}

		_StrategyBase.Withdraw(recipient, token, amountShares)
	})
}

func Fuzz_Nosy_StrategyBaseTransactor_Deposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseTransactor
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var token common.Address
		fill_err = tp.Fill(&token)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || opts == nil || amount == nil {
			return
		}

		_StrategyBase.Deposit(opts, token, amount)
	})
}

func Fuzz_Nosy_StrategyBaseTransactor_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseTransactor
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _underlyingToken common.Address
		fill_err = tp.Fill(&_underlyingToken)
		if fill_err != nil {
			return
		}
		var _pauserRegistry common.Address
		fill_err = tp.Fill(&_pauserRegistry)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || opts == nil {
			return
		}

		_StrategyBase.Initialize(opts, _underlyingToken, _pauserRegistry)
	})
}

func Fuzz_Nosy_StrategyBaseTransactor_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseTransactor
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || opts == nil || newPausedStatus == nil {
			return
		}

		_StrategyBase.Pause(opts, newPausedStatus)
	})
}

func Fuzz_Nosy_StrategyBaseTransactor_PauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseTransactor
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || opts == nil {
			return
		}

		_StrategyBase.PauseAll(opts)
	})
}

func Fuzz_Nosy_StrategyBaseTransactor_SetPauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseTransactor
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newPauserRegistry common.Address
		fill_err = tp.Fill(&newPauserRegistry)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || opts == nil {
			return
		}

		_StrategyBase.SetPauserRegistry(opts, newPauserRegistry)
	})
}

func Fuzz_Nosy_StrategyBaseTransactor_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseTransactor
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || opts == nil || newPausedStatus == nil {
			return
		}

		_StrategyBase.Unpause(opts, newPausedStatus)
	})
}

func Fuzz_Nosy_StrategyBaseTransactor_UserUnderlying__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseTransactor
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var user common.Address
		fill_err = tp.Fill(&user)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || opts == nil {
			return
		}

		_StrategyBase.UserUnderlying(opts, user)
	})
}

func Fuzz_Nosy_StrategyBaseTransactor_Withdraw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseTransactor
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var recipient common.Address
		fill_err = tp.Fill(&recipient)
		if fill_err != nil {
			return
		}
		var token common.Address
		fill_err = tp.Fill(&token)
		if fill_err != nil {
			return
		}
		var amountShares *big.Int
		fill_err = tp.Fill(&amountShares)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || opts == nil || amountShares == nil {
			return
		}

		_StrategyBase.Withdraw(opts, recipient, token, amountShares)
	})
}

// skipping Fuzz_Nosy_StrategyBaseTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_StrategyBaseTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseTransactorRaw
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || opts == nil {
			return
		}

		_StrategyBase.Transfer(opts)
	})
}

func Fuzz_Nosy_StrategyBaseTransactorSession_Deposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseTransactorSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var token common.Address
		fill_err = tp.Fill(&token)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || amount == nil {
			return
		}

		_StrategyBase.Deposit(token, amount)
	})
}

func Fuzz_Nosy_StrategyBaseTransactorSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseTransactorSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var _underlyingToken common.Address
		fill_err = tp.Fill(&_underlyingToken)
		if fill_err != nil {
			return
		}
		var _pauserRegistry common.Address
		fill_err = tp.Fill(&_pauserRegistry)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.Initialize(_underlyingToken, _pauserRegistry)
	})
}

func Fuzz_Nosy_StrategyBaseTransactorSession_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseTransactorSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || newPausedStatus == nil {
			return
		}

		_StrategyBase.Pause(newPausedStatus)
	})
}

func Fuzz_Nosy_StrategyBaseTransactorSession_PauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseTransactorSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.PauseAll()
	})
}

func Fuzz_Nosy_StrategyBaseTransactorSession_SetPauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseTransactorSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var newPauserRegistry common.Address
		fill_err = tp.Fill(&newPauserRegistry)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.SetPauserRegistry(newPauserRegistry)
	})
}

func Fuzz_Nosy_StrategyBaseTransactorSession_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseTransactorSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || newPausedStatus == nil {
			return
		}

		_StrategyBase.Unpause(newPausedStatus)
	})
}

func Fuzz_Nosy_StrategyBaseTransactorSession_UserUnderlying__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseTransactorSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var user common.Address
		fill_err = tp.Fill(&user)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil {
			return
		}

		_StrategyBase.UserUnderlying(user)
	})
}

func Fuzz_Nosy_StrategyBaseTransactorSession_Withdraw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyBase *StrategyBaseTransactorSession
		fill_err = tp.Fill(&_StrategyBase)
		if fill_err != nil {
			return
		}
		var recipient common.Address
		fill_err = tp.Fill(&recipient)
		if fill_err != nil {
			return
		}
		var token common.Address
		fill_err = tp.Fill(&token)
		if fill_err != nil {
			return
		}
		var amountShares *big.Int
		fill_err = tp.Fill(&amountShares)
		if fill_err != nil {
			return
		}
		if _StrategyBase == nil || amountShares == nil {
			return
		}

		_StrategyBase.Withdraw(recipient, token, amountShares)
	})
}

func Fuzz_Nosy_StrategyBaseUnpausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyBaseUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_StrategyBaseUnpausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyBaseUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_StrategyBaseUnpausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyBaseUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_StrategyManagerCaller_CalculateWithdrawalRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCaller
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var queuedWithdrawal IStrategyManagerDeprecatedStructQueuedWithdrawal
		fill_err = tp.Fill(&queuedWithdrawal)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.CalculateWithdrawalRoot(opts, queuedWithdrawal)
	})
}

func Fuzz_Nosy_StrategyManagerCaller_DEPOSITTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCaller
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.DEPOSITTYPEHASH(opts)
	})
}

func Fuzz_Nosy_StrategyManagerCaller_DOMAINTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCaller
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.DOMAINTYPEHASH(opts)
	})
}

func Fuzz_Nosy_StrategyManagerCaller_Delegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCaller
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.Delegation(opts)
	})
}

func Fuzz_Nosy_StrategyManagerCaller_DomainSeparator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCaller
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.DomainSeparator(opts)
	})
}

func Fuzz_Nosy_StrategyManagerCaller_EigenPodManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCaller
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.EigenPodManager(opts)
	})
}

func Fuzz_Nosy_StrategyManagerCaller_GetDeposits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCaller
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.GetDeposits(opts, staker)
	})
}

func Fuzz_Nosy_StrategyManagerCaller_Nonces__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCaller
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.Nonces(opts, arg0)
	})
}

func Fuzz_Nosy_StrategyManagerCaller_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCaller
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.Owner(opts)
	})
}

func Fuzz_Nosy_StrategyManagerCaller_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCaller
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var index uint8
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.Paused(opts, index)
	})
}

func Fuzz_Nosy_StrategyManagerCaller_Paused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCaller
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.Paused0(opts)
	})
}

func Fuzz_Nosy_StrategyManagerCaller_PauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCaller
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.PauserRegistry(opts)
	})
}

func Fuzz_Nosy_StrategyManagerCaller_Slasher__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCaller
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.Slasher(opts)
	})
}

func Fuzz_Nosy_StrategyManagerCaller_StakerStrategyList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCaller
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 *big.Int
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil || arg1 == nil {
			return
		}

		_StrategyManager.StakerStrategyList(opts, arg0, arg1)
	})
}

func Fuzz_Nosy_StrategyManagerCaller_StakerStrategyListLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCaller
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.StakerStrategyListLength(opts, staker)
	})
}

func Fuzz_Nosy_StrategyManagerCaller_StakerStrategyShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCaller
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 common.Address
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.StakerStrategyShares(opts, arg0, arg1)
	})
}

func Fuzz_Nosy_StrategyManagerCaller_StrategyIsWhitelistedForDeposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCaller
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.StrategyIsWhitelistedForDeposit(opts, arg0)
	})
}

func Fuzz_Nosy_StrategyManagerCaller_StrategyWhitelister__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCaller
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.StrategyWhitelister(opts)
	})
}

func Fuzz_Nosy_StrategyManagerCaller_ThirdPartyTransfersForbidden__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCaller
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.ThirdPartyTransfersForbidden(opts, arg0)
	})
}

func Fuzz_Nosy_StrategyManagerCaller_WithdrawalRootPending__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCaller
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.WithdrawalRootPending(opts, arg0)
	})
}

// skipping Fuzz_Nosy_StrategyManagerCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_StrategyManagerCallerSession_CalculateWithdrawalRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCallerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var queuedWithdrawal IStrategyManagerDeprecatedStructQueuedWithdrawal
		fill_err = tp.Fill(&queuedWithdrawal)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.CalculateWithdrawalRoot(queuedWithdrawal)
	})
}

func Fuzz_Nosy_StrategyManagerCallerSession_DEPOSITTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCallerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.DEPOSITTYPEHASH()
	})
}

func Fuzz_Nosy_StrategyManagerCallerSession_DOMAINTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCallerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.DOMAINTYPEHASH()
	})
}

func Fuzz_Nosy_StrategyManagerCallerSession_Delegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCallerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.Delegation()
	})
}

func Fuzz_Nosy_StrategyManagerCallerSession_DomainSeparator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCallerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.DomainSeparator()
	})
}

func Fuzz_Nosy_StrategyManagerCallerSession_EigenPodManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCallerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.EigenPodManager()
	})
}

func Fuzz_Nosy_StrategyManagerCallerSession_GetDeposits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCallerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.GetDeposits(staker)
	})
}

func Fuzz_Nosy_StrategyManagerCallerSession_Nonces__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCallerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.Nonces(arg0)
	})
}

func Fuzz_Nosy_StrategyManagerCallerSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCallerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.Owner()
	})
}

func Fuzz_Nosy_StrategyManagerCallerSession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCallerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var index uint8
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.Paused(index)
	})
}

func Fuzz_Nosy_StrategyManagerCallerSession_Paused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCallerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.Paused0()
	})
}

func Fuzz_Nosy_StrategyManagerCallerSession_PauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCallerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.PauserRegistry()
	})
}

func Fuzz_Nosy_StrategyManagerCallerSession_Slasher__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCallerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.Slasher()
	})
}

func Fuzz_Nosy_StrategyManagerCallerSession_StakerStrategyList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCallerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 *big.Int
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || arg1 == nil {
			return
		}

		_StrategyManager.StakerStrategyList(arg0, arg1)
	})
}

func Fuzz_Nosy_StrategyManagerCallerSession_StakerStrategyListLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCallerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.StakerStrategyListLength(staker)
	})
}

func Fuzz_Nosy_StrategyManagerCallerSession_StakerStrategyShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCallerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 common.Address
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.StakerStrategyShares(arg0, arg1)
	})
}

func Fuzz_Nosy_StrategyManagerCallerSession_StrategyIsWhitelistedForDeposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCallerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.StrategyIsWhitelistedForDeposit(arg0)
	})
}

func Fuzz_Nosy_StrategyManagerCallerSession_StrategyWhitelister__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCallerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.StrategyWhitelister()
	})
}

func Fuzz_Nosy_StrategyManagerCallerSession_ThirdPartyTransfersForbidden__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCallerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.ThirdPartyTransfersForbidden(arg0)
	})
}

func Fuzz_Nosy_StrategyManagerCallerSession_WithdrawalRootPending__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerCallerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.WithdrawalRootPending(arg0)
	})
}

func Fuzz_Nosy_StrategyManagerDepositIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerDepositIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_StrategyManagerDepositIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerDepositIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_StrategyManagerDepositIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerDepositIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_StrategyManagerFilterer_FilterDeposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerFilterer
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.FilterDeposit(opts)
	})
}

func Fuzz_Nosy_StrategyManagerFilterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerFilterer
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_StrategyManagerFilterer_FilterOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerFilterer
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var previousOwner []common.Address
		fill_err = tp.Fill(&previousOwner)
		if fill_err != nil {
			return
		}
		var newOwner []common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.FilterOwnershipTransferred(opts, previousOwner, newOwner)
	})
}

func Fuzz_Nosy_StrategyManagerFilterer_FilterPaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerFilterer
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var account []common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.FilterPaused(opts, account)
	})
}

func Fuzz_Nosy_StrategyManagerFilterer_FilterPauserRegistrySet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerFilterer
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.FilterPauserRegistrySet(opts)
	})
}

func Fuzz_Nosy_StrategyManagerFilterer_FilterStrategyAddedToDepositWhitelist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerFilterer
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.FilterStrategyAddedToDepositWhitelist(opts)
	})
}

func Fuzz_Nosy_StrategyManagerFilterer_FilterStrategyRemovedFromDepositWhitelist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerFilterer
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.FilterStrategyRemovedFromDepositWhitelist(opts)
	})
}

func Fuzz_Nosy_StrategyManagerFilterer_FilterStrategyWhitelisterChanged__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerFilterer
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.FilterStrategyWhitelisterChanged(opts)
	})
}

func Fuzz_Nosy_StrategyManagerFilterer_FilterUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerFilterer
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var account []common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.FilterUnpaused(opts, account)
	})
}

func Fuzz_Nosy_StrategyManagerFilterer_FilterUpdatedThirdPartyTransfersForbidden__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerFilterer
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.FilterUpdatedThirdPartyTransfersForbidden(opts)
	})
}

func Fuzz_Nosy_StrategyManagerFilterer_ParseDeposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerFilterer
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.ParseDeposit(log)
	})
}

func Fuzz_Nosy_StrategyManagerFilterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerFilterer
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.ParseInitialized(log)
	})
}

func Fuzz_Nosy_StrategyManagerFilterer_ParseOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerFilterer
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.ParseOwnershipTransferred(log)
	})
}

func Fuzz_Nosy_StrategyManagerFilterer_ParsePaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerFilterer
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.ParsePaused(log)
	})
}

func Fuzz_Nosy_StrategyManagerFilterer_ParsePauserRegistrySet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerFilterer
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.ParsePauserRegistrySet(log)
	})
}

func Fuzz_Nosy_StrategyManagerFilterer_ParseStrategyAddedToDepositWhitelist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerFilterer
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.ParseStrategyAddedToDepositWhitelist(log)
	})
}

func Fuzz_Nosy_StrategyManagerFilterer_ParseStrategyRemovedFromDepositWhitelist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerFilterer
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.ParseStrategyRemovedFromDepositWhitelist(log)
	})
}

func Fuzz_Nosy_StrategyManagerFilterer_ParseStrategyWhitelisterChanged__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerFilterer
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.ParseStrategyWhitelisterChanged(log)
	})
}

func Fuzz_Nosy_StrategyManagerFilterer_ParseUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerFilterer
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.ParseUnpaused(log)
	})
}

func Fuzz_Nosy_StrategyManagerFilterer_ParseUpdatedThirdPartyTransfersForbidden__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerFilterer
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.ParseUpdatedThirdPartyTransfersForbidden(log)
	})
}

// skipping Fuzz_Nosy_StrategyManagerFilterer_WatchDeposit__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.StrategyManagerDeposit

// skipping Fuzz_Nosy_StrategyManagerFilterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.StrategyManagerInitialized

// skipping Fuzz_Nosy_StrategyManagerFilterer_WatchOwnershipTransferred__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.StrategyManagerOwnershipTransferred

// skipping Fuzz_Nosy_StrategyManagerFilterer_WatchPaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.StrategyManagerPaused

// skipping Fuzz_Nosy_StrategyManagerFilterer_WatchPauserRegistrySet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.StrategyManagerPauserRegistrySet

// skipping Fuzz_Nosy_StrategyManagerFilterer_WatchStrategyAddedToDepositWhitelist__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.StrategyManagerStrategyAddedToDepositWhitelist

// skipping Fuzz_Nosy_StrategyManagerFilterer_WatchStrategyRemovedFromDepositWhitelist__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.StrategyManagerStrategyRemovedFromDepositWhitelist

// skipping Fuzz_Nosy_StrategyManagerFilterer_WatchStrategyWhitelisterChanged__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.StrategyManagerStrategyWhitelisterChanged

// skipping Fuzz_Nosy_StrategyManagerFilterer_WatchUnpaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.StrategyManagerUnpaused

// skipping Fuzz_Nosy_StrategyManagerFilterer_WatchUpdatedThirdPartyTransfersForbidden__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.StrategyManagerUpdatedThirdPartyTransfersForbidden

func Fuzz_Nosy_StrategyManagerInitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_StrategyManagerInitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_StrategyManagerInitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_StrategyManagerOwnershipTransferredIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_StrategyManagerOwnershipTransferredIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_StrategyManagerOwnershipTransferredIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_StrategyManagerPausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerPausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_StrategyManagerPausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerPausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_StrategyManagerPausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerPausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_StrategyManagerPauserRegistrySetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerPauserRegistrySetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_StrategyManagerPauserRegistrySetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerPauserRegistrySetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_StrategyManagerPauserRegistrySetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerPauserRegistrySetIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

// skipping Fuzz_Nosy_StrategyManagerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_StrategyManagerRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_StrategyManagerRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerRaw
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.Transfer(opts)
	})
}

func Fuzz_Nosy_StrategyManagerSession_AddShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var token common.Address
		fill_err = tp.Fill(&token)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var shares *big.Int
		fill_err = tp.Fill(&shares)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || shares == nil {
			return
		}

		_StrategyManager.AddShares(staker, token, strategy, shares)
	})
}

func Fuzz_Nosy_StrategyManagerSession_AddStrategiesToDepositWhitelist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var strategiesToWhitelist []common.Address
		fill_err = tp.Fill(&strategiesToWhitelist)
		if fill_err != nil {
			return
		}
		var thirdPartyTransfersForbiddenValues []bool
		fill_err = tp.Fill(&thirdPartyTransfersForbiddenValues)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.AddStrategiesToDepositWhitelist(strategiesToWhitelist, thirdPartyTransfersForbiddenValues)
	})
}

func Fuzz_Nosy_StrategyManagerSession_CalculateWithdrawalRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var queuedWithdrawal IStrategyManagerDeprecatedStructQueuedWithdrawal
		fill_err = tp.Fill(&queuedWithdrawal)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.CalculateWithdrawalRoot(queuedWithdrawal)
	})
}

func Fuzz_Nosy_StrategyManagerSession_DEPOSITTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.DEPOSITTYPEHASH()
	})
}

func Fuzz_Nosy_StrategyManagerSession_DOMAINTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.DOMAINTYPEHASH()
	})
}

func Fuzz_Nosy_StrategyManagerSession_Delegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.Delegation()
	})
}

func Fuzz_Nosy_StrategyManagerSession_DepositIntoStrategy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var token common.Address
		fill_err = tp.Fill(&token)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || amount == nil {
			return
		}

		_StrategyManager.DepositIntoStrategy(strategy, token, amount)
	})
}

func Fuzz_Nosy_StrategyManagerSession_DepositIntoStrategyWithSignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var token common.Address
		fill_err = tp.Fill(&token)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var expiry *big.Int
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		var signature []byte
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || amount == nil || expiry == nil {
			return
		}

		_StrategyManager.DepositIntoStrategyWithSignature(strategy, token, amount, staker, expiry, signature)
	})
}

func Fuzz_Nosy_StrategyManagerSession_DomainSeparator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.DomainSeparator()
	})
}

func Fuzz_Nosy_StrategyManagerSession_EigenPodManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.EigenPodManager()
	})
}

func Fuzz_Nosy_StrategyManagerSession_GetDeposits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.GetDeposits(staker)
	})
}

func Fuzz_Nosy_StrategyManagerSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var initialOwner common.Address
		fill_err = tp.Fill(&initialOwner)
		if fill_err != nil {
			return
		}
		var initialStrategyWhitelister common.Address
		fill_err = tp.Fill(&initialStrategyWhitelister)
		if fill_err != nil {
			return
		}
		var _pauserRegistry common.Address
		fill_err = tp.Fill(&_pauserRegistry)
		if fill_err != nil {
			return
		}
		var initialPausedStatus *big.Int
		fill_err = tp.Fill(&initialPausedStatus)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || initialPausedStatus == nil {
			return
		}

		_StrategyManager.Initialize(initialOwner, initialStrategyWhitelister, _pauserRegistry, initialPausedStatus)
	})
}

func Fuzz_Nosy_StrategyManagerSession_MigrateQueuedWithdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var queuedWithdrawal IStrategyManagerDeprecatedStructQueuedWithdrawal
		fill_err = tp.Fill(&queuedWithdrawal)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.MigrateQueuedWithdrawal(queuedWithdrawal)
	})
}

func Fuzz_Nosy_StrategyManagerSession_Nonces__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.Nonces(arg0)
	})
}

func Fuzz_Nosy_StrategyManagerSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.Owner()
	})
}

func Fuzz_Nosy_StrategyManagerSession_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || newPausedStatus == nil {
			return
		}

		_StrategyManager.Pause(newPausedStatus)
	})
}

func Fuzz_Nosy_StrategyManagerSession_PauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.PauseAll()
	})
}

func Fuzz_Nosy_StrategyManagerSession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var index uint8
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.Paused(index)
	})
}

func Fuzz_Nosy_StrategyManagerSession_Paused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.Paused0()
	})
}

func Fuzz_Nosy_StrategyManagerSession_PauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.PauserRegistry()
	})
}

func Fuzz_Nosy_StrategyManagerSession_RemoveShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var shares *big.Int
		fill_err = tp.Fill(&shares)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || shares == nil {
			return
		}

		_StrategyManager.RemoveShares(staker, strategy, shares)
	})
}

func Fuzz_Nosy_StrategyManagerSession_RemoveStrategiesFromDepositWhitelist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var strategiesToRemoveFromWhitelist []common.Address
		fill_err = tp.Fill(&strategiesToRemoveFromWhitelist)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.RemoveStrategiesFromDepositWhitelist(strategiesToRemoveFromWhitelist)
	})
}

func Fuzz_Nosy_StrategyManagerSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.RenounceOwnership()
	})
}

func Fuzz_Nosy_StrategyManagerSession_SetPauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var newPauserRegistry common.Address
		fill_err = tp.Fill(&newPauserRegistry)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.SetPauserRegistry(newPauserRegistry)
	})
}

func Fuzz_Nosy_StrategyManagerSession_SetStrategyWhitelister__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var newStrategyWhitelister common.Address
		fill_err = tp.Fill(&newStrategyWhitelister)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.SetStrategyWhitelister(newStrategyWhitelister)
	})
}

func Fuzz_Nosy_StrategyManagerSession_SetThirdPartyTransfersForbidden__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var value bool
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.SetThirdPartyTransfersForbidden(strategy, value)
	})
}

func Fuzz_Nosy_StrategyManagerSession_Slasher__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.Slasher()
	})
}

func Fuzz_Nosy_StrategyManagerSession_StakerStrategyList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 *big.Int
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || arg1 == nil {
			return
		}

		_StrategyManager.StakerStrategyList(arg0, arg1)
	})
}

func Fuzz_Nosy_StrategyManagerSession_StakerStrategyListLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.StakerStrategyListLength(staker)
	})
}

func Fuzz_Nosy_StrategyManagerSession_StakerStrategyShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 common.Address
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.StakerStrategyShares(arg0, arg1)
	})
}

func Fuzz_Nosy_StrategyManagerSession_StrategyIsWhitelistedForDeposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.StrategyIsWhitelistedForDeposit(arg0)
	})
}

func Fuzz_Nosy_StrategyManagerSession_StrategyWhitelister__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.StrategyWhitelister()
	})
}

func Fuzz_Nosy_StrategyManagerSession_ThirdPartyTransfersForbidden__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.ThirdPartyTransfersForbidden(arg0)
	})
}

func Fuzz_Nosy_StrategyManagerSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_StrategyManagerSession_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || newPausedStatus == nil {
			return
		}

		_StrategyManager.Unpause(newPausedStatus)
	})
}

func Fuzz_Nosy_StrategyManagerSession_WithdrawSharesAsTokens__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var recipient common.Address
		fill_err = tp.Fill(&recipient)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var shares *big.Int
		fill_err = tp.Fill(&shares)
		if fill_err != nil {
			return
		}
		var token common.Address
		fill_err = tp.Fill(&token)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || shares == nil {
			return
		}

		_StrategyManager.WithdrawSharesAsTokens(recipient, strategy, shares, token)
	})
}

func Fuzz_Nosy_StrategyManagerSession_WithdrawalRootPending__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.WithdrawalRootPending(arg0)
	})
}

func Fuzz_Nosy_StrategyManagerStrategyAddedToDepositWhitelistIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerStrategyAddedToDepositWhitelistIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_StrategyManagerStrategyAddedToDepositWhitelistIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerStrategyAddedToDepositWhitelistIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_StrategyManagerStrategyAddedToDepositWhitelistIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerStrategyAddedToDepositWhitelistIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_StrategyManagerStrategyRemovedFromDepositWhitelistIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerStrategyRemovedFromDepositWhitelistIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_StrategyManagerStrategyRemovedFromDepositWhitelistIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerStrategyRemovedFromDepositWhitelistIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_StrategyManagerStrategyRemovedFromDepositWhitelistIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerStrategyRemovedFromDepositWhitelistIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_StrategyManagerStrategyWhitelisterChangedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerStrategyWhitelisterChangedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_StrategyManagerStrategyWhitelisterChangedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerStrategyWhitelisterChangedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_StrategyManagerStrategyWhitelisterChangedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerStrategyWhitelisterChangedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_StrategyManagerTransactor_AddShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactor
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var token common.Address
		fill_err = tp.Fill(&token)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var shares *big.Int
		fill_err = tp.Fill(&shares)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil || shares == nil {
			return
		}

		_StrategyManager.AddShares(opts, staker, token, strategy, shares)
	})
}

func Fuzz_Nosy_StrategyManagerTransactor_AddStrategiesToDepositWhitelist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactor
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var strategiesToWhitelist []common.Address
		fill_err = tp.Fill(&strategiesToWhitelist)
		if fill_err != nil {
			return
		}
		var thirdPartyTransfersForbiddenValues []bool
		fill_err = tp.Fill(&thirdPartyTransfersForbiddenValues)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.AddStrategiesToDepositWhitelist(opts, strategiesToWhitelist, thirdPartyTransfersForbiddenValues)
	})
}

func Fuzz_Nosy_StrategyManagerTransactor_DepositIntoStrategy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactor
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var token common.Address
		fill_err = tp.Fill(&token)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil || amount == nil {
			return
		}

		_StrategyManager.DepositIntoStrategy(opts, strategy, token, amount)
	})
}

func Fuzz_Nosy_StrategyManagerTransactor_DepositIntoStrategyWithSignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactor
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var token common.Address
		fill_err = tp.Fill(&token)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var expiry *big.Int
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		var signature []byte
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil || amount == nil || expiry == nil {
			return
		}

		_StrategyManager.DepositIntoStrategyWithSignature(opts, strategy, token, amount, staker, expiry, signature)
	})
}

func Fuzz_Nosy_StrategyManagerTransactor_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactor
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var initialOwner common.Address
		fill_err = tp.Fill(&initialOwner)
		if fill_err != nil {
			return
		}
		var initialStrategyWhitelister common.Address
		fill_err = tp.Fill(&initialStrategyWhitelister)
		if fill_err != nil {
			return
		}
		var _pauserRegistry common.Address
		fill_err = tp.Fill(&_pauserRegistry)
		if fill_err != nil {
			return
		}
		var initialPausedStatus *big.Int
		fill_err = tp.Fill(&initialPausedStatus)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil || initialPausedStatus == nil {
			return
		}

		_StrategyManager.Initialize(opts, initialOwner, initialStrategyWhitelister, _pauserRegistry, initialPausedStatus)
	})
}

func Fuzz_Nosy_StrategyManagerTransactor_MigrateQueuedWithdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactor
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var queuedWithdrawal IStrategyManagerDeprecatedStructQueuedWithdrawal
		fill_err = tp.Fill(&queuedWithdrawal)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.MigrateQueuedWithdrawal(opts, queuedWithdrawal)
	})
}

func Fuzz_Nosy_StrategyManagerTransactor_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactor
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil || newPausedStatus == nil {
			return
		}

		_StrategyManager.Pause(opts, newPausedStatus)
	})
}

func Fuzz_Nosy_StrategyManagerTransactor_PauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactor
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.PauseAll(opts)
	})
}

func Fuzz_Nosy_StrategyManagerTransactor_RemoveShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactor
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var shares *big.Int
		fill_err = tp.Fill(&shares)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil || shares == nil {
			return
		}

		_StrategyManager.RemoveShares(opts, staker, strategy, shares)
	})
}

func Fuzz_Nosy_StrategyManagerTransactor_RemoveStrategiesFromDepositWhitelist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactor
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var strategiesToRemoveFromWhitelist []common.Address
		fill_err = tp.Fill(&strategiesToRemoveFromWhitelist)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.RemoveStrategiesFromDepositWhitelist(opts, strategiesToRemoveFromWhitelist)
	})
}

func Fuzz_Nosy_StrategyManagerTransactor_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactor
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.RenounceOwnership(opts)
	})
}

func Fuzz_Nosy_StrategyManagerTransactor_SetPauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactor
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newPauserRegistry common.Address
		fill_err = tp.Fill(&newPauserRegistry)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.SetPauserRegistry(opts, newPauserRegistry)
	})
}

func Fuzz_Nosy_StrategyManagerTransactor_SetStrategyWhitelister__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactor
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newStrategyWhitelister common.Address
		fill_err = tp.Fill(&newStrategyWhitelister)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.SetStrategyWhitelister(opts, newStrategyWhitelister)
	})
}

func Fuzz_Nosy_StrategyManagerTransactor_SetThirdPartyTransfersForbidden__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactor
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var value bool
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.SetThirdPartyTransfersForbidden(opts, strategy, value)
	})
}

func Fuzz_Nosy_StrategyManagerTransactor_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactor
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.TransferOwnership(opts, newOwner)
	})
}

func Fuzz_Nosy_StrategyManagerTransactor_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactor
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil || newPausedStatus == nil {
			return
		}

		_StrategyManager.Unpause(opts, newPausedStatus)
	})
}

func Fuzz_Nosy_StrategyManagerTransactor_WithdrawSharesAsTokens__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactor
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var recipient common.Address
		fill_err = tp.Fill(&recipient)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var shares *big.Int
		fill_err = tp.Fill(&shares)
		if fill_err != nil {
			return
		}
		var token common.Address
		fill_err = tp.Fill(&token)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil || shares == nil {
			return
		}

		_StrategyManager.WithdrawSharesAsTokens(opts, recipient, strategy, shares, token)
	})
}

// skipping Fuzz_Nosy_StrategyManagerTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_StrategyManagerTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactorRaw
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || opts == nil {
			return
		}

		_StrategyManager.Transfer(opts)
	})
}

func Fuzz_Nosy_StrategyManagerTransactorSession_AddShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactorSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var token common.Address
		fill_err = tp.Fill(&token)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var shares *big.Int
		fill_err = tp.Fill(&shares)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || shares == nil {
			return
		}

		_StrategyManager.AddShares(staker, token, strategy, shares)
	})
}

func Fuzz_Nosy_StrategyManagerTransactorSession_AddStrategiesToDepositWhitelist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactorSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var strategiesToWhitelist []common.Address
		fill_err = tp.Fill(&strategiesToWhitelist)
		if fill_err != nil {
			return
		}
		var thirdPartyTransfersForbiddenValues []bool
		fill_err = tp.Fill(&thirdPartyTransfersForbiddenValues)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.AddStrategiesToDepositWhitelist(strategiesToWhitelist, thirdPartyTransfersForbiddenValues)
	})
}

func Fuzz_Nosy_StrategyManagerTransactorSession_DepositIntoStrategy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactorSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var token common.Address
		fill_err = tp.Fill(&token)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || amount == nil {
			return
		}

		_StrategyManager.DepositIntoStrategy(strategy, token, amount)
	})
}

func Fuzz_Nosy_StrategyManagerTransactorSession_DepositIntoStrategyWithSignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactorSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var token common.Address
		fill_err = tp.Fill(&token)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var expiry *big.Int
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		var signature []byte
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || amount == nil || expiry == nil {
			return
		}

		_StrategyManager.DepositIntoStrategyWithSignature(strategy, token, amount, staker, expiry, signature)
	})
}

func Fuzz_Nosy_StrategyManagerTransactorSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactorSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var initialOwner common.Address
		fill_err = tp.Fill(&initialOwner)
		if fill_err != nil {
			return
		}
		var initialStrategyWhitelister common.Address
		fill_err = tp.Fill(&initialStrategyWhitelister)
		if fill_err != nil {
			return
		}
		var _pauserRegistry common.Address
		fill_err = tp.Fill(&_pauserRegistry)
		if fill_err != nil {
			return
		}
		var initialPausedStatus *big.Int
		fill_err = tp.Fill(&initialPausedStatus)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || initialPausedStatus == nil {
			return
		}

		_StrategyManager.Initialize(initialOwner, initialStrategyWhitelister, _pauserRegistry, initialPausedStatus)
	})
}

func Fuzz_Nosy_StrategyManagerTransactorSession_MigrateQueuedWithdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactorSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var queuedWithdrawal IStrategyManagerDeprecatedStructQueuedWithdrawal
		fill_err = tp.Fill(&queuedWithdrawal)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.MigrateQueuedWithdrawal(queuedWithdrawal)
	})
}

func Fuzz_Nosy_StrategyManagerTransactorSession_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactorSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || newPausedStatus == nil {
			return
		}

		_StrategyManager.Pause(newPausedStatus)
	})
}

func Fuzz_Nosy_StrategyManagerTransactorSession_PauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactorSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.PauseAll()
	})
}

func Fuzz_Nosy_StrategyManagerTransactorSession_RemoveShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactorSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var shares *big.Int
		fill_err = tp.Fill(&shares)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || shares == nil {
			return
		}

		_StrategyManager.RemoveShares(staker, strategy, shares)
	})
}

func Fuzz_Nosy_StrategyManagerTransactorSession_RemoveStrategiesFromDepositWhitelist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactorSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var strategiesToRemoveFromWhitelist []common.Address
		fill_err = tp.Fill(&strategiesToRemoveFromWhitelist)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.RemoveStrategiesFromDepositWhitelist(strategiesToRemoveFromWhitelist)
	})
}

func Fuzz_Nosy_StrategyManagerTransactorSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactorSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.RenounceOwnership()
	})
}

func Fuzz_Nosy_StrategyManagerTransactorSession_SetPauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactorSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var newPauserRegistry common.Address
		fill_err = tp.Fill(&newPauserRegistry)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.SetPauserRegistry(newPauserRegistry)
	})
}

func Fuzz_Nosy_StrategyManagerTransactorSession_SetStrategyWhitelister__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactorSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var newStrategyWhitelister common.Address
		fill_err = tp.Fill(&newStrategyWhitelister)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.SetStrategyWhitelister(newStrategyWhitelister)
	})
}

func Fuzz_Nosy_StrategyManagerTransactorSession_SetThirdPartyTransfersForbidden__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactorSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var value bool
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.SetThirdPartyTransfersForbidden(strategy, value)
	})
}

func Fuzz_Nosy_StrategyManagerTransactorSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactorSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil {
			return
		}

		_StrategyManager.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_StrategyManagerTransactorSession_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactorSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || newPausedStatus == nil {
			return
		}

		_StrategyManager.Unpause(newPausedStatus)
	})
}

func Fuzz_Nosy_StrategyManagerTransactorSession_WithdrawSharesAsTokens__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _StrategyManager *StrategyManagerTransactorSession
		fill_err = tp.Fill(&_StrategyManager)
		if fill_err != nil {
			return
		}
		var recipient common.Address
		fill_err = tp.Fill(&recipient)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var shares *big.Int
		fill_err = tp.Fill(&shares)
		if fill_err != nil {
			return
		}
		var token common.Address
		fill_err = tp.Fill(&token)
		if fill_err != nil {
			return
		}
		if _StrategyManager == nil || shares == nil {
			return
		}

		_StrategyManager.WithdrawSharesAsTokens(recipient, strategy, shares, token)
	})
}

func Fuzz_Nosy_StrategyManagerUnpausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_StrategyManagerUnpausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_StrategyManagerUnpausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerUnpausedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_StrategyManagerUpdatedThirdPartyTransfersForbiddenIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerUpdatedThirdPartyTransfersForbiddenIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_StrategyManagerUpdatedThirdPartyTransfersForbiddenIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerUpdatedThirdPartyTransfersForbiddenIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_StrategyManagerUpdatedThirdPartyTransfersForbiddenIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *StrategyManagerUpdatedThirdPartyTransfersForbiddenIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_TransparentUpgradeableProxyAdminChangedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *TransparentUpgradeableProxyAdminChangedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_TransparentUpgradeableProxyAdminChangedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *TransparentUpgradeableProxyAdminChangedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_TransparentUpgradeableProxyAdminChangedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *TransparentUpgradeableProxyAdminChangedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

// skipping Fuzz_Nosy_TransparentUpgradeableProxyCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_TransparentUpgradeableProxyFilterer_FilterAdminChanged__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer
		fill_err = tp.Fill(&_TransparentUpgradeableProxy)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _TransparentUpgradeableProxy == nil || opts == nil {
			return
		}

		_TransparentUpgradeableProxy.FilterAdminChanged(opts)
	})
}

func Fuzz_Nosy_TransparentUpgradeableProxyFilterer_FilterUpgraded__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer
		fill_err = tp.Fill(&_TransparentUpgradeableProxy)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var implementation []common.Address
		fill_err = tp.Fill(&implementation)
		if fill_err != nil {
			return
		}
		if _TransparentUpgradeableProxy == nil || opts == nil {
			return
		}

		_TransparentUpgradeableProxy.FilterUpgraded(opts, implementation)
	})
}

func Fuzz_Nosy_TransparentUpgradeableProxyFilterer_ParseAdminChanged__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer
		fill_err = tp.Fill(&_TransparentUpgradeableProxy)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _TransparentUpgradeableProxy == nil {
			return
		}

		_TransparentUpgradeableProxy.ParseAdminChanged(log)
	})
}

func Fuzz_Nosy_TransparentUpgradeableProxyFilterer_ParseUpgraded__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer
		fill_err = tp.Fill(&_TransparentUpgradeableProxy)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _TransparentUpgradeableProxy == nil {
			return
		}

		_TransparentUpgradeableProxy.ParseUpgraded(log)
	})
}

// skipping Fuzz_Nosy_TransparentUpgradeableProxyFilterer_WatchAdminChanged__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.TransparentUpgradeableProxyAdminChanged

// skipping Fuzz_Nosy_TransparentUpgradeableProxyFilterer_WatchUpgraded__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.TransparentUpgradeableProxyUpgraded

// skipping Fuzz_Nosy_TransparentUpgradeableProxyRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_TransparentUpgradeableProxyRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_TransparentUpgradeableProxyRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _TransparentUpgradeableProxy *TransparentUpgradeableProxyRaw
		fill_err = tp.Fill(&_TransparentUpgradeableProxy)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _TransparentUpgradeableProxy == nil || opts == nil {
			return
		}

		_TransparentUpgradeableProxy.Transfer(opts)
	})
}

func Fuzz_Nosy_TransparentUpgradeableProxySession_Fallback__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _TransparentUpgradeableProxy *TransparentUpgradeableProxySession
		fill_err = tp.Fill(&_TransparentUpgradeableProxy)
		if fill_err != nil {
			return
		}
		var calldata []byte
		fill_err = tp.Fill(&calldata)
		if fill_err != nil {
			return
		}
		if _TransparentUpgradeableProxy == nil {
			return
		}

		_TransparentUpgradeableProxy.Fallback(calldata)
	})
}

func Fuzz_Nosy_TransparentUpgradeableProxyTransactor_Fallback__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactor
		fill_err = tp.Fill(&_TransparentUpgradeableProxy)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var calldata []byte
		fill_err = tp.Fill(&calldata)
		if fill_err != nil {
			return
		}
		if _TransparentUpgradeableProxy == nil || opts == nil {
			return
		}

		_TransparentUpgradeableProxy.Fallback(opts, calldata)
	})
}

// skipping Fuzz_Nosy_TransparentUpgradeableProxyTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_TransparentUpgradeableProxyTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorRaw
		fill_err = tp.Fill(&_TransparentUpgradeableProxy)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _TransparentUpgradeableProxy == nil || opts == nil {
			return
		}

		_TransparentUpgradeableProxy.Transfer(opts)
	})
}

func Fuzz_Nosy_TransparentUpgradeableProxyTransactorSession_Fallback__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorSession
		fill_err = tp.Fill(&_TransparentUpgradeableProxy)
		if fill_err != nil {
			return
		}
		var calldata []byte
		fill_err = tp.Fill(&calldata)
		if fill_err != nil {
			return
		}
		if _TransparentUpgradeableProxy == nil {
			return
		}

		_TransparentUpgradeableProxy.Fallback(calldata)
	})
}

func Fuzz_Nosy_TransparentUpgradeableProxyUpgradedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *TransparentUpgradeableProxyUpgradedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_TransparentUpgradeableProxyUpgradedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *TransparentUpgradeableProxyUpgradedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_TransparentUpgradeableProxyUpgradedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *TransparentUpgradeableProxyUpgradedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_UpgradeCaller_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeCaller
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil || opts == nil {
			return
		}

		_Upgrade.Owner(opts)
	})
}

// skipping Fuzz_Nosy_UpgradeCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_UpgradeCallerSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeCallerSession
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil {
			return
		}

		_Upgrade.Owner()
	})
}

func Fuzz_Nosy_UpgradeCancelUpgradeIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *UpgradeCancelUpgradeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_UpgradeCancelUpgradeIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *UpgradeCancelUpgradeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_UpgradeCancelUpgradeIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *UpgradeCancelUpgradeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_UpgradeFilterer_FilterCancelUpgrade__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeFilterer
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil || opts == nil {
			return
		}

		_Upgrade.FilterCancelUpgrade(opts)
	})
}

func Fuzz_Nosy_UpgradeFilterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeFilterer
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil || opts == nil {
			return
		}

		_Upgrade.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_UpgradeFilterer_FilterOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeFilterer
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var previousOwner []common.Address
		fill_err = tp.Fill(&previousOwner)
		if fill_err != nil {
			return
		}
		var newOwner []common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil || opts == nil {
			return
		}

		_Upgrade.FilterOwnershipTransferred(opts, previousOwner, newOwner)
	})
}

func Fuzz_Nosy_UpgradeFilterer_FilterPlanUpgrade__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeFilterer
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil || opts == nil {
			return
		}

		_Upgrade.FilterPlanUpgrade(opts)
	})
}

func Fuzz_Nosy_UpgradeFilterer_ParseCancelUpgrade__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeFilterer
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil {
			return
		}

		_Upgrade.ParseCancelUpgrade(log)
	})
}

func Fuzz_Nosy_UpgradeFilterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeFilterer
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil {
			return
		}

		_Upgrade.ParseInitialized(log)
	})
}

func Fuzz_Nosy_UpgradeFilterer_ParseOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeFilterer
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil {
			return
		}

		_Upgrade.ParseOwnershipTransferred(log)
	})
}

func Fuzz_Nosy_UpgradeFilterer_ParsePlanUpgrade__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeFilterer
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil {
			return
		}

		_Upgrade.ParsePlanUpgrade(log)
	})
}

// skipping Fuzz_Nosy_UpgradeFilterer_WatchCancelUpgrade__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.UpgradeCancelUpgrade

// skipping Fuzz_Nosy_UpgradeFilterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.UpgradeInitialized

// skipping Fuzz_Nosy_UpgradeFilterer_WatchOwnershipTransferred__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.UpgradeOwnershipTransferred

// skipping Fuzz_Nosy_UpgradeFilterer_WatchPlanUpgrade__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.UpgradePlanUpgrade

func Fuzz_Nosy_UpgradeInitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *UpgradeInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_UpgradeInitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *UpgradeInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_UpgradeInitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *UpgradeInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_UpgradeOwnershipTransferredIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *UpgradeOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_UpgradeOwnershipTransferredIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *UpgradeOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_UpgradeOwnershipTransferredIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *UpgradeOwnershipTransferredIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_UpgradePlanUpgradeIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *UpgradePlanUpgradeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_UpgradePlanUpgradeIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *UpgradePlanUpgradeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_UpgradePlanUpgradeIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *UpgradePlanUpgradeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

// skipping Fuzz_Nosy_UpgradeRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_UpgradeRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_UpgradeRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeRaw
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil || opts == nil {
			return
		}

		_Upgrade.Transfer(opts)
	})
}

func Fuzz_Nosy_UpgradeSession_CancelUpgrade__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeSession
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil {
			return
		}

		_Upgrade.CancelUpgrade()
	})
}

func Fuzz_Nosy_UpgradeSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeSession
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		var owner_ common.Address
		fill_err = tp.Fill(&owner_)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil {
			return
		}

		_Upgrade.Initialize(owner_)
	})
}

func Fuzz_Nosy_UpgradeSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeSession
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil {
			return
		}

		_Upgrade.Owner()
	})
}

func Fuzz_Nosy_UpgradeSession_PlanUpgrade__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeSession
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		var plan UpgradePlan
		fill_err = tp.Fill(&plan)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil {
			return
		}

		_Upgrade.PlanUpgrade(plan)
	})
}

func Fuzz_Nosy_UpgradeSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeSession
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil {
			return
		}

		_Upgrade.RenounceOwnership()
	})
}

func Fuzz_Nosy_UpgradeSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeSession
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil {
			return
		}

		_Upgrade.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_UpgradeTransactor_CancelUpgrade__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeTransactor
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil || opts == nil {
			return
		}

		_Upgrade.CancelUpgrade(opts)
	})
}

func Fuzz_Nosy_UpgradeTransactor_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeTransactor
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var owner_ common.Address
		fill_err = tp.Fill(&owner_)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil || opts == nil {
			return
		}

		_Upgrade.Initialize(opts, owner_)
	})
}

func Fuzz_Nosy_UpgradeTransactor_PlanUpgrade__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeTransactor
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var plan UpgradePlan
		fill_err = tp.Fill(&plan)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil || opts == nil {
			return
		}

		_Upgrade.PlanUpgrade(opts, plan)
	})
}

func Fuzz_Nosy_UpgradeTransactor_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeTransactor
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil || opts == nil {
			return
		}

		_Upgrade.RenounceOwnership(opts)
	})
}

func Fuzz_Nosy_UpgradeTransactor_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeTransactor
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil || opts == nil {
			return
		}

		_Upgrade.TransferOwnership(opts, newOwner)
	})
}

// skipping Fuzz_Nosy_UpgradeTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_UpgradeTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeTransactorRaw
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil || opts == nil {
			return
		}

		_Upgrade.Transfer(opts)
	})
}

func Fuzz_Nosy_UpgradeTransactorSession_CancelUpgrade__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeTransactorSession
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil {
			return
		}

		_Upgrade.CancelUpgrade()
	})
}

func Fuzz_Nosy_UpgradeTransactorSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeTransactorSession
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		var owner_ common.Address
		fill_err = tp.Fill(&owner_)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil {
			return
		}

		_Upgrade.Initialize(owner_)
	})
}

func Fuzz_Nosy_UpgradeTransactorSession_PlanUpgrade__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeTransactorSession
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		var plan UpgradePlan
		fill_err = tp.Fill(&plan)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil {
			return
		}

		_Upgrade.PlanUpgrade(plan)
	})
}

func Fuzz_Nosy_UpgradeTransactorSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeTransactorSession
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil {
			return
		}

		_Upgrade.RenounceOwnership()
	})
}

func Fuzz_Nosy_UpgradeTransactorSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Upgrade *UpgradeTransactorSession
		fill_err = tp.Fill(&_Upgrade)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _Upgrade == nil {
			return
		}

		_Upgrade.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_WOmniApprovalIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *WOmniApprovalIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_WOmniApprovalIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *WOmniApprovalIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_WOmniApprovalIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *WOmniApprovalIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_WOmniCaller_Allowance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniCaller
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 common.Address
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _WOmni == nil || opts == nil {
			return
		}

		_WOmni.Allowance(opts, arg0, arg1)
	})
}

func Fuzz_Nosy_WOmniCaller_BalanceOf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniCaller
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _WOmni == nil || opts == nil {
			return
		}

		_WOmni.BalanceOf(opts, arg0)
	})
}

func Fuzz_Nosy_WOmniCaller_Decimals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniCaller
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _WOmni == nil || opts == nil {
			return
		}

		_WOmni.Decimals(opts)
	})
}

func Fuzz_Nosy_WOmniCaller_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniCaller
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _WOmni == nil || opts == nil {
			return
		}

		_WOmni.Name(opts)
	})
}

func Fuzz_Nosy_WOmniCaller_Symbol__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniCaller
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _WOmni == nil || opts == nil {
			return
		}

		_WOmni.Symbol(opts)
	})
}

func Fuzz_Nosy_WOmniCaller_TotalSupply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniCaller
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _WOmni == nil || opts == nil {
			return
		}

		_WOmni.TotalSupply(opts)
	})
}

// skipping Fuzz_Nosy_WOmniCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_WOmniCallerSession_Allowance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniCallerSession
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 common.Address
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _WOmni == nil {
			return
		}

		_WOmni.Allowance(arg0, arg1)
	})
}

func Fuzz_Nosy_WOmniCallerSession_BalanceOf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniCallerSession
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _WOmni == nil {
			return
		}

		_WOmni.BalanceOf(arg0)
	})
}

func Fuzz_Nosy_WOmniCallerSession_Decimals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniCallerSession
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		if _WOmni == nil {
			return
		}

		_WOmni.Decimals()
	})
}

func Fuzz_Nosy_WOmniCallerSession_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniCallerSession
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		if _WOmni == nil {
			return
		}

		_WOmni.Name()
	})
}

func Fuzz_Nosy_WOmniCallerSession_Symbol__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniCallerSession
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		if _WOmni == nil {
			return
		}

		_WOmni.Symbol()
	})
}

func Fuzz_Nosy_WOmniCallerSession_TotalSupply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniCallerSession
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		if _WOmni == nil {
			return
		}

		_WOmni.TotalSupply()
	})
}

func Fuzz_Nosy_WOmniDepositIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *WOmniDepositIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_WOmniDepositIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *WOmniDepositIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_WOmniDepositIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *WOmniDepositIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_WOmniFilterer_FilterApproval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniFilterer
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var src []common.Address
		fill_err = tp.Fill(&src)
		if fill_err != nil {
			return
		}
		var guy []common.Address
		fill_err = tp.Fill(&guy)
		if fill_err != nil {
			return
		}
		if _WOmni == nil || opts == nil {
			return
		}

		_WOmni.FilterApproval(opts, src, guy)
	})
}

func Fuzz_Nosy_WOmniFilterer_FilterDeposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniFilterer
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var dst []common.Address
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		if _WOmni == nil || opts == nil {
			return
		}

		_WOmni.FilterDeposit(opts, dst)
	})
}

func Fuzz_Nosy_WOmniFilterer_FilterTransfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniFilterer
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var src []common.Address
		fill_err = tp.Fill(&src)
		if fill_err != nil {
			return
		}
		var dst []common.Address
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		if _WOmni == nil || opts == nil {
			return
		}

		_WOmni.FilterTransfer(opts, src, dst)
	})
}

func Fuzz_Nosy_WOmniFilterer_FilterWithdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniFilterer
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var src []common.Address
		fill_err = tp.Fill(&src)
		if fill_err != nil {
			return
		}
		if _WOmni == nil || opts == nil {
			return
		}

		_WOmni.FilterWithdrawal(opts, src)
	})
}

func Fuzz_Nosy_WOmniFilterer_ParseApproval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniFilterer
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _WOmni == nil {
			return
		}

		_WOmni.ParseApproval(log)
	})
}

func Fuzz_Nosy_WOmniFilterer_ParseDeposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniFilterer
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _WOmni == nil {
			return
		}

		_WOmni.ParseDeposit(log)
	})
}

func Fuzz_Nosy_WOmniFilterer_ParseTransfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniFilterer
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _WOmni == nil {
			return
		}

		_WOmni.ParseTransfer(log)
	})
}

func Fuzz_Nosy_WOmniFilterer_ParseWithdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniFilterer
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _WOmni == nil {
			return
		}

		_WOmni.ParseWithdrawal(log)
	})
}

// skipping Fuzz_Nosy_WOmniFilterer_WatchApproval__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.WOmniApproval

// skipping Fuzz_Nosy_WOmniFilterer_WatchDeposit__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.WOmniDeposit

// skipping Fuzz_Nosy_WOmniFilterer_WatchTransfer__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.WOmniTransfer

// skipping Fuzz_Nosy_WOmniFilterer_WatchWithdrawal__ because parameters include func, chan, or unsupported interface: chan<- *github.com/omni-network/omni/contracts/bindings.WOmniWithdrawal

// skipping Fuzz_Nosy_WOmniRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_WOmniRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_WOmniRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniRaw
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _WOmni == nil || opts == nil {
			return
		}

		_WOmni.Transfer(opts)
	})
}

func Fuzz_Nosy_WOmniSession_Allowance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniSession
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 common.Address
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _WOmni == nil {
			return
		}

		_WOmni.Allowance(arg0, arg1)
	})
}

func Fuzz_Nosy_WOmniSession_Approve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniSession
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var guy common.Address
		fill_err = tp.Fill(&guy)
		if fill_err != nil {
			return
		}
		var wad *big.Int
		fill_err = tp.Fill(&wad)
		if fill_err != nil {
			return
		}
		if _WOmni == nil || wad == nil {
			return
		}

		_WOmni.Approve(guy, wad)
	})
}

func Fuzz_Nosy_WOmniSession_BalanceOf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniSession
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _WOmni == nil {
			return
		}

		_WOmni.BalanceOf(arg0)
	})
}

func Fuzz_Nosy_WOmniSession_Decimals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniSession
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		if _WOmni == nil {
			return
		}

		_WOmni.Decimals()
	})
}

func Fuzz_Nosy_WOmniSession_Deposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniSession
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		if _WOmni == nil {
			return
		}

		_WOmni.Deposit()
	})
}

func Fuzz_Nosy_WOmniSession_Fallback__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniSession
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var calldata []byte
		fill_err = tp.Fill(&calldata)
		if fill_err != nil {
			return
		}
		if _WOmni == nil {
			return
		}

		_WOmni.Fallback(calldata)
	})
}

func Fuzz_Nosy_WOmniSession_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniSession
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		if _WOmni == nil {
			return
		}

		_WOmni.Name()
	})
}

func Fuzz_Nosy_WOmniSession_Receive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniSession
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		if _WOmni == nil {
			return
		}

		_WOmni.Receive()
	})
}

func Fuzz_Nosy_WOmniSession_Symbol__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniSession
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		if _WOmni == nil {
			return
		}

		_WOmni.Symbol()
	})
}

func Fuzz_Nosy_WOmniSession_TotalSupply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniSession
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		if _WOmni == nil {
			return
		}

		_WOmni.TotalSupply()
	})
}

func Fuzz_Nosy_WOmniSession_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniSession
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var dst common.Address
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var wad *big.Int
		fill_err = tp.Fill(&wad)
		if fill_err != nil {
			return
		}
		if _WOmni == nil || wad == nil {
			return
		}

		_WOmni.Transfer(dst, wad)
	})
}

func Fuzz_Nosy_WOmniSession_TransferFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniSession
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var src common.Address
		fill_err = tp.Fill(&src)
		if fill_err != nil {
			return
		}
		var dst common.Address
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var wad *big.Int
		fill_err = tp.Fill(&wad)
		if fill_err != nil {
			return
		}
		if _WOmni == nil || wad == nil {
			return
		}

		_WOmni.TransferFrom(src, dst, wad)
	})
}

func Fuzz_Nosy_WOmniSession_Withdraw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniSession
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var wad *big.Int
		fill_err = tp.Fill(&wad)
		if fill_err != nil {
			return
		}
		if _WOmni == nil || wad == nil {
			return
		}

		_WOmni.Withdraw(wad)
	})
}

func Fuzz_Nosy_WOmniTransactor_Approve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniTransactor
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var guy common.Address
		fill_err = tp.Fill(&guy)
		if fill_err != nil {
			return
		}
		var wad *big.Int
		fill_err = tp.Fill(&wad)
		if fill_err != nil {
			return
		}
		if _WOmni == nil || opts == nil || wad == nil {
			return
		}

		_WOmni.Approve(opts, guy, wad)
	})
}

func Fuzz_Nosy_WOmniTransactor_Deposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniTransactor
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _WOmni == nil || opts == nil {
			return
		}

		_WOmni.Deposit(opts)
	})
}

func Fuzz_Nosy_WOmniTransactor_Fallback__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniTransactor
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var calldata []byte
		fill_err = tp.Fill(&calldata)
		if fill_err != nil {
			return
		}
		if _WOmni == nil || opts == nil {
			return
		}

		_WOmni.Fallback(opts, calldata)
	})
}

func Fuzz_Nosy_WOmniTransactor_Receive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniTransactor
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _WOmni == nil || opts == nil {
			return
		}

		_WOmni.Receive(opts)
	})
}

func Fuzz_Nosy_WOmniTransactor_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniTransactor
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var dst common.Address
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var wad *big.Int
		fill_err = tp.Fill(&wad)
		if fill_err != nil {
			return
		}
		if _WOmni == nil || opts == nil || wad == nil {
			return
		}

		_WOmni.Transfer(opts, dst, wad)
	})
}

func Fuzz_Nosy_WOmniTransactor_TransferFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniTransactor
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var src common.Address
		fill_err = tp.Fill(&src)
		if fill_err != nil {
			return
		}
		var dst common.Address
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var wad *big.Int
		fill_err = tp.Fill(&wad)
		if fill_err != nil {
			return
		}
		if _WOmni == nil || opts == nil || wad == nil {
			return
		}

		_WOmni.TransferFrom(opts, src, dst, wad)
	})
}

func Fuzz_Nosy_WOmniTransactor_Withdraw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniTransactor
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var wad *big.Int
		fill_err = tp.Fill(&wad)
		if fill_err != nil {
			return
		}
		if _WOmni == nil || opts == nil || wad == nil {
			return
		}

		_WOmni.Withdraw(opts, wad)
	})
}

// skipping Fuzz_Nosy_WOmniTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_WOmniTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniTransactorRaw
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _WOmni == nil || opts == nil {
			return
		}

		_WOmni.Transfer(opts)
	})
}

func Fuzz_Nosy_WOmniTransactorSession_Approve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniTransactorSession
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var guy common.Address
		fill_err = tp.Fill(&guy)
		if fill_err != nil {
			return
		}
		var wad *big.Int
		fill_err = tp.Fill(&wad)
		if fill_err != nil {
			return
		}
		if _WOmni == nil || wad == nil {
			return
		}

		_WOmni.Approve(guy, wad)
	})
}

func Fuzz_Nosy_WOmniTransactorSession_Deposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniTransactorSession
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		if _WOmni == nil {
			return
		}

		_WOmni.Deposit()
	})
}

func Fuzz_Nosy_WOmniTransactorSession_Fallback__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniTransactorSession
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var calldata []byte
		fill_err = tp.Fill(&calldata)
		if fill_err != nil {
			return
		}
		if _WOmni == nil {
			return
		}

		_WOmni.Fallback(calldata)
	})
}

func Fuzz_Nosy_WOmniTransactorSession_Receive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniTransactorSession
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		if _WOmni == nil {
			return
		}

		_WOmni.Receive()
	})
}

func Fuzz_Nosy_WOmniTransactorSession_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniTransactorSession
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var dst common.Address
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var wad *big.Int
		fill_err = tp.Fill(&wad)
		if fill_err != nil {
			return
		}
		if _WOmni == nil || wad == nil {
			return
		}

		_WOmni.Transfer(dst, wad)
	})
}

func Fuzz_Nosy_WOmniTransactorSession_TransferFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniTransactorSession
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var src common.Address
		fill_err = tp.Fill(&src)
		if fill_err != nil {
			return
		}
		var dst common.Address
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var wad *big.Int
		fill_err = tp.Fill(&wad)
		if fill_err != nil {
			return
		}
		if _WOmni == nil || wad == nil {
			return
		}

		_WOmni.TransferFrom(src, dst, wad)
	})
}

func Fuzz_Nosy_WOmniTransactorSession_Withdraw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _WOmni *WOmniTransactorSession
		fill_err = tp.Fill(&_WOmni)
		if fill_err != nil {
			return
		}
		var wad *big.Int
		fill_err = tp.Fill(&wad)
		if fill_err != nil {
			return
		}
		if _WOmni == nil || wad == nil {
			return
		}

		_WOmni.Withdraw(wad)
	})
}

func Fuzz_Nosy_WOmniTransferIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *WOmniTransferIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_WOmniTransferIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *WOmniTransferIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_WOmniTransferIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *WOmniTransferIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_WOmniWithdrawalIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *WOmniWithdrawalIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_WOmniWithdrawalIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *WOmniWithdrawalIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_WOmniWithdrawalIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *WOmniWithdrawalIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

// skipping Fuzz_Nosy_DeployAVSDirectory__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

// skipping Fuzz_Nosy_DeployAdmin__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

// skipping Fuzz_Nosy_DeployAllocPredeploys__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

// skipping Fuzz_Nosy_DeployCreate3__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

// skipping Fuzz_Nosy_DeployDelegationManager__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

// skipping Fuzz_Nosy_DeployFeeOracleV1__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

// skipping Fuzz_Nosy_DeployMockERC20__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

// skipping Fuzz_Nosy_DeployOmni__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

// skipping Fuzz_Nosy_DeployOmniAVS__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

// skipping Fuzz_Nosy_DeployOmniBridgeL1__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

// skipping Fuzz_Nosy_DeployOmniBridgeNative__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

// skipping Fuzz_Nosy_DeployOmniPortal__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

// skipping Fuzz_Nosy_DeployPingPong__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

// skipping Fuzz_Nosy_DeployPortalRegistry__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

// skipping Fuzz_Nosy_DeployProxyAdmin__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

// skipping Fuzz_Nosy_DeploySlashing__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

// skipping Fuzz_Nosy_DeployStaking__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

// skipping Fuzz_Nosy_DeployStrategyBase__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

// skipping Fuzz_Nosy_DeployStrategyManager__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

// skipping Fuzz_Nosy_DeployTransparentUpgradeableProxy__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

// skipping Fuzz_Nosy_DeployUpgrade__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

// skipping Fuzz_Nosy_DeployWOmni__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend
