package array

/*
1071. Greatest Common Divisor of Strings
Easy
3.5K
551
Companies
For two strings s and t, we say "t divides s" if and only if s = t + ... + t (i.e., t is concatenated with itself one or more times).

Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.

Example 1:
Input: str1 = "ABCABC", str2 = "ABC"
Output: "ABC"

Example 2:
Input: str1 = "ABABAB", str2 = "ABAB"
Output: "AB"

Example 3:
Input: str1 = "LEET", str2 = "CODE"
Output: ""
 
Constraints:

1 <= str1.length, str2.length <= 1000
str1 and str2 consist of English uppercase letters.

*/

func GcdOfStrings(str1 string, str2 string) string {


    l1, l2 := len(str1), len(str2)
    if l1>l2 {
        str1, str2 = str2, str1
    }


    for i:=len(str1); i>0; i-- {

        if l1%i==0 && l2%i==0 {
            if isValid(str1, str2, i) {
                return str1[:i]
            }
        }
    }
    
    return ""
 
}

func isValid(str1, str2 string, ln int) bool {

    makeSlice:=func(str string, ln int) []string {
        var res []string
        start:=0
        for end:=ln; len(str)>=end; end=end+ln {
            res=append(res, str[start:end])
            start=end
        }
        return res
    }

    isAllSame:=func(inp []string, str string) bool {
        for _, val:=range inp {
            if val!=str {
                return false
            }
        }
        return true
    }

    s1, s2 := makeSlice(str1, ln), makeSlice(str2, ln)
    //fmt.Println(str1, str2, ln)
    //fmt.Println(s1, s2)


    return isAllSame(s1, str1[:ln]) && isAllSame(s2, str1[:ln])
}
