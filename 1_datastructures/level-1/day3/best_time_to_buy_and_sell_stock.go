package array

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

func print(table [][]int) {

	for i := 0; i < len(table); i++ {
		fmt.Println(table[i])
	}

}

func MaxProfit(prices []int) int {

	//return solution1(prices)
	return solution2_dp(prices)

}

func solution2_dp(prices []int) int {

	if len(prices) <= 1 {
		return 0
	}

	tot := len(prices)

	table := make([][]int, tot+1)
	for i := range table {
		table[i] = make([]int, tot+1)
	}

	for buy := 0; buy <= len(prices); buy++ {

		for sell := 0; sell <= len(prices); sell++ {

			if buy == 0 || sell == 0 || sell < buy {
				continue
			}

			strikePrice := prices[sell-1] - prices[buy-1]

			if strikePrice > max(table[buy-1][sell], table[buy][sell-1]) {
				table[buy][sell] = strikePrice
			} else {
				table[buy][sell] = max(table[buy-1][sell], table[buy][sell-1])
			}

		}

	}

	print(table)

	return table[len(prices)][len(prices)]

}

func max(val1, val2 int) int {

	if val1 > val2 {
		return val1
	}

	return val2
}

func solution1_array(prices []int) int {

	if len(prices) <= 1 {
		return 0
	}

	buy, profit := prices[0], 0
	for _, sell := range prices[1:] {

		if sell-buy > profit {
			profit = sell - buy
		}

		if buy > sell {
			buy = sell
		}

	}

	return profit

}
