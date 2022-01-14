package main

import (
	"fmt"
	"sort"
	"time"
)

type Elevator struct {
	currentFloor int   //电梯所在楼层
	targetFloor  []int //电梯要去楼层
	floor        []int //总楼层
	downFloor    int
	person       *Person
}

func (e *Elevator) Run() {
	fmt.Printf("楼层有%d层\n", len(e.floor))
	if len(e.targetFloor) <= 0 {
		fmt.Println("所有电梯楼层没有人请求电梯，电梯不动")
		return
	}
	fmt.Printf("电梯在%d层\n", e.currentFloor)
	e.Up()
	if len(e.targetFloor) > 1 {
		e.Down()
	}

	return

}

func (e *Elevator) Up() {
	floors := e.floor
	flag := false
	for _, floor := range floors {
		if e.currentFloor > floor {
			continue
		}
		time.Sleep(1 * time.Second)
		for _, num := range e.targetFloor {

			if floor != num {
				e.downFloor = num
			}
			if floor == num {
				e.currentFloor = floor
				fmt.Printf("楼层%d\n", floor)
				fmt.Printf("电梯停在%d楼\n", e.currentFloor)
				e.Open()
				e.Close()
				if e.currentFloor != num {
					flag = true
					break
				}
			}
		}
		if flag {
			break
		}
	}
}

func (e *Elevator) Down() {
	floors := e.floor
	e.targetFloor = []int{e.downFloor}
	sort.Sort(sort.Reverse(sort.IntSlice(floors)))
	flag := false
	for _, floor := range floors {
		if e.currentFloor < floor && e.currentFloor != floor {
			continue
		}
		fmt.Printf("楼层%d\n", floor)
		time.Sleep(1 * time.Second)
		for _, num := range e.targetFloor {
			if floor == num {
				e.currentFloor = floor
				fmt.Printf("电梯停在%d楼\n", e.currentFloor)
				e.Open()
				e.Close()
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
}

func (Elevator) Open() {
	time.Sleep(1 * time.Second)
	fmt.Println("开门")
}

func (Elevator) Close() {
	time.Sleep(1 * time.Second)
	fmt.Println("关门")
}

func (Elevator) SliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
