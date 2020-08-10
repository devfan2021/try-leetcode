import "math"

/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 *
 * https://leetcode.com/problems/minimum-size-subarray-sum/description/
 *
 * algorithms
 * Medium (37.69%)
 * Likes:    2304
 * Dislikes: 106
 * Total Accepted:    265.5K
 * Total Submissions: 702.4K
 * Testcase Example:  '7\n[2,3,1,2,4,3]'
 *
 * Given an array of n positive integers and a positive integer s, find the
 * minimal length of a contiguous subarray of which the sum ≥ s. If there isn't
 * one, return 0 instead.
 *
 * Example:
 *
 *
 * Input: s = 7, nums = [2,3,1,2,4,3]
 * Output: 2
 * Explanation: the subarray [4,3] has the minimal length under the problem
 * constraint.
 *
 * Follow up:
 *
 * If you have figured out the O(n) solution, try coding another solution of
 * which the time complexity is O(n log n).
 *
 */

// @lc code=start
// 方法暂时还没消化
func minSubArrayLen(s int, nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	ans := math.MaxInt32
	sum := 0
	start := 0

	for end := 0; end < len(nums); end++ {

		sum += nums[end]
		for sum >= s {
			ans = min(end-start+1, ans)
			sum -= nums[start]
			start++
		}
	}

	if ans == math.MaxInt32 {
		return 0
	}

	return ans
}

func min(i, j int) int {

	if i < j {
		return i
	}

	return j
}

// @lc code=end

