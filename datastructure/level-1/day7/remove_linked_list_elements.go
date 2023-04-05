package linkedlist

/*
*

Given the head of a linked list and an integer val, remove all the nodes of the linked list that has Node.val == val, and return the new head.

Example 1:
Input: head = [1,2,6,3,4,5,6], val = 6
Output: [1,2,3,4,5]

Example 2:
Input: head = [], val = 1
Output: []

Example 3:
Input: head = [7,7,7,7], val = 7
Output: []

  - Definition for singly-linked list.
  - type ListNode struct {
  - Val int
  - Next *ListNode
  - }
*/
func RemoveElements(head *ListNode, val int) *ListNode {

	if head == nil {
		return nil
	}

	curNode, prev := head, head

	for curNode != nil {

		if curNode.Val == val {

			if head == curNode {
				head = curNode.Next
				prev = curNode.Next
				curNode = curNode.Next
				continue
			}

			if curNode.Next == nil {
				prev.Next = nil
				break
			} else if curNode.Next.Val != val {
				prev.Next = curNode.Next
				curNode = curNode.Next
			} else {
				curNode = curNode.Next
			}

		} else {
			prev = curNode
			curNode = curNode.Next
		}

	}

	return head

}
