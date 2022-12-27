package easy

func IsPalindrome(inp int) bool {
	
	if (inp < 0) {
		return false
	}

	if (inp == 0) {
		return true
	}

	tempCopy := inp
	var remSlice []int

	for (tempCopy > 0) {
		rem := tempCopy%10
		remSlice = append(remSlice, rem)
		tempCopy=tempCopy/10	
	}

	for i,j := 0, len(remSlice)-1; i<j; i,j = i+1, j-1 {
		if (remSlice[i] != remSlice[j]) {
			return false
		}
	}
	
	return true
}
