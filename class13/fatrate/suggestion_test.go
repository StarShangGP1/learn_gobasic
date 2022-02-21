package main

import (
	"testing"
)

func Test_fatRateSuggestion_GetSuggestion(t *testing.T) {
	sugg := GetFatRateSuggestion()
	tests := []Person{
		{
			sex:     "男",
			age:     35,
			fatRate: 0.24,
		},
	}
	if got := sugg.GetSuggestion(&tests[0]); got != "肥胖" {
		t.Fail()
	}
}

func Test_fatRateSuggestion_GetSuggestion1(t *testing.T) {
	sugg := GetFatRateSuggestion()
	type args struct {
		person *Person
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.0}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.01}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.02}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.03}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.04}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.05}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.06}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.07}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.08}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.09}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.10}}, want: "标准"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.11}}, want: "标准"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.12}}, want: "标准"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.13}}, want: "标准"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.14}}, want: "标准"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.15}}, want: "标准"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.16}}, want: "偏重"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.17}}, want: "偏重"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.18}}, want: "偏重"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.19}}, want: "偏重"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.20}}, want: "偏重"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.21}}, want: "肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.22}}, want: "肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.23}}, want: "肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.24}}, want: "肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.25}}, want: "肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.26}}, want: "非常肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.27}}, want: "非常肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.28}}, want: "非常肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.29}}, want: "非常肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.30}}, want: "非常肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.31}}, want: "非常肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.32}}, want: "非常肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.33}}, want: "非常肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.34}}, want: "非常肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.35}}, want: "非常肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.36}}, want: "非常肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.37}}, want: "非常肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.38}}, want: "非常肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.39}}, want: "非常肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.40}}, want: "非常肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.41}}, want: "非常肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.42}}, want: "非常肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.43}}, want: "非常肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.44}}, want: "非常肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.94}}, want: "非常肥胖"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := sugg
			if got := s.GetSuggestion(tt.args.person); got != tt.want {
				t.Errorf("GetSuggestion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_BasalMetabolism(t *testing.T) {
	p := &Person{
		name:   "小强",
		sex:    "男",
		tall:   1.73,
		weight: 73,
		age:    26,
	}
	calc := &Calc{}
	calc.BasalMetabolism(p)
	t.Log(p.bm)
	//2100 kcal
	// 制定计划
	// 一个月减4kg，一周减1kg
	// 1kg的脂肪等于 7700 kcal 除以7天，每天减少 1100 kcal
	// 2100-1100=1000 kcal
	// 1 kcal = 4.184 j; 1000 = 4184 j
	// 人体组成部分三大要素：蛋白质(40%)400kcal、碳水(40%)400kcal、脂肪(20%)200kcal
}
