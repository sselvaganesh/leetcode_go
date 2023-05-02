package string

/*
43. Multiply Strings
Medium
6.1K
2.8K
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

func Multiply(num1 string, num2 string) string {

	return multiplySolution(num1, num2)

}

func multiplySolution(num1, num2 string) string {

	if num1 == "0" || num2 == "0" {
		return "0"
	}

	l1, l2 := len(num1), len(num2)
	res := make([]byte, l1+l2)

	pos := len(res) - 1
	for p2 := len(num2) - 1; p2 >= 0; p2-- {
		relPos := pos
		for p1 := len(num1) - 1; p1 >= 0; p1-- {
			sum := res[relPos] + ((num1[p1] - '0') * (num2[p2] - '0'))
			res[relPos] = sum % 10
			res[relPos-1] += (sum / 10)
			relPos--
		}
		pos--
	}

	for i := 0; i < len(res); i++ {
		res[i] += '0'
	}

	if res[0] == '0' {
		res = res[1:]
	}

	return string(res)
}
