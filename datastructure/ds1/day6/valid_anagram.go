package string

/*
Given two strings s and t, return true if t is an anagram of s, and false otherwise.
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:
Input: s = "anagram", t = "nagaram"
Output: true

Example 2:
Input: s = "rat", t = "car"
Output: false
*/

import(
	"reflect"
)

func IsAnagram(s string, t string) bool {

	return reflect.DeepEqual(countChars(s), countChars(t))

}

func countChars(s string) [26]int {

	result := [26]int{}
	for i := 0; i < len(s); i++ {
		result[s[i]-'a']++
	}

	return result
}
