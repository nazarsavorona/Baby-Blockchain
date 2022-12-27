package Vote

import (
	"baby-blockchain/pkg/Account"
	"baby-blockchain/pkg/Signature"
	"baby-blockchain/pkg/Voting"
	"fmt"
	"log"
	"strconv"
)

type Vote struct {
	sender    *Account.Account
	voting    *Voting.Voting
	vote      int
	signature *Signature.Signature
}

func (v *Vote) verify() bool {
	return v.signature.VerifySignature(v.sender.ToString()+v.voting.ToString()+strconv.Itoa(v.vote), v.sender.KeyPair.PublicKey)
}

func (v *Vote) ToString() string {
	return fmt.Sprintf("{Vote: account %s, %s, vote %d, %s}",
		v.sender.ToString(),
		v.voting.ToString(),
		v.vote,
		v.signature.ToString())
}

func (v *Vote) Print() {
	log.Println(v.ToString())
}

func CreateVote(sender *Account.Account, voting *Voting.Voting, vote int) (*Vote, error) {
	if vote < 0 || vote >= voting.OptionsAmount() {
		return nil, fmt.Errorf("invalid vote index")
	}

	signature := Signature.SignData(sender.ToString()+voting.ToString()+strconv.Itoa(vote), sender.KeyPair.PrivateKey)

	return &Vote{sender: sender,
		voting:    voting,
		vote:      vote,
		signature: signature}, nil
}

func VerifyVote(vote *Vote) bool {
	return vote.verify()
}
