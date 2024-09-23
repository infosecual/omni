package netconf

import (
	"context"
	"testing"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/registry/types"
	"github.com/omni-network/omni/lib/xchain"
	"github.com/spf13/pflag"
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

func Fuzz_Nosy_Chain_ChainVersions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var shards []xchain.ShardID
		fill_err = tp.Fill(&shards)
		if fill_err != nil {
			return
		}

		c := mustSimnetChain(id, shards...)
		c.ChainVersions()
	})
}

func Fuzz_Nosy_Chain_ConfLevels__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var shards []xchain.ShardID
		fill_err = tp.Fill(&shards)
		if fill_err != nil {
			return
		}

		c := mustSimnetChain(id, shards...)
		c.ConfLevels()
	})
}

func Fuzz_Nosy_Chain_ShardsUint64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var shards []xchain.ShardID
		fill_err = tp.Fill(&shards)
		if fill_err != nil {
			return
		}

		c := mustSimnetChain(id, shards...)
		c.ShardsUint64()
	})
}

func Fuzz_Nosy_Chain_Verify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var shards []xchain.ShardID
		fill_err = tp.Fill(&shards)
		if fill_err != nil {
			return
		}

		c := mustSimnetChain(id, shards...)
		c.Verify()
	})
}

func Fuzz_Nosy_ID_IsEphemeral__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i ID
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		i.IsEphemeral()
	})
}

func Fuzz_Nosy_ID_IsProtected__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i ID
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		i.IsProtected()
	})
}

func Fuzz_Nosy_ID_Static__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i ID
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		i.Static()
	})
}

func Fuzz_Nosy_ID_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i ID
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		i.String()
	})
}

func Fuzz_Nosy_ID_Verify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i ID
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		i.Verify()
	})
}

func Fuzz_Nosy_ID_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i ID
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		i.Version()
	})
}

func Fuzz_Nosy_Network_Chain__(f *testing.F) {
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
		var network ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var portals []bindings.PortalRegistryDeployment
		fill_err = tp.Fill(&portals)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		n, err := networkFromPortals(ctx, network, portals)
		if err != nil {
			return
		}
		n.Chain(id)
	})
}

func Fuzz_Nosy_Network_ChainIDs__(f *testing.F) {
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
		var network ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var portals []bindings.PortalRegistryDeployment
		fill_err = tp.Fill(&portals)
		if fill_err != nil {
			return
		}

		n, err := networkFromPortals(ctx, network, portals)
		if err != nil {
			return
		}
		n.ChainIDs()
	})
}

func Fuzz_Nosy_Network_ChainName__(f *testing.F) {
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
		var network ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var portals []bindings.PortalRegistryDeployment
		fill_err = tp.Fill(&portals)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		n, err := networkFromPortals(ctx, network, portals)
		if err != nil {
			return
		}
		n.ChainName(id)
	})
}

func Fuzz_Nosy_Network_ChainNamesByIDs__(f *testing.F) {
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
		var network ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var portals []bindings.PortalRegistryDeployment
		fill_err = tp.Fill(&portals)
		if fill_err != nil {
			return
		}

		n, err := networkFromPortals(ctx, network, portals)
		if err != nil {
			return
		}
		n.ChainNamesByIDs()
	})
}

func Fuzz_Nosy_Network_ChainVersionName__(f *testing.F) {
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
		var network ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var portals []bindings.PortalRegistryDeployment
		fill_err = tp.Fill(&portals)
		if fill_err != nil {
			return
		}
		var chainVer xchain.ChainVersion
		fill_err = tp.Fill(&chainVer)
		if fill_err != nil {
			return
		}

		n, err := networkFromPortals(ctx, network, portals)
		if err != nil {
			return
		}
		n.ChainVersionName(chainVer)
	})
}

func Fuzz_Nosy_Network_ChainVersionNames__(f *testing.F) {
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
		var network ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var portals []bindings.PortalRegistryDeployment
		fill_err = tp.Fill(&portals)
		if fill_err != nil {
			return
		}

		n, err := networkFromPortals(ctx, network, portals)
		if err != nil {
			return
		}
		n.ChainVersionNames()
	})
}

