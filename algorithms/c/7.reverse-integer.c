#include <limits.h>
/*
 * @lc app=leetcode id=7 lang=c
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
int reverse(int x)
{
	long long retVal = 0;
	while (x != 0)
	{
		retVal = retVal * 10 + x % 10;
		x = x / 10;
	}

	return (retVal < INT_MIN || retVal > INT_MAX) ? 0 : retVal;
}

// @lc code=end