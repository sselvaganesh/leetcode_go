package tree

/*
102. Binary Tree Level Order Traversal
Medium
12.8K
256
Companies

Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]

Example 2:
Input: root = [1]
Output: [[1]]

Example 3:
Input: root = []
Output: []


Constraints:
The number of nodes in the tree is in the range [0, 2000].
-1000 <= Node.val <= 1000
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

func LevelOrder(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	nodes := []*TreeNode{root}
	var result [][]int

	childNodes, values := eachLevel(nodes)

	for {

		if values != nil {
			result = append(result, values)
		}

		if childNodes == nil {
			break
		}

		childNodes, values = eachLevel(childNodes)

	}

	return result
}

func eachLevel(nodes []*TreeNode) ([]*TreeNode, []int) {

	if nodes == nil {
		return nil, nil
	}

	var childNodes []*TreeNode
	var result []int

	for _, node := range nodes {

		result = append(result, node.Val)

		if node.Left != nil {
			childNodes = append(childNodes, node.Left)
		}

		if node.Right != nil {
			childNodes = append(childNodes, node.Right)
		}

	}

	return childNodes, result
}
