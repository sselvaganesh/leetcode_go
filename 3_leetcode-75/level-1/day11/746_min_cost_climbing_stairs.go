package dynamicprogramming

/*
746. Min Cost Climbing Stairs
Easy
9.1K
1.4K
Companies

You are given an integer array cost where cost[i] is the cost of ith step on a staircase. Once you pay the cost, you can either climb one or two steps.

You can either start from the step with index 0, or the step with index 1.

Return the minimum cost to reach the top of the floor.


Example 1:
Input: cost = [10,15,20]
Output: 15
Explanation: You will start at index 1.
- Pay 15 and climb two steps to reach the top.
The total cost is 15.

Example 2:
Input: cost = [1,100,1,1,1,100,1,1,100,1]
Output: 6
Explanation: You will start at index 0.
- Pay 1 and climb two steps to reach index 2.
- Pay 1 and climb two steps to reach index 4.
- Pay 1 and climb two steps to reach index 6.
- Pay 1 and climb one step to reach index 7.
- Pay 1 and climb two steps to reach index 9.
- Pay 1 and climb one step to reach the top.
The total cost is 6.


Constraints:
2 <= cost.length <= 1000
0 <= cost[i] <= 999

*/

func MinCostClimbingStairs(cost []int) int {

	m := make(map[int]int)
	return minCost(cost, m, 0)

}

func minCost(cost []int, m map[int]int, idx int) int {

	if (idx + 1) >= len(cost) {
		return 0
	}

	if val, ok := m[idx]; ok {
		return val
	}

	res := min(cost[idx]+minCost(cost, m, idx+1), cost[idx+1]+minCost(cost, m, idx+2))
	m[idx] = res
	return res

}

func min(val1, val2 int) int {
	if val1 < val2 {
		return val1
	}
	return val2
}
