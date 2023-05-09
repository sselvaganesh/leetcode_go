package array

/*
15. 3Sum
Medium
25K
2.3K
Companies

Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

Example 1:
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.

Example 2:
Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.

Example 3:
Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.


Constraints:
3 <= nums.length <= 3000
-105 <= nums[i] <= 105
*/

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {

	//return bruteForceThreeSum(nums)

	//return usingHashMap(nums)

	return twoPointers(nums)

}

func sortingSolution(nums []int) [][]int {

	var res [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {

		for j, k := i+1, len(nums)-1; k > j; {

			sum := nums[i] + nums[j] + nums[k]

			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})

				for j+1 < k && nums[j] == nums[j+1] {
					j++
				}

				for k-1 > j && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			} else if sum > 0 {
				k--
			} else {
				j++
			}

		}

		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}

	}

	return res
}

func twoPointers(nums []int) [][]int {

	var result [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums); {

		for j, k := i+1, len(nums)-1; k > j; {

			if nums[i]+nums[j]+nums[k] == 0 {

				result = append(result, []int{nums[i], nums[j], nums[k]})

				j++
				for j < len(nums) {
					if nums[j-1] == nums[j] {
						j++
						continue
					} else {
						break
					}
				}

				k--
				for k > j {
					if nums[k+1] == nums[k] {
						k--
						continue
					} else {
						break
					}
				}

			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else {
				j++
			}
		}

		i++
		for i < len(nums) {
			if nums[i-1] == nums[i] {
				i++
				continue
			} else {
				break
			}
		}

	}

	return result
}

func usingHashMap(nums []int) [][]int {

	var result [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums); {

		m := make(map[int]struct{})
		for j := i + 1; j < len(nums); {

			rem := -(nums[i] + nums[j])
			if _, ok := m[rem]; ok {
				result = append(result, []int{nums[i], nums[j], rem})

				j++
				for j < len(nums) {
					if nums[j-1] == nums[j] {
						j++
						continue
					} else {
						break
					}
				}

			} else {
				m[nums[j]] = struct{}{}
				j++
			}
		}

		i++
		for i < len(nums) {
			if nums[i-1] == nums[i] {
				i++
				continue
			} else {
				break
			}

		}

	}

	return result
}

func bruteForceThreeSum(nums []int) [][]int {

	var res [][]int

	sort.Ints(nums)
	for i := 0; i < len(nums); {

		for j := i + 1; j < len(nums); {

			for k := j + 1; k < len(nums); {

				if nums[i]+nums[j]+nums[k] == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}

				k++
				for k < len(nums) {
					if nums[k-1] == nums[k] {
						k++
						continue
					} else {
						break
					}

				}

			}

			j++
			for j < len(nums) {
				if nums[j-1] == nums[j] {
					j++
					continue
				} else {
					break
				}
			}

		}

		i++
		for i < len(nums) {
			if nums[i-1] == nums[i] {
				i++
				continue
			} else {
				break
			}
		}

	}

	return res
}
