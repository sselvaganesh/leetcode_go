package binary_search

/*
34. Find First and Last Position of Element in Sorted Array
Medium
16.5K
393
Companies

Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.

Example 1:
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]

Example 2:
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]

Example 3:
Input: nums = [], target = 0
Output: [-1,-1]

Constraints:
0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums is a non-decreasing array.
-109 <= target <= 109
*/

func SearchRange(nums []int, target int) []int {

	//return bruteForceSearchRange(nums, target)

	//return binarySearch1(nums, target)      //Worst cast is O(n)

	return binarySearch2(nums, target)

}

func binarySearch2(nums []int, target int) []int {

	return []int{binarySearchWithFlag(nums, target, true), binarySearchWithFlag(nums, target, false)}

}

func binarySearch1(nums []int, target int) []int {

	lower := getRelatedIndex(nums, target)
	if lower == -1 {
		return []int{-1, -1}
	}
	upper := lower

	for lower-1 >= 0 && nums[lower-1] == target {
		lower--
	}

	for upper+1 < len(nums) && nums[upper+1] == target {
		upper++
	}

	return []int{lower, upper}
}

func bruteForceSearchRange(nums []int, target int) []int {

	start, end := -1, -1
	for i := 0; i < len(nums); i++ {

		if start == -1 && nums[i] == target {
			start = i
		}

		if start != -1 && nums[i] == target {
			end = i
		}

	}

	return []int{start, end}
}

func binarySearchWithFlag(nums []int, target int, left bool) int {

	low, high := 0, len(nums)-1
	res := -1

	for high >= low {
		mid := (low + high) / 2
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			res = mid
			if left {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}

	return res

}

func getRelatedIndex(nums []int, target int) int {

	low, mid, high := 0, 0, len(nums)-1
	for high >= low {

		mid = (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}

	return -1
}
