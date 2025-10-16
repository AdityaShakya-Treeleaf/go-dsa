package dsa

import (
	"fmt"
	"slices"
)

//3349. Adjacent Increasing Sub arrays Detection I

//Given an array nums of n integers and an integer k, determine whether there exist two adjacent sub arrays of length k such that both sub arrays are strictly increasing. Specifically, check if there are two sub arrays starting at indices a and b (a < b), where:
//Both sub arrays nums[a..a + k - 1] and nums[b..b + k - 1] are strictly increasing.
//The sub arrays must be adjacent, meaning b = a + k.
//Return true if it is possible to find two such sub arrays, and false otherwise.

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
