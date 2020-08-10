/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (38.68%)
 * Likes:    5130
 * Dislikes: 224
 * Total Accepted:    1M
 * Total Submissions: 2.7M
 * Testcase Example:  '"()"'
 *
 * Given a string containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 *
 * An input string is valid if:
 *
 *
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 *
 *
 * Note that an empty string is also considered valid.
 *
 * Example 1:
 *
 *
 * Input: "()"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: "()[]{}"
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: "(]"
 * Output: false
 *
 *
 * Example 4:
 *
 *
 * Input: "([)]"
 * Output: false
 *
 *
 * Example 5:
 *
 *
 * Input: "{[]}"
 * Output: true
 *
 *
 */

// @lc code=start
// 40 = ‘(’, 41 = ‘)’ // 123 = ‘{’, 125 = ‘}’ // 91 = ‘[’, 93 = ‘]’
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	hash := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	stack := make([]rune, 0, 1)
	for _, ch := range s {
		switch ch {
		case '(', '{', '[':
			// case 40, 123, 91:
			stack = append(stack, ch)
		case ')', '}', ']':
			// case 41, 125, 93:
			if len(stack) == 0 || stack[len(stack)-1] != hash[ch] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}

// @lc code=end