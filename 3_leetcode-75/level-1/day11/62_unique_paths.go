package dynamicprogramming

/*
62. Unique Paths
Medium
13.4K
377
Companies

There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]). The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either down or right at any point in time.

Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.

The test cases are generated so that the answer will be less than or equal to 2 * 109.


Example 1:
Input: m = 3, n = 7
Output: 28

Example 2:
Input: m = 3, n = 2
Output: 3
Explanation: From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Down -> Down
2. Down -> Down -> Right
3. Down -> Right -> Down


Constraints:
1 <= m, n <= 100

*/

import (
	"strconv"
)

func UniquePaths(m int, n int) int {

	cache := make(map[string]int)
	return dfs(cache, m, n, 0, 0)

}

func dfs(cache map[string]int, m, n, row, col int) int {

	if row == m-1 && col == n-1 {
		return 1
	} else if row >= m || row < 0 || col >= n || col < 0 {
		return 0
	}

	if val, ok := cache[key(row, col)]; ok {
		return val
	}

	val1 := dfs(cache, m, n, row+1, col)
	val2 := dfs(cache, m, n, row, col+1)

	cache[key(row+1, col)] = val1
	cache[key(row, col+1)] = val2

	cache[key(row, col)] = val1 + val2
	return val1 + val2

}

func key(row, col int) string {
	return strconv.Itoa(row) + "|" + strconv.Itoa(col)
}
