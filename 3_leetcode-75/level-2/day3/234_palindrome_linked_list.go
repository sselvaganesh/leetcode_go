package linked_list

/*
234. Palindrome Linked List
Easy
13.7K
749
Companies

Given the head of a singly linked list, return true if it is a
palindrome
 or false otherwise.

Example 1:
Input: head = [1,2,2,1]
Output: true

Example 2:
Input: head = [1,2]
Output: false


Constraints:
The number of nodes in the list is in the range [1, 105].
0 <= Node.val <= 9

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func IsPalindrome(head *ListNode) bool {

	if head == nil {
		return true
	}

	var elems []int
	curNode := head

	for curNode != nil {
		elems = append(elems, curNode.Val)
		curNode = curNode.Next
	}

	for i, j := 0, len(elems)-1; j > i; i, j = i+1, j-1 {
		if elems[i] != elems[j] {
			return false
		}

	}
	return true

}
