package main

import (
	"fmt"
	"reflect"
)

func main() {
	//arr()
	//mulArr()
	//slices()
	//maps()
	maps := map[string]map[float64]string{
		"18~39": {
			0.1:  "偏瘦",
			0.16: "标准",
			0.21: "偏胖",
			0.26: "肥胖",
		},
	}
	for ages, val := range maps {
		fmt.Println(ages[:2], ages[3:5], val)
		fmt.Println(reflect.TypeOf(ages[:2]), reflect.TypeOf(ages[3:5]), val)
		for k, v := range val {
			fmt.Println(reflect.TypeOf(k), v)
		}
	}
}

// map
func maps() {
	var map1 map[string]string
	map2 := map[string]string{}
	map3 := map[string]string{"aa": "00", "bb": "11", "cc": "22"}
	// 增
	map3["dd"] = "33"
	// 删
	delete(map3, "aa")
	// 改
	map3["bb"] = "00"
	fmt.Println(map1, map2, map3)

	for k, v := range map3 {
		fmt.Println(k, v)
	}

	map4, ok := map3["aa"]
	if ok == false {
		map3["aa"] = "00"
	}
	fmt.Println(map4, ok)
	fmt.Println(map3)
}

// 切片
func slices() {
	a := []int{1, 2, 3, 4, 5}
	// 增
	a = append(a, 6)
	// 删
	a = append(a[:2], a[3:]...)
	// 改
	a[0] = 0
	fmt.Println(a)

	a2 := make([]int, 0)
	fmt.Println(a2, len(a2), cap(a2))

	s := "hello"
	fmt.Println(s)
	sByte := []byte(s)
	sByte[0] = 'a'
	s = string(sByte)
	fmt.Println(s)

	s2 := "你好"
	fmt.Println(s2)
	sByte2 := []rune(s2)
	sByte2[0] = 'a'
	s2 = string(sByte2)
	fmt.Println(s2)

}

// 多维数组
func mulArr() {

	mulArr1 := [3]string{"Tom", "man", "online"}
	fmt.Println(mulArr1)

	mulArr2 := [3][3]string{
		[3]string{"Tom", "man", "online"},
		[3]string{"Tom1", "man", "online"},
		[3]string{"Tom2", "man", "online"},
	}
	fmt.Println(mulArr2)

	mulArr3 := [...][3]string{
		[3]string{"Tom", "man", "online"},
		[3]string{"Tom1", "man", "online"},
		[3]string{"Tom2", "man", "online"},
		[3]string{"Tom3", "man", "online"},
		[3]string{"Tom4", "man", "online"},
	}
	fmt.Println(mulArr3)
}

// 数组声明
func arr() {
	var a [3]int = [3]int{}
	b := [3]int{1, 2, 3}
	c := [...]int{1, 2, 3, 4, 5}

	var d [3]int
	d = [3]int{1, 2, 3}
	d[0] = 0
	d[1] = 1
	d[2] = 2

	fmt.Println(a, b, c, d)
	fmt.Println(len(d))

	for i, val := range d {
		fmt.Println(i, val)
	}
}
