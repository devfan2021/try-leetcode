/*
 * @lc app=leetcode id=344 lang=golang
 *
 * [344] Reverse String
 *
 * https://leetcode.com/problems/reverse-string/description/
 *
 * algorithms
 * Easy (67.94%)
 * Likes:    1422
 * Dislikes: 682
 * Total Accepted:    753.6K
 * Total Submissions: 1.1M
 * Testcase Example:  '["h","e","l","l","o"]'
 *
 * Write a function that reverses a string. The input string is given as an
 * array of characters char[].
 *
 * Do not allocate extra space for another array, you must do this by modifying
 * the input array in-place with O(1) extra memory.
 *
 * You may assume all the characters consist of printable ascii
 * characters.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: ["h","e","l","l","o"]
 * Output: ["o","l","l","e","h"]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: ["H","a","n","n","a","h"]
 * Output: ["h","a","n","n","a","H"]
 *
 *
 *
 *
 */

// @lc code=start
func reverseString(s []byte) {
	reverseString2(s)
}

// solution by recursive function：Time complexity: O(n), Space complexity: O(n)
func reverseString2(s []byte) {
	helper(s, 0, len(s)-1)
}

func helper(s []byte, start, end int) {
	if start >= end {
		return
	}
	s[start], s[end] = s[end], s[start]
	helper(s, start+1, end-1)
}

// two pointers, Iteration, exchange first and last element: Time complexity: O(n), Space complexity: O(1)
func reverseString1(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left, right = left+1, right-1
	}
}

// @lc code=end