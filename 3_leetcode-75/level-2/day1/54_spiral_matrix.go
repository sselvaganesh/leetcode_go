package simulation

/*
54. Spiral Matrix
Medium
11K
1K
Companies
Given an m x n matrix, return all elements of the matrix in spiral order.

Example 1:
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]

Example 2:
Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]

Constraints:
m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100
*/

func SpiralOrder(matrix [][]int) []int {

	rows := len(matrix)
	cols := len(matrix[0])

	trace := make([][]bool, rows)
	for i := range trace {
		trace[i] = make([]bool, cols)
	}

	row, col, dir := 0, 0, 0
	row, col, elems := getElems(matrix, trace, rows, cols, row, col, dir)
	var res []int

	for elems != nil {

		res = append(res, elems...)
		dir++
		dir = dir % 4

		if dir == 0 {
			col++
		} else if dir == 1 {
			row++
		} else if dir == 2 {
			col--
		} else {
			row--
		}

		row, col, elems = getElems(matrix, trace, rows, cols, row, col, dir)
	}

	return res
}

func getElems(mat [][]int, trace [][]bool, rows, cols, row, col, dir int) (int, int, []int) {

	var elems []int
	if dir == 0 { // Right
		for ; col < cols && !trace[row][col]; col++ {
			elems = append(elems, mat[row][col])
			trace[row][col] = true
		}
		col--
	} else if dir == 1 { // Down
		for ; row < rows && !trace[row][col]; row++ {
			elems = append(elems, mat[row][col])
			trace[row][col] = true
		}
		row--
	} else if dir == 2 { // Left
		for ; col >= 0 && !trace[row][col]; col-- {
			elems = append(elems, mat[row][col])
			trace[row][col] = true
		}
		col++
	} else { // Up
		for ; row >= 0 && !trace[row][col]; row-- {
			elems = append(elems, mat[row][col])
			trace[row][col] = true
		}
		row++
	}
	return row, col, elems
}

func SpiralOrderSolution(matrix [][]int) []int {

	var res []int

	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[0]))
	}
	row, col, newRow, newCol, direction := 0, 0, 0, 0, 0

	for {

		tot := len(res)
		newRow, newCol = traverse(matrix, visited, &res, row, col, direction)
		if tot == len(res) {
			break
		}

		direction++
		direction %= 4
		row, col = newRow, newCol
	}

	return res
}

func traverse(matrix [][]int, visited [][]bool, res *[]int, startRow, startCol, direction int) (int, int) {

	const (
		right, down, left, up = 0, 1, 2, 3
	)
	rows, cols := len(matrix), len(matrix[0])

	row, col := startRow, startCol
	if direction == right {
		for col < cols && !visited[row][col] {
			*res = append(*res, matrix[row][col])
			visited[row][col] = true
			col++
		}
		col--
		row++
	} else if direction == down {
		for row < rows && !visited[row][col] {
			*res = append(*res, matrix[row][col])
			visited[row][col] = true
			row++
		}
		row--
		col--
	} else if direction == left {
		for col >= 0 && !visited[row][col] {
			*res = append(*res, matrix[row][col])
			visited[row][col] = true
			col--
		}
		col++
		row--
	} else {
		for row >= 0 && !visited[row][col] {
			*res = append(*res, matrix[row][col])
			visited[row][col] = true
			row--
		}
		row++
		col++
	}

	return row, col
}
