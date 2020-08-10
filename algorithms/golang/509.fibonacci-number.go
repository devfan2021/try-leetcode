/*
 * @lc app=leetcode id=509 lang=golang
 *
 * [509] Fibonacci Number
 *
 * https://leetcode.com/problems/fibonacci-number/description/
 *
 * algorithms
 * Easy (67.07%)
 * Likes:    656
 * Dislikes: 193
 * Total Accepted:    213.6K
 * Total Submissions: 317.9K
 * Testcase Example:  '2'
 *
 * The Fibonacci numbers, commonly denoted F(n) form a sequence, called the
 * Fibonacci sequence, such that each number is the sum of the two preceding
 * ones, starting from 0 and 1. That is,
 *
 *
 * F(0) = 0,   F(1) = 1
 * F(N) = F(N - 1) + F(N - 2), for N > 1.
 *
 *
 * Given N, calculate F(N).
 *
 *
 *
 * Example 1:
 *
 *
 * Input: 2
 * Output: 1
 * Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.
 *
 *
 * Example 2:
 *
 *
 * Input: 3
 * Output: 2
 * Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.
 *
 *
 * Example 3:
 *
 *
 * Input: 4
 * Output: 3
 * Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.
 *
 *
 *
 *
 * Note:
 *
 * 0 ≤ N ≤ 30.
 *
 */

// @lc code=start
func fib(N int) int {
	return fib2(N)
}

// solution by recursive with memorization
func fib2(N int) int {
	cache := map[int]int{}
	return helper(N, &cache)
}

func helper(N int, cache *map[int]int) int {
	if v, ok := (*cache)[N]; ok {
		return v
	}

	if N == 0 || N == 1 {
		return N
	}
	retVal := helper(N-1, cache) + helper(N-2, cache)
	(*cache)[N] = retVal
	return retVal
}

// solution by recursive
func fib1(N int) int {
	if N == 0 || N == 1 {
		return N
	}
	return fib(N-1) + fib(N-2)
}

// @lc code=end