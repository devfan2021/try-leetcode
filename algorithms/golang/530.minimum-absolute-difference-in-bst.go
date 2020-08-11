import (
	"sort"
)

/*
 * @lc app=leetcode id=530 lang=golang
 *
 * [530] Minimum Absolute Difference in BST
 *
 * https://leetcode.com/problems/minimum-absolute-difference-in-bst/description/
 *
 * algorithms
 * Easy (53.31%)
 * Likes:    888
 * Dislikes: 68
 * Total Accepted:    91.6K
 * Total Submissions: 170.3K
 * Testcase Example:  '[1,null,3,2]'
 *
 * Given a binary search tree with non-negative values, find the minimum
 * absolute difference between values of any two nodes.
 *
 * Example:
 *
 *
 * Input:
 *
 * ⁠  1
 * ⁠   \
 * ⁠    3
 * ⁠   /
 * ⁠  2
 *
 * Output:
 * 1
 *
 * Explanation:
 * The minimum absolute difference is 1, which is the difference between 2 and
 * 1 (or between 2 and 3).
 *
 *
 *
 *
 * Note:
 *
 *
 * There are at least two nodes in this BST.
 * This question is the same as 783:
 * https://leetcode.com/problems/minimum-distance-between-bst-nodes/
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	return getMinimumDifference3(root)
}

// https://leetcode.com/problems/minimum-absolute-difference-in-bst/discuss/376454/Go-golang-2-solutions
func getMinimumDifference3(root *TreeNode) int {
	var dfs func(root *TreeNode)
	prev, minVal := -1, 2<<31-1

	dfs = func(root *TreeNode) {
		if root != nil {
			dfs(root.Left)
			if prev != -1 {
				if root.Val-prev < minVal {
					minVal = root.Val - prev
				}
			}
			prev = root.Val
			dfs(root.Right)
		}
	}
	dfs(root)
	return minVal
}

// dfs, inorder
func getMinimumDifference2(root *TreeNode) int {
	cache := []int{}
	dfs2(root, &cache)
	minVal := 2<<31 - 1
	for i := 0; i < len(cache)-1; i++ {
		minVal = min(minVal, cache[i+1]-cache[i])
	}
	return minVal
}

func dfs2(root *TreeNode, vals *[]int) {
	if root == nil {
		return
	}
	dfs2(root.Left, vals)
	*vals = append(*vals, root.Val)
	dfs2(root.Right, vals)
}

// dfs + sort: time complexity: nlog(n)
func getMinimumDifference1(root *TreeNode) int {
	cache := []int{}
	dfs(root, &cache)
	sort.Ints(cache)
	minVal := 2<<31 - 1
	for i := 0; i < len(cache)-1; i++ {
		minVal = min(minVal, cache[i+1]-cache[i])
	}
	return minVal
}

func dfs(root *TreeNode, vals *[]int) {
	if root == nil {
		return
	}
	*vals = append(*vals, root.Val)
	dfs(root.Left, vals)
	dfs(root.Right, vals)
}

func min(val1, val2 int) int {
	if val1 < val2 {
		return val1
	}
	return val2
}

// @lc code=end