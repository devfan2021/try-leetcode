/*
 * @lc app=leetcode id=796 lang=golang
 *
 * [796] Rotate String
 *
 * https://leetcode.com/problems/rotate-string/description/
 *
 * algorithms
 * Easy (50.73%)
 * Likes:    1620
 * Dislikes: 79
 * Total Accepted:    131.9K
 * Total Submissions: 258.1K
 * Testcase Example:  '"abcde"\n"cdeab"'
 *
 * Given two strings s and goal, return true if and only if s can become goal
 * after some number of shifts on s.
 *
 * A shift on s consists of moving the leftmost character of s to the rightmost
 * position.
 *
 *
 * For example, if s = "abcde", then it will be "bcdea" after one shift.
 *
 *
 *
 * Example 1:
 * Input: s = "abcde", goal = "cdeab"
 * Output: true
 * Example 2:
 * Input: s = "abcde", goal = "abced"
 * Output: false
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length, goal.length <= 100
 * s and goal consist of lowercase English letters.
 *
 *
 */
// @lc code=start
func rotateString(s string, goal string) bool {
	return solution2(s, goal)
}

func solution1(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	if s == goal {
		return true
	}

	for i := 0; i <= len(s); i++ {
		if strings.Compare(s[i:]+s[:i], goal) == 0 {
			return true
		}
	}
	return false
}

func solution2(s string, goal string) bool {
	return len(s) == len(goal) && strings.Contains(s+s, goal)
}

// @lc code=end