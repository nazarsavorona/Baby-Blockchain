package Voting

import (
	"encoding/json"
	"log"
)

type Voting struct {
	Question string   `json:"question"`
	Options  []string `json:"options"`
}

func (v *Voting) ToString() string {
	str, _ := json.Marshal(v)
	return string(str)
}
func (v *Voting) Print() {
	log.Println(v.ToString())
}

func (v *Voting) OptionsAmount() int {
	return len(v.Options)
}

func CreateVoting(question string, options []string) *Voting {
	return &Voting{Question: question,
		Options: options}
}
