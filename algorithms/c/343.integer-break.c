#include <stdlib.h>
/*
 * @lc app=leetcode id=343 lang=c
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

int max(int val1, int val2);

int integerBreak(int n)
{
	if (n < 2)
	{
		return 0;
	}
	else if (n == 2)
	{
		return 1;
	}
	else if (n == 3)
	{
		return 2;
	}

	int *dp = (int *)malloc(sizeof(int) * (n + 1));
	dp[0] = 0;
	dp[1] = 0;
	dp[2] = 1;
	dp[3] = 2;

	for (int i = 4; i <= n; i++)
	{
		int curMax = 0;
		for (int j = 1; j <= i / 2; j++)
		{
			curMax = max(curMax, max(j * (i - j), j * dp[i - j]));
		}
		dp[i] = curMax;
	}

	return dp[n];
}

int max(int val1, int val2)
{
	if (val1 > val2)
	{
		return val1;
	}
	return val2;
}

// @lc code=end