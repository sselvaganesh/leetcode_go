package twopointers

/*
Given a string s, reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

Example 1:
Input: s = "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"

Example 2:
Input: s = "God Ding"
Output: "doG gniD"
*/


import(
	"strings"
)

func reverseWords(s string) string {

	var result []byte
	space := []byte(" ")
	words := strings.Split(s, " ")

	for i, word := range words {
		result = append(result, reverse([]byte(word))...)

		if i < len(words)-1 {
			result = append(result, space...)
		}
	}

	return string(result)
}

func reverse(inp []byte) []byte {
	var result []byte
	for j := len(inp) - 1; j >= 0; j-- {
		result = append(result, inp[j])
	}
	return result
}
