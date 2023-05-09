package linkedlist

/*
24. Swap Nodes in Pairs
Medium
9.6K
365
Companies
Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)

Example 1:
Input: head = [1,2,3,4]
Output: [2,1,4,3]

Example 2:
Input: head = []
Output: []

Example 3:
Input: head = [1]
Output: [1]


Constraints:
The number of nodes in the list is in the range [0, 100].
0 <= Node.val <= 100
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func SwapPairs(head *ListNode) *ListNode {

	return swap(head)

}

func swap(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	curNode, nextNode := head, head.Next
	var newHead, last *ListNode

	for nextNode != nil {

		next := nextNode.Next
		first, second := swapHelp(curNode, nextNode)

		if newHead == nil {
			newHead = first
		} else {
			last.Next = first
		}
		last = second

		if next == nil {
			break
		} else if next.Next == nil {
			last.Next = next
			break
		}

		curNode, nextNode = next, next.Next
	}

	return newHead
}

func swapHelp(node1, node2 *ListNode) (*ListNode, *ListNode) {

	node1.Next = nil
	if node2 == nil {
		return node1, nil
	}

	node2.Next = node1
	return node2, node1
}
