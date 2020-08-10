import "math"

/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 *
 * https://leetcode.com/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (27.57%)
 * Likes:    4023
 * Dislikes: 539
 * Total Accepted:    708.6K
 * Total Submissions: 2.6M
 * Testcase Example:  '[2,1,3]'
 *
 * Given a binary tree, determine if it is a valid binary search tree (BST).
 *
 * Assume a BST is defined as follows:
 *
 *
 * The left subtree of a node contains only nodes with keys less than the
 * node's key.
 * The right subtree of a node contains only nodes with keys greater than the
 * node's key.
 * Both the left and right subtrees must also be binary search trees.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * ⁠   2
 * ⁠  / \
 * ⁠ 1   3
 *
 * Input: [2,1,3]
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * ⁠   5
 * ⁠  / \
 * ⁠ 1   4
 * / \
 * 3   6
 *
 * Input: [5,1,4,null,null,3,6]
 * Output: false
 * Explanation: The root node's value is 5 but its right child's value is 4.
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
// [10,5,15,null,null,6,20]
func isValidBST(root *TreeNode) bool {
	return isValidBST2(root)
}

// solution by tree inorder
func isValidBST2(root *TreeNode) bool {
	prev, stack := math.MinInt64, []*TreeNode{}

	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pop.Val <= prev {
				return false
			}
			prev = pop.Val
			root = pop.Right
		}
	}

	return true
}

// solution by recursive
func isValidBST1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	res, v := true, root.Val
	helper(&res, root.Left, &v, nil)
	helper(&res, root.Right, nil, &v)
	return res
}

func helper(res *bool, root *TreeNode, upperBound, lowerBound *int) {
	if root == nil || !*res {
		return
	}

	if (upperBound != nil && root.Val >= *upperBound) || (lowerBound != nil && root.Val <= *lowerBound) {
		*res = false
		return
	}

	v := root.Val
	helper(res, root.Left, &v, lowerBound)
	helper(res, root.Right, upperBound, &v)
}

// @lc code=end