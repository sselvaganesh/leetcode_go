package simulation

/*
202. Happy Number
Easy
8.5K
1.1K
Companies
Write an algorithm to determine if a number n is happy.

A happy number is a number defined by the following process:

Starting with any positive integer, replace the number by the sum of the squares of its digits.
Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
Those numbers for which this process ends in 1 are happy.
Return true if n is a happy number, and false if not.

Example 1:
Input: n = 19
Output: true
Explanation:
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1

Example 2:
Input: n = 2
Output: false

Constraints:
1 <= n <= 231 - 1
*/

func IsHappy(n int) bool {

	num := n
	m := make(map[int]struct{})

	for num != 1 {
		if _, ok := m[num]; ok {
			return false
		}
		m[num] = struct{}{}
		num = sum(squareOfdigits(num)...)
	}

	return true

}

func HappyNumberSolution(n int) bool {

	m := make(map[int]struct{})
	for {
		val := squareOfDigits1(n)
		if val == 1 {
			return true
		}

		if _, ok := m[val]; ok {
			break
		}
		m[val] = struct{}{}
		n = val
	}

	return false
}

func squareOfDigits1(n int) int {
	res := 0
	for n > 0 {
		last := n % 10
		res += (last * last)
		n /= 10
	}
	return res
}

func squareOfdigits(num int) []int {
	var res []int
	for num > 0 {
		t := num % 10
		res = append(res, t*t)
		num = num / 10
	}

	return res
}

func sum(nums ...int) int {
	res := 0
	for _, val := range nums {
		res += val
	}
	return res
}
