package backtracking

import (
	"testing"
)

func TestLetterCasePermutation(t *testing.T) {

	input := "mQe"

	result := LetterCasePermutation(input)
	t.Logf("result: %v", result)

}
