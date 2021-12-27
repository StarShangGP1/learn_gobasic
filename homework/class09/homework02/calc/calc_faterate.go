package calculator

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
)

func CalcFatRate(bmi float64, age int, sex string) (fatRate float64) {
	if bmi <= 0 {
		fmt.Errorf("bmi不能录入0或负数, %.2f%%", bmi)
		return
	}
	if age <= 0 || age > 150 {
		fmt.Errorf("年龄不能录入0或负数以及超过150的年龄, %d", age)
		return
	}
	if sex != "男" && sex != "女" {
		fmt.Errorf("性别为非男女的性别输入, %s", sex)
		return
	}
	return gobmi.CalcFatRate(bmi, age, sex)
}
