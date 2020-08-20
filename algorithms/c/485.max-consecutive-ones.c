/*
 * @lc app=leetcode id=485 lang=c
 *
 * [485] Max Consecutive Ones
 *
 * https://leetcode.com/problems/max-consecutive-ones/description/
 *
 * algorithms
 * Easy (55.19%)
 * Likes:    631
 * Dislikes: 360
 * Total Accepted:    223.2K
 * Total Submissions: 404.6K
 * Testcase Example:  '[1,0,1,1,0,1]'
 *
 * Given a binary array, find the maximum number of consecutive 1s in this
 * array.
 *
 * Example 1:
 *
 * Input: [1,1,0,1,1,1]
 * Output: 3
 * Explanation: The first two digits or the last three digits are consecutive
 * 1s.
 * ⁠   The maximum number of consecutive 1s is 3.
 *
 *
 *
 * Note:
 *
 * The input array will only contain 0 and 1.
 * The length of input array is a positive integer and will not exceed 10,000
 *
 *
 */

// @lc code=start
int findMaxConsecutiveOnes(int *nums, int numsSize)
{
	int max_count = 0, cur_count = 0;
	for (int i = 0; i < numsSize; i++)
	{
		if (nums[i] == 1)
		{
			cur_count++;
		}
		else
		{
			cur_count = 0;
		}

		if (cur_count > max_count)
		{
			max_count = cur_count;
		}
	}
	return max_count;
}

// @lc code=end