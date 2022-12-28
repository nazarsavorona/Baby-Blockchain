package Account

import (
	"baby-blockchain/pkg/Hash"
	"baby-blockchain/pkg/KeyPair"
	"baby-blockchain/pkg/Signature"
	"encoding/json"
	"io"
	"log"
)

type Account struct {
	AccountID string           `json:"account_id"`
	KeyPair   *KeyPair.KeyPair `json:"key_pair"`
}

func GenAccount(rand io.Reader) *Account {
	keyPair := KeyPair.GenKeyPair(rand)
	return &Account{AccountID: Hash.ToSHA256(keyPair.ToString()), KeyPair: keyPair}
}

func (acc *Account) signData(message string) *Signature.Signature {
	return Signature.SignData(message, acc.KeyPair.PrivateKey)
}

func (acc *Account) ToString() string {
	str, _ := json.Marshal(acc)
	return string(str)
}

func (acc *Account) Print() {
	log.Println(acc.ToString())
}
