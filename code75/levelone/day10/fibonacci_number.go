package dynamicprogramming

/*
509. Fibonacci Number
Easy
6.5K
311
Companies


The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,

F(0) = 0, F(1) = 1
F(n) = F(n - 1) + F(n - 2), for n > 1.
Given n, calculate F(n).

Example 1:
Input: n = 2
Output: 1
Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.

Example 2:
Input: n = 3
Output: 2
Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.

Example 3:
Input: n = 4
Output: 3
Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.

*/

func Fib(n int) int {

	//return fib_recursion(n)

	return dp(n)

}

func dp(n int) int {

	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	result := make([]int, n+1)
	result[0], result[1] = 0, 1

	for idx := 2; idx <= n; idx++ {
		result[idx] = result[idx-2] + result[idx-1]
	}

	return result[n]
}

func fib_recursion(n int) int {

	m := make(map[int]int)
	m[0] = 0
	m[1] = 1

	if val, ok := m[n]; ok {
		return val
	}

	result := fib_recursion(n-1) + fib_recursion(n-2)
	m[n] = result

	return result

}
