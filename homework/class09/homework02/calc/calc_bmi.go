package calculator

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
)

func CalcBMI(tall float64, weight float64) (bmi float64) {
	if tall == 0 || tall < 0 {
		fmt.Errorf("身高不能录入0或负数, %.2f", tall)
		return
	}
	if weight == 0 || weight < 0 {
		fmt.Errorf("体重不能录入0或负数, %.2f", tall)
		return
	}
	bmi, _ = gobmi.BMI(weight, tall)
	return
}
