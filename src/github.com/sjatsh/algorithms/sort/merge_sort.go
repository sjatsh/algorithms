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

// 归并排序：基本思想，数组分为左右两边，对左右两边分别排序，再对两个有序数组进行合并，结束标志：分割的数组只剩下一个元素
// 1：二分法找出中间值，划分左右两个数组，对左右两个数组分别进行排序
// 2：重复1继续对数组进行分割，直到只剩下一个元素
// 3：对每个划分后的有序数组进行合并
func mergeSort(array, temp []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	// 左边排序
	mergeSort(array, temp, left, mid)
	// 右边排序
	mergeSort(array, temp, mid+1, right)
	// 将两个有序序列合并
	merge(array, temp, left, mid, right)
}

// 将有二个有序数列a[left...mid]和a[mid+1...last]合并。
func merge(array []int, temp []int, left, mid, right int) {

	// 第一个有序序列第一个下标
	f1 := left
	// 第二个有序序列第一个下标
	f2 := mid + 1

	// temp结果数组下标
	k := 0

	// 遍历两个有序数组并排序后放入temp数组
	for ; f1 <= mid && f2 <= right; k++ {
		if array[f1] <= array[f2] {
			temp[k] = array[f1]
			f1++
			continue
		}
		temp[k] = array[f2]
		f2++
	}

	// 拷贝可能剩余的数据
	for ; f1 <= mid; f1, k = f1+1, k+1 {
		temp[k] = array[f1]
	}
	for ; f2 <= right; f2, k = f2+1, k+1 {
		temp[k] = array[f2]
	}

	// 排好序的数据重新放回原数组
	for i := 0; i < k; i++ {
		array[left+i] = temp[i]
	}
}
