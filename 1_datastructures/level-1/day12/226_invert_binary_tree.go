package tree

/*
226. Invert Binary Tree
Easy
11.9K
169
Companies

Given the root of a binary tree, invert the tree, and return its root.

 Example 1:
Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]

Example 2:
Input: root = [2,1,3]
Output: [2,3,1]

Example 3:
Input: root = []
Output: []


Constraints:
The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InvertTree(root *TreeNode) *TreeNode {

	invert(root)
	return root
}

func invert(node *TreeNode) {

	if node == nil {
		return
	}

	node.Left, node.Right = node.Right, node.Left
	invert(node.Left)
	invert(node.Right)

}
