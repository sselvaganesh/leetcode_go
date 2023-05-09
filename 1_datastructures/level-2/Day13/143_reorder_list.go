package linkedlist

/*
143. Reorder List
Medium
8.7K
291
Companies

You are given the head of a singly linked-list. The list can be represented as:

L0 → L1 → … → Ln - 1 → Ln
Reorder the list to be on the following form:

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
You may not modify the values in the list's nodes. Only nodes themselves may be changed.


Example 1:
Input: head = [1,2,3,4]
Output: [1,4,2,3]

Example 2:
Input: head = [1,2,3,4,5]
Output: [1,5,2,4,3]


Constraints:
The number of nodes in the list is in the range [1, 5 * 104].
1 <= Node.val <= 1000
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func ReorderList(head *ListNode) {

	if head == nil {
		return
	}

	var deq []*ListNode
	curNode := head
	for curNode != nil {
		deq = append(deq, curNode)
		curNode = curNode.Next
	}

	var last *ListNode

	for i, j := 0, len(deq)-1; j >= i; i, j = i+1, j-1 {

		if last != nil {
			last.Next = deq[i]
			if i == j {
				last = deq[i]
				break
			}
		}
		deq[i].Next = deq[j]
		last = deq[j]
	}

	last.Next = nil
	return
}
