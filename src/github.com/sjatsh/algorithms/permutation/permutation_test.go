package permulation

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var n = 5
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

// TestCombination
func TestCombination(t *testing.T) {

	TestRandNum(t)

	// n 选 m所有组合数
	combinationNum := combinationNum(n, m).Int64()

	start := time.Now()
	// 所有组合数的索引数组
	indexS := combinationNumOfIndex(n, m)
	// 通过组合索引获取结果
	result := findByIndexS(nums, indexS)
	fmt.Println("combination use time:", time.Since(start))
	//fmt.Println(result)

	if int64(len(result)) == combinationNum {
		fmt.Printf("combination result success, count=%d\n", combinationNum)
	} else {
		fmt.Printf("combination result fail, right result must be %d\n", combinationNum)
	}
}

// TestRecursionPermutation
func TestRecursionPermutation(t *testing.T) {

	TestRandNum(t)

	// n 选 m所有组合数
	combinationNum := combinationNum(n, m).Int64()

	start := time.Now()
	// 所有组合数的索引数组
	indexS := combinationNumOfIndex(n, m)
	// 通过组合索引获取结果
	result := findByIndexS(nums, indexS)
	fmt.Println("combination use time:", time.Since(start))

	if int64(len(result)) != combinationNum {
		fmt.Errorf("combination result fail, right result must be %d\n", combinationNum)
	}

	// 计算所有排列数
	permutationNum := permutationNum(n, m).Int64()

	recursionPermutationResult := make([][]int, 0, permutationNum)
	for _, nums := range result {

		//递归加回溯法全排列
		result := recursionPermutation(nums, 0)
		recursionPermutationResult = append(recursionPermutationResult, result...)
	}
	fmt.Println("recursion permutation use time:", time.Since(start))

	//fmt.Println(recursionPermutationResult)

	if int64(len(recursionPermutationResult)) == permutationNum {
		fmt.Printf("recursion permutation result success, count=%d\n", permutationNum)
	} else {
		fmt.Errorf("recursion permutation result fail, right result must be %d\n", permutationNum)
	}
}

// TestDictionaryPermutation
func TestDictionaryPermutation(t *testing.T) {

	TestRandNum(t)

	// n 选 m所有组合数
	combinationNum := combinationNum(n, m).Int64()

	start := time.Now()
	// 所有组合数的索引数组
	indexS := combinationNumOfIndex(n, m)
	// 通过组合索引获取结果
	result := findByIndexS(nums, indexS)
	fmt.Println("combination use time:", time.Since(start))

	//fmt.Println(result)
	if int64(len(result)) != combinationNum {
		fmt.Errorf("combination result fail, right result must be %d\n", combinationNum)
	}

	// 计算所有排列数
	permutationNum := permutationNum(n, m).Int64()

	// 通过对所有组合进行全排列得到所有排列组合
	permutationResult := make([][]int, 0, permutationNum)
	for _, nums := range result {
		result := permutation(nums)
		permutationResult = append(permutationResult, result...)
	}
	fmt.Println("permutation use time:", time.Since(start))

	//fmt.Println(permutationResult)
	if int64(len(permutationResult)) == permutationNum {
		fmt.Printf("permutation result success, count=%d\n", permutationNum)
	} else {
		fmt.Printf("permutation result fail, right result must be %d\n", permutationNum)
	}
}
