package main

import (
	"fmt"
)

/*
	判断两条线是否平行
	提示:
	• 两点决定一条直线
	• 两条线是否平行取决于两条线的斜率是否一样
*/
func main() {
	fmt.Println("homework")
	linesCalc()
}

var p1x, p1y, p2x, p2y, p3x, p3y, p4x, p4y float64

func linesCalc() {

	points := []float64{p1x, p1y, p2x, p2y, p3x, p3y, p4x, p4y}

	// 输入坐标点
	points = inputPoints(points)
	fmt.Println("坐标点：", points)

	// 计算坐标点
	result := calcPoints(points)
	k1 := result[0]
	k2 := result[1]
	fmt.Println("平行线k1: ", k1)
	fmt.Println("平行线k2: ", k2)

	// 判断是否平行
	ifParallel(k1, k2)

}

func ifParallel(k1, k2 float64) {
	if k1 != k2 {
		fmt.Println("两条线不平行")
	} else {
		if k1 == 0 && k2 == 0 {
			fmt.Println("两条线在原点上")
		} else {
			fmt.Println("两条线平行")
		}
	}
}

func calcPoints(p []float64) []float64 {
	var result []float64
	var k1, k2 float64

	if len(p) == 0 {
		fmt.Println("没有传入变量")
		return []float64{}
	}

	p1x, p1y = p[0], p[1]
	p2x, p2y = p[2], p[3]
	p3x, p3y = p[4], p[5]
	p4x, p4y = p[6], p[7]

	numerator1 := p2y - p1y
	denominator1 := p2x - p1x
	numerator2 := p4y - p3y
	denominator2 := p4x - p3x

	if denominator1 != 0 || denominator2 != 0 {
		k1 = numerator1 / denominator1
		k2 = numerator2 / denominator2
	} else {
		k1 = numerator1 * denominator2
		k2 = numerator2 * denominator1
	}

	result = append(result, k1, k2)

	return result
}

func inputPoints(p []float64) []float64 {
	var points []float64
	if len(p) == 0 {
		fmt.Println("没有传入变量")
		return []float64{}
	}
	for i, p := range p {
		i++
		if i%2 == 0 {
			fmt.Printf("请输入第%d个点的y:", i)
		} else {
			fmt.Printf("请输入第%d个点的x:", i)
		}
		fmt.Scanln(&p)
		points = append(points, p)
	}
	return points
}
