package main

import (
	"fmt"
)

func main() {
	array := []int{6, 7, 9, 3, 6, 8, 1, 9, 3}
	temp := make([]int, len(array))
	mergeSort(array, temp, 0, len(array)-1)
	fmt.Println(array)
}

func mergeSort(array []int, temp []int, first, last int) {

	if first < last {
		mid := (first + last) / 2
		// 左边排序
		mergeSort(array, temp, first, mid)
		// 右边排序
		mergeSort(array, temp, mid+1, last)
		// 将两个有序序列合并
		merge(array, temp, first, mid, last)
	}
}

// 将有二个有序数列a[first...mid]和a[mid+1...last]合并。
func merge(array []int, temp []int, first, mid, last int) {

	f := first
	mid1 := mid + 1
	k := 0

	for ; f <= mid && mid1 <= last; k++ {
		if array[f] <= array[mid1] {
			temp[k] = array[f]
			f++
		} else {
			temp[k] = array[mid1]
			mid1++
		}
	}

	for ; f <= mid; f, k = f+1, k+1 {
		temp[k] = array[f]
	}
	for ; mid1 <= last; mid1, k = mid1+1, k+1 {
		temp[k] = array[mid1]
	}
	for i := 0; i < k; i++ {
		array[first+i] = temp[i]
	}
}
