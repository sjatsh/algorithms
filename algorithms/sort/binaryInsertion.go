package main

import (
	"fmt"
)

// 二分插入排序
// 算法思想简单描述：
// 在插入第i个元素时，对前面的0～i-1元素进行折半，先跟他们
// 中间的那个元素比，如果小，则对前半再进行折半，否则对后半
// 进行折半，直到left>right，然后再把第i个元素前1位与目标位置之间
// 的所有元素后移，再把第i个元素放在目标位置上。
func main() {

	array := []int{6, 7, 9, 3, 6, 8, 1, 9, 3}

	var left, mid, right, temp int
	length := len(array)

	for i := 1; i < length; i++ {

		left = 0
		right = i - 1
		temp = array[i]

		// 循环拿当前元素与中间值比较，不断调整left|right直到left>right
		// 此时left就是需要插入的位置
		for left <= right {
			mid = (left + right) / 2
			if temp < array[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		// 循环完后，left=right+1,此时left为当前插入数字所待坑位
		// 把left以及右侧数据右移，空出坑位
		for j := i; j > left; j-- {
			array[j] = array[j-1]
		}
		// 将当前插入数字挪入它该待的坑位
		array[left] = temp
	}

	fmt.Println(array)
}
