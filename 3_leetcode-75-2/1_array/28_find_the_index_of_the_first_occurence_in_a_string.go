package array

/*
28. Find the Index of the First Occurrence in a String
Easy
4.3K
227
Companies

Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:
Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.

Example 2:
Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" did not occur in "leetcode", so we return -1.
 

Constraints:

1 <= haystack.length, needle.length <= 104
haystack and needle consist of only lowercase English characters.

*/

import (
	"strings"
)

func strStr(haystack string, needle string) int {

    //return usingBuiltInFn(haystack, needle)

    return bruteForce(haystack, needle)

}

func bruteForce(haystack, needle string) int {
    
    for pos:=0; pos<len(haystack); pos++ {
        if haystack[pos]==needle[0] && isFullMatch(haystack, needle, pos) {
                return pos
        }
    }

    return -1
}

func isFullMatch(haystack, needle string, pos int) bool {

    for i,j:=pos,0; i<len(haystack) && j<len(needle) && haystack[i]==needle[j]; i,j=i+1,j+1 {
        if j==len(needle)-1 {
            return true
        }
    }

    return false
}

func usingBuiltInFn(haystack, needle string) int {

    return strings.Index(haystack, needle)
}