func Fuzz_Nosy_Network_ChainVersionsTo__(f *testing.F) {
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
		var network ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var portals []bindings.PortalRegistryDeployment
		fill_err = tp.Fill(&portals)
		if fill_err != nil {
			return
		}
		var dstChainID uint64
		fill_err = tp.Fill(&dstChainID)
		if fill_err != nil {
			return
		}

		n, err := networkFromPortals(ctx, network, portals)
		if err != nil {
			return
		}
		n.ChainVersionsTo(dstChainID)
	})
}

func Fuzz_Nosy_Network_EVMChains__(f *testing.F) {
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
		var network ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var portals []bindings.PortalRegistryDeployment
		fill_err = tp.Fill(&portals)
		if fill_err != nil {
			return
		}

		n, err := networkFromPortals(ctx, network, portals)
		if err != nil {
			return
		}
		n.EVMChains()
	})
}

func Fuzz_Nosy_Network_EVMStreams__(f *testing.F) {
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
		var network ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var portals []bindings.PortalRegistryDeployment
		fill_err = tp.Fill(&portals)
		if fill_err != nil {
			return
		}

		n, err := networkFromPortals(ctx, network, portals)
		if err != nil {
			return
		}
		n.EVMStreams()
	})
}

func Fuzz_Nosy_Network_EthereumChain__(f *testing.F) {
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
		var network ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var portals []bindings.PortalRegistryDeployment
		fill_err = tp.Fill(&portals)
		if fill_err != nil {
			return
		}

		n, err := networkFromPortals(ctx, network, portals)
		if err != nil {
			return
		}
		n.EthereumChain()
	})
}

func Fuzz_Nosy_Network_OmniConsensusChain__(f *testing.F) {
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
		var network ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var portals []bindings.PortalRegistryDeployment
		fill_err = tp.Fill(&portals)
		if fill_err != nil {
			return
		}

		n, err := networkFromPortals(ctx, network, portals)
		if err != nil {
			return
		}
		n.OmniConsensusChain()
	})
}

func Fuzz_Nosy_Network_OmniEVMChain__(f *testing.F) {
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
		var network ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var portals []bindings.PortalRegistryDeployment
		fill_err = tp.Fill(&portals)
		if fill_err != nil {
			return
		}

		n, err := networkFromPortals(ctx, network, portals)
		if err != nil {
			return
		}
		n.OmniEVMChain()
	})
}

func Fuzz_Nosy_Network_StreamName__(f *testing.F) {
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
		var network ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var portals []bindings.PortalRegistryDeployment
		fill_err = tp.Fill(&portals)
		if fill_err != nil {
			return
		}
		var stream xchain.StreamID
		fill_err = tp.Fill(&stream)
		if fill_err != nil {
			return
		}

		n, err := networkFromPortals(ctx, network, portals)
		if err != nil {
			return
		}
		n.StreamName(stream)
	})
}

func Fuzz_Nosy_Network_StreamsBetween__(f *testing.F) {
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
		var network ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var portals []bindings.PortalRegistryDeployment
		fill_err = tp.Fill(&portals)
		if fill_err != nil {
			return
		}
		var srcChainID uint64
		fill_err = tp.Fill(&srcChainID)
		if fill_err != nil {
			return
		}
		var dstChainID uint64
		fill_err = tp.Fill(&dstChainID)
		if fill_err != nil {
			return
		}

		n, err := networkFromPortals(ctx, network, portals)
		if err != nil {
			return
		}
		n.StreamsBetween(srcChainID, dstChainID)
	})
}

func Fuzz_Nosy_Network_StreamsFrom__(f *testing.F) {
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
		var network ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var portals []bindings.PortalRegistryDeployment
		fill_err = tp.Fill(&portals)
		if fill_err != nil {
			return
		}
		var srcChainID uint64
		fill_err = tp.Fill(&srcChainID)
		if fill_err != nil {
			return
		}

		n, err := networkFromPortals(ctx, network, portals)
		if err != nil {
			return
		}
		n.StreamsFrom(srcChainID)
	})
}

