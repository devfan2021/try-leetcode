/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
 *
 * https://leetcode.com/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (37.14%)
 * Likes:    1531
 * Dislikes: 704
 * Total Accepted:    427.7K
 * Total Submissions: 1.1M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, find its minimum depth.
 *
 * The minimum depth is the number of nodes along the shortest path from the
 * root node down to the nearest leaf node.
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
 * return its minimum depth = 2.
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
func minDepth(root *TreeNode) int {
	return minDepth2(root)
}

func minDepth3(root *TreeNode) int {
	if root == nil {
		return 0
	}

	level, stack := 0, []*TreeNode{root}
	for len(stack) > 0 {
		level++
		nextStack := []*TreeNode{}
		for i := 0; i < len(stack); i++ {
			curNode := stack[i]
			if curNode.Left == nil && curNode.Right == nil {
				return level
			}

			if curNode.Left != nil {
				nextStack = append(nextStack, curNode.Left)
			}
			if curNode.Right != nil {
				nextStack = append(nextStack, curNode.Right)
			}
		}
		stack = nextStack
	}
	return level
}

// similar solution1, concise, clean code
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left, right := minDepth2(root.Left), minDepth2(root.Right)
	if left == 0 || right == 0 {
		return left + right + 1
	}
	return min(left, right) + 1
}

// [1,2]
func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	} else if root.Left != nil && root.Right != nil {
		return 1 + min(minDepth1(root.Left), minDepth1(root.Right))
	} else if root.Left == nil {
		return 1 + minDepth1(root.Right)
	} else {
		return 1 + minDepth1(root.Left)
	}
}

func min(val1, val2 int) int {
	if val1 < val2 {
		return val1
	}

	return val2
}

// @lc code=end