package tree

/*
113. Path Sum II
Medium
7K
139
Companies

Given the root of a binary tree and an integer targetSum, return all root-to-leaf paths where the sum of the node values in the path equals targetSum. Each path should be returned as a list of the node values, not node references.

A root-to-leaf path is a path starting from the root and ending at any leaf node. A leaf is a node with no children.

Example 1:
Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
Output: [[5,4,11,2],[5,8,4,5]]
Explanation: There are two paths whose sum equals targetSum:
5 + 4 + 11 + 2 = 22
5 + 8 + 4 + 5 = 22

Example 2:
Input: root = [1,2,3], targetSum = 5
Output: []

Example 3:
Input: root = [1,2], targetSum = 0
Output: []

Constraints:
The number of nodes in the tree is in the range [0, 5000].
-1000 <= Node.val <= 1000
-1000 <= targetSum <= 1000

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func PathSum(root *TreeNode, targetSum int) [][]int {

	if root == nil {
		return nil
	}

	var result [][]int
	helper(root, targetSum, 0, []int{}, &result)
	return result

}

func helper(curNode *TreeNode, target, curSum int, path []int, result *[][]int) {

	path = append(path, curNode.Val)
	curSum += curNode.Val

	if curNode.Right == nil && curNode.Left == nil {
		if curSum == target {
			*result = append(*result, append([]int{}, path...))
		}
		return
	}

	if curNode.Left != nil {
		helper(curNode.Left, target, curSum, path, result)
	}

	if curNode.Right != nil {
		helper(curNode.Right, target, curSum, path, result)
	}

	curSum -= curNode.Val
	path = path[:len(path)-1]

	return
}
