package main

import "fmt"

func main() {
	// 计算两个圆的面积之和并输出，精确到小数点后 3 位。
	const PI = 3.14
	var r1, r2 float64 = 4.0, 5.0
	var s1 = PI * r1 * r1
	var s2 = PI * r2 * r2
	fmt.Printf("%.3f", s1+s2)
}
