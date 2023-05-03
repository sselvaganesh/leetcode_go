package bfs_dfs

/*
1091. Shortest Path in Binary Matrix
Medium
4.6K
182
Companies

Given an n x n binary matrix grid, return the length of the shortest clear path in the matrix. If there is no clear path, return -1.

A clear path in a binary matrix is a path from the top-left cell (i.e., (0, 0)) to the bottom-right cell (i.e., (n - 1, n - 1)) such that:

All the visited cells of the path are 0.
All the adjacent cells of the path are 8-directionally connected (i.e., they are different and they share an edge or a corner).
The length of a clear path is the number of visited cells of this path.

Example 1:
Input: grid = [[0,1],[1,0]]
Output: 2

Example 2:
Input: grid = [[0,0,0],[1,1,0],[1,1,0]]
Output: 4

Example 3:
Input: grid = [[1,0,0],[1,1,0],[1,1,0]]
Output: -1

Constraints:
n == grid.length
n == grid[i].length
1 <= n <= 100
grid[i][j] is 0 or 1

*/

type co struct {
	row, col int
}

func shortestPathBinaryMatrix(grid [][]int) int {

	if grid[0][0] == 1 {
		return -1
	}
	path, n := 1, len(grid)
	indices := []co{{0, 0}}

	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	for indices != nil {

		if isReached(indices, n) {
			return path
		}
		path++
		indices = nextMoves(indices, grid, visited, n)
	}

	return -1

}

func nextMoves(inp []co, grid [][]int, visited [][]bool, n int) []co {

	var res []co

	for _, pos := range inp {
		row, col := pos.row, pos.col
		if col-1 >= 0 && !visited[row][col-1] && grid[row][col-1] == 0 {
			visited[row][col-1] = true
			res = append(res, co{row, col - 1})
		}
		if col+1 < n && !visited[row][col+1] && grid[row][col+1] == 0 {
			visited[row][col+1] = true
			res = append(res, co{row, col + 1})
		}
		if row-1 >= 0 && !visited[row-1][col] && grid[row-1][col] == 0 {
			visited[row-1][col] = true
			res = append(res, co{row - 1, col})
		}
		if row+1 < n && !visited[row+1][col] && grid[row+1][col] == 0 {
			visited[row+1][col] = true
			res = append(res, co{row + 1, col})
		}

		//Diagnols
		if row-1 >= 0 && col-1 >= 0 && !visited[row-1][col-1] && grid[row-1][col-1] == 0 {
			visited[row-1][col-1] = true
			res = append(res, co{row - 1, col - 1})
		}
		if row-1 >= 0 && col+1 < n && !visited[row-1][col+1] && grid[row-1][col+1] == 0 {
			visited[row-1][col+1] = true
			res = append(res, co{row - 1, col + 1})
		}
		if row+1 < n && col-1 >= 0 && !visited[row+1][col-1] && grid[row+1][col-1] == 0 {
			visited[row+1][col-1] = true
			res = append(res, co{row + 1, col - 1})
		}
		if row+1 < n && col+1 < n && !visited[row+1][col+1] && grid[row+1][col+1] == 0 {
			visited[row+1][col+1] = true
			res = append(res, co{row + 1, col + 1})
		}
	}

	return res
}

func isReached(inp []co, n int) bool {

	for _, pos := range inp {
		if pos.row == n-1 && pos.col == n-1 {
			return true
		}
	}

	return false
}
