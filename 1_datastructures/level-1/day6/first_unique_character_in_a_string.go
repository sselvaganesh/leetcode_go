package string

/*
Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.

Example 1:
Input: s = "leetcode"
Output: 0

Example 2:
Input: s = "loveleetcode"
Output: 2

Example 3:
Input: s = "aabb"
Output: -1

*/

func FirstUniqChar(s string) int {

	if len(s) == 0 {
		return -1
	} else if len(s) == 1 {
		return 0
	}

	m := hash1(s)
	for i := 0; i < len(s); i++ {

		if m[s[i]] == 1 {
			return i
		}

	}
	return -1

}

func hash1(s string) map[byte]int {

	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {

		if _, ok := m[s[i]]; ok {
			m[s[i]]++
		} else {
			m[s[i]] = 1
		}
	}

	return m
}
