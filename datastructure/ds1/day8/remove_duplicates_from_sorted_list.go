package linkedlist

/*
Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

Example 1:
Input: head = [1,1,2]
Output: [1,2]

Example 2:
Input: head = [1,1,2,3,3]
Output: [1,2,3]

*/



/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func DeleteDuplicates(head *ListNode) *ListNode {
    
    return solution1(head)

}

func solution1(head *ListNode) *ListNode {

    if head==nil || head.Next==nil {
        return head
    }

    p1, p2:=head, head.Next
    for p2!=nil {

        if p1.Val==p2.Val {
            p2=p2.Next
            continue
        }

        p1.Next=p2

        p1=p2
        p2=p2.Next
    }

    p1.Next=p2
    return head
}



