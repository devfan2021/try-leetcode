/*
 * @lc app=leetcode id=59 lang=golang
 *
 * [59] Spiral Matrix II
 *
 * https://leetcode.com/problems/spiral-matrix-ii/description/
 *
 * algorithms
 * Medium (54.07%)
 * Likes:    1106
 * Dislikes: 119
 * Total Accepted:    202K
 * Total Submissions: 372K
 * Testcase Example:  '3'
 *
 * Given a positive integer n, generate a square matrix filled with elements
 * from 1 to n^2 in spiral order.
 *
 * Example:
 *
 *
 * Input: 3
 * Output:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 8, 9, 4 ],
 * ⁠[ 7, 6, 5 ]
 * ]
 *
 *
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}

	retVal := make([][]int, n)
	for i := 0; i < n; i++ {
		retVal[i] = make([]int, n)
	}

	total, val := n*n, 1
	left, right, top, bottom := 0, n-1, 0, n-1

	for val <= total {
		// from left to right
		for i := left; i <= right; i++ {
			retVal[top][i] = val
			val++
		}
		top++
		if top > bottom {
			break
		}

		// from top to bottom
		for i := top; i <= bottom; i++ {
			retVal[i][right] = val
			val++
		}
		right--
		if left > right {
			break
		}

		// from right to left
		for i := right; i >= left; i-- {
			retVal[bottom][i] = val
			val++
		}
		bottom--
		if top > bottom {
			break
		}

		// from bottom to top
		for i := bottom; i >= top; i-- {
			retVal[i][left] = val
			val++
		}
		left++
		if left > right {
			break
		}
	}

	return retVal
}

// @lc code=end