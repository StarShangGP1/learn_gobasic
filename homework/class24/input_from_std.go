package main

import (
	"fmt"
)

type InputFromStd struct {
}

func (InputFromStd) GetInput() *PersonalInformationHW {
	// 录入各项
	var name string
	fmt.Print("姓名：")
	fmt.Scanln(&name)

	var weight float64
	fmt.Print("体重（千克）：")
	fmt.Scanln(&weight)

	var tall float64
	fmt.Print("身高（米）：")
	fmt.Scanln(&tall)
	var age int
	fmt.Print("年龄：")
	fmt.Scanln(&age)

	sex := "男"
	fmt.Print("性别（男/女）：")
	fmt.Scanln(&sex)

	return &PersonalInformationHW{
		Name:   name,
		Sex:    sex,
		Tall:   tall,
		Weight: weight,
		Age:    int64(age),
	}
}

func (InputFromStd) GetInputFake() *PersonalInformationHW {
	return &PersonalInformationHW{
		Name:   "张三",
		Sex:    "男",
		Tall:   1.75,
		Weight: 73,
		Age:    31,
	}
}
