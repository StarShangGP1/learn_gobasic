package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(111)
}

var personFatRate = map[string]float64{}

func inputRecord(name string, fatRate ...float64) {
	minFatRate := math.MaxFloat64
	for _, fr := range fatRate {
		if minFatRate > fr {
			minFatRate = fr
		}
	}
	personFatRate[name] = minFatRate
}

func getRank(name string) (rank int, fatRate float64) {
	fatRateToPersonMap := map[float64][]string{}
	rankArr := make([]float64, 0, len(personFatRate))
	for name, fr := range personFatRate {
		fatRateToPersonMap[fr] = append(fatRateToPersonMap[fr], name)
		rankArr = append(rankArr, fr)
	}
	sort.Float64s(rankArr)
	for i, fr := range rankArr {
		_names := fatRateToPersonMap[fr]
		for _, _name := range _names {
			if _name == name {
				rank = i + 1
				fatRate = fr
				return
			}
		}

	}
	return 0, 0
}
