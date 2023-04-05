package paranthesisvalidation

import (
	"testing"
)


func TestValidateParanthesisUsingRune(t *testing.T) {

	type testcase struct {
		testId int
		inp string
		result bool
	}

	runtest := func(tc testcase) {
		if actual := ValidateParanthesisUsingRune(tc.inp); actual != tc.result {		
			t.Fatalf("%+v; Actual: [%v]", tc, actual)
		}
		
		if actual := ValidateParanthesisUsingByte(tc.inp); actual != tc.result {
			t.Fatalf("%+v; Actual: [%v]", tc, actual)
		}
	}

	tests := []testcase{
		{testId: 1, inp: "", result: true},
		{testId: 2, inp: "()", result: true},
		{testId: 3, inp: "()[]{}", result: true},
		{testId: 4, inp: "(]", result: false},
	}

	for _, testObj := range tests {
		runtest(testObj)
	}
}

