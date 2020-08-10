/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 *
 * https://leetcode.com/problems/plus-one/description/
 *
 * algorithms
 * Easy (42.18%)
 * Likes:    1466
 * Dislikes: 2331
 * Total Accepted:    592.8K
 * Total Submissions: 1.4M
 * Testcase Example:  '[1,2,3]'
 *
 * Given a non-empty array of digits representing a non-negative integer, plus
 * one to the integer.
 *
 * The digits are stored such that the most significant digit is at the head of
 * the list, and each element in the array contain a single digit.
 *
 * You may assume the integer does not contain any leading zero, except the
 * number 0 itself.
 *
 * Example 1:
 *
 *
 * Input: [1,2,3]
 * Output: [1,2,4]
 * Explanation: The array represents the integer 123.
 *
 *
 * Example 2:
 *
 *
 * Input: [4,3,2,1]
 * Output: [4,3,2,2]
 * Explanation: The array represents the integer 4321.
 *
 */

// @lc code=start
func plusOne(digits []int) []int {
	return plusOne2(digits)
}

func plusOne2(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		// 巧妙的进行了加1操作
		if digits[i] < 9 {
			digits[i] = digits[i] + 1
			return digits
		} else {
			digits[i] = 0
		}
	}

	retVal := make([]int, len(digits)+1)
	retVal[0] = 1
	return retVal
}

// 解法有点啰嗦，对每个位都进行了curry位的添加，其实可以进行简化
func plusOne1(digits []int) []int {
	retVal := make([]int, len(digits)+1, len(digits)+1)
	curry := 0
	for i := len(digits) - 1; i >= 0; i-- {
		val := 0
		if i == len(digits)-1 {
			val = digits[i] + 1 + curry
		} else {
			val = digits[i] + curry
		}
		if val >= 10 {
			retVal[i+1] = val - 10
			curry = 1
		} else {
			retVal[i+1] = val
			curry = 0
		}
	}
	if curry == 1 {
		retVal[0] = 1
		return retVal
	} else {
		return retVal[1:len(retVal)]
	}
}

// @lc code=end

