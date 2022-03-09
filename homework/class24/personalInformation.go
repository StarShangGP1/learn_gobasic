package main

import (
	"encoding/json"
	"fmt"
	"learn_gobasic/pkg/apis"
	"learn_gobasic/pkg/fatRate"
	"learn_gobasic/pkg/format"
	"log"
)

type PersonalInformation struct {
	Name   string  `json:"name"`
	Sex    string  `json:"sex"`
	Tall   float64 `json:"tall"`
	Weight float64 `json:"weight"`
	Age    int64   `json:"age"`
	apis.PersonalInformation
	fatRate.InputFromStd
	format.Record
	FatRateRank
}

//注册个人信息
func (p *PersonalInformation) registerPersonalInformation() *PersonalInformation {
	pi := p.GetInput()
	p.Name = pi.Name
	p.Sex = pi.Sex
	p.Tall = float64(pi.Tall)
	p.Weight = float64(pi.Weight)
	p.Age = pi.Age
	return p
}

func (p *PersonalInformation) registerPersonalInformationFake() *PersonalInformation {
	pi := p.GetInputFake()
	p.Name = pi.Name
	p.Sex = pi.Sex
	p.Tall = float64(pi.Tall)
	p.Weight = float64(pi.Weight)
	p.Age = pi.Age
	return p
}

//更新身高、体重、年龄
func (p *PersonalInformation) updatePersonalInformation(tall, weight float64, age int64) *PersonalInformation {
	p.Tall = tall
	p.Weight = weight
	p.Age = age
	return p
}

//获取体脂排行
func (p *PersonalInformation) getFatRateRank(name string) (rank int, fatRate float64) {
	r, f := p.GetRank(name)
	return r, f
}

//	注册时，保存所有注册信息到文件中
//	使用 JSON 格式保存
//	每行一条记录
func (p *PersonalInformation) savePersonalInformation(filePath string) {
	p.NewRecord(filePath)
	if err := p.SavePInfo(p); err != nil {
		fmt.Println("保存个人信息失败: ", err)
	}
	fmt.Println("保存个人信息成功。")
}

func (p *PersonalInformation) SavePInfo(pInfo *PersonalInformation) error {
	{
		data, err := json.Marshal(pInfo)
		if err != nil {
			fmt.Println("marshal 出错：", err)
			return err
		}
		if err := p.WriteFileWithJson(data); err != nil {
			log.Println("写入JSON时出错：", err)
			return err
		}
	}
	return nil
}
