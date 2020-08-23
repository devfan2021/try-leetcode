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
	return longestPalindrome2(s)
}

// dp:  p(i, j) = p(i+1, j-1) ^ (Si = Sj)
func longestPalindrome2(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}

	maxLen, begin, dp := 1, 0, make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
		dp[i][i] = true
	}

	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

// brute force, time complexity:O(n^3)
func longestPalindrome1(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}

	begin, maxLength := 0, 1
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if j-i+1 > maxLength && validPalindromic(s, i, j) {
				maxLength = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLength]
}

func validPalindromic(s string, begin, end int) bool {
	for begin < end {
		if s[begin] != s[end] {
			return false
		}
		begin++
		end--
	}
	return true
}

// @lc code=end