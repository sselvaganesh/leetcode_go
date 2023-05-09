package searchalgo

/*
542. 01 Matrix
Medium
6.9K
326
Companies


Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell.

The distance between two adjacent cells is 1.



Example 1:


Input: mat = [[0,0,0],[0,1,0],[0,0,0]]
Output: [[0,0,0],[0,1,0],[0,0,0]]
Example 2:


Input: mat = [[0,0,0],[0,1,0],[1,1,1]]
Output: [[0,0,0],[0,1,0],[1,2,1]]


Constraints:

m == mat.length
n == mat[i].length
1 <= m, n <= 104
1 <= m * n <= 104
mat[i][j] is either 0 or 1.
There is at least one 0 in mat.


*/

import (
// "fmt"
)

func UpdateMatrix(mat [][]int) [][]int {

	//return bfs(mat)

	return dp(mat)

}

func dp(mat [][]int) [][]int {

	if mat == nil {
		return nil
	}

	rows := len(mat)
	cols := len(mat[0])

	// Prepare answers
	dist := make([][]int, rows)
	for i := range dist {
		dist[i] = make([]int, cols)
	}
	max := rows * cols
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			dist[row][col] = max
		}
	}

	// Fill while comparing Left and Up
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if mat[row][col] == 0 {
				dist[row][col] = 0
				continue
			}

			up, left := max, max
			if col-1 >= 0 {
				left = dist[row][col-1]
			}
			if row-1 >= 0 {
				up = dist[row-1][col]
			}
			dist[row][col] = 1 + min(up, left)
		}
	}

	// Fill while comparing Right and Down
	for row := rows - 1; row >= 0; row-- {
		for col := cols - 1; col >= 0; col-- {
			if mat[row][col] == 0 {
				dist[row][col] = 0
				continue
			}

			right, down := max, max
			if col+1 < cols {
				right = dist[row][col+1]
			}
			if row+1 < rows {
				down = dist[row+1][col]
			}
			dist[row][col] = min(dist[row][col], 1+min(right, down))

		}
	}

	return dist

}

func min(val1, val2 int) int {
	if val1 < val2 {
		return val1
	}
	return val2
}

type co struct{ row, col int }

func bfs(mat [][]int) [][]int {

	if mat == nil {
		return nil
	}

	rows, cols := len(mat), len(mat[0])
	result := make([][]int, rows)

	for i := range result {
		result[i] = make([]int, cols)
	}

	for row := 0; row < rows; row++ {

		for col := 0; col < cols; col++ {

			if mat[row][col] == 1 {
				result[row][col] = nearest(mat, rows, cols, row, col)
			}

		}

	}

	return result
}

func nearest(mat [][]int, rows, cols, row, col int) int {

	if mat[row][col] == 0 {
		return 0
	}

	paths := make([][]bool, rows)
	for i := range paths {
		paths[i] = make([]bool, cols)
	}

	getNextSteps := func(paths [][]bool, rows, cols int, points []co) []co {

		var result []co
		for _, p := range points {

			if p.row-1 >= 0 && !paths[p.row-1][p.col] {
				result = append(result, co{p.row - 1, p.col})
				paths[p.row-1][p.col] = true
			}

			if p.row+1 < rows && !paths[p.row+1][p.col] {
				result = append(result, co{p.row + 1, p.col})
				paths[p.row+1][p.col] = true
			}

			if p.col-1 >= 0 && !paths[p.row][p.col-1] {
				result = append(result, co{p.row, p.col - 1})
				paths[p.row][p.col-1] = true
			}

			if p.col+1 < cols && !paths[p.row][p.col+1] {
				result = append(result, co{p.row, p.col + 1})
				paths[p.row][p.col+1] = true
			}

		}

		return result
	}

	steps := 1
	nextSteps := []co{{row, col}}
	paths[row][col] = true

	for {

		nextSteps = getNextSteps(paths, rows, cols, nextSteps)
		if nextSteps == nil {
			break
		}

		for _, p := range nextSteps {
			if mat[p.row][p.col] == 0 {
				return steps
			}
		}

		steps++

	}

	return steps
}

func getSource(mat [][]int, rows, cols int) []co {

	var result []co

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {

			if mat[row][col] == 0 {
				result = append(result, co{row, col})
			}

		}
	}

	return result
}
