package longestsubstring

func LengthOfLongestSubstring1(s string) int {

	if len(s) <= 1 {
		return len(s)
	}

	var byteSlice []byte
	m := make(map[byte]struct{})
	max := 0

	for i := 0; i < len(s); i++ {

		if _, ok := m[s[i]]; ok {

			if len(byteSlice) > max {
				max = len(byteSlice)
			}

			j := 0
			for ; j < len(byteSlice); j++ {
				delete(m, byteSlice[j])
				if s[i] == byteSlice[j] {
					break
				}
			}
			byteSlice = byteSlice[j+1:]
		}

		byteSlice = append(byteSlice, s[i])
		m[s[i]] = struct{}{}

	}

	if len(byteSlice) > max {
		return len(byteSlice)
	}

	return max

}
