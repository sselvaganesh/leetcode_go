package linkedlist

/**
You are given the heads of two sorted linked lists list1 and list2.
Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.
Return the head of the merged linked list.

Example 1:
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

Example 2:
Input: list1 = [], list2 = []
Output: []

Example 3:
Input: list1 = [], list2 = [0]
Output: [0]

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

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	var list1Node = list1
	var list2Node = list2
	if list1.Val > list2.Val {
		list2Node = list1
		list1Node = list2
	}
	for list1Node.Next != nil && list2Node != nil {
		if list1Node.Next.Val < list2Node.Val {
			list1Node = list1Node.Next
		} else {
			var temp = list2Node.Next
			list2Node.Next = list1Node.Next
			list1Node.Next = list2Node
			list2Node = temp
			list1Node = list1Node.Next
		}
	}
	for list2Node != nil {
		list1Node.Next = list2Node
		list1Node = list1Node.Next
		list2Node = list2Node.Next
	}
	if list1.Val > list2.Val {
		return list2
	}
	return list1
}
