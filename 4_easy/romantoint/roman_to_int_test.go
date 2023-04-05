package romantoint

import(
	"testing"
)

func TestRomanToInt(t *testing.T){

	type testcase struct {
		testId uint8
		roman string
		val int
	}
	
	
	runtest := func(tc testcase) {
		if val := RomanToInt(tc.roman); val != tc.val {
			t.Fatalf("Failed: [%+v] Actual: [%v]", tc, val)
		}
	}
	
	tests := []testcase{
		{testId: 1, roman: "III", val: 3},
		//{testId: 2, roman: "IV", val: 4},
		{testId: 3, roman: "V", val: 5},
		{testId: 4, roman: "VII", val: 7},
		{testId: 5, roman: "LVIII", val: 58},
		//{testId: 6, roman: "MCMXCIV", val: 1994},
		//{testId: 7, roman: "XIIV", val: 13},
				
	}

	for _, test := range tests {
		runtest(test)
	}

}
