package dsa

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

// Unique Length-3 Palindromic Subsequences
// Given a string s, return the number of unique palindromes of length three that are a subsequence of s.
//
// Note that even if there are multiple ways to obtain the same subsequence, it is still only counted once.
// A palindrome is a string that reads the same forwards and backwards.
// A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.
//
// For example, "ace" is a subsequence of "abcde".
// Example 1:
//
// Input: s = "aabca"
// Output: 3
// Explanation: The 3 palindromic subsequences of length 3 are:
// - "aba" (subsequence of "aabca")
// - "aaa" (subsequence of "aabca")
// - "aca" (subsequence of "aabca")
// Example 2:
//
// Input: s = "adc"
// Output: 0
// Explanation: There are no palindromic subsequences of length 3 in "adc".
// Example 3:
//
// Input: s = "bbcbaba"
// Output: 4
// Explanation: The 4 palindromic subsequences of length 3 are:
// - "bbb" (subsequence of "bbcbaba")
// - "bcb" (subsequence of "bbcbaba")
// - "bab" (subsequence of "bbcbaba")
// - "aba" (subsequence of "bbcbaba")
type CountPalindromeSubsequence struct{}

func (c CountPalindromeSubsequence) CountPalindromicSubsequence(s string) int {
	n := len(s)
	if n < 2 {
		return 0
	}

	if possible, returnValue := validateIfPalindromePossible(s); !possible {
		return returnValue
	}

	palindromes := checkPalindrome(n, s)
	fmt.Println(len(palindromes))
	return len(palindromes)
}

func checkPalindrome(n int, s string) map[string]bool {
	palindromes := make(map[string]bool)
	checkMap := make(map[string]int)
	for i := 0; i < n; i++ {
		curr := string(s[i])
		occurrence, ok := checkMap[curr]
		if !ok {
			checkMap[curr] = 0
		} else if occurrence >= 26 {
			continue
		}

		if i+2 <= n {
			checkHalf := s[i+2:]
			foundInd := strings.LastIndex(checkHalf, curr)
			if foundInd != -1 {
				foundInd += i + 2
				added := make(map[string]bool)
				for j := i + 1; j < foundInd; j++ {
					currChar := string(s[j])
					if added[currChar] {
						continue
					}

					if checkMap[curr] > 26 {
						break
					}
					currPalindrome := curr + currChar + curr
					_, exists := palindromes[curr+":"+currPalindrome]
					if !exists {
						added[currChar] = true
						checkMap[curr]++
						palindromes[curr+":"+currPalindrome] = true
					}
				}
			}
		}
	}

	fmt.Println(checkMap)
	return palindromes
}

func validateIfPalindromePossible(s string) (bool, int) {
	sRunes := []rune(s)
	sort.Slice(sRunes, func(i, j int) bool {
		return sRunes[i] > sRunes[j]
	})

	n := len(slices.Compact(sRunes))
	fmt.Println(n)
	if n < 3 {
		//all characters are same
		return false, n
	} else if n == len(sRunes) {
		// all characters are difference
		return false, 0
	}

	return true, 0
}

func countPalindromeSubsequence(s string) int {
	n := len(s)
	if n < 3 {
		return 0
	}

	first := make([]int, 26)
	last := make([]int, 26)

	for i := 0; i < 26; i++ {
		first[i] = -1
		last[i] = -1
	}

	// Record first & last occurrence
	for i := 0; i < n; i++ {
		idx := int(s[i] - 'a')
		if first[idx] == -1 {
			first[idx] = i
		}
		last[idx] = i
	}

	result := 0

	// For each character a..z
	for c := 0; c < 26; c++ {
		if first[c] == -1 || last[c] == -1 || last[c]-first[c] < 2 {
			continue
		}

		// Collect distinct middle characters
		middle := make(map[byte]bool)
		for i := first[c] + 1; i < last[c]; i++ {
			middle[s[i]] = true
		}

		result += len(middle)
	}

	return result
}
