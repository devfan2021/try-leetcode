/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 *
 * https://leetcode.com/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (46.95%)
 * Likes:    4428
 * Dislikes: 143
 * Total Accepted:    696.9K
 * Total Submissions: 1.5M
 * Testcase Example:  '2'
 *
 * You are climbing a stair case. It takes n steps to reach to the top.
 *
 * Each time you can either climb 1 or 2 steps. In how many distinct ways can
 * you climb to the top?
 *
 * Example 1:
 *
 *
 * Input: 2
 * Output: 2
 * Explanation: There are two ways to climb to the top.
 * 1. 1 step + 1 step
 * 2. 2 steps
 *
 *
 * Example 2:
 *
 *
 * Input: 3
 * Output: 3
 * Explanation: There are three ways to climb to the top.
 * 1. 1 step + 1 step + 1 step
 * 2. 1 step + 2 steps
 * 3. 2 steps + 1 step
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 45
 *
 *
 */

// @lc code=start
func climbStairs(n int) int {
	return climbStairs3(n)
}

// dp, time complexity: O(n), space complexity: O(n)
func climbStairs3(n int) int {
	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// solution by recursive with memorization
func climbStairs2(n int) int {
	cache := map[int]int{}
	return helper(n, &cache)
}

func helper(n int, cache *map[int]int) int {
	if v, ok := (*cache)[n]; ok {
		return v
	}
	if n <= 2 {
		(*cache)[n] = n
		return n
	}

	retVal := helper(n-1, cache) + helper(n-2, cache)
	(*cache)[n] = retVal
	return retVal
}

// solution recursive, time limit exceeded
func climbStairs1(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

// @lc code=end