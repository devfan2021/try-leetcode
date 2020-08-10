/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
 *
 * https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (44.86%)
 * Likes:    1663
 * Dislikes: 33
 * Total Accepted:    220.5K
 * Total Submissions: 485.4K
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * Given inorder and postorder traversal of a tree, construct the binary tree.
 *
 * Note:
 * You may assume that duplicates do not exist in the tree.
 *
 * For example, given
 *
 *
 * inorder = [9,3,15,20,7]
 * postorder = [9,15,7,20,3]
 *
 * Return the following binary tree:
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	// find the root value's index in the inorder array
	var index int
	for i, v := range inorder {
		if v == postorder[len(postorder)-1] {
			index = i
			break
		}
	}

	// This is how to devide the original array by the index we get
	// post_left, post_right :=postorder[0:idx], postorder[idx: -1]
	// in_left, in_right := inorder[0: idx], inorder[idx+1:]
	return &TreeNode{Val: postorder[len(postorder)-1], Left: buildTree(inorder[0:index], postorder[0:index]), Right: buildTree(inorder[index+1:], postorder[index:len(postorder)-1])}
}

// @lc code=end