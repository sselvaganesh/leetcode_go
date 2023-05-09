package tree

/*
543. Diameter of Binary Tree
Easy
11.2K
703
Companies
Given the root of a binary tree, return the length of the diameter of the tree.

The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

The length of a path between two nodes is represented by the number of edges between them.


Example 1:
Input: root = [1,2,3,4,5]
Output: 3
Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].

Example 2:
Input: root = [1,2]
Output: 1


Constraints:
The number of nodes in the tree is in the range [1, 104].
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

func DiameterOfBinaryTree(root *TreeNode) int {

	return diameter(root)

}

func diameter(curNode *TreeNode) int {

	if curNode == nil {
		return 0
	}

	lHeight := height(curNode.Left)
	rHeight := height(curNode.Right)

	lDiameter := diameter(curNode.Left)
	rDiameter := diameter(curNode.Right)

	return max(lHeight+rHeight, max(lDiameter, rDiameter))

}

func height(curNode *TreeNode) int {

	if curNode == nil {
		return 0
	}

	return 1 + max(height(curNode.Left), height(curNode.Right))
}

func max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

var maxDiameter int

func diameterOfBinaryTreeRec(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ld := diameterOfBinaryTreeRec(root.Left)
	rd := diameterOfBinaryTreeRec(root.Right)

	d := ld + rd
	if d > maxDiameter {
		maxDiameter = d
	}

	return 1 + max(ld, rd)
}
