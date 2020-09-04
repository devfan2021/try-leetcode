/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 *
 * https://leetcode.com/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (54.60%)
 * Likes:    3419
 * Dislikes: 64
 * Total Accepted:    446.9K
 * Total Submissions: 816.3K
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * Given a m x n grid filled with non-negative numbers, find a path from top
 * left to bottom right which minimizes the sum of all numbers along its path.
 *
 * Note: You can only move either down or right at any point in time.
 *
 * Example:
 *
 *
 * Input:
 * [
 * [1,3,1],
 * ⁠ [1,5,1],
 * ⁠ [4,2,1]
 * ]
 * Output: 7
 * Explanation: Because the path 1→3→1→1→1 minimizes the sum.
 *
 *
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	return minPathSum2(grid)
}

func minPathSum2(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	row, col := len(grid), len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 && j > 0 {
				grid[i][j] = grid[i][j] + grid[i][j-1]
			} else if i > 0 && j == 0 {
				grid[i][j] = grid[i][j] + grid[i-1][j]
			} else if i > 0 && j > 0 {
				if grid[i-1][j] < grid[i][j-1] {
					grid[i][j] = grid[i-1][j] + grid[i][j]
				} else {
					grid[i][j] = grid[i][j-1] + grid[i][j]
				}
			}
		}
	}

	return grid[row-1][col-1]
}

// dp[i][j] = min(dp[i-1][j], dp[i][j]) + grid[i][j]
func minPathSum1(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	row, col := len(grid), len(grid[0])
	sum := make([][]int, row)
	for i := 0; i < row; i++ {
		sum[i] = make([]int, col)
	}

	sum[0][0] = grid[0][0]
	for i := 1; i < row; i++ {
		sum[i][0] = sum[i-1][0] + grid[i][0]
	}

	for i := 1; i < col; i++ {
		sum[0][i] = sum[0][i-1] + grid[0][i]
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			sum[i][j] = min(sum[i-1][j], sum[i][j-1]) + grid[i][j]
		}
	}

	return sum[row-1][col-1]
}

func min(val1, val2 int) int {
	if val1 < val2 {
		return val1
	}
	return val2
}

// @lc code=end