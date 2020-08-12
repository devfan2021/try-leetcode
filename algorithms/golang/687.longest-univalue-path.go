/*
 * @lc app=leetcode id=687 lang=golang
 *
 * [687] Longest Univalue Path
 *
 * https://leetcode.com/problems/longest-univalue-path/description/
 *
 * algorithms
 * Easy (35.88%)
 * Likes:    1758
 * Dislikes: 477
 * Total Accepted:    94.1K
 * Total Submissions: 259.8K
 * Testcase Example:  '[5,4,5,1,1,5]'
 *
 * Given a binary tree, find the length of the longest path where each node in
 * the path has the same value. This path may or may not pass through the
 * root.
 *
 * The length of path between two nodes is represented by the number of edges
 * between them.
 *
 *
 *
 * Example 1:
 *
 * Input:
 *
 *
 * ⁠             5
 * ⁠            / \
 * ⁠           4   5
 * ⁠          / \   \
 * ⁠         1   1   5
 *
 *
 * Output: 2
 *
 *
 *
 * Example 2:
 *
 * Input:
 *
 *
 * ⁠             1
 * ⁠            / \
 * ⁠           4   5
 * ⁠          / \   \
 * ⁠         4   4   5
 *
 *
 * Output: 2
 *
 *
 *
 * Note: The given binary tree has not more than 10000 nodes. The height of the
 * tree is not more than 1000.
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
func longestUnivaluePath(root *TreeNode) int {
	return longestUnivaluePath1(root)
}

// [1,2,2,2,2], [1,null,1,1,1,1,1,1]
func longestUnivaluePath1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxVal := 0
	arrowLength(root, &maxVal)
	return maxVal
}

func arrowLength(root *TreeNode, maxVal *int) int {
	if root == nil {
		return 0
	}
	left := arrowLength(root.Left, maxVal)
	right := arrowLength(root.Right, maxVal)
	arrowLeft, arrowRight := 0, 0
	if root.Left != nil && root.Val == root.Left.Val {
		arrowLeft = left + 1
	}
	if root.Right != nil && root.Val == root.Right.Val {
		arrowRight = right + 1
	}

	*maxVal = max(*maxVal, arrowLeft+arrowRight)
	return max(arrowLeft, arrowRight)
}

func max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

// @lc code=end