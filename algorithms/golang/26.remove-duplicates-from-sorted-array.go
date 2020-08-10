/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 *
 * https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
 *
 * algorithms
 * Easy (44.66%)
 * Likes:    2488
 * Dislikes: 5091
 * Total Accepted:    973.3K
 * Total Submissions: 2.2M
 * Testcase Example:  '[1,1,2]'
 *
 * Given a sorted array nums, remove the duplicates in-place such that each
 * element appear only once and return the new length.
 *
 * Do not allocate extra space for another array, you must do this by modifying
 * the input array in-place with O(1) extra memory.
 *
 * Example 1:
 *
 *
 * Given nums = [1,1,2],
 *
 * Your function should return length = 2, with the first two elements of nums
 * being 1 and 2 respectively.
 *
 * It doesn't matter what you leave beyond the returned length.
 *
 * Example 2:
 *
 *
 * Given nums = [0,0,1,1,1,2,2,3,3,4],
 *
 * Your function should return length = 5, with the first five elements of nums
 * being modified to 0, 1, 2, 3, and 4 respectively.
 *
 * It doesn't matter what values are set beyond the returned length.
 *
 *
 * Clarification:
 *
 * Confused why the returned value is an integer but your answer is an array?
 *
 * Note that the input array is passed in by reference, which means
 * modification to the input array will be known to the caller as well.
 *
 * Internally you can think of this:
 *
 *
 * // nums is passed in by reference. (i.e., without making a copy)
 * int len = removeDuplicates(nums);
 *
 * // any modification to nums in your function would be known by the caller.
 * // using the length returned by your function, it prints the first len
 * elements.
 * for (int i = 0; i < len; i++) {
 * print(nums[i]);
 * }
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	return removeDuplicates3(nums)
}

// 因为是已排序数组，可以直接进行比较
func removeDuplicates3(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	index := 1
	cmpVal := nums[0]
	for i := 1; i < len(nums); i++ {
		if cmpVal != nums[i] {
			nums[index] = nums[i]
			cmpVal = nums[i]
			index++
		}
	}
	nums = nums[0:index]
	return len(nums)
}

// 因为是已排序数组，可以直接进行比较
// 增加一个curVal变量，执行效率是removeDuplicates3的2倍
func removeDuplicates2(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	index := 1
	cmpVal := nums[0]
	for i := 1; i < len(nums); i++ {
		curVal := nums[i]
		if cmpVal != curVal {
			nums[index] = curVal
			cmpVal = curVal
			index++
		}
	}
	nums = nums[0:index]
	return len(nums)
}

// 借助一个Map进行缓存
func removeDuplicates1(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	index := 0
	valMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		curVal := nums[i]
		_, ok := valMap[curVal]
		if !ok {
			nums[index] = nums[i]
			index++
			valMap[curVal] = 1
		}
	}
	nums = nums[0:index]
	return len(nums)
}

// @lc code=end

