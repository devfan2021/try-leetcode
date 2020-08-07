
/*
 * @lc app=leetcode id=100 lang=golang
 *
 * [100] Same Tree
 *
 * https://leetcode.com/problems/same-tree/description/
 *
 * algorithms
 * Easy (52.50%)
 * Likes:    2269
 * Dislikes: 64
 * Total Accepted:    587.6K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,3]\n[1,2,3]'
 *
 * Given two binary trees, write a function to check if they are the same or
 * not.
 *
 * Two binary trees are considered the same if they are structurally identical
 * and the nodes have the same value.
 *
 * Example 1:
 *
 *
 * Input:     1         1
 * ⁠         / \       / \
 * ⁠        2   3     2   3
 *
 * ⁠       [1,2,3],   [1,2,3]
 *
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input:     1         1
 * ⁠         /           \
 * ⁠        2             2
 *
 * ⁠       [1,2],     [1,null,2]
 *
 * Output: false
 *
 *
 * Example 3:
 *
 *
 * Input:     1         1
 * ⁠         / \       / \
 * ⁠        2   1     1   2
 *
 * ⁠       [1,2,1],   [1,1,2]
 *
 * Output: false
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
func isSameTree(p *TreeNode, q *TreeNode) bool {
	return isSameTree2(p, q)
}

func isSameTree2(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if !checkTreeNode(p, q) {
		return false
	}

	deqP, deqQ := []*TreeNode{}, []*TreeNode{}
	deqP = append(deqP, p)
	deqQ = append(deqQ, q)
	for len(deqP) > 0 {
		tmpP, tmpQ := deqP[len(deqP)-1], deqQ[len(deqQ)-1]
		deqP, deqQ = deqP[:len(deqP)-1], deqQ[:len(deqQ)-1]
		if !checkTreeNode(tmpP, tmpQ) {
			return false
		}
		if tmpP != nil {
			if !checkTreeNode(tmpP.Left, tmpQ.Left) {
				return false
			}
			if tmpP.Left != nil {
				deqP = append(deqP, tmpP.Left)
				deqQ = append(deqQ, tmpQ.Left)
			}

			if !checkTreeNode(tmpP.Right, tmpQ.Right) {
				return false
			}
			if tmpP.Right != nil {
				deqP = append(deqP, tmpP.Right)
				deqQ = append(deqQ, tmpQ.Right)
			}
		}
	}
	return true
}

func checkTreeNode(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return true
}

// recursive
func isSameTree1(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree1(p.Left, q.Left) && isSameTree1(p.Right, q.Right)
}

// @lc code=end