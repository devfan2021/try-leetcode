/*
 * @lc app=leetcode id=747 lang=golang
 *
 * [747] Largest Number At Least Twice of Others
 *
 * https://leetcode.com/problems/largest-number-at-least-twice-of-others/description/
 *
 * algorithms
 * Easy (41.70%)
 * Likes:    300
 * Dislikes: 559
 * Total Accepted:    86.4K
 * Total Submissions: 206.9K
 * Testcase Example:  '[0,0,0,1]'
 *
 * In a given integer array nums, there is always exactly one largest element.
 *
 * Find whether the largest element in the array is at least twice as much as
 * every other number in the array.
 *
 * If it is, return the index of the largest element, otherwise return -1.
 *
 * Example 1:
 *
 *
 * Input: nums = [3, 6, 1, 0]
 * Output: 1
 * Explanation: 6 is the largest integer, and for every other number in the
 * array x,
 * 6 is more than twice as big as x.  The index of value 6 is 1, so we return
 * 1.
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1, 2, 3, 4]
 * Output: -1
 * Explanation: 4 isn't at least as big as twice the value of 3, so we return
 * -1.
 *
 *
 *
 *
 * Note:
 *
 *
 * nums will have a length in the range [1, 50].
 * Every nums[i] will be an integer in the range [0, 99].
 *
 *
 *
 *
 */

// @lc code=start
func dominantIndex(nums []int) int {
	return dominantIndex2(nums)
}

// 网上其他方案， 通过一层循环搞定，定义了最大值，第二大值变量
func dominantIndex2(nums []int) int {
	if len(nums) == 0 {
		return -1
	} else if len(nums) == 1 {
		return 0
	}

	first, second := -1, -1
	firstIndex, secondIndex := 0, 0

	for i, val := range nums {
		if val > first {
			second, first = first, val
			secondIndex, firstIndex = firstIndex, i
		} else if first > val && val > second {
			second = val
			secondIndex = i
		}
	}

	if first >= nums[secondIndex]*2 {
		return firstIndex
	}
	return -1
}

// 自己思考的解法：1.遍历先求最大值  2.遍历判断是否大于等于2倍
func dominantIndex1(nums []int) int {
	index := -1
	maxVal := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
			index = i
		}
	}

	for _, val := range nums {
		if val != 0 && maxVal != val && maxVal < 2*val {
			return -1
		}
	}
	return index
}

// [0,0,0,1]
// @lc code=end

