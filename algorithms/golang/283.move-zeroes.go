/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 *
 * https://leetcode.com/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (57.54%)
 * Likes:    3731
 * Dislikes: 121
 * Total Accepted:    825.9K
 * Total Submissions: 1.4M
 * Testcase Example:  '[0,1,0,3,12]'
 *
 * Given an array nums, write a function to move all 0's to the end of it while
 * maintaining the relative order of the non-zero elements.
 *
 * Example:
 *
 *
 * Input: [0,1,0,3,12]
 * Output: [1,3,12,0,0]
 *
 * Note:
 *
 *
 * You must do this in-place without making a copy of the array.
 * Minimize the total number of operations.
 *
 */

// @lc code=start
func moveZeroes(nums []int) {
	moveZeroes2(nums)
}

// 简化版本，直接两两交互
func moveZeroes2(nums []int) {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != j { // 如果i,j相同，不用进行替换
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
}

func moveZeroes1(nums []int) {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if count != i {
				nums[count] = nums[i]
			}
			count++
		}
	}

	for i := count; i < len(nums); i++ {
		nums[i] = 0
	}
}

// @lc code=end

