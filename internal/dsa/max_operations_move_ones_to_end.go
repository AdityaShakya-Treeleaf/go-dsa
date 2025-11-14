package dsa

//Maximum Number of Operations to Move Ones to the End

//You are given a binary string s.

//You can perform the following operation on the string any number of times:

//Choose any index i from the string where i + 1 < s.length such that s[i] == '1' and s[i + 1] == '0'.
//Move the character s[i] to the right until it reaches the end of the string or another '1'. For example, for s = "010010", if we choose i = 1, the resulting string will be s = "000110".
//Return the maximum number of operations that you can perform.

//Example 1:
//Input: s = "1001101"
//Output: 4
//Explanation:
//We can perform the following operations:
//Choose index i = 0. The resulting string is s = "0011101".
//Choose index i = 4. The resulting string is s = "0011011".
//Choose index i = 3. The resulting string is s = "0010111".
//Choose index i = 2. The resulting string is s = "0001111".

//Example 2:
//Input: s = "00111"
//Output: 0

type MaxOperationsMoveOnesToEnd struct{}

func (m MaxOperationsMoveOnesToEnd) MaxOperations(s string) int {
	operations := 0
	onesCount := 0
	n := len(s)

	for i := 0; i < n; i++ {
		if s[i] == '1' {
			onesCount++
		} else if s[i] == '0' {
			// If we see a '0' and there were '1's before it
			// All those '1's will need to pass this '0'
			if i > 0 && s[i-1] == '1' {
				operations += onesCount
			}
		}
	}

	return operations
}

//
//func (m MaxOperationsMoveOnesToEnd) MaxOperations(s string) int {
//	strArr := []rune(s)
//	n := len(strArr)
//	if n <= 1 {
//		return 0
//	}
//
//	return getMaxOps(s, 0)
//}
//
//func getMaxOps(strArr string, mo int) int {
//	reLoop := false
//	//48 - 0, 49 - 1
//	prevVal := 0
//	findOne := false
//	firstOne := -1
//	resultArr := []rune(strArr)
//	for ind, val := range strArr {
//		if findOne {
//			if val == 49 || ind == len(resultArr) {
//				resultArr = slices.Insert(resultArr, ind, 49)
//				break
//			}
//		}
//		if val == 48 {
//			if prevVal == 49 {
//				resultArr = slices.Delete(resultArr, ind-1, ind)
//				mo++
//				reLoop = true
//				findOne = true
//			}
//			prevVal = 48
//		} else if val == 49 {
//			if firstOne == -1 {
//				firstOne = ind
//			}
//
//			prevVal = 49
//		}
//	}
//
//	if reLoop {
//		return getMaxOps(string(resultArr[firstOne:]), mo)
//	} else {
//		return mo
//	}
//}
