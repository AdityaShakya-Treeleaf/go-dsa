package dsa

import (
	"fmt"
	"log"
	"slices"
)

//Adjacent Increasing Sub Arrays Detection II

//Given an array nums of n integers, your task is to find the maximum value of k for which there exist two adjacent sub arrays of length k each, such that both sub arrays are strictly increasing. Specifically, check if there are two sub arrays of length k starting at indices a and b (a < b), where:
//Both sub arrays nums[a..a + k - 1] and nums[b..b + k - 1] are strictly increasing.
//The sub arrays must be adjacent, meaning b = a + k.
//Return the maximum possible value of k.

type AdjacentIncreasingSubArraysDetectionV2 struct{}

func (ad AdjacentIncreasingSubArraysDetectionV2) Execute() {
	log.Println("AdjacentIncreasingSubArraysDetectionV2............")
	nums := []int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}
	expected := 3

	actual := maxIncreasingSubArraysV2(nums)

	if expected != actual {
		log.Fatalf("TestCase1: Failed.")
	} else {
		log.Printf("TestCase1: Success (input: %d, expected: %d, actual: %d)", nums, expected, actual)
	}

	nums2 := []int{20, -401, -84, 680, -975, 89, 413, -6, 232, -241, -886, 835, 168, 933, 396, -770, -670, -448, 613, 264, -534, -160, 394, -361, -662, -877, -368, -584, -249, -851, 553, 125, 184, 750, 36, -420, 589, 693, 694, 698, -632, -626, -623, -472, 397, -760, 775, -561, 645, 867, -577, -13, 105, -7, 598, 309, -323, 26, 577, 269, 419, -100, -175, -894, -862, -762, 263, 66, 560, -969, -451, -948, 445, -306, -492, -859, 189, 31, 973, -191, 735, 864, 673, -517, -927, 560, 762, -661, 733, -129, -40, -351}
	expected2 := 5

	actual2 := maxIncreasingSubArraysV2(nums2)

	if expected2 != actual2 {
		log.Fatalf("TestCase2: Failed. Expected: %d, Actual: %d", expected2, actual2)
	} else {
		log.Printf("TestCase2: Success (input: %d, expected: %d, actual: %d)", nums2, expected2, actual2)
	}
	log.Println("-------------------------------------------")
}

func maxIncreasingSubArrays(nums []int) int {
	n := len(nums)

	if res, invalid := validateV2(n); invalid {
		return res
	}

	for j := n / 2; j > 0; j-- {
		if res, hasOutput := process(n, j, nums); hasOutput {
			return res
		}
	}

	return 0
}

func process(n, j int, nums []int) (int, bool) {
	fmt.Println("For: ", j)
	for l := 0; l < n-j; l++ {
		chunked := slices.Collect(slices.Chunk(nums[l:], j))
		fmt.Println(chunked)
		for i := 0; i < len(chunked)-1; i++ {
			log.Printf("Checking: %d and %d", chunked[i], chunked[i+1])
			if len(chunked[i]) != len(chunked[i+1]) {
				continue
			}
			if checkStrictlyIncreasingV2(chunked[i]) && checkStrictlyIncreasingV2(chunked[i+1]) {
				return j, true
			}
		}
	}

	return 0, false
}

func checkStrictlyIncreasingV2(nums []int) bool {
	return slices.IsSortedFunc(nums, func(a, b int) int {
		if a == b {
			return -1
		}

		return a - b
	})
}

func validateV2(n int) (int, bool) {
	if n == 0 || n == 1 {
		return 0, true
	} else if n == 2 {
		return 1, true
	}

	return 0, false
}

func maxIncreasingSubArraysV2(nums []int) int {
	n := len(nums)

	if res, invalid := validateV2(n); invalid {
		return res
	}

	maxStrictlyIncreasing := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		if i == n-1 || nums[i] >= nums[i+1] {
			maxStrictlyIncreasing[i] = 1
		} else {
			maxStrictlyIncreasing[i] = maxStrictlyIncreasing[i+1] + 1
		}
	}

	for k := n / 2; k > 0; k-- {
		for i := 0; i+2*k <= n; i++ {
			if maxStrictlyIncreasing[i] >= k && maxStrictlyIncreasing[i+k] >= k {
				return k
			}
		}
	}

	return 0
}
