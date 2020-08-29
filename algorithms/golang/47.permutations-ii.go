import "sort"

/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 *
 * https://leetcode.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (46.58%)
 * Likes:    2089
 * Dislikes: 63
 * Total Accepted:    366.7K
 * Total Submissions: 785.5K
 * Testcase Example:  '[1,1,2]'
 *
 * Given a collection of numbers that might contain duplicates, return all
 * possible unique permutations.
 *
 * Example:
 *
 *
 * Input: [1,1,2]
 * Output:
 * [
 * ⁠ [1,1,2],
 * ⁠ [1,2,1],
 * ⁠ [2,1,1]
 * ]
 *
 *
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	return permuteUnique1(nums)
}

func permuteUnique1(nums []int) [][]int {
	sort.Ints(nums)
	output, track := [][]int{}, []int{}
	backtrack(&output, nums, track, 0)
	return output
}

func backtrack(output *[][]int, nums, track []int, index int) {
	if len(track) == len(nums) {
		path := make([]int, len(nums))
		copy(path, track)
		*output = append(*output, path)
		return
	}

	for i := 0; i < len(nums); i++ {
		if i == index {
			continue
		}
		track = append(track, nums[i])
		backtrack(output, nums, track, index+1)
		track = track[:len(track)-1]
	}
}

// @lc code=end