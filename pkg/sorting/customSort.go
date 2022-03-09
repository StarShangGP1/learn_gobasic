package sorting

// Bubble 冒泡排序
func Bubble(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
		//fmt.Println("中间状态：", *arr)
	}
	//fmt.Println("最终状态：", *arr)
}

// QuickSort 快速排序
func QuickSort(arr *[]int, start, end int) *[]int {
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
		QuickSort(arr, start, right)
	}
	if left < end {
		QuickSort(arr, left, end)
	}
	return arr
}
