/*
 * @lc app=leetcode id=228 lang=golang
 *
 * [228] Summary Ranges
 *
 * https://leetcode.com/problems/summary-ranges/description/
 *
 * algorithms
 * Easy (44.31%)
 * Likes:    1273
 * Dislikes: 831
 * Total Accepted:    241.4K
 * Total Submissions: 544.9K
 * Testcase Example:  '[0,1,2,4,5,7]'
 *
 * You are given a sorted unique integer array nums.
 *
 * Return the smallest sorted list of ranges that cover all the numbers in the
 * array exactly. That is, each element of nums is covered by exactly one of
 * the ranges, and there is no integer x such that x is in one of the ranges
 * but not in nums.
 *
 * Each range [a,b] in the list should be output as:
 *
 *
 * "a->b" if a != b
 * "a" if a == b
 *
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [0,1,2,4,5,7]
 * Output: ["0->2","4->5","7"]
 * Explanation: The ranges are:
 * [0,2] --> "0->2"
 * [4,5] --> "4->5"
 * [7,7] --> "7"
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [0,2,3,4,6,8,9]
 * Output: ["0","2->4","6","8->9"]
 * Explanation: The ranges are:
 * [0,0] --> "0"
 * [2,4] --> "2->4"
 * [6,6] --> "6"
 * [8,9] --> "8->9"
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= nums.length <= 20
 * -2^31 <= nums[i] <= 2^31 - 1
 * All the values of nums are unique.
 * nums is sorted in ascending order.
 *
 *
 */

// @lc code=start
func summaryRanges(nums []int) []string {
	return solution1(nums)
}

// slice, slice append, string format,
func solution1(nums []int) []string {
	var retVals []string
	if len(nums) == 0 {
		return retVals
	}

	if len(nums) == 1 {
		return append(retVals, fmt.Sprintf("%d", nums[0]))
	}

	for begin, cmpVal, index := nums[0], nums[0], 1; index <= len(nums)-1; index++ {
		if (cmpVal + 1) == nums[index] {
			cmpVal = nums[index]
		} else {
			if begin == cmpVal {
				retVals = append(retVals, fmt.Sprintf("%d", begin))
			} else {
				retVals = append(retVals, fmt.Sprintf("%d->%d", begin, cmpVal))
			}
			begin, cmpVal = nums[index], nums[index]
		}

		if index == len(nums)-1 {
			if begin == cmpVal {
				retVals = append(retVals, fmt.Sprintf("%d", begin))
			} else {
				retVals = append(retVals, fmt.Sprintf("%d->%d", begin, cmpVal))
			}
		}
	}
	return retVals
}

// simple method, using array index
func solution2(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	var retVals []string
	head := 0
	for index := range nums {
		if index < len(nums)-1 && (nums[index]+1) == nums[index+1] {
			continue
		}

		if head == index {
			retVals = append(retVals, strconv.Itoa(nums[index]))
		} else {
			retVals = append(retVals, strconv.Itoa(nums[head])+"->"+strconv.Itoa(nums[index]))
		}
		head = index + 1
	}
	return retVals
}

// @lc code=end