package dsa

import (
	"log"
)

type SmallestMissingNonNegativeInteger struct{}

//You are given a 0-indexed integer array nums and an integer value.

//In one operation, you can add or subtract value from any element of nums.

//For example, if nums = [1,2,3] and value = 2, you can choose to subtract value from nums[0] to make nums = [-1,2,3].
//The MEX (minimum excluded) of an array is the smallest missing non-negative integer in it.

// For example, the MEX of [-1,2,3] is 0 while the MEX of [1,0,3] is 2.
// Return the maximum MEX of nums after applying the mentioned operation any number of times.
func (s SmallestMissingNonNegativeInteger) Execute(nums []int, value int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	remainderMap := make(map[int]int)
	for _, num := range nums {
		remainder := ((num % value) + value) % value
		remainderMap[remainder]++
	}

	log.Println(remainderMap)
	mex := 0
	for {
		i := mex % value
		log.Println(i)
		if remainderMap[i] > 0 {
			remainderMap[i]--
			mex++
		} else {
			log.Println("Before break: ", mex)
			break
		}
		log.Println(remainderMap)
	}

	return mex
}
