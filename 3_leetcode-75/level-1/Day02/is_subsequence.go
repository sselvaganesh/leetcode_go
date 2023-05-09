package daytwo

/*
Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

Input: s = "abc", t = "ahbgdc"
Output: true

Input: s = "axc", t = "ahbgdc"
Output: false
*/

func isSubsequence(s string, t string) bool {

	if s == "" && t == "" {
		return true
	} else if s == "" {
		return true
	}

	sIdx := 0
	for i := 0; i < len(t); i++ {
		if s[sIdx] == t[i] {
			if sIdx+1 == len(s) {
				return true
			} else {
				sIdx++
			}
		}
	}

	return false

}
