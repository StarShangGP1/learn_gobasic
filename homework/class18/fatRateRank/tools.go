package fatRateRank

import "math/rand"

func RandFR(min, max float64, num int) (result float64) {
	tempArr := make([]float64, num)
	for i := range tempArr {
		tempArr[i] = min + rand.Float64()*(max-min)
	}
	return tempArr[0]
}
