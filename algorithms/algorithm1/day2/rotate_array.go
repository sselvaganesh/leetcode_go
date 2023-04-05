package twopointers

/*
Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

Example 1:
Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]

Example 2:
Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]
*/

func Rotate(nums []int, k int) {

	//solution1(nums, k)

	solution2(nums, k)

}

func solution2(nums []int, k int) {

	if len(nums) <= 1 || k%len(nums) == 0 {
		return
	}

	n := k % len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, n-1)
	reverse(nums, n, len(nums)-1)

	return
}

func reverse(nums []int, start, end int) {

	for end > start {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func solution1(nums []int, k int) {

	if len(nums) <= 1 || k%len(nums) == 0 {
		return
	}

	n := len(nums) - (k % len(nums))
	var result []int

	result = append(result, nums[n:]...)
	result = append(result, nums[:n]...)
	copy(nums, result)

	return
}
