package regex

import (
	"testing"
)

func TestIsMatch(t *testing.T) {

	type testcase struct {
		id     int
		s      string
		p      string
		result bool
	}

	runtest := func(tc testcase) {

		if result := Leetcode(tc.s, tc.p); result != tc.result {
			t.Fatalf("%+v, Actual: [%v]", tc, result)
		}

	}

	tests := []testcase{
		//{id: 1, s: "aa", p: "a", result: false},
		//{id: 2, s: "aa", p: "a*", result: true},
		//{id: 3, s: "ab", p: ".*", result: true},
		{id: 4, s: "aab", p: "c*a*b", result: true},
		//{id: 5, s: "aabcbcbcaccbcaabc", p: ".*a*aa*.*b*.c*.*a*", result: true},
		//{id: 5, s: "aaa", p: "a*a", result: true},
		//{id: 6, s: "aaaa", p: "a*aa.a", result: true},
		//{id: 7, s: "aaaaaaaaaaaaaaaaaa", p: "a*aaaa*aa.aaa.aaaa.a", result: true},
	}

	for _, testObj := range tests {
		runtest(testObj)
	}

}
