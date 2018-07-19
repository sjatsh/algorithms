package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var n = 10
var m = 3

var nums = make([]int, n)

func TestRandNum(t *testing.T) {

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < n; i++ {

		randNum := r.Intn(n)
		for {
			find := false
			for j := 0; j < i; j++ {
				if randNum == nums[j] {
					find = true
					randNum = r.Intn(n)
				}
			}
			if !find {
				break
			}
		}
		nums[i] = randNum
	}
	fmt.Println(nums)
}

func TestCombination(t *testing.T) {

	TestRandNum(t)

	// n 选 m所有组合数
	combinationNum := combinationNum(len(nums), m)

	start := time.Now()
	// 所有组合数的索引数组
	indexS := combinationNumOfIndex(len(nums), m)
	// 通过组合索引获取结果
	result := findByIndexS(nums, indexS)
	fmt.Println("combination use time:", time.Since(start))

	fmt.Println(result)
	if len(result) == combinationNum {
		fmt.Println("combination result success")
	} else {
		fmt.Printf("combination result fail, right result must be %d\n", combinationNum)
	}
}

func TestRecursionPermutation(t *testing.T) {

	TestRandNum(t)

	// n 选 m所有组合数
	combinationNum := combinationNum(len(nums), m)

	start := time.Now()
	// 所有组合数的索引数组
	indexS := combinationNumOfIndex(len(nums), m)
	// 通过组合索引获取结果
	result := findByIndexS(nums, indexS)
	fmt.Println("combination use time:", time.Since(start))

	fmt.Println(result)
	if len(result) == combinationNum {
		fmt.Println("combination result success")
	} else {
		fmt.Printf("combination result fail, right result must be %d\n", combinationNum)
	}

	// 计算所有排列数
	permutationNum := permutationNum(len(nums), m)

	recursionPermutationResult := make([][]int, 0, permutationNum)
	for _, nums := range result {

		//递归加回溯法全排列
		result := recursionPermutation(nums, 0)
		recursionPermutationResult = append(recursionPermutationResult, result...)
	}
	fmt.Println("recursion permutation use time:", time.Since(start))

	fmt.Println(recursionPermutationResult)
	if len(recursionPermutationResult) == permutationNum {
		fmt.Println("recursion permutation result success")
	} else {
		fmt.Printf("recursion permutation result fail, right result must be %d\n", permutationNum)
	}
}

func TestDictionaryPermutation(t *testing.T) {

	TestRandNum(t)

	// n 选 m所有组合数
	combinationNum := combinationNum(len(nums), m)

	start := time.Now()
	// 所有组合数的索引数组
	indexS := combinationNumOfIndex(len(nums), m)
	// 通过组合索引获取结果
	result := findByIndexS(nums, indexS)
	fmt.Println("combination use time:", time.Since(start))

	fmt.Println(result)
	if len(result) == combinationNum {
		fmt.Println("combination result success")
	} else {
		fmt.Printf("combination result fail, right result must be %d\n", combinationNum)
	}

	// 计算所有排列数
	permutationNum := permutationNum(len(nums), m)

	// 通过对所有组合进行全排列得到所有排列组合
	permutationResult := make([][]int, 0, permutationNum)
	for _, nums := range result {
		result := permutation(nums)
		permutationResult = append(permutationResult, result...)
	}
	fmt.Println("permutation use time:", time.Since(start))

	fmt.Println(permutationResult)
	if len(permutationResult) == permutationNum {
		fmt.Println("permutation result success")
	} else {
		fmt.Printf("permutation result fail, right result must be %d\n", permutationNum)
	}
}
