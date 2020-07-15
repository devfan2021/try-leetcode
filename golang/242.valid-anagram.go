/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 *
 * https://leetcode.com/problems/valid-anagram/description/
 *
 * algorithms
 * Easy (56.37%)
 * Likes:    1543
 * Dislikes: 143
 * Total Accepted:    575.4K
 * Total Submissions: 1M
 * Testcase Example:  '"anagram"\n"nagaram"'
 *
 * Given two strings s and t , write a function to determine if t is an anagram
 * of s.
 *
 * Example 1:
 *
 *
 * Input: s = "anagram", t = "nagaram"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: s = "rat", t = "car"
 * Output: false
 *
 *
 * Note:
 * You may assume the string contains only lowercase alphabets.
 *
 * Follow up:
 * What if the inputs contain unicode characters? How would you adapt your
 * solution to such case?
 *
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	return isAnagram2(s, t)
}

// 因为只有小写字母，因此只需要建一个包含26个int的数组进行比较
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make([]int, 26)
	for i := 0; i < len(s); i++ {
		count[int(s[i]-'a')]++
		count[int(t[i]-'a')]--
	}
	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}

// 用一个hash统计每个字母的数字
func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := map[byte]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}
	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}

// @lc code=end