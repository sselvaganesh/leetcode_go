package stack

/*
844. Backspace String Compare
Easy
6.2K
286
Companies

Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.

Note that after backspacing an empty text, the text will continue empty.

Example 1:
Input: s = "ab#c", t = "ad#c"
Output: true
Explanation: Both s and t become "ac".

Example 2:
Input: s = "ab##", t = "c#d#"
Output: true
Explanation: Both s and t become "".

Example 3:
Input: s = "a#c", t = "b"
Output: false
Explanation: s becomes "c" while t becomes "b".

*/

import (
	"reflect"
)


func BackspaceCompare(s string, t string) bool {
    //return solution1(s, t)
    return solution2(s, t)
}


func solution2(s, t string) bool {

    lastIndex:=func(i int) int {
        if i<=0{
            return 0
        }
        return i-1
    }

    for i:=0; i<len(s); {
        if s[i]=='#' {
            s=s[:lastIndex(i)] + s[i+1:]
            i=lastIndex(i)
        } else {
            i++
        }
    }

    for i:=0; i<len(t); {
        if t[i]=='#' {
            t=t[:lastIndex(i)] + t[i+1:]
            i=lastIndex(i)
        } else {
            i++
        }
    }

    return s==t

}

func solution1(s, t string) bool {
    return reflect.DeepEqual(excludeBackspace(s), excludeBackspace(t))
}

func excludeBackspace(inp string) ([]byte) {

    var result []byte
    for i:=0; i<len(inp); i++{
        if inp[i]=='#'{
            if len(result)==0{
                continue
            }
            result=result[:len(result)-1]
        } else {
            result=append(result, inp[i])
        }
    }
    return result
}
