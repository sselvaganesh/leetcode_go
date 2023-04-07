package slidingwindow

/*
438. Find All Anagrams in a String
Medium
11.1K
312
Companies

Given two strings s and p, return an array of all the start indices of p's anagrams in s. You may return the answer in any order.
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:
Input: s = "cbaebabacd", p = "abc"
Output: [0,6]
Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".

Example 2:
Input: s = "abab", p = "ab"
Output: [0,1,2]
Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".


Constraints:
1 <= s.length, p.length <= 3 * 104
s and p consist of lowercase English letters.

*/

import (
	"reflect"
)

func FindAnagrams(s string, p string) []int {

	return solution1(s, p)
}

func solution1(s, p string) []int {

	if len(p) > len(s) {
		return nil
	}

	var result []int
	sMap, pMap := initMap(s[:len(p)]), initMap(p)

	for i, j := 0, len(p); j < len(s); i, j = i+1, j+1 {

		if reflect.DeepEqual(sMap, pMap) {
			result = append(result, i)
		}

		sMap[s[i]]--
		sMap[s[j]]++
	}

	if reflect.DeepEqual(sMap, pMap) {
		result = append(result, len(s)-len(p))
	}

	return result
}

func initMap(s string) map[byte]int {

	alphabets := "abcdefghijklmnopqrstuvwxyz"

	m := make(map[byte]int)
	for i := 0; i < len(alphabets); i++ {
		m[alphabets[i]] = 0
	}

	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	return m
}
