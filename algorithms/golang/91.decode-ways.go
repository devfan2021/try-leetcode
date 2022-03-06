/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 *
 * https://leetcode.com/problems/decode-ways/description/
 *
 * algorithms
 * Medium (29.18%)
 * Likes:    6439
 * Dislikes: 3742
 * Total Accepted:    740.7K
 * Total Submissions: 2.5M
 * Testcase Example:  '"12"'
 *
 * A message containing letters from A-Z can be encoded into numbers using the
 * following mapping:
 *
 *
 * 'A' -> "1"
 * 'B' -> "2"
 * ...
 * 'Z' -> "26"
 *
 *
 * To decode an encoded message, all the digits must be grouped then mapped
 * back into letters using the reverse of the mapping above (there may be
 * multiple ways). For example, "11106" can be mapped into:
 *
 *
 * "AAJF" with the grouping (1 1 10 6)
 * "KJF" with the grouping (11 10 6)
 *
 *
 * Note that the grouping (1 11 06) is invalid because "06" cannot be mapped
 * into 'F' since "6" is different from "06".
 *
 * Given a string s containing only digits, return the number of ways to decode
 * it.
 *
 * The test cases are generated so that the answer fits in a 32-bit integer.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "12"
 * Output: 2
 * Explanation: "12" could be decoded as "AB" (1 2) or "L" (12).
 *
 *
 * Example 2:
 *
 *
 * Input: s = "226"
 * Output: 3
 * Explanation: "226" could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2
 * 2 6).
 *
 *
 * Example 3:
 *
 *
 * Input: s = "06"
 * Output: 0
 * Explanation: "06" cannot be mapped to "F" because of the leading zero ("6"
 * is different from "06").
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 100
 * s contains only digits and may contain leading zero(s).
 *
 *
 */

// @lc code=start
func numDecodings(s string) int {
	return soluntion1(s)
}

/* case1: f(n) = f(n-1), when s[n] != 0
 * case2: f(n) = f(n-2), when s[n-1] != 0 and 10*s[n-1]+s[n] <= 26
 */
func soluntion1(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

func soluntion2(s string) int {
	n := len(s)
	// a=f(n-2), b=f(n-1), c=f(n)
	a, b, c := 0, 1, 0
	for i := 1; i <= n; i++ {
		c = 0
		if s[i-1] != '0' {
			c += b
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
			c += a
		}
		a, b = b, c
	}
	return c
}

// @lc code=end