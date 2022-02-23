package main

import (
	"fmt"
	"sync"
)

// 生产 100 个 go ,然后退出，而不是sleep
func productionGo100() {
	pCount := 100
	count := 0
	wg := sync.WaitGroup{}
	wg.Add(pCount)
	for i := 0; i < pCount; i++ {
		go func(i int) {
			wg.Done()
			count++
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
	//time.Sleep(10 * time.Second)
	fmt.Println("生产 100 个 go", count)

}
