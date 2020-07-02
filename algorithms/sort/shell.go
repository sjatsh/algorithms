package main

import (
	"fmt"
)

/**
1 设置gap序列即增量序列，最后一次gap必须是1
2 将相距gap的一组数按照插入排序（注意 插入排序从第二个开始）
3 插入排序 增量为gap 而不是1
*/
func main() {
	// https://upload.wikimedia.org/wikipedia/commons/d/d8/Sorting_shellsort_anim.gif
	array := []int{6, 7, 9, 3, 6, 8, 1, 9, 3}
	n := len(array)

	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			// 前一个元素
			preIndex := i - gap
			// 当前元素
			current := array[i]

			// 有序序列从尾部开始遍历，前一个元素大于当前元素往后移动，并且每次下标往前移动gap位
			for ; preIndex >= 0 && array[preIndex] > current; preIndex -= gap {
				// 前一个元素移到后一个元素位置
				array[preIndex+gap] = array[preIndex]
			}
			// 把当前元素放入已排序序列合适位置
			array[preIndex+gap] = current
		}
	}

	fmt.Println(array)
}
