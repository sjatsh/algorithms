package main

import "fmt"

/**
n个记录的直接选择排序可经过n-1趟直接选择排序得到有序结果。具体算法描述如下：
1. 初始状态：无序区为R[1..n]，有序区为空；
2. 第i趟排序(i=1,2,3…n-1)开始时，当前有序区和无序区分别为R[1..i-1]和R(i..n）。
3. 该趟排序从当前无序区中-选出关键字最小的记录 R[k]，将它与有序区的第1个记录R交换，使R[1..i]和R[i+1..n)分别变为记录个数增加1个的新有序区和记录个数减少1个的新无序区；
4. n-1趟结束，数组有序化了。
*/
func main() {

	//https://images2017.cnblogs.com/blog/849589/201710/849589-20171015224719590-1433219824.gif
	var minIndex int
	array := []int{6, 7, 9, 3, 6, 8, 1, 9, 3}

	for i := 0; i < len(array)-1; i++ {
		minIndex = i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}
		array[i], array[minIndex] = array[minIndex], array[i]
	}

	fmt.Println(array)
}
