package calculator

import (
	"testing"
)

func TestCalcBMI(t *testing.T) {
	//var tall, weight float64 = 1.75, 90
	var tall, weight float64 = 0, 0
	if tall == 0 || tall < 0 {
		t.Fatalf("身高不能录入0或负数, %.2f", tall)
	}
	if weight == 0 || weight < 0 {
		t.Fatalf("体重不能录入0或负数, %.2f", weight)
	}
	bmi := CalcBMI(tall, weight)
	t.Logf("体脂为: %.2f%%", bmi)
}
