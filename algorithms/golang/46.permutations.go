/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 *
 * https://leetcode.com/problems/permutations/description/
 *
 * algorithms
 * Medium (62.54%)
 * Likes:    4079
 * Dislikes: 106
 * Total Accepted:    627.6K
 * Total Submissions: 988.7K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a collection of distinct integers, return all possible permutations.
 *
 * Example:
 *
 *
 * Input: [1,2,3]
 * Output:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 *
 *
 */

// @lc code=start
func permute(nums []int) [][]int {
	return permute1(nums)
}

func permute1(nums []int) [][]int {
	output, track := [][]int{}, []int{}
	backtrack(&output, nums, track)
	return output
}

func backtrack(output *[][]int, nums, track []int) {
	if len(track) == len(nums) {
		path := make([]int, len(track))
		copy(path, track)
		*output = append(*output, path)
		return
	}

	for i := 0; i < len(nums); i++ {
		if contains(track, nums[i]) {
			continue
		}

		track = append(track, nums[i])
		backtrack(output, nums, track)
		track = track[:len(track)-1]
	}
}

func contains(nums []int, value int) bool {
	if len(nums) == 0 {
		return false
	}
	for _, v := range nums {
		if v == value {
			return true
		}
	}
	return false
}

// @lc code=end