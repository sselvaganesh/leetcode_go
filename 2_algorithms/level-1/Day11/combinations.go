package backtracking

/*
77. Combinations
Medium
5.9K
182
Companies

Given two integers n and k, return all possible combinations of k numbers chosen from the range [1, n].
You may return the answer in any order.


Example 1:
Input: n = 4, k = 2
Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
Explanation: There are 4 choose 2 = 6 total combinations.
Note that combinations are unordered, i.e., [1,2] and [2,1] are considered to be the same combination.

Example 2:
Input: n = 1, k = 1
Output: [[1]]
Explanation: There is 1 choose 1 = 1 total combination.

Constraints:

1 <= n <= 20
1 <= k <= n
*/

func Combine(n int, k int) [][]int {

	if n == 0 || k == 0 {
		return nil
	}

	inp := make([]int, n)
	for i := 0; i < n; i++ {
		inp[i] = i + 1
	}

	var result [][]int
	var elem []int
	solution1(inp, &result, elem, 0, 0, n, k)
	return result

}

func solution1(input []int, result *[][]int, elem []int, inpIdx, start, n, k int) {

	if len(elem) == k {
		r := append([]int{}, elem...)
		*result = append(*result, r)
		return
	} else if inpIdx >= n {
		return
	}

	for i := inpIdx; i < n; i++ {
		elem = append(elem, input[i])
		solution1(input, result, elem, i+1, start+1, n, k)
		elem = elem[:start]
	}

}
