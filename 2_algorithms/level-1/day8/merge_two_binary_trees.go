package bfsdfs

/**

You are given two binary trees root1 and root2.

Imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not. You need to merge the two trees into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of the new tree.

Return the merged tree.

Note: The merging process must start from the root nodes of both trees.


Example 1:
Input: root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
Output: [3,4,5,5,4,null,7]

Example 2:
Input: root1 = [1], root2 = [1,2]
Output: [2,2]

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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	//return solution1(root1, root2)  //New binary tree
	return solution2(root1, root2) //In-Place solution

}

func solution2(root1, root2 *TreeNode) *TreeNode {

	if root1 == nil && root2 == nil {
		return nil
	} else if root1 == nil && root2 != nil {
		return root2
	} else if root1 != nil && root2 == nil {
		return root1
	}

	root1.Val = root1.Val + root2.Val

	root1.Left = solution2(root1.Left, root2.Left)
	root1.Right = solution2(root1.Right, root2.Right)

	return root1
}

func solution1(root1, root2 *TreeNode) *TreeNode {

	if root1 == nil && root2 == nil {
		return nil
	} else if root1 == nil && root2 != nil {
		return root2
	} else if root1 != nil && root2 == nil {
		return root1
	}

	root := new(TreeNode)
	root.Val = root1.Val + root2.Val

	root.Left = solution1(root1.Left, root2.Left)
	root.Right = solution1(root1.Right, root2.Right)

	return root
}
