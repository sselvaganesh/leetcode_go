package linkedlist

/*
Given the head of a singly linked list, reverse the list, and return the reversed list.

Example 1:
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

Example 2:
Input: head = [1,2]
Output: [2,1]

Example 3:
Input: head = []
Output: []

*/


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */ 

type ListNode struct {
     Val int
     Next *ListNode
}
 
func ReverseList(head *ListNode) *ListNode {
    
    return inPlace(head)
}


func inPlace(head *ListNode) *ListNode {

    if head==nil || head.Next == nil {
        return head
    }

    var prev *ListNode
    curNode:=head

    for curNode!=nil {
        next:=curNode.Next
        curNode.Next=prev
        prev=curNode
        curNode=next
    }

    return prev

}


