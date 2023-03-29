package linkedlist

/*
Given the head of a singly linked list, reverse the list, and return the reversed list.

Example 1:
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

Example 2:
Input: head = [1,2]
Output: [2,1]

Example 3:
Input: head = []
Output: []

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func ReverseList(head *ListNode) *ListNode {

	return solution2(head)
}

func solution2(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var (
		curNode           = head
		prev    *ListNode = nil
		next    *ListNode = nil
	)

	for curNode != nil {
		next = curNode.Next
		curNode.Next = prev
		prev = curNode
		curNode = next
	}

	return prev

}

func solution1(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	curNode := head
	var q []int
	for curNode != nil {
		q = append(q, curNode.Val)
		curNode = curNode.Next
	}

	var result *ListNode
	for i := len(q) - 1; i >= 0; i-- {

		node := newNode(q[i])
		if result == nil {
			result = node
		} else {
			curNode.Next = node
		}
		curNode = node

	}

	return result

}

func newNode(val int) *ListNode {
	node := new(ListNode)
	node.Val = val
	node.Next = nil
	return node
}
