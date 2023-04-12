package array

/*
56. Merge Intervals
Medium
18.6K
631
Companies

Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.


Example 1:
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].

Example 2:
Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.


Constraints:
1 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 104

*/

import (
	"sort"
)

func Merge(intervals [][]int) [][]int {

	if intervals == nil {
		return nil
	}

	if len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i+1 < len(intervals); {

		if isOverlap(intervals[i], intervals[i+1]) {
			merged := mergeTwoIntervals(intervals[i], intervals[i+1])
			intervals = append(append(append([][]int{}, intervals[:i]...), merged), intervals[i+2:]...)
		} else {
			i++
		}

	}

	return intervals

}

func mergeTwoIntervals(val1, val2 []int) []int {

	min := func(v1, v2 int) int {
		if v1 > v2 {
			return v2
		}
		return v1
	}

	max := func(v1, v2 int) int {
		if v1 > v2 {
			return v1
		}
		return v2
	}

	return []int{min(val1[0], val2[0]), max(val1[1], val2[1])}

}

func isOverlap(val1, val2 []int) bool {
	return val2[0] >= val1[0] && val2[0] <= val1[1]
}
