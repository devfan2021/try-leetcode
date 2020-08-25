import "fmt"

/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
 *
 * https://leetcode.com/problems/binary-tree-paths/description/
 *
 * algorithms
 * Easy (51.57%)
 * Likes:    1856
 * Dislikes: 111
 * Total Accepted:    329.6K
 * Total Submissions: 638.5K
 * Testcase Example:  '[1,2,3,null,5]'
 *
 * Given a binary tree, return all root-to-leaf paths.
 *
 * Note: A leaf is a node with no children.
 *
 * Example:
 *
 *
 * Input:
 *
 * ⁠  1
 * ⁠/   \
 * 2     3
 * ⁠\
 * ⁠ 5
 *
 * Output: ["1->2->5", "1->3"]
 *
 * Explanation: All root-to-leaf paths are: 1->2->5, 1->3
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
func binaryTreePaths(root *TreeNode) []string {
	return binaryTreePaths1(root)
}

// recursive dfs
func binaryTreePaths1(root *TreeNode) []string {
	retVals := []string{}
	dfs(root, "", &retVals)
	return retVals
}

func dfs(root *TreeNode, path string, paths *[]string) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		nextPath := fmt.Sprintf("%s%d", path, root.Val)
		*paths = append(*paths, nextPath)
	} else {
		nextPath := fmt.Sprintf("%s%d->", path, root.Val)
		if root.Left != nil {
			dfs(root.Left, nextPath, paths)
		}

		if root.Right != nil {
			dfs(root.Right, nextPath, paths)
		}
	}
}

// @lc code=end