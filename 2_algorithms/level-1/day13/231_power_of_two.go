package bit_manupulation

/*
231. Power of Two
Easy
5.1K
357
Companies

Given an integer n, return true if it is a power of two. Otherwise, return false.

An integer n is a power of two, if there exists an integer x such that n == 2x.


Example 1:
Input: n = 1
Output: true
Explanation: 20 = 1

Example 2:
Input: n = 16
Output: true
Explanation: 24 = 16

Example 3:
Input: n = 3
Output: false
 
Constraints:
-231 <= n <= 231 - 1
 
Follow up: Could you solve it without loops/recursion?
*/


import (
    "fmt"
    "strings"
    "unsafe"
)



func IsPowerOfTwo(n int) bool {

   //return solution1(n)
    return solution2(n)
}

func solution2(n int) bool {

    return len(strings.Trim(fmt.Sprintf("%b", n), "0"))==1

}


func solution1(n int) bool {

    loop:= int(unsafe.Sizeof(n)*4)

    for i:=0; i<=loop; i++ {
        if n==(1<<i) {
            return true
        }

    }

    return false
}
