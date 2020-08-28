/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 *
 * https://leetcode.com/problems/combination-sum/description/
 *
 * algorithms
 * Medium (55.22%)
 * Likes:    3974
 * Dislikes: 119
 * Total Accepted:    552.3K
 * Total Submissions: 986.9K
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * Given a set of candidate numbers (candidates) (without duplicates) and a
 * target number (target), find all unique combinations in candidates where the
 * candidate numbers sums to target.
 *
 * The same repeated number may be chosen from candidates unlimited number of
 * times.
 *
 * Note:
 *
 *
 * All numbers (including target) will be positive integers.
 * The solution set must not contain duplicate combinations.
 *
 *
 * Example 1:
 *
 *
 * Input: candidates = [2,3,6,7], target = 7,
 * A solution set is:
 * [
 * ⁠ [7],
 * ⁠ [2,2,3]
 * ]
 *
 *
 * Example 2:
 *
 *
 * Input: candidates = [2,3,5], target = 8,
 * A solution set is:
 * [
 * [2,2,2,2],
 * [2,3,3],
 * [3,5]
 * ]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= candidates.length <= 30
 * 1 <= candidates[i] <= 200
 * Each element of candidate is unique.
 * 1 <= target <= 500
 *
 *
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	return combinationSum1(candidates, target)
}

func combinationSum1(candidates []int, target int) [][]int {
	output, track := [][]int{}, []int{}
	backtrack(&output, candidates, track, target, 0)
	return output
}

func backtrack(output *[][]int, candidates, track []int, target, start int) {
	if target == 0 {
		path := make([]int, len(track))
		copy(path, track)
		*output = append(*output, path)
		return
	}
	if target < 0 {
		return
	}

	for i := start; i < len(candidates); i++ {
		remain := target - candidates[i]
		if remain >= 0 {
			track = append(track, candidates[i])
			backtrack(output, candidates, track, remain, i)
			track = track[:len(track)-1]
		}
	}
}

// @lc code=end