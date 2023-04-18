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

	m1 := make(map[byte]string)
	m2 := make(map[string]byte)

	for i, j := 0, 0; i < len(pattern) && j < len(words); i, j = i+1, j+1 {

		val1, ok1 := m1[pattern[i]]
		val2, ok2 := m2[words[j]]

		if !ok1 && !ok2 {
			m1[pattern[i]] = words[j]
			m2[words[j]] = pattern[i]
			continue
		} else if ok1 && ok2 {
			if val1 != words[j] || val2 != pattern[i] {
				return false
			}
		} else {
			return false
		}

	}

	return true

}
