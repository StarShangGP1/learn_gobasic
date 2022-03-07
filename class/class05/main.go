package main

import "fmt"

func main() {
	hello1 := hello("Tom")
	fmt.Println(hello1)
}

// 函数
func hello(name string) string {
	fmt.Println("hello, ", name)
	// TODO  做一些事
	// FIXME 修复我
	return name
}
