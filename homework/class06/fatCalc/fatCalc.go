package main

import (
	"fmt"
	calc2 "learn_go/homework/class06/calc"
	"os"
)

/*
   这是一个基于 BMI 的体脂计算器。
   体脂率是指人体内脂肪重量在人体总体重中所占的比例，
   又称 体脂百分数，它反映人体内脂肪含量的多少。

   BMI 计算法
   BMI=体重(公斤)÷(身高×身高)(米)
   体脂率:1.2×BMI+0.23×年龄-5.4-10.8×性别(男为1，女为0)
   要求:
    计算出一个人的体脂
    告诉他是偏瘦、标准、偏重、肥胖、严重肥胖

   新增需求:
   	能够主动录入姓名、性别、身高、体重、年龄信息得出这个人的体脂率
	自行控制是否继续录入信息去计算，还是退出程序

   计算多个人的平均体脂: 连续输入多人的体脂计算信息，最后输出所有人的平均体脂
--------------------------------------------------------------------------
    基于 BMI 的体脂计算器的新需求:
	不要求输入性别，同时计算两个性别的体脂
	不要求输入年龄，计算在当前身高、体重下不同年龄的体脂
	亚洲、欧洲、美洲的体质不同，年龄所占体脂比重不同，计算在相同身高、体重不变的情况下各洲的体脂
--------------------------------------------------------------------------
	(新)体脂计算器新需求:
		• 不同年龄的体脂比重不同，并且会不定期调整
		提示:
		• 计算公式中的数字是常量，需要把比重提取为函数

*/
var name, sex string
var weight, height float64
var age, gender int
var people []float64

func main() {
	fmt.Println("homework")
	// 自行控制是否继续录入信息去计算，还是退出程序
	for {
		var ifContinue string
		fmt.Print("是否录入下一个（y/n）？")
		fmt.Scanln(&ifContinue)
		if ifContinue != "y" && ifContinue != "Y" {
			//计算多个人的平均体脂:连续输入多人的体脂计算信息，最后输出所有人的平均体脂
			calcManyPeople()
			os.Exit(0)
		} else {
			fatCalc()
		}
	}

}

func fatCalc() {

	//默认数据
	//name, sex = "Tom", "男"
	//weight, height = 90.0, 1.75
	//age = 25

	//  能够主动录入姓名、性别、身高、体重、年龄信息得出这个人的体脂率
	name, sex, weight, height, age = inputFromInfo(name, sex, weight, height, age)

	if sex == "男" {
		gender = 1
	} else {
		gender = 0
	}
	// 计算体脂率
	//bodyFatRatio := calcBodyFatRatio(weight, height, age, gender)
	bmi := calc2.CalcBMI(weight, height)
	bodyFatRatio := calc2.CalcBodyFatRatio(sex, age, bmi)

	// 计算结果
	result := suggest(gender, age, bodyFatRatio)
	bodyFatRatio = bodyFatRatio * 100
	people = append(people, bodyFatRatio)

	fmt.Printf(name+" 体脂率：%0.2f%%\n", bodyFatRatio)
	fmt.Printf("身体的情况：%s", result)

}

func calcManyPeople() {
	if len(people) != 0 {
		var result float64
		for _, v := range people {
			result += v
		}
		result = result / float64(len(people))
		fmt.Printf("所有人的平均体脂率：%0.2f%%\n", result)
	} else {
		fmt.Println("没有输入数据，无法计算")
	}

}

func inputFromInfo(name, sex string, weight, height float64, age int) (string, string, float64, float64, int) {

	fmt.Print("姓名：")
	fmt.Scanln(&name)
	fmt.Print("体重：")
	fmt.Scanln(&weight)
	fmt.Print("身高：")
	fmt.Scanln(&height)
	fmt.Print("年龄：")
	fmt.Scanln(&age)
	fmt.Print("性别：")
	fmt.Scanln(&sex)

	return name, sex, weight, height, age
}

