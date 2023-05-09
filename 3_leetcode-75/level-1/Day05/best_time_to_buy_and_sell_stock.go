package greedy

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.
You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.


Example 1:
Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

Example 2:
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.
*/

import (
	"fmt"
)

func print(table [][]int, prices []int) {

	for i := 0; i < len(table); i++ {
		if i == 0 {
			fmt.Println(0, " : ", table[i])
		} else {
			fmt.Println(prices[i-1], " : ", table[i])
		}

	}

}

func MaxProfit(prices []int) int {

	return greedy(prices)
	//return dynamicProgramming(prices)

}

func dynamicProgramming(prices []int) int {

	if len(prices) <= 1 {
		return 0
	}

	table := make([][]int, len(prices)+1)
	for i := range table {
		table[i] = make([]int, len(prices)+1)
	}

	for buy := 0; buy <= len(prices); buy++ {

		for sell := 0; sell <= len(prices); sell++ {

			if buy == 0 || sell == 0 {
				continue
			}

			if sell <= buy {
				table[buy][sell] = table[buy-1][sell]
				continue
			}

			buyPrice := prices[buy-1]
			sellPrice := prices[sell-1]

			if (sellPrice - buyPrice) > max(table[buy-1][sell], table[buy][sell-1]) {
				table[buy][sell] = sellPrice - buyPrice
			} else {
				table[buy][sell] = max(table[buy-1][sell], table[buy][sell-1])
			}

		}

	}

	return table[len(prices)][len(prices)]

}

func max(val1, val2 int) int {

	if val1 > val2 {
		return val1
	}

	return val2
}

func greedy(prices []int) int {

	if len(prices) <= 1 {
		return 0
	}

	buyPrice, profit := prices[0], 0
	for _, sellPrice := range prices[1:] {

		if (sellPrice - buyPrice) > profit {
			profit = sellPrice - buyPrice
		}

		if sellPrice < buyPrice {
			buyPrice = sellPrice
		}

	}

	return profit
}
