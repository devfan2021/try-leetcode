/*
 * @lc app=leetcode id=542 lang=golang
 *
 * [542] 01 Matrix
 *
 * https://leetcode.com/problems/01-matrix/description/
 *
 * algorithms
 * Medium (39.36%)
 * Likes:    1486
 * Dislikes: 107
 * Total Accepted:    87.7K
 * Total Submissions: 220.9K
 * Testcase Example:  '[[0,0,0],[0,1,0],[0,0,0]]'
 *
 * Given a matrix consists of 0 and 1, find the distance of the nearest 0 for
 * each cell.
 *
 * The distance between two adjacent cells is 1.
 *
 *
 *
 * Example 1:
 *
 *
 * Input:
 * [[0,0,0],
 * ⁠[0,1,0],
 * ⁠[0,0,0]]
 *
 * Output:
 * [[0,0,0],
 * [0,1,0],
 * [0,0,0]]
 *
 *
 * Example 2:
 *
 *
 * Input:
 * [[0,0,0],
 * ⁠[0,1,0],
 * ⁠[1,1,1]]
 *
 * Output:
 * [[0,0,0],
 * ⁠[0,1,0],
 * ⁠[1,2,1]]
 *
 *
 *
 *
 * Note:
 *
 *
 * The number of elements of the given matrix will not exceed 10,000.
 * There are at least one 0 in the given matrix.
 * The cells are adjacent in only four directions: up, down, left and right.
 *
 *
 */

// @lc code=start
func updateMatrix(matrix [][]int) [][]int {
	return updateMatrix1(matrix)
}

// using bfs
func updateMatrix1(matrix [][]int) [][]int {
	dist := make([][]int, len(matrix))
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return dist
	}

	queue := make([][]int, 0)
	for i := 0; i < len(matrix); i++ {
		dist[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				dist[i][j] = 0
				queue = append(queue, []int{i, j})
			} else {
				dist[i][j] = -1 // using visited flag
			}
		}
	}

	maxRows, maxColums := len(matrix), len(matrix[0])
	directions := [][]int{{-1, 0}, {0, 1}, {0, -1}, {1, 0}}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		x, y := top[0], top[1]
		for _, v := range directions {
			nextX, nextY := x+v[0], y+v[1]
			if nextX >= 0 && nextX < maxRows && nextY >= 0 && nextY < maxColums && dist[nextX][nextY] == -1 {
				dist[nextX][nextY] = dist[x][y] + 1
				queue = append(queue, []int{nextX, nextY})
			}
		}
	}
	return dist
}

// @lc code=end