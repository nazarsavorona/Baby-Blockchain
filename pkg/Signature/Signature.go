package Signature

import (
	"crypto/ed25519"
	"fmt"
	"log"
)

type Signature struct {
	SignatureValue []byte
}

func SignData(message string, privateKey ed25519.PrivateKey) *Signature {
	return &Signature{SignatureValue: ed25519.Sign(privateKey, []byte(message))}
}

func (sig *Signature) VerifySignature(message string, publicKey ed25519.PublicKey) bool {
	return ed25519.Verify(publicKey, []byte(message), sig.SignatureValue)
}

func (sig *Signature) ToString() string {
	return fmt.Sprintf("Signature value: %x", sig.SignatureValue)
}

func (sig *Signature) PrintSignature() {
	log.Println(sig.ToString())
}
