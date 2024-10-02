package main

import (
	"math/big"

	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-iden3-crypto/poseidon"
	"github.com/iden3/go-iden3-crypto/utils"
)

type Deposit struct {
	BJJKey babyjub.PrivateKey
}

const userPK = "28156abe7fe2fd433dc9df969286b96666489bac508612d0e16593e944c4f69f"

func NewDeposit() *Deposit {
	var k babyjub.PrivateKey

	intKey, _ := big.NewInt(0).SetString(userPK, 16)

	field := utils.CheckBigIntInField(intKey)
	if field == false {
		panic("key not in field")
	}

	bytes := utils.BigIntLEBytes(intKey)
	// copy bytes to private key
	copy(k[:], bytes[:])

	return &Deposit{
		BJJKey: k,
	}
}

func (d *Deposit) Sign(msg *big.Int) *babyjub.Signature {
	return d.BJJKey.SignPoseidon(msg)
}

func (d *Deposit) PublicKeyHash() *big.Int {
	pubKey := d.BJJKey.Public()
	println("pubKey:", pubKey.X.String(), pubKey.Y.String())

	bytes := utils.BigIntLEBytes(pubKey.X)

	println("ax le:", big.NewInt(0).SetBytes(bytes[:]).String())

	pubKeyHash, err := poseidon.Hash([]*big.Int{pubKey.X, pubKey.Y})
	if err != nil {
		panic(err)
	}
	return pubKeyHash
}

type AuthorizationMessage struct {
	Amount    *big.Int
	Receiver  *big.Int
	Nonce     *big.Int
	PaymentID *big.Int
}

func NewAuthorizationMessage(amount, receiver, nonce, paymentID *big.Int) *AuthorizationMessage {
	return &AuthorizationMessage{
		Amount:    amount,
		Receiver:  receiver,
		Nonce:     nonce,
		PaymentID: paymentID,
	}
}

func (am *AuthorizationMessage) Hash() *big.Int {
	hash, err := poseidon.Hash([]*big.Int{am.Amount, am.Receiver, am.Nonce, am.PaymentID})
	if err != nil {
		panic(err)
	}
	return hash
}
