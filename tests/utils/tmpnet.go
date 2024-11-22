// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"encoding/json"
	"os"

	"github.com/cryft-labs/cryftgo/config"
	"github.com/cryft-labs/cryftgo/ids"
	"github.com/cryft-labs/cryftgo/tests/fixture/tmpnet"

	"github.com/shubhamdubey02/subnet-evm/plugin/evm"
)

var DefaultChainConfig = tmpnet.FlagsMap{
	"log-level":         "debug",
	"warp-api-enabled":  true,
	"local-txs-enabled": true,
}

func NewTmpnetNodes(count int) []*tmpnet.Node {
	nodes := make([]*tmpnet.Node, count)
	for i := range nodes {
		node := tmpnet.NewNode("")
		node.EnsureKeys()
		nodes[i] = node
	}
	return nodes
}

func NewTmpnetNetwork(owner string, nodes []*tmpnet.Node, flags tmpnet.FlagsMap, subnets ...*tmpnet.Subnet) *tmpnet.Network {
	defaultFlags := tmpnet.FlagsMap{}
	defaultFlags.SetDefaults(flags)
	defaultFlags.SetDefaults(tmpnet.FlagsMap{
		config.ProposerVMUseCurrentHeightKey: true,
	})
	return &tmpnet.Network{
		Owner:        owner,
		DefaultFlags: defaultFlags,
		Nodes:        nodes,
		Subnets:      subnets,
	}
}

// Create the configuration that will enable creation and access to a
// subnet created on a temporary network.
func NewTmpnetSubnet(name string, genesisPath string, chainConfig tmpnet.FlagsMap, nodes ...*tmpnet.Node) *tmpnet.Subnet {
	if len(nodes) == 0 {
		panic("a subnet must be validated by at least one node")
	}

	validatorIDs := make([]ids.NodeID, len(nodes))
	for i, node := range nodes {
		validatorIDs[i] = node.NodeID
	}

	genesisBytes, err := os.ReadFile(genesisPath)
	if err != nil {
		panic(err)
	}

	chainConfigBytes, err := json.Marshal(chainConfig)
	if err != nil {
		panic(err)
	}

	return &tmpnet.Subnet{
		Name: name,
		Chains: []*tmpnet.Chain{
			{
				VMID:         evm.ID,
				Genesis:      genesisBytes,
				Config:       string(chainConfigBytes),
				PreFundedKey: tmpnet.HardhatKey,
			},
		},
		ValidatorIDs: validatorIDs,
	}
}
