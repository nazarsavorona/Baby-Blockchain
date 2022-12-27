package Account

import (
	"baby-blockchain/pkg/Hash"
	"baby-blockchain/pkg/KeyPair"
	"baby-blockchain/pkg/Signature"
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
	return &Account{AccountID: Hash.ToSHA256(keyPair.ToString()), KeyPair: keyPair}
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
