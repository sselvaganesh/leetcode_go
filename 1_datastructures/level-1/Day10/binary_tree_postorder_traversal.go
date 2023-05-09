package tree

/*
145. Binary Tree Postorder Traversal
Easy
5.8K
170
Companies

Given the root of a binary tree, return the postorder traversal of its nodes' values.

Example 1:
Input: root = [1,null,2,3]
Output: [3,2,1]

Example 2:
Input: root = []
Output: []

Example 3:
Input: root = [1]
Output: [1]

Constraints:
The number of the nodes in the tree is in the range [0, 100].
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
func PostorderTraversal(root *TreeNode) []int {

	result := make([]int, 0)
	postOrder(root, &result)
	return result
}

func postOrder(node *TreeNode, result *[]int) {

	if node == nil {
		return
	}

	postOrder(node.Left, result)
	postOrder(node.Right, result)
	*result = append(*result, node.Val)

}
