package main

import (
	"baby-blockchain/pkg/Account"
	"baby-blockchain/pkg/Block"
	"baby-blockchain/pkg/Blockchain"
	"baby-blockchain/pkg/Hash"
	"baby-blockchain/pkg/KeyPair"
	"baby-blockchain/pkg/Signature"
	"baby-blockchain/pkg/Transaction"
	"baby-blockchain/pkg/Vote"
	"baby-blockchain/pkg/Voting"
	"log"
	"math/rand"
	"os"
)

func main() {
	f, err := os.OpenFile("output.txt", os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	log.SetOutput(f)

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

	//voting.Print()

	vote, err := Vote.CreateVote(account, voting, 2)

	if err != nil {
		log.Fatal(err)
	}

	//vote.Print()
	//log.Printf("Vote verified: %v", Vote.VerifyVote(vote))

	//votes := []*Vote.Vote{
	//	vote,
	//	vote,
	//	vote,
	//}

	//transaction := Transaction.CreateTransaction(vote, rand.Int())
	transactions := []*Transaction.Transaction{
		Transaction.CreateTransaction(vote, rand.Int()),
		Transaction.CreateTransaction(vote, rand.Int()),
		Transaction.CreateTransaction(vote, rand.Int()),
	}

	block1 := Block.CreateBlock(transactions, "nil")
	//block1.Print()

	transactions = []*Transaction.Transaction{
		Transaction.CreateTransaction(vote, rand.Int()),
		Transaction.CreateTransaction(vote, rand.Int()),
		Transaction.CreateTransaction(vote, rand.Int()),
	}

	block2 := Block.CreateBlock(transactions, Hash.ToSHA256(block1.ToString()))
	//block2.Print()

	blockChain := Blockchain.InitBlockchain()

	err = blockChain.AddBlock(block1)
	if err != nil {
		log.Fatal(err)
	}
	//blockChain.Print()

	err = blockChain.AddBlock(block2)
	if err != nil {
		log.Fatal(err)
	}
	blockChain.Print()
}
