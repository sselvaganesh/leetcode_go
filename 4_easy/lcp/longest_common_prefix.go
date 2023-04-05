package lcp

import(
	"strings"
)

func max(lcp1, lcp2 string) string {

	if len(lcp1) >= len(lcp2) {
		return lcp1
	} 
	
	return lcp2
}


func helper(str1, str2 string) string {


	if len(str2) > len(str1) {
		str1, str2 = str2, str1
	}

	if strings.Contains(str1, str2) {
		return str2
	}

	return helper(str1, str2[:len(str2)-1])
}


func GetLcp(input []string) string {

	if len(input)==0 {
		return ""
	} else if len(input)==1 {
		return input[0]
	}

	result := helper(input[0], input[1])
	if result == ""{
		return ""
	}

	for i:=2; i<len(input); i++ {
		if result=helper(result, input[i]); result=="" {
			return ""
		}
	}
	
	return result
}
