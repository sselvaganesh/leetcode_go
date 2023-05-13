package heap

/*
451. Sort Characters By Frequency
Medium
6.6K
228
Companies
Given a string s, sort it in decreasing order based on the frequency of the characters. The frequency of a character is the number of times it appears in the string.

Return the sorted string. If there are multiple answers, return any of them.


Example 1:
Input: s = "tree"
Output: "eert"
Explanation: 'e' appears twice while 'r' and 't' both appear once.
So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.

Example 2:
Input: s = "cccaaa"
Output: "aaaccc"
Explanation: Both 'c' and 'a' appear three times, so both "cccaaa" and "aaaccc" are valid answers.
Note that "cacaca" is incorrect, as the same characters must be together.

Example 3:
Input: s = "Aabb"
Output: "bbAa"
Explanation: "bbaA" is also a valid answer, but "Aabb" is incorrect.
Note that 'A' and 'a' are treated as two different characters.

Constraints:
1 <= s.length <= 5 * 105
s consists of uppercase and lowercase English letters and digits.
*/

import (
	"bytes"
	"sort"
)

func FrequencySort(s string) string {

	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	type elem struct {
		alpha byte
		occur int
	}
	var elems []elem

	for k, v := range m {
		elems = append(elems, elem{k, v})
	}

	sort.Slice(elems, func(i, j int) bool {
		return elems[i].occur > elems[j].occur
	})

	var res []byte
	for i := 0; i < len(elems); i++ {
		alpha := elems[i].alpha
		occur := elems[i].occur
		res = append(res, bytes.Repeat([]byte{alpha}, occur)...)
	}

	return string(res)
}
