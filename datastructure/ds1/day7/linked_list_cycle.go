package linkedlist

/**
Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.

Example 1:
Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).

Example 2:
Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.

Example 3:
Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.


	1	2
	2	4
	3	3
	4
	5
	3




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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func HasCycle(head *ListNode) bool {

	//return usingHashMap(head)

	return usingTwoPointers(head)

}

func usingTwoPointers(head *ListNode) bool {

	if head == nil {
		return false
	}

	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil && fast.Next.Next != nil {

		if slow == fast {
			return true
		}

		slow = slow.Next
		fast = fast.Next.Next

	}

	return false

}

func usingHashMap(head *ListNode) bool {

	if head == nil {
		return false
	}

	m := make(map[*ListNode]struct{})
	curNode := head

	for curNode != nil {

		if _, ok := m[curNode]; ok {
			return true
		} else {
			m[curNode] = struct{}{}
		}

		curNode = curNode.Next
	}

	return false
}
