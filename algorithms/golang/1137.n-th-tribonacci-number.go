/*
 * @lc app=leetcode id=1137 lang=golang
 *
 * [1137] N-th Tribonacci Number
 *
 * https://leetcode.com/problems/n-th-tribonacci-number/description/
 *
 * algorithms
 * Easy (55.96%)
 * Likes:    317
 * Dislikes: 41
 * Total Accepted:    44.6K
 * Total Submissions: 79.8K
 * Testcase Example:  '4'
 *
 * The Tribonacci sequence Tn is defined as follows:
 *
 * T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0.
 *
 * Given n, return the value of Tn.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 4
 * Output: 4
 * Explanation:
 * T_3 = 0 + 1 + 1 = 2
 * T_4 = 1 + 1 + 2 = 4
 *
 *
 * Example 2:
 *
 *
 * Input: n = 25
 * Output: 1389537
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= n <= 37
 * The answer is guaranteed to fit within a 32-bit integer, ie. answer <= 2^31
 * - 1.
 *
 */

// @lc code=start
func tribonacci(n int) int {
	return tribonacci3(n)
}

// cursive, time limit exceeded
func tribonacci1(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}

	return tribonacci1(n-1) + tribonacci1(n-2) + tribonacci1(n-3)
}

// recusive with memory
func tribonacci2(n int) int {
	hash := map[int]int{}
	return helper(n, hash)
}

func helper(n int, hash map[int]int) int {
	if v, ok := hash[n]; ok {
		return v
	}

	sum := 0
	if n == 0 {
		sum = 0
	} else if n == 1 || n == 2 {
		sum = 1
	} else {
		sum = helper(n-1, hash) + helper(n-2, hash) + helper(n-3, hash)
	}
	hash[n] = sum

	return sum
}

// dynamic program
func tribonacci3(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}

	dp := map[int]int{}
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-3] + dp[i-2] + dp[i-1]
	}
	return dp[n]
}

// @lc code=end