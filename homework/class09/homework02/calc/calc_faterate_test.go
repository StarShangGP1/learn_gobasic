package calculator

import "testing"

func TestCalcFatRate(t *testing.T) {
	{
		var tall, weight float64 = 1.75, 90
		var age int = 35
		var sex string = "男"
		if age <= 0 || age > 150 {
			t.Fatalf("年龄不能录入0或负数以及超过150的年龄, %d", age)
		}
		if sex != "男" && sex != "女" {
			t.Fatalf("性别为非男女的性别输入, %s", sex)
		}

		bmi := CalcBMI(tall, weight)
		if bmi <= 0 {
			t.Fatalf("bmi不能录入0或负数, %.2f%%", bmi)
		}
		fatRate := CalcFatRate(bmi, age, sex)
		t.Logf("体脂率为:%.2f%%", fatRate)
	}
}
