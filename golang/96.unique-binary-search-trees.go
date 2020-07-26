/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 *
 * https://leetcode.com/problems/unique-binary-search-trees/description/
 *
 * algorithms
 * Medium (51.05%)
 * Likes:    3470
 * Dislikes: 124
 * Total Accepted:    307.1K
 * Total Submissions: 582.8K
 * Testcase Example:  '3'
 *
 * Given n, how many structurally unique BST's (binary search trees) that store
 * values 1 ... n?
 *
 * Example:
 *
 *
 * Input: 3
 * Output: 5
 * Explanation:
 * Given n = 3, there are a total of 5 unique BST's:
 *
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 19
 *
 *
 */

// @lc code=start
type position struct {
	start, end int
}

func numTrees(n int) int {
	if n == 0 {
		return 0
	}
	cache := map[position]int{}
	return helper(1, n, cache)
}

func helper(start, end int, cache map[position]int) int {
	if start > end {
		return 1
	}
	key := position{start, end}
	if val, ok := cache[key]; ok {
		return val
	}

	retVal := 0
	for i := start; i <= end; i++ {
		retVal = retVal + helper(start, i-1, cache)*helper(i+1, end, cache)
	}
	cache[key] = retVal
	return retVal
}

// @lc code=end