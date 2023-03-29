package twopointer

/*
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
Note that you must do this in-place without making a copy of the array.

Example 1:
Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

Example 2:
Input: nums = [0]
Output: [0]
*/

func MoveZeroes(nums []int) {

	//solution1(nums)
	solution2(nums)

}

func solution2(nums []int) {

	for zeroPos, idx := 0, 0; idx < len(nums); idx++ {

		if nums[idx] != 0 {
			nums[zeroPos], nums[idx] = nums[idx], nums[zeroPos]
			zeroPos++
		}

	}

	return

}

func solution1(nums []int) {

	if len(nums) <= 1 {
		return
	}

	zeroPos, idx := 0, 0

	for idx < len(nums) && zeroPos < len(nums) {

		for idx < len(nums) {
			if nums[idx] == 0 {
				idx++
			} else {
				break
			}
		}

		for zeroPos < len(nums) {
			if nums[zeroPos] != 0 {
				zeroPos++
			} else {
				break
			}
		}

		if zeroPos < len(nums) && idx < len(nums) && idx > zeroPos {
			nums[zeroPos] = nums[idx]
			nums[idx] = 0
			idx++
		} else {
			idx++
		}

	}

	return

}
