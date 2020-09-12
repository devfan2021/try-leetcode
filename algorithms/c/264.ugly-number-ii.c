#include <stdlib.h>

/*
 * @lc app=leetcode id=264 lang=c
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
#define min(a, b) (a > b ? b : a)

int nthUglyNumber(int n)
{
	int *dp = malloc(sizeof(int) * n);
	int twoStep = 0, threeStep = 0, fiveStep = 0;
	dp[0] = 1;
	for (int i = 1; i < n; i++)
	{
		int twoVal = dp[twoStep] * 2;
		int threeVal = dp[threeStep] * 3;
		int fiveVal = dp[fiveStep] * 5;
		int minVal = min(min(twoVal, threeVal), fiveVal);

		if (minVal == twoVal)
			twoStep++;
		if (minVal == threeVal)
			threeStep++;
		if (minVal == fiveVal)
			fiveStep++;

		dp[i] = minVal;
	}

	return dp[n - 1];
}

// @lc code=end