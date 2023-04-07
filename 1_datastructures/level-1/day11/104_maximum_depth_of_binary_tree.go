package tree

/*
104. Maximum Depth of Binary Tree
Easy
10.5K
167
Companies

Given the root of a binary tree, return its maximum depth.
A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.


Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: 3

Example 2:
Input: root = [1,null,2]
Output: 2

Constraints:
The number of nodes in the tree is in the range [0, 104].
-100 <= Node.val <= 100
*/

func MaxDepth(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, depth int) int {

	if node == nil {
		return depth
	}

	return max(dfs(node.Left, depth+1), dfs(node.Right, depth+1))

}

func max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}
