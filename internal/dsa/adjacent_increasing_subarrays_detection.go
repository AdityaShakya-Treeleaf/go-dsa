package dsa

import (
	"fmt"
	"slices"
)

type AdjacentIncreasingSubArraysDetection struct{}

func (ad AdjacentIncreasingSubArraysDetection) Execute() {
	nums := []int{4, 3, 2, 1}
	isSorted := slices.IsSortedFunc(nums, func(a, b int) int {
		if a == b {
			return -1
		}

		return a - b
	})

	fmt.Println(isSorted)

	testCase1 := []int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}
	k := 3
	fmt.Println(fmt.Sprintf("TestCase1: Expected: True, Actual: %t", hasIncreasingSubarrays(testCase1, k)))
}

func hasIncreasingSubarrays(nums []int, k int) bool {
	n := len(nums)

	if validate(n, k) {
		return false
	}

	for i := 0; i < n-k; i++ {
		if processCheck(nums, i, k) {
			return true
		}
	}

	return false
}

func processCheck(nums []int, i, k int) bool {
	subNums := slices.Collect(slices.Chunk(nums[i:], k))
	if len(subNums) > 1 {
		for j := 0; j < len(subNums)-1; j++ {
			if len(subNums[j]) != len(subNums[j+1]) {
				continue
			}

			if checkStrictlyIncreasing(subNums[j]) && checkStrictlyIncreasing(subNums[j+1]) {
				return true
			}
		}
	}

	return false
}

func checkStrictlyIncreasing(checkArr []int) bool {
	return slices.IsSortedFunc(checkArr, func(a, b int) int {
		if a == b {
			return -1
		}

		return a - b
	})
}

func validate(n, k int) bool {
	return n == 0 || k == 0 || n < k
}
