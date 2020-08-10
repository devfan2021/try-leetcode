/*
 * @lc app=leetcode id=590 lang=golang
 *
 * [590] N-ary Tree Postorder Traversal
 *
 * https://leetcode.com/problems/n-ary-tree-postorder-traversal/description/
 *
 * algorithms
 * Easy (71.50%)
 * Likes:    689
 * Dislikes: 64
 * Total Accepted:    96.4K
 * Total Submissions: 133.7K
 * Testcase Example:  '[1,null,3,2,4,null,5,6]'
 *
 * Given an n-ary tree, return the postorder traversal of its nodes' values.
 *
 * Nary-Tree input serialization is represented in their level order traversal,
 * each group of children is separated by the null value (See examples).
 *
 *
 *
 * Follow up:
 *
 * Recursive solution is trivial, could you do it iteratively?
 *
 *
 * Example 1:
 *
 *
 *
 *
 * Input: root = [1,null,3,2,4,null,5,6]
 * Output: [5,6,3,2,4,1]
 *
 *
 * Example 2:
 *
 *
 *
 *
 * Input: root =
 * [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
 * Output: [2,6,14,11,7,3,12,8,4,13,9,10,5,1]
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

func postorder(root *Node) []int {
	return postorder2(root)
}

func postorder2(root *Node) []int {
	retVals := []int{}
	if root == nil {
		return retVals
	}

	stack := []*Node{root}
	for len(stack) > 0 {
		lastNode := stack[len(stack)-1]

		retVals = append(retVals, lastNode.Val)
		stack = stack[:len(stack)-1]

		childs := lastNode.Children
		if len(childs) > 0 {
			for i := 0; i < len(childs); i++ {
				stack = append(stack, childs[i])
			}
		}
	}

	// reverse result
	for i, j := 0, len(retVals)-1; i < j; i, j = i+1, j-1 {
		retVals[i], retVals[j] = retVals[j], retVals[i]
	}

	return retVals
}

func postorder1(root *Node) []int {
	retVals := []int{}
	helper(root, &retVals)
	return retVals
}

func helper(root *Node, retVals *[]int) {
	if root == nil {
		return
	}
	childs := root.Children
	for i := 0; i < len(childs); i++ {
		helper(childs[i], retVals)
	}
	*retVals = append(*retVals, root.Val)
}

// @lc code=end