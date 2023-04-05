package greedy

/*
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
*/

func LongestPalindrome(s string) int {

	return hashTable(s)

}

func hashTable(s string) int {

	if len(s) <= 1 {
		return len(s)
	}

	m := make(map[byte]int) // map[char]occurences
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			m[s[i]]++
		} else {
			m[s[i]] = 1
		}
	}

	result := 0
	includeChar := false

	for _, occurences := range m {

		if occurences%2 == 0 {
			result += occurences
		} else {
			result = result + occurences - 1
			includeChar = true
		}
	}

	if includeChar {
		return result + 1
	}

	return result
}
