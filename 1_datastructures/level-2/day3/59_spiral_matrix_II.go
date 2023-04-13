package array

/*
59. Spiral Matrix II
Medium
4.8K
209
Companies

Given a positive integer n, generate an n x n matrix filled with elements from 1 to n2 in spiral order.

Example 1:
Input: n = 3
Output: [[1,2,3],[8,9,4],[7,6,5]]

Example 2:
Input: n = 1
Output: [[1]]


Constraints:
1 <= n <= 20
*/

func GenerateMatrix(n int) [][]int {
	return solution_generateSpiralMatrix(n)
}

func solution_generateSpiralMatrix(n int) [][]int {

	res := make([][]int, n)
	trace := make([][]bool, n)
	for i := range res {
		res[i] = make([]int, n)
		trace[i] = make([]bool, n)
	}

	num, dir, row, col := 1, 0, 0, 0
	for {
		dir = dir % 4
		oldNum := num
		row, col, num = fill(res, trace, dir, n, n, row, col, num)
		if num == oldNum {
			break
		}
		dir++
	}

	return res
}

func fill(mat [][]int, trace [][]bool, dir, rows, cols, startRow, startCol, num int) (int, int, int) {

	row, col := startRow, startCol

	// 0 -> Right : col+1
	if dir == 0 {
		for col < cols && !trace[row][col] {
			mat[row][col] = num
			trace[row][col] = true
			num++
			col++
		}
		col--
		row++
	}

	// 1 -> Down : row+1
	if dir == 1 {
		for row < rows && !trace[row][col] {
			mat[row][col] = num
			trace[row][col] = true
			num++
			row++
		}
		row--
		col--
	}

	// 2 -> Left : col-1
	if dir == 2 {
		for col >= 0 && !trace[row][col] {
			mat[row][col] = num
			trace[row][col] = true
			num++
			col--
		}
		col++
		row--
	}

	// 3 -> Up : row-1
	if dir == 3 {
		for row < rows && !trace[row][col] {
			mat[row][col] = num
			trace[row][col] = true
			num++
			row--
		}
		row++
		col++
	}

	return row, col, num
}
