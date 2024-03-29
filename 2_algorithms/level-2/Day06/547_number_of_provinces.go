package bfs_dfs

/*
547. Number of Provinces
Medium
7.5K
286
Companies

There are n cities. Some of them are connected, while some are not. If city a is connected directly with city b, and city b is connected directly with city c, then city a is connected indirectly with city c.

A province is a group of directly or indirectly connected cities and no other cities outside of the group.

You are given an n x n matrix isConnected where isConnected[i][j] = 1 if the ith city and the jth city are directly connected, and isConnected[i][j] = 0 otherwise.

Return the total number of provinces.

Example 1:
Input: isConnected = [[1,1,0],[1,1,0],[0,0,1]]
Output: 2

Example 2:
Input: isConnected = [[1,0,0],[0,1,0],[0,0,1]]
Output: 3


Constraints:
1 <= n <= 200
n == isConnected.length
n == isConnected[i].length
isConnected[i][j] is 1 or 0.
isConnected[i][i] == 1
isConnected[i][j] == isConnected[j][i]
*/

func findCircleNum(isConnected [][]int) int {

	provinces := 0

	for row := 0; row < len(isConnected); row++ {
		for col := 0; col < len(isConnected[row]); col++ {
			if isConnected[row][col] == 1 {
				provinces++
				connect(isConnected, col)
			}
		}
	}

	return provinces
}

func connect(cities [][]int, fromCity int) {

	for toCity := 0; toCity < len(cities[fromCity]); toCity++ {

		if cities[fromCity][toCity] == 1 {
			cities[fromCity][toCity] = 0
			connect(cities, toCity)
		}

	}

}
