package easy

import (
	"testing"
) 

func TestIsPalindrome(t *testing.T) {

	type testcase struct {
		inp int
		result bool
	}	
	
	runtest := func(tc testcase) {	
		if output := IsPalindrome(tc.inp); output != tc.result {
			t.Fatalf("Failed: [%+v  Actual [%+v]]", tc, output)
		}
	}
	
	tests := []testcase {
		{inp: 1, result: true},
		{inp: 22, result: true},
		{inp: -1, result: false},
		{inp: 121, result: true},
		{inp: -121, result: false},
		{inp: 333, result: true},
		{inp: 0, result: true},
		//{inp: , result: },
	}
		
	for _, testObj := range tests {
		runtest(testObj)
	}
}
