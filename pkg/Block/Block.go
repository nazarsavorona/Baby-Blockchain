package Block

import (
	"baby-blockchain/pkg/Hash"
	"baby-blockchain/pkg/Transaction"
	"baby-blockchain/pkg/Vote"
	"encoding/json"
	"fmt"
	"log"
)

type Block struct {
	BlockID      string                     `json:"block_id"`
	PrevHash     string                     `json:"previous_hash"`
	Transactions []*Transaction.Transaction `json:"transactions"`
}

func CreateBlock(transactions []*Transaction.Transaction, prevHash string) *Block {
	return &Block{
		BlockID:      Hash.ToSHA256(fmt.Sprintf("%v", transactions)),
		PrevHash:     prevHash,
		Transactions: transactions,
	}
}

func (b *Block) AddTransaction(transaction *Transaction.Transaction) error {
	if !Vote.VerifyVote(transaction.Vote) {
		return fmt.Errorf("invalid transaction")
	}

	b.Transactions = append(b.Transactions, transaction)
	return nil
}

func (b *Block) ToString() string {
	str, _ := json.Marshal(b)
	return string(str)
}

func (b *Block) Print() {
	log.Println(b.ToString())
}
