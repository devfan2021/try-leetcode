/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 *
 * https://leetcode.com/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Medium (62.54%)
 * Likes:    3189
 * Dislikes: 135
 * Total Accepted:    742.3K
 * Total Submissions: 1.2M
 * Testcase Example:  '[1,null,2,3]'
 *
 * Given a binary tree, return the inorder traversal of its nodes' values.
 *
 * Example:
 *
 *
 * Input: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 *
 * Output: [1,3,2]
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
func inorderTraversal(root *TreeNode) []int {
	return inorderTraversal2(root)
}

// using stack, time complexity:O(n), space complexity:O(n)
func inorderTraversal2(root *TreeNode) []int {
	retVal, stack := []int{}, []*TreeNode{}
	// don't add root node to stack, only assign root node to curr
	curr := root
	for curr != nil || len(stack) != 0 {
		// add left child node to the end of stack
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		retVal = append(retVal, curr.Val)
		stack = stack[:len(stack)-1] // pop
		curr = curr.Right
	}
	return retVal
}

// using recursion, time complexity:O(n), space complexity:O(n)
func inorderTraversal1(root *TreeNode) []int {
	retVal := []int{}
	traversalNode(root, &retVal)
	return retVal
}

func traversalNode(root *TreeNode, ret *[]int) {
	if root != nil {
		if root.Left != nil {
			traversalNode(root.Left, ret)
		}
		*ret = append(*ret, root.Val)
		if root.Right != nil {
			traversalNode(root.Right, ret)
		}
	}
}

// @lc code=end