/*
 * @lc app=leetcode id=263 lang=golang
 *
 * [263] Ugly Number
 *
 * https://leetcode.com/problems/ugly-number/description/
 *
 * algorithms
 * Easy (41.31%)
 * Likes:    527
 * Dislikes: 655
 * Total Accepted:    211.2K
 * Total Submissions: 508K
 * Testcase Example:  '6'
 *
 * Write a program to check whether a given number is an ugly number.
 *
 * Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.
 *
 * Example 1:
 *
 *
 * Input: 6
 * Output: true
 * Explanation: 6 = 2 × 3
 *
 * Example 2:
 *
 *
 * Input: 8
 * Output: true
 * Explanation: 8 = 2 × 2 × 2
 *
 *
 * Example 3:
 *
 *
 * Input: 14
 * Output: false
 * Explanation: 14 is not ugly since it includes another prime factor 7.
 *
 *
 * Note:
 *
 *
 * 1 is typically treated as an ugly number.
 * Input is within the 32-bit signed integer range: [−2^31,  2^31 − 1].
 *
 */

// @lc code=start
func isUgly(num int) bool {
	return isUgly3(num)
}

func isUgly3(num int) bool {
	if num < 1 {
		return false
	}
	factors := [3]int{2, 3, 5}
	for _, v := range factors {
		for num%v == 0 {
			num /= v
		}
	}
	return num == 1
}

func isUgly2(num int) bool {
	if num == 0 {
		return false
	}
	for i := 2; i < 6; i++ {
		for num%i == 0 {
			num /= i
		}
	}
	return num == 1
}

func isUgly1(num int) bool {
	if num == 0 {
		return false
	}
	if num == 1 || num == 2 || num == 3 || num == 5 {
		return true
	}
	if (num%2 == 0 && isUgly1(num/2)) || (num%3 == 0 && isUgly1(num/3)) || (num%5 == 0 && isUgly1(num/5)) {
		return true
	}
	return false
}

// @lc code=end