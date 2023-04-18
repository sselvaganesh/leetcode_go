package strings

/*
409. Longest Palindrome
Easy
4.6K
283
Companies

Given a string s which consists of lowercase or uppercase letters, return the length of the longest palindrome that can be built with those letters.

Letters are case sensitive, for example, "Aa" is not considered a palindrome here.

Example 1:
Input: s = "abccccdd"
Output: 7
Explanation: One longest palindrome that can be built is "dccaccd", whose length is 7.

Example 2:
Input: s = "a"
Output: 1
Explanation: The longest palindrome that can be built is "a", whose length is 1.

Constraints:
1 <= s.length <= 2000
s consists of lowercase and/or uppercase English letters only.
*/

func LongestPalindrome(s string) int {

	m := initMap()

	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	res := 0
	addOne := false

	for _, charCount := range m {

		if charCount%2 == 0 {
			res += charCount
			continue
		}

		res += charCount - 1
		if !addOne {
			addOne = true
		}
	}

	if addOne {
		return res + 1
	}

	return res
}

func initMap() map[byte]int {

	m := make(map[byte]int)
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i := 0; i < len(lower); i++ {
		m[lower[i]] = 0
	}

	for i := 0; i < len(upper); i++ {
		m[upper[i]] = 0
	}
	return m
}
