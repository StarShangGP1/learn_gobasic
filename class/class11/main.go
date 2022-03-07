package main

import (
	"fmt"
	fatrates2 "learn_go/class/class11/fatrates"
)

func main() {
	frSvc := fatrates2.FatRateService{S: fatrates2.GetFatRateSuggestion()}
	p := &fatrates2.Person{
		Name:   "小强",
		Sex:    "男",
		Tall:   1.7,
		Weight: 70,
		Age:    35,
	}
	fmt.Println(frSvc.GiveSuggestionToPerson(p))
}
