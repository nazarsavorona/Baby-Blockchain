package Account

import (
	"baby-blockchain/pkg/KeyPair"
	"baby-blockchain/pkg/Signature"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
)

type Account struct {
	AccountID string
	KeyPair   *KeyPair.KeyPair
}

func GenAccount(rand io.Reader) *Account {
	keyPair := KeyPair.GenKeyPair(rand)
	return &Account{AccountID: fmt.Sprintf("%x", sha256.Sum256([]byte(keyPair.ToString()))), KeyPair: keyPair}
}

func (acc *Account) signData(message string) *Signature.Signature {
	return Signature.SignData(message, acc.KeyPair.PrivateKey)
}

func (acc *Account) ToString() string {
	return fmt.Sprintf("{AccountID: %s; KeyPair: %s}", acc.AccountID, acc.KeyPair.ToString())
}

func (acc *Account) Print() {
	log.Println(acc.ToString())
}
