package main

import (
	"fmt"
	"sort"
)

// permutationOfIndex n选m组合索引
func combinationNumOfIndex(n, m int) [][]int {

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

// permutation 全排列（字典法）
func permutation(nums []int) [][]int {

	result := make([][]int, 0)

	sort.Slice(nums, func(j, k int) bool {
		return nums[j] < nums[k]
	})
	result = addTo(result, nums)

	for {

		find := false
		pos1, pos2 := 0, 0
		for j := len(nums) - 2; j >= 0; j-- {

			if nums[j] < nums[j+1] {
				find = true
				pos1 = j
				break
			}
		}

		if !find {
			break
		}

		for j := len(nums) - 1; j > pos1; j-- {

			if nums[j] >= nums[pos1] {
				pos2 = j
				break
			}
		}

		nums[pos1], nums[pos2] = nums[pos2], nums[pos1]

		for j, k := pos1+1, len(nums)-1; j < k; j, k = j+1, k-1 {

			nums[j], nums[k] = nums[k], nums[j]
		}

		result = addTo(result, nums)

	}

	return result
}

// recursionPermutation 递归排序
func recursionPermutation(nums []int, index int) [][]int {

	result := make([][]int, 0)
	if index == len(nums)-1 {
		result = addTo(result, nums)
		return result
	}

	for i := index; i < len(nums); i++ {

		nums[i], nums[index] = nums[index], nums[i]
		result = append(result, recursionPermutation(nums, index+1)...)
		nums[i], nums[index] = nums[index], nums[i]
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
