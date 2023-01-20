package binarysearch

import (
	"testing"
)

func TestFindInsertPosition(t *testing.T) {

	type testcase struct {
		testId int
		nums []int
		target int
		result int
	}
	
	runtest := func(tc testcase) {
		if actual := FindInsertPosition(tc.nums, tc.target); actual != tc.result {
			t.Fatalf("%+v; Actual: [%d]", tc, actual)
		}
	}
	
	tests := []testcase{
		{testId: 1, nums: []int{1, 5, 7, 10, 11, 16, 18}, target:12, result: 5},
		{testId: 2, nums: []int{1,3,5,6}, target: 5, result: 2},
		{testId: 3, nums: []int{1,3,5,6}, target: 2, result: 1},
		{testId: 4, nums: []int{1,3,5,6}, target: 7, result: 4},
		//{testId: 5, nums: []int{}, target: result: },		
	}

	for _, testObj := range tests {
		runtest(testObj)
	}
}
