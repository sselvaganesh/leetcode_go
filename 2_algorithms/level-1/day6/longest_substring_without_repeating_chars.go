package slidingwindow

/*
Given a string s, find the length of the longest
substring
 without repeating characters.

Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func LengthOfLongestSubstring(s string) int {

	if len(s) <= 1 {
		return len(s)
	}

	m := make(map[byte]int)
	start, longest := 0, 0

	for idx := 0; idx < len(s); idx++ {

		if val, ok := m[s[idx]]; ok && val >= start {
			if idx-start > longest {
				longest = idx - start
			}
			start = val + 1
		}

		m[s[idx]] = idx
	}

	if len(s)-start > longest {
		return len(s) - start
	}

	return longest
}
