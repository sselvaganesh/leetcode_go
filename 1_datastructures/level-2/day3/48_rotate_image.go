package array

/*
48. Rotate Image
Medium
14.2K
640
Companies

You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

Example 1:
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]

Example 2:
Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

Constraints:
n == matrix.length == matrix[i].length
1 <= n <= 20
-1000 <= matrix[i][j] <= 1000
*/

func Rotate(matrix [][]int) {

	left, right := 0, len(matrix)-1

	for right > left {

		for i := 0; i < (right - left); i++ {

			top, bottom := left, right

			// Save top left
			topLeft := matrix[top][left+i]

			// Move bottom left to top left
			matrix[top][left+i] = matrix[bottom-i][left]

			// Move bottom right to bottom left
			matrix[bottom-i][left] = matrix[bottom][right-i]

			// Move top right to bottom right
			matrix[bottom][right-i] = matrix[top+i][right]

			// Move top left to top right
			matrix[top+i][right] = topLeft
		}

		left++
		right--
	}

	return

}

func Rotate2(matrix [][]int) {
	transpose(matrix)
	reverseRows(matrix)
}

func reverseRows(matrix [][]int) {
	for i := range matrix {
		for j := 0; j < len(matrix)/2; j++ {
			matrix[i][j], matrix[i][len(matrix)-1-j] = matrix[i][len(matrix)-1-j], matrix[i][j]
		}
	}
}

func transpose(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

}
