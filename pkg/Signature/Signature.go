package Signature

import (
	"crypto/ed25519"
	"encoding/json"
	"log"
)

type Signature struct {
	SignatureValue []byte `json:"signature_value"`
}

func SignData(message string, privateKey ed25519.PrivateKey) *Signature {
	return &Signature{SignatureValue: ed25519.Sign(privateKey, []byte(message))}
}

func (sig *Signature) VerifySignature(message string, publicKey ed25519.PublicKey) bool {
	return ed25519.Verify(publicKey, []byte(message), sig.SignatureValue)
}

func (sig *Signature) ToString() string {
	str, _ := json.Marshal(sig)
	return string(str)
}

func (sig *Signature) Print() {
	log.Println(sig.ToString())
}
