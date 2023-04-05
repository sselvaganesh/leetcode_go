package linkedlist

/**

21. Merge Two Sorted Lists
Easy
17.8K
1.6K
Companies
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

	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	dummy := new(ListNode)
	head, l1, l2 := dummy, list1, list2

	for l1 != nil && l2 != nil {

		if l1.Val <= l2.Val {
			dummy.Next = l1
			l1 = l1.Next
		} else if l1.Val > l2.Val {
			dummy.Next = l2
			l2 = l2.Next
		}
		dummy = dummy.Next
	}

	if l1 != nil {
		dummy.Next = l1
	} else if l2 != nil {
		dummy.Next = l2
	}

	return head.Next
}

// 10 10 20 30
// 11
