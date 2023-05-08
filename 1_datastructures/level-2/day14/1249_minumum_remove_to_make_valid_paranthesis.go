package stack

/*
1249. Minimum Remove to Make Valid Parentheses
Medium
5.7K
102
Companies
Given a string s of '(' , ')' and lowercase English characters.

Your task is to remove the minimum number of parentheses ( '(' or ')', in any positions ) so that the resulting parentheses string is valid and return any valid string.

Formally, a parentheses string is valid if and only if:

It is the empty string, contains only lowercase characters, or
It can be written as AB (A concatenated with B), where A and B are valid strings, or
It can be written as (A), where A is a valid string.

Example 1:
Input: s = "lee(t(c)o)de)"
Output: "lee(t(c)o)de"
Explanation: "lee(t(co)de)" , "lee(t(c)ode)" would also be accepted.

Example 2:
Input: s = "a)b(c)d"
Output: "ab(c)d"

Example 3:
Input: s = "))(("
Output: ""
Explanation: An empty string is also valid.


Constraints:
1 <= s.length <= 105
s[i] is either'(' , ')', or lowercase English letter.
*/

func MinRemoveToMakeValid(s string) string {

	open := 0
	var res []byte

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			open++
		} else if s[i] == ')' {
			if open == 0 {
				continue
			}
			open--
		}
		res = append(res, s[i])
	}

	if open == 0 {
		return string(res)
	}

	var result []byte

	for i := len(res) - 1; i >= 0; i-- {
		if open > 0 && res[i] == '(' {
			open--
			continue
		}
		result = append([]byte{res[i]}, result...)
	}

	return string(result)
}
