package main

import "fmt"

/**
1. 找到数组里最小的元素
2. 让它和数组的第一个元素交换
3. 在剩下元素中找到最小值，与数组的第二个元素交换
4. 如此往复，直到将整个数组排序
*/
func main() {

	//https://images2017.cnblogs.com/blog/849589/201710/849589-20171015224719590-1433219824.gif
	var minIndex int
	array := []int{6, 7, 9, 3, 6, 8, 1, 9, 3}

	for i := 0; i < len(array)-1; i++ {
		minIndex = i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}
		array[i], array[minIndex] = array[minIndex], array[i]
	}

	fmt.Println(array)
}
