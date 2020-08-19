/*
 * @lc app=leetcode id=258 lang=golang
 *
 * [258] Add Digits
 *
 * https://leetcode.com/problems/add-digits/description/
 *
 * algorithms
 * Easy (56.26%)
 * Likes:    878
 * Dislikes: 1146
 * Total Accepted:    318.3K
 * Total Submissions: 551.8K
 * Testcase Example:  '38'
 *
 * Given a non-negative integer num, repeatedly add all its digits until the
 * result has only one digit.
 *
 * Example:
 *
 *
 * Input: 38
 * Output: 2
 * Explanation: The process is like: 3 + 8 = 11, 1 + 1 = 2.
 * Since 2 has only one digit, return it.
 *
 *
 * Follow up:
 * Could you do it without any loop/recursion in O(1) runtime?
 */

// @lc code=start
func addDigits(num int) int {
	return addDigits3(num)
}

func addDigits1(num int) int {
	digitalRoot := 0
	for num > 0 {
		digitalRoot += num % 10
		num = num / 10
		if num == 0 && digitalRoot > 9 {
			num = digitalRoot
			digitalRoot = 0
		}
	}
	return digitalRoot
}

// time complexity:O(n), space complexity:O(1)
func addDigits2(num int) int {
	for num >= 10 {
		num = num/10 + num%10
	}
	return num
}

// 	num= a1 + a2*10    + a3*100   + ... + an*(10*n)
//	num= a1 + a2*(9+1) + a3(99+1) + ... + an*(10*n-1+1)
//   	= a1 + a2 + a3 + ... + an（正符合该题，如果 num > 10，重复上面操作）
// time complexity:O(1), space complexity:O(1)
func addDigits3(num int) int {
	return (num-1)%9 + 1
}

// @lc code=end