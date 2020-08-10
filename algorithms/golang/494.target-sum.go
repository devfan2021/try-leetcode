/*
 * @lc app=leetcode id=494 lang=golang
 *
 * [494] Target Sum
 *
 * https://leetcode.com/problems/target-sum/description/
 *
 * algorithms
 * Medium (46.56%)
 * Likes:    2620
 * Dislikes: 111
 * Total Accepted:    167.4K
 * Total Submissions: 360.4K
 * Testcase Example:  '[1,1,1,1,1]\n3'
 *
 * You are given a list of non-negative integers, a1, a2, ..., an, and a
 * target, S. Now you have 2 symbols + and -. For each integer, you should
 * choose one from + and - as its new symbol.
 *
 * Find out how many ways to assign symbols to make sum of integers equal to
 * target S.
 *
 * Example 1:
 *
 *
 * Input: nums is [1, 1, 1, 1, 1], S is 3.
 * Output: 5
 * Explanation:
 *
 * -1+1+1+1+1 = 3
 * +1-1+1+1+1 = 3
 * +1+1-1+1+1 = 3
 * +1+1+1-1+1 = 3
 * +1+1+1+1-1 = 3
 *
 * There are 5 ways to assign symbols to make the sum of nums be target 3.
 *
 *
 *
 * Constraints:
 *
 *
 * The length of the given array is positive and will not exceed 20.
 * The sum of elements in the given array will not exceed 1000.
 * Your output answer is guaranteed to be fitted in a 32-bit integer.
 *
 *
 */

// @lc code=start
func findTargetSumWays(nums []int, S int) int {
	return findTargetSumWays1(nums, S)
}

// brute force, time complexity:O(2^n) space complexity:O(n)
func findTargetSumWays1(nums []int, S int) int {
	count := 0
	calculate(nums, 0, 0, S, &count)
	return count
}

func calculate(nums []int, index, sum, target int, count *int) {
	if index == len(nums) {
		if sum == target {
			*count = *count + 1
		}
	} else {
		calculate(nums, index+1, sum+nums[index], target, count)
		calculate(nums, index+1, sum-nums[index], target, count)
	}
}

// @lc code=end