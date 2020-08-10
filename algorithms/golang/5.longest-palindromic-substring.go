/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 *
 * https://leetcode.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (29.26%)
 * Likes:    7334
 * Dislikes: 554
 * Total Accepted:    979.1K
 * Total Submissions: 3.3M
 * Testcase Example:  '"babad"'
 *
 * Given a string s, find the longest palindromic substring in s. You may
 * assume that the maximum length of s is 1000.
 *
 * Example 1:
 *
 *
 * Input: "babad"
 * Output: "bab"
 * Note: "aba" is also a valid answer.
 *
 *
 * Example 2:
 *
 *
 * Input: "cbbd"
 * Output: "bb"
 *
 *
 */

// @lc code=start
func longestPalindrome(s string) string {
	return longestPalindrome1(s)
}

// dp:  p(i, j) = p(i+1, j-1) ^ (Si = Sj)
func longestPalindrome1(s string) string {
	retVals, length, dp := "", len(s), make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {

		}
	}
}

// @lc code=end