package string

/*
5. Longest Palindromic Substring
Medium
24.9K
1.5K
Companies

Given a string s, return the longest palindromic substring in s.

Example 1:
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.

Example 2:
Input: s = "cbbd"
Output: "bb"


Constraints:
1 <= s.length <= 1000
s consist of only digits and English letters.
*/

func LongestPalindrome(s string) string {

	m := make(map[string]bool)
	res := ""
	return helper(s, &m, &res)

}

func helper(s string, cache *map[string]bool, res *string) string {

	result := *res
	m := *cache

	if len(result) > 0 && len(result) >= len(s) {
		return result
	}

	if palindrome, ok := m[s]; ok && palindrome {
		return s
	} else if ok && !palindrome {
		return ""
	}

	if isPalindrome(s) {
		m[s] = true
		return s
	} else {
		m[s] = false
	}

	return max(helper(s[1:], cache, res), helper(s[:len(s)-1], cache, res))

}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; j > i; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func max(s1, s2 string) string {
	if len(s1) > len(s2) {
		return s1
	}
	return s2
}
