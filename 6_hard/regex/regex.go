package regex

import(
	//"log"
	"fmt"
)

func print(table [][]bool, s, p string) {

	fmt.Printf("string [%v] pattern: [%v]\n", s, p)

	for i:= range table {
		fmt.Println(table[i])		
	}
}

func IsMatch(s string, p string) bool {
	m := make(map[string]bool)
	return isOk(s, p, m)
}

func isOk(s, p string, m map[string]bool) bool {

	if len(s) == 0 && len(p) == 0 {
		return true
	} else if len(s) > 0 && len(p) == 0 {
		return false
	} else if len(s) == 0 && len(p) == 1 {
		return false
	} else if len(s) == 0 && len(p) == 2 && p[1] == '*' {
		return true
	} else if len(s) == 1 && len(p) == 1 {
		if p[0] == '.' || s[0] == p[0] {
			return true
		} else {
			return false
		}

	}

	if len(p) >= 2 && p[1] == '*' {

		if len(s) == 0 {
			return isOk(s, p[2:], m)
		} else {
			if p[0] == '.' || (s[0] == p[0]) {
				return isOk(s, p[2:], m) || isOk(s[1:], p, m)
			} else {
				return isOk(s, p[2:], m)
			}
		}

	}

	if len(p) > 0 && len(s) > 0 {
		if p[0] == '.' || s[0] == p[0] {
			return isOk(s[1:], p[1:], m)
		} else if s[0] != p[0] {
			return false
		}
	}

	return false

}

func Leetcode(s, p string) bool {

	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2] || (i > 0 && dp[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.'))
			} else {
				dp[i][j] = i > 0 && dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.')
			}
		}
	}
	print(dp, s, p)
	return dp[m][n]

}
