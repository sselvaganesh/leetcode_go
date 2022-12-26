package easy


func TwoSum1(nums []int, target int) []int {

	//Time Complexity: O(n^2)
	//Space Complexity: O(1) - No additional space

	var result []int
	var ok bool

	for i:=0; i<len(nums); i++ {
		for j:=i+1; j<len(nums); j++ {
			if (nums[i]+nums[j] == target) {
				result = append(result, i)
				result = append(result, j)
				ok = true
				break
			}
		}

		if ok {
			break
		}

	}

	return result
}


func TwoSum2(nums []int, target int) []int {
	
	var result []int
	inpMap := make(map[int][]int)

	// Time Complexity: O(n) 
	// Space Complexity: O(n)
	for i:=0; i<len(nums); i++ {
		
		val, ok := inpMap[nums[i]]
		if ok {
			if (nums[i]*2 == target) {
				result = append(result, val[0])
				result = append(result, i)
				break
			}		
		} else {		
			val, ok := inpMap[target-nums[i]]
			if ok {
				result = append(result, val[0])
				result = append(result, i)
				break
			}		
		}
		
		val = append(val, i)
		inpMap[nums[i]] = val		
	}

	return result
}




























