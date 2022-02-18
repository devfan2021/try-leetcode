/*
 * @lc app=leetcode id=1941 lang=golang
 *
 * [1941] Check if All Characters Have Equal Number of Occurrences
 *
 * https://leetcode.com/problems/check-if-all-characters-have-equal-number-of-occurrences/description/
 *
 * algorithms
 * Easy (76.77%)
 * Likes:    277
 * Dislikes: 9
 * Total Accepted:    29.1K
 * Total Submissions: 37.8K
 * Testcase Example:  '"abacbc"'
 *
 * Given a string s, return true if s is a good string, or false otherwise.
 *
 * A string s is good if all the characters that appear in s have the same
 * number of occurrences (i.e., the same frequency).
 *
 *
 * Example 1:
 *
 *
 * Input: s = "abacbc"
 * Output: true
 * Explanation: The characters that appear in s are 'a', 'b', and 'c'. All
 * characters occur 2 times in s.
 *
 *
 * Example 2:
 *
 *
 * Input: s = "aaabb"
 * Output: false
 * Explanation: The characters that appear in s are 'a' and 'b'.
 * 'a' occurs 3 times while 'b' occurs 2 times, which is not the same number of
 * times.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 1000
 * s consists of lowercase English letters.
 *
 *
 */

// @lc code=start
func areOccurrencesEqual(s string) bool {
	hashMap := make(map[rune]int)
	for _, v := range s {
		hashMap[v] = hashMap[v] + 1
	}

	count := -1
	for _, v := range hashMap {
		if count == -1 {
			count = v
		}
		if count != v {
			return false
		}
	}
	return true
}

// @lc code=end