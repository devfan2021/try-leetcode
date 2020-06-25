/*
 * @lc app=leetcode id=905 lang=golang
 *
 * [905] Sort Array By Parity
 *
 * https://leetcode.com/problems/sort-array-by-parity/description/
 *
 * algorithms
 * Easy (73.91%)
 * Likes:    938
 * Dislikes: 73
 * Total Accepted:    194.1K
 * Total Submissions: 262.5K
 * Testcase Example:  '[3,1,2,4]'
 *
 * Given an array A of non-negative integers, return an array consisting of all
 * the even elements of A, followed by all the odd elements of A.
 *
 * You may return any answer array that satisfies this condition.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [3,1,2,4]
 * Output: [2,4,3,1]
 * The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= A.length <= 5000
 * 0 <= A[i] <= 5000
 *
 *
 *
 */

// @lc code=start
func sortArrayByParity(A []int) []int {
	return sortArrayByParity1(A)
}

//采用位运算，重用输入参数数组， 按位与，只有1&1 == 1， 其他情况都等于0
func sortArrayByParity1(A []int) []int {
	for low, high := 0, len(A)-1; low < high; {
		for ; A[low]&1 == 0 && low < high; low++ {
		}
		for ; A[high]&1 == 1 && high > low; high-- {
		}
		A[low], A[high] = A[high], A[low]
	}
	return A
}

// 增加两个指示值，进行遍历搜索查询
func sortArrayByParity2(A []int) []int {
	retArr := make([]int, len(A), len(A))
	startIndex := 0
	endIndex := len(A) - 1
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			retArr[startIndex] = A[i]
			startIndex++
		} else {
			retArr[endIndex] = A[i]
			endIndex--
		}
	}
	return retArr
}

// @lc code=end

