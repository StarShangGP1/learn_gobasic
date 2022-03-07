package calc

func CalcBodyFatRatio(sex string, age int, bmi float64) (bodyFatRatio float64) {
	gender := 0
	if sex == "男" {
		gender = 1
	} else {
		gender = 0
	}
	bodyFatRatio = (1.2*bmi + ageWeight(age)*float64(age) - 5.4 - 10.8*float64(gender)) / 100
	return
}

func ageWeight(age int) (ageWeight float64) {
	ageWeight = 0.23
	if age >= 30 {
		ageWeight = 0.22
	}
	return
}
