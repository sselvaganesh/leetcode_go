package tree

/*
105. Construct Binary Tree from Preorder and Inorder Traversal
Medium
12.7K
372
Companies

Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.

Example 1:
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]

Example 2:
Input: preorder = [-1], inorder = [-1]
Output: [-1]


Constraints:

1 <= preorder.length <= 3000
inorder.length == preorder.length
-3000 <= preorder[i], inorder[i] <= 3000
preorder and inorder consist of unique values.
Each value of inorder also appears in preorder.
preorder is guaranteed to be the preorder traversal of the tree.
inorder is guaranteed to be the inorder traversal of the tree.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func BuildTree(preorder []int, inorder []int) *TreeNode {

	return tree(preorder, inorder)

}

func tree(preorder, inorder []int) *TreeNode {

	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{preorder[0], nil, nil}
	idx := index(inorder, preorder[0])

	root.Left = tree(preorder[1:idx+1], inorder[:idx])
	root.Right = tree(preorder[idx+1:], inorder[idx+1:])

	return root
}

func index(inorder []int, val int) int {

	for i := 0; i < len(inorder); i++ {
		if inorder[i] == val {
			return i
		}
	}
	return -1
}
