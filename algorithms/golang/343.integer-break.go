/*
 * @lc app=leetcode id=343 lang=golang
 *
 * [343] Integer Break
 *
 * https://leetcode.com/problems/integer-break/description/
 *
 * algorithms
 * Medium (50.44%)
 * Likes:    1146
 * Dislikes: 230
 * Total Accepted:    112.5K
 * Total Submissions: 222.7K
 * Testcase Example:  '2'
 *
 * Given a positive integer n, break it into the sum of at least two positive
 * integers and maximize the product of those integers. Return the maximum
 * product you can get.
 *
 * Example 1:
 *
 *
 *
 * Input: 2
 * Output: 1
 * Explanation: 2 = 1 + 1, 1 × 1 = 1.
 *
 *
 * Example 2:
 *
 *
 * Input: 10
 * Output: 36
 * Explanation: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.
 *
 * Note: You may assume that n is not less than 2 and not larger than 58.
 *
 *
 */

// @lc code=start
func integerBreak(n int) int {
	return integerBreak2(n)
}

// greedy: if n>=5, more cut 3, if n == 4,  2*2 > 1*3
func integerBreak3(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	timeOf3 := n / 3

	// check last 4 elements
	if n-timeOf3*3 == 1 {
		timeOf3 -= 1 // should last 4 element to 2 * 2
	}

	timeOf2 := (n - timeOf3*3) / 2

	retVal := 1
	for i := timeOf3; i > 0; i-- {
		retVal *= 3
	}

	for i := timeOf2; i > 0; i-- {
		retVal *= 2
	}
	return retVal
}

// dp, dp[i] = max(max(j*(i -j), j * dp[i-j])
func integerBreak2(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0
	dp[2] = 1
	dp[3] = 2

	for i := 4; i <= n; i++ {
		curMax := 0
		for j := 1; j <= i/2; j++ { // half of i: 1*3, 3*1
			curMax = max(curMax, max(j*(i-j), j*dp[i-j]))
		}
		dp[i] = curMax
	}

	return dp[n]
}

// dp, dp[i] = max(max(j*(i -j), j * dp[i-j])
func integerBreak1(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0

	for i := 2; i <= n; i++ {
		curMax := 0
		for j := 1; j < i; j++ {
			curMax = max(curMax, max(j*(i-j), j*dp[i-j]))
		}
		dp[i] = curMax
	}

	return dp[n]
}

func max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

// @lc code=end