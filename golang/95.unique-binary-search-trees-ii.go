/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
 *
 * https://leetcode.com/problems/unique-binary-search-trees-ii/description/
 *
 * algorithms
 * Medium (39.90%)
 * Likes:    2238
 * Dislikes: 160
 * Total Accepted:    194K
 * Total Submissions: 480.2K
 * Testcase Example:  '3'
 *
 * Given an integer n, generate all structurally unique BST's (binary search
 * trees) that store values 1 ... n.
 *
 * Example:
 *
 *
 * Input: 3
 * Output:
 * [
 * [1,null,3,2],
 * [3,2,null,1],
 * [3,1,null,null,2],
 * [2,1,3],
 * [1,null,2,null,3]
 * ]
 * Explanation:
 * The above output corresponds to the 5 unique BST's shown below:
 *
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= n <= 8
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
func generateTrees(n int) []*TreeNode {
	return generateTrees2(n)
}

type position struct {
	start, end int
}

// solution by recursive with memorization
func generateTrees2(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	cache := map[position][]*TreeNode{}
	return geneateTreesRecursiveWithMemorization(1, n, cache)
}

func geneateTreesRecursiveWithMemorization(start, end int, cache map[position][]*TreeNode) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	pos := position{start, end}
	if t, ok := cache[pos]; ok {
		return t
	}

	nodes := []*TreeNode{}
	for i := start; i <= end; i++ {
		leftNodes := geneateTreesRecursiveWithMemorization(start, i-1, cache)
		rightNodes := geneateTreesRecursiveWithMemorization(i+1, end, cache)
		for j := 0; j < len(leftNodes); j++ {
			for k := 0; k < len(rightNodes); k++ {
				root := &TreeNode{Val: i}
				root.Left = leftNodes[j]
				root.Right = rightNodes[k]
				nodes = append(nodes, root)
			}
		}
	}
	cache[pos] = nodes
	return nodes
}

// solution by recursive
func generateTrees1(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	return geneateTreesRecursive(1, n)
}

func geneateTreesRecursive(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	nodes := []*TreeNode{}
	for i := start; i <= end; i++ {
		leftNodes := geneateTreesRecursive(start, i-1)
		rightNodes := geneateTreesRecursive(i+1, end)
		for j := 0; j < len(leftNodes); j++ {
			for k := 0; k < len(rightNodes); k++ {
				root := &TreeNode{Val: i}
				root.Left = leftNodes[j]
				root.Right = rightNodes[k]
				nodes = append(nodes, root)
			}
		}
	}
	return nodes
}

// @lc code=end