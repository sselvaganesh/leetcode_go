package strings

/*
415. Add Strings
Easy
4.4K
647
Companies

Given two non-negative integers, num1 and num2 represented as string, return the sum of num1 and num2 as a string.

You must solve the problem without using any built-in library for handling large integers (such as BigInteger). You must also not convert the inputs to integers directly.

Example 1:
Input: num1 = "11", num2 = "123"
Output: "134"

Example 2:
Input: num1 = "456", num2 = "77"
Output: "533"

Example 3:
Input: num1 = "0", num2 = "0"
Output: "0"


Constraints:
1 <= num1.length, num2.length <= 104
num1 and num2 consist of only digits.
num1 and num2 don't have any leading zeros except for the zero itself.
*/

func AddStrings(num1 string, num2 string) string {

	var res []byte
	if len(num1) > len(num2) {
		res = make([]byte, len(num1)+1)
	} else {
		res = make([]byte, len(num2)+1)
	}

	i, j, idx := len(num1)-1, len(num2)-1, len(res)-1
	var sum, rem uint8

	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		sum = (num1[i] - '0') + (num2[j] - '0') + rem
		res[idx] = sum % 10
		rem = sum / 10
		idx--
	}

	for ; i >= 0; i-- {
		sum = num1[i] - '0' + rem
		res[idx] = sum % 10
		rem = sum / 10
		idx--
	}

	for ; j >= 0; j-- {
		sum = num2[j] - '0' + rem
		res[idx] = sum % 10
		rem = sum / 10
		idx--
	}

	if rem > 0 {
		res[idx] = rem
	}

	for i := 0; i < len(res); i++ {
		res[i] = res[i] + '0'
	}

	if res[0] == '0' {
		res = res[1:]
	}

	return string(res)
}

func addStringsSolution(num1, num2 string) string {

	max := 0
	if len(num1) > len(num2) {
		max = len(num1) + 1
	} else {
		max = len(num2) + 1
	}

	result := make([]byte, max)
	pos := max - 1
	var sum uint8

	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {

		if i >= 0 && j >= 0 {
			sum = result[pos] + (num1[i] - '0') + (num2[j] - '0')
		} else if i >= 0 {
			sum = result[pos] + (num1[i] - '0')
		} else {
			sum = result[pos] + (num2[j] - '0')
		}
		result[pos] = sum % 10
		result[pos-1] = result[pos-1] + (sum / 10)
		pos--
	}

	for i := 0; i < len(result); i++ {
		result[i] += '0'
	}

	if result[0] == '0' {
		result = result[1:]
	}

	return string(result)

}
