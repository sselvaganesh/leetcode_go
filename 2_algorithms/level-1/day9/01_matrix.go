package searchalgo

/*

Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell.
The distance between two adjacent cells is 1.

Example 1:
Input: mat = [[0,0,0],[0,1,0],[0,0,0]]
Output: [[0,0,0],[0,1,0],[0,0,0]]

Example 2:
Input: mat = [[0,0,0],[0,1,0],[1,1,1]]
Output: [[0,0,0],[0,1,0],[1,2,1]]

*/

import (
//"fmt"
)

func UpdateMatrix(mat [][]int) [][]int {

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

	type co struct{ row, col int }

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

func closest(mat [][]int, rows, cols int) (int, int) {

	for row := 0; row < rows; row++ {

		for col := 0; col < cols; col++ {

			if mat[row][col] == 0 {
				return row, col
			}
		}

	}

	return -1, -1
}
