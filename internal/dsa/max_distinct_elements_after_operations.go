package dsa

import (
	"log"
	"slices"
)

//Maximum Number of Distinct Elements After Operations

//You are given an integer array nums and an integer k.
//You are allowed to perform the following operation on each element of the array at most once:
//
//Add an integer in the range [-k, k] to the element.

//Return the maximum possible number of distinct elements in nums after performing the operations.
//Example

//Input: nums = [1,2,2,3,3,4], k = 2
//Output: 6

//Explanation:
//nums changes to [-1, 0, 1, 2, 3, 4] after performing operations on the first four elements.

type MaxDistinctElementsAfterOperations struct{}

func (mo MaxDistinctElementsAfterOperations) MaxDistinctElements(nums []int, k int) int {
	log.Println("Input: ", nums)
	uniqNum := slices.Compact(nums)
	n := len(uniqNum)

	if n == 0 {
		return 0
	}

	uniqMap := make(map[int]bool)
	for i := 0; i < n; i++ {
		for j := 0; j <= 2*k; j++ {
			distinct := uniqNum[i] - k + j
			uniqMap[distinct] = true
		}
	}

	log.Println(uniqMap)
	res := len(uniqMap)
	log.Println(res)
	return res
}
