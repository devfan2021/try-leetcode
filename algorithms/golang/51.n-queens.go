/*
 * @lc app=leetcode id=51 lang=golang
 *
 * [51] N-Queens
 *
 * https://leetcode.com/problems/n-queens/description/
 *
 * algorithms
 * Hard (45.66%)
 * Likes:    1941
 * Dislikes: 73
 * Total Accepted:    205.2K
 * Total Submissions: 442.3K
 * Testcase Example:  '4'
 *
 * The n-queens puzzle is the problem of placing n queens on an n×n chessboard
 * such that no two queens attack each other.
 *
 *
 *
 * Given an integer n, return all distinct solutions to the n-queens puzzle.
 *
 * Each solution contains a distinct board configuration of the n-queens'
 * placement, where 'Q' and '.' both indicate a queen and an empty space
 * respectively.
 *
 * Example:
 *
 *
 * Input: 4
 * Output: [
 * ⁠[".Q..",  // Solution 1
 * ⁠ "...Q",
 * ⁠ "Q...",
 * ⁠ "..Q."],
 *
 * ⁠["..Q.",  // Solution 2
 * ⁠ "Q...",
 * ⁠ "...Q",
 * ⁠ ".Q.."]
 * ]
 * Explanation: There exist two distinct solutions to the 4-queens puzzle as
 * shown above.
 *
 *
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	return solveNQueens1(n)
}

// 回溯方法，还未完全理解，待进行深入学习
func solveNQueens1(n int) [][]string {
	res := make([][]string, 0)
	sol := initialSolution(n)

	backtrack(&res, sol, 0, n)
	return res
}

func backtrack(res *[][]string, sol [][]byte, col int, n int) {
	if col == n {
		collectSolution(res, sol)
		return
	}
	for i := 0; i < n; i++ {
		if isValid(sol, i, col, n) {
			sol[i][col] = 'Q'
			backtrack(res, sol, col+1, n)
			sol[i][col] = '.'
		}
	}
}

func initialSolution(n int) [][]byte {
	res := make([][]byte, 0)
	for i := 0; i < n; i++ {
		tmp := make([]byte, n)
		for j := range tmp {
			tmp[j] = '.'
		}
		res = append(res, tmp)
	}
	return res
}

func collectSolution(res *[][]string, sol [][]byte) {
	tmp := make([]string, 0)
	for i := range sol {
		tmp = append(tmp, string(sol[i]))
	}
	*res = append(*res, tmp)
}

func isValid(tmp [][]byte, row, col, n int) bool {
	for i := 0; i < col; i++ {
		if tmp[row][i] == 'Q' {
			return false
		}
	}
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if tmp[i][j] == 'Q' {
			return false
		}
	}
	for i, j := row, col; i < n && j >= 0; i, j = i+1, j-1 {
		if tmp[i][j] == 'Q' {
			return false
		}
	}
	return true
}

// @lc code=end