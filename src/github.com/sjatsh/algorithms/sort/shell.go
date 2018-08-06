package main

import "fmt"

func main() {

	count := 0
	array := []int{4, 3, 5, 6}
	n := len(array)
	var gap, i, j int

	for gap = n / 2; gap > 0; gap /= 2 {
		count++
		//插入排序简洁写法
		for i = gap; i < n; i++ {
			count++
			num := array[i]
			for j = i - gap; j >= 0 && array[j] > num; j -= gap {
				count++
				array[j+gap] = array[j]
			}
			array[j+gap] = num
		}
	}

	fmt.Println(array)
	fmt.Printf("count:%d\n", count)
}
