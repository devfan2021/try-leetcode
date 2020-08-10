/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
 *
 * https://leetcode.com/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (46.45%)
 * Likes:    4213
 * Dislikes: 104
 * Total Accepted:    662.9K
 * Total Submissions: 1.4M
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * Given a binary tree, check whether it is a mirror of itself (ie, symmetric
 * around its center).
 *
 * For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
 *
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠/ \ / \
 * 3  4 4  3
 *
 *
 *
 *
 * But the following [1,2,2,null,3,null,3] is not:
 *
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠  \   \
 * ⁠  3    3
 *
 *
 *
 *
 * Follow up: Solve it both recursively and iteratively.
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
func isSymmetric(root *TreeNode) bool {
	return isSymmetric1(root)
}

func isSymmetric2(root *TreeNode) bool {
	stack := []*TreeNode{root, root}
	for len(stack) > 0 {
		node1, node2 := stack[len(stack)-1], stack[len(stack)-2]
		if node1 == nil && node2 == nil {
			stack = stack[:len(stack)-2]
		} else {
			if node1 == nil || node2 == nil {
				return false
			}
			if node1.Val != node2.Val {
				return false
			}
			stack = stack[:len(stack)-2]
			stack = append(stack, node1.Left)
			stack = append(stack, node2.Right)

			stack = append(stack, node1.Right)
			stack = append(stack, node2.Left)
		}
	}
	return true
}

// solution by recursion, time complexity: O(n), space complexity: O(n)
func isSymmetric1(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	return node1.Val == node2.Val && isMirror(node1.Left, node2.Right) && isMirror(node1.Right, node2.Left)
}

// @lc code=end