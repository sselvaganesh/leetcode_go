package stringtoint

import (
	"fmt"
	//"strconv"
	"strings"
)

func MyAtoi(s string) int {

	if s == "" {
		return 0
	}

	validStr, isNeg := getValidStr(strings.TrimSpace(s))
	if validStr == "" {
		return 0
	}
	
	result := getInt(validStr)

	if isNeg {
		return -result
	}
	return result
}


func getInt(s string) int {

	var result int
	for i:=0; i<len(s); i++ {
		result = (result * 10) + int(s[i] - '0')
	}

	return result
}


func getValidStr(s string) (string, bool) {

	if len(s) == 0 {
		return "", false
	}

	var result []byte
	var isNeg bool
	
	if s[0] == '-' {
		isNeg = true
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	for i:=0; i<len(s); i++ {
		if isNum(s[i]) {
			result = append(result, s[i])
		} else {
			break
		}
	}
	
	fmt.Printf("Valid str: [%v] isNeg: [%v]\n", string(result), isNeg)
	return string(result), isNeg
}


func isNum(inp byte) bool {
	return inp >= '0' && inp <= '9'
}
