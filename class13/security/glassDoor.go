package main

import "fmt"

//强制绑定接口Door，必须全部实现接口Door的方法，否则编译失败
var _ Door = &GlassDoor{}

type GlassDoor struct {
}

func (d *GlassDoor) Unlock() {
	fmt.Println("GlassDoor Unlock")
}

func (d *GlassDoor) Lock() {
	fmt.Println("GlassDoor Lock")
}

func (*GlassDoor) Open() {
	fmt.Println("GlassDoor Open")
}
func (*GlassDoor) Close() {
	fmt.Println("GlassDoor Close")
}
