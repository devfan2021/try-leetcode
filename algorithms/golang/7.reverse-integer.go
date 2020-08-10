import "math"

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 *
 * https://leetcode.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (25.75%)
 * Likes:    3516
 * Dislikes: 5529
 * Total Accepted:    1.2M
 * Total Submissions: 4.5M
 * Testcase Example:  '123'
 *
 * Given a 32-bit signed integer, reverse digits of an integer.
 *
 * Example 1:
 *
 *
 * Input: 123
 * Output: 321
 *
 *
 * Example 2:
 *
 *
 * Input: -123
 * Output: -321
 *
 *
 * Example 3:
 *
 *
 * Input: 120
 * Output: 21
 *
 *
 * Note:
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose
 * of this problem, assume that your function returns 0 when the reversed
 * integer overflows.
 *
 */

// @lc code=start
func reverse(x int) int {
	return reverse2(x)
}

func reverse2(x int) int {
	const (
		MaxInt = 1<<31 - 1
		MinInt = -1 << 31
	)
	retVal := 0
	for x != 0 {
		retVal = retVal*10 + x%10
		x = x / 10
	}
	if retVal > MaxInt || retVal < MinInt {
		return 0
	}
	return retVal
}

// 1534236469
func reverse1(x int) int {
	retVal := 0
	for x != 0 {
		retVal = retVal*10 + x%10
		x = x / 10
	}
	if retVal > math.MaxInt32 || retVal < math.MinInt32 {
		return 0
	}
	return retVal
}

// @lc code=end