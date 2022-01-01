package main

import (
	"fmt"
)

func main() {
	person := Person{
		name:   "Tom",
		sex:    "男",
		age:    35,
		tall:   1.75,
		weight: 70.0,
	}
	fmt.Println(person)
}

type Person struct {
	name   string
	sex    string
	age    int
	tall   float64
	weight float64
}

type Calculator struct {
	left, right int
	op          string
}

func (c *Calculator) Add() int {
	return c.left + c.right
}

func (c *Calculator) Sub() int {
	return c.left - c.right
}

func (c *Calculator) Multiple() int {
	return c.left * c.right
}

func (c *Calculator) Divide() int {
	return c.left / c.right
}

func (c *Calculator) Reminder() int {
	return c.left % c.right
}
