/*
 * @lc app=leetcode id=977 lang=golang
 *
 * [977] Squares of a Sorted Array
 *
 * https://leetcode.com/problems/squares-of-a-sorted-array/description/
 *
 * algorithms
 * Easy (72.28%)
 * Likes:    1098
 * Dislikes: 85
 * Total Accepted:    227.7K
 * Total Submissions: 315.1K
 * Testcase Example:  '[-4,-1,0,3,10]'
 *
 * Given an array of integers A sorted in non-decreasing order, return an array
 * of the squares of each number, also in sorted non-decreasing order.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [-4,-1,0,3,10]
 * Output: [0,1,9,16,100]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [-7,-3,2,3,11]
 * Output: [4,9,9,49,121]
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= A.length <= 10000
 * -10000 <= A[i] <= 10000
 * A is sorted in non-decreasing order.
 *
 *
 *
 */

// @lc code=start
func sortedSquares(A []int) []int {
	return sortedSquares2(A)
}

func sortedSquares2(A []int) []int {
	sortArr := make([]int, len(A), len(A))
	i, j, k := 0, len(A)-1, len(A)-1
	for ; i <= j; k-- {
		a, b := A[i]*A[i], A[j]*A[j]
		if a >= b {
			sortArr[k] = a
			i++
		} else {
			sortArr[k] = b
			j--
		}
	}
	return sortArr
}

func sortedSquares1(A []int) []int {
	sortArr := make([]int, len(A))

	for index, val := range A {
		sortArr[index] = val * val
	}

	// 数组排序
	quickSort(sortArr)
	return sortArr
}

func quickSort(A []int) {
	quickSortPart(A, 0, len(A)-1)
}

func quickSortPart(A []int, start, end int) {
	if (end - start) < 1 {
		return
	}

	pivot := A[end]
	splitIndex := start
	for i := start; i < end; i++ {
		if A[i] < pivot {
			tmp := A[splitIndex]
			A[splitIndex] = A[i]
			A[i] = tmp

			splitIndex++
		}
	}

	A[end] = A[splitIndex]
	A[splitIndex] = pivot

	quickSortPart(A, start, splitIndex-1)
	quickSortPart(A, splitIndex+1, end)
}

// @lc code=end
