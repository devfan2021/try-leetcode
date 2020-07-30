import "sort"

/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
 *
 * https://leetcode.com/problems/first-missing-positive/description/
 *
 * algorithms
 * Hard (31.60%)
 * Likes:    3635
 * Dislikes: 811
 * Total Accepted:    344.2K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,0]'
 *
 * Given an unsorted integer array, find the smallest missing positive
 * integer.
 *
 * Example 1:
 *
 *
 * Input: [1,2,0]
 * Output: 3
 *
 *
 * Example 2:
 *
 *
 * Input: [3,4,-1,1]
 * Output: 2
 *
 *
 * Example 3:
 *
 *
 * Input: [7,8,9,11,12]
 * Output: 1
 *
 *
 * Note:
 *
 * Your algorithm should run in O(n) time and uses constant extra space.
 *
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	return firstMissingPositive1(nums)
}

//[], [0,2,2,1,1]
func firstMissingPositive1(nums []int) int {
	// handle: []
	if len(nums) == 0 {
		return 1
	}

	// sort and find first postive number index
	sort.Ints(nums)
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			index = i
			break
		}
	}

	// check first postive number index is postive or may be value on
	if nums[index] < 0 || (nums[index] > 0 && nums[index] != 1) {
		return 1
	}

	pre := nums[index]
	for i := index + 1; i < len(nums); i++ {
		if pre == nums[i] { // handle: 0,1,1,2,2
			continue
		}

		if nums[i] != (pre + 1) {
			return pre + 1
		}
		pre = nums[i]
	}
	return pre + 1
}

// @lc code=end