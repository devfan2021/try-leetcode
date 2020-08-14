/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
 *
 * https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/
 *
 * algorithms
 * Easy (57.09%)
 * Likes:    2609
 * Dislikes: 220
 * Total Accepted:    421.9K
 * Total Submissions: 727.5K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * Given an array where elements are sorted in ascending order, convert it to a
 * height balanced BST.
 *
 * For this problem, a height-balanced binary tree is defined as a binary tree
 * in which the depth of the two subtrees of every node never differ by more
 * than 1.
 *
 * Example:
 *
 *
 * Given the sorted array: [-10,-3,0,5,9],
 *
 * One possible answer is: [0,-3,9,-10,null,5], which represents the following
 * height balanced BST:
 *
 * ⁠     0
 * ⁠    / \
 * ⁠  -3   9
 * ⁠  /   /
 * ⁠-10  5
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
func sortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBST1(nums)
}

func sortedArrayToBST2(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	return &TreeNode{Val: nums[mid], Left: sortedArrayToBST2(nums[:mid]), Right: sortedArrayToBST2(nums[mid+1:])}
}

func sortedArrayToBST1(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	return dfs(nums, 0, len(nums)-1)
}

func dfs(nums []int, low, high int) *TreeNode {
	if low > high {
		return nil
	}

	mid := (low + high) / 2
	curNode := &TreeNode{Val: nums[mid]}
	curNode.Left = dfs(nums, low, mid-1)
	curNode.Right = dfs(nums, mid+1, high)
	return curNode
}

// @lc code=end