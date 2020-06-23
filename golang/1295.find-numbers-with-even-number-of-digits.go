/*
 * @lc app=leetcode id=1295 lang=golang
 *
 * [1295] Find Numbers with Even Number of Digits
 *
 * https://leetcode.com/problems/find-numbers-with-even-number-of-digits/description/
 *
 * algorithms
 * Easy (82.83%)
 * Likes:    321
 * Dislikes: 51
 * Total Accepted:    113.5K
 * Total Submissions: 137.1K
 * Testcase Example:  '[12,345,2,6,7896]'
 *
 * Given an array nums of integers, return how many of them contain an even
 * number of digits.
 *
 * Example 1:
 *
 *
 * Input: nums = [12,345,2,6,7896]
 * Output: 2
 * Explanation:
 * 12 contains 2 digits (even number of digits).
 * 345 contains 3 digits (odd number of digits).
 * 2 contains 1 digit (odd number of digits).
 * 6 contains 1 digit (odd number of digits).
 * 7896 contains 4 digits (even number of digits).
 * Therefore only 12 and 7896 contain an even number of digits.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [555,901,482,1771]
 * Output: 1
 * Explanation:
 * Only 1771 contains an even number of digits.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 500
 * 1 <= nums[i] <= 10^5
 *
 *
 */

import "strconv"

// @lc code=start
func findNumbers(nums []int) int {
	return findNumbers2(nums)
}

func findNumbers2(nums []int) int {
	count := 0
	for _, val := range nums {
		digit := 0
		for val >= 1 {
			val /= 10
			digit++
		}

		if digit%2 == 0 {
			count++
		}
	}
	return count
}

func findNumbers1(nums []int) int {
	total := 0
	for _, val := range nums {
		valStr := strconv.Itoa(val)
		if val != 0 && len(valStr)%2 == 0 {
			total++
		}
	}
	return total
}

// @lc code=end

