/*
 * @lc app=leetcode id=204 lang=golang
 *
 * [204] Count Primes
 *
 * https://leetcode.com/problems/count-primes/description/
 *
 * algorithms
 * Easy (31.23%)
 * Likes:    2064
 * Dislikes: 619
 * Total Accepted:    366.6K
 * Total Submissions: 1.2M
 * Testcase Example:  '10'
 *
 * Count the number of prime numbers less than a non-negative number, n.
 *
 * Example:
 *
 *
 * Input: 10
 * Output: 4
 * Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
 *
 *
 */

// @lc code=start
func countPrimes(n int) int {
	count := 0
	for i := 2; i < n; i++ {
		if (i != 2 && i%2 == 0) || (i != 3 && i%3 == 0) || (i != 5 && i%5 == 0) || (i != 7 && i%7 == 0) {
			continue
		}
		count++
	}
	return count
}

// @lc code=end