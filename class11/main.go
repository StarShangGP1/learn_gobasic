package main

import (
	"fmt"
	"learn_go/class11/fatrates"
)

func main() {
	frSvc := fatrates.FatRateService{S: fatrates.GetFatRateSuggestion()}
	p := &fatrates.Person{
		Name:   "小强",
		Sex:    "男",
		Tall:   1.7,
		Weight: 70,
		Age:    35,
	}
	fmt.Println(frSvc.GiveSuggestionToPerson(p))
}
