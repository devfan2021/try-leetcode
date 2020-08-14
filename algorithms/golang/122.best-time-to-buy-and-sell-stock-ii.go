/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 *
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/
 *
 * algorithms
 * Easy (56.35%)
 * Likes:    2655
 * Dislikes: 1784
 * Total Accepted:    639.7K
 * Total Submissions: 1.1M
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * Say you have an array prices for which the i^th element is the price of a
 * given stock on day i.
 *
 * Design an algorithm to find the maximum profit. You may complete as many
 * transactions as you like (i.e., buy one and sell one share of the stock
 * multiple times).
 *
 * Note: You may not engage in multiple transactions at the same time (i.e.,
 * you must sell the stock before you buy again).
 *
 * Example 1:
 *
 *
 * Input: [7,1,5,3,6,4]
 * Output: 7
 * Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit
 * = 5-1 = 4.
 * Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 =
 * 3.
 *
 *
 * Example 2:
 *
 *
 * Input: [1,2,3,4,5]
 * Output: 4
 * Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit
 * = 5-1 = 4.
 * Note that you cannot buy on day 1, buy on day 2 and sell them later, as you
 * are
 * engaging multiple transactions at the same time. You must sell before buying
 * again.
 *
 *
 * Example 3:
 *
 *
 * Input: [7,6,4,3,1]
 * Output: 0
 * Explanation: In this case, no transaction is done, i.e. max profit = 0.
 *
 *
 * Constraints:
 *
 *
 * 1 <= prices.length <= 3 * 10 ^ 4
 * 0 <= prices[i] <= 10 ^ 4
 *
 *
 */

// @lc code=start
func maxProfit(prices []int) int {
	return maxProfit2(prices)
}

// [2,4,1], [1,2,3,4,5]
func maxProfit2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	maxProfit, minPrice, sellPrice := 0, prices[0], prices[0]
	for i := 1; i < len(prices); {
		if prices[i] < minPrice && minPrice >= sellPrice { // handle this case [2,4,1]
			minPrice = prices[i]
			sellPrice = prices[i]
			i++
		} else {
			if prices[i] < sellPrice {
				maxProfit += sellPrice - minPrice
				minPrice, sellPrice = prices[i], prices[i]
			} else {
				sellPrice = prices[i]
				i++
				if i >= len(prices) { // last element  [1,2,3,4,5]
					maxProfit += sellPrice - minPrice
				}
			}
		}
	}
	return maxProfit
}

func maxProfit1(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	maxProfit := 0
	for i := 0; i < len(prices); {
		preVal := prices[i]
		maxVal := preVal
		curIndex := i
		for j := i + 1; j < len(prices); j++ {
			if prices[j] > maxVal {
				curIndex = j
				maxVal = prices[j]
			} else {
				break
			}
		}
		if maxVal > preVal {
			maxProfit += maxVal - preVal
		}
		i = curIndex + 1
	}

	return maxProfit
}

// @lc code=end