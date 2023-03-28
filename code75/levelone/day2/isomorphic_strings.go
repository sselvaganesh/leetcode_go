package daytwo

/*
Given two strings s and t, determine if they are isomorphic.
Two strings s and t are isomorphic if the characters in s can be replaced to get t.
All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.


Input: s = "egg", t = "add"
Output: true

Input: s = "foo", t = "bar"
Output: false

Input: s = "paper", t = "title"
Output: true

aaabbbab
aaabbbba

*/

func IsIsomorphic(s string, t string) bool {

	sMap := make(map[byte]byte)
	tMap := make(map[byte]byte)

	for i := 0; i < len(s); i++ {

		tVal, sOk := sMap[s[i]]
		sVal, tOk := tMap[t[i]]

		if sOk && tVal != t[i] {
			return false
		} else if tOk && sVal != s[i] {
			return false
		} else {
			sMap[s[i]] = t[i]
			tMap[t[i]] = s[i]
		}

	}

	return true

}
