/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 *
 * https://leetcode.com/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (33.65%)
 * Likes:    2277
 * Dislikes: 561
 * Total Accepted:    363.3K
 * Total Submissions: 1.1M
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * Given a matrix of m x n elements (m rows, n columns), return all elements of
 * the matrix in spiral order.
 *
 * Example 1:
 *
 *
 * Input:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 4, 5, 6 ],
 * ⁠[ 7, 8, 9 ]
 * ]
 * Output: [1,2,3,6,9,8,7,4,5]
 *
 *
 * Example 2:
 *
 * Input:
 * [
 * ⁠ [1, 2, 3, 4],
 * ⁠ [5, 6, 7, 8],
 * ⁠ [9,10,11,12]
 * ]
 * Output: [1,2,3,4,8,12,11,10,9,5,6,7]
 *
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	return spiralOrder2(matrix)
}

func spiralOrder2(matrix [][]int) []int {
	if nil == matrix || len(matrix) == 0 {
		return []int{}
	}

	row, col := len(matrix), len(matrix[0])
	top, bottom, left, right := 0, row-1, 0, col-1
	result, index := make([]int, row*col), 0 // make enough array
	for {
		// from left to right
		for i := left; i <= right; i++ {
			result[index] = matrix[top][i]
			index++
		}
		top++
		if top > bottom {
			break
		}

		// from top to bottom
		for i := top; i <= bottom; i++ {
			result[index] = matrix[i][right]
			index++
		}
		right--
		if left > right {
			break
		}

		// from right to left
		for i := right; i >= left; i-- {
			result[index] = matrix[bottom][i]
			index++
		}
		bottom--
		if top > bottom {
			break
		}

		// from bottom to top
		for i := bottom; i >= top; i-- {
			result[index] = matrix[i][left]
			index++
		}
		left++
		if left > right {
			break
		}
	}

	return result
}

func spiralOrder1(matrix [][]int) []int {
	if nil == matrix || len(matrix) == 0 {
		return []int{}
	}

	result := make([]int, 0)
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for {
		// right
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++
		if top > bottom {
			break
		}

		// down
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--
		if left > right {
			break
		}

		// left
		for i := right; i >= left; i-- {
			result = append(result, matrix[bottom][i])
		}
		bottom--
		if top > bottom {
			break
		}

		// up
		for i := bottom; i >= top; i-- {
			result = append(result, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}

	return result
}

// @lc code=end