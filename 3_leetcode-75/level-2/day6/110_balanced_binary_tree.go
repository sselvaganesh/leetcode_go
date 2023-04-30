package tree

/*
110. Balanced Binary Tree
Easy
8.8K
494
Companies
Given a binary tree, determine if it is
height-balanced

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: true

Example 2:
Input: root = [1,2,2,3,3,null,null,4,4]
Output: false

Example 3:
Input: root = []
Output: true


Constraints:
The number of nodes in the tree is in the range [0, 5000].
-104 <= Node.val <= 104
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func IsBalanced(root *TreeNode) bool {

	return isHeightBalanced(root)

}

func isHeightBalanced(curNode *TreeNode) bool {

	if curNode == nil {
		return true
	}

	if abs(maxDepth(curNode.Left), maxDepth(curNode.Right)) > 1 {
		return false
	}

	return isHeightBalanced(curNode.Left) && isHeightBalanced(curNode.Right)
}

func maxDepth(curNode *TreeNode) int {

	if curNode == nil {
		return 0
	}

	return 1 + max(maxDepth(curNode.Left), maxDepth(curNode.Right))
}

func max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

func abs(val1, val2 int) int {
	res := val1 - val2
	if res < 0 {
		res = -res
	}
	return res
}
