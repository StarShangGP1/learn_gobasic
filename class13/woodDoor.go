package main

import "fmt"

//强制绑定接口Door，必须全部实现接口Door的方法，否则编译失败
var _ Door = &WoodDoor{}

type WoodDoor struct {
}

func (d *WoodDoor) Unlock() {
	fmt.Println("WoodDoor Unlock")
}

func (d *WoodDoor) Lock() {
	fmt.Println("WoodDoor Lock")
}

func (*WoodDoor) Open() {
	fmt.Println("WoodDoor Open")
}
func (*WoodDoor) Close() {
	fmt.Println("WoodDoor Close")
}
