package array

/*
435. Non-overlapping Intervals
Medium
5.7K
156
Companies
Given an array of intervals intervals where intervals[i] = [starti, endi], return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.

Example 1:
Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
Output: 1
Explanation: [1,3] can be removed and the rest of the intervals are non-overlapping.

Example 2:
Input: intervals = [[1,2],[1,2],[1,2]]
Output: 2
Explanation: You need to remove two [1,2] to make the rest of the intervals non-overlapping.

Example 3:
Input: intervals = [[1,2],[2,3]]
Output: 0
Explanation: You don't need to remove any of the intervals since they're already non-overlapping.

Constraints:
1 <= intervals.length <= 105
intervals[i].length == 2
-5 * 104 <= starti < endi <= 5 * 104
*/

import (
	"sort"
)

func EraseOverlapIntervals(intervals [][]int) int {

	if intervals == nil {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	removed := 0
	lastEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {

		curStart, curEnd := intervals[i][0], intervals[i][1]
		if curStart >= lastEnd {
			lastEnd = curEnd
		} else {
			removed++
			if lastEnd > curEnd {
				lastEnd = curEnd
			}
		}
	}

	return removed
}

func eraseOverlapIntervalsSolution(intervals [][]int) int {

	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	removed := 0
	start, end := intervals[0][0], intervals[0][1]
	for _, interval := range intervals[1:] {

		if interval[0] >= start && interval[0] < end {

			removed++

			if end >= interval[1] {
				start, end = interval[0], interval[1]
			}
		} else {
			start, end = interval[0], interval[1]
		}

	}

	return removed

}
