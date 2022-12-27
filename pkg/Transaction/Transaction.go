package Transaction

import (
	"baby-blockchain/pkg/Vote"
	"crypto/sha256"
	"fmt"
	"log"
	"strconv"
)

type Transaction struct {
	TransactionID string
	Vote          *Vote.Vote
	Nonce         int
}

func CreateTransaction(vote *Vote.Vote, nonce int) *Transaction {
	transactionID := fmt.Sprintf("%x", sha256.Sum256([]byte(vote.ToString()+strconv.Itoa(nonce))))
	return &Transaction{TransactionID: transactionID, Vote: vote, Nonce: nonce}
}

func (tr *Transaction) ToString() string {
	return fmt.Sprintf("{TransactionID: %s; %s; Nonce: %d}", tr.TransactionID, tr.Vote.ToString(), tr.Nonce)
}

func (tr *Transaction) Print() {
	log.Println(tr.ToString())
}
