/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal-ii/description/
 *
 * algorithms
 * Easy (51.46%)
 * Likes:    1556
 * Dislikes: 216
 * Total Accepted:    361.3K
 * Total Submissions: 674.7K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the bottom-up level order traversal of its
 * nodes' values. (ie, from left to right, level by level from leaf to root).
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
 * return its bottom-up level order traversal as:
 *
 * [
 * ⁠ [15,7],
 * ⁠ [9,20],
 * ⁠ [3]
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
func levelOrderBottom(root *TreeNode) [][]int {
	return levelOrderBottom2(root)
}

// dfs solution
func levelOrderBottom2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	retVals, stack := [][]int{}, []*TreeNode{root}
	for len(stack) > 0 {
		nextStack := []*TreeNode{}
		retVals = append(retVals, []int{})
		for i := 0; i < len(stack); i++ {
			curNode := stack[i]
			retVals[len(retVals)-1] = append(retVals[len(retVals)-1], curNode.Val)

			if curNode.Left != nil {
				nextStack = append(nextStack, curNode.Left)
			}
			if curNode.Right != nil {
				nextStack = append(nextStack, curNode.Right)
			}
		}
		stack = nextStack
	}
	for i, j := 0, len(retVals)-1; i < j; i, j = i+1, j-1 {
		retVals[i], retVals[j] = retVals[j], retVals[i]
	}
	return retVals
}

// dfs solution
func levelOrderBottom1(root *TreeNode) [][]int {
	retVals, level := [][]int{}, 0
	dfs(root, &retVals, level)
	for i, j := 0, len(retVals)-1; i < j; i, j = i+1, j-1 {
		retVals[i], retVals[j] = retVals[j], retVals[i]
	}
	return retVals
}

func dfs(root *TreeNode, retVals *[][]int, level int) {
	if root == nil {
		return
	}
	if len(*retVals) <= level {
		*retVals = append(*retVals, []int{})
	}

	if root.Left != nil {
		dfs(root.Left, retVals, level+1)
	}
	if root.Right != nil {
		dfs(root.Right, retVals, level+1)
	}

	(*retVals)[level] = append((*retVals)[level], root.Val)
}

// @lc code=end