/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
 *
 * https://leetcode.com/problems/balanced-binary-tree/description/
 *
 * algorithms
 * Easy (43.18%)
 * Likes:    2397
 * Dislikes: 176
 * Total Accepted:    461.6K
 * Total Submissions: 1.1M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, determine if it is height-balanced.
 *
 * For this problem, a height-balanced binary tree is defined as:
 *
 *
 * a binary tree in which the left and right subtrees of every node differ in
 * height by no more than 1.
 *
 *
 *
 *
 * Example 1:
 *
 * Given the following tree [3,9,20,null,null,15,7]:
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * Return true.
 *
 * Example 2:
 *
 * Given the following tree [1,2,2,3,3,null,null,4,4]:
 *
 *
 * ⁠      1
 * ⁠     / \
 * ⁠    2   2
 * ⁠   / \
 * ⁠  3   3
 * ⁠ / \
 * ⁠4   4
 *
 *
 * Return false.
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
func isBalanced(root *TreeNode) bool {
	return isBalanced2(root)
}

// https://leetcode.com/problems/balanced-binary-tree/discuss/35691/The-bottom-up-O(N)-solution-would-be-better
// [1,null,2,null,3]
// [1,2,2,3,3,3,3,4,4,4,4,4,4,null,null,5,5]
// time complexity:O(N)
func isBalanced2(root *TreeNode) bool {
	return dfsHeight(root) != -1
}

func dfsHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := dfsHeight(root.Left)
	if leftHeight == -1 {
		return -1
	}
	rightHeight := dfsHeight(root.Right)
	if rightHeight == -1 {
		return -1
	}
	if abs(leftHeight, rightHeight) > 1 {
		return -1
	}
	return max(leftHeight, rightHeight) + 1
}

// time complexity: O(N*N)
func isBalanced1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	return abs(left, right) <= 1 && isBalanced1(root.Left) && isBalanced1(root.Right)
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(dfs(root.Left), dfs(root.Right)) + 1
}

func abs(val1, val2 int) int {
	minus := val1 - val2
	if minus > 0 {
		return minus
	}
	return -minus
}

func max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

// @lc code=end