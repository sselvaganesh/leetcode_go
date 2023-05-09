package array

/*
560. Subarray Sum Equals K
Medium
17.8K
519
Companies

Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.

A subarray is a contiguous non-empty sequence of elements within an array.

Example 1:
Input: nums = [1,1,1], k = 2
Output: 2

Example 2:
Input: nums = [1,2,3], k = 3
Output: 2


Constraints:
1 <= nums.length <= 2 * 104
-1000 <= nums[i] <= 1000
-107 <= k <= 107
*/

func SubarraySum(nums []int, k int) int {

	m := make(map[int]int)
	m[0] = 1
	res, sum := 0, 0

	for i := 0; i < len(nums); i++ {

		sum += nums[i]
		if val, ok := m[sum-k]; ok {
			res += val
		}

		if _, ok := m[sum]; !ok {
			m[sum] = 1
		} else {
			m[sum]++
		}

	}

	return res
}
