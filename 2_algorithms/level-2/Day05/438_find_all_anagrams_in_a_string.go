package sliding_window

/*
438. Find All Anagrams in a String
Medium
11.1K
313
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

	var res []int
	sMap, pMap := make(map[byte]int), getCharCount(p)
	start := 0

	for i := 0; i < len(s); i++ {

		if _, ok := sMap[s[i]]; !ok {
			sMap[s[i]] = 1
		} else {
			sMap[s[i]]++
		}

		if i >= len(p)-1 {
			if reflect.DeepEqual(sMap, pMap) {
				res = append(res, start)
			}
			if val, _ := sMap[s[start]]; val == 1 {
				delete(sMap, s[start])
			} else {
				sMap[s[start]]--
			}
			start++
		}

	}
	return res

}

func getCharCount(p string) map[byte]int {
	m := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		if _, ok := m[p[i]]; !ok {
			m[p[i]] = 1
		} else {
			m[p[i]]++
		}
	}
	return m
}
