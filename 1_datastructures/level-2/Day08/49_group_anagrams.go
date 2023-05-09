package string

/*
49. Group Anagrams

Medium
15K
433
Companies
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Example 2:
Input: strs = [""]
Output: [[""]]

Example 3:
Input: strs = ["a"]
Output: [["a"]]


Constraints:
1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] consists of lowercase English letters.
*/

import (
	"strconv"
)

func GroupAnagrams(strs []string) [][]string {

	m := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		k := key(strs[i])
		if val, ok := m[k]; !ok {
			m[k] = []string{strs[i]}
		} else {
			val = append(val, strs[i])
			m[k] = val
		}
	}
	return getResult(m)
}

func getResult(inp map[string][]string) [][]string {

	var result [][]string
	for _, strSlice := range inp {
		result = append(result, strSlice)
	}
	return result
}

func key(s string) string {
	slice := [26]byte{}
	for i := 0; i < len(s); i++ {
		slice[s[i]-'a']++
	}

	var res string
	for i := 0; i < len(slice); i++ {
		res = res + strconv.Itoa(int(i)) + "|" + strconv.Itoa(int(slice[i]))
	}
	return res
}
