package array

/*
238. Product of Array Except Self
Medium
17.1K
942
Companies
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.



Example 1:
Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Example 2:
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]

Constraints:
2 <= nums.length <= 105
-30 <= nums[i] <= 30
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
*/

func ProductExceptSelf(nums []int) []int {

	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = 1
	}

	pfix := 1
	for i := 0; i < len(nums); i++ {
		result[i] = pfix
		pfix = pfix * nums[i]
	}

	pfix = 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = pfix * result[i]
		pfix = nums[i] * pfix
	}

	return result

}

func productExceptSelfSolution(nums []int) []int {

	result := make([]int, len(nums))

	prefix := 1
	for i := 0; i < len(nums); i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	prefix = 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= prefix
		prefix *= nums[i]
	}

	return result
}
