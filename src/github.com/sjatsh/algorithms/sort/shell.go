package main

import "fmt"

/**
1 设置gap序列即增量序列，最后一次gap必须是1
2 将相距gap的一组数按照插入排序（注意 插入排序从第二个开始）
3 插入排序 增量为gap 而不是1
*/
func main() {

	//https://upload.wikimedia.org/wikipedia/commons/d/d8/Sorting_shellsort_anim.gif
	array := []int{6, 7, 9, 3, 6, 8, 1, 9, 3}
	n := len(array)

	for gap := n / 2; gap > 0; gap /= 2 {

		for i := gap; i < n; i++ {

			preIndex := i - gap
			current := array[i]
			for ; preIndex >= 0 && array[preIndex] > current; {
				array[preIndex+gap] = array[preIndex]
				preIndex -= gap
			}
			array[preIndex+gap] = current
		}
	}

	fmt.Println(array)
}
