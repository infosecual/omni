package types

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/omni-network/omni/e2e"
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

func Fuzz_Nosy_DeployInfos_Addr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var file string
		fill_err = tp.Fill(&file)
		if fill_err != nil {
			return
		}
		var chainID uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var contract ContractName
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}

		i, err := LoadDeployInfos(file)
		if err != nil {
			return
		}
		i.Addr(chainID, contract)
	})
}

func Fuzz_Nosy_DeployInfos_Save__(f *testing.F) {
	f.Fuzz(func(t *testing.T, f1 string, f2 string) {
		i, err := LoadDeployInfos(f1)
		if err != nil {
			return
		}
		i.Save(f2)
	})
}

func Fuzz_Nosy_DeployInfos_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var file string
		fill_err = tp.Fill(&file)
		if fill_err != nil {
			return
		}
		var chainID uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var contract ContractName
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}

		i, err := LoadDeployInfos(file)
		if err != nil {
			return
		}
		i.Set(chainID, contract, addr, height)
	})
}

func Fuzz_Nosy_EVMChain_AttestInterval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n1 netconf.ID
		fill_err = tp.Fill(&n1)
		if fill_err != nil {
			return
		}
		var n2 netconf.ID
		fill_err = tp.Fill(&n2)
		if fill_err != nil {
			return
		}

		c := OmniEVMByNetwork(n1)
		c.AttestInterval(n2)
	})
}

func Fuzz_Nosy_EVMChain_ShardsUint64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var network netconf.ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}

		c := OmniEVMByNetwork(network)
		c.ShardsUint64()
	})
}

// skipping Fuzz_Nosy_InfraProvider_Clean__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/e2e/types.InfraProvider

// skipping Fuzz_Nosy_InfraProvider_Restart__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/e2e/types.InfraProvider

// skipping Fuzz_Nosy_InfraProvider_Upgrade__ because parameters include func, chan, or unsupported interface: github.com/omni-network/omni/e2e/types.InfraProvider

func Fuzz_Nosy_InfrastructureData_ServicesByInstance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d InfrastructureData
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var d2 e2e.InstanceData
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}

		d.ServicesByInstance(d2)
	})
}

func Fuzz_Nosy_Manifest_OmniEVMs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m Manifest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}

		m.OmniEVMs()
	})
}

func Fuzz_Nosy_Manifest_Seeds__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m Manifest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}

		m.Seeds()
	})
}

func Fuzz_Nosy_OmniEVM_NodeKeyHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o OmniEVM
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}

		o.NodeKeyHex()
	})
}

func Fuzz_Nosy_PublicChain_Chain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain EVMChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var rpcAddresses []string
		fill_err = tp.Fill(&rpcAddresses)
		if fill_err != nil {
			return
		}

		c := NewPublicChain(chain, rpcAddresses)
		c.Chain()
	})
}

func Fuzz_Nosy_PublicChain_NextRPCAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain EVMChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var rpcAddresses []string
		fill_err = tp.Fill(&rpcAddresses)
		if fill_err != nil {
			return
		}

		c := NewPublicChain(chain, rpcAddresses)
		c.NextRPCAddress()
	})
}

func Fuzz_Nosy_Testnet_ArchiveNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 Testnet
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		t1.ArchiveNode()
	})
}

func Fuzz_Nosy_Testnet_BroadcastNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 Testnet
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		t1.BroadcastNode()
	})
}

func Fuzz_Nosy_Testnet_BroadcastOmniEVM__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 Testnet
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		t1.BroadcastOmniEVM()
	})
}

func Fuzz_Nosy_Testnet_HasOmniEVM__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 Testnet
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		t1.HasOmniEVM()
	})
}

func Fuzz_Nosy_Testnet_HasPerturbations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 Testnet
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		t1.HasPerturbations()
	})
}

func Fuzz_Nosy_Testnet_RandomHaloAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 Testnet
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		t1.RandomHaloAddr()
	})
}

func Fuzz_Nosy_AnvilChainsByNames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var names []string
		fill_err = tp.Fill(&names)
		if fill_err != nil {
			return
		}

		AnvilChainsByNames(names)
	})
}

func Fuzz_Nosy_PublicRPCByName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, name string) {
		PublicRPCByName(name)
	})
}

func Fuzz_Nosy_instancesEqual__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a e2e.InstanceData
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b e2e.InstanceData
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		instancesEqual(a, b)
	})
}

func Fuzz_Nosy_intervalFromPeriod__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var network netconf.ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var period time.Duration
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}

		intervalFromPeriod(network, period)
	})
}
