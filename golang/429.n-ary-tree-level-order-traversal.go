/*
 * @lc app=leetcode id=429 lang=golang
 *
 * [429] N-ary Tree Level Order Traversal
 *
 * https://leetcode.com/problems/n-ary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (64.21%)
 * Likes:    621
 * Dislikes: 46
 * Total Accepted:    78.3K
 * Total Submissions: 120.5K
 * Testcase Example:  '[1,null,3,2,4,null,5,6]'
 *
 * Given an n-ary tree, return the level order traversal of its nodes' values.
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
 * Output: [[1],[3,2,4],[5,6]]
 *
 *
 * Example 2:
 *
 *
 *
 *
 * Input: root =
 * [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
 * Output: [[1],[2,3,4,5],[6,7,8,9,10],[11,12,13],[14]]
 *
 *
 *
 * Constraints:
 *
 *
 * The height of the n-ary tree is less than or equal to 1000
 * The total number of nodes is between [0, 10^4]
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

func levelOrder(root *Node) [][]int {
	return levelOrder1(root)
}

// bfs
func levelOrder2(root *Node) [][]int {
	retVals := [][]int{}
	if root == nil {
		return retVals
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		nextStack := []*Node{}
		curRet := []int{}
		for i := 0; i < len(stack); i++ {
			curNode := stack[i]
			if len(curNode.Children) > 0 {
				nextStack = append(nextStack, curNode.Children...)
			}
			curRet = append(curRet, curNode.Val)
		}
		stack = nextStack
		retVals = append(retVals, curRet)
	}
	return retVals
}

// dfs
func levelOrder1(root *Node) [][]int {
	retVals := [][]int{}
	if root == nil {
		return retVals
	}
	dfs(root, 0, &retVals)
	return retVals
}

func dfs(root *Node, level int, retVals *[][]int) {
	if len(*retVals) <= level {
		*retVals = append(*retVals, []int{root.Val})
	} else {
		(*retVals)[level] = append((*retVals)[level], root.Val)
	}

	childs := root.Children
	if len(childs) > 0 {
		for i := 0; i < len(childs); i++ {
			dfs(childs[i], level+1, retVals)
		}
	}
}

// @lc code=end