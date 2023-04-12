package binary_search

/*
74. Search a 2D Matrix
Medium
11.9K
343
Companies

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

Constraints:
m == matrix.length
n == matrix[i].length
1 <= m, n <= 100
-104 <= matrix[i][j], target <= 104

*/

func SearchMatrix(matrix [][]int, target int) bool {

	return searchMatrixSolution(matrix, target)

}

func searchMatrixSolution(mat [][]int, target int) bool {

	if mat == nil {
		return false
	}

	rows := len(mat)
	cols := len(mat[0])
	low, high := 0, (rows*cols)-1

	for high >= low {

		mid := (low + high) / 2
		row, col := index(cols, mid)

		if mat[row][col] == target {
			return true
		} else if target > mat[row][col] {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}

	return false

}

func index(cols, pos int) (int, int) {
	return pos / cols, pos % cols
}