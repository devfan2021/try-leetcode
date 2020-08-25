/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 *
 * https://leetcode.com/problems/string-to-integer-atoi/description/
 *
 * algorithms
 * Medium (15.27%)
 * Likes:    1648
 * Dislikes: 9562
 * Total Accepted:    585.6K
 * Total Submissions: 3.8M
 * Testcase Example:  '"42"'
 *
 * Implement atoi which converts a string to an integer.
 *
 * The function first discards as many whitespace characters as necessary until
 * the first non-whitespace character is found. Then, starting from this
 * character, takes an optional initial plus or minus sign followed by as many
 * numerical digits as possible, and interprets them as a numerical value.
 *
 * The string can contain additional characters after those that form the
 * integral number, which are ignored and have no effect on the behavior of
 * this function.
 *
 * If the first sequence of non-whitespace characters in str is not a valid
 * integral number, or if no such sequence exists because either str is empty
 * or it contains only whitespace characters, no conversion is performed.
 *
 * If no valid conversion could be performed, a zero value is returned.
 *
 * Note:
 *
 *
 * Only the space character ' ' is considered as whitespace character.
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−2^31,  2^31 − 1]. If the numerical
 * value is out of the range of representable values, INT_MAX (2^31 − 1) or
 * INT_MIN (−2^31) is returned.
 *
 *
 * Example 1:
 *
 *
 * Input: "42"
 * Output: 42
 *
 *
 * Example 2:
 *
 *
 * Input: "   -42"
 * Output: -42
 * Explanation: The first non-whitespace character is '-', which is the minus
 * sign.
 * Then take as many numerical digits as possible, which gets 42.
 *
 *
 * Example 3:
 *
 *
 * Input: "4193 with words"
 * Output: 4193
 * Explanation: Conversion stops at digit '3' as the next character is not a
 * numerical digit.
 *
 *
 * Example 4:
 *
 *
 * Input: "words and 987"
 * Output: 0
 * Explanation: The first non-whitespace character is 'w', which is not a
 * numerical
 * digit or a +/- sign. Therefore no valid conversion could be performed.
 *
 * Example 5:
 *
 *
 * Input: "-91283472332"
 * Output: -2147483648
 * Explanation: The number "-91283472332" is out of the range of a 32-bit
 * signed integer.
 * Thefore INT_MIN (−2^31) is returned.
 *
 */

// @lc code=start
func myAtoi(str string) int {
	return myAtoi2(str)
}

func myAtoi2(str string) int {
	if len(str) == 0 {
		return 0
	}

	idx, sign := 0, 1
	for idx < len(str) && str[idx] == ' ' {
		idx++
	}

	if idx < len(str) && (str[idx] == '-' || str[idx] == '+') {
		if str[idx] == '-' {
			sign = -1
		}
		idx++
	}

	result := 0
	for idx < len(str) && isNumerical(str[idx]) {
		n := int(str[idx] - '0')

		if sign == 1 && result > (2147483647-n)/10 {
			return 2147483647
		}

		if sign == -1 && -result < (-2147483648+n)/10 {
			return -2147483648
		}

		result = result*10 + n
		idx++
	}
	return result * sign
}

func isNumerical(b byte) bool {
	return b >= '0' && b <= '9'
}

func myAtoi1(str string) int {
	if len(str) == 0 {
		return 0
	}

	index := 0
	for index < len(str) {
		if str[index] == ' ' {
			index++
		} else {
			break
		}
	}

	sign := 1
	if index < len(str) && (str[index] == '-' || str[index] == '+') {
		if str[index] == '-' {
			sign = -1
		}
		index++
	}

	// Build the result and check for overflow/underflow condition
	minVal, maxVal, result := -1<<31, 1<<31-1, 0
	for index < len(str) && str[index] >= '0' && str[index] <= '9' {
		if result > maxVal/10 || result == maxVal/10 && int(str[index]-'0') > maxVal%10 {
			if sign == 1 {
				return maxVal
			} else {
				return minVal
			}
		}
		result = result*10 + int(str[index]-'0')
		index++
	}
	return result * sign
}

// @lc code=end