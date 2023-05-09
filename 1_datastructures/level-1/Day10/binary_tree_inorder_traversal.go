package tree

/*

94. Binary Tree Inorder Traversal
Easy
11.2K
549
Companies
Given the root of a binary tree, return the inorder traversal of its nodes' values.


Example 1:
Input: root = [1,null,2,3]
Output: [1,3,2]

Example 2:
Input: root = []
Output: []

Example 3:
Input: root = [1]
Output: [1]


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

func InorderTraversal(root *TreeNode) []int {

	result := make([]int, 0)
	inOrder(root, &result)

	return result
}

func inOrder(node *TreeNode, result *[]int) {

	if node == nil {
		return
	}

	inOrder(node.Left, result)
	*result = append(*result, node.Val)
	inOrder(node.Right, result)

}
