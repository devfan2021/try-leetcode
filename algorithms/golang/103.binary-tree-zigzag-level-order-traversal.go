/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
 *
 * https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/description/
 *
 * algorithms
 * Medium (52.40%)
 * Likes:    5624
 * Dislikes: 168
 * Total Accepted:    691K
 * Total Submissions: 1.3M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given the root of a binary tree, return the zigzag level order traversal of
 * its nodes' values. (i.e., from left to right, then right to left for the
 * next level and alternate between).
 *
 *
 * Example 1:
 *
 *
 * Input: root = [3,9,20,null,null,15,7]
 * Output: [[3],[20,9],[15,7]]
 *
 *
 * Example 2:
 *
 *
 * Input: root = [1]
 * Output: [[1]]
 *
 *
 * Example 3:
 *
 *
 * Input: root = []
 * Output: []
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [0, 2000].
 * -100 <= Node.val <= 100
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	return solution1(root)
}

func solution1(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	retVal, stack, level := [][]int{}, []*TreeNode{root}, 0
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

		if level%2 == 0 {
			retVal = append(retVal, layNode)
		} else {
			reversNode := []int{}
			for i := len(layNode) - 1; i >= 0; i-- {
				reversNode = append(reversNode, layNode[i])
			}
			retVal = append(retVal, reversNode)
		}
		level++
	}
	return retVal
}

func solution2(root *TreeNode) [][]int {
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
	if level%2 == 0 {
		*nodeVals[level] = append(*nodeVals[level], root.Val)
	} else {
		*nodeVals[level] = append([]int{root.Val}, *nodeVals[level]...)
	}
	levelOrderByRecursion(root.Left, nodeVals, level+1)
	levelOrderByRecursion(root.Right, nodeVals, level+1)
}

// @lc code=end