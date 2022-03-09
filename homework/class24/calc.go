package main

import (
	"log"

	gobmi "github.com/armstrongli/go-bmi"
)

type Calc struct {
	continental string
}

func (c *Calc) BMI(person *PersonalInformationHW) (float64, error) {
	bmi, err := gobmi.BMI(person.Weight, person.Tall)
	if err != nil {
		log.Println("error when calculating bmi:", err)
		return -1, err
	}
	return bmi, nil
}

func (c *Calc) FatRate(person *PersonalInformationHW) (float64, error) {
	bmi, err := c.BMI(person)
	if err != nil {
		return -1, err
	}
	return gobmi.CalcFatRate(bmi, int(person.Age), person.Sex), nil
}
