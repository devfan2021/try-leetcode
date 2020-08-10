import "strings"

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
	return longestCommonPrefix2(strs)
}

// Divide and conquer
func longestCommonPrefix4(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	return longestCommonPrefix4(strs, 0, len(strs)-1)
}

func longestCommonPrefix4(strs []string, start, end int) string {
	if start == end {
		return strs[start]
	}

	mid := start + (end-start)/2
	leftPrefix := longestCommonPrefix4(strs, start, mid)
	rightPrefix := longestCommonPrefix4(strs, mid+1, end)

	return commonPrefix(leftPrefix, rightPrefix)
}

func commonPrefix(left, right string) string {
	minLength := len(left)
	if len(right) < minLength {
		minLength = len(right)
	}

	for i := 0; i < minLength; i++ {
		if left[i] != right[i] {
			return left[:i]
		}
	}
	return left[:minLength]
}

// horizontal scanning
func longestCommonPrefix3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

// vertical scanning
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		character := strs[0][i]
		for j := 1; j < len(strs); j++ {
			// reach end or character not equal
			if i == len(strs[j]) || strs[j][i] != character {
				return strs[0][0:i]
			}
		}
	}
	return strs[0]
}

// not efficiency
func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 1.find smallest length str
	smallLen := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < smallLen {
			smallLen = len(strs[i])
		}
	}

	// 2. check small str length whether zero
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

