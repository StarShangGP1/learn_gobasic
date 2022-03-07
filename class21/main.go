package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	// 顺序查找
	//sequentialSearch()

	// 二分查找
	//binarySearch()

	// 树查找
	//treeSearch()

	// 锁复习
	lockReview()

	// channel 复习
	channelReview()
}

var totalCompare = 0

// 二分查找
func binarySearch() {
	totalCompare = 0
	newArr := append([]int64{}, sampleData...)
	startTime := time.Now()
	quickSort(&newArr, 0, len(newArr)-1)
	for i := 0; i < 1000; i++ {
		searchBinary(&newArr, 501)
		searchBinary(&newArr, 888)
		searchBinary(&newArr, 900)
		searchBinary(&newArr, 3)
	}
	finishTime := time.Since(startTime)

	//总比较次数： 37000
	//总用时： 140.39µs
	fmt.Println("二分查找")
	fmt.Println("总比较次数：", totalCompare)
	fmt.Println("总用时：", finishTime)

}

func searchBinary(arrP *[]int64, targetNum int64) bool {
	return searchHalf(arrP, 0, len(*arrP)-1, targetNum)
}

func searchHalf(arrP *[]int64, left, right int, targetNum int64) bool {
	middleIdx := (left + right) / 2
	data := (*arrP)[middleIdx]
	totalCompare++
	if data < targetNum {
		if left == right {
			return false
		}
		return searchHalf(arrP, middleIdx+1, right, targetNum)
	} else if data > targetNum {
		if left == right {
			return false
		}
		return searchHalf(arrP, left, middleIdx-1, targetNum)
	} else {
		return true
	}
}

// 顺序查找
func sequentialSearch() {
	arr := sampleData
	startTime := time.Now()
	for i := 0; i < 1000; i++ {
		search(&arr, 501)
		search(&arr, 888)
		search(&arr, 900)
		search(&arr, 3)
	}
	finishTime := time.Since(startTime)
	// 总比较次数： 3165000
	// 总用时： 5.184901ms
	fmt.Println("顺序查找")
	fmt.Println("总比较次数：", totalCompare)
	fmt.Println("总用时：", finishTime)
}

func search(arrP *[]int64, targetNum int64) bool {
	for _, v := range *arrP {
		totalCompare++
		if v == targetNum {
			return true
		}
	}
	return false
}

func generateRandomData(size int) []int64 {
	arr := make([]int64, 0, size)
	for i := 0; i < size; i++ {
		i, _ := rand.Int(rand.Reader, big.NewInt(3000))
		arr = append(arr, i.Int64())
	}
	return arr
}

// 快排
func quickSort(arr *[]int64, start, end int) *[]int64 {
	pivotIdx := (start + end) / 2
	pivotV := (*arr)[pivotIdx]
	left, right := start, end
	for left <= right {
		for (*arr)[left] < pivotV {
			left++
		}
		for (*arr)[right] > pivotV {
			right--
		}
		if left >= right {
			break
		}
		(*arr)[left], (*arr)[right] = (*arr)[right], (*arr)[left]
		left++
		right--
	}
	if left == right {
		left++
		right--
	}
	if right > start {
		quickSort(arr, start, right)
	}
	if left < end {
		quickSort(arr, left, end)
	}
	return arr
}
