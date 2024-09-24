package main

import (
	"context"
	"math/big"

	"github.com/iden3/go-merkletree-sql/v2"
	"github.com/iden3/go-merkletree-sql/v2/db/memory"
)

const LevelsDepth = 4

type OnChainTreeStore struct {
	Tree *merkletree.MerkleTree
}

func NewOnChainTreeStore() *OnChainTreeStore {
	tree, err := merkletree.NewMerkleTree(context.Background(), memory.NewMemoryStorage(), LevelsDepth)
	if err != nil {
		panic(err)
	}
	return &OnChainTreeStore{Tree: tree}
}

func (ts *OnChainTreeStore) Add(k, v *big.Int) error {
	return ts.Tree.Add(context.Background(), k, v)
}

func (ts *OnChainTreeStore) Proof(k *big.Int) *merkletree.Proof {
	proof, _, err := ts.Tree.GenerateProof(context.Background(), k, ts.Tree.Root())
	if err != nil {
		panic(err)
	}
	return proof
}
