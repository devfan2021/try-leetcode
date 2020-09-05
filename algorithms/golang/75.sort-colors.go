/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 *
 * https://leetcode.com/problems/sort-colors/description/
 *
 * algorithms
 * Medium (47.45%)
 * Likes:    3853
 * Dislikes: 235
 * Total Accepted:    545.3K
 * Total Submissions: 1.1M
 * Testcase Example:  '[2,0,2,1,1,0]'
 *
 * Given an array with n objects colored red, white or blue, sort them in-place
 * so that objects of the same color are adjacent, with the colors in the order
 * red, white and blue.
 *
 * Here, we will use the integers 0, 1, and 2 to represent the color red,
 * white, and blue respectively.
 *
 * Note: You are not suppose to use the library's sort function for this
 * problem.
 *
 * Example:
 *
 *
 * Input: [2,0,2,1,1,0]
 * Output: [0,0,1,1,2,2]
 *
 * Follow up:
 *
 *
 * A rather straight forward solution is a two-pass algorithm using counting
 * sort.
 * First, iterate the array counting number of 0's, 1's, and 2's, then
 * overwrite array with total number of 0's, then 1's and followed by 2's.
 * Could you come up with a one-pass algorithm using only constant space?
 *
 *
 */

// @lc code=start
func sortColors(nums []int) {
	sortColors2(nums)
}

func sortColors2(nums []int) {
	if len(nums) == 0 {
		return
	}

	zero, two, index := 0, len(nums)-1, 0
	for index <= two {
		if nums[index] == 2 {
			nums[index], nums[two] = nums[two], nums[index]
			two--
		} else if nums[index] == 0 {
			nums[index], nums[zero] = nums[zero], nums[index]
			zero++
			index++
		} else {
			index++
		}
	}
}

func sortColors1(nums []int) {
	if len(nums) == 0 {
		return
	}

	hash := map[int]int{}
	for _, v := range nums {
		hash[v]++
	}

	index := 0
	for i := 0; i <= 2; i++ {
		for hash[i] > 0 {
			hash[i]--
			nums[index] = i
			index++
		}
	}
}

// @lc code=end