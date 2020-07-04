import (
	"sort"
)

/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 *
 * https://leetcode.com/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (45.87%)
 * Likes:    2074
 * Dislikes: 140
 * Total Accepted:    471.1K
 * Total Submissions: 1M
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * Given an array nums of n integers and an integer target, find three integers
 * in nums such that the sum is closest to target. Return the sum of the three
 * integers. You may assume that each input would have exactly one solution.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [-1,2,1,-4], target = 1
 * Output: 2
 * Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 =
 * 2).
 *
 *
 *
 * Constraints:
 *
 *
 * 3 <= nums.length <= 10^3
 * -10^3 <= nums[i] <= 10^3
 * -10^4 <= target <= 10^4
 *
 *
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	diff := target - nums[0] - nums[1] - nums[2]
	for i := 0; i < len(nums); i++ {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if abs(target-sum) < abs(diff) {
				diff = target - sum
			}
			if sum > target {
				k--
				for j < k && nums[k] == nums[k+1] { // 重复值
					k--
				}
			} else if sum < target {
				j++
				for j < k && nums[j-1] == nums[j] { // 重复值
					j++
				}
			} else {
				return sum
			}
		}
	}
	return target - diff
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

// @lc code=end

