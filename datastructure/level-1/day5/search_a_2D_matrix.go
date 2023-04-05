package array

/*
You are given an m x n integer matrix matrix with the following two properties:

Each row is sorted in non-decreasing order.
The first integer of each row is greater than the last integer of the previous row.
Given an integer target, return true if target is in matrix or false otherwise.

You must write a solution in O(log(m * n)) time complexity.

Example 1:
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true

Example 2:
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false

*/

func SearchMatrix(matrix [][]int, target int) bool {

	if matrix == nil {
		return false
	}

	totRows := len(matrix)
	totCols := len(matrix[0])

	low, high := 0, totRows*totCols
	var row, col int

	for high >= low {

		mid := (low + high) / 2
		if mid == 0 {
			row, col = 0, 0
		} else {
			row, col = index(mid, totRows, totCols)
		}

		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}

	return false

}

func index(pos int, totRows, totCols int) (int, int) {

	row := (pos - 1) / totCols
	col := (pos - 1) - (row * totCols)
	return row, col

}
