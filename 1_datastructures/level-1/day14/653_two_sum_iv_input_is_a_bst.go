package tree

/**
653. Two Sum IV - Input is a BST
Easy
5.8K
240
Companies

Given the root of a binary search tree and an integer k, return true if there exist two elements in the BST such that their sum is equal to k, or false otherwise.

Example 1:
Input: root = [5,3,6,2,4,null,7], k = 9
Output: true

Example 2:
Input: root = [5,3,6,2,4,null,7], k = 28
Output: false

Constraints:
The number of nodes in the tree is in the range [1, 104].
-104 <= Node.val <= 104
root is guaranteed to be a valid binary search tree.
-105 <= k <= 105

 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/

func FindTarget(root *TreeNode, k int) bool {

	m := make(map[int]struct{})
	return isFound(root, k, m)

}

func isFound(node *TreeNode, target int, m map[int]struct{}) bool {

	if node == nil {
		return false
	}

	if _, ok := m[target-node.Val]; ok {
		return true
	}
	m[node.Val] = struct{}{}

	return isFound(node.Left, target, m) || isFound(node.Right, target, m)

}
