package strings

/*
43. Multiply Strings
Medium
6.1K
2.7K
Companies

Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.

Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.

Example 1:
Input: num1 = "2", num2 = "3"
Output: "6"

Example 2:
Input: num1 = "123", num2 = "456"
Output: "56088"


Constraints:
1 <= num1.length, num2.length <= 200
num1 and num2 consist of digits only.
Both num1 and num2 do not contain any leading zero, except the number 0 itself.
*/

import (
	"strconv"
	"strings"
)

func Multiply(num1 string, num2 string) string {

	res := make([]int, len(num1)+len(num2))

	for l1 := len(num1) - 1; l1 >= 0; l1-- {
		for l2 := len(num2) - 1; l2 >= 0; l2-- {
			idx := l1 + l2 + 1
			mul := byteToInt(num1[l1]) * byteToInt(num2[l2])
			addToRes(res, idx, mul)
			idx--
		}
	}

	result := ""
	for i := 0; i < len(res); i++ {
		result += strconv.Itoa(res[i])
	}

	mult := strings.TrimLeft(result, "0")
	if mult == "" {
		return "0"
	}

	return mult

}

func addToRes(res []int, start, mul int) {
	rem := 0
	for mul > 0 {
		sum := res[start] + (mul % 10) + rem
		res[start] = sum % 10
		rem = sum / 10
		mul /= 10
		start--
	}

	for rem > 0 {
		sum := res[start] + rem
		res[start] = sum % 10
		sum = sum / 10
		rem = sum
		start--
	}

}

func byteToInt(inp byte) int {
	return int(inp - '0')
}

func multiply_another_person_solution(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	n, m := len(num1), len(num2)
	res := make([]byte, n+m)
	for i := n - 1; i >= 0; i-- { // 12
		for j := m - 1; j >= 0; j-- { // 14
			curr := (num1[i] - '0') * (num2[j] - '0')
			res[i+j+1] += curr
			if res[i+j+1] > 9 {
				res[i+j] += res[i+j+1] / 10
				res[i+j+1] %= 10
			}
		}
	}

	if res[0] == 0 {
		res = res[1:]
	}

	for i := range res {
		res[i] += '0'
	}

	return string(res)
}
