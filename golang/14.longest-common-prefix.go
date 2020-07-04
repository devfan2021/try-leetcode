import "fmt"

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 *
 * https://leetcode.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (35.21%)
 * Likes:    2519
 * Dislikes: 1840
 * Total Accepted:    747.6K
 * Total Submissions: 2.1M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * Write a function to find the longest common prefix string amongst an array
 * of strings.
 *
 * If there is no common prefix, return an empty string "".
 *
 * Example 1:
 *
 *
 * Input: ["flower","flow","flight"]
 * Output: "fl"
 *
 *
 * Example 2:
 *
 *
 * Input: ["dog","racecar","car"]
 * Output: ""
 * Explanation: There is no common prefix among the input strings.
 *
 *
 * Note:
 *
 * All given inputs are in lowercase letters a-z.
 *
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {

}

// not efficiency
func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	smallLen := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < smallLen {
			smallLen = len(strs[i])
		}
	}

	if smallLen == 0 {
		return ""
	}

	startIndex := 0
	for startIndex < smallLen {
		if isMatch(strs, startIndex) {
			startIndex++
		} else {
			break
		}
	}
	fmt.Println(startIndex)
	if startIndex > 0 {
		return strs[0][:startIndex]
	}
	return ""
}

func isMatch(strs []string, startIndex int) bool {
	cmpVal := strs[0][:startIndex+1]
	for i := 1; i < len(strs); i++ {
		if cmpVal != strs[i][:startIndex+1] {
			return false
		}
	}
	return true
}

// @lc code=end

