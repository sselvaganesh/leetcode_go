package easy

import (
	"testing"
)

type testcase struct {
	testId int
	input []int
	target int
	expected []int
}
var tests = []testcase {
	{ testId: 1, input: []int{2,7,11,15}, target: 9, expected: []int{0,1}, }, 
	{ testId: 2, input: []int{3,2,4}, target: 6, expected: []int{1,2}, }, 
	{ testId: 3, input: []int{3,3}, target: 6, expected: []int{0,1}, }, 		
}


func TestTwoSum1(t *testing.T){

	runtest := func(tc testcase) {		
		result := TwoSum1(tc.input, tc.target)
		if result == nil {
			t.Fatalf("Failed: [%+v]. Output is nil", tc)
		} else if (len(result) < 2) {
			t.Fatalf("Failed: [%+v]. Output is not a 2 element array", tc)
		} else if (result[0] != tc.expected[0] || result[1] != tc.expected[1]) {
			t.Fatalf("Failed: [%+v]. Output [%v]", tc, result)
		} else {
			t.Logf("Passed: [%+v]", tc)
		}
	}
	
	for _, testObj := range tests {
		runtest(testObj)
	}
}


func TestTwoSum2(t *testing.T){

	runtest := func(tc testcase) {		
		result := TwoSum2(tc.input, tc.target)
		if result == nil {
			t.Fatalf("Failed: [%+v]. Output is nil", tc)
		} else if (len(result) < 2) {
			t.Fatalf("Failed: [%+v]. Output is not a 2 element array", tc)
		} else if (result[0] != tc.expected[0] || result[1] != tc.expected[1]) {
			t.Fatalf("Failed: [%+v]. Output [%v]", tc, result)
		} else {
			t.Logf("Passed: [%+v]", tc)
		}
	}
	
	for _, testObj := range tests {
		runtest(testObj)
	}
}


