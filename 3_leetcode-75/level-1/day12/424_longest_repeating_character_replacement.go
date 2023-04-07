package slidingwindow

/*
424. Longest Repeating Character Replacement
Medium
7.8K
328
Companies

You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most k times.

Return the length of the longest substring containing the same letter you can get after performing the above operations.

Example 1:
Input: s = "ABAB", k = 2
Output: 4
Explanation: Replace the two 'A's with two 'B's or vice versa.

Example 2:
Input: s = "AABABBA", k = 1
Output: 4
Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.


Constraints:
1 <= s.length <= 105
s consists of only uppercase English letters.
0 <= k <= s.length

*/

func CharacterReplacement(s string, k int) int {

	if len(s) == 0 {
		return 0
	}

	left, right, result, windowSize, maxOccur := 0, 0, 0, 0, 0
	m := emptyMap()
	m[s[0]] = 1

	for left, right = 0, 1; right < len(s); {

		windowSize = right - left
		maxOccur = max(m)

		if windowSize-maxOccur <= k {

			if windowSize > result {
				result = windowSize
			}

			m[s[right]]++
			right++
		} else {
			m[s[left]]--
			left++
			m[s[right]]++
			right++
		}

	}

	windowSize = (right - left)
	if val := windowSize - max(m); val <= k && windowSize > result {
		return windowSize
	}

	return result

}

func emptyMap() map[byte]int {

	alphabets := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	m := make(map[byte]int)
	for i := 0; i < len(alphabets); i++ {
		m[alphabets[i]] = 0
	}
	return m
}

func getOccurences(s string) map[byte]int {

	alphabets := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	m := make(map[byte]int)
	for i := 0; i < len(alphabets); i++ {
		m[alphabets[i]] = 0
	}

	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	return m
}

func max(m map[byte]int) int {

	maxVal := 0
	for _, val := range m {

		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}
