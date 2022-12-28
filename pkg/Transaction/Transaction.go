package Transaction

import (
	"baby-blockchain/pkg/Hash"
	"baby-blockchain/pkg/Vote"
	"encoding/json"
	"log"
	"strconv"
	"time"
)

type Transaction struct {
	TransactionID string     `json:"transaction_id"`
	Vote          *Vote.Vote `json:"vote"`
	Nonce         int        `json:"nonce"`
	Date          time.Time  `json:"date"`
}

func CreateTransaction(vote *Vote.Vote, nonce int) *Transaction {
	now := time.Now()
	transactionID := Hash.ToSHA256(vote.ToString() + strconv.Itoa(nonce) + now.String())
	return &Transaction{TransactionID: transactionID, Vote: vote, Nonce: nonce, Date: now}
}

func (tr *Transaction) ToString() string {
	str, _ := json.Marshal(tr)
	return string(str)
}

func (tr *Transaction) Print() {
	log.Println(tr.ToString())
}
