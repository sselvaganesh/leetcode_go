package array

/*
240. Search a 2D Matrix II
Medium
10.3K
169
Companies

Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix. This matrix has the following properties:

Integers in each row are sorted in ascending from left to right.
Integers in each column are sorted in ascending from top to bottom.

Example 1:
Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
Output: true

Example 2:
Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
Output: false


Constraints:
m == matrix.length
n == matrix[i].length
1 <= n, m <= 300
-109 <= matrix[i][j] <= 109
All the integers in each row are sorted in ascending order.
All the integers in each column are sorted in ascending order.
-109 <= target <= 109

*/

func SearchMatrix(matrix [][]int, target int) bool {

	//return solution1(matrix, target)
	return solution2(matrix, target)

}

func solution2(matrix [][]int, target int) bool {

	row, col := 0, len(matrix[0])-1
	for row < len(matrix) && col >= 0 {

		if matrix[row][col] == target {
			return true
		} else if target > matrix[row][col] {
			row++
		} else {
			col--
		}

	}

	return false

}

func solution1(matrix [][]int, target int) bool {

	if matrix == nil {
		return false
	}

	for i := range matrix {
		if binarySearchRow(matrix[i], target) {
			return true
		}
	}
	return false
}

func binarySearchCol(mat [][]int, col, target int) bool {

	low, high := 0, len(mat)-1
	for high >= low {
		mid := (low + high) / 2
		if mat[mid][col] == target {
			return true
		} else if target > mat[mid][col] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false
}

func binarySearchRow(nums []int, target int) bool {

	low, high := 0, len(nums)-1
	for high >= low {
		mid := (low + high) / 2
		if nums[mid] == target {
			return true
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}
