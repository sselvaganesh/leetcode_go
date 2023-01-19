package removedupfromsortedarraay

import(
	"fmt"
)

func RemoveDuplicates(nums []int) int {

	if len(nums) == 1 {
		return 1
	}

	lastElem := nums[0]
	swapPos := 1
	traversePos := 1
	
	for ; len(nums)>traversePos; {
	
		if lastElem == nums[traversePos] {
			traversePos++
		} else {
			nums[traversePos], nums[swapPos] = nums[swapPos], nums[traversePos]
			lastElem = nums[swapPos]
			swapPos++
			traversePos++
		}
	}

	fmt.Printf("Actual: %v", nums[:swapPos])
	return swapPos
}
