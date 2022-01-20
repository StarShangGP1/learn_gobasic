package main

import (
	"fmt"
	"sort"
	"time"
)

type Elevator struct {
	floors        []int //总楼层
	currentFloor  int   //电梯所在楼层
	upFloorsArr   []int //向上数组
	downFloorsArr []int //向下数组
	upOrDown      int   //电梯运行方向 -1 暂停 0 下 1 上
	person        *Person
}

// FloorsArrCompare 扩展
func (e *Elevator) FloorsArrCompare(arr []int) {
	if len(e.upFloorsArr) >= len(e.downFloorsArr) {
		e.upOrDown = 1
	} else {
		e.upOrDown = 0
	}
}

func (e *Elevator) Run(p *Person) {
	// 处理数组
	e.SaveArray(p)
	//方向判断
	e.FloorsArrCompare(p.floorNum)
	//e.UpOrDown(p)

	fmt.Printf("楼层有%d层\n", len(e.floors))
	fmt.Printf("电梯在%d层\n", e.currentFloor)

	for {
		if len(e.upFloorsArr) != 0 || len(e.downFloorsArr) != 0 {
			switch e.upOrDown {
			case -1:
				break
			case 0:
				e.Down()
			case 1:
				e.Up()
			default:
				break
			}
		} else {
			fmt.Println("所有电梯楼层没有人请求电梯，电梯不动")
			break
		}
	}
	return

}

func (e *Elevator) UpOrDown(p *Person) {
	if len(p.floorNum) != 0 {
		switch {
		case p.floorNum[0] < e.currentFloor:
			// down
			e.upOrDown = 0
		case p.floorNum[0] > e.currentFloor:
			// up
			e.upOrDown = 1
		default:
			// stop
			e.upOrDown = -1
		}
	}

}

func (e *Elevator) SaveArray(p *Person) {
	sort.Slice(p.floorNum, func(i, j int) bool {
		return p.floorNum[i] < p.floorNum[j]
	})
	for _, num := range p.floorNum {
		if num >= e.currentFloor {
			e.upFloorsArr = append(e.upFloorsArr, num)
		} else {
			e.downFloorsArr = append(e.downFloorsArr, num)
		}
	}
}

func (e *Elevator) Up() {

	// 对楼层从小到大排序
	sort.Sort(sort.IntSlice(e.floors))
	sort.Sort(sort.IntSlice(e.upFloorsArr))

	for _, floor := range e.floors {

		if floor < e.currentFloor {
			continue
		}
		// 到顶楼了
		if floor == len(e.floors) {
			e.upOrDown = 0
		}
		if len(e.upFloorsArr) != 0 {
			time.Sleep(1 * time.Second)
			fmt.Printf("楼层%d\n", floor)
		}

		for _, num := range e.upFloorsArr {
			if floor == num {
				e.currentFloor = num
				fmt.Printf("电梯停在%d楼\n", e.currentFloor)
				e.Open()
				e.Close()
				//移除上一个元素
				e.upFloorsArr = append(e.upFloorsArr[:0], e.upFloorsArr[1:]...)
			}
		}
	}
}

func (e *Elevator) Down() {

	// 对楼层从大到小排序
	sort.Sort(sort.Reverse(sort.IntSlice(e.floors)))
	sort.Sort(sort.Reverse(sort.IntSlice(e.downFloorsArr)))

	for _, floor := range e.floors {

		if floor > e.currentFloor {
			continue
		}
		// 到底楼了
		if floor == 1 {
			e.upOrDown = -1
		}

		time.Sleep(1 * time.Second)
		fmt.Printf("楼层%d\n", floor)

		for _, num := range e.downFloorsArr {
			if floor == num {
				e.currentFloor = floor
				fmt.Printf("电梯停在%d楼\n", e.currentFloor)
				e.Open()
				e.Close()
				//移除上一个元素
				e.downFloorsArr = append(e.downFloorsArr[:0], e.downFloorsArr[1:]...)
			}
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
