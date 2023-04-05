package longestsubstring

import (
	"fmt"
)

var l = fmt.Println
var result = make(map[string]bool)

func LengthOfLongestSubstring(s string) int {

	if len(s) <= 1 {
		return len(s)
	}

	if len(longestSubstring(s)) > 0 {
		return len(s)
	}

	return max(LengthOfLongestSubstring(s[1:]), LengthOfLongestSubstring(s[:len(s)-1]))
}

func longestSubstring(s string) string {


	if len(s) <= 1 {
		return s
	}

	if valid, ok := result[s]; valid && ok {
		return s
	}

	m := make(map[byte]struct{})
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			result[s] = false
			return ""
		}
		m[s[i]] = struct{}{}
	}

	result[s] = true
	return s

}

func max(s1, s2 int) int {

	if s1 >= s2 {
		return s1
	}

	return s2
}
