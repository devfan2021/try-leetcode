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
// 10000
// @lc code=start
func countPrimes(n int) int {
	return countPrimes2(n)
}

func countPrimes2(n int) int {
	if n < 2 {
		return 0
	}

	// line array overmatch map
	// hash := make(map[int]bool, n)
	hash := make([]bool, n)
	for i := 2; i < n; i++ {
		hash[i] = true
	}

	for i := 2; i*i < n; i++ {
		if hash[i] {
			for j := 2 * i; j < n; j += i {
				hash[j] = false
			}
		}
	}

	count := 0
	for i := 2; i < n; i++ {
		if hash[i] {
			count++
		}
	}

	return count
}

// time limit exceeded: 499979
func countPrimes1(n int) int {
	hash := map[int]bool{}
	for i := 2; i < n; i++ {
		hash[i] = true
	}

	count := 0
	for i := 2; i < n; i++ {
		if isPrime1(i) {
			count++
		}
	}
	return count
}

func isPrime1(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// @lc code=end