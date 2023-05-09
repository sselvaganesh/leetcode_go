package dp

/*
198. House Robber
Medium
17.4K
330
Companies

You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.


Example 1:
Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.

Example 2:
Input: nums = [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
Total amount you can rob = 2 + 9 + 1 = 12.


Constraints:
1 <= nums.length <= 100
0 <= nums[i] <= 400


*/

func Rob(nums []int) int {

	//m:=make(map[int]int)
	//return dfs_rob(nums, m, 0)

	//return dp1(nums)

	return dp2(nums)
}

func dp2(nums []int) int {

	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	cash := make([]int, len(nums))
	cash[0] = nums[0]
	cash[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {

		if cash[i-2]+nums[i] > cash[i-1] {
			cash[i] = cash[i-2] + nums[i]
		} else {
			cash[i] = cash[i-1]
		}

	}

	return cash[len(cash)-1]
}

func dfs_rob(nums []int, m map[int]int, idx int) int {

	if idx >= len(nums) {
		return 0
	} else if idx == len(nums)-1 {
		return nums[idx]
	} else if idx == len(nums)-2 {
		return max(nums[idx], nums[idx+1])
	}

	if val, ok := m[idx]; ok {
		return val
	}

	val := max(nums[idx]+dfs_rob(nums, m, idx+2), dfs_rob(nums, m, idx+1))
	m[idx] = val

	return val

}

func max(val1, val2 int) int {

	if val1 > val2 {
		return val1
	}
	return val2
}
