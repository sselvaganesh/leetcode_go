package tree

/*
230. Kth Smallest Element in a BST
Medium
9.6K
174
Companies
Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.

Example 1:
Input: root = [3,1,4,null,2], k = 1
Output: 1

Example 2:
Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3


Constraints:
The number of nodes in the tree is n.
1 <= k <= n <= 104
0 <= Node.val <= 104
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func KthSmallest(root *TreeNode, k int) int {

	elem := make([]int, 2)
	getNum(root, k, &elem)
	return elem[1]

}

func getNum(curNode *TreeNode, k int, elem *[]int) {

	if curNode == nil || (*elem)[0] == k {
		return
	}

	getNum(curNode.Left, k, elem)
	(*elem)[0]++
	if (*elem)[0] == k {
		(*elem)[1] = curNode.Val
		return
	} else if (*elem)[0] >= k {
		return
	}

	getNum(curNode.Right, k, elem)

}
