package longestpalindrome

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {

	type testcase struct {
		testId int
		inp    string
		result string
	}

	runtest := func(tc testcase) {
		if actual := LongestPalindrome(tc.inp); actual != tc.result {
			t.Fatalf("%+v Actual [%v]", tc, actual)
		}
	}

	tests := []testcase{
		{testId: 1, inp: "babad", result: "bab"},
		{testId: 2, inp: "cbbd", result: "bb"},
		{testId: 3, inp: "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth", result: "ranynar"},
		//{testId: , inp: "", result: ""},
	}

	for _, testObj := range tests {
		runtest(testObj)
	}

}
