package main

import (
	"math/big"

	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-iden3-crypto/poseidon"
)

type Deposit struct {
	BJJKey babyjub.PrivateKey
}

func NewDeposit() *Deposit {
	return &Deposit{
		BJJKey: babyjub.NewRandPrivKey(),
	}
}

func (d *Deposit) Sign(msg *big.Int) *babyjub.Signature {
	return d.BJJKey.SignPoseidon(msg)
}

func (d *Deposit) PublicKeyHash() *big.Int {
	pubKey := d.BJJKey.Public()
	pubKeyHash, err := poseidon.Hash([]*big.Int{pubKey.X, pubKey.Y})
	if err != nil {
		panic(err)
	}
	return pubKeyHash
}

type AuthorizationMessage struct {
	Amount   *big.Int
	Receiver *big.Int
	Nonce    *big.Int
}

func NewAuthorizationMessage(amount, receiver, nonce *big.Int) *AuthorizationMessage {
	return &AuthorizationMessage{
		Amount:   amount,
		Receiver: receiver,
		Nonce:    nonce,
	}
}

func (am *AuthorizationMessage) Hash() *big.Int {
	hash, err := poseidon.Hash([]*big.Int{am.Amount, am.Receiver, am.Nonce})
	if err != nil {
		panic(err)
	}
	return hash
}
