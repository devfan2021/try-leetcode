/*
 * @lc app=leetcode id=700 lang=golang
 *
 * [700] Search in a Binary Search Tree
 *
 * https://leetcode.com/problems/search-in-a-binary-search-tree/description/
 *
 * algorithms
 * Easy (72.92%)
 * Likes:    852
 * Dislikes: 118
 * Total Accepted:    192.1K
 * Total Submissions: 263K
 * Testcase Example:  '[4,2,7,1,3]\n2'
 *
 * Given the root node of a binary search tree (BST) and a value. You need to
 * find the node in the BST that the node's value equals the given value.
 * Return the subtree rooted with that node. If such node doesn't exist, you
 * should return NULL.
 *
 * For example,
 *
 *
 * Given the tree:
 * ⁠       4
 * ⁠      / \
 * ⁠     2   7
 * ⁠    / \
 * ⁠   1   3
 *
 * And the value to search: 2
 *
 *
 * You should return this subtree:
 *
 *
 * ⁠     2
 * ⁠    / \
 * ⁠   1   3
 *
 *
 * In the example above, if we want to search the value 5, since there is no
 * node with value 5, we should return NULL.
 *
 * Note that an empty tree is represented by NULL, therefore you would see the
 * expected output (serialized tree format) as [], not null.
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
func searchBST(root *TreeNode, val int) *TreeNode {
	return searchBST3(root, val)
}

// solution by iterative, use BST characteristic, left node's value less than right node's value
func searchBST3(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		}
		if val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return nil
}

// solution by bfs
func searchBST2(root *TreeNode, val int) *TreeNode {
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		nextStack := []*TreeNode{}
		for _, v := range stack {
			if v.Val == val {
				return v
			}

			if v.Left != nil {
				nextStack = append(nextStack, v.Left)
			}
			if v.Right != nil {
				nextStack = append(nextStack, v.Right)
			}
		}
		stack = nextStack
	}
	return nil
}

// solution by recursive
func searchBST1(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}

	left := searchBST1(root.Left, val)
	if left == nil {
		return searchBST1(root.Right, val)
	}
	return left
}

// @lc code=end