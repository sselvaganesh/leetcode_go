package easy


func TwoSum(nums []int, target int) []int {
	return twoSum1(nums, target)
}


func twoSum1(nums []int, target int) []int {

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


