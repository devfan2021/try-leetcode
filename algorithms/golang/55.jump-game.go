/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 *
 * https://leetcode.com/problems/jump-game/description/
 *
 * algorithms
 * Medium (34.66%)
 * Likes:    4655
 * Dislikes: 350
 * Total Accepted:    507.2K
 * Total Submissions: 1.5M
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * Given an array of non-negative integers, you are initially positioned at the
 * first index of the array.
 *
 * Each element in the array represents your maximum jump length at that
 * position.
 *
 * Determine if you are able to reach the last index.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [2,3,1,1,4]
 * Output: true
 * Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last
 * index.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [3,2,1,0,4]
 * Output: false
 * Explanation: You will always arrive at index 3 no matter what. Its maximum
 * jump length is 0, which makes it impossible to reach the last index.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 3 * 10^4
 * 0 <= nums[i][j] <= 10^5
 *
 *
 */

// @lc code=start
func canJump(nums []int) bool {
	return canJump2(nums)
}

// Greedy
func canJump2(nums []int) bool {
	lastPos := len(nums) - 1
	for i := lastPos; i >= 0; i-- {
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}
	return lastPos == 0
}

// Time Limit Exceeded
// backtracking
func canJump1(nums []int) bool {
	return canJumpFromPosition(0, nums)
}

func canJumpFromPosition(position int, nums []int) bool {
	if position == len(nums)-1 {
		return true
	}

	furthestJump := min(position+nums[position], len(nums)-1)
	// from the max step
	for nextPosition := furthestJump; nextPosition > position; nextPosition-- {
		if canJumpFromPosition(nextPosition, nums) {
			return true
		}
	}

	return false
}

func min(val1, val2 int) int {
	if val1 < val2 {
		return val1
	}

	return val2
}

// @lc code=end