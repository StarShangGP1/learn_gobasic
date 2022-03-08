package main

import "fmt"

func CheckSomething() error {
	if !operation1() {
		return fmt.Errorf("op1 fails")
	}
	if !operation2() {
		return fmt.Errorf("op2 fails")
	}
	if !operation3() {
		return fmt.Errorf("op3 fails")
	}
	// ...
	return nil
}

func operation3() bool {
	fmt.Println("op3...")
	return false
}

func operation2() bool {
	fmt.Println("op2...")
	return false
}

func operation1() bool {
	fmt.Println("op1...")
	return false
}
