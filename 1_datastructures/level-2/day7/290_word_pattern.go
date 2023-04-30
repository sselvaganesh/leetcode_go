package strings

/*
290. Word Pattern
Easy
6.1K
722
Companies

Given a pattern and a string s, find if s follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.

Example 1:
Input: pattern = "abba", s = "dog cat cat dog"
Output: true

Example 2:
Input: pattern = "abba", s = "dog cat cat fish"
Output: false

Example 3:
Input: pattern = "aaaa", s = "dog cat cat dog"
Output: false


Constraints:
1 <= pattern.length <= 300
pattern contains only lower-case English letters.
1 <= s.length <= 3000
s contains only lowercase English letters and spaces ' '.
s does not contain any leading or trailing spaces.
All the words in s are separated by a single space.
*/

import (
	"strings"
)

func WordPattern(pattern string, s string) bool {

	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	pMap, sMap := make(map[byte]string), make(map[string]byte)
	for i := 0; i < len(pattern); i++ {

		pVal, pOk := pMap[pattern[i]]
		sVal, sOk := sMap[words[i]]

		if !pOk && !sOk {
			pMap[pattern[i]] = words[i]
			sMap[words[i]] = pattern[i]
		} else if pOk && sOk {
			if pVal != words[i] || sVal != pattern[i] {
				return false
			}
		} else {
			return false
		}
	}

	return true

}
