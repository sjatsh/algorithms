package main

import "fmt"

func main() {

	count := 0
	var minIndex int
	array := []int{6, 7, 9, 3, 6, 8, 1, 9, 3}

	for i := 0; i < len(array)-1; i++ {
		minIndex = i
		count++
		for j := i + 1; j < len(array); j++ {
			count++
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}
		array[i], array[minIndex] = array[minIndex], array[i]
	}

	fmt.Println(array)
	fmt.Printf("count:%d\n", count)
}
