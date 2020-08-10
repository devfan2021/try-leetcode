/*
 * @lc app=leetcode id=117 lang=golang
 *
 * [117] Populating Next Right Pointers in Each Node II
 *
 * https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/description/
 *
 * algorithms
 * Medium (38.42%)
 * Likes:    1587
 * Dislikes: 177
 * Total Accepted:    255K
 * Total Submissions: 656.2K
 * Testcase Example:  '[1,2,3,4,5,null,7]'
 *
 * Given a binary tree
 *
 *
 * struct Node {
 * ⁠ int val;
 * ⁠ Node *left;
 * ⁠ Node *right;
 * ⁠ Node *next;
 * }
 *
 *
 * Populate each next pointer to point to its next right node. If there is no
 * next right node, the next pointer should be set to NULL.
 *
 * Initially, all next pointers are set to NULL.
 *
 *
 *
 * Follow up:
 *
 *
 * You may only use constant extra space.
 * Recursive approach is fine, you may assume implicit stack space does not
 * count as extra space for this problem.
 *
 *
 *
 * Example 1:
 *
 *
 *
 *
 * Input: root = [1,2,3,4,5,null,7]
 * Output: [1,#,2,3,#,4,5,7,#]
 * Explanation: Given the above binary tree (Figure A), your function should
 * populate each next pointer to point to its next right node, just like in
 * Figure B. The serialized output is in level order as connected by the next
 * pointers, with '#' signifying the end of each level.
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the given tree is less than 6000.
 * -100 <= node.val <= 100
 *
 *
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	return connect2(root)
}

// solution with recusive
func connect2(root *Node) *Node {
	if root == nil {
		return root
	}
	if root.Left != nil {
		if root.Right != nil {
			root.Left.Next = root.Right
		} else {
			root.Left.Next = nodeNext(root.Next)
		}
	}

	if root.Right != nil {
		root.Right.Next = nodeNext(root.Next)
	}
	connect2(root.Right)
	connect2(root.Left)
	return root
}

func nodeNext(root *Node) *Node {
	if root == nil {
		return root
	}
	if root.Left != nil {
		return root.Left
	}
	if root.Right != nil {
		return root.Right
	}
	return nodeNext(root.Next)
}

// solution with bfs, using stack
func connect1(root *Node) *Node {
	if root == nil {
		return root
	}

	stack := []*Node{root}
	for len(stack) > 0 {
		nextStack := []*Node{}
		var nextItem *Node
		for i := 0; i < len(stack); i++ {
			curNode := stack[i]
			curNode.Next = nextItem
			if curNode.Right != nil {
				nextStack = append(nextStack, curNode.Right)
			}
			if curNode.Left != nil {
				nextStack = append(nextStack, curNode.Left)
			}
			nextItem = curNode
		}
		stack = nextStack
	}
	return root
}

// @lc code=end