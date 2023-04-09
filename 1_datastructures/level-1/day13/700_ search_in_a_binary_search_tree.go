package tree

/**

700. Search in a Binary Search Tree
Easy
4.8K
164
Companies

You are given the root of a binary search tree (BST) and an integer val.

Find the node in the BST that the node's value equals val and return the subtree rooted with that node. If such a node does not exist, return null.


Example 1:
Input: root = [4,2,7,1,3], val = 2
Output: [2,1,3]

Example 2:
Input: root = [4,2,7,1,3], val = 5
Output: []

Constraints:
The number of nodes in the tree is in the range [1, 5000].
1 <= Node.val <= 107
root is a binary search tree.
1 <= val <= 107
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

func SearchBST(root *TreeNode, val int) *TreeNode {

	if root == nil {
		return nil
	} else if root.Val == val {
		return root
	}

	if left := SearchBST(root.Left, val); left != nil {
		return left
	}

	return SearchBST(root.Right, val)

}
