/*
 * @lc app=leetcode id=1446 lang=golang
 *
 * [1446] Consecutive Characters
 *
 * https://leetcode.com/problems/consecutive-characters/description/
 *
 * algorithms
 * Easy (62.06%)
 * Likes:    1159
 * Dislikes: 23
 * Total Accepted:    113.7K
 * Total Submissions: 183.4K
 * Testcase Example:  '"leetcode"'
 *
 * The power of the string is the maximum length of a non-empty substring that
 * contains only one unique character.
 *
 * Given a string s, return the power of s.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "leetcode"
 * Output: 2
 * Explanation: The substring "ee" is of length 2 with the character 'e'
 * only.
 *
 *
 * Example 2:
 *
 *
 * Input: s = "abbcccddddeeeeedcba"
 * Output: 5
 * Explanation: The substring "eeeee" is of length 5 with the character 'e'
 * only.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 500
 * s consists of only lowercase English letters.
 *
 *
 */

// @lc code=start
func maxPower(s string) int {
	return solution1(s)
}

func solution1(s string) int {
	max, cMax, cFirst := 0, 0, rune(s[0])
	for _, v := range s {
		if cFirst == v {
			cMax += 1
		} else {
			if cMax > max {
				max = cMax
			}
			cFirst = v
			cMax = 1
		}
	}
	if cMax > max {
		return cMax
	}
	return max
}

// @lc code=end