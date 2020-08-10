/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
 *
 * https://leetcode.com/problems/path-sum/description/
 *
 * algorithms
 * Easy (40.75%)
 * Likes:    1950
 * Dislikes: 499
 * Total Accepted:    479.6K
 * Total Submissions: 1.2M
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,null,1]\n22'
 *
 * Given a binary tree and a sum, determine if the tree has a root-to-leaf path
 * such that adding up all the values along the path equals the given sum.
 *
 * Note: A leaf is a node with no children.
 *
 * Example:
 *
 * Given the below binary tree and sum = 22,
 *
 *
 * ⁠     5
 * ⁠    / \
 * ⁠   4   8
 * ⁠  /   / \
 * ⁠ 11  13  4
 * ⁠/  \      \
 * 7    2      1
 *
 *
 * return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.
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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	return hasPathSum1(root, sum)
}

// solution by stack
func hasPathSum2(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	stack := []*TreeNode{root}
	for len(stack) > 0 {
		nextStack := []*TreeNode{}
		for _, v := range stack {
			if v.Left == nil && v.Right == nil {
				if v.Val == sum {
					return true
				}
			} else {
				if v.Left != nil {
					v.Left.Val = v.Left.Val + v.Val
					nextStack = append(nextStack, v.Left)
				}
				if v.Right != nil {
					v.Right.Val = v.Right.Val + v.Val
					nextStack = append(nextStack, v.Right)
				}
			}
		}
		stack = nextStack
	}
	return false
}

func hasPathSum1(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && root.Val == sum {
		return true
	}
	return hasPathSum1(root.Left, sum-root.Val) || hasPathSum1(root.Right, sum-root.Val)
}

// @lc code=end