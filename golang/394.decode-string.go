import "strconv"

/*
 * @lc app=leetcode id=394 lang=golang
 *
 * [394] Decode String
 *
 * https://leetcode.com/problems/decode-string/description/
 *
 * algorithms
 * Medium (49.26%)
 * Likes:    3180
 * Dislikes: 163
 * Total Accepted:    210.4K
 * Total Submissions: 423K
 * Testcase Example:  '"3[a]2[bc]"'
 *
 * Given an encoded string, return its decoded string.
 *
 * The encoding rule is: k[encoded_string], where the encoded_string inside the
 * square brackets is being repeated exactly k times. Note that k is guaranteed
 * to be a positive integer.
 *
 * You may assume that the input string is always valid; No extra white spaces,
 * square brackets are well-formed, etc.
 *
 * Furthermore, you may assume that the original data does not contain any
 * digits and that digits are only for those repeat numbers, k. For example,
 * there won't be input like 3a or 2[4].
 *
 *
 * Example 1:
 * Input: s = "3[a]2[bc]"
 * Output: "aaabcbc"
 * Example 2:
 * Input: s = "3[a2[c]]"
 * Output: "accaccacc"
 * Example 3:
 * Input: s = "2[abc]3[cd]ef"
 * Output: "abcabccdcdcdef"
 * Example 4:
 * Input: s = "abc3[cd]xyz"
 * Output: "abccdcdcdxyz"
 *
 */

// @lc code=start
func decodeString(s string) string {
	return decodeString1(s)
}

// solution, help with stack
func decodeString1(s string) string {
	if len(s) == 0 {
		return s
	}

	stack := make([]byte, 0)
	n := len(s)
	for i := 0; i < n; i++ {
		switch s[i] {
		case ']':
			temp := make([]byte, 0)
			// find from inverse direction
			for stack[len(stack)-1] != '[' && len(stack) > 0 { // save charater
				temp = append(temp, stack[len(stack)-1]) // inverse direction
				stack = stack[:len(stack)-1]
			}
			// pop '['
			stack = stack[:len(stack)-1]

			idx := 1
			for len(stack) >= idx && isDigit(stack[len(stack)-idx]) {
				idx++
			}

			nIndex := len(stack) - idx + 1 // number index
			num := stack[nIndex:]          // number
			stack = stack[:nIndex]         //pop number
			count, _ := strconv.Atoi(string(num))
			for j := 0; j < count; j++ {
				for k := len(temp) - 1; k >= 0; k-- { // inverse direction
					stack = append(stack, temp[k])
				}
			}
		default:
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// @lc code=end