package main

import (
	"fmt"
)

func main() {

	var numS = []int{1, 2, 3, 4, 5}
	m := 3

	num := combinationNum(len(numS), m)
	indexS := permutationOfIndex(len(numS), m)
	result := findByIndexS(numS, indexS)
	fmt.Println(result)

	if len(result) == num {
		fmt.Println("result success")
	} else {
		fmt.Printf("result fail, right result must be %d\n", num)
	}

}

// permutationOfIndex n选m组合索引
func permutationOfIndex(n, m int) [][]int {

	if m < 1 || m > n {
		fmt.Println("Illegal argument. Param m must between 1 and len(nums).")
		return [][]int{}
	}

	// all permutation result
	result := make([][]int, 0, combinationNum(n, m))
	// save all of index result
	indexS := make([]int, n)
	// first index array
	for i := 0; i < n; i++ {
		if i < m {
			indexS[i] = 1
		} else {
			indexS[i] = 0
		}
	}
	// 第一组组合索引
	result = addTo(result, indexS)

	for {

		find := false

		for i := 0; i < n-1; i++ {

			// if index i==1 and index i+1==0
			if 1 == indexS[i] && 0 == indexS[i+1] {

				find = true
				indexS[i], indexS[i+1] = 0, 1
				if i > 1 {
					moveOneToLeft(indexS[:i])
				}
				result = addTo(result, indexS)
				break
			}
		}

		if !find {
			break
		}
	}

	return result
}

// addTo addTo
func addTo(result [][]int, indexS []int) [][]int {

	newIndexS := make([]int, len(indexS))
	copy(newIndexS, indexS)
	return append(result, newIndexS)
}

// moveOneToLeft moveOneToLeft
func moveOneToLeft(indexS []int) {

	num := 0
	for i := 0; i < len(indexS); i++ {
		if 1 == indexS[i] {
			num++
		}
	}

	for i := 0; i < len(indexS); i++ {
		if i < num {
			indexS[i] = 1
		} else {
			indexS[i] = 0
		}
	}
}

// findByIndexS findByIndexS
func findByIndexS(numS []int, indexS [][]int) [][]int {

	if 0 == len(indexS) {
		return [][]int{}
	}

	result := make([][]int, len(indexS))

	for i, index := range indexS {

		line := make([]int, 0)
		for j, num := range index {
			if 1 == num {
				line = append(line, numS[j])
			}
		}

		result[i] = line
	}

	return result
}

// permutationNum 排列数
func permutationNum(n, m int) int {

	return factorial(n) / factorial(n-m)
}

// combinationNum 组合数
func combinationNum(n, m int) int {
	return factorial(n) / (factorial(n-m) * factorial(m))
}

// factorial 阶乘
func factorial(n int) int {

	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
