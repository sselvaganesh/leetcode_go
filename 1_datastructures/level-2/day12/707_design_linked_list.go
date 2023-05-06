package linkedlist

/*
707. Design Linked List
Medium
2.2K
1.4K
Companies
Design your implementation of the linked list. You can choose to use a singly or doubly linked list.
A node in a singly linked list should have two attributes: val and next. val is the value of the current node, and next is a pointer/reference to the next node.
If you want to use the doubly linked list, you will need one more attribute prev to indicate the previous node in the linked list. Assume all nodes in the linked list are 0-indexed.

Implement the MyLinkedList class:

MyLinkedList() Initializes the MyLinkedList object.
int get(int index) Get the value of the indexth node in the linked list. If the index is invalid, return -1.
void addAtHead(int val) Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
void addAtTail(int val) Append a node of value val as the last element of the linked list.
void addAtIndex(int index, int val) Add a node of value val before the indexth node in the linked list. If index equals the length of the linked list, the node will be appended to the end of the linked list. If index is greater than the length, the node will not be inserted.
void deleteAtIndex(int index) Delete the indexth node in the linked list, if the index is valid.


Example 1:

Input
["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"]
[[], [1], [3], [1, 2], [1], [1], [1]]
Output
[null, null, null, null, 2, null, 3]

Explanation
MyLinkedList myLinkedList = new MyLinkedList();
myLinkedList.addAtHead(1);
myLinkedList.addAtTail(3);
myLinkedList.addAtIndex(1, 2);    // linked list becomes 1->2->3
myLinkedList.get(1);              // return 2
myLinkedList.deleteAtIndex(1);    // now the linked list is 1->3
myLinkedList.get(1);              // return 3


Constraints:

0 <= index, val <= 1000
Please do not use the built-in LinkedList library.
At most 2000 calls will be made to get, addAtHead, addAtTail, addAtIndex and deleteAtIndex.

*/

type MyLinkedList struct {
	head *Node
}

type Node struct {
	Val  int
	Next *Node
}

func Constructor() MyLinkedList {

	return MyLinkedList{}

}

func getNewNode(val int) *Node {
	return &Node{val, nil}
}

func (this *MyLinkedList) Get(index int) int {

	curNode := this.head
	idx := 0

	for curNode != nil {

		if idx == index {
			return curNode.Val
		}

		curNode = curNode.Next
		idx++
	}

	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {

	newNode := getNewNode(val)

	if this.head == nil {
		this.head = newNode
		return
	}

	newNode.Next = this.head
	this.head = newNode

}

func (this *MyLinkedList) AddAtTail(val int) {

	newNode := getNewNode(val)
	if this.head == nil {
		this.head = newNode
		return
	}

	lastNode := this.head
	for ; lastNode.Next != nil; lastNode = lastNode.Next {
	}
	lastNode.Next = newNode
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {

	if index == 0 {
		this.AddAtHead(val)
	}

	if this.head == nil {
		return
	}

	idx := 0
	curNode, nextNode := this.head, this.head.Next
	for nextNode != nil {
		if idx == (index - 1) {
			break
		}
		curNode, nextNode = nextNode, nextNode.Next
		idx++
	}

	if idx != (index - 1) {
		return
	}

	newNode := getNewNode(val)
	newNode.Next = nextNode
	curNode.Next = newNode

}

func (this *MyLinkedList) DeleteAtIndex(index int) {

	if this.head == nil {
		return
	}

	if index == 0 {
		this.head = this.head.Next
		return
	}

	idx := 0
	var prevNode *Node
	curNode := this.head

	for curNode != nil {

		if idx == index {
			break
		}

		prevNode, curNode = curNode, curNode.Next
		idx++
	}

	if idx != index {
		return
	}

	if prevNode == nil {
		this.head = this.head.Next
	} else if curNode == nil {
		prevNode.Next = nil
	} else {
		prevNode.Next = curNode.Next
	}

}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
