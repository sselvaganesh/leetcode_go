package binarysearchtree

/**

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



 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
*/

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if p.Val < q.Val {
		return solution(root, p, q)
	}

	return solution(root, q, p)

}

func solution(curNode, low, high *TreeNode) *TreeNode {

	if curNode == nil {
		return nil
	}

	if curNode.Val == low.Val || curNode.Val == high.Val || (curNode.Val > low.Val && curNode.Val < high.Val) {
		return curNode
	}

	if low.Val > curNode.Val {
		return solution(curNode.Right, low, high)
	}

	return solution(curNode.Left, low, high)

}
