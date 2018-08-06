package main

import "fmt"

func main() {

	//https://images2018.cnblogs.com/blog/849589/201803/849589-20180331170017421-364506073.gif
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
