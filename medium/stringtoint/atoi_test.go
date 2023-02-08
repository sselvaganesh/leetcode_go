package stringtoint

import (
	"testing"
)

func TestMyAtoi(t *testing.T) {

	type testcase struct {
		id int
		inp string
		result int
	}

	runtest := func(tc testcase) {
		if result := MyAtoi(tc.inp); result != tc.result {
			t.Fatalf("[%+v] Actual [%v]", tc, result)
		}
	}

	tests := []testcase{
		{id: 1, inp: "", result: 0},
		{id: 2, inp: " ", result: 0},
		{id: 3, inp: "+", result: 0},
		{id: 4, inp: "-", result: 0},
		{id: 5, inp: "+-+-+-", result: 0},
		{id: 6, inp: "achsdfd", result: 0},
		{id: 7, inp: "0", result: 0},
		{id: 8, inp: "1", result: 1},
		{id: 9, inp: "1234", result: 1234},
		{id: 10, inp: "+1234", result: 1234},
		{id: 11, inp: "-1234", result: -1234},
		{id: 12, inp: " +1234", result: 1234},
		{id: 13, inp: " +1234 ", result: 1234},
		{id: 14, inp: " +1234 ", result: 1234},
		{id: 15, inp: " +12  34", result: 12},
		{id: 16, inp: "-91283472332", result: -91283472332},
		//{id: , inp: "", result: },
	}

	for _, testObj := range tests {
		runtest(testObj)
	}
	
}
