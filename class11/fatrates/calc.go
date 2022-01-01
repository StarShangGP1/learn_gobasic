package fatrates

import (
	gobmi "github.com/armstrongli/go-bmi"
	"log"
)

type Calc struct{}

func (c *Calc) calcBmi(p *Person) error {
	bmi, err := gobmi.BMI(p.Weight, p.Tall)
	if err != nil {
		log.Println("err:", err)
		return err
	}
	p.Bmi = bmi
	return nil
}

func (c *Calc) calcFatRate(p *Person) {
	p.FatRate = gobmi.CalcFatRate(p.Bmi, p.Age, p.Sex)
}

//func (Calc) BMI(person *Person) error {
//	bmi, err := gobmi.BMI(person.weight, person.tall)
//	if err != nil {
//		log.Println("err:", err)
//		return err
//	}
//	person.bmi = bmi
//	return nil
//}
//
//func (Calc) FatRate(person *Person) error {
//	person.fatRate = gobmi.CalcFatRate(person.bmi, person.age, person.sex)
//	return nil
//}
