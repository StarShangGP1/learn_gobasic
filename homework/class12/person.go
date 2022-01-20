package main

import "fmt"

type Person struct {
	floorNum []int //楼层数字
}

func (p *Person) PressTheFloorNumber(num ...int) {
	fmt.Printf("有人按了%d楼的电梯\n", num)
	p.floorNum = num
}
