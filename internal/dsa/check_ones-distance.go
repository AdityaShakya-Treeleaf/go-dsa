package dsa

//Check If All 1's Are at Least Length K Places Away
//Given an binary array nums and an integer k, return true if all 1's are at least k places away from each other, otherwise return false.
//
//Example 1:
//Input: nums = [1,0,0,0,1,0,0,1], k = 2
//Output: true
//Explanation: Each of the 1s are at least 2 places away from each other.

type CheckOnesDistance struct{}

func (c CheckOnesDistance) KLengthApart(nums []int, k int) bool {
	onesFound := false
	dist := 0
	for _, num := range nums {
		if onesFound {
			if num == 1 {
				if dist < k {
					return false
				} else {
					dist = 0
				}
			} else {
				dist++
			}
		} else {
			onesFound = checkIfFirstOne(num)
		}
	}

	return true
}

func checkIfFirstOne(num int) bool {
	if num == 1 {
		return true
	}

	return false
}
