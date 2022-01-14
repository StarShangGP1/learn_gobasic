package main

import "fmt"

type Person struct {
	floorNum []int //楼层数字
}

func (p *Person) PressTheFloorNumber(e *Elevator) bool {
	fmt.Printf("有人按了%d楼的电梯\n", p.floorNum)
	if len(p.floorNum) > 0 {
		e.targetFloor = p.floorNum
		return true
	}
	return false
}
