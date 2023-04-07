package dp

/*
70. Climbing Stairs
Easy
17.7K
548
Companies

You are climbing a staircase. It takes n steps to reach the top.
Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example 1:
Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps


Example 2:
Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step


Constraints:
1 <= n <= 45
*/

func ClimbStairs(n int) int {

	//m:=make(map[int]int)
	//return recursion(n, m)

	return dfs1(n)
}

func dfs1(n int) int {

	steps := []int{1, 2}
	m := make(map[int]int)
	return dfs_helper(n, steps, m)
}

func dfs_helper(rem int, steps []int, m map[int]int) int {

	if rem == 0 {
		return 1
	} else if rem < 0 {
		return 0
	}

	if val, ok := m[rem]; ok {
		return val
	}

	result := 0
	for i := 0; i < len(steps); i++ {
		temp := dfs_helper(rem-steps[i], steps, m)
		m[rem-steps[i]] = temp
		result += temp
	}

	m[rem] = result
	return result
}

func recursion(n int, m map[int]int) int {

	if n == 0 {
		return 1
	} else if n < 0 {
		return 0
	}

	if val, ok := m[n]; ok {
		return val
	}

	val := recursion(n-1, m) + recursion(n-2, m)
	m[n] = val

	return val

}
