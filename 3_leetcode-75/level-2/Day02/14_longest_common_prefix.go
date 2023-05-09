package strings

/*
14. Longest Common Prefix
Easy
13.4K
3.9K
Companies

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".


Example 1:
Input: strs = ["flower","flow","flight"]
Output: "fl"

Example 2:
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.


Constraints:
1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters.
*/

func LongestCommonPrefix(strs []string) string {

	return bruteForceLCP(strs)
}

func LongestCommonPrefixSolution(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	var res []byte
	for i := 0; i < len(strs[0]); i++ {

		success := true
		for j := 1; j < len(strs); j++ {

			if i < len(strs[j]) && strs[0][i] == strs[j][i] {
				continue
			} else {
				success = false
				break
			}

		}

		if success {
			res = append(res, strs[0][i])
		} else {
			break
		}

	}

	return string(res)
}

func bruteForceLCP(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	var res []byte
	for i := 0; i < len(strs[0]); i++ {

		success := false
		for j := 0; j < len(strs); j++ {

			if len(strs[j]) < i+1 {
				break
			}

			if strs[0][i] != strs[j][i] {
				success = false
				break
			}

			if j == len(strs)-1 {
				success = true
			}
		}

		if success {
			res = append(res, strs[0][i])
		} else {
			break
		}

	}

	return string(res)
}
