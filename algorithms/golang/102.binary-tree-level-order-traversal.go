/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (53.80%)
 * Likes:    3083
 * Dislikes: 76
 * Total Accepted:    623.3K
 * Total Submissions: 1.1M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the level order traversal of its nodes' values.
 * (ie, from left to right, level by level).
 *
 *
 * For example:
 * Given binary tree [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 *
 * return its level order traversal as:
 *
 * [
 * ⁠ [3],
 * ⁠ [9,20],
 * ⁠ [15,7]
 * ]
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
func levelOrder(root *TreeNode) [][]int {
	return levelOrder2(root)
}

// solution with stack
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	retVal, stack := [][]int{}, []*TreeNode{root}
	for len(stack) > 0 {
		layNode, nextStack := []int{}, []*TreeNode{}
		for i := 0; i < len(stack); i++ {
			node := stack[i]
			if node.Left != nil {
				nextStack = append(nextStack, node.Left)
			}
			if node.Right != nil {
				nextStack = append(nextStack, node.Right)
			}
			layNode = append(layNode, node.Val)
		}
		stack = nextStack
		retVal = append(retVal, layNode)
	}
	return retVal
}

// solution with recursion
func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	retVal := [][]int{}
	levelOrderByRecursion(root, &retVal, 0)
	return retVal
}

func levelOrderByRecursion(root *TreeNode, nodeVals *[][]int, level int) {
	if root == nil {
		return
	}

	if len(*nodeVals) <= level {
		*nodeVals = append(*nodeVals, []int{})
	}
	(*nodeVals)[level] = append((*nodeVals)[level], root.Val)

	levelOrderByRecursion(root.Left, nodeVals, level+1)
	levelOrderByRecursion(root.Right, nodeVals, level+1)
}

// @lc code=end