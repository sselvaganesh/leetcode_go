package twopointers

/*
Write a function that reverses a string. The input string is given as an array of characters s.
You must do this by modifying the input array in-place with O(1) extra memory.

Example 1:
Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]

Example 2:
Input: s = ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]

*/

func ReverseString(s []byte) {

	if s == nil {
		return
	}

	for i, j := 0, len(s)-1; j > i; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return
}
