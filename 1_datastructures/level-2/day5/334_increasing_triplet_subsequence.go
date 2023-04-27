package array

/*
334. Increasing Triplet Subsequence
Medium
6.4K
277
Companies

Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.

Example 1:
Input: nums = [1,2,3,4,5]
Output: true
Explanation: Any triplet where i < j < k is valid.

Example 2:
Input: nums = [5,4,3,2,1]
Output: false
Explanation: No triplet exists.

Example 3:
Input: nums = [2,1,5,0,4,6]
Output: true
Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4] == 4 < nums[5] == 6.


Constraints:
1 <= nums.length <= 5 * 105
-231 <= nums[i] <= 231 - 1
*/

import (
	"math"
)

func IncreasingTriplet(nums []int) bool {

	var first, second bool
	low, mid := math.MaxInt, math.MinInt

	for _, num := range nums {

		if first && second {
			break
		}

		if low > num {
			low = num
			continue
		}

		if !first && num > low {
			first = true
			mid = num
			continue
		}

		if first && num > mid {
			second = true
			break
		}

		if num > low && num < mid {
			mid = num
		}

	}

	return first && second

}

func increasingTripletSolution(nums []int) bool {

	first, second := false, false
	min, max := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {

		if first && second {
			break
		}

		if min > nums[i] {
			min = nums[i]
			continue
		}

		if !first && nums[i] > min {
			first = true
			max = nums[i]
			continue
		}

		if first && nums[i] > max {
			second = true
		}

		if nums[i] > min && nums[i] < max {
			max = nums[i]
		}

	}

	return first && second

}
