/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 *
 * https://leetcode.com/problems/word-search/description/
 *
 * algorithms
 * Medium (35.70%)
 * Likes:    4250
 * Dislikes: 199
 * Total Accepted:    522.2K
 * Total Submissions: 1.5M
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * Given a 2D board and a word, find if the word exists in the grid.
 *
 * The word can be constructed from letters of sequentially adjacent cell,
 * where "adjacent" cells are those horizontally or vertically neighboring. The
 * same letter cell may not be used more than once.
 *
 * Example:
 *
 *
 * board =
 * [
 * ⁠ ['A','B','C','E'],
 * ⁠ ['S','F','C','S'],
 * ⁠ ['A','D','E','E']
 * ]
 *
 * Given word = "ABCCED", return true.
 * Given word = "SEE", return true.
 * Given word = "ABCB", return false.
 *
 *
 *
 * Constraints:
 *
 *
 * board and word consists only of lowercase and uppercase English letters.
 * 1 <= board.length <= 200
 * 1 <= board[i].length <= 200
 * 1 <= word.length <= 10^3
 *
 *
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	return exist1(board, word)
}

func exist1(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if dfs(board, i, j, word, 0) {
				return true
			}
		}
	}

	return false
}

func dfs(board [][]byte, row int, col int, word string, index int) bool {
	if index == len(word) {
		return true
	}

	if row < 0 || row >= len(board) || col < 0 || col >= len(board[row]) || word[index] != board[row][col] {
		return false
	}

	tmp := board[row][col]
	// mark visited flag
	board[row][col] = '%'

	flag := dfs(board, row+1, col, word, index+1) || dfs(board, row-1, col, word, index+1) || dfs(board, row, col+1, word, index+1) || dfs(board, row, col-1, word, index+1)

	// reset visited flag
	board[row][col] = tmp

	return flag
}

// @lc code=end