import "sort"

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 *
 * https://leetcode.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (38.89%)
 * Likes:    4398
 * Dislikes: 298
 * Total Accepted:    615.1K
 * Total Submissions: 1.6M
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * Given a collection of intervals, merge all overlapping intervals.
 *
 * Example 1:
 *
 *
 * Input: [[1,3],[2,6],[8,10],[15,18]]
 * Output: [[1,6],[8,10],[15,18]]
 * Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into
 * [1,6].
 *
 *
 * Example 2:
 *
 *
 * Input: [[1,4],[4,5]]
 * Output: [[1,5]]
 * Explanation: Intervals [1,4] and [4,5] are considered overlapping.
 *
 * NOTE: input types have been changed on April 15, 2019. Please reset to
 * default code definition to get new method signature.
 *
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	return merge1(intervals)
}

func merge1(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	retVals := [][]int{}
	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		startI, endI := intervals[i][0], intervals[i][1]
		if end >= startI {
			if endI > end { // merge overlapping, avoid case: [1, 5], [4, 4]
				end = endI
			}
		} else {
			retVals = append(retVals, []int{start, end})
			start, end = startI, endI
		}
		if i == len(intervals)-1 {
			retVals = append(retVals, []int{start, end})
		}
	}
	return retVals
}

// @lc code=end