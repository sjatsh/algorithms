package main

import "fmt"

func main() {

	count := 0
	var preIndex, current int
	array := []int{6, 7, 9, 3, 6, 8, 1, 9, 3}

	for i := 1; i < len(array); i++ {
		count++
		preIndex = i - 1
		current = array[i]
		for preIndex >= 0 && array[preIndex] > current {
			count++
			array[preIndex+1] = array[preIndex]
			preIndex--
		}
		array[preIndex+1] = current
	}

	fmt.Println(array)
	fmt.Printf("count:%d\n", count)
}
