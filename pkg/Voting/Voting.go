package Voting

import (
	"encoding/json"
	"fmt"
	"log"
)

type Voting struct {
	question string
	options  []string
}

func (v *Voting) ToString() string {
	options, _ := json.Marshal(v.options)
	return fmt.Sprintf("{Voting: %q, options: %s}", v.question, options)
}
func (v *Voting) Print() {
	log.Println(v.ToString())
}

func (v *Voting) OptionsAmount() int {
	return len(v.options)
}

func CreateVoting(question string, options []string) *Voting {
	return &Voting{question: question,
		options: options}
}
