package array

/*
345. Reverse Vowels of a String
Easy
3.4K
2.4K
Companies
Given a string s, reverse only all the vowels in the string and return it.

The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.


Example 1:
Input: s = "hello"
Output: "holle"

Example 2:
Input: s = "leetcode"
Output: "leotcede"
 

Constraints:

1 <= s.length <= 3 * 105
s consist of printable ASCII characters.
*/

func ReverseVowels(s string) string {

    m:=map[byte]struct{} {
        'a': struct{}{},
        'e': struct{}{},
        'i': struct{}{},
        'o': struct{}{},
        'u': struct{}{},
        'A': struct{}{},
        'E': struct{}{},
        'I': struct{}{},
        'O': struct{}{},
        'U': struct{}{},
    }

    goRight:=func(pos, last int) int {
        for i:=pos; last>i; i++ {
            if _, ok:=m[s[i]]; ok {
                return i
            }
        }
        return -1
    }

    goLeft:=func(pos, begin int) int {
        for i:=pos; i>begin; i-- {
            if _, ok:=m[s[i]]; ok {
                return i
            }
        }
        return -1
    }

    res:=[]byte(s)
    i, j :=0, len(s)-1
    for j>i {

        l:=goRight(i, j)
        if l==-1 {
            break
        }

        r:=goLeft(j, l)
        if r==-1 {
            break
        }

        res[l], res[r] = res[r], res[l]
        i=l+1
        j=r-1
    }

    return string(res)
}
