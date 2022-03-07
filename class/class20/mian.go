package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arrSize := 10
	var arr []int
	for i := 0; i < arrSize; i++ {
		arr = append(arr, rand.Intn(50))
	}
	fmt.Println(arr)
	start := time.Now()
	// 冒泡排序
	//SortBubble()
	// 归并排序
	//result := mergeSort(arr)
	// 快速排序
	// 性能比较： 快排 > 归并 > 冒泡
	result := quickSort(&arr, 0, arrSize-1)
	finish := time.Since(start)
	fmt.Println(finish)
	fmt.Println(result)

	// chan 复习
	ChannelStudy()

}

// 快速排序
func quickSort(arr *[]int, start, end int) *[]int {
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

// 归并排序
func mergeSort(arr []int) []int {
	left, right := (arr)[:len(arr)/2], (arr)[len(arr)/2:]
	if len(arr) <= 2 {
		return mergeArr(left, right)
	} else {
		return mergeArr(mergeSort(left), mergeSort(right))
	}
}

func mergeArr(left, right []int) []int {
	var out []int
	leftIndex, rightIndex := 0, 0
	for {
		if leftIndex == len(left) || rightIndex == len(right) {
			break
		}
		if left[leftIndex] < right[rightIndex] {
			out = append(out, left[leftIndex])
			leftIndex++
			continue
		} else {
			out = append(out, right[rightIndex])
			rightIndex++
			continue
		}
	}
	for ; leftIndex < len(left); leftIndex++ {
		out = append(out, left[leftIndex])
	}
	for ; rightIndex < len(right); rightIndex++ {
		out = append(out, right[rightIndex])
	}
	return out
}
