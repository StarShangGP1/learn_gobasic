package main

import (
	"testing"
)

//案例1
//• 楼层有5层，所有电梯楼层没有人请求电梯，电梯不动
func TestElevatorCase1(t *testing.T) {
	floor := []int{1, 2, 3, 4, 5}
	person := &Person{}
	e := &Elevator{floor: floor, person: person}
	if len(e.floor) != 5 {
		t.Fatalf("楼层不是5层，%d", e.floor)
	}
	flag := person.PressTheFloorNumber(e)
	if flag {
		t.Fatalf("所有电梯楼层应该没有人请求电梯 %t", flag)
	}
	e.Run()
}

//案例2
//• 楼层有5层，电梯在1层。三楼按电梯。电梯向三楼行进，并停在三楼。
func TestElevatorCase2(t *testing.T) {

	floor := []int{1, 2, 3, 4, 5}
	person := &Person{floorNum: 3}
	currFloor := 1
	e := &Elevator{floor: floor, person: person, currentFloor: currFloor}
	if len(e.floor) != 5 {
		t.Fatalf("楼层不是5层，%d", e.floor)
	}
	if e.currentFloor != 1 {
		t.Fatalf("电梯不在1层，%d", e.currentFloor)
	}
	person.PressTheFloorNumber(e)
	if e.targetFloor != 3 {
		t.Fatalf("应该是三楼按电梯，%d", e.targetFloor)
	}
	e.Run()
}

//案例3
//• 楼层有5层，电梯在3层。上来一些人后，
//  目标楼层:4楼、2楼。电梯先向上到4楼，然后转 头到2楼，最后停在2楼。
//• 楼层有5层，电梯在3层。上来一些人后，
//  目标楼层:4楼、5楼、2楼。电梯先向上到4楼，然 后到5楼，之后转头到2楼，最后停在2楼。
func TestElevatorCase3(t *testing.T) {
	//{
	//	cf := 3
	//	f := []int{1, 2, 3, 4, 5}
	//	p := &Person{floorNum: []int{4, 2}}
	//	e := &Elevator{currentFloor: cf, floor: f, person: p}
	//	e.Run()
	//}
	//{
	//	cf := 3
	//	f := []int{1, 2, 3, 4, 5}
	//	p := &Person{floorNum: []int{4, 5, 2}}
	//	e := &Elevator{currentFloor: cf, floor: f, person: p}
	//	e.Run()
	//}
}
