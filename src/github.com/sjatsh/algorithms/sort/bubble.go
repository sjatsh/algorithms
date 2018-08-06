package main

import "fmt"

func main() {
	// https://images2017.cnblogs.com/blog/849589/201710/849589-20171015223238449-2146169197.gif
	count := 0
	array := []int{6, 7, 9, 3, 6, 8, 1, 9, 3}
	for i := 0; i < len(array); i++ {
		count++
		for j := i + 1; j < len(array); j++ {
			count++
			if array[j] < array[i] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	fmt.Println(array)
	fmt.Printf("count:%d\n", count)
}
