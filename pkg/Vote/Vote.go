package Vote

import (
	"baby-blockchain/pkg/Account"
	"baby-blockchain/pkg/Signature"
	"baby-blockchain/pkg/Voting"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

type Vote struct {
	Sender    *Account.Account     `json:"sender"`
	Voting    *Voting.Voting       `json:"voting"`
	Vote      int                  `json:"vote"`
	Signature *Signature.Signature `json:"signature"`
}

func (v *Vote) verify() bool {
	return v.Signature.VerifySignature(v.Sender.ToString()+v.Voting.ToString()+strconv.Itoa(v.Vote), v.Sender.KeyPair.PublicKey)
}

func (v *Vote) ToString() string {
	str, _ := json.Marshal(v)
	return string(str)
}

func (v *Vote) Print() {
	log.Println(v.ToString())
}

func CreateVote(sender *Account.Account, voting *Voting.Voting, vote int) (*Vote, error) {
	if vote < 0 || vote >= voting.OptionsAmount() {
		return nil, fmt.Errorf("invalid vote index")
	}

	signature := Signature.SignData(sender.ToString()+voting.ToString()+strconv.Itoa(vote), sender.KeyPair.PrivateKey)

	return &Vote{Sender: sender,
		Voting:    voting,
		Vote:      vote,
		Signature: signature}, nil
}

func VerifyVote(vote *Vote) bool {
	return vote.verify()
}
