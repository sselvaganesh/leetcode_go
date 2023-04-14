package two_pointers

/*
844. Backspace String Compare
Easy
6.3K
287
Companies

Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.

Note that after backspacing an empty text, the text will continue empty.

Example 1:
Input: s = "ab#c", t = "ad#c"
Output: true
Explanation: Both s and t become "ac".

Example 2:
Input: s = "ab##", t = "c#d#"
Output: true
Explanation: Both s and t become "".

Example 3:
Input: s = "a#c", t = "b"
Output: false
Explanation: s becomes "c" while t becomes "b".


Constraints:
1 <= s.length, t.length <= 200
s and t only contain lowercase letters and '#' characters.
*/

func BackspaceCompare(s string, t string) bool {

	return prune(s) == prune(t)

}

func prune(str string) string {

	var res []byte

	for i := 0; i < len(str); i++ {
		if str[i] == '#' {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		} else {
			res = append(res, str[i])
		}
	}
	return string(res)
}
