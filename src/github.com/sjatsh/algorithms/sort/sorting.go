package main

import "fmt"

// 二分排序
// 思想上和插入排序差不多，加入了二分的思想。在前面已排序的数组中找中间数进行对比。
func main() {

	array := []int{6, 7, 9, 3, 6, 8, 1, 9, 3}

	var start, end, temp, mid int

	for i := 1; i < len(array); i++ {

		start = 0
		end = i - 1
		temp = array[i]

		for start <= end {
			mid = (start + end) / 2
			if temp < array[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
		//循环完后，start=end+1,此时start为当前插入数字所待坑位
		//把坑位给当前插入的数据挪出来
		for j := i - 1; j >= start; j-- {
			array[j+1] = array[j]
		}
		//将当前插入数字挪入它该待的坑位
		array[start] = temp
	}

	fmt.Println(array)
}
