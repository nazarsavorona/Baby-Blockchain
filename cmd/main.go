package main

import (
	"baby-blockchain/pkg/Account"
	"baby-blockchain/pkg/KeyPair"
	"baby-blockchain/pkg/Signature"
	"baby-blockchain/pkg/Transaction"
	"baby-blockchain/pkg/Vote"
	"baby-blockchain/pkg/Voting"
	"log"
	"math/rand"
)

func main() {
	keyPair := KeyPair.GenKeyPair(nil)
	keyPair.Print()

	message := "Love GoLang"
	signature := Signature.SignData(message, keyPair.PrivateKey)
	log.Printf("Signature verified: %v", signature.VerifySignature(message, keyPair.PublicKey))
	log.Printf("Signature verified: %v", signature.VerifySignature(message, KeyPair.GenKeyPair(nil).PublicKey))
	signature.Print()

	account := Account.GenAccount(nil)

	voting := Voting.CreateVoting("Do you love Distributed Lab?",
		[]string{
			"Yes",
			"No",
			"What is Distributed Lab?",
		})

	voting.Print()

	vote, err := Vote.CreateVote(account, voting, 2)

	if err != nil {
		log.Fatal(err)
	}

	vote.Print()
	log.Printf("Vote verified: %v", Vote.VerifyVote(vote))

	transaction := Transaction.CreateTransaction(vote, rand.Int())
	transaction.Print()
}
