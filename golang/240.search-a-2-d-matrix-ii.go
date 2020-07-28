/*
 * @lc app=leetcode id=240 lang=golang
 *
 * [240] Search a 2D Matrix II
 *
 * https://leetcode.com/problems/search-a-2d-matrix-ii/description/
 *
 * algorithms
 * Medium (42.90%)
 * Likes:    3202
 * Dislikes: 75
 * Total Accepted:    324.4K
 * Total Submissions: 752.6K
 * Testcase Example:  '[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]\n' +
  '5'
 *
 * Write an efficient algorithm that searches for a value in an m x n matrix.
 * This matrix has the following properties:
 *
 *
 * Integers in each row are sorted in ascending from left to right.
 * Integers in each column are sorted in ascending from top to bottom.
 *
 *
 * Example:
 *
 * Consider the following matrix:
 *
 *
 * [
 * ⁠ [1,   4,  7, 11, 15],
 * ⁠ [2,   5,  8, 12, 19],
 * ⁠ [3,   6,  9, 16, 22],
 * ⁠ [10, 13, 14, 17, 24],
 * ⁠ [18, 21, 23, 26, 30]
 * ]
 *
 *
 * Given target = 5, return true.
 *
 * Given target = 20, return false.
 *
*/
// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	return searchMatrix2(matrix, target)
}

// use each row and each col order character, do iterate from top-right
func searchMatrix2(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) < 1 || len(matrix[0]) < 1 {
		return false
	}
	row, col := 0, len(matrix[0])-1
	for row <= len(matrix)-1 && col >= 0 {
		if target == matrix[row][col] {
			return true
		} else if target > matrix[row][col] {
			row++
		} else {
			col--
		}
	}
	return false
}

// do binary search for each matrix row
func searchMatrix1(matrix [][]int, target int) bool {
	for _, v := range matrix {
		if binarySearch(v, target) {
			return true
		}
	}
	return false
}

func binarySearch(matrix []int, target int) bool {
	left, right := 0, len(matrix)-1
	for left < right {
		mid := left + (right-left)/2
		if matrix[mid] == target {
			return true
		} else if matrix[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left < len(matrix) && matrix[left] == target {
		return true
	}
	return false
}

// @lc code=end