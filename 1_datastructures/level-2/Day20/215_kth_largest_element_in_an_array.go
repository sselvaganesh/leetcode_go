package heap

/*
215. Kth Largest Element in an Array
Medium
13.8K
683
Companies

Given an integer array nums and an integer k, return the kth largest element in the array.

Note that it is the kth largest element in the sorted order, not the kth distinct element.

You must solve it in O(n) time complexity.

Example 1:
Input: nums = [3,2,1,5,6,4], k = 2
Output: 5

Example 2:
Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4
 
Constraints:
1 <= k <= nums.length <= 105
-104 <= nums[i] <= 104
*/

func FindKthLargest(nums []int, k int) int {

    tot:=len(nums)
    for i:=(tot/2)-1; i>=0; i-- {
        buildMaxHeap(nums, tot, i)
    }

    for i:=1; i<=k-1; i++ {
        nums[0], nums[tot-1] = nums[tot-1], nums[0]
        tot--
        buildMaxHeap(nums, tot, 0)        
    }

    return nums[0]
}

func buildMaxHeap(nums []int, tot, idx int) {

    l,r:=(2*idx)+1, (2*idx)+2
    largest:=idx

    if l<tot && nums[l]>nums[largest] {
        largest=l
    }

    if r<tot && nums[r]>nums[largest] {
        largest=r
    }

    if largest!=idx {
        nums[idx], nums[largest] = nums[largest], nums[idx]
        buildMaxHeap(nums, tot, largest)
    }

}

