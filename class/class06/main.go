package main

import "fmt"

func main() {
	//fib := fibonacci(10)
	//fmt.Println(fib)
	//guess(1,50)

	defer func() {
		fmt.Println("hello")
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("发生严重错误，以截获")
		}
	}()
}

// 递归
// 斐波那契数列
func fibonacci(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 猜数字
func guess(left, right uint) {
	guessed := (left + right) / 2
	var getFromInput string
	fmt.Println("我猜是：", guessed)
	fmt.Print("如果高了，输入 1，如果低了，输入 0；对了，输入 9：")
	fmt.Scanln(&getFromInput)
	switch getFromInput {
	case "1":
		if left == right {
			fmt.Println("你耍赖了")
			return
		}
		guess(left, guessed-1)
	case "0":
		if left == right {
			fmt.Println("你耍赖了")
			return
		}
		guess(guessed+1, right)
	case "9":
		fmt.Println("你猜的数字是：", guessed)
	default:
		fmt.Println("未知错误")
	}
}
