package array

/*
169. Majority Element
Easy
14.3K
439
Companies
Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

Example 1:
Input: nums = [3,2,3]
Output: 3

Example 2:
Input: nums = [2,2,1,1,1,2,2]
Output: 2


Constraints:
n == nums.length
1 <= n <= 5 * 104
-109 <= nums[i] <= 109

Follow-up: Could you solve the problem in linear time and in O(1) space?
*/

func MajorityElement(nums []int) int {

	//return solution1(nums)
	return solution2(nums)

}

func solution2(nums []int) int {

	res := nums[0]
	cnt := 1

	for _, num := range nums[1:] {
		if num == res {
			cnt++
		} else {
			cnt--
		}

		if cnt == 0 {
			res = num
			cnt = 1
		}
	}

	return res
}

func solution1(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	doesntExist := -1
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {

		if val, ok := m[nums[i]]; !ok {
			m[nums[i]] = 1
		} else {
			if val+1 > (len(nums) / 2) {
				return nums[i]
			}
			m[nums[i]]++
		}
	}
	return doesntExist
}
