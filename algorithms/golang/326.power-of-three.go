/*
 * @lc app=leetcode id=326 lang=golang
 *
 * [326] Power of Three
 *
 * https://leetcode.com/problems/power-of-three/description/
 *
 * algorithms
 * Easy (42.99%)
 * Likes:    791
 * Dislikes: 101
 * Total Accepted:    439.2K
 * Total Submissions: 1M
 * Testcase Example:  '27'
 *
 * Given an integer n, return true if it is a power of three. Otherwise, return
 * false.
 *
 * An integer n is a power of three, if there exists an integer x such that n
 * == 3^x.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 27
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: n = 0
 * Output: false
 *
 *
 * Example 3:
 *
 *
 * Input: n = 9
 * Output: true
 *
 *
 *
 * Constraints:
 *
 *
 * -2^31 <= n <= 2^31 - 1
 *
 *
 *
 * Follow up: Could you solve it without loops/recursion?
 */

// @lc code=start
func isPowerOfThree(n int) bool {
	return soluntion2()
}

// Loop Iteration
func soluntion1(n int) bool {
	if n < 1 {
		return false
	}
	// x is zero
	if n == 1 {
		return true
	}
	for n > 3 {
		if n%3 != 0 {
			return false
		}
		n = n / 3
	}
	if n == 3 {
		return true
	}
	return false
}

func soluntion2(n int) bool {
	if n < 1 {
		return false
	}
	for n%3 == 0 {
		n = n / 3
	}
	// x is zero
	return n == 1
}

// Integer Limitations
func soluntion3(n int) bool {
	// 1162261467 is 3^19,  3^20 is bigger than int
	return n > 0 && 1162261467%n == 0
}

// @lc code=end