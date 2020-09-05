/*
 * @lc app=leetcode id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 *
 * https://leetcode.com/problems/set-matrix-zeroes/description/
 *
 * algorithms
 * Medium (43.21%)
 * Likes:    2436
 * Dislikes: 326
 * Total Accepted:    340.2K
 * Total Submissions: 785.4K
 * Testcase Example:  '[[1,1,1],[1,0,1],[1,1,1]]'
 *
 * Given an m x n matrix. If an element is 0, set its entire row and column to
 * 0. Do it in-place.
 *
 * Follow up:
 *
 *
 * A straight forward solution using O(mn) space is probably a bad idea.
 * A simple improvement uses O(m + n) space, but still not the best
 * solution.
 * Could you devise a constant space solution?
 *
 *
 *
 * Example 1:
 *
 *
 * Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
 * Output: [[1,0,1],[0,0,0],[1,0,1]]
 *
 *
 * Example 2:
 *
 *
 * Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
 * Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
 *
 *
 *
 * Constraints:
 *
 *
 * m == matrix.length
 * n == matrix[0].length
 * 1 <= m, n <= 200
 * -2^31 <= matrix[i][j] <= 2^31 - 1
 *
 *
 */

// [[0,0,0,5],[4,3,1,4],[0,1,1,4],[1,2,1,3],[0,0,1,1]]
// [[1,1,1],[1,0,1],[1,1,1]]
// [[1],[0]]

// @lc code=start
func setZeroes(matrix [][]int) {
	setZeroes2(matrix)
}

func setZeroes2(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	row, col := len(matrix), len(matrix[0])
	rowHash := map[int]bool{}
	colHash := map[int]bool{}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				rowHash[i] = true
				colHash[j] = true
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if rowHash[i] || colHash[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func setZeroes1(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	row, col := len(matrix), len(matrix[0])
	zeroColHash := map[int]bool{}

	for i := 0; i < row; i++ {

		zeroInCurRow := false
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				zeroInCurRow = true
				zeroColHash[j] = true
			}
		}

		if len(zeroColHash) > 0 {
			for j := 0; j < col; j++ {
				if zeroInCurRow {
					matrix[i][j] = 0
				}

				if zeroColHash[j] {
					for k := 0; k <= i; k++ {
						matrix[k][j] = 0
					}
				}
			}
		}
	}
}

// @lc code=end