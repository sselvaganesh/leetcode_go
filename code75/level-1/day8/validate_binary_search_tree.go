package binarysearchtree

/**

Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left
subtree
 of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.


Example 1:
Input: root = [2,1,3]
Output: true


Example 2:
Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.


 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {

	return check(root, math.MinInt, math.MaxInt)

}

func check(root *TreeNode, min, max int) bool {

	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	if root.Left != nil && root.Left.Val >= root.Val {
		return false
	}

	if root.Right != nil && root.Right.Val <= root.Val {
		return false
	}

	return check(root.Left, min, root.Val) && check(root.Right, root.Val, max)

}
