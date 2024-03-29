package linked_list

/*
142. Linked List Cycle II
Medium
11.5K
844
Companies
Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to (0-indexed). It is -1 if there is no cycle. Note that pos is not passed as a parameter.

Do not modify the linked list.

Example 1:
Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1
Explanation: There is a cycle in the linked list, where tail connects to the second node.

Example 2:
Input: head = [1,2], pos = 0
Output: tail connects to node index 0
Explanation: There is a cycle in the linked list, where tail connects to the first node.

Example 3:
Input: head = [1], pos = -1
Output: no cycle
Explanation: There is no cycle in the linked list.


Constraints:
The number of the nodes in the list is in the range [0, 104].
-105 <= Node.val <= 105
pos is -1 or a valid index in the linked-list.

Follow up: Can you solve it using O(1) (i.e. constant) memory?
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func DetectCycle(head *ListNode) *ListNode {

	//return usingHashMap(head)

	return usingTwoPointers(head)

}

func usingTwoPointers(head *ListNode) *ListNode {

	slow, fast := head, head

	for fast != nil && fast.Next != nil {

		slow, fast = slow.Next, fast.Next.Next

		if slow == fast {
			newSlow := head
			for newSlow != fast {
				newSlow, fast = newSlow.Next, fast.Next
			}
			return fast
		}

	}

	return nil
}

func usingHashMap(head *ListNode) *ListNode {

	m := make(map[*ListNode]struct{})
	curNode := head

	for curNode != nil {
		if _, ok := m[curNode]; ok {
			return curNode
		} else {
			m[curNode] = struct{}{}
		}

		curNode = curNode.Next
	}

	return nil
}
