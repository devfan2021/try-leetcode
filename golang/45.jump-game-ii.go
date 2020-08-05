/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 *
 * https://leetcode.com/problems/jump-game-ii/description/
 *
 * algorithms
 * Hard (30.30%)
 * Likes:    2658
 * Dislikes: 140
 * Total Accepted:    265.4K
 * Total Submissions: 869.4K
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * Given an array of non-negative integers, you are initially positioned at the
 * first index of the array.
 *
 * Each element in the array represents your maximum jump length at that
 * position.
 *
 * Your goal is to reach the last index in the minimum number of jumps.
 *
 * Example:
 *
 *
 * Input: [2,3,1,1,4]
 * Output: 2
 * Explanation: The minimum number of jumps to reach the last index is 2.
 * ⁠   Jump 1 step from index 0 to 1, then 3 steps to the last index.
 *
 * Note:
 *
 * You can assume that you can always reach the last index.
 *
 */

// @lc code=start
// [0], [1,2,3]
func jump(nums []int) int {
	return jump2(nums)
}

// search from positive direction. time complexity:O(n), space complexity:O(1)
func jump2(nums []int) int {
	step, end, length := 0, 0, len(nums)
	maxPostion := 0
	for i := 0; i < length-1; i++ {
		maxPostion = max(maxPostion, i+nums[i])
		if i == end { // perfect
			end = maxPostion
			step++
		}
	}
	return step
}

// greedy method, search from reverse direction. time complexity:O(n*n), space complexity:O(1)
func jump1(nums []int) int {
	step, position := 0, len(nums)-1
	for position > 0 {
		for i := 0; i < position; i++ {
			if (i + nums[i]) >= position {
				position = i
				step++
				break
			}
		}
	}
	return step
}

func max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

// @lc code=end