/*
 * @lc app=leetcode id=938 lang=golang
 *
 * [938] Range Sum of BST
 *
 * https://leetcode.com/problems/range-sum-of-bst/description/
 *
 * algorithms
 * Easy (80.78%)
 * Likes:    1298
 * Dislikes: 228
 * Total Accepted:    230.2K
 * Total Submissions: 283.2K
 * Testcase Example:  '[10,5,15,3,7,null,18]\n7\n15'
 *
 * Given the root node of a binary search tree, return the sum of values of all
 * nodes with value between L and R (inclusive).
 *
 * The binary search tree is guaranteed to have unique values.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: root = [10,5,15,3,7,null,18], L = 7, R = 15
 * Output: 32
 *
 *
 *
 * Example 2:
 *
 *
 * Input: root = [10,5,15,3,7,13,18,1,null,6], L = 6, R = 10
 * Output: 23
 *
 *
 *
 *
 * Note:
 *
 *
 * The number of nodes in the tree is at most 10000.
 * The final answer is guaranteed to be less than 2^31.
 *
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
func rangeSumBST(root *TreeNode, L int, R int) int {
	return rangeSumBST2(root, L, R)
}

// iterative implementation
func rangeSumBST2(root *TreeNode, L int, R int) int {
	sum, stack := 0, []*TreeNode{root}
	for len(stack) > 0 {
		nextStack := []*TreeNode{}
		for _, v := range stack {
			if v.Val >= L && v.Val <= R {
				sum += v.Val
			}

			if v.Left != nil {
				nextStack = append(nextStack, v.Left)
			}
			if v.Right != nil {
				nextStack = append(nextStack, v.Right)
			}
		}
		stack = nextStack
	}
	return sum
}

// dfs
func rangeSumBST1(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	sum := 0
	if root.Val >= L && root.Val <= R {
		sum += root.Val
	}

	if root.Left != nil {
		sum += rangeSumBST(root.Left, L, R)
	}
	if root.Right != nil {
		sum += rangeSumBST(root.Right, L, R)
	}

	return sum
}

// @lc code=end