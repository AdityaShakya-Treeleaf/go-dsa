package dsa

//Minimum Number of Operations to Make All Array Elements Equal to 1
//You are given a 0-indexed array nums consisting of positive integers. You can do the following operation on the array any number of times:
//
//Select an index i such that 0 <= i < n - 1 and replace either of nums[i] or nums[i+1] with their gcd value.
//Return the minimum number of operations to make all elements of nums equal to 1. If it is impossible, return -1.
//
//The gcd of two integers is the greatest common divisor of the two integers.
//Example 1:
//
//Input: nums = [2,6,3,4]
//Output: 4
//Explanation: We can do the following operations:
//- Choose index i = 2 and replace nums[2] with gcd(3,4) = 1. Now we have nums = [2,6,1,4].
//- Choose index i = 1 and replace nums[1] with gcd(6,1) = 1. Now we have nums = [2,1,1,4].
//- Choose index i = 0 and replace nums[0] with gcd(2,1) = 1. Now we have nums = [1,1,1,4].
//- Choose index i = 2 and replace nums[3] with gcd(1,4) = 1. Now we have nums = [1,1,1,1].
//Example 2:
//Input: nums = [2,10,6,14]
//Output: -1
//Explanation: It can be shown that it is impossible to make all the elements equal to 1.

type MinOperationsToOne struct{}

func (m MinOperationsToOne) MinOperations(nums []int) int {
	n := len(nums)

	// Step 1: Count existing 1s
	countOnes := 0
	for _, num := range nums {
		if num == 1 {
			countOnes++
		}
	}

	// Step 2: If all are 1s, no operations needed
	if countOnes == n {
		return 0
	}

	// Step 3: If at least one 1 exists
	// gcd(1, x) = 1, so we need n - countOnes operations
	if countOnes > 0 {
		return n - countOnes
	}

	// Step 4: No 1s exist - check if possible
	overallGCD := nums[0]
	for i := 1; i < n; i++ {
		overallGCD = gcd(overallGCD, nums[i])
		if overallGCD == 1 {
			break // Early exit if GCD becomes 1
		}
	}

	if overallGCD > 1 {
		return -1 // Impossible to create 1
	}

	// Step 5: Find minimum operations to create first 1
	// Check all subarrays for their GCD
	minOpsToCreateOne := n

	for i := 0; i < n; i++ {
		currentGCD := nums[i]
		for j := i + 1; j < n; j++ {
			currentGCD = gcd(currentGCD, nums[j])

			if currentGCD == 1 {
				// Operations to reduce subarray [i..j] to a single 1
				// We need (j - i) operations
				minOpsToCreateOne = min(minOpsToCreateOne, j-i)
				break
			}
		}
	}

	// Total: operations to create first 1 + operations to convert rest
	return minOpsToCreateOne + (n - 1)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
