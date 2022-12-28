package Blockchain

import (
	"baby-blockchain/pkg/Block"
	"baby-blockchain/pkg/Hash"
	"baby-blockchain/pkg/Transaction"
	"encoding/json"
	"fmt"
	"log"
)

type Blockchain struct {
	BlockHistory []*Block.Block             `json:"block_history"`
	Transactions []*Transaction.Transaction `json:"transactions"`
}

func InitBlockchain() *Blockchain {
	return &Blockchain{
		BlockHistory: []*Block.Block{},
		Transactions: []*Transaction.Transaction{},
	}
}

func isTransactionExists(transaction *Transaction.Transaction, transactionArray []*Transaction.Transaction, index ...int) bool {
	for i, t := range transactionArray {
		if len(index) > 0 && i == index[0] {
			continue
		}
		if t.TransactionID == transaction.TransactionID {
			return true
		}
	}
	return false
}

func (b *Blockchain) validateBlock(block *Block.Block) bool {
	historyLength := len(b.BlockHistory)
	if historyLength != 0 && block.PrevHash != Hash.ToSHA256(b.BlockHistory[historyLength-1].ToString()) {
		return false
	}

	for i, transaction := range block.Transactions {
		if isTransactionExists(transaction, block.Transactions, i) || isTransactionExists(transaction, b.Transactions) {
			return false
		}
	}

	return true
}

func (b *Blockchain) AddBlock(block *Block.Block) error {
	if !b.validateBlock(block) {
		return fmt.Errorf("invalid block")
	}

	for _, transaction := range block.Transactions {
		b.Transactions = append(b.Transactions, transaction)
	}
	b.BlockHistory = append(b.BlockHistory, block)

	return nil
}

func (b *Blockchain) ToString() string {
	str, _ := json.MarshalIndent(b, "", "\t")
	return string(str)
}

func (b *Blockchain) Print() {
	log.Println(b.ToString())
}