func calcBodyFatRatio(weight, height float64, age, gender int) float64 {
	bmi := weight / (height * height)
	bodyFatRatio := (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(gender)) / 100
	return bodyFatRatio
}

func suggest(gender, age int, bodyFatRatio float64) string {
	var result string
	// 男
	if gender == 1 {
		if age >= 18 && age <= 39 {
			switch {
			case bodyFatRatio <= 0.1:
				result = "偏瘦"
			case bodyFatRatio > 0.1 && bodyFatRatio <= 0.16:
				result = "标准"
			case bodyFatRatio > 0.16 && bodyFatRatio <= 0.21:
				result = "偏胖"
			case bodyFatRatio > 0.21 && bodyFatRatio <= 0.26:
				result = "肥胖"
			case bodyFatRatio > 0.26:
				result = "非常肥胖"
			default:
				fmt.Println("不再计算之内，无法计算")
			}
		} else if age >= 40 && age <= 59 {
			switch {
			case bodyFatRatio <= 0.11:
				result = "偏瘦"
			case bodyFatRatio > 0.11 && bodyFatRatio <= 0.17:
				result = "标准"
			case bodyFatRatio > 0.17 && bodyFatRatio <= 0.22:
				result = "偏胖"
			case bodyFatRatio > 0.22 && bodyFatRatio <= 0.27:
				result = "肥胖"
			case bodyFatRatio > 0.27:
				result = "非常肥胖"
			default:
				fmt.Println("不再计算之内，无法计算")
			}
		} else if age >= 60 {
			switch {
			case bodyFatRatio <= 0.13:
				result = "偏瘦"
			case bodyFatRatio > 0.13 && bodyFatRatio <= 0.19:
				result = "标准"
			case bodyFatRatio > 0.19 && bodyFatRatio <= 0.24:
				result = "偏胖"
			case bodyFatRatio > 0.24 && bodyFatRatio <= 0.29:
				result = "肥胖"
			case bodyFatRatio > 0.29:
				result = "非常肥胖"
			default:
				fmt.Println("不再计算之内，无法计算")
			}
		} else {
			fmt.Println("不再计算之内，无法计算")
		}
		// 女
	} else {
		if age >= 18 && age <= 39 {
			switch {
			case bodyFatRatio <= 0.2:
				result = "偏瘦"
			case bodyFatRatio > 0.2 && bodyFatRatio <= 0.27:
				result = "标准"
			case bodyFatRatio > 0.27 && bodyFatRatio <= 0.34:
				result = "偏胖"
			case bodyFatRatio > 0.34 && bodyFatRatio <= 0.39:
				result = "肥胖"
			case bodyFatRatio > 0.39:
				result = "非常肥胖"
			default:
				fmt.Println("不再计算之内，无法计算")
			}
		} else if age >= 40 && age <= 59 {
			switch {
			case bodyFatRatio <= 0.21:
				result = "偏瘦"
			case bodyFatRatio > 0.21 && bodyFatRatio <= 0.28:
				result = "标准"
			case bodyFatRatio > 0.28 && bodyFatRatio <= 0.35:
				result = "偏胖"
			case bodyFatRatio > 0.35 && bodyFatRatio <= 0.40:
				result = "肥胖"
			case bodyFatRatio > 0.40:
				result = "非常肥胖"
			default:
				fmt.Println("不再计算之内，无法计算")
			}
		} else if age >= 60 {
			switch {
			case bodyFatRatio <= 0.22:
				result = "偏瘦"
			case bodyFatRatio > 0.22 && bodyFatRatio <= 0.29:
				result = "标准"
			case bodyFatRatio > 0.29 && bodyFatRatio <= 0.36:
				result = "偏胖"
			case bodyFatRatio > 0.36 && bodyFatRatio <= 0.41:
				result = "肥胖"
			case bodyFatRatio > 0.41:
				result = "非常肥胖"
			default:
				fmt.Println("不再计算之内，无法计算")
			}
		} else {
			fmt.Println("不再计算之内，无法计算")
		}
	}
	return result
}
