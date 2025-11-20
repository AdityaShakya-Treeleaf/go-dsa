package dsa

import "log"

//Set Intersection Size At Least Two
//You are given a 2D integer array intervals where intervals[i] = [starti, endi] represents all the integers from starti to endi inclusively.
//A containing set is an array nums where each interval from intervals has at least two integers in nums.

//For example, if intervals = [[1,3], [3,7], [8,9]], then [1,2,4,7,8,9] and [2,3,4,8,9] are containing sets.
//Return the minimum possible size of a containing set.

type SetIntersectionSizeTwo struct{}

func (s SetIntersectionSizeTwo) IntersectionSizeTwo(intervals [][]int) int {
	log.Printf("Length of intervals: %d", len(intervals))
	counterMap := make(map[int][]int)
	maxContained := 1
	for ind, interval := range intervals {
		n := len(interval)
		if n > 0 {
			i := interval[0]
			for i <= interval[n-1] {
				counterMap[i] = append(counterMap[i], ind)
				maxContained = maximum(maxContained, len(counterMap[i]))
				i++
			}
		}
	}

	log.Println("CounterMap: ", counterMap)
	log.Println("Max: ", maxContained)

	if maxContained == 1 {
		return len(intervals) * 2
	}

	return (len(intervals) * 2) - (len(intervals) - maxContained)
}

func maximum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
