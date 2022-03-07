package calc

func CalcBMI(height, weight float64) (bmi float64) {
	if height <= 0 || weight <= 0 {
		panic("参数不能小于等于0")
	}
	return weight / (height * weight)
}
