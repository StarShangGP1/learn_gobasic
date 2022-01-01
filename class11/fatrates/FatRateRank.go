package fatrates

import (
	"math"
	"sort"
)

type FatRateRank struct {
	items []RankItem
}

func (frk *FatRateRank) InputRecord(name string, fatRate ...float64) {
	minFatRate := math.MaxFloat64
	for _, fr := range fatRate {
		if minFatRate > fr {
			minFatRate = fr
		}
	}
	found := false
	for i, item := range frk.items {
		if item.Name == name {
			if item.FatRate >= minFatRate {
				item.FatRate = minFatRate
			}
			frk.items[i] = item
			found = true
			break
		}
	}
	if !found {
		frk.items = append(frk.items, RankItem{
			Name:    name,
			FatRate: minFatRate,
		})
	}
}

func (frk *FatRateRank) GetRank(name string) (rank int, fatRate float64) {
	sort.Slice(frk.items, func(i, j int) bool {
		return frk.items[i].FatRate < frk.items[j].FatRate
	})

	frs := map[float64]struct{}{}
	for _, item := range frk.items {
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
