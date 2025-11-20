package dsa

import (
	"fmt"
	"sort"
)

//Set Intersection Size At Least Two
//You are given a 2D integer array intervals where intervals[i] = [starti, endi] represents all the integers from starti to endi inclusively.
//A containing set is an array nums where each interval from intervals has at least two integers in nums.

//For example, if intervals = [[1,3], [3,7], [8,9]], then [1,2,4,7,8,9] and [2,3,4,8,9] are containing sets.
//Return the minimum possible size of a containing set.

type SetIntersectionSizeTwo struct{}

func (s SetIntersectionSizeTwo) IntersectionSizeTwo(intervals [][]int) int {
	// Sort intervals by end point, and by start point in descending order for same end
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] > intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})

	fmt.Println(intervals)
	minLength := 0
	// We'll track the two largest numbers we've chosen for the current coverage
	secondHighest, highest := -1, -1

	for _, interval := range intervals {
		start, end := interval[0], interval[1]

		// If current interval is already covered by our two numbers
		if start <= secondHighest {
			continue
		}

		// If we need one more number for this interval
		if start > highest {
			// Add two largest numbers from this interval
			minLength += 2
			highest = end
			secondHighest = end - 1
		} else {
			// We need exactly one more number
			minLength += 1
			secondHighest = highest
			highest = end
			// Make sure first is at least start
			if secondHighest < start {
				secondHighest = start
			}
		}
	}

	return minLength
}
