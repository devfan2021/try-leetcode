/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 *
 * https://leetcode.com/problems/house-robber/description/
 *
 * algorithms
 * Easy (41.85%)
 * Likes:    5095
 * Dislikes: 156
 * Total Accepted:    545.3K
 * Total Submissions: 1.3M
 * Testcase Example:  '[1,2,3,1]'
 *
 * You are a professional robber planning to rob houses along a street. Each
 * house has a certain amount of money stashed, the only constraint stopping
 * you from robbing each of them is that adjacent houses have security system
 * connected and it will automatically contact the police if two adjacent
 * houses were broken into on the same night.
 *
 * Given a list of non-negative integers representing the amount of money of
 * each house, determine the maximum amount of money you can rob tonight
 * without alerting the police.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,2,3,1]
 * Output: 4
 * Explanation: Rob house 1 (money = 1) and then rob house 3 (money =
 * 3).
 * Total amount you can rob = 1 + 3 = 4.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [2,7,9,3,1]
 * Output: 12
 * Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house
 * 5 (money = 1).
 * Total amount you can rob = 2 + 9 + 1 = 12.
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= nums.length <= 100
 * 0 <= nums[i] <= 400
 *
 *
 */

// @lc code=start
func rob(nums []int) int {
	return rob2(nums)
}

// dynamic program dp[k] = max(dp[k-2]+num[k], dp[k-1])
func rob3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		tmp := second
		second = max(first+nums[i], second)
		first = tmp
	}
	return second
}

// recursive with memory
func rob2(nums []int) int {
	hash := map[int]int{}
	return helper2(nums, len(nums)-1, hash)
}

func helper2(nums []int, index int, hash map[int]int) int {
	if index < 0 {
		return 0
	}

	if v, ok := hash[index]; ok {
		return v
	}

	retVal := max(helper2(nums, index-2, hash)+nums[index], helper2(nums, index-1, hash))
	hash[index] = retVal
	return retVal
}

// recursive, Time Limit Exceeded
func rob1(nums []int) int {
	return helper(nums, len(nums)-1)
}

func helper(nums []int, index int) int {
	if index < 0 {
		return 0
	}
	return max(helper(nums, index-2)+nums[index], helper(nums, index-1))
}

func max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

// @lc code=end