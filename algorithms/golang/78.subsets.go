/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 *
 * https://leetcode.com/problems/subsets/description/
 *
 * algorithms
 * Medium (60.12%)
 * Likes:    4015
 * Dislikes: 85
 * Total Accepted:    604.2K
 * Total Submissions: 975.9K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a set of distinct integers, nums, return all possible subsets (the
 * power set).
 *
 * Note: The solution set must not contain duplicate subsets.
 *
 * Example:
 *
 *
 * Input: nums = [1,2,3]
 * Output:
 * [
 * ⁠ [3],
 * [1],
 * [2],
 * [1,2,3],
 * [1,3],
 * [2,3],
 * [1,2],
 * []
 * ]
 *
 */

// @lc code=start
func subsets(nums []int) [][]int {
	return subsets1(nums)
}

func subsets1(nums []int) [][]int {
	output := [][]int{}
	track := []int{} // save visit path
	backtrack(&output, nums, []int{}, 0)
	return output
}

func backtrack(output *[][]int, nums []int, track []int, start int) {
	// should do copy and save current path
	path := make([]int, len(track))
	copy(path, track)
	*output = append(*output, path)

	for i := start; i < len(nums); i++ {
		// do action
		track = append(track, nums[i])
		// do backtrack for next i+1
		backtrack(output, nums, track, i+1)
		// roback add action
		track = track[:len(track)-1]
	}
}

// @lc code=end