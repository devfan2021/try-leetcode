import "sort"

/*
 * @lc app=leetcode id=912 lang=golang
 *
 * [912] Sort an Array
 *
 * https://leetcode.com/problems/sort-an-array/description/
 *
 * algorithms
 * Medium (63.46%)
 * Likes:    463
 * Dislikes: 276
 * Total Accepted:    83.6K
 * Total Submissions: 131.1K
 * Testcase Example:  '[5,2,3,1]'
 *
 * Given an array of integers nums, sort the array in ascending order.
 *
 *
 * Example 1:
 * Input: nums = [5,2,3,1]
 * Output: [1,2,3,5]
 * Example 2:
 * Input: nums = [5,1,1,2,0,0]
 * Output: [0,0,1,1,2,5]
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 50000
 * -50000 <= nums[i] <= 50000
 *
 *
 */

// @lc code=start
func sortArray(nums []int) []int {
	return sortArray1(nums)
}

// use merge sort, from top voted solution, should review and study
func sortArray2(nums []int) []int {
	mergeSort(nums, 0, len(nums))
	return nums
}

// merge sort, top-down approach
func mergeSort(nums []int, start, end int) {
	if start >= end-1 {
		return
	}
	middle := (start + end) / 2
	mergeSort(nums, start, middle)
	mergeSort(nums, middle, end)
	merge(nums, start, end)
}

func merge(nums []int, start, end int) {
	arr, mid := make([]int, end-start), (start+end)/2
	for i, j, k := start, mid, 0; k < len(arr); k++ {
		if j == end || (i < mid && nums[i] < nums[j]) {
			arr[k] = nums[i]
			i++
		} else {
			arr[k] = nums[j]
			j++
		}
	}
	copy(nums[start:end], arr)
}

// use system built-in sort method
func sortArray1(nums []int) []int {
	sort.Ints(nums)
	return nums
}

// @lc code=end