package main

import (
	gobmi "github.com/armstrongli/go-bmi"
	"learn_gobasic/pkg/fatRate"
	"log"
)

type FatRateRank struct {
	Name    string
	FatRate float64
	fatRate.Rank
}

//注册个人信息
func (fr *FatRateRank) registerPersonalInformationFake(p *PersonalInformation) {
	f, err := fr.FatRates(p)
	if err != nil {
		log.Fatal("计算体脂率失败：", err)
	}
	fr.InputRecord(p.Name, f)
}

//获取体脂排行
func (fr *FatRateRank) getFatRateRank(p *PersonalInformation) (rank int, fatRate float64) {
	rank, f := fr.GetRank(p.Name)
	return rank, f
}

// GetRankWithSort 使用冒泡、快速排序分别实现体脂排序功能
func (fr *FatRateRank) GetRankWithSort(name string) (rank int, fatRate float64) {
	return 1, 1.0
}

func (fr *FatRateRank) BMI(person *PersonalInformation) (float64, error) {
	bmi, err := gobmi.BMI(person.Weight, person.Tall)
	if err != nil {
		log.Println("error when calculating bmi:", err)
		return -1, err
	}
	return bmi, nil
}

func (fr *FatRateRank) FatRates(person *PersonalInformation) (float64, error) {
	bmi, err := fr.BMI(person)
	if err != nil {
		return -1, err
	}
	return gobmi.CalcFatRate(bmi, int(person.Age), person.Sex), nil
}
