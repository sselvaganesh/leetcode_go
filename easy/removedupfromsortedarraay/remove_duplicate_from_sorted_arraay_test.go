package removedupfromsortedarraay

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {

	type testcase struct {
		testId int
		nums []int
		result int
	}

	runtest := func(tc testcase){
		
		if actual := RemoveDuplicates(tc.nums); actual != tc.result {
			t.Fatalf("%+v; Actual: %v", tc, actual)
		}
	}
	
	tests := []testcase {
		{testId: 1, nums: []int{1,1,2}, result: 2},	//[1,2]
		{testId: 2, nums: []int{0,0,1,1,1,2,2,3,3,4}, result: 5}, //[0,1,2,3,4]
	}


	for _, testObj := range tests {
		runtest(testObj)
	}

}
