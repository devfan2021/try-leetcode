/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 *
 * https://leetcode.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (34.35%)
 * Likes:    5033
 * Dislikes: 461
 * Total Accepted:    732.8K
 * Total Submissions: 2.1M
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * Suppose an array sorted in ascending order is rotated at some pivot unknown
 * to you beforehand.
 *
 * (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
 *
 * You are given a target value to search. If found in the array return its
 * index, otherwise return -1.
 *
 * You may assume no duplicate exists in the array.
 *
 * Your algorithm's runtime complexity must be in the order of O(log n).
 *
 * Example 1:
 *
 *
 * Input: nums = [4,5,6,7,0,1,2], target = 0
 * Output: 4
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [4,5,6,7,0,1,2], target = 3
 * Output: -1
 *
 */
// @lc code=start
func search(nums []int, target int) int {
	return search2(nums, target)
}

// 常规解法，直接遍历O(n)
func search2(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		}
	}
	return -1
}

func search1(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	low, high := 0, len(nums)-1
	// 二分查找
	for low <= high {
		mid := (high-low)/2 + low
		if nums[mid] == target {
			return mid
		}

		// 主要思路，确认哪一边是已确认的递增序列，通过nums[mid]与target的巧妙比较
		if nums[mid] >= nums[low] { // 说明左边是已排序好的
			if target >= nums[low] && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else { // 右边是已排序好的
			if target > nums[mid] && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

// @lc code=end

