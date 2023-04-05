package string

/*
Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.
Each letter in magazine can only be used once in ransomNote.

Example 1:
Input: ransomNote = "a", magazine = "b"
Output: false

Example 2:
Input: ransomNote = "aa", magazine = "ab"
Output: false

Example 3:
Input: ransomNote = "aa", magazine = "aab"
Output: true

*/

func CanConstruct(ransomNote string, magazine string) bool {

	//return hashMapSolution(ransomNote, magazine)

	return arraySolution(ransomNote, magazine)

}

func arraySolution(ransomNote, magazine string) bool {

	if (len(ransomNote) == 0 && len(magazine) == 0) || len(ransomNote) == 0 {
		return true
	} else if len(magazine) == 0 {
		return false
	}

	charCount := countChars1(magazine)
	for i := 0; i < len(ransomNote); i++ {

		idx := ransomNote[i] - 'a'
		if charCount[idx] == 0 {
			return false
		} else {
			charCount[idx]--
		}
	}

	return true

}

func countChars1(s string) [26]int {

	var result [26]int
	for i := 0; i < len(s); i++ {
		result[s[i]-'a']++
	}
	return result
}

func hashMapSolution(ransomNote string, magazine string) bool {

	if (len(ransomNote) == 0 && len(magazine) == 0) || len(ransomNote) == 0 {
		return true
	} else if len(magazine) == 0 {
		return false
	}

	m := hash(magazine)
	for i := 0; i < len(ransomNote); i++ {

		val, ok := m[ransomNote[i]]
		if !ok {
			return false
		} else {
			if val == 1 {
				delete(m, ransomNote[i])
			} else {
				m[ransomNote[i]]--
			}
		}
	}

	return true

}

func hash(s string) map[byte]int {

	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {

		if _, ok := m[s[i]]; ok {
			m[s[i]]++
		} else {
			m[s[i]] = 1
		}

	}

	return m
}
