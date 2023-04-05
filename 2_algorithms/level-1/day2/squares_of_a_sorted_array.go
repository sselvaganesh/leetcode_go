package twopointers

/*
Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.

Example 1:
Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].

Example 2:
Input: nums = [-7,-3,2,3,11]
Output: [4,9,9,49,121]
*/

import (
	"sort"
)

func SortedSquares(nums []int) []int {

	//return bruteForce(nums)
	return twoPointers(nums)

}

func twoPointers(nums []int) []int {

	if len(nums) == 0 {
		return nil
	}

	result := make([]int, len(nums))
	for left, right, idx := 0, len(nums)-1, len(nums)-1; right >= left; {

		lVal := nums[left] * nums[left]
		rVal := nums[right] * nums[right]

		if lVal >= rVal {
			result[idx] = lVal
			left++
		} else if rVal > lVal {
			result[idx] = rVal
			right--
		}

		idx--
		if idx < 0 {
			break
		}

	}

	return result

}

func bruteForce(nums []int) []int {

	if len(nums) == 0 {
		return nil
	}

	var result []int
	for i := 0; i < len(nums); i++ {
		result = append(result, nums[i]*nums[i])
	}

	sort.Ints(result)
	return result

}
