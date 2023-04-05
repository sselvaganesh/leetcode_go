package longestpalindrome

import (
	//"fmt"
	//"time"
)

func LongestPalindrome(s string) string {
	m := make(map[string]bool)
	longestStr := ""
	return helper(s, &m, &longestStr)
}

func helper(s string, m *map[string]bool, longestStr *string) string {

	if len(*longestStr) > 0 && len(*longestStr) >= len(s) {
		return *longestStr
	}

	if success, ok := (*m)[s]; ok && success {
		return s
	} else if ok && !success {
		return ""
	}

	if len(s) <= 1 {
		*longestStr = s
		(*m)[s] = true
		return s
	}

	if isPalindrome(s) {
		*longestStr = s
		(*m)[s] = true
		return s
	} else {
		(*m)[s] = false
	}

	p1 := s[:len(s)-1]
	p2 := s[1:]

	return longest(helper(p1, m, longestStr), helper(p2, m, longestStr))
}

func isPalindrome(inp string) bool {


	for i, j := 0, len(inp)-1; j > i; i, j = i+1, j-1 {
		if inp[i] != inp[j] {
			return false
		}
	}

	//fmt.Println("==== > Palindrom found: ", inp)
	//time.Sleep(1*time.Second)
	return true
}

func longest(s1, s2 string) string {

	if len(s1) >= len(s2) {
		return s1
	}
	return s2
}