func Fuzz_Nosy_Network_StreamsTo__(f *testing.F) {
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
		var network ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var portals []bindings.PortalRegistryDeployment
		fill_err = tp.Fill(&portals)
		if fill_err != nil {
			return
		}
		var dstChainID uint64
		fill_err = tp.Fill(&dstChainID)
		if fill_err != nil {
			return
		}

		n, err := networkFromPortals(ctx, network, portals)
		if err != nil {
			return
		}
		n.StreamsTo(dstChainID)
	})
}

func Fuzz_Nosy_Network_Verify__(f *testing.F) {
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
		var network ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var portals []bindings.PortalRegistryDeployment
		fill_err = tp.Fill(&portals)
		if fill_err != nil {
			return
		}

		n, err := networkFromPortals(ctx, network, portals)
		if err != nil {
			return
		}
		n.Verify()
	})
}

func Fuzz_Nosy_Static_ConsensusRPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s Static
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.ConsensusRPC()
	})
}

func Fuzz_Nosy_Static_ConsensusSeeds__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s Static
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.ConsensusSeeds()
	})
}

func Fuzz_Nosy_Static_ExecutionRPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s Static
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.ExecutionRPC()
	})
}

func Fuzz_Nosy_Static_ExecutionSeeds__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s Static
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.ExecutionSeeds()
	})
}

func Fuzz_Nosy_Static_OmniConsensusChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s Static
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.OmniConsensusChain()
	})
}

func Fuzz_Nosy_Static_OmniConsensusChainIDStr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s Static
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.OmniConsensusChainIDStr()
	})
}

func Fuzz_Nosy_Static_OmniConsensusChainIDUint64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s Static
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.OmniConsensusChainIDUint64()
	})
}

func Fuzz_Nosy_Static_OmniExecutionChainName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s Static
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.OmniExecutionChainName()
	})
}

func Fuzz_Nosy_Static_PortalDeployment__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s Static
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var chainID uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}

		s.PortalDeployment(chainID)
	})
}

func Fuzz_Nosy_BindFlag__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var flags *pflag.FlagSet
		fill_err = tp.Fill(&flags)
		if fill_err != nil {
			return
		}
		var network *ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		if flags == nil || network == nil {
			return
		}

		BindFlag(flags, network)
	})
}

func Fuzz_Nosy_ChainNamer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var network ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}

		ChainNamer(network)
	})
}

func Fuzz_Nosy_ChainVersionNamer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var network ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}

		ChainVersionNamer(network)
	})
}

func Fuzz_Nosy_ConsensusChainIDStr2Uint64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id string) {
		ConsensusChainIDStr2Uint64(id)
	})
}

func Fuzz_Nosy_IsAny__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var targets []ID
		fill_err = tp.Fill(&targets)
		if fill_err != nil {
			return
		}

		IsAny(id, targets...)
	})
}

func Fuzz_Nosy_IsEthereumChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var network ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var chainID uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}

		IsEthereumChain(network, chainID)
	})
}

func Fuzz_Nosy_IsOmniConsensus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var network ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var chainID uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}

		IsOmniConsensus(network, chainID)
	})
}

func Fuzz_Nosy_IsOmniExecution__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var network ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var chainID uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}

		IsOmniExecution(network, chainID)
	})
}

func Fuzz_Nosy_MetadataByName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var network ID
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}

		MetadataByName(network, name)
	})
}

func Fuzz_Nosy_containsAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var network Network
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var expected []string
		fill_err = tp.Fill(&expected)
		if fill_err != nil {
			return
		}

		containsAll(network, expected)
	})
}

func Fuzz_Nosy_toPortalBindings__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var portals []*types.Portal
		fill_err = tp.Fill(&portals)
		if fill_err != nil {
			return
		}

		toPortalBindings(portals)
	})
}

func Fuzz_Nosy_toShardIDs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var shards []uint64
		fill_err = tp.Fill(&shards)
		if fill_err != nil {
			return
		}

		toShardIDs(shards)
	})
}
