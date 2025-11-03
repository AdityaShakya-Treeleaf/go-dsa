package dsa

import (
	"slices"
)

// Maximize the Number of Partitions After Operations
//
// You are given a string s and an integer k.
//
// First, you are allowed to change at most one index in s to another lowercase English letter.
//
// After that, do the following partitioning operation until s is empty:
//
// Choose the longest prefix of s containing at most k distinct characters.
// Delete the prefix from s and increase the number of partitions by one. The remaining characters (if any) in s maintain their initial order.
// Return an integer denoting the maximum number of resulting partitions after the operations by optimally choosing at most one index to change.
// 1 <= s.length <= 104
// s consists only of lowercase English letters.
// 1 <= k <= 26
//func maxPartitionsAfterOperations(s string, k int) int {
//	n := len(s)
//	memo := make(map[string]int)
//
//	var dp func(pos int, changed bool, mask int) int
//	dp = func(pos int, changed bool, mask int) int {
//		// Base case
//		if pos >= n {
//			if mask > 0 {
//				return 1
//			}
//			return 0
//		}
//
//		// Memoization key
//		key := fmt.Sprintf("%d-%t-%d", pos, changed, mask)
//		if val, ok := memo[key]; ok {
//			return val
//		}
//
//		charIdx := int(s[pos] - 'a')
//		charBit := 1 << charIdx
//
//		// Count distinct characters in current mask
//		countBits := func(m int) int {
//			count := 0
//			for m > 0 {
//				count += m & 1
//				m >>= 1
//			}
//			return count
//		}
//
//		result := 0
//
//		// Option 1: Don't change this character
//		newMask := mask | charBit
//		if countBits(newMask) <= k {
//			// Continue current partition
//			result = max(result, dp(pos+1, changed, newMask))
//		} else {
//			// Start new partition
//			result = max(result, 1+dp(pos+1, changed, charBit))
//		}
//
//		// Option 2: Change this character (if not changed yet)
//		if !changed {
//			for ch := 0; ch < 26; ch++ {
//				if ch == charIdx {
//					continue
//				}
//
//				bit := 1 << ch
//				newMask := mask | bit
//
//				if countBits(newMask) <= k {
//					result = max(result, dp(pos+1, true, newMask))
//				} else {
//					result = max(result, 1+dp(pos+1, true, bit))
//				}
//			}
//		}
//
//		memo[key] = result
//		return result
//	}
//
//	return dp(0, false, 0)
//}

func maxPartitionsAfterOperations(s string, k int) int {
	if k < 1 || k > 26 {
		return 0
	}

	totalPartitions := 1
	runes := []rune(s)
	chunked := slices.Collect(slices.Chunk(runes, k))
	for _, chunk := range chunked {
		unique := getUnique(chunk)
		if len(chunk) == len(unique) {
			totalPartitions += 1
			continue
		}

		if len(chunk)-len(unique) != 1 {
			return 1
		}

	}

	return totalPartitions
}

func getUnique(nums []int32) []int32 {
	slices.Sort(nums)
	return slices.Compact(nums)
}
