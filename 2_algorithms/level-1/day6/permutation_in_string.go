package slidingwindow

/*
Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.

In other words, return true if one of s1's permutations is the substring of s2.


Example 1:
Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Explanation: s2 contains one permutation of s1 ("ba").

Example 2:
Input: s1 = "ab", s2 = "eidboaoo"
Output: false

*/

import (
	"reflect"
)

func CheckInclusion(s1 string, s2 string) bool {

	if len(s1) == 0 {
		return true
	} else if len(s2) == 0 || len(s1) > len(s2) {
		return false
	}

	s1Hash := hash(s1)
	s2Hash := hash(s2[:len(s1)])
	start := 0

	for i := len(s1); i < len(s2); i++ {

		if reflect.DeepEqual(s1Hash, s2Hash) {
			return true
		}

		s2Hash[s2[start]]--
		s2Hash[s2[i]]++
		start++

	}

	if reflect.DeepEqual(s1Hash, s2Hash) {
		return true
	}

	return false
}

func hash(s string) map[byte]int {

	m := make(map[byte]int)

	alphabets := "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i < len(alphabets); i++ {
		m[alphabets[i]] = 0
	}

	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	return m
}
