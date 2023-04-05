package binarysearch

import(

)

func FindInsertPosition(nums []int, target int) int {

	start := 0
	end := len(nums)-1
	
	for ; end>=start ; {
	
		if target > nums[end] {
			return end+1
		} else if target < nums[start] {
			return start
		}
	
		mid := (start+end)/2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			start = mid+1
		} else {
			end = mid-1
		}
	
	}	
	
	
	return 0
}
