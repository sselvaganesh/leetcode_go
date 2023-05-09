package tree

/*
235. Lowest Common Ancestor of a Binary Search Tree
Medium
9.2K
256
Companies

Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

Example 1:
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6
Explanation: The LCA of nodes 2 and 8 is 6.

Example 2:
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2
Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.

Example 3:
Input: root = [2,1], p = 2, q = 1
Output: 2

Constraints:
The number of nodes in the tree is in the range [2, 105].
-109 <= Node.val <= 109
All Node.val are unique.
p != q
p and q will exist in the BST.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	low, high := minMax(p.Val, q.Val)

	return solution(root, low, high)

}

func solution(node *TreeNode, low, high int) *TreeNode {

	if node == nil {
		return nil
	}

	if node.Val == low || node.Val == high || node.Val > low && node.Val < high {
		return node
	}

	if node.Val > high {
		return solution(node.Left, low, high)
	}

	return solution(node.Right, low, high)
}

func minMax(val1, val2 int) (int, int) {
	if val1 < val2 {
		return val1, val2
	}
	return val2, val1
}
