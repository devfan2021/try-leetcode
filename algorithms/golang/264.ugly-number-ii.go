import "math"

/*
 * @lc app=leetcode id=264 lang=golang
 *
 * [264] Ugly Number II
 *
 * https://leetcode.com/problems/ugly-number-ii/description/
 *
 * algorithms
 * Medium (42.09%)
 * Likes:    2072
 * Dislikes: 128
 * Total Accepted:    184.2K
 * Total Submissions: 436.7K
 * Testcase Example:  '10'
 *
 * Write a program to find the n-th ugly number.
 *
 * Ugly numbers are positive numbers whose prime factors only include 2, 3,
 * 5.
 *
 * Example:
 *
 *
 * Input: n = 10
 * Output: 12
 * Explanation: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 is the sequence of the first 10
 * ugly numbers.
 *
 * Note:
 *
 *
 * 1 is typically treated as an ugly number.
 * n does not exceed 1690.
 *
 */

// @lc code=start
func nthUglyNumber(n int) int {
	return nthUglyNumber1(n)
}

func nthUglyNumber1(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	twoStep, threeStep, fiveStep := 0, 0, 0
	for i := 1; i < n; i++ {
		twoVal := dp[twoStep] * 2
		threeVal := dp[threeStep] * 3
		fiveVal := dp[fiveStep] * 5

		dp[i] = min(twoVal, threeVal, fiveVal)
		if dp[i] == twoVal {
			twoStep++
		}
		if dp[i] == threeVal {
			threeStep++
		}
		if dp[i] == fiveVal {
			fiveStep++
		}
	}

	return dp[n-1]
}

func min(args ...int) int {
	min := math.MaxInt32
	for _, v := range args {
		if v < min {
			min = v
		}
	}
	return min
}

// @lc code=end