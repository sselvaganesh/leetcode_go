package tree

/*
144. Binary Tree Preorder Traversal
Easy
6.6K
174
Companies

Given the root of a binary tree, return the preorder traversal of its nodes' values.


Example 1:
Input: root = [1,null,2,3]
Output: [1,2,3]

Example 2:
Input: root = []
Output: []

Example 3:
Input: root = [1]
Output: [1]


Constraints:
The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100

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

func PreorderTraversal(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	var result []int
	preOrder(root, &result)
	return result

}

func preOrder(node *TreeNode, result *[]int) {

	if node == nil {
		return
	}

	*result = append(*result, node.Val)
	preOrder(node.Left, result)
	preOrder(node.Right, result)
}
