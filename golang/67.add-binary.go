/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 *
 * https://leetcode.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (43.80%)
 * Likes:    1731
 * Dislikes: 270
 * Total Accepted:    444.5K
 * Total Submissions: 1M
 * Testcase Example:  '"11"\n"1"'
 *
 * Given two binary strings, return their sum (also a binary string).
 *
 * The input strings are both non-empty and contains only characters 1 or 0.
 *
 * Example 1:
 *
 *
 * Input: a = "11", b = "1"
 * Output: "100"
 *
 * Example 2:
 *
 *"1010"\n"1011"
 * Input: a = "1010", b = "1011"
 * Output: "10101"
 *
 *
 * Constraints:
 *
 *
 * Each string consists only of '0' or '1' characters.
 * 1 <= a.length, b.length <= 10^4
 * Each string is either "0" or doesn't contain any leading zero.
 *
 *
 */

// @lc code=start
// 彻底失败，浪费的时间比较长
func addBinary(a string, b string) string {
	return addBinary2(a, b)
}

// 自己写的一直有问题，暂时先放一放
func addBinary2(a string, b string) string {
	if a == "0" {
		return b
	} else if b == "0" {
		return a
	}

	maxLen := 0
	if len(a) > len(b) {
		maxLen = len(a)
	} else {
		maxLen = len(b)
	}

	var retByte = make([]byte, maxLen+1, maxLen+1)
	var curry byte
	// var one byte = '1'
	for i, j, index := len(a)-1, len(b)-1, maxLen-1; index >= 0; index-- {
		var aInt byte
		if i >= 0 {
			aInt = a[i] - '0'
		}
		var bInt byte
		if j >= 0 {
			bInt = a[j] - '0'
		}

		// 0: 0（0,0）0（1,1）1（0,1）
		// 1: 0（0,1）1（0,0）1（1,1）
		var sum byte
		sum = aInt ^ bInt ^ curry
		retByte[index+1] = sum + '0'

		// 0 (0 0), 0 (0 1), 1 (0 0),
		// 0 (1 1) 1 (1 1) 1 (0, 1)
		// curry = aInt&bInt | (aInt ^ bInt&curry)
		curry = aInt&bInt | curry&(aInt^bInt)
		retByte[i+1] = sum + '0'
		i--
		j--
	}
	if curry == 1 {
		retByte[0] = '1'
		return string(retByte)
	}
	return string(retByte[1:])
}

// 网上例子
func addBinary1(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	ans := make([]byte, len(a)+1)
	var carry byte
	for i, j := len(a), len(b); i >= 1 || j >= 1; {
		i, j = i-1, j-1
		var a2Digit, b2Digit byte
		if i >= 0 {
			a2Digit = a[i] - '0'
		}
		if j >= 0 {
			b2Digit = b[j] - '0'
		}
		// sum and carry of full adder
		sum := a2Digit ^ b2Digit ^ carry
		carry = a2Digit&b2Digit | carry&(a2Digit^b2Digit)
		ans[i+1] = sum + '0'
	}
	if carry == 1 {
		ans[0] = '1'
		return string(ans)
	}
	return string(ans[1:])
}

// @lc code=end

