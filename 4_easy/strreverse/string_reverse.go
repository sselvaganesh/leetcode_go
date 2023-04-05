package strreverse

func StringReverse(inp string) string {

	strReverse := []rune(inp)
	for i,j :=0, len(inp)-1; j>i; i,j =i+1, j-1 {
		strReverse[i], strReverse[j] = strReverse[j], strReverse[i]
	}
	return string(strReverse)
}


