package KeyPair

import (
	"crypto/ed25519"
	"fmt"
	"io"
	"log"
)

type KeyPair struct {
	PrivateKey ed25519.PrivateKey
	PublicKey  ed25519.PublicKey
}

func GenKeyPair(rand io.Reader) *KeyPair {
	publicKey, privateKey, err := ed25519.GenerateKey(rand)
	if err != nil {
		log.Fatal("Key generation went wrong")
	}

	return &KeyPair{PrivateKey: privateKey, PublicKey: publicKey}
}

func (kp *KeyPair) ToString() string {
	return fmt.Sprintf("{Private Key: %x; Public Key: %x}", kp.PrivateKey, kp.PublicKey)
}

func (kp *KeyPair) PrintKeyPair() {
	log.Println(kp.ToString())
}
