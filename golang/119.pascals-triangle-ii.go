/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 *
 * https://leetcode.com/problems/pascals-triangle-ii/description/
 *
 * algorithms
 * Easy (48.37%)
 * Likes:    793
 * Dislikes: 196
 * Total Accepted:    280.4K
 * Total Submissions: 578.1K
 * Testcase Example:  '3'
 *
 * Given a non-negative index k where k ≤ 33, return the k^th index row of the
 * Pascal's triangle.
 *
 * Note that the row index starts from 0.
 *
 *
 * In Pascal's triangle, each number is the sum of the two numbers directly
 * above it.
 *
 * Example:
 *
 *
 * Input: 3
 * Output: [1,3,3,1]
 *
 *
 * Follow up:
 *
 * Could you optimize your algorithm to use only O(k) extra space?
 *
 */

// @lc code=start
func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	var retVal [][]int
	oneVal := 1
	retVal = append(retVal, []int{oneVal})
	for i := 1; i <= rowIndex; i++ {
		var rowVal []int
		rowVal = append(rowVal, oneVal) // current row: add one in first
		lastRowVal := retVal[i-1]       // last row element
		for j := 0; j+1 < i; j++ {
			rowVal = append(rowVal, lastRowVal[j]+lastRowVal[j+1])
		}
		rowVal = append(rowVal, oneVal) // current row: add one in last

		retVal = append(retVal, rowVal) // add current row to return slice
	}
	return retVal[rowIndex]
}

// @lc code=end

