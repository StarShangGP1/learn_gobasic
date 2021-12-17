package main

import "fmt"

func main() {
	//forLoop()
	//forRange()
	arr()
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

func forLoop() {
	for i := 0; i < 100; i++ {
		fmt.Printf("hello, %d\n", i)
	}
}

func forRange() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index, val := range arr {
		fmt.Println(index, val)
	}
}
