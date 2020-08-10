/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 *
 * https://leetcode.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (46.23%)
 * Likes:    8293
 * Dislikes: 397
 * Total Accepted:    1.1M
 * Total Submissions: 2.3M
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * Given an integer array nums, find the contiguous subarray (containing at
 * least one number) which has the largest sum and return its sum.
 *
 * Example:
 *
 *
 * Input: [-2,1,-3,4,-1,2,1,-5,4],
 * Output: 6
 * Explanation: [4,-1,2,1] has the largest sum = 6.
 *
 *
 * Follow up:
 *
 * If you have figured out the O(n) solution, try coding another solution using
 * the divide and conquer approach, which is more subtle.
 *
 */

// @lc code=start
func maxSubArray(nums []int) int {
	return maxSubArray2(nums)
}

// dp, method is clear
func maxSubArray2(nums []int) int {
	size := len(nums)
	dp := make([]int, size)
	dp[0] = nums[0]
	maxSum := dp[0]
	for i := 1; i < size; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		maxSum = max(max, dp[i])
	}
	return maxSum
}

func maxSubArray1(nums []int) int {
	preSum, maxSum := 0, nums[0]
	for i := 0; i < len(nums); i++ {
		preSum = max(preSum+nums[i], nums[i])
		maxSum = max(maxSum, preSum)
	}
	return maxSum
}

func max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

// @lc code=end