package dp

/*
120. Triangle
Medium
7.9K
473
Companies

Given a triangle array, return the minimum path sum from top to bottom.

For each step, you may move to an adjacent number of the row below. More formally, if you are on index i on the current row, you may move to either index i or index i + 1 on the next row.


Example 1:
Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
Output: 11
Explanation: The triangle looks like:
   2
  3 4
 6 5 7
4 1 8 3
The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined above).


Example 2:
Input: triangle = [[-10]]
Output: -10


Constraints:
1 <= triangle.length <= 200
triangle[0].length == 1
triangle[i].length == triangle[i - 1].length + 1
-104 <= triangle[i][j] <= 104

*/

import(
	"strconv"
	"math"
)

func minimumTotal(triangle [][]int) int {

	//cache:=make(map[string]int)
	//return dfs_minTotal(triangle, 0, 0, cache)
	return dp(triangle)

}

func dp(mat [][]int) int {

	if mat == nil {
		return 0
	}

	result := append([][]int{}, mat...) //Duplicating the input - Deep copy
	for row := len(mat) - 2; row >= 0; row-- {

		for col := 0; col < len(mat[row]); col++ {

			result[row][col] += min(result[row+1][col], result[row+1][col+1])

		}
	}

	return result[0][0]

}

func dfs_minTotal(mat [][]int, row, col int, cache map[string]int) int {

	if row >= len(mat) || col >= len(mat[row]) {
		return 0
	}

	if val, ok := cache[key(row, col)]; ok {
		return val
	}

	val := mat[row][col] + min(dfs_minTotal(mat, row+1, col, cache), dfs_minTotal(mat, row+1, col+1, cache))
	cache[key(row, col)] = val
	return val

}

func key(row, col int) string {
	return strconv.Itoa(row) + "|" + strconv.Itoa(col)
}

func min(input ...int) int {

	if len(input) == 0 {
		return 0
	}

	minVal := input[0]
	for _, val := range input {
		if minVal > val {
			minVal = val
		}
	}

	return minVal
}

func minMatrix(mat [][]int) (int, int) {

	row, col := -1, -1
	curMin := math.MinInt

	for i := 0; i < len(mat); i++ {

		for j := 0; j < len(mat[i]); j++ {

			if curMin >= mat[i][j] {
				curMin = mat[i][j]
				row, col = i, j
			}

		}

	}

	return row, col
}
