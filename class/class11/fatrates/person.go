package fatrates

type Person struct {
	Name   string
	Sex    string
	Age    int
	Tall   float64
	Weight float64

	Bmi     float64
	FatRate float64
}

//func (p *Person) calcBmi() error {
//	bmi, err := gobmi.BMI(p.Weight, p.Tall)
//	if err != nil {
//		log.Println("err:", err)
//		return err
//	}
//	p.Bmi = bmi
//	return nil
//}
//
//func (p *Person) calcFatRate() {
//	p.FatRate = gobmi.CalcFatRate(p.Bmi, p.Age, p.Sex)
//}
