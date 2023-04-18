package strings

/*
763. Partition Labels
Medium
9.4K
345
Companies

You are given a string s. We want to partition the string into as many parts as possible so that each letter appears in at most one part.

Note that the partition is done so that after concatenating all the parts in order, the resultant string should be s.

Return a list of integers representing the size of these parts.

Example 1:
Input: s = "ababcbacadefegdehijhklij"
Output: [9,7,8]
Explanation:
The partition is "ababcbaca", "defegde", "hijhklij".
This is a partition so that each letter appears in at most one part.
A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits s into less parts.

Example 2:
Input: s = "eccbbbbdec"
Output: [10]


Constraints:
1 <= s.length <= 500
s consists of lowercase English letters.

*/

func PartitionLabels(s string) []int {

	var res []int
	m := initMap(s)
	length := 0
	localMap := make(map[byte]struct{})

	for i := 0; i < len(s); i++ {

		m[s[i]]--
		length++
		localMap[s[i]] = struct{}{}
		if isAllZero(m, localMap) {
			res = append(res, length)
			length = 0
		}

	}
	return res
}

func initMap(s string) map[byte]int {

	m := make(map[byte]int)
	lower := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < len(lower); i++ {
		m[lower[i]] = 0
	}

	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	return m
}

func isAllZero(m map[byte]int, curStr map[byte]struct{}) bool {

	for k, _ := range curStr {
		if val, _ := m[k]; val != 0 {
			return false
		}
	}

	return true
}
