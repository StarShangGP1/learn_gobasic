package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"
)

func main() {
	countDict()
	//prime()
}

func countDict() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	lock := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(500)
	for p := 0; p < 500; p++ {
		go func() {
			defer wg.Done()
			//fmt.Print("正在统计第", p, "页, ")
			r, _ := rand.Int(rand.Reader, big.NewInt(100))
			//fmt.Println("有", r.Int64(), "字")
			lock.Lock()
			totalCount += r.Int64()
			defer lock.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("总共有：", totalCount, "字")
}

func prime() {
	startTime := time.Now()
	var result []int
	for num := 2; num <= 100; num++ {
		if isPrime(num) {
			result = append(result, num)
		}
	}
	finishTime := time.Now()
	fmt.Println(len(result))
	fmt.Println("共耗时：", finishTime.Sub(startTime))
}

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
