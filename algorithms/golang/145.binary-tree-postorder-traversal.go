/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
 *
 * https://leetcode.com/problems/binary-tree-postorder-traversal/description/
 *
 * algorithms
 * Hard (54.15%)
 * Likes:    1797
 * Dislikes: 94
 * Total Accepted:    383.5K
 * Total Submissions: 701.5K
 * Testcase Example:  '[1,null,2,3]'
 *
 * Given a binary tree, return the postorder traversal of its nodes' values.
 *
 * Example:
 *
 *
 * Input: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 *
 * Output: [3,2,1]
 *
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
func postorderTraversal(root *TreeNode) []int {
	return postorderTraversal2(root)
}

// do by stack
func postorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	retVal, stack, visited := []int{}, []*TreeNode{root}, map[*TreeNode]bool{}
	visited[root] = true
	for len(stack) > 0 {
		emptyFlag, top := true, stack[len(stack)-1]
		if top.Right != nil && !visited[top.Right] {
			visited[top.Right] = true
			stack = append(stack, top.Right)
			emptyFlag = false
		}
		if top.Left != nil && !visited[top.Left] {
			visited[top.Left] = true
			stack = append(stack, top.Left)
			emptyFlag = false
		}
		if emptyFlag {
			top = stack[len(stack)-1]
			if _, ok := visited[top]; ok {
				stack = stack[:len(stack)-1]
				retVal = append(retVal, top.Val)
			}
		}
	}
	return retVal
}

// do by recursion
func postorderTraversal1(root *TreeNode) []int {
	retVal := []int{}
	doRecursive(root, &retVal)
	return retVal
}

func doRecursive(root *TreeNode, vals *[]int) {
	if root != nil {
		if root.Left != nil {
			doRecursive(root.Left, vals)
		}
		if root.Right != nil {
			doRecursive(root.Right, vals)
		}
		*vals = append(*vals, root.Val)
	}
}

// @lc code=end