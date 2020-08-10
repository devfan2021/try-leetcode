import "sort"

/*
 * @lc app=leetcode id=498 lang=golang
 *
 * [498] Diagonal Traverse
 *
 * https://leetcode.com/problems/diagonal-traverse/description/
 *
 * algorithms
 * Medium (47.78%)
 * Likes:    722
 * Dislikes: 329
 * Total Accepted:    81.6K
 * Total Submissions: 170.8K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * Given a matrix of M x N elements (M rows, N columns), return all elements of
 * the matrix in diagonal order as shown in the below image.
 *
 *
 *
 * Example:
 *
 *
 * Input:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 4, 5, 6 ],
 * ⁠[ 7, 8, 9 ]
 * ]
 *
 * Output:  [1,2,4,7,5,3,6,8,9]
 *
 * Explanation:
 *
 *
 *
 *
 *
 * Note:
 *
 * The total number of elements of the given matrix will not exceed 10,000.
 *
 */

// @lc code=start
// 参考别人的实现，效率不高，待改进算法
func findDiagonalOrder(matrix [][]int) []int {
	cacheMap := make(map[int][]int)
	count := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			cacheMap[i+j] = append(cacheMap[i+j], matrix[i][j])
			count++
		}
	}

	var keys []int
	for k, _ := range cacheMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var retArr []int
	for _, val := range keys {
		child := cacheMap[val]
		if val%2 == 0 {
			for i := len(child) - 1; i >= 0; i-- {
				retArr = append(retArr, child[i])
			}
		} else {
			for i := 0; i < len(child); i++ {
				retArr = append(retArr, child[i])
			}
		}
	}
	return retArr
}

// @lc code=end
