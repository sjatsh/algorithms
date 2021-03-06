package main

import "fmt"

/**
冒泡排序
1. 比较相邻的元素。如果第一个比第二个大，就交换它们两个；
2. 对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对，这样在最后的元素应该会是最大的数；
3. 针对所有的元素重复以上的步骤，除了最后一个；
4. 重复步骤1~3，直到排序完成。
*/
func main() {
	// https://images2017.cnblogs.com/blog/849589/201710/849589-20171015223238449-2146169197.gif
	array := []int{6, 7, 9, 3, 6, 8, 1, 9, 3}
	l := len(array)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	fmt.Println(array)
}
