/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 *
 * https://leetcode.com/problems/combinations/description/
 *
 * algorithms
 * Medium (53.85%)
 * Likes:    1531
 * Dislikes: 69
 * Total Accepted:    296.2K
 * Total Submissions: 542.1K
 * Testcase Example:  '4\n2'
 *
 * Given two integers n and k, return all possible combinations of k numbers
 * out of 1 ... n.
 *
 * Example:
 *
 *
 * Input: n = 4, k = 2
 * Output:
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 *
 *
 */

// @lc code=start
func combine(n int, k int) [][]int {
	return combine2(n, k)
}

// backtrack
func combine2(n int, k int) [][]int {
	if n == 0 || k == 0 || k > n {
		return [][]int{}
	}

	output, cur := [][]int{}, []int{}
	backtrack(&output, cur, 1, n, k)
	return output
}

func backtrack(output *[][]int, cur []int, begin, n, k int) {
	if len(cur) == k {
		// using copy for new slice
		path := make([]int, k)
		copy(path, cur)
		*output = append(*output, path)
		return
	}

	for i := begin; i < n+1; i++ {
		cur = append(cur, i)
		backtrack(output, cur, i+1, n, k)
		cur = (cur)[:len(cur)-1]
	}
}

// mathematical formula: C(n,k) = C(n-1, k-1) + C(n-1, k)
func combine1(n int, k int) [][]int {
	if n == k || k == 0 {
		childs := []int{}
		for i := 1; i <= k; i++ {
			childs = append(childs, i)
		}
		retArrs := [][]int{}
		retArrs = append(retArrs, childs)
		return retArrs
	}

	preArrs1 := combine1(n-1, k-1)
	for i, _ := range preArrs1 {
		// add last n
		preArrs1[i] = append(preArrs1[i], n)
	}

	preArrs2 := combine1(n-1, k)
	preArrs1 = append(preArrs1, preArrs2...)
	return preArrs1
}

// @lc code=end