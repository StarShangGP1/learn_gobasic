package main

import (
	"testing"
)

//案例1
//• 楼层有5层，所有电梯楼层没有人请求电梯，电梯不动
func TestElevatorCase1(t *testing.T) {

	p := &Person{}
	e := &Elevator{
		floors:       []int{1, 2, 3, 4, 5},
		currentFloor: 1,
	}
	if len(e.floors) != 5 {
		t.Fatalf("楼层不是5层，%d", e.floors)
	}
	p.PressTheFloorNumber()
	if len(p.floorNum) != 0 {
		t.Fatalf("所有电梯楼层应该没有人请求电梯 %d", p.floorNum)
	}
	e.Run(p)
}

//案例2
//• 楼层有5层，电梯在1层。三楼按电梯。电梯向三楼行进，并停在三楼。
func TestElevatorCase2(t *testing.T) {
	p := &Person{}
	e := &Elevator{
		floors:       []int{1, 2, 3, 4, 5},
		currentFloor: 1,
	}
	if len(e.floors) != 5 {
		t.Fatalf("楼层不是5层，%d", e.floors)
	}
	if e.currentFloor != 1 {
		t.Fatalf("电梯不在1层，%d", e.currentFloor)
	}
	p.PressTheFloorNumber(3)
	if p.floorNum[0] != 3 {
		t.Fatalf("应该是三楼按电梯，%d", p.floorNum)
	}
	e.Run(p)
}

//案例3
//• 楼层有5层，电梯在3层。上来一些人后，
//  目标楼层:4楼、2楼。电梯先向上到4楼，然后转 头到2楼，最后停在2楼。
//
//• 楼层有5层，电梯在3层。上来一些人后，
//  目标楼层:4楼、5楼、2楼。电梯先向上到4楼，然 后到5楼，之后转头到2楼，最后停在2楼。
func TestElevatorCase3(t *testing.T) {
	{
		p := &Person{}
		e := &Elevator{
			floors:       []int{1, 2, 3, 4, 5},
			currentFloor: 3,
		}
		if len(e.floors) != 5 {
			t.Fatalf("楼层不是5层，%d", e.floors)
		}
		if e.currentFloor != 3 {
			t.Fatalf("电梯不在3层，%d", e.currentFloor)
		}
		p.PressTheFloorNumber(4, 2)
		e.Run(p)
		if e.currentFloor != 2 {
			t.Fatalf("电梯最后停在2楼，但是目前停留在%d楼", e.currentFloor)
		}
	}

	{
		p := &Person{}
		e := &Elevator{
			floors:       []int{1, 2, 3, 4, 5},
			currentFloor: 3,
		}
		if len(e.floors) != 5 {
			t.Fatalf("楼层不是5层，%d", e.floors)
		}
		if e.currentFloor != 3 {
			t.Fatalf("电梯不在3层，%d", e.currentFloor)
		}
		p.PressTheFloorNumber(4, 5, 2)
		e.Run(p)
		if e.currentFloor != 2 {
			t.Fatalf("电梯最后停在2楼，但是目前停留在%d楼", e.currentFloor)
		}
	}
}

// 电梯行进规则改进版:
// 电梯根据当前的位置，优先向目标多的方向行进，
// 直到该方向全部送达后重新根据当前的电梯位置做 相同规则的选择。
func TestElevatorCase4(t *testing.T) {

	p := &Person{}
	e := &Elevator{
		floors:       []int{1, 2, 3, 4, 5},
		currentFloor: 2,
	}
	if len(e.floors) != 5 {
		t.Fatalf("楼层不是5层，%d", e.floors)
	}
	if e.currentFloor != 2 {
		t.Fatalf("电梯不在3层，%d", e.currentFloor)
	}
	p.PressTheFloorNumber(5, 4, 1)
	e.Run(p)
	if e.currentFloor != 1 {
		t.Fatalf("电梯最后停在1楼，但是目前停留在%d楼", e.currentFloor)
	}

}
