package linkedlist

/*
82. Remove Duplicates from Sorted List II
Medium
7.6K
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

func DeleteDuplicates(head *ListNode) *ListNode {

	return delete(head)

}

func delete(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	curNode, nextNode := head, head.Next

	for nextNode != nil {

		if curNode.Val == nextNode.Val {
			for nextNode != nil && curNode.Val == nextNode.Val {
				curNode, nextNode = nextNode, nextNode.Next
			}

			curNode = nextNode
			if nextNode != nil {
				nextNode = nextNode.Next
			}

			if prev != nil {
				prev.Next = curNode
			}

		} else {

			if prev == nil {
				head = curNode
			}
			prev, curNode, nextNode = curNode, nextNode, nextNode.Next
		}
	}

	if prev == nil {
		head = curNode
	}

	return head
}
