/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 *
 * https://leetcode.com/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (46.29%)
 * Likes:    5701
 * Dislikes: 195
 * Total Accepted:    739.7K
 * Total Submissions: 1.6M
 * Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
 *
 * Given a 2d grid map of '1's (land) and '0's (water), count the number of
 * islands. An island is surrounded by water and is formed by connecting
 * adjacent lands horizontally or vertically. You may assume all four edges of
 * the grid are all surrounded by water.
 *
 *
 * Example 1:
 *
 *
 * Input: grid = [
 * ⁠ ["1","1","1","1","0"],
 * ⁠ ["1","1","0","1","0"],
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["0","0","0","0","0"]
 * ]
 * Output: 1
 *
 *
 * Example 2:
 *
 *
 * Input: grid = [
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["0","0","1","0","0"],
 * ⁠ ["0","0","0","1","1"]
 * ]
 * Output: 3
 *
 *
 */
// 给出一个0-1矩阵，1代表陆地，0代表水。相连的1可被看作一块陆地。求陆地的数量。
// @lc code=start
func numIslands(grid [][]byte) int {
	return numIslands2(grid)
}

// 采用dfs访问，标记已经访问的节点, 空间换时间
func numIslands2(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	result := 0
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				result++
				dfsMark2WithVisited(grid, visited, i, j)
			}
		}
	}
	return result
}

func dfsMark2WithVisited(grid [][]byte, visited [][]bool, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
		return
	}
	visited[i][j] = true
	if grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'
	dfsMark2WithVisited(grid, visited, i-1, j)
	dfsMark2WithVisited(grid, visited, i+1, j)
	dfsMark2WithVisited(grid, visited, i, j-1)
	dfsMark2WithVisited(grid, visited, i, j+1)
}

// 采用dfs访问，但是里面有太多重复的访问判断
func numIslands1(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				result++
				dfsMark1(grid, i, j)
			}
		}
	}
	return result
}

func dfsMark1(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfsMark1(grid, i-1, j)
	dfsMark1(grid, i+1, j)
	dfsMark1(grid, i, j-1)
	dfsMark1(grid, i, j+1)
}

// @lc code=end