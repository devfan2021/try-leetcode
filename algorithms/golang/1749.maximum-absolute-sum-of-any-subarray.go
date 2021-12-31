/*
 * @lc app=leetcode id=1749 lang=golang
 *
 * [1749] Maximum Absolute Sum of Any Subarray
 *
 * https://leetcode.com/problems/maximum-absolute-sum-of-any-subarray/description/
 *
 * algorithms
 * Medium (55.98%)
 * Likes:    573
 * Dislikes: 9
 * Total Accepted:    17.6K
 * Total Submissions: 31.4K
 * Testcase Example:  '[1,-3,2,3,-4]'
 *
 * You are given an integer array nums. The absolute sum of a subarray [numsl,
 * numsl+1, ..., numsr-1, numsr] is abs(numsl + numsl+1 + ... + numsr-1 +
 * numsr).
 *
 * Return the maximum absolute sum of any (possibly empty) subarray of nums.
 *
 * Note that abs(x) is defined as follows:
 *
 *
 * If x is a negative integer, then abs(x) = -x.
 * If x is a non-negative integer, then abs(x) = x.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,-3,2,3,-4]
 * Output: 5
 * Explanation: The subarray [2,3] has absolute sum = abs(2+3) = abs(5) = 5.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [2,-5,1,-4,3,-2]
 * Output: 8
 * Explanation: The subarray [-5,1,-4] has absolute sum = abs(-5+1-4) = abs(-8)
 * = 8.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 *
 *
 */
// [-7,-1,0,-2,1,3,8,-2,-6,-1,-10,-6,-6,8,-4,-9,-4,1,4,-9]
// @lc code=start
func maxAbsoluteSum(nums []int) int {
	return solution2(nums)
}

func solution1(nums []int) int {
	bestMaxSum, curMaxSum := 0, 0
	for _, val := range nums {
		curMaxSum = maxVal(curMaxSum+val, val)
		bestMaxSum = maxVal(bestMaxSum, curMaxSum)
	}

	bestMinSum, curMinSum := 0, 0
	for _, val := range nums {
		curMinSum = minVal(curMinSum+val, val)
		bestMinSum = minVal(bestMinSum, curMinSum)
	}

	if bestMinSum < 0 {
		if abs(bestMinSum) > bestMaxSum {
			return abs(bestMinSum)
		}
	}
	return bestMaxSum
}

func solution2(nums []int) int {
	result, curMaxSum, curMinSum := 0, 0, 0
	for _, val := range nums {
		curMaxSum = maxVal(curMaxSum+val, val)
		curMinSum = minVal(curMinSum+val, val)
		result = maxVal(result, maxVal(abs(curMaxSum), abs(curMinSum)))
	}
	return result
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func maxVal(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

func minVal(val1, val2 int) int {
	if val1 < val2 {
		return val1
	}
	return val2
}

// @lc code=end