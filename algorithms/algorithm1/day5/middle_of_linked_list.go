package linkedlist

/**

Given the head of a singly linked list, return the middle node of the linked list.
If there are two middle nodes, return the second middle node.

Input: head = [1,2,3,4,5]
Output: [3,4,5]
Explanation: The middle node of the list is node 3.

Input: head = [1,2,3,4,5,6]
Output: [4,5,6]
Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.


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

func MiddleNode(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	curNode, middleNode := head, head
	i := 1

	for curNode != nil {
		curNode = curNode.Next
		if i%2 == 0 {
			middleNode = middleNode.Next
		}
		i++
	}

	return middleNode

}


