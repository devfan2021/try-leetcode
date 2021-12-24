/*
 * @lc app=leetcode id=654 lang=golang
 *
 * [654] Maximum Binary Tree
 *
 * https://leetcode.com/problems/maximum-binary-tree/description/
 *
 * algorithms
 * Medium (82.85%)
 * Likes:    3209
 * Dislikes: 289
 * Total Accepted:    193.1K
 * Total Submissions: 233K
 * Testcase Example:  '[3,2,1,6,0,5]'
 *
 * You are given an integer array nums with no duplicates. A maximum binary
 * tree can be built recursively from nums using the following algorithm:
 *
 *
 * Create a root node whose value is the maximum value in nums.
 * Recursively build the left subtree on the subarray prefix to the left of the
 * maximum value.
 * Recursively build the right subtree on the subarray suffix to the right of
 * the maximum value.
 *
 *
 * Return the maximum binary tree built from nums.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [3,2,1,6,0,5]
 * Output: [6,3,5,null,2,0,null,null,1]
 * Explanation: The recursive calls are as follow:
 * - The largest value in [3,2,1,6,0,5] is 6. Left prefix is [3,2,1] and right
 * suffix is [0,5].
 * ⁠   - The largest value in [3,2,1] is 3. Left prefix is [] and right suffix
 * is [2,1].
 * ⁠       - Empty array, so no child.
 * ⁠       - The largest value in [2,1] is 2. Left prefix is [] and right
 * suffix is [1].
 * ⁠           - Empty array, so no child.
 * ⁠           - Only one element, so child is a node with value 1.
 * ⁠   - The largest value in [0,5] is 5. Left prefix is [0] and right suffix
 * is [].
 * ⁠       - Only one element, so child is a node with value 0.
 * ⁠       - Empty array, so no child.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [3,2,1]
 * Output: [3,null,2,null,1]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 1000
 * 0 <= nums[i] <= 1000
 * All integers in nums are unique.
 *
 *
 */
// don't understand the O(n) solution using map or stack

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	return solution2(nums)
}

func solution1(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	maxIndex := maxValIndex(nums)
	// init tree node
	root := &TreeNode{nums[maxIndex], nil, nil}
	recursiveMaxTree(root, nums[:maxIndex], nums[maxIndex+1:])
	return root
}

func solution2(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxIndex := maxValIndex(nums)
	root := &TreeNode{nums[maxIndex], nil, nil}
	root.Left = solution2(nums[:maxIndex])
	root.Right = solution2(nums[maxIndex+1:])
	return root
}

func recursiveMaxTree(root *TreeNode, lNums []int, rNums []int) {
	lMaxIndex, rMaxIndex := maxValIndex(lNums), maxValIndex(rNums)
	if lMaxIndex == -1 && rMaxIndex == -1 {
		return
	}

	if lMaxIndex > -1 {
		root.Left = &TreeNode{lNums[lMaxIndex], nil, nil}
		recursiveMaxTree(root.Left, lNums[:lMaxIndex], lNums[lMaxIndex+1:])
	}

	if rMaxIndex > -1 {
		root.Right = &TreeNode{rNums[rMaxIndex], nil, nil}
		recursiveMaxTree(root.Right, rNums[:rMaxIndex], rNums[rMaxIndex+1:])
	}
}

func maxValIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	// find max value
	maxIndex, maxVal := 0, 0
	for index, val := range nums {
		if val > maxVal {
			maxIndex, maxVal = index, val
		}
	}
	return maxIndex
}

// @lc code=end