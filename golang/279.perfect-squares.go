/*
 * @lc app=leetcode id=279 lang=golang
 *
 * [279] Perfect Squares
 *
 * https://leetcode.com/problems/perfect-squares/description/
 *
 * algorithms
 * Medium (45.84%)
 * Likes:    2978
 * Dislikes: 184
 * Total Accepted:    311.2K
 * Total Submissions: 660.7K
 * Testcase Example:  '12'
 *
 * Given a positive integer n, find the least number of perfect square numbers
 * (for example, 1, 4, 9, 16, ...) which sum to n.
 *
 * Example 1:
 *
 *
 * Input: n = 12
 * Output: 3
 * Explanation: 12 = 4 + 4 + 4.
 *
 * Example 2:
 *
 *
 * Input: n = 13
 * Output: 2
 * Explanation: 13 = 4 + 9.
 */

// @lc code=start
func numSquares(n int) int {
	return numSquares2(n)
}

func numSquares2(n int) int {
	// to make n length visited
	remainders, visited, result := make([]int, 0), make([]bool, n), 1

	remainders = append(remainders, n)
	for len(remainders) > 0 {
		// add length to avoid each make remainders
		length := len(remainders)
		for _, curr := range remainders {
			for i := 1; i*i <= curr; i++ {
				if curr == i*i {
					return result
				}
				if !visited[curr-i*i] {
					remainders = append(remainders, curr-i*i)
					visited[curr-i*i] = true
				}
			}
		}
		remainders = remainders[length:]
		result++
	}
	return result
}

// bfs, 1. find all child squares of n, 2. do bfs
func numSquares1(n int) int {
	squares := make([]int, 0)
	// careful boundary, i <= n/2 and tmp <=n
	for i := 1; i <= n/2; i++ {
		tmp := i * i
		if tmp <= n {
			squares = append(squares, tmp)
		}
	}

	// use set to avoid duplicate check
	remainders := map[int]bool{}
	remainders[n] = true
	step := 0
	for len(remainders) > 0 {
		step++
		nextRemainders := map[int]bool{}
		for k, _ := range remainders {
			for _, v := range squares {
				if k == v {
					return step // find the node!
				} else if k < v { // remainder less square
					break
				} else {
					if _, ok := nextRemainders[k-v]; !ok {
						nextRemainders[k-v] = true
					}
				}
			}
		}
		remainders = nextRemainders
	}
	return step
}

// @lc code=end