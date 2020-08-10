import "sort"

/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 *
 * https://leetcode.com/problems/single-number/description/
 *
 * algorithms
 * Easy (65.15%)
 * Likes:    4444
 * Dislikes: 160
 * Total Accepted:    878.8K
 * Total Submissions: 1.3M
 * Testcase Example:  '[2,2,1]'
 *
 * Given a non-empty array of integers, every element appears twice except for
 * one. Find that single one.
 *
 * Note:
 *
 * Your algorithm should have a linear runtime complexity. Could you implement
 * it without using extra memory?
 *
 * Example 1:
 *
 *
 * Input: [2,2,1]
 * Output: 1
 *
 *
 * Example 2:
 *
 *
 * Input: [4,1,2,1,2]
 * Output: 4
 *
 *
 */

// @lc code=start
func singleNumber(nums []int) int {
	return singleNumber2(nums)
}

// a ^ a ^ b = 0 ^ b = b,  0与任何数异或等于原数
func singleNumber1(nums []int) int {
	val := 0
	for _, v := range nums {
		val ^= v
	}
	return val
}

// 直接用map进行处理, 增加，删除，最后一个即为所求值
func singleNumber2(nums []int) int {
	valMap := make(map[int]bool)
	for _, val := range nums {
		_, ok := valMap[val]
		if ok {
			delete(valMap, val)
		} else {
			valMap[val] = true
		}
	}
	for key, _ := range valMap {
		return key
	}
	return 0
}

// 排序，比较前后两个值
func singleNumber3(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i += 2 {
		if nums[i] == nums[i+1] {
			continue
		}
		return nums[i]
	}
	return nums[len(nums)-1]
}

// @lc code=end

