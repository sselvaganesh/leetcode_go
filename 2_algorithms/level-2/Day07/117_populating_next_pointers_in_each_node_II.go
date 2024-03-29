package bfs_dfs

/*
117. Populating Next Right Pointers in Each Node II
Medium
5.2K
296
Companies

Given a binary tree

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.

Initially, all next pointers are set to NULL.

Example 1:
Input: root = [1,2,3,4,5,null,7]
Output: [1,#,2,3,#,4,5,7,#]
Explanation: Given the above binary tree (Figure A), your function should populate each next pointer to point to its next right node, just like in Figure B. The serialized output is in level order as connected by the next pointers, with '#' signifying the end of each level.

Example 2:
Input: root = []
Output: []


Constraints:
The number of nodes in the tree is in the range [0, 6000].
-100 <= Node.val <= 100


Follow-up:
You may only use constant extra space.
The recursive approach is fine. You may assume implicit stack space does not count as extra space for this problem.
*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func Connect(root *Node) *Node {

	if root == nil {
		return nil
	}

	nextLevel := getNextLevelNodes([]*Node{root})
	for nextLevel != nil {
		for i := 0; i < len(nextLevel)-1; i++ {
			nextLevel[i].Next = nextLevel[i+1]
		}
		nextLevel = getNextLevelNodes(nextLevel)
	}

	return root
}

func getNextLevelNodes(nodes []*Node) []*Node {
	var result []*Node
	for _, curNode := range nodes {
		if curNode.Left != nil {
			result = append(result, curNode.Left)
		}

		if curNode.Right != nil {
			result = append(result, curNode.Right)
		}
	}
	return result
}
