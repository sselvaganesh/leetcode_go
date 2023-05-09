package searchalgo

/*
You are given an m x n grid where each cell can have one of three values:

0 representing an empty cell,
1 representing a fresh orange, or
2 representing a rotten orange.
Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.

Example 1:
Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
Output: 4

Example 2:
Input: grid = [[2,1,1],[0,1,1],[1,0,1]]
Output: -1
Explanation: The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally.

Example 3:
Input: grid = [[0,2]]
Output: 0
Explanation: Since there are already no fresh oranges at minute 0, the answer is just 0.
*/

func OrangesRotting(grid [][]int) int {

	if grid == nil {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	minute := 0
	nextMin, good, oldRotten, newRotten := minutePlusOne(grid, rows, cols)

	for {

		if good == 0 {
			return minute
		} else if newRotten == 0 || oldRotten == 0 {
			return -1
		}

		minute++
		nextMin, good, oldRotten, newRotten = minutePlusOne(nextMin, rows, cols)
	}

	return minute
}

func minutePlusOne(minute [][]int, rows, cols int) ([][]int, int, int, int) {

	good, oldRotten, newRotten := 0, 0, 0
	nextMinute := make([][]int, rows)
	for i := range nextMinute {
		nextMinute[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			nextMinute[i][j] = minute[i][j]
		}
	}

	for row := 0; row < rows; row++ {

		for col := 0; col < cols; col++ {

			if minute[row][col] == 1 {
				good++
				continue
			}

			isBad := false
			if minute[row][col] == 2 {
				isBad = true
				oldRotten++
			}

			//upwards
			if isBad && row-1 >= 0 && minute[row-1][col] == 1 {
				nextMinute[row-1][col] = 2
				newRotten++
			}

			//Downwards
			if isBad && row+1 < rows && minute[row+1][col] == 1 {
				nextMinute[row+1][col] = 2
				newRotten++
			}

			//Left
			if isBad && col-1 >= 0 && minute[row][col-1] == 1 {
				nextMinute[row][col-1] = 2
				newRotten++
			}

			//right
			if isBad && col+1 < cols && minute[row][col+1] == 1 {
				nextMinute[row][col+1] = 2
				newRotten++
			}

		}

	}

	return nextMinute, good, oldRotten, newRotten
}
