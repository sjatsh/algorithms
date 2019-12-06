package main

import "fmt"

/**
1. 从数列中挑出一个元素，称为"基准"（pivot）
2. 重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作。
3. 递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序。
*/
func main() {
	// http://jbcdn2.b0.upaiyun.com/2012/01/Visual-and-intuitive-feel-of-7-common-sorting-algorithms.gif
	array := []int{6, 7, 9, 3, 6, 8, 1, 9, 3}
	sort(array, 0, len(array)-1)
	fmt.Println(array)
}

// 递归排序
func sort(array []int, left, right int) {
	// 数组左边下标大于右边下标时退出
	if left > right {
		return
	}
	// 分区并排序，最终返回基准下标
	var pivotIdx = partition(array, left, right)
	// 分别递归排序基准左边和右边的数组
	sort(array, left, pivotIdx-1)
	sort(array, pivotIdx+1, right)
}

// 分区并排序
func partition(array []int, left, right int) int {
	// 记录下一个待交换位置
	var storeIndex = left
	// 直接选最右边的元素为基准元素
	var pivot = array[right]

	for i := left; i < right; i++ {
		if array[i] < pivot {
			if i != storeIndex {
				array[storeIndex], array[i] = array[i], array[storeIndex]
			}
			// 交换位置后，storeIndex 自增 1，代表下一个可能要交换的位置
			storeIndex++
		}
	}

	// 将基准元素放置到最后的正确位置上
	array[right], array[storeIndex] = array[storeIndex], array[right]
	return storeIndex
}
