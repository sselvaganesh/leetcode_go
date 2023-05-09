package two_pointers

/*
82. Remove Duplicates from Sorted List II
Medium
7.5K
206
Companies
Given the head of a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list. Return the linked list sorted as well.

Example 1:
Input: head = [1,2,3,3,4,4,5]
Output: [1,2,5]

Example 2:
Input: head = [1,1,1,2,3]
Output: [2,3]


Constraints:
The number of nodes in the list is in the range [0, 300].
-100 <= Node.val <= 100
The list is guaranteed to be sorted in ascending order.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	prev, cur, next := head, head, head.Next

	for next != nil {

		if cur.Val == next.Val {

			for next != nil && cur.Val == next.Val {
				cur = next
				next = next.Next
			}

			if head.Val == cur.Val {
				head = next
			} else {
				prev.Next = next
			}

			cur = next
			if next != nil {
				next = next.Next
			}

		} else {
			prev = cur
			cur = next
			next = next.Next
		}

	}

	return head
}

func deleteDuplicatesSolution(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var prevNode *ListNode
	curNode, nextNode := head, head.Next

	for nextNode != nil {

		if curNode.Val == nextNode.Val {

			for ; nextNode != nil && curNode.Val == nextNode.Val; nextNode = nextNode.Next {
			}

			if nextNode != nil {
				curNode = nextNode
				nextNode = nextNode.Next
			} else {
				curNode = nil
			}

			if prevNode != nil {
				prevNode.Next = curNode
			}

		} else {

			if prevNode == nil {
				head = curNode
			}

			prevNode, curNode, nextNode = curNode, nextNode, nextNode.Next
		}

	}

	if prevNode == nil {
		head = curNode
	}

	return head
}
