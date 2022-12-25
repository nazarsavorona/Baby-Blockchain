package main

import (
	"baby-blockchain/pkg/KeyPair"
	"baby-blockchain/pkg/Signature"
	"log"
)

func main() {
	keyPair := KeyPair.GenKeyPair(nil)
	keyPair.PrintKeyPair()

	message := "Love GoLang"
	signature := Signature.SignData(message, keyPair.PrivateKey)
	log.Printf("Signature verified: %v", signature.VerifySignature(message, keyPair.PublicKey))
	log.Printf("Signature verified: %v", signature.VerifySignature(message, KeyPair.GenKeyPair(nil).PublicKey))
	signature.PrintSignature()
}
