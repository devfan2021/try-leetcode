/*
 * @lc app=leetcode id=783 lang=golang
 *
 * [783] Minimum Distance Between BST Nodes
 *
 * https://leetcode.com/problems/minimum-distance-between-bst-nodes/description/
 *
 * algorithms
 * Easy (52.24%)
 * Likes:    711
 * Dislikes: 200
 * Total Accepted:    64.4K
 * Total Submissions: 122.4K
 * Testcase Example:  '[4,2,6,1,3,null,null]'
 *
 * Given a Binary Search Tree (BST) with the root node root, return the minimum
 * difference between the values of any two different nodes in the tree.
 *
 * Example :
 *
 *
 * Input: root = [4,2,6,1,3,null,null]
 * Output: 1
 * Explanation:
 * Note that root is a TreeNode object, not an array.
 *
 * The given tree [4,2,6,1,3,null,null] is represented by the following
 * diagram:
 *
 * ⁠         4
 * ⁠       /   \
 * ⁠     2      6
 * ⁠    / \
 * ⁠   1   3
 *
 * while the minimum difference in this tree is 1, it occurs between node 1 and
 * node 2, also between node 3 and node 2.
 *
 *
 * Note:
 *
 *
 * The size of the BST will be between 2 and 100.
 * The BST is always valid, each node's value is an integer, and each node's
 * value is different.
 * This question is the same as 530:
 * https://leetcode.com/problems/minimum-absolute-difference-in-bst/
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
func minDiffInBST(root *TreeNode) int {
	return minDiffInBST2(root)
}

// dfs, direct compare each node value
func minDiffInBST2(root *TreeNode) int {
	var dfs func(root *TreeNode)
	prev, minVal := -1, 2<<31-1

	dfs = func(root *TreeNode) {
		if root != nil {
			dfs(root.Left)
			if prev != -1 {
				if root.Val-prev < minVal {
					minVal = root.Val - prev
				}
			}
			prev = root.Val
			dfs(root.Right)
		}
	}
	dfs(root)
	return minVal
}

// dfs + save each node vals, compare each node value to get minimum value
func minDiffInBST1(root *TreeNode) int {
	vals := []int{}
	dfs1(root, &vals)
	minVal := 2<<31 - 1
	for i := 1; i < len(vals); i++ {
		minVal = min(minVal, vals[i]-vals[i-1])
	}
	return minVal
}

func dfs1(root *TreeNode, vals *[]int) {
	if root == nil {
		return
	}

	dfs1(root.Left, vals)
	*vals = append(*vals, root.Val)
	dfs1(root.Right, vals)
}

func min(val1, val2 int) int {
	if val1 < val2 {
		return val1
	}
	return val2
}

// @lc code=end