/*
 * @lc app=leetcode id=485 lang=golang
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
func findMaxConsecutiveOnes(nums []int) int {
	return findMaxConsecutiveOnes2(nums)
}

func findMaxConsecutiveOnes2(nums []int) int {
	maxCount, currentCount := 0, 0

	for _, num := range nums {
		if num == 1 {
			currentCount++
		} else {
			currentCount = 0
		}

		if currentCount > maxCount {
			maxCount = currentCount
		}
	}

	return maxCount
}

func findMaxConsecutiveOnes1(nums []int) int {
	firstLen, secondLen := 0, 0
	for _, item := range nums {
		if item == 1 {
			secondLen++
		} else {
			if secondLen > firstLen {
				firstLen = secondLen
			}
			secondLen = 0
		}
	}
	if secondLen > firstLen {
		return secondLen
	}
	return firstLen
}

// @lc code=end