package linkedlist

/*
25. Reverse Nodes in k-Group
Hard
11K
572
Companies

Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.

You may not alter the values in the list's nodes, only nodes themselves may be changed.

Example 1:
Input: head = [1,2,3,4,5], k = 2
Output: [2,1,4,3,5]

Example 2:
Input: head = [1,2,3,4,5], k = 3
Output: [3,2,1,4,5]


Constraints:
The number of nodes in the list is n.
1 <= k <= n <= 5000
0 <= Node.val <= 1000
*/

func ReverseKGroup(head *ListNode, k int) *ListNode {

	if head == nil || k == 0 {
		return head
	}

	counter := 1
	curNode := head
	var newHead, first, last *ListNode

	for curNode != nil {

		if first == nil {
			first = curNode
		}

		if counter == k {

			next := curNode.Next
			curNode.Next = nil
			revHead, revTail := reverseKNodes(first)

			if newHead == nil {
				newHead = revHead
			} else {
				last.Next = revHead
			}
			last = revTail

			first = nil
			counter = 1
			curNode = next
			continue
		}

		counter++
		curNode = curNode.Next

	}

	if last != nil {
		last.Next = first
	}

	return newHead

}

func reverseKNodes(first *ListNode) (*ListNode, *ListNode) {

	/* Contains k nodes always (atleast 1 node)
	   last node is has nil value (Next)
	*/

	head, curNode := first, first
	var prev *ListNode

	for curNode != nil {
		next := curNode.Next
		curNode.Next = prev
		prev, curNode = curNode, next
	}

	return prev, head
}
