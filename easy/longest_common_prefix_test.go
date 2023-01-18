package easy

import(
	"testing"
)


func TestFindLcp(t *testing.T) {

	type testcase struct{
		testId int
		input []string
		lcp string
	}


	runtest := func(tc testcase) {
		if lcp:=GetLcp(tc.input); lcp!=tc.lcp {
			t.Fatalf("Failed: %+v Result: %v", tc, lcp)
		}
	}

	tests:= []testcase{
		{testId: 1, input: []string{"flower","flow","flight"}, lcp: "fl"},
		{testId: 2, input: []string{"dog","racecar","car"}, lcp: ""},
		{testId: 3, input: []string{"cir","car"}, lcp: "c"},
		{testId: 4, input: []string{"reflower","flow","flight"}, lcp: "fl"},
		//{testId: 5, input: []string{}, lcp: ""},
		//{testId: 6, input: []string{}, lcp: ""},		
	}


	for _, test := range tests{
		runtest(test)
	}
	
}
