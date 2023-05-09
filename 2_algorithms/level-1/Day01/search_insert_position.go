package binarysearch

/*
Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
You must write an algorithm with O(log n) runtime complexity.

Example 1:
Input: nums = [1,3,5,6], target = 5
Output: 2

Example 2:
Input: nums = [1,3,5,6], target = 2
Output: 1

Example 3:
Input: nums = [1,3,5,6], target = 7
Output: 4
*/

func SearchInsert(nums []int, target int) int {

	start := 0
	end := len(nums) - 1
	mid := (start + end) / 2

	for start != end && end > start {

		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}

		mid = (start + end) / 2
	}

	if target > nums[mid] {
		return mid + 1
	} else if mid == 0 {
		return 0
	} else {
		return mid
	}

}
