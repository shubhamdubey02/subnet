// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"fmt"

	"github.com/cryft-labs/cryftgo/version"
	"github.com/shubhamdubey02/subnet/plugin/evm"
	"github.com/shubhamdubey02/subnet/plugin/runner"
)

func main() {
	versionString := fmt.Sprintf("Subnet-EVM/%s [AvalancheGo=%s, rpcchainvm=%d]", evm.Version, version.Current, version.RPCChainVMProtocol)
	runner.Run(versionString)
}