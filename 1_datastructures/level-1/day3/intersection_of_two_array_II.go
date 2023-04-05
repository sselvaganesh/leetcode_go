package array

/*
Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.

Example 1:
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]

Example 2:
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]
Explanation: [9,4] is also accepted.
*/

import (
	"sort"
)

func Intersect(nums1 []int, nums2 []int) []int {

	//return solution1(nums1, nums2)

	return solution2(nums1, nums2)

}

func solution2(nums1, nums2 []int) []int {

	if len(nums1) == 0 && len(nums2) == 0 {
		return nil
	}

	var result []int
	m := make(map[int]int) //element vs occurences
	for i := 0; i < len(nums1); i++ {
		if _, ok := m[nums1[i]]; ok {
			m[nums1[i]]++
		} else {
			m[nums1[i]] = 1
		}
	}

	for j := 0; j < len(nums2); j++ {

		if val, ok := m[nums2[j]]; ok {

			result = append(result, nums2[j])

			if val == 1 {
				delete(m, nums2[j])
			} else {
				m[nums2[j]]--
			}
		}

	}

	return result
}

func solution1(nums1, nums2 []int) []int {

	if len(nums1) == 0 && len(nums2) == 0 {
		return nil
	}

	var result []int

	sort.Ints(nums1)
	sort.Ints(nums2)

	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {

		if nums1[i] == nums2[j] {
			result = append(result, nums1[i])
			i++
			j++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}

	}

	return result
}
