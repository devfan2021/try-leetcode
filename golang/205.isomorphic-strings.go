/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 *
 * https://leetcode.com/problems/isomorphic-strings/description/
 *
 * algorithms
 * Easy (39.54%)
 * Likes:    1366
 * Dislikes: 354
 * Total Accepted:    292K
 * Total Submissions: 737.2K
 * Testcase Example:  '"egg"\n"add"'
 *
 * Given two strings s and t, determine if they are isomorphic.
 *
 * Two strings are isomorphic if the characters in s can be replaced to get t.
 *
 * All occurrences of a character must be replaced with another character while
 * preserving the order of characters. No two characters may map to the same
 * character but a character may map to itself.
 *
 * Example 1:
 *
 *
 * Input: s = "egg", t = "add"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: s = "foo", t = "bar"
 * Output: false
 *
 * Example 3:
 *
 *
 * Input: s = "paper", t = "title"
 * Output: true
 *
 * Note:
 * You may assume both s and t have the same length.
 *
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	return isIsomorphic1(s, t)
}

func isIsomorphic1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap, tMap := map[uint8]int{}, map[uint8]int{}
	for index := range s {
		if sMap[s[index]] != tMap[t[index]] {
			return false
		}

		sMap[s[index]] = index + 1
		tMap[t[index]] = index + 1
	}
	return true
}

func isIsomorphic2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap, tMap := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if sMap[s[i]] != tMap[t[i]] {
			return false
		}

		sMap[s[i]] = i + 1
		tMap[t[i]] = i + 1
	}
	return true
}

// @lc code=end

