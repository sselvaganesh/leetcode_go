package linked_list

/*
328. Odd Even Linked List
Medium
8.1K
445
Companies

Given the head of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return the reordered list.

The first node is considered odd, and the second node is even, and so on.

Note that the relative order inside both the even and odd groups should remain as it was in the input.

You must solve the problem in O(1) extra space complexity and O(n) time complexity.

Example 1:
Input: head = [1,2,3,4,5]
Output: [1,3,5,2,4]

Example 2:
Input: head = [2,1,3,5,6,4,7]
Output: [2,3,6,7,1,5,4]


Constraints:
The number of nodes in the linked list is in the range [0, 104].
-106 <= Node.val <= 106
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func OddEvenList(head *ListNode) *ListNode {

	var evenHead, evenTail *ListNode
	curNode := head

	for curNode != nil {

		evenNode := curNode.Next
		if evenNode != nil {
			if evenHead == nil {
				evenHead = evenNode
			} else {
				evenTail.Next = evenNode
			}

			evenTail = evenNode
			curNode.Next = evenNode.Next
		}

		if curNode.Next == nil {
			curNode.Next = evenHead
			if evenTail != nil {
				evenTail.Next = nil
			}
			break
		}

		curNode = curNode.Next
	}

	return head
}

func oddEvenListSolution(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	var oddTail, evenHead, evenTail *ListNode
	curNode := head
	counter := 1

	for curNode != nil {
		if counter%2 == 1 {
			if oddTail == nil {
				oddTail = curNode
			} else {
				oddTail.Next = curNode
				oddTail = curNode
			}
		} else {
			if evenHead == nil {
				evenHead, evenTail = curNode, curNode
			} else {
				evenTail.Next = curNode
				evenTail = curNode
			}
		}
		curNode = curNode.Next
		counter++
	}

	if evenTail != nil {
		evenTail.Next = nil
	}
	oddTail.Next = evenHead

	return head

}
