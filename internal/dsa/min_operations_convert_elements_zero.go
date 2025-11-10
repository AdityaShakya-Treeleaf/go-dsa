package dsa

//Minimum Operations to Convert All Elements to Zero

//You are given an array nums of size n, consisting of non-negative integers. Your task is to apply some (possibly zero) operations on the array so that all elements become 0.
//In one operation, you can select a subarray [i, j] (where 0 <= i <= j < n) and set all occurrences of the minimum non-negative integer in that subarray to 0.
//Return the minimum number of operations required to make all elements in the array 0.

//Example:
//Input: nums = [3,1,2,1]
//
//Output: 3
//
//Explanation:
//
//Select subarray [1,3] (which is [1,2,1]), where the minimum non-negative integer is 1. Setting all occurrences of 1 to 0 results in [3,0,2,0].
//Select subarray [2,2] (which is [2]), where the minimum non-negative integer is 2. Setting all occurrences of 2 to 0 results in [3,0,0,0].
//Select subarray [0,0] (which is [3]), where the minimum non-negative integer is 3. Setting all occurrences of 3 to 0 results in [0,0,0,0].
//Thus, the minimum number of operations required is 3.

type MinOperationConvertToZero struct{}

func (m MinOperationConvertToZero) MinOperations(nums []int) int {
	operations := 0
	i := 0
	n := len(nums)

	for i < n {
		// Skip all zeros
		if nums[i] == 0 {
			i++
			continue
		}

		// Found a non-zero segment
		// Count how many distinct values are in this segment
		distinctValues := make(map[int]bool)

		// Process entire contiguous non-zero segment
		for i < n && nums[i] != 0 {
			distinctValues[nums[i]] = true
			i++
		}

		// Each distinct value requires one operation to eliminate
		operations += len(distinctValues)
	}

	return operations
}
