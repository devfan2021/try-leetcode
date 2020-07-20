import "sort"

/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 *
 * https://leetcode.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (58.23%)
 * Likes:    3264
 * Dislikes: 220
 * Total Accepted:    658.1K
 * Total Submissions: 1.1M
 * Testcase Example:  '[3,2,3]'
 *
 * Given an array of size n, find the majority element. The majority element is
 * the element that appears more than ⌊ n/2 ⌋ times.
 *
 * You may assume that the array is non-empty and the majority element always
 * exist in the array.
 *
 * Example 1:
 *
 *
 * Input: [3,2,3]
 * Output: 3
 *
 * Example 2:
 *
 *
 * Input: [2,2,1,1,1,2,2]
 * Output: 2
 *
 *
 */

// @lc code=start
func majorityElement(nums []int) int {
	return majorityElement2(nums)
}

// use sort, index: len(nums)/2
func majorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// use hash, time complexity: O(n), space complexity: O(n)
func majorityElement1(nums []int) int {
	minVal := len(nums) / 2.0
	hash := make(map[int]int, len(nums))
	for _, v := range nums {
		hash[v]++
		if hash[v] > minVal {
			return v
		}
	}
	return 0
}

// @lc code=end