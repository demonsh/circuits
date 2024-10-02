package main

import (
	"encoding/json"
	"math/big"

	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-merkletree-sql/v2"
)

var PAYMENTID = big.NewInt(1)

func main() {
	println("Started")

	ts := NewOnChainTreeStore()

	deposit := NewDeposit()
	pubKeyHash := deposit.PublicKeyHash()
	println("pubKeyHash:", pubKeyHash.String())

	ts.Add(PAYMENTID, pubKeyHash)

	println("Root:", ts.Tree.Root().BigInt().String())

	proof := ts.Proof(big.NewInt(1))
	proofJson, _ := json.Marshal(proof)
	println(string(proofJson))

	amount := big.NewInt(1)
	receiver := big.NewInt(2)
	nonce := big.NewInt(3)
	authMsg := NewAuthorizationMessage(amount, receiver, nonce, PAYMENTID)
	authMsgHash := authMsg.Hash()

	signature := deposit.Sign(authMsgHash)

	s := struct {
		Amount     string             `json:"amount"`
		Receiver   string             `json:"receiver"`
		Nonce      string             `json:"nonce"`
		S          string             `json:"signatureS"`
		X          string             `json:"signatureR8X"`
		Y          string             `json:"signatureR8Y"`
		PaymentID  string             `json:"paymentID"`
		PrivateKey string             `json:"paymentPrivateKey"`
		Root       string             `json:"root"`
		Tree       []*merkletree.Hash `json:"tree"`
	}{
		X:          signature.R8.X.String(),
		Y:          signature.R8.Y.String(),
		S:          signature.S.String(),
		Amount:     amount.String(),
		Receiver:   receiver.String(),
		Nonce:      nonce.String(),
		Root:       ts.Tree.Root().BigInt().String(),
		PaymentID:  PAYMENTID.String(),
		PrivateKey: babyjub.SkToBigInt(&deposit.BJJKey).String(),
		Tree:       ts.Siblings(PAYMENTID),
	}
	sJson, err := json.Marshal(s)
	if err != nil {
		println("Error marshalling struct:", err)
	} else {
		println(string(sJson))
	}

}
