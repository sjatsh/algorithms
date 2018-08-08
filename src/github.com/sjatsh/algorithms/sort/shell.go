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
	var gap, i, j int

	for gap = n / 2; gap > 0; gap /= 2 {
		//插入排序简洁写法
		for i = gap; i < n; i++ {
			num := array[i]
			for j = i - gap; j >= 0 && array[j] > num; j -= gap {
				array[j+gap] = array[j]
			}
			array[j+gap] = num
		}
	}

	fmt.Println(array)
}
