/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 *
 * https://leetcode.com/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (51.68%)
 * Likes:    1401
 * Dislikes: 107
 * Total Accepted:    374.1K
 * Total Submissions: 722.4K
 * Testcase Example:  '5'
 *
 * Given a non-negative integer numRows, generate the first numRows of Pascal's
 * triangle.
 *
 *
 * In Pascal's triangle, each number is the sum of the two numbers directly
 * above it.
 *
 * Example:
 *
 *
 * Input: 5
 * Output:
 * [
 * ⁠    [1],
 * ⁠   [1,1],
 * ⁠  [1,2,1],
 * ⁠ [1,3,3,1],
 * ⁠[1,4,6,4,1]
 * ]
 *
 *
 */

// @lc code=start
func generate(numRows int) [][]int {
	if numRows == 0 {
		return make([][]int, 0)
	}

	var retVal [][]int
	oneVal := 1
	retVal = append(retVal, []int{oneVal})
	for i := 1; i < numRows; i++ {
		var rowVal []int
		rowVal = append(rowVal, oneVal) // current row: add one in first
		lastRowVal := retVal[i-1]       // last row element
		for j := 0; j+1 < i; j++ {
			rowVal = append(rowVal, lastRowVal[j]+lastRowVal[j+1])
		}
		rowVal = append(rowVal, oneVal) // current row: add one in last

		retVal = append(retVal, rowVal) // add current row to return slice
	}
	return retVal
}

// @lc code=end

