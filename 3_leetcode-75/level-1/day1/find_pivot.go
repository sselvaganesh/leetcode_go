package dayone

/*

Given an array of integers nums, calculate the pivot index of this array.
The pivot index is the index where the sum of all the numbers strictly to the left of the index is equal to the sum of all the numbers strictly to the index's right.
If the index is on the left edge of the array, then the left sum is 0 because there are no elements to the left. This also applies to the right edge of the array.
Return the leftmost pivot index. If no such index exists, return -1.


Input: nums = [1,7,3,6,5,6]
Output: 3
Explanation:
The pivot index is 3.
Left sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11
Right sum = nums[4] + nums[5] = 5 + 6 = 11


Input: nums = [1,2,3]
Output: -1
Explanation:
There is no index that satisfies the conditions in the problem statement.

Input: nums = [2,1,-1]
Output: 0
Explanation:
The pivot index is 0.
Left sum = 0 (no elements to the left of index 0)
Right sum = nums[1] + nums[2] = 1 + -1 = 0

*/

func PivotIndex(nums []int) int {

	if len(nums) == 0 {
		return -1
	}

	leftSum, rightSum := 0, sum(nums[1:]...)
	for pivot := 0; pivot < len(nums); pivot++ {
		if leftSum == rightSum {
			return pivot
		}
		leftSum += nums[pivot]
		if pivot+1 < len(nums) {
			rightSum -= nums[pivot+1]
		} else {
			rightSum -= nums[pivot]
		}

	}

	return -1
}

func sum(inp ...int) int {

	result := 0
	for _, val := range inp {
		result += val
	}
	return result
}
