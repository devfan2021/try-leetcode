import "sort"

/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 *
 * https://leetcode.com/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (48.33%)
 * Likes:    1914
 * Dislikes: 71
 * Total Accepted:    340.6K
 * Total Submissions: 703.5K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * Given a collection of candidate numbers (candidates) and a target number
 * (target), find all unique combinations in candidates where the candidate
 * numbers sums to target.
 *
 * Each number in candidates may only be used once in the combination.
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
 * Input: candidates = [10,1,2,7,6,1,5], target = 8,
 * A solution set is:
 * [
 * ⁠ [1, 7],
 * ⁠ [1, 2, 5],
 * ⁠ [2, 6],
 * ⁠ [1, 1, 6]
 * ]
 *
 *
 * Example 2:
 *
 *
 * Input: candidates = [2,5,2,1,2], target = 5,
 * A solution set is:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
 *
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	return combinationSum21(candidates, target)
}

func combinationSum21(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	dfs(&res, candidates, []int{}, 0, target)
	return res
}

func dfs(res *[][]int, candidates, cur []int, index, target int) {
	if target < 0 {
		return
	} else if target == 0 {
		*res = append(*res, append([]int{}, cur...))
	} else {
		for i := index; i < len(candidates); i++ {
			if i == index || candidates[i] != candidates[i-1] {
				dfs(res, candidates, append(cur, candidates[i]), i+1, target-candidates[i])
			}
		}
	}
}

// @lc code=end