package two_pointers

/*
986. Interval List Intersections
Medium
5.1K
99
Companies

You are given two lists of closed intervals, firstList and secondList, where firstList[i] = [starti, endi] and secondList[j] = [startj, endj]. Each list of intervals is pairwise disjoint and in sorted order.

Return the intersection of these two interval lists.

A closed interval [a, b] (with a <= b) denotes the set of real numbers x with a <= x <= b.

The intersection of two closed intervals is a set of real numbers that are either empty or represented as a closed interval. For example, the intersection of [1, 3] and [2, 4] is [2, 3].


Example 1:
Input: firstList = [[0,2],[5,10],[13,23],[24,25]], secondList = [[1,5],[8,12],[15,24],[25,26]]
Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]

Example 2:
Input: firstList = [[1,3],[5,9]], secondList = []
Output: []


Constraints:
0 <= firstList.length, secondList.length <= 1000
firstList.length + secondList.length >= 1
0 <= starti < endi <= 109
endi < starti+1
0 <= startj < endj <= 109
endj < startj+1
*/

func IntervalIntersection(firstList [][]int, secondList [][]int) [][]int {

	var res [][]int

	for i, j := 0, 0; i < len(firstList) && j < len(secondList); {

		if isOverlap(firstList[i], secondList[j]) {

			inSection := intersection(firstList[i], secondList[j])
			if inSection[1] == firstList[i][1] {
				i++
			} else {
				j++
			}
			res = append(res, inSection)

		} else {

			if firstList[i][1] > secondList[j][1] {
				j++
			} else {
				i++
			}

		}

	}

	return res
}

func intersection(val1, val2 []int) []int {

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

	return []int{max(val1[0], val2[0]), min(val1[1], val2[1])}

}

func isOverlap(val1, val2 []int) bool {

	if val1[0] > val2[0] {
		val1, val2 = val2, val1
	}

	return (val2[0] >= val1[0] && val2[0] <= val1[1]) || (val2[1] >= val1[0] && val2[1] <= val1[1])

}
