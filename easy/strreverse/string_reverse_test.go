package strreverse

import(
	"testing"
)


func TestStringReverse(t *testing.T){

	type testcase struct {
		testId uint8
		inp string
		rev string
	}

	runtest := func(tc testcase) {		
		if rev:=StringReverse(tc.inp); rev != tc.rev {
			t.Fatalf("TestCase [%+v] Actual [%s]", tc, rev)
		}
	}

	tests := []testcase{
		{testId: 1, inp: "selva", rev: "avles"},
		{testId: 2, inp: "hsenag", rev: "ganesh"},
		{testId: 3, inp: "", rev: ""},
		{testId: 4, inp: "1234", rev: "4321"},
		{testId: 5, inp: "golang", rev: "gnalog"},
		{testId: 6, inp: "stack", rev: "kcats"},		
	}
	
	for _, test := range tests {
		runtest(test)
	}


}
