package array

/*
Given an integer array nums, find the
subarray with the largest sum, and return its sum.

Example 1:
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.

Example 2:
Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.

Example 3:
Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
*/

func MaxSubArray(nums []int) int {

	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	curSum, result := nums[0], nums[0]
	for _, val := range nums[1:] {

		if curSum < 0 {
			curSum = 0
		}

		curSum += val
		result = max(result, curSum)
	}

	return result
}

func max(val1, val2 int) int {

	if val1 > val2 {
		return val1
	}

	return val2
}

func Solution2(nums []int) int {
	max := nums[0]
	curMax := nums[0]
	for _, num := range nums[1:] {
		if max > 0 {
			max = max + num
		} else {
			max = num
		}
		if max > curMax {
			curMax = max
		}
	}
	return curMax
}
