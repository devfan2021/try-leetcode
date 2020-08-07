import (
	"strings"
)

/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 *
 * https://leetcode.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (32.51%)
 * Likes:    683
 * Dislikes: 2495
 * Total Accepted:    385.3K
 * Total Submissions: 1.2M
 * Testcase Example:  '"Hello World"'
 *
 * Given a string s consists of upper/lower-case alphabets and empty space
 * characters ' ', return the length of last word (last word means the last
 * appearing word if we loop from left to right) in the string.
 *
 * If the last word does not exist, return 0.
 *
 * Note: A word is defined as a maximal substring consisting of non-space
 * characters only.
 *
 * Example:
 *
 *
 * Input: "Hello World"
 * Output: 5
 *
 *
 *
 *
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	return lengthOfLastWord2(s)
}

func lengthOfLastWord2(s string) int {
	length, tail := 0, len(s)-1
	for tail >= 0 && s[tail] == ' ' {
		tail--
	}

	for tail >= 0 && s[tail] != ' ' {
		length++
		tail--
	}
	return length
}

func lengthOfLastWord1(s string) int {
	if len(s) == 0 {
		return 0
	}
	childs := strings.Split(s, " ")

	// handle "a "
	for i := len(childs) - 1; i >= 0; i-- {
		item := childs[i]
		if item != "" {
			return len(item)
		}
	}
	return 0
}

// @lc code=end