/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 *
 * https://leetcode.com/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (47.87%)
 * Likes:    2289
 * Dislikes: 1544
 * Total Accepted:    927.8K
 * Total Submissions: 1.9M
 * Testcase Example:  '121'
 *
 * Determine whether an integer is a palindrome. An integer is a palindrome
 * when it reads the same backward as forward.
 *
 * Example 1:
 *
 *
 * Input: 121
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: -121
 * Output: false
 * Explanation: From left to right, it reads -121. From right to left, it
 * becomes 121-. Therefore it is not a palindrome.
 *
 *
 * Example 3:
 *
 *
 * Input: 10
 * Output: false
 * Explanation: Reads 01 from right to left. Therefore it is not a
 * palindrome.
 *
 *
 * Follow up:
 *
 * Coud you solve it without converting the integer to a string?
 *
 */

// @lc code=start
func isPalindrome(x int) bool {
	return isPalindrome1(x)
}

// topest version
func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	res, xo := 0, x

	// 反向相除，再累加，看是否等于开始值
	for xo != 0 {
		pop := xo % 10
		xo /= 10
		res = res*10 + pop
	}

	return res == x
}

// myself write version
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}

	val := x
	factor := 1

	for x >= 10 {
		x = x / 10
		factor *= 10
	}

	for val > 0 {
		if val < factor {
			for val > factor {
				last := val % 10
				if last != 0 {
					return false
				}
				factor /= 10
			}
		}

		last := val % 10
		first := val / factor
		if last != first {
			return false
		}
		val = (val - first*factor) / 10
		factor /= 10 * 10
		if val < 0 {
			return false
		}
	}
	return true
}

// fmt.Println(isPalindrome(9999))
// fmt.Println(isPalindrome(99))
// fmt.Println(isPalindrome(1))
// fmt.Println(isPalindrome(10))
// fmt.Println(isPalindrome(1000021))
// fmt.Println(isPalindrome(1001))
// fmt.Println(isPalindrome(1000030001)) // 异常通不过
// fmt.Println(isPalindrome(1000031111)) // 异常通不过

// @lc code=end