package heap

/*
347. Top K Frequent Elements
Medium
13.4K
494
Companies

Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order. 

Example 1:
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

Example 2:
Input: nums = [1], k = 1
Output: [1]
 

Constraints:
1 <= nums.length <= 105
-104 <= nums[i] <= 104
k is in the range [1, the number of unique elements in the array].
It is guaranteed that the answer is unique.
 
Follow up: Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
*/


func TopKFrequent(nums []int, k int) []int {

    var res []int

    m:=make(map[int]int)
    for i:=0; i<len(nums); i++ {
        if _, ok:=m[nums[i]]; !ok {
            m[nums[i]]=1
        } else {
            m[nums[i]]++
        }
    }

    for i:=0; i<k; i++ {
        max:=getMax(m)
        res=append(res, max)
        delete(m, max)
    }

    return res
}

func getMax(m map[int]int) int {
    maxNum, maxOccur :=-1, -1
    for num, occur := range m {
        if occur>maxOccur {
            maxNum=num
            maxOccur=occur
        }
    }

    return maxNum
}

