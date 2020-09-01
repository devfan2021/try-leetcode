/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 *
 * https://leetcode.com/problems/divide-two-integers/description/
 *
 * algorithms
 * Medium (16.42%)
 * Likes:    1279
 * Dislikes: 5497
 * Total Accepted:    293.6K
 * Total Submissions: 1.8M
 * Testcase Example:  '10\n3'
 *
 * Given two integers dividend and divisor, divide two integers without using
 * multiplication, division and mod operator.
 *
 * Return the quotient after dividing dividend by divisor.
 *
 * The integer division should truncate toward zero, which means losing its
 * fractional part. For example, truncate(8.345) = 8 and truncate(-2.7335) =
 * -2.
 *
 * Example 1:
 *
 *
 * Input: dividend = 10, divisor = 3
 * Output: 3
 * Explanation: 10/3 = truncate(3.33333..) = 3.
 *
 *
 * Example 2:
 *
 *
 * Input: dividend = 7, divisor = -3
 * Output: -2
 * Explanation: 7/-3 = truncate(-2.33333..) = -2.
 *
 *
 * Note:
 *
 *
 * Both dividend and divisor will be 32-bit signed integers.
 * The divisor will never be 0.
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose
 * of this problem, assume that your function returns 2^31 − 1 when the
 * division result overflows.
 *
 *
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	return divide2(dividend, divisor)
}

// test case: 1\n-1
func divide2(dividend int, divisor int) int {
	minInt32, maxInt32 := -1<<31, 1<<31-1
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if dividend == minInt32 && divisor == -1 {
		return maxInt32
	}

	isPostive := true
	if dividend < 0 {
		dividend = -dividend
		isPostive = !isPostive
	}
	if divisor < 0 {
		divisor = -divisor
		isPostive = !isPostive
	}

	// divide some part to approach the answer
	quotient := div(dividend, divisor)
	if !isPostive {
		return -quotient
	}

	return quotient
}

// similar binary search
func div(val1 int, val2 int) int {
	if val1 < val2 {
		return 0
	}

	count, divisor := 1, val2
	for (divisor + divisor) < val1 {
		count += count     // double
		divisor += divisor // double
	}
	return count + div(val1-divisor, val2)
}

// test case: 1\n-1
func divide1(dividend int, divisor int) int {
	minInt32, maxInt32 := -1<<31, 1<<31-1
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if dividend == minInt32 && divisor == -1 {
		return maxInt32
	}

	isPostive := true
	if dividend < 0 {
		dividend = -dividend
		isPostive = !isPostive
	}
	if divisor < 0 {
		divisor = -divisor
		isPostive = !isPostive
	}

	// divide some part to approach the answer
	quotient := 0
	for dividend >= divisor { // todo: thinking, why for iterator
		multiple := divisor
		for i := 0; dividend >= multiple; i++ {
			dividend -= multiple
			quotient += 1 << uint(i)
			multiple <<= 1
		}
	}

	if !isPostive {
		return -quotient
	}

	return quotient
}

// @lc code=end