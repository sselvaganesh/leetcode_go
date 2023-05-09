package backtracking

/*
784. Letter Case Permutation
Medium
4.2K
151
Companies

Given a string s, you can transform every letter individually to be lowercase or uppercase to create another string.

Return a list of all possible strings we could create. Return the output in any order.

Example 1:
Input: s = "a1b2"
Output: ["a1b2","a1B2","A1b2","A1B2"]

Example 2:
Input: s = "3z4"
Output: ["3z4","3Z4"]

Constraints:
1 <= s.length <= 12
s consists of lowercase English letters, uppercase English letters, and digits.


*/

import (
//"strings"
)

func LetterCasePermutation(s string) []string {

	return letterCombination([]byte(s))
}

func letterCombination(str []byte) []string {

	if len(str) == 0 {
		return nil
	}

	//fmt.Println("str: ", string(str))
	var result []string

	for i := 0; i < len(str); i++ {

		if (str[i] >= 'a' && str[i] <= 'z') || (str[i] >= 'A' && str[i] <= 'Z') {

			var c1, c2 string
			c1 = strCopy(str[:i+1])
			if str[i] >= 'a' && str[i] <= 'z' {
				c2 = strCopy(str[:i]) + string(str[i]-32)
			} else {
				c2 = strCopy(str[:i]) + string(str[i]+32)
			}

			temp := letterCombination(str[i+1:])
			if len(temp) > 0 {
				for i := 0; i < len(temp); i++ {
					result = append(result, c1+strCopy([]byte(temp[i])))
					result = append(result, c2+strCopy([]byte(temp[i])))
				}
			} else {
				result = append(result, c1)
				result = append(result, c2)
			}

			break
		}

	}

	if len(result) == 0 {
		return []string{string(str)}
	}

	//fmt.Println("result: ", result)
	return result
}

func strCopy(inp []byte) string {
	return string(append([]byte{}, inp...))
}
