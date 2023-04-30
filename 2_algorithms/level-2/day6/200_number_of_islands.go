package bfs_dfs

/*
200. Number of Islands
Medium
19.6K
435
Companies

Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.


Example 1:
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1

Example 2:
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3

Constraints:
m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] is '0' or '1'.
*/

func NumIslands(grid [][]byte) int {

	if grid == nil {
		return 0
	}

	rows, cols, islands := len(grid), len(grid[0]), 0
	for row := 0; row < rows; row++ {

		for col := 0; col < cols; col++ {

			if grid[row][col] == '1' {
				islands++
				sink(grid, rows, cols, row, col)
			}

		}

	}

	return islands
}

func sink(grid [][]byte, rows, cols, row, col int) {

	if row >= rows || row < 0 || col >= cols || col < 0 || grid[row][col] == '0' {
		return
	}
	grid[row][col] = '0'
	sink(grid, rows, cols, row, col+1)
	sink(grid, rows, cols, row+1, col)
	sink(grid, rows, cols, row, col-1)
	sink(grid, rows, cols, row-1, col)
}
