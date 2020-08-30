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
	output, track, visit := [][]int{}, []int{}, make([]bool, len(nums))
	backtrack(&output, nums, track, visit, 0)
	return output
}

func backtrack(output *[][]int, nums, track []int, visited []bool, index int) {
	if len(track) == len(nums) {
		path := make([]int, len(nums))
		copy(path, track)
		*output = append(*output, path)
		return
	}

	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}

		// https://leetcode.com/problems/permutations-ii/discuss/18594/Really-easy-Java-solution-much-easier-than-the-solutions-with-very-high-vote
		// With inputs as [1a, 1b, 2a],
		/// If we don't handle the duplicates, the results would be: [1a, 1b, 2a], [1b, 1a, 2a]...,
		// so we must make sure 1a goes before 1b to avoid duplicates
		// By using nums[i-1]==nums[i] && !used[i-1], we can make sure that 1b cannot be choosed before 1a
		if i > 0 && nums[i-1] == nums[i] && !visited[i-1] {
			continue
		}

		track = append(track, nums[i])
		visited[i] = true
		backtrack(output, nums, track, visited, index+1)
		track = track[:len(track)-1]
		visited[i] = false
	}
}

// @lc code=end