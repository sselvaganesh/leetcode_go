package linkedlist

/**

Given the head of a linked list, remove the nth node from the end of the list and return its head.

Example 1:
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

Example 2:
Input: head = [1], n = 1
Output: []

Example 3:
Input: head = [1,2], n = 1
Output: [1]


 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {

    if head==nil || n<1{
        return head
    }


    curNode, prevNode := head, head
    counter:=1

    for curNode!=nil {

        if counter==(n+1) {
            prevNode=head
        } else if prevNode!=nil {
            prevNode=prevNode.Next
        } 

        curNode=curNode.Next
        counter++
    }

    if prevNode==nil{
        return head.Next
    }

    if prevNode.Next.Next==nil {
        prevNode.Next=nil
    } else {
        prevNode.Next=prevNode.Next.Next
    }

    return head

}


