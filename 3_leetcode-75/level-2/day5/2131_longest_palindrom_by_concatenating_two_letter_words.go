package greedy

/*
2131. Longest Palindrome by Concatenating Two Letter Words
Medium
2.2K
44
Companies

You are given an array of strings words. Each element of words consists of two lowercase English letters.

Create the longest possible palindrome by selecting some elements from words and concatenating them in any order. Each element can be selected at most once.

Return the length of the longest palindrome that you can create. If it is impossible to create any palindrome, return 0.

A palindrome is a string that reads the same forward and backward.


Example 1:
Input: words = ["lc","cl","gg"]
Output: 6
Explanation: One longest palindrome is "lc" + "gg" + "cl" = "lcggcl", of length 6.
Note that "clgglc" is another longest palindrome that can be created.

Example 2:
Input: words = ["ab","ty","yt","lc","cl","ab"]
Output: 8
Explanation: One longest palindrome is "ty" + "lc" + "cl" + "yt" = "tylcclyt", of length 8.
Note that "lcyttycl" is another longest palindrome that can be created.

Example 3:
Input: words = ["cc","ll","xx"]
Output: 2
Explanation: One longest palindrome is "cc", of length 2.
Note that "ll" is another longest palindrome that can be created, and so is "xx".

Constraints:

1 <= words.length <= 105
words[i].length == 2
words[i] consists of lowercase English letters.
*/

func LongestPalindrome(words []string) int {

	m1 := make(map[string]int)
	m2 := make(map[string]int)
	rev := func(word string) string {
		return string([]byte{word[1], word[0]})
	}

	res := 0
	for _, word := range words {
		revWord := rev(word)
		if word == revWord {
			if _, ok := m1[word]; !ok {
				m1[word] = 1
			} else {
				m1[word]++
			}
		} else {
			if occur, ok := m2[revWord]; ok && occur > 0 {
				res += 4
				m2[revWord]--
			} else if occur, ok := m2[word]; ok && occur > 0 {
				m2[word]++
			} else {
				m2[word] = 1
			}
		}
	}

	addExtra := false
	for _, occur := range m1 {
		if occur%2 == 0 {
			res += (occur * 2)
		} else {
			addExtra = true
			res += ((occur - 1) * 2)
		}
	}

	if addExtra {
		return res + 2
	}

	return res
}
