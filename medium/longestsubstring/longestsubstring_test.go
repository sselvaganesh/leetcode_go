package longestsubstring

import (
	"testing"
	"time"
	"log"
)


func TestLengthOfLongestSubstring(t *testing.T) {

	type testcase struct {
		testId int
		inp string
		longest int
	}

	runtest := func(tc testcase) {
		if longest := LengthOfLongestSubstring(tc.inp); longest != tc.longest {
			t.Fatalf("%+v Actual: %v", tc, longest)
		}
	}

	tests := []testcase{
		{testId: 1, inp: "", longest: 0},
		{testId: 2, inp: "a", longest: 1},
		{testId: 3, inp: "ab", longest: 2},
		{testId: 4, inp: "abc", longest: 3},
		{testId: 5, inp: "abcda", longest: 4},
		{testId: 6, inp: "abcabcbb", longest: 3},
		{testId: 7, inp: "dihwofdwnhsvldnzsdxclgeka", longest: 11},
		//{testId: 1, inp: "", longest: },
	}


	start := time.Now()
	for _, testObj := range tests {	
		runtest(testObj)
	}
	
	duration := time.Since(start)
	log.Println("Execution time: ", duration)
}




