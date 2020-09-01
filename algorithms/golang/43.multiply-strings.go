/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 *
 * https://leetcode.com/problems/multiply-strings/description/
 *
 * algorithms
 * Medium (33.95%)
 * Likes:    1857
 * Dislikes: 833
 * Total Accepted:    308.4K
 * Total Submissions: 906.2K
 * Testcase Example:  '"2"\n"3"'
 *
 * Given two non-negative integers num1 and num2 represented as strings, return
 * the product of num1 and num2, also represented as a string.
 *
 * Example 1:
 *
 *
 * Input: num1 = "2", num2 = "3"
 * Output: "6"
 *
 * Example 2:
 *
 *
 * Input: num1 = "123", num2 = "456"
 * Output: "56088"
 *
 *
 * Note:
 *
 *
 * The length of both num1 and num2 is < 110.
 * Both num1 and num2 contain only digits 0-9.
 * Both num1 and num2 do not contain any leading zero, except the number 0
 * itself.
 * You must not use any built-in BigInteger library or convert the inputs to
 * integer directly.
 *
 *
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	return multiply1(num1, num2)
}

// "123"\n"456"
func multiply1(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	l1, l2 := len(num1), len(num2)
	res := make([]byte, l1+l2)

	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			val := (num1[i] - '0') * (num2[j] - '0')
			sum := val + res[i+j+1] // add old value

			res[i+j+1] = sum % 10
			res[i+j] += sum / 10 // add curry
		}
	}

	if res[0] == 0 {
		res = res[1:]
	}
	for i := range res {
		res[i] += '0'
	}
	return string(res)
}

// @lc code=end