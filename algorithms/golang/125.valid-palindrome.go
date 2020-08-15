import "strings"

/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 *
 * https://leetcode.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (35.55%)
 * Likes:    1337
 * Dislikes: 3067
 * Total Accepted:    653.9K
 * Total Submissions: 1.8M
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * Given a string, determine if it is a palindrome, considering only
 * alphanumeric characters and ignoring cases.
 *
 * Note: For the purpose of this problem, we define empty string as valid
 * palindrome.
 *
 * Example 1:
 *
 *
 * Input: "A man, a plan, a canal: Panama"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: "race a car"
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * s consists only of printable ASCII characters.
 *
 *
 */

// @lc code=start
func isPalindrome(s string) bool {
	return isPalindrome1(s)
}

func isPalindrome2(s string) bool {
	if len(s) <= 1 {
		return true
	}

	left, right := 0, len(s)-1
	for left < right {
		if !isChar(s[left]) {
			left++
		} else if !isChar(s[right]) {
			right--
		} else if isEq(s[left], s[right]) {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

func isChar(b byte) bool {
	if b >= 'a' && b <= 'z' {
		return true
	} else if b >= 'A' && b <= 'Z' {
		return true
	} else if b >= '0' && b <= '9' {
		return true
	}
	return false
}

func isEq(a, b byte) bool {
	i, j := string(a), string(b)
	return strings.EqualFold(i, j)
}

func isPalindrome1(s string) bool {
	l := 0
	r := len(s) - 1
	for l < r {
		if !isAlphanumeric(s[l]) {
			l++
			continue
		}
		if !isAlphanumeric(s[r]) {
			r--
			continue
		}
		if upper(s[l]) == upper(s[r]) {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}

func isAlphanumeric(char byte) bool {
	return (char >= '0' && char <= '9') || (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func upper(char byte) byte {
	if char >= 'a' && char <= 'z' {
		return char - 32
	}
	return char
}

// @lc code=end