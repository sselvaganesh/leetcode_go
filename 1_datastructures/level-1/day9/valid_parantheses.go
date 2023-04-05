package stackqueue

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

Example 1:
Input: s = "()"
Output: true

Example 2:
Input: s = "()[]{}"
Output: true

Example 3:
Input: s = "(]"
Output: false

*/

func IsValid(s string) bool {

	if len(s) == 0 {
		return true
	} else if len(s) == 1 {
		return false
	}

	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	var queue []byte
	for i := 0; i < len(s); i++ {

		val, ok := m[s[i]]
		if !ok {
			queue = append(queue, s[i])
		} else {

			if len(queue) == 0 || queue[len(queue)-1] != val {
				return false
			}
			queue = queue[:len(queue)-1]
		}
	}

	if len(queue) == 0 {
		return true
	}

	return false
}
