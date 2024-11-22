// (c) 2021-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package handlers

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shubhamdubey02/subnet/core/state/snapshot"
	"github.com/shubhamdubey02/subnet/core/types"
)

var (
	_ BlockProvider    = &TestBlockProvider{}
	_ SnapshotProvider = &TestSnapshotProvider{}
)

type TestBlockProvider struct {
	GetBlockFn func(common.Hash, uint64) *types.Block
}

func (t *TestBlockProvider) GetBlock(hash common.Hash, number uint64) *types.Block {
	return t.GetBlockFn(hash, number)
}

type TestSnapshotProvider struct {
	Snapshot *snapshot.Tree
}

func (t *TestSnapshotProvider) Snapshots() *snapshot.Tree {
	return t.Snapshot
}
