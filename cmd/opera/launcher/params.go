package launcher

import (
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/params"

	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/go-opera/opera/genesis"
	"github.com/Fantom-foundation/go-opera/opera/genesisstore"
)

// TODO: put enodes here
var (
	Bootnodes = map[string][]string{
		"main": {
			"enode://be2e97ef580b4835932ad7956ff039c62ba7e92735bb62ca6c7627064ca4a03688820c643b87df8d8ff442ccb7f41313d645ed01444bebf30efd7174b173eb23@172.105.183.170:5566",
			"enode://a2248c2359b858b5698f0cb5880badaec8ef1101729a4c18f57cd485a3a541b0bc848267be3d3cbb511d66919593c35c3948cc9760b0b89c79295a6584dee0fd@139.162.161.97:5050",
			"enode://349f106b2e0083e00013c1470721613b9dcb15d94db3b01bcc1c7d69c99a332f67ac76733aa032a54ac84d551010548b741c90831afe0245b2ef58db402d25f4@85.159.211.72:5050",
			"enode://c19eb328c146a32f15125343302f70e65a2b127b1042a01ab87d5996eef034f849f720dc985d4d8710b997a22b7111323be2f2f7622366606ba621a015eb7ae9@172.105.154.248:5050",
		},
		"test": {},
	}

	// asDefault is slice with one empty element which indicates that network default bootnodes should be substituted
	asDefault = []*enode.Node{{}}

	mainnetHeader = genesis.Header{
		GenesisID:   hash.HexToHash("0x1435856b8ee4e65296202b469aa9b18266d65ebce979020d7949b20c9fc7258b"),
		NetworkID:   opera.MainNetworkID,
		NetworkName: "main",
	}

	testnetHeader = genesis.Header{
		GenesisID:   hash.HexToHash("0xc4a5fc96e575a16a9a0c7349d44dc4d0f602a54e0a8543360c2fee4c3937b49e"),
		NetworkID:   opera.TestNetworkID,
		NetworkName: "test",
	}

	AllowedOperaGenesis = []GenesisTemplate{
		{
			Name:   "Mainnet-5577 with pruned MPT",
			Header: mainnetHeader,
			Hashes: genesis.Hashes{
				genesisstore.EpochsSection: hash.HexToHash("0xfca05fe92da9cdd1d3f10a73e1b7f552e4b401601f8f2ef25faa43b249b13d4b"),
				genesisstore.BlocksSection: hash.HexToHash("0xedf38b734737152b355a4043cae3a62372dc5e3208b7ee5c1763ce232253f254"),
				genesisstore.EvmSection:    hash.HexToHash("0x512343461b1b251314500e37515a160e56d705e0ce48533f5cf311a74120351d"),
			},
		},
		{
			Name:   "Mainnet-5577 with full MPT",
			Header: mainnetHeader,
			Hashes: genesis.Hashes{
				genesisstore.EpochsSection: hash.HexToHash("0xfca05fe92da9cdd1d3f10a73e1b7f552e4b401601f8f2ef25faa43b249b13d4b"),
				genesisstore.BlocksSection: hash.HexToHash("0xedf38b734737152b355a4043cae3a62372dc5e3208b7ee5c1763ce232253f254"),
				genesisstore.EvmSection:    hash.HexToHash("0x512343461b1b251314500e37515a160e56d705e0ce48533f5cf311a74120351d"),
			},
		},
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
