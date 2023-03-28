package dayone

/*
Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
Return the running sum of nums.

Input: nums = [1,2,3,4]
Output: [1,3,6,10]
Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].

Input: nums = [1,1,1,1,1]
Output: [1,2,3,4,5]
Explanation: Running sum is obtained as follows: [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1].

Input: nums = [3,1,2,10,1]
Output: [3,4,6,16,17]

*/

func RunningSum(nums []int) []int {

	if len(nums) == 0 {
		return nil
	}

	var result []int
	temp := 0

	for i := 0; i < len(nums); i++ {
		temp += nums[i]
		result = append(result, temp)
	}

	return result

}
