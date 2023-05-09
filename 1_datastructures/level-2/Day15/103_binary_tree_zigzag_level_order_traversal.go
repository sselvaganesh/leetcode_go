package tree

/*
103. Binary Tree Zigzag Level Order Traversal
Medium
9.2K
241
Companies
Given the root of a binary tree, return the zigzag level order traversal of its nodes' values. (i.e., from left to right, then right to left for the next level and alternate between).

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[20,9],[15,7]]

Example 2:
Input: root = [1]
Output: [[1]]

Example 3:
Input: root = []
Output: []


Constraints:
The number of nodes in the tree is in the range [0, 2000].
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

func ZigzagLevelOrder(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	var res [][]int
	nodes := []*TreeNode{root}
	left := true

	for len(nodes) > 0 {

		if left {
			res = append(res, getNodeVal(nodes, true))
		} else {
			res = append(res, getNodeVal(nodes, false))
		}

		left = !left
		nodes = getNextLevel(nodes)
	}
	return res
}

func getNextLevel(nodes []*TreeNode) []*TreeNode {

	var res []*TreeNode
	for i := 0; i < len(nodes); i++ {
		if nodes[i].Left != nil {
			res = append(res, nodes[i].Left)
		}
		if nodes[i].Right != nil {
			res = append(res, nodes[i].Right)
		}
	}

	return res
}

func getNodeVal(nodes []*TreeNode, fwd bool) []int {

	var res []int
	if fwd {
		for i := 0; i < len(nodes); i++ {
			res = append(res, nodes[i].Val)
		}
	} else {
		for i := len(nodes) - 1; i >= 0; i-- {
			res = append(res, nodes[i].Val)
		}
	}

	return res
}
