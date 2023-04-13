package array

/*
119. Pascal's Triangle II
Easy
3.7K
297
Companies
Given an integer rowIndex, return the rowIndexth (0-indexed) row of the Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:

Example 1:
Input: rowIndex = 3
Output: [1,3,3,1]

Example 2:
Input: rowIndex = 0
Output: [1]

Example 3:
Input: rowIndex = 1
Output: [1,1]

Constraints:
0 <= rowIndex <= 33

Follow up: Could you optimize your algorithm to use only O(rowIndex) extra space?
*/

func GetRow(rowIndex int) []int {

	//return bruteForce(rowIndex)
	return recursion(rowIndex)

}

func recursion(rowIndex int) []int {

	if rowIndex == 0 {
		return []int{1}
	} else if rowIndex == 1 {
		return []int{1, 1}
	}

	prevRow := recursion(rowIndex - 1)

	curRow := make([]int, rowIndex+1)
	curRow[0], curRow[rowIndex] = 1, 1
	for i := 1; i < len(curRow)-1; i++ {
		curRow[i] = prevRow[i-1] + prevRow[i]
	}

	return curRow
}

func bruteForce(rowIndex int) []int {

	if rowIndex == 0 {
		return getInitRowN(0)
	}

	prev := []int{1}
	var curRow []int
	for row := 1; row <= rowIndex; row++ {
		curRow = getInitRowN(row)
		updateRow(prev, curRow)
		prev = append([]int{}, curRow...)
	}

	return curRow
}

func updateRow(prev, cur []int) {
	for i := 1; i < len(cur)-1; i++ {
		cur[i] = prev[i-1] + prev[i]
	}
}

func getInitRowN(row int) []int {

	if row == 0 {
		return []int{1}
	}

	res := make([]int, row+1)
	res[0], res[row] = 1, 1
	return res
}
