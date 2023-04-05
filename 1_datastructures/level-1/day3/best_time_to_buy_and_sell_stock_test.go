package array

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {

	prices := []int{7, 1, 5, 3, 6, 4}

	profit := MaxProfit(prices)

	t.Logf("MaxProfit [%v]", profit)

}
