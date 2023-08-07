package array

/*
443. String Compression

Medium
4.1K
6.5K
Companies

Given an array of characters chars, compress it using the following algorithm:

Begin with an empty string s. For each group of consecutive repeating characters in chars:

If the group's length is 1, append the character to s.
Otherwise, append the character followed by the group's length.
The compressed string s should not be returned separately, but instead, be stored in the input character array chars. Note that group lengths that are 10 or longer will be split into multiple characters in chars.

After you are done modifying the input array, return the new length of the array.

You must write an algorithm that uses only constant extra space.

Example 1:
Input: chars = ["a","a","b","b","c","c","c"]
Output: Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]
Explanation: The groups are "aa", "bb", and "ccc". This compresses to "a2b2c3".

Example 2:
Input: chars = ["a"]
Output: Return 1, and the first character of the input array should be: ["a"]
Explanation: The only group is "a", which remains uncompressed since it's a single character.

Example 3:
Input: chars = ["a","b","b","b","b","b","b","b","b","b","b","b","b"]
Output: Return 4, and the first 4 characters of the input array should be: ["a","b","1","2"].
Explanation: The groups are "a" and "bbbbbbbbbbbb". This compresses to "ab12".
 

Constraints:
1 <= chars.length <= 2000
chars[i] is a lowercase English letter, uppercase English letter, digit, or symbol.

*/


import (
	"fmt"
	"strconv"
)

func Compress(chars []byte) int {
    
    if len(chars)==1 {
        return 1
    }

    // Assumption: Input slice is sorted
    start, count, prev := 0, 1, chars[0]
    for i:=1; i<len(chars); i++ {

        if chars[i]==prev {
            count++
        } else {

            numByteArr:=numArray(count)
            if count>1 {
                chars=append(append(chars[:start+1], numByteArr...), chars[i:]...)
                start=start + len(numByteArr) + 1        
                i=start
            } else {
                start=i
            }
            prev=chars[i]
            count=1
        }
    }

    if count>1 {
        chars=append(chars[:start+1], numArray(count)...)
    }

    return len(chars)
}

func numArray(count int) []byte {

    var res []byte

    //fmt.Println(count)
    for count>0 {
        res=append([]byte(strconv.Itoa(count%10)), res...)
        count=count/10
    }

    return res
}

func print(chars []byte) {
    for _, each := range chars {
        fmt.Printf("%c ", each)
    }
    fmt.Println()
}





