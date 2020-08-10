/*
 * @lc app=leetcode id=559 lang=golang
 *
 * [559] Maximum Depth of N-ary Tree
 *
 * https://leetcode.com/problems/maximum-depth-of-n-ary-tree/description/
 *
 * algorithms
 * Easy (68.18%)
 * Likes:    883
 * Dislikes: 48
 * Total Accepted:    117.2K
 * Total Submissions: 170.7K
 * Testcase Example:  '[1,null,3,2,4,null,5,6]'
 *
 * Given a n-ary tree, find its maximum depth.
 *
 * The maximum depth is the number of nodes along the longest path from the
 * root node down to the farthest leaf node.
 *
 * Nary-Tree input serialization is represented in their level order traversal,
 * each group of children is separated by the null value (See examples).
 *
 *
 * Example 1:
 *
 *
 *
 *
 * Input: root = [1,null,3,2,4,null,5,6]
 * Output: 3
 *
 *
 * Example 2:
 *
 *
 *
 *
 * Input: root =
 * [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
 * Output: 5
 *
 *
 *
 * Constraints:
 *
 *
 * The depth of the n-ary tree is less than or equal to 1000.
 * The total number of nodes is between [0, 10^4].
 *
 *
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
	return maxDepth1(root)
}

func maxDepth2(root *Node) int {
	count := 0
	if root == nil {
		return count
	}

	stack := []*Node{root}
	for len(stack) > 0 {
		count++
		nextStack := []*Node{}
		for i := 0; i < len(stack); i++ {
			curNode := stack[i]
			if len(curNode.Children) > 0 {
				nextStack = append(nextStack, curNode.Children...)
			}
		}
		stack = nextStack
	}
	return count
}

func maxDepth1(root *Node) int {
	if root == nil {
		return 0
	}
	level := 0
	for _, child := range root.Children {
		level = max(level, maxDepth1(child))
	}
	return level + 1
}

func max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

// @lc code=end