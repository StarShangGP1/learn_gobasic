package main

import (
	"encoding/json"
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
	"io/ioutil"
	"learn_gobasic/pkg/apis"
	"learn_gobasic/pkg/fatRate"
	"learn_gobasic/pkg/format"
	"log"
	"os"
)

func main() {
	//caseOne()
	caseTwo()
}

var filePath = "class/class23/file.json"

func caseTwo() {
	input := &fatRate.InputFromStd{}
	calc := &fatRate.Calc{}
	rank := &fatRate.Rank{}
	record := format.Record{}
	records := record.NewRecord(filePath)

	pInfo := input.GetInputFake()
	if err := records.SavePInfo(pInfo); err != nil {
		log.Fatal("保存失败：", err)
	}
	fr, err := calc.FatRate(pInfo)
	if err != nil {
		log.Fatal("计算体脂率失败：", err)
	}
	rank.InputRecord(pInfo.Name, fr)
	rankResult, _ := rank.GetRank(pInfo.Name)
	fmt.Println("排名结果：", rankResult)
}

func caseOne() {

	pInfo := apis.PersonalInformation{
		Name:   "小强",
		Sex:    "男",
		Tall:   1.70,
		Weight: 71,
		Age:    35,
	}
	fmt.Printf("%+v\n", pInfo)
	data, err := json.Marshal(pInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("marshal 的结果是(原生)：", data)
	fmt.Println("marshal 的结果是（string）：", string(data))
	writeFile(filePath, data)
	readFile(filePath)
}

func readFile(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}
	fmt.Println("读取出来的内容是：", string(data))
	pInfo := apis.PersonalInformation{}
	err = json.Unmarshal(data, &pInfo)
	if err != nil {
		fmt.Println("反序列化失败：", err)
		return
	}
	fmt.Println("开始计算体脂信息：", pInfo)
	bmi, _ := gobmi.BMI(float64(pInfo.Weight), float64(pInfo.Tall))
	fmt.Printf("%s 的 BMI是：%v\n", pInfo.Name, bmi)
	fatRate := gobmi.CalcFatRate(bmi, int(pInfo.Age), pInfo.Sex)
	fmt.Printf("%s 的 体脂率是：%v\n", pInfo.Name, fatRate)
}

func writeFile(filePath string, data []byte) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("无法打开文件", filePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer file.Close()
	_, err = file.Write(data)
	fmt.Println(err)
}
