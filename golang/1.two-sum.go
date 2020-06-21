/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 *
 * https://leetcode.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (45.47%)
 * Likes:    15290
 * Dislikes: 556
 * Total Accepted:    3M
 * Total Submissions: 6.6M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * Given an array of integers, return indices of the two numbers such that they
 * add up to a specific target.
 *
 * You may assume that each input would have exactly one solution, and you may
 * not use the same element twice.
 *
 * Example:
 *
 *
 * Given nums = [2, 7, 11, 15], target = 9,
 *
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 *
 *
 */

// @lc code=start

// 时间复杂度为n(方法与twoSum2类似)
func twoSum1(nums []int, target int) []int {
	vMap := make(map[int]int)
	for index, item := range nums {
		remainVal := target - item
		value, exist := vMap[remainVal]
		if exist {
			return []int{value, index}
		}
		vMap[item] = index
	}
	return nil
}

// 时间复杂度为n
func twoSum2(nums []int, target int) []int {
	vMap := make(map[int]int)
	for index, item := range nums {
		if _, ok := vMap[target-item]; ok {
			return []int{vMap[target-item], index}
		}
		vMap[item] = index
	}
	return []int{}
}

// 时间复杂度为n*n
func twoSum1(nums []int, target int) []int {
	result := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		var item = nums[i]
		var nextItem = target - item
		for j := i + 1; j < len(nums); j++ {
			if nextItem == nums[j] {
				result[0] = i
				result[1] = j
				break
			}
		}
	}
	return result
}

// @lc code=end