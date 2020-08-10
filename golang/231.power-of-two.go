/*
 * @lc app=leetcode id=231 lang=golang
 *
 * [231] Power of Two
 *
 * https://leetcode.com/problems/power-of-two/description/
 *
 * algorithms
 * Easy (43.62%)
 * Likes:    952
 * Dislikes: 198
 * Total Accepted:    359.2K
 * Total Submissions: 821.8K
 * Testcase Example:  '1'
 *
 * Given an integer, write a function to determine if it is a power of two.
 *
 * Example 1:
 *
 *
 * Input: 1
 * Output: true
 * Explanation: 2^0 = 1
 *
 *
 * Example 2:
 *
 *
 * Input: 16
 * Output: true
 * Explanation: 2^4 = 16
 *
 * Example 3:
 *
 *
 * Input: 218
 * Output: false
 *
 */

// @lc code=start
func isPowerOfTwo(n int) bool {
	return isPowerOfTwo3(n)
}

func isPowerOfTwo3(n int) bool {
	i := 1
	for i < n {
		i *= 2
	}
	return i == n
}

func isPowerOfTwo2(n int) bool {
	if n <= 0 {
		return false
	}
	for n&1 == 0 {
		n >>= 1
	}
	return n == 1
}

// & bit operation: n & (n-1), clear last 1 bit
// 0&0=0 1&0=0 0&1=0 1&1= 1
func isPowerOfTwo1(n int) bool {
	return n > 0 && n&(n-1) == 0
}

// @lc code=end