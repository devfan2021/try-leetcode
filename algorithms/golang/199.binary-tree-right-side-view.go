/*
 * @lc app=leetcode id=199 lang=golang
 *
 * [199] Binary Tree Right Side View
 *
 * https://leetcode.com/problems/binary-tree-right-side-view/description/
 *
 * algorithms
 * Medium (53.36%)
 * Likes:    2445
 * Dislikes: 142
 * Total Accepted:    306.1K
 * Total Submissions: 565.3K
 * Testcase Example:  '[1,2,3,null,5,null,4]'
 *
 * Given a binary tree, imagine yourself standing on the right side of it,
 * return the values of the nodes you can see ordered from top to bottom.
 *
 * Example:
 *
 *
 * Input: [1,2,3,null,5,null,4]
 * Output: [1, 3, 4]
 * Explanation:
 *
 * ⁠  1            <---
 * ⁠/   \
 * 2     3         <---
 * ⁠\     \
 * ⁠ 5     4       <---
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
func rightSideView(root *TreeNode) []int {
	return rightSideView1(root)
}

// [1,2,3,4]
func rightSideView2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	retVal, stack := []int{}, []*TreeNode{root}
	for len(stack) > 0 {
		lastNode := stack[len(stack)-1]
		retVal = append(retVal, lastNode.Val)

		nextStack := []*TreeNode{}
		for i := 0; i < len(stack); i++ {
			curNode := stack[i]
			if curNode.Left != nil {
				nextStack = append(nextStack, curNode.Left)
			}
			if curNode.Right != nil {
				nextStack = append(nextStack, curNode.Right)
			}
		}
		stack = nextStack
	}
	return retVal
}

func rightSideView1(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	retVal := []int{}
	helper(root, &retVal, 0)
	return retVal
}

func helper(root *TreeNode, retVal *[]int, depth int) {
	if root == nil {
		return
	}

	if len(*retVal) == depth {
		*retVal = append(*retVal, root.Val)
	}

	helper(root.Right, retVal, depth+1)
	helper(root.Left, retVal, depth+1)
}

// @lc code=end