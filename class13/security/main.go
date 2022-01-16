package main

import (
	"fmt"
)

func main() {
	security := Assets{assets: []Asset{
		&GlassDoor{},
		&WoodDoor{},
	}}
	fmt.Println("开始上班")
	security.StartWork()
	fmt.Println("8小时候，下班")
	security.StopWork()
	fmt.Println("over")
}
