/*
 * @lc app=leetcode id=162 lang=golang
 *
 * [162] Find Peak Element
 *
 * https://leetcode.com/problems/find-peak-element/description/
 *
 * algorithms
 * Medium (43.08%)
 * Likes:    1726
 * Dislikes: 2064
 * Total Accepted:    362.7K
 * Total Submissions: 840.1K
 * Testcase Example:  '[1,2,3,1]'
 *
 * A peak element is an element that is greater than its neighbors.
 *
 * Given an input array nums, where nums[i] ≠ nums[i+1], find a peak element
 * and return its index.
 *
 * The array may contain multiple peaks, in that case return the index to any
 * one of the peaks is fine.
 *
 * You may imagine that nums[-1] = nums[n] = -∞.
 *
 * Example 1:
 *
 *
 * Input: nums = [1,2,3,1]
 * Output: 2
 * Explanation: 3 is a peak element and your function should return the index
 * number 2.
 *
 * Example 2:
 *
 *
 * Input: nums = [1,2,1,3,5,6,4]
 * Output: 1 or 5
 * Explanation: Your function can return either index number 1 where the peak
 * element is 2,
 * or index number 5 where the peak element is 6.
 *
 *
 * Follow up: Your solution should be in logarithmic complexity.
 *
 */

// @lc code=start
func findPeakElement(nums []int) int {
	return findPeakElement3(nums)
}

func findPeakElement3(nums []int) int {
	return search(nums, 0, len(nums)-1)
}

func search(nums []int, left, right int) int {
	if left == right {
		return left
	}
	mid := left + (right-left)/2
	if nums[mid] > nums[mid+1] {
		return search(nums, left, mid)
	} else {
		return search(nums, mid+1, right)
	}
}

// 二分查找，注意nums[mid] > nums[mid+1]及最后返回low
func findPeakElement2(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] > nums[mid+1] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func findPeakElement1(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return len(nums) - 1
}

// @lc code=end