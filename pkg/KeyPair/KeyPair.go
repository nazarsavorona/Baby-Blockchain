package KeyPair

import (
	"crypto/ed25519"
	"encoding/json"
	"io"
	"log"
)

type KeyPair struct {
	PrivateKey ed25519.PrivateKey `json:"private_key"`
	PublicKey  ed25519.PublicKey  `json:"public_key"`
}

func GenKeyPair(rand io.Reader) *KeyPair {
	publicKey, privateKey, err := ed25519.GenerateKey(rand)
	if err != nil {
		log.Fatal("Key generation went wrong")
	}

	return &KeyPair{PrivateKey: privateKey, PublicKey: publicKey}
}

func (kp *KeyPair) ToString() string {
	str, _ := json.Marshal(kp)
	return string(str)
}

func (kp *KeyPair) Print() {
	log.Println(kp.ToString())
}
