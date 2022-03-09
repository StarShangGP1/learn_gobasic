package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type PersonalInformationHW struct {
	Name    string  `json:"name"`
	Sex     string  `json:"sex"`
	Tall    float64 `json:"tall"`
	Weight  float64 `json:"weight"`
	Age     int64   `json:"age"`
	FatRate float64 `json:"fat_rate"`
}

// RegisterPersonalInformation 注册个人信息
func (p *PersonalInformationHW) RegisterPersonalInformation() *PersonalInformationHW {
	pi := InputFromStd{}.GetInput()
	p.Name = pi.Name
	p.Sex = pi.Sex
	p.Tall = pi.Tall
	p.Weight = pi.Weight
	p.Age = pi.Age
	return p
}

func (p *PersonalInformationHW) RegisterPersonalInformationFake() *PersonalInformationHW {
	pi := InputFromStd{}.GetInputFake()
	p.Name = pi.Name
	p.Sex = pi.Sex
	p.Tall = pi.Tall
	p.Weight = pi.Weight
	p.Age = pi.Age
	return p
}

//更新身高、体重、年龄
func (p *PersonalInformationHW) updatePersonalInformation(tall, weight float64, age int64) {
	p.Tall = tall
	p.Weight = weight
	p.Age = age
}

//获取体脂排行
func (p *PersonalInformationHW) getFatRateRank(name string) (rank int, fatRate float64) {
	fr := FatRateRank{}
	rank, fatRate = fr.GetRank(name)
	return rank, fatRate
}

//	注册时，保存所有注册信息到文件中
//	使用 JSON 格式保存
//	每行一条记录
func (p *PersonalInformationHW) savePersonalInformation(filePath string) {
	record := &Record{}
	record = record.NewRecord(filePath)
	if err := p.SavePInfo(record, p); err != nil {
		fmt.Println("保存个人信息失败: ", err)
	}
	fmt.Println("保存个人信息成功。")
}

func (p *PersonalInformationHW) SavePInfo(record *Record, pInfo *PersonalInformationHW) error {

	data, err := json.Marshal(pInfo)
	if err != nil {
		fmt.Println("marshal 出错：", err)
		return err
	}

	if err := record.WriteFileWithJson(data); err != nil {
		log.Println("写入JSON时出错：", err)
		return err
	}

	return nil
}
