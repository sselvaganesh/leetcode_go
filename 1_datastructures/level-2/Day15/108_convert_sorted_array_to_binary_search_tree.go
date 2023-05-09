package tree

/*
108. Convert Sorted Array to Binary Search Tree
Easy
9.5K
474
Companies
Given an integer array nums where the elements are sorted in ascending order, convert it to a
height-balanced
 binary search tree.


Example 1:
Input: nums = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Explanation: [0,-10,5,null,-3,null,9] is also accepted:

Example 2:
Input: nums = [1,3]
Output: [3,1]
Explanation: [1,null,3] and [3,1] are both height-balanced BSTs.


Constraints:
1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums is sorted in a strictly increasing order.

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

func SortedArrayToBST(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}

	return getTree(nums, 0, len(nums)-1)
}

func getTree(nums []int, low, high int) *TreeNode {

	if low > high {
		return nil
	}

	mid := (low + high) / 2
	root := &TreeNode{nums[mid], nil, nil}
	root.Left = getTree(nums, low, mid-1)
	root.Right = getTree(nums, mid+1, high)

	return root
}
