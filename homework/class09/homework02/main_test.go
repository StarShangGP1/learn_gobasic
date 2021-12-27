package main

import (
	calc "learn_go/homework/class09/homework02/calc"
	"testing"
)

func TestHealthSuggest(t *testing.T) {
	sex, age, tall, weight := "男", 35, 1.75, 90.0
	bmi := calc.CalcBMI(tall, weight)
	fatRate := calc.CalcFatRate(bmi, age, sex)

	var checkHealthinessFunc func(age int, fatRate float64)
	if sex == "男" {
		checkHealthinessFunc = getHealthinessSuggestionsForMale
	} else {
		checkHealthinessFunc = getHealthinessSuggestionsForFemale
	}
	getHealthinessSuggestions(age, fatRate, checkHealthinessFunc)
}
