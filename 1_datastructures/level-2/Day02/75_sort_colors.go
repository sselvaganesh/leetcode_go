package array

/*
75. Sort Colors
Medium
14.6K
519
Companies

Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function.


Example 1:
Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]

Example 2:
Input: nums = [2,0,1]
Output: [0,1,2]


Constraints:
n == nums.length
1 <= n <= 300
nums[i] is either 0, 1, or 2.

Follow up: Could you come up with a one-pass algorithm using only constant extra space?

*/

import (
	"sort"
)

func SortColors(nums []int) {

	red, blue := 0, len(nums)-1
	for idx := 0; blue >= idx; {

		//fmt.Println(nums)
		if nums[idx] == 0 {
			nums[idx], nums[red] = nums[red], nums[idx]
			red++
			idx++
		} else if nums[idx] == 2 {
			nums[idx], nums[blue] = nums[blue], nums[idx]
			blue--
		} else {
			idx++
		}

	}

	return

}

func InPlaceSolution(nums []int) {

	const (
		red, white, blue = 0, 1, 2
	)

	redPos, bluePos := 0, len(nums)-1

	for i := 0; i <= bluePos; {
		if nums[i] == red {
			nums[i], nums[redPos] = nums[redPos], nums[i]
			redPos++
			i++
		} else if nums[i] == blue {
			nums[i], nums[bluePos] = nums[bluePos], nums[i]
			bluePos--
		} else {
			i++
		}
	}

}

func UsingLibFunc(nums []int) {
	sort.Ints(nums)
}
