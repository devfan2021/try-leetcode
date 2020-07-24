/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
 *
 * https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (65.41%)
 * Likes:    2567
 * Dislikes: 75
 * Total Accepted:    836.4K
 * Total Submissions: 1.3M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, find its maximum depth.
 *
 * The maximum depth is the number of nodes along the longest path from the
 * root node down to the farthest leaf node.
 *
 * Note: A leaf is a node with no children.
 *
 * Example:
 *
 * Given binary tree [3,9,20,null,null,15,7],
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * return its depth = 3.
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
func maxDepth(root *TreeNode) int {
	return maxDepth3(root)
}

// solution by recursion, simple method
func maxDepth3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth3(root.Left), maxDepth3(root.Right)) + 1
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

// solution by stack, BFS
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	visited, stack, level := map[*TreeNode]bool{}, []*TreeNode{root}, 0
	visited[root] = true

	for len(stack) > 0 {
		level++
		nextStack := []*TreeNode{}
		for _, v := range stack {
			if v.Left != nil && !visited[v.Left] {
				visited[v.Left] = true
				nextStack = append(nextStack, v.Left)
			}
			if v.Right != nil && !visited[v.Right] {
				visited[v.Right] = true
				nextStack = append(nextStack, v.Right)
			}
		}
		stack = nextStack
	}
	return level
}

// solution by recursion
func maxDepth1(root *TreeNode) int {
	return maxDepthByRecursion(root, 0)
}

func maxDepthByRecursion(root *TreeNode, level int) int {
	if root != nil {
		leftLevel, rightLevel := level+1, level+1
		if root.Left != nil {
			tmp := maxDepthByRecursion(root.Left, level+1)
			if tmp > leftLevel {
				leftLevel = tmp
			}
		}
		if root.Right != nil {
			tmp := maxDepthByRecursion(root.Right, level+1)
			if tmp > rightLevel {
				rightLevel = tmp
			}
		}
		if leftLevel > rightLevel {
			return leftLevel
		} else {
			return rightLevel
		}
	}
	return level
}

// @lc code=end