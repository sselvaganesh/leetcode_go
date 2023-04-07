package tree

/*
101. Symmetric Tree
Easy
13.2K
296
Companies

Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

Example 1:
Input: root = [1,2,2,3,4,4,3]
Output: true

Example 2:
Input: root = [1,2,2,null,3,null,3]
Output: false


Constraints:
The number of nodes in the tree is in the range [1, 1000].
-100 <= Node.val <= 100

*/

func IsSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	return isOk(root.Left, root.Right)

}

func isOk(left, right *TreeNode) bool {

	if left == nil && right == nil {
		return true
	}

	if (left == nil && right != nil) || (left != nil && right == nil) || left.Val != right.Val {
		return false
	}

	return isOk(left.Left, right.Right) && isOk(left.Right, right.Left)

}
