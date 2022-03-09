package main

import (
	"log"
	"math"
	"sort"
)

type FatRateRank struct {
	Name    string
	FatRate float64
	Rank    int
	items   []FatRateRank
}

//注册个人信息
func (fr *FatRateRank) registerPersonalInformationFake(p *PersonalInformationHW) {
	calc := &Calc{}
	f, err := calc.FatRate(p)
	if err != nil {
		log.Fatal("计算体脂率失败：", err)
	}
	fr.InputRecord(p.Name, f)
}

//获取体脂排行
func (fr *FatRateRank) getFatRateRank(p *PersonalInformationHW) (int, float64) {
	rank, fatRate := fr.GetRank(p.Name)
	fr.Rank = rank
	return rank, fatRate
}

func (fr *FatRateRank) InputRecord(name string, fatRate ...float64) {
	minFatRate := math.MaxFloat64
	for _, item := range fatRate {
		if minFatRate > item {
			minFatRate = item
		}
	}

	found := false
	for i, item := range fr.items {
		if item.Name == name {
			if item.FatRate >= minFatRate {
				item.FatRate = minFatRate
			}
			fr.items[i] = item
			found = true
			break
		}
	}
	if !found {
		fr.items = append(fr.items, FatRateRank{
			Name:    name,
			FatRate: minFatRate,
		})
	}
}

func (fr *FatRateRank) GetRank(name string) (rank int, fatRate float64) {
	//sort.Slice(fr.items, func(i, j int) bool {
	//	return fr.items[i].FatRate < fr.items[j].FatRate
	//})

	fr.Bubble(&fr.items)
	//fr.QuickSort(&fr.items, 1, len(fr.items)-1)

	frs := map[float64]struct{}{}
	for _, item := range fr.items {
		frs[item.FatRate] = struct{}{}
		if item.Name == name {
			fatRate = item.FatRate
		}
	}
	rankArr := make([]float64, 0, len(frs))
	for k := range frs {
		rankArr = append(rankArr, k)
	}
	sort.Float64s(rankArr)
	for i, frItem := range rankArr {
		if frItem == fatRate {
			rank = i + 1
			break
		}
	}

	return
}

// Bubble 冒泡排序
func (fr *FatRateRank) Bubble(arr *[]FatRateRank) {
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j].FatRate > (*arr)[j+1].FatRate {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}

// QuickSort 快速排序
func (fr *FatRateRank) QuickSort(arr *[]FatRateRank, start, end int) *[]FatRateRank {
	pivotIdx := (start + end) / 2
	pivotV := (*arr)[pivotIdx]
	left, right := start, end
	for left <= right {
		for (*arr)[left].FatRate < pivotV.FatRate {
			left++
		}
		for (*arr)[right].FatRate > pivotV.FatRate {
			right--
		}
		if left >= right {
			break
		}
		(*arr)[left], (*arr)[right] = (*arr)[right], (*arr)[left]
		left++
		right--
	}
	if left == right {
		left++
		right--
	}
	if right > start {
		fr.QuickSort(arr, start, right)
	}
	if left < end {
		fr.QuickSort(arr, left, end)
	}
	return arr
}
