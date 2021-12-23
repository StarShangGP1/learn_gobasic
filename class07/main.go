package main

import (
	"fmt"
	"learn_go/class07/calc"

	"github.com/spf13/cobra"
	// learn_go_tools "learn.go.tools"

)

func main() {
	// 录入
	var (
		name   string
		sex    string
		tall   float64
		weight float64
		age    int
	)

	cmd := &cobra.Command{
		Use:   "healthcheck",
		Short: "体脂计算器，根据身高、体重、性别、年龄计算体脂比，并给出健康建议",
		Long:  "该体脂计算器是基于BMI的体脂计算器。。。。。。。",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("name: ", name)
			fmt.Println("sex: ", sex)
			fmt.Println("tall: ", tall)
			fmt.Println("weight: ", weight)
			fmt.Println("age: ", age)
			// 计算
			bmi := calc.CalcBMI(tall, weight)
			fatRate := calc.CalcBodyFatRatio(sex, age, bmi)
			fmt.Println("fatRate: ", fatRate)
			// 评估结果
			// todo
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "姓名")
	cmd.Flags().StringVar(&sex, "sex", "", "性别")
	cmd.Flags().Float64Var(&tall, "tall", 0, "身高")
	cmd.Flags().Float64Var(&weight, "weight", 0, "体重")
	cmd.Flags().IntVar(&age, "age", 0, "年龄")

	cmd.Execute()
}
