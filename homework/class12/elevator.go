package main

import (
	"fmt"
	"time"
)

type Elevator struct {
	currentFloor int   //电梯所在楼层
	targetFloor  int   //电梯要去楼层
	floor        []int //总楼层
	person       *Person
}

func (e *Elevator) Run() {
	fmt.Printf("楼层有%d层\n", len(e.floor))
	if e.targetFloor <= 0 {
		fmt.Println("所有电梯楼层没有人请求电梯，电梯不动")
		return
	}

	fmt.Printf("电梯在%d层\n", e.currentFloor)
	for _, floor := range e.floor {
		fmt.Printf("楼层%d\n", floor)
		time.Sleep(1 * time.Second)
		if floor == e.targetFloor {
			e.currentFloor = floor
			fmt.Printf("电梯停在%d楼\n", e.currentFloor)
			break
		}
	}
	return

}
