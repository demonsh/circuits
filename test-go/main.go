package main

import (
	"encoding/json"
	"math/big"
)

func main() {
	println("Started")

	ts := NewOnChainTreeStore()

	deposit := NewDeposit()
	pubKeyHash := deposit.PublicKeyHash()

	ts.Add(big.NewInt(1), pubKeyHash)

	println(ts.Tree.Root())

	proof := ts.Proof(big.NewInt(1))
	proofJson, _ := json.Marshal(proof)
	println(string(proofJson))

	amount := big.NewInt(1)
	receiver := big.NewInt(2)
	nonce := big.NewInt(3)
	authMsg := NewAuthorizationMessage(amount, receiver, nonce)
	authMsgHash := authMsg.Hash()

	signature := deposit.Sign(authMsgHash)

	s := struct {
		X        *big.Int `json:"x"`
		Y        *big.Int `json:"y"`
		S        *big.Int `json:"s"`
		Amount   *big.Int `json:"amount"`
		Receiver *big.Int `json:"receiver"`
		Nonce    *big.Int `json:"nonce"`
	}{
		X:        signature.R8.X,
		Y:        signature.R8.Y,
		S:        signature.S,
		Amount:   amount,
		Receiver: receiver,
		Nonce:    nonce,
	}
	sJson, err := json.Marshal(s)
	if err != nil {
		println("Error marshalling struct:", err)
	} else {
		println(string(sJson))
	}

}
