package main

import (
	"log"

	gobmi "github.com/armstrongli/go-bmi"
)

type Calc struct {
	continental string
}

func (Calc) BMI(person *Person) error {
	bmi, err := gobmi.BMI(person.weight, person.tall)
	if err != nil {
		log.Println("error when calculating bmi:", err)
		return err
	}
	person.bmi = bmi
	return nil
}

func (Calc) FatRate(person *Person) error {
	person.fatRate = gobmi.CalcFatRate(person.bmi, person.age, person.sex)
	return nil
}

// BasalMetabolism 基础代谢
//[90+4.8X身高+13.4X体重-5.7X年龄]X1.2
func (Calc) BasalMetabolism(person *Person) error {
	person.bm = (90 + 4.8*person.tall*100 + 13.4*person.weight - 5.7*float64(person.age)) * 1.2
	return nil
}
