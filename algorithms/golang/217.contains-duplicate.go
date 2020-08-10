import "sort"

/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 *
 * https://leetcode.com/problems/contains-duplicate/description/
 *
 * algorithms
 * Easy (55.71%)
 * Likes:    860
 * Dislikes: 726
 * Total Accepted:    564.9K
 * Total Submissions: 1M
 * Testcase Example:  '[1,2,3,1]'
 *
 * Given an array of integers, find if the array contains any duplicates.
 *
 * Your function should return true if any value appears at least twice in the
 * array, and it should return false if every element is distinct.
 *
 * Example 1:
 *
 *
 * Input: [1,2,3,1]
 * Output: true
 *
 * Example 2:
 *
 *
 * Input: [1,2,3,4]
 * Output: false
 *
 * Example 3:
 *
 *
 * Input: [1,1,1,3,3,4,3,2,4,2]
 * Output: true
 *
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	return containsDuplicate1(nums)
}

// 先排序，再比较
func containsDuplicate1(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return false
	}

	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

// 借助map
func containsDuplicate2(nums []int) bool {
	valMap := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if valMap[nums[i]] {
			return true
		}
		valMap[nums[i]] = true
	}
	return false
}

// @lc code=end

