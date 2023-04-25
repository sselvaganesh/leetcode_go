package linked_list

/**
19. Remove Nth Node From End of List
Medium
15.4K
642
Companies
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


Constraints:
The number of nodes in the list is sz.
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz

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

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {

	if head == nil {
		return nil
	}

	prev, curNode := &ListNode{}, head
	counter := 1

	for curNode != nil {

		if counter == n+1 {
			prev = head
		} else if prev != nil {
			prev = prev.Next
		}

		curNode = curNode.Next
		counter++
	}

	if prev == nil {
		head = head.Next
	} else if prev.Next.Next == nil {
		prev.Next = nil
	} else {
		prev.Next = prev.Next.Next
	}

	return head

}


func RemoveNthFromEndSolution(head *ListNode, n int) *ListNode {

    if head==nil {
        return nil
    }

    curNode, nextNode := &ListNode{}, head
    counter:=0

    for nextNode!=nil {

        if curNode!=nil {
            curNode=curNode.Next
        }

        nextNode=nextNode.Next
        counter++

        if counter==n+1 {
            curNode=head
        }

    }

    if curNode==nil {
        head=head.Next
    } else if curNode.Next.Next==nil {
        curNode.Next=nil
    } else {
        curNode.Next=curNode.Next.Next
    }


    return head

}

