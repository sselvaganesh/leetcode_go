package bfs_dfs

/*
797. All Paths From Source to Target
Medium
6.5K
134
Companies

Given a directed acyclic graph (DAG) of n nodes labeled from 0 to n - 1, find all possible paths from node 0 to node n - 1 and return them in any order.

The graph is given as follows: graph[i] is a list of all nodes you can visit from node i (i.e., there is a directed edge from node i to node graph[i][j]).

Example 1:
Input: graph = [[1,2],[3],[3],[]]
Output: [[0,1,3],[0,2,3]]
Explanation: There are two paths: 0 -> 1 -> 3 and 0 -> 2 -> 3.

Example 2:
Input: graph = [[4,3,1],[3,2,4],[3],[4],[]]
Output: [[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]


Constraints:
n == graph.length
2 <= n <= 15
0 <= graph[i][j] < n
graph[i][j] != i (i.e., there will be no self-loops).
All the elements of graph[i] are unique.
The input graph is guaranteed to be a DAG.
*/

func AllPathsSourceTarget(graph [][]int) [][]int {

	var result [][]int
	toDestination(graph, 0, len(graph)-1, &result, []int{})
	return result
}

func toDestination(graph [][]int, src, dest int, res *[][]int, path []int) {

	if src == dest {
		*res = append(*res, append([]int{0}, path...))
		return
	}

	nextHop := graph[src]
	if len(nextHop) == 0 {
		return
	}

	for i := 0; i < len(nextHop); i++ {
		path = append(path, nextHop[i])
		toDestination(graph, nextHop[i], dest, res, path)
		path = path[:len(path)-1]
	}

}
