package main

import "fmt"

// Operation 技术部分
type Operation interface {
	Execute() error
}

func CheckSomething2(op []Operation) error {
	for _, v := range op {
		if err := v.Execute(); err != nil {
			return err
		}
		// ...
	}
	return nil
}

// 业务部分
var operations = []Operation{
	&operation11{},
	&operation12{},
	&operation13{},
}

type operation11 struct{}
type operation12 struct{}
type operation13 struct{}

func (o operation13) Execute() error {
	fmt.Println("op3")
	return nil
}

func (o operation12) Execute() error {
	fmt.Println("op2")
	return nil
}

func (o operation11) Execute() error {
	fmt.Println("op1")
	return nil
}
