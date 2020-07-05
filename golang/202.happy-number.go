/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 *
 * https://leetcode.com/problems/happy-number/description/
 *
 * algorithms
 * Easy (50.13%)
 * Likes:    2120
 * Dislikes: 422
 * Total Accepted:    509.4K
 * Total Submissions: 1M
 * Testcase Example:  '19'
 *
 * Write an algorithm to determine if a number n is "happy".
 *
 * A happy number is a number defined by the following process: Starting with
 * any positive integer, replace the number by the sum of the squares of its
 * digits, and repeat the process until the number equals 1 (where it will
 * stay), or it loops endlessly in a cycle which does not include 1. Those
 * numbers for which this process ends in 1 are happy numbers.
 *
 * Return True if n is a happy number, and False if not.
 *
 * Example:
 *
 *
 * Input: 19
 * Output: true
 * Explanation:
 * 1^2 + 9^2 = 82
 * 8^2 + 2^2 = 68
 * 6^2 + 8^2 = 100
 * 1^2 + 0^2 + 0^2 = 1
 *
 *
 */

// @lc code=start
func isHappy(n int) bool {
	valMap := make(map[int]bool)
	valMap[n] = true
	for n != 1 {
		sum := 0
		for n >= 10 {
			sum += (n % 10) * (n % 10)
			n = n / 10
		}
		sum += (n % 10) * (n % 10)
		n = sum
		if valMap[n] {
			return false
		}
		valMap[n] = true
	}
	return true
}

// @lc code=end