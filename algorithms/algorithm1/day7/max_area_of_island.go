package searchalgo

import (
	"fmt"
)

func print(inp [][]int) {

	for i := 0; i < len(inp); i++ {
		fmt.Println(inp[i])
	}

}

func MaxAreaOfIsland(grid [][]int) int {

	if grid == nil {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])

	result := make([][]int, rows)
	for i := range result {
		result[i] = make([]int, cols)
	}
	copy(result, grid)

	max := 0
	for row := 0; row < rows; row++ {

		for col := 0; col < cols; col++ {

			if result[row][col] == 1 {

				tot := sink(result, rows, cols, row, col)
				if tot > max {
					max = tot
				}

			}

		}

	}

	return max
}

func sink(result [][]int, rows, cols, row, col int) int {

	if row < 0 || row >= rows || col < 0 || col >= cols || result[row][col] == 0 {
		return 0
	}

	total := 1

	result[row][col] = 0

	total += sink(result, rows, cols, row+1, col)
	total += sink(result, rows, cols, row-1, col)
	total += sink(result, rows, cols, row, col+1)
	total += sink(result, rows, cols, row, col-1)

	return total

}
