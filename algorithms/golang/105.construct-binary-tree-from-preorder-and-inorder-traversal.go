/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
 *
 * https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (47.80%)
 * Likes:    3455
 * Dislikes: 93
 * Total Accepted:    364.5K
 * Total Submissions: 752.9K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * Given preorder and inorder traversal of a tree, construct the binary tree.
 *
 * Note:
 * You may assume that duplicates do not exist in the tree.
 *
 * For example, given
 *
 *
 * preorder = [3,9,20,15,7]
 * inorder = [9,3,15,20,7]
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// find the root's value index in the inorder arr
	var index int
	for i, v := range inorder {
		if v == preorder[0] {
			index = i
			break
		}
	}

	leftNode := buildTree(preorder[1:index+1], inorder[:index])
	rightNode := buildTree(preorder[index+1:], inorder[index+1:])

	return &TreeNode{Val: preorder[0], Left: leftNode, Right: rightNode}
}

// @lc code=end