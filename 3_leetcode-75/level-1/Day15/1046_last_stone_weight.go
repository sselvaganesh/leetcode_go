package heap

/*
1046. Last Stone Weight
Easy
4.4K
83
Companies

You are given an array of integers stones where stones[i] is the weight of the ith stone.

We are playing a game with the stones. On each turn, we choose the heaviest two stones and smash them together. Suppose the heaviest two stones have weights x and y with x <= y. The result of this smash is:

If x == y, both stones are destroyed, and
If x != y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.
At the end of the game, there is at most one stone left.

Return the weight of the last remaining stone. If there are no stones left, return 0.

Example 1:
Input: stones = [2,7,4,1,8,1]
Output: 1
Explanation:
We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of the last stone.

Example 2:
Input: stones = [1]
Output: 1

Constraints:
1 <= stones.length <= 30
1 <= stones[i] <= 1000

*/

import (
	"sort"
)

func LastStoneWeight(stones []int) int {

	//return solution1(stones)

	return solution2(stones)

}

func solution2(stones []int) int {

	max := func(val1, val2 int) (int, int) {
		if val1 > val2 {
			return val1, val2
		}
		return val2, val1
	}

	for len(stones) > 1 {

		maxHeapify(stones)

		if len(stones) > 2 {
			stones[1], stones[2] = max(stones[1], stones[2])
		}

		if stones[0] == stones[1] {
			stones = append([]int{}, stones[2:]...)
		} else {
			stones = append([]int{stones[0] - stones[1]}, stones[2:]...)
		}

	}

	if len(stones) == 0 {
		return 0
	}

	return stones[0]
}

func maxHeapify(stones []int) {

	for i := len(stones) - 1; i >= 0; i-- {
		heapify(stones, i)
	}

}

func heapify(stones []int, curElem int) {

	totElem := len(stones)
	left, right := (2*curElem)+1, (2*curElem)+2

	high := curElem
	if left < totElem && stones[left] >= stones[high] {
		high = left
	}

	if right < totElem && stones[right] >= stones[high] {
		high = right
	}

	if high != curElem {
		stones[high], stones[curElem] = stones[curElem], stones[high]
		heapify(stones, high)
	}

}

func solution1(stones []int) int {

	sortFunc := func(i, j int) bool {
		return stones[i] > stones[j]
	}
	sort.Slice(stones, sortFunc)

	for len(stones) > 1 {
		if stones[0] == stones[1] {
			stones = append([]int{}, stones[2:]...)
		} else {
			stones = append([]int{stones[0] - stones[1]}, stones[2:]...)
		}
		sort.Slice(stones, sortFunc)
	}

	if len(stones) == 0 {
		return 0
	}

	return stones[0]

}
