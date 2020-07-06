/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 *
 * https://leetcode.com/problems/first-unique-character-in-a-string/description/
 *
 * algorithms
 * Easy (53.18%)
 * Likes:    1911
 * Dislikes: 116
 * Total Accepted:    535.5K
 * Total Submissions: 1M
 * Testcase Example:  '"leetcode"'
 *
 * Given a string, find the first non-repeating character in it and return its
 * index. If it doesn't exist, return -1.
 *
 * Examples:
 *
 *
 * s = "leetcode"
 * return 0.
 *
 * s = "loveleetcode"
 * return 2.
 *
 *
 *
 *
 * Note: You may assume the string contains only lowercase English letters.
 *
 */

// @lc code=start
func firstUniqChar(s string) int {
	return firstUniqChar2(s)
}

func firstUniqChar2(s string) int {
	hash := make([]int, 26, 26)
	for _, v := range s {
		hash[v-'a']++
	}
	for i, c := range s {
		if hash[c-'a'] == 1 {
			return i
		}
	}
	return -1
}

func firstUniqChar1(s string) int {
	valMap := make(map[byte]int)
	for _, v := range s {
		valMap[byte(v)] = valMap[byte(v)] + 1
	}

	for i := 0; i < len(s); i++ {
		val, _ := valMap[s[i]]
		if val == 1 {
			return i
		}
	}
	return -1
}

// @lc code=end