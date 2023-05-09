package backtracking

/*
46. Permutations
Medium
15.2K
255
Companies

Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.

Example 1:
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Example 2:
Input: nums = [0,1]
Output: [[0,1],[1,0]]

Example 3:
Input: nums = [1]
Output: [[1]]


Constraints:
1 <= nums.length <= 6
-10 <= nums[i] <= 10
All the integers of nums are unique.
*/

func Permute(nums []int) [][]int {

	if len(nums) == 0 {
		return nil
	}

	return solution(nums)

}

func solution(nums []int) [][]int {

	if len(nums) == 1 {
		r := append([]int{}, nums...)
		return [][]int{r}
	}

	var result [][]int
	for _ = range nums {

		zeroPos := nums[0]
		nums = nums[1:]
		temp := solution(nums)

		for j := range temp {
			r := append([]int{}, temp[j]...)
			r = append(r, zeroPos)
			result = append(result, r)
		}

		nums = append(nums, zeroPos)
	}

	return result
}
