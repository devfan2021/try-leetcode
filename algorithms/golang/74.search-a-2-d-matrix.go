/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 *
 * https://leetcode.com/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (36.54%)
 * Likes:    2049
 * Dislikes: 165
 * Total Accepted:    344.1K
 * Total Submissions: 940.4K
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,50]]\n3'
 *
 * Write an efficient algorithm that searches for a value in an m x n matrix.
 * This matrix has the following properties:
 *
 *
 * Integers in each row are sorted from left to right.
 * The first integer of each row is greater than the last integer of the
 * previous row.
 *
 *
 * Example 1:
 *
 *
 * Input:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 3
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 13
 * Output: false
 *
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	return searchMatrix2(matrix, target)
}

func searchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	cols := len(matrix[0])
	left, right := 0, len(matrix)*cols-1
	for left <= right {
		mid := (left + right) / 2
		curVal := matrix[mid/cols][mid%cols]

		if target == curVal {
			return true
		}

		if target > curVal {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

func searchMatrix1(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rowIndex, colIndex := 0, len(matrix[0])-1
	for rowIndex < len(matrix) && colIndex >= 0 {
		curVal := matrix[rowIndex][colIndex]
		if target > curVal {
			rowIndex++
		} else if target < curVal {
			colIndex--
		} else {
			return true
		}
	}
	return false
}

// @lc code=end