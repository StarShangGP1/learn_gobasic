package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 变量的多种声明方式
	var hello string = "hello, golang!"
	fmt.Println(hello)

	var int2 = 33
	fmt.Println(int2)

	float3 := 3.14
	fmt.Println(float3)

	var int4, int5 = 44, 55
	fmt.Println(int4, int5)

	float4, float5 := 4.0, 5.0
	fmt.Println(float4, float5)

	var (
		int6, int7 = 6, 7
	)
	fmt.Println(int6, int7)

	reflect111(hello)

	// 常量
	// 只能用 bool，string，数字
	const PI = 3.1415926
	fmt.Println(PI)

	var a, b int = 3, 10
	fmt.Println("a+b", a+b)
	fmt.Println("a-b", a-b)
	fmt.Println("a*b", a*b)
	fmt.Println("a/b", a/b)
	fmt.Println("a%b", a%b)
	var c, d int = 1, 2
	fmt.Println(c & d)
	fmt.Println(c | d)
	fmt.Println(c ^ d)
	fmt.Println("----")
	findMaxNum()
	ifElse()
}

// 类型推断
func reflect111(i interface{}) {
	fmt.Println(reflect.TypeOf(i))
}

// 找到最大的数字
func findMaxNum() {
	arr := []int{3, 4, 4, 5, 6, 8, 3, 5, 6}
	result := -1
	for _, item := range arr {
		if result < 0 {
			result = item
			fmt.Println(result)
			//fmt.Println("----1")
		} else {
			result = result ^ item
			fmt.Println(result)
			//fmt.Println("----2")
		}
	}
	fmt.Println(result)
}

func ifElse() {
	a, b := 10, 20
	if a > b {
		fmt.Println("a>b")
	} else {
		fmt.Println("a<b")
	}
}
