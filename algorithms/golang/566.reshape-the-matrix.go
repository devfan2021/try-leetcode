/*
 * @lc app=leetcode id=566 lang=golang
 *
 * [566] Reshape the Matrix
 *
 * https://leetcode.com/problems/reshape-the-matrix/description/
 *
 * algorithms
 * Easy (60.55%)
 * Likes:    813
 * Dislikes: 100
 * Total Accepted:    103.4K
 * Total Submissions: 170.7K
 * Testcase Example:  '[[1,2],[3,4]]\n1\n4'
 *
 * In MATLAB, there is a very useful function called 'reshape', which can
 * reshape a matrix into a new one with different size but keep its original
 * data.
 *
 *
 *
 * You're given a matrix represented by a two-dimensional array, and two
 * positive integers r and c representing the row number and column number of
 * the wanted reshaped matrix, respectively.
 *
 * ⁠The reshaped matrix need to be filled with all the elements of the original
 * matrix in the same row-traversing order as they were.
 *
 *
 *
 * If the 'reshape' operation with given parameters is possible and legal,
 * output the new reshaped matrix; Otherwise, output the original matrix.
 *
 *
 * Example 1:
 *
 * Input:
 * nums =
 * [[1,2],
 * ⁠[3,4]]
 * r = 1, c = 4
 * Output:
 * [[1,2,3,4]]
 * Explanation:The row-traversing of nums is [1,2,3,4]. The new reshaped matrix
 * is a 1 * 4 matrix, fill it row by row by using the previous list.
 *
 *
 *
 * Example 2:
 *
 * Input:
 * nums =
 * [[1,2],
 * ⁠[3,4]]
 * r = 2, c = 4
 * Output:
 * [[1,2],
 * ⁠[3,4]]
 * Explanation:There is no way to reshape a 2 * 2 matrix to a 2 * 4 matrix. So
 * output the original matrix.
 *
 *
 *
 * Note:
 *
 * The height and width of the given matrix is in range [1, 100].
 * The given r and c are all positive.
 *
 *
 */

// @lc code=start
// [[1,2],[3,4]]\n2\n4
func matrixReshape(nums [][]int, r int, c int) [][]int {
	return matrixReshape2(nums, r, c)
}

func matrixReshape2(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 || len(nums)*len(nums[0]) != r*c {
		return nums
	}

	retVals := [][]int{}
	totalCol, curRow, curCol := len(nums[0]), 0, 0
	for i := 0; i < r; i++ {
		// make enough space to store data
		retVals = append(retVals, make([]int, c, c))
		for j := 0; j < c; j++ {
			retVals[i][j] = nums[curRow][curCol]
			curCol++
			if curCol == totalCol {
				curRow++
				curCol = 0
			}
		}
	}
	return retVals
}

func matrixReshape1(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 || len(nums)*len(nums[0]) != r*c {
		return nums
	}

	retVals := [][]int{}
	totalCol, curRow, curCol := len(nums[0]), 0, 0
	for i := 0; i < r; i++ {
		retVals = append(retVals, []int{})
		for j := 0; j < c; j++ {
			retVals[i] = append(retVals[i], nums[curRow][curCol])
			curCol++
			if curCol == totalCol {
				curRow++
				curCol = 0
			}
		}
	}
	return retVals
}

// @lc code=end