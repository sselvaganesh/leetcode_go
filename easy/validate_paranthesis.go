package easy

import (
	//"fmt"
)

func ValidateParanthesisUsingRune(inp string) bool {

	if len(inp)%2 != 0 {
		return false
	}
	
	var clsRune []rune
	openVsClose := map[rune]rune {
		 '(' : ')' ,
		 '{' : '}' ,
		 '[' : ']' ,
	}
	
	for _, each := range inp {
	
		cls, ok := openVsClose[each]
		if ok {
			clsRune = append(clsRune, cls)
			continue
		} 
		
		if len(clsRune) == 0 || clsRune[len(clsRune)-1] != each {
			return false
		}
		
		clsRune = clsRune[:len(clsRune)-1]
	}
	
	return len(clsRune) == 0
}



func ValidateParanthesisUsingByte(inp string) bool {

	if len(inp)%2 != 0 {
		return false
	}
	
	openVsClose := map[byte]byte{
		'(' : ')',
		'{' : '}',
		'[' : ']',
	}	
	clsByte := []byte{}

	for i:=0; i<len(inp); i++ {

		cls, ok := openVsClose[inp[i]]
		if ok {
			clsByte = append(clsByte, cls)
			continue			
		}
		
		if len(clsByte) == 0 || clsByte[len(clsByte)-1] != inp[i] {
			return false
		}

		clsByte = clsByte[:len(clsByte)-1]
	}


	return len(clsByte) == 0
}
