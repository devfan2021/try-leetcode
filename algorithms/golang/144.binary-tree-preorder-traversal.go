/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
 *
 * https://leetcode.com/problems/binary-tree-preorder-traversal/description/
 *
 * algorithms
 * Medium (55.11%)
 * Likes:    1555
 * Dislikes: 58
 * Total Accepted:    502.1K
 * Total Submissions: 905.2K
 * Testcase Example:  '[1,null,2,3]'
 *
 * Given a binary tree, return the preorder traversal of its nodes' values.
 *
 * Example:
 *
 *
 * Input: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 *
 * Output: [1,2,3]
 *
 *
 * Follow up: Recursive solution is trivial, could you do it iteratively?
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
func preorderTraversal(root *TreeNode) []int {
	return preorderTraversal2(root)
}

func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	retVal := []int{}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		// pop top node
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		retVal = append(retVal, top.Val)
		// do first check right
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
	return retVal
}

// do with recursion
func preorderTraversal1(root *TreeNode) []int {
	vals := []int{}
	doRecursive(root, &vals)
	return vals
}

func doRecursive(root *TreeNode, vals *[]int) {
	if root != nil {
		*vals = append(*vals, root.Val)
		if root.Left != nil {
			doRecursive(root.Left, vals)
		}
		if root.Right != nil {
			doRecursive(root.Right, vals)
		}
	}
}

// @lc code=end