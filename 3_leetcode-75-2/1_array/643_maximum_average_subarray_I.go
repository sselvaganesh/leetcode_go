package array

/*
643. Maximum Average Subarray I
Easy
2.8K
232
Companies

You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.

Example 1:
Input: nums = [1,12,-5,-6,50,3], k = 4
Output: 12.75000
Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75

Example 2:
Input: nums = [5], k = 1
Output: 5.00000
 
Constraints:

n == nums.length
1 <= k <= n <= 105
-104 <= nums[i] <= 104
*/


func findMaxAverage(nums []int, k int) float64 {

    var max float64 = -99999.99

    left, right, tot := 0, 0, 0

    for right<len(nums) {        

        if right<k {
            tot+=nums[right]
            right++
            continue
        }

        if tmp:=float64(tot)/float64(k); tmp>max {
            max=tmp
        }

        tot-=nums[left]
        tot+=nums[right]
        left++
        right++
    }

    if tmp:=float64(tot)/float64(k); tmp>max {
        max=tmp
    }

    return max
}
