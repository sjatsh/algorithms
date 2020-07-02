package main

import "fmt"

/**
一般来说，插入排序都采用in-place在数组上实现。具体算法描述如下：
1. 从第一个元素开始，该元素可以认为已经被排序；
2. 取出下一个元素，在已经排序的元素序列中从后向前扫描；
3. 如果该元素（已排序）大于新元素，将该元素移到下一位置；
4. 重复步骤3，直到找到已排序的元素小于或者等于新元素的位置；
5. 将新元素插入到该位置后；
6. 重复步骤2~5。
*/
func main() {

	// https://images2017.cnblogs.com/blog/849589/201710/849589-20171015225645277-1151100000.gif
	var preIndex, current int
	array := []int{6, 7, 9, 3, 6, 8, 1, 9, 3}

	for i := 1; i < len(array); i++ {

		preIndex = i - 1
		current = array[i]

		for preIndex >= 0 && array[preIndex] > current {
			array[preIndex+1] = array[preIndex]
			preIndex--
		}

		array[preIndex+1] = current
	}

	fmt.Println(array)
}
