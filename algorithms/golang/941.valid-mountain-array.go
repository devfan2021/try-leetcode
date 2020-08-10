/*
 * @lc app=leetcode id=941 lang=golang
 *
 * [941] Valid Mountain Array
 *
 * https://leetcode.com/problems/valid-mountain-array/description/
 *
 * algorithms
 * Easy (34.03%)
 * Likes:    380
 * Dislikes: 66
 * Total Accepted:    57.8K
 * Total Submissions: 170.9K
 * Testcase Example:  '[2,1]'
 *
 * Given an array A of integers, return true if and only if it is a valid
 * mountain array.
 *
 * Recall that A is a mountain array if and only if:
 *
 *
 * A.length >= 3
 * There exists some i with 0 < i < A.length - 1 such that:
 *
 * A[0] < A[1] < ... A[i-1] < A[i]
 * A[i] > A[i+1] > ... > A[A.length - 1]
 *
 *
 *
 *
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [2,1]
 * Output: false
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [3,5,5]
 * Output: false
 *
 *
 *
 * Example 3:
 *
 *
 * Input: [0,3,2,1]
 * Output: true
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * 0 <= A.length <= 10000
 * 0 <= A[i] <= 10000
 *
 *
 *
 *
 *
 *
 *
 *
 *
 */

// @lc code=start
func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	// 第一个和第二个相等，或者第一个一直大于后续
	if A[0] == A[1] || A[0] > A[1] {
		return false
	}

	isUp := true
	cmpVal := A[0]
	for i := 1; i < len(A); i++ {
		if A[i] == cmpVal {
			return false
		} else {
			// 上升趋势
			if isUp {
				if A[i] < cmpVal {
					isUp = false
				}
			} else {
				// 下降趋势
				if A[i] > cmpVal {
					return false
				}
			}
			cmpVal = A[i]
		}
	}
	if !isUp {
		return true
	}
	return false
}

// @lc code=end

