/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 *
 * https://leetcode.com/problems/implement-strstr/description/
 *
 * algorithms
 * Easy (34.27%)
 * Likes:    1558
 * Dislikes: 1888
 * Total Accepted:    662.1K
 * Total Submissions: 1.9M
 * Testcase Example:  '"hello"\n"ll"'
 *
 * Implement strStr().
 *
 * Return the index of the first occurrence of needle in haystack, or -1 if
 * needle is not part of haystack.
 *
 * Example 1:
 *
 *
 * Input: haystack = "hello", needle = "ll"
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: haystack = "aaaaa", needle = "bba"
 * Output: -1
 *
 *
 * Clarification:
 *
 * What should we return when needle is an empty string? This is a great
 * question to ask during an interview.
 *
 * For the purpose of this problem, we will return 0 when needle is an empty
 * string. This is consistent to C's strstr() and Java's indexOf().
 *
 */

// @lc code=start
// 充分利用了slice的直接比较， 54行
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	length := len(needle)
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:i+length] == needle {
			return i
		}
	}
	return -1
}

// @lc code=end

