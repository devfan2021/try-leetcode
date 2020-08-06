/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 *
 * https://leetcode.com/problems/triangle/description/
 *
 * algorithms
 * Medium (43.54%)
 * Likes:    2085
 * Dislikes: 246
 * Total Accepted:    252.4K
 * Total Submissions: 572.3K
 * Testcase Example:  '[[2],[3,4],[6,5,7],[4,1,8,3]]'
 *
 * Given a triangle, find the minimum path sum from top to bottom. Each step
 * you may move to adjacent numbers on the row below.
 *
 * For example, given the following triangle
 *
 *
 * [
 * ⁠    [2],
 * ⁠   [3,4],
 * ⁠  [6,5,7],
 * ⁠ [4,1,8,3]
 * ]
 *
 *
 * The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).
 *
 * Note:
 *
 * Bonus point if you are able to do this using only O(n) extra space, where n
 * is the total number of rows in the triangle.
 *
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	return minimumTotal2(triangle)
}

type point struct {
	x, y int
}

func minimumTotal2(triangle [][]int) int {
	cache := map[point]int{}
	return dfsWithMemory(triangle, 0, 0, cache)
}

func dfsWithMemory(triangle [][]int, i, j int, cache map[point]int) int {
	if i == len(triangle) {
		return 0
	}

	p := point{i, j}
	if v, ok := cache[p]; ok {
		return v
	}

	left, right := dfsWithMemory(triangle, i+1, j, cache), dfsWithMemory(triangle, i+1, j+1, cache)
	cur := min(left, right) + triangle[i][j]
	cache[p] = cur
	return cur
}

func minimumTotal1(triangle [][]int) int {
	return dfs(triangle, 0, 0)
}

func dfs(triangle [][]int, i, j int) int {
	if i == len(triangle) {
		return 0
	}

	adjacent1, adjacent2 := dfs(triangle, i+1, j), dfs(triangle, i+1, j+1)
	return min(adjacent1, adjacent2) + triangle[i][j]
}

func min(val1, val2 int) int {
	if val1 < val2 {
		return val1
	}
	return val2
}

// @lc code=end