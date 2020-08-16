/*
 * @lc app=leetcode id=172 lang=golang
 *
 * [172] Factorial Trailing Zeroes
 *
 * https://leetcode.com/problems/factorial-trailing-zeroes/description/
 *
 * algorithms
 * Easy (37.73%)
 * Likes:    918
 * Dislikes: 1151
 * Total Accepted:    215.5K
 * Total Submissions: 570.4K
 * Testcase Example:  '3'
 *
 * Given an integer n, return the number of trailing zeroes in n!.
 *
 * Example 1:
 *
 *
 * Input: 3
 * Output: 0
 * Explanation: 3! = 6, no trailing zero.
 *
 * Example 2:
 *
 *
 * Input: 5
 * Output: 1
 * Explanation: 5! = 120, one trailing zero.
 *
 * Note: Your solution should be in logarithmic time complexity.
 *
 */

// https://leetcode.com/problems/factorial-trailing-zeroes/discuss/52367/my-explanation-of-the-logn-solution
// 10 is the product of 2 and 5. In n!, we need to know how many 2 and 5, and the number of zeros is the minimum of the number of 2 and the number of 5.
//Since multiple of 2 is more than multiple of 5, the number of zeros is dominant by the number of 5.

// @lc code=start
func trailingZeroes(n int) int {
	return trailingZeroes2(n)
}

// The ZERO comes from 10.
// The 10 comes from 2 x 5
// And we need to account for all the products of 5 and 2. likes 4×5 = 20 ...
// So, if we take all the numbers with 5 as a factor, we'll have way more than enough even numbers to pair with them to get factors of 10
func trailingZeroes3(n int) int {
	retVal := 0
	for n > 0 {
		retVal += n / 5
		n /= 5
	}
	return retVal
}

func trailingZeroes2(n int) int {
	if n == 0 {
		return 0
	}
	return n/5 + trailingZeroes2(n/5)
}

// Time Limit Exceeded
func trailingZeroes1(n int) int {
	retVal := 0
	for i := 1; i <= n; i++ {
		for j := i; j%5 == 0; {
			retVal++
			j /= 5
		}
	}
	return retVal
}

// @lc code=end