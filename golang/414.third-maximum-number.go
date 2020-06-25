import (
	"math"
	"math/bits"
)

/*
 * @lc app=leetcode id=414 lang=golang
 *
 * [414] Third Maximum Number
 *
 * https://leetcode.com/problems/third-maximum-number/description/
 *
 * algorithms
 * Easy (30.39%)
 * Likes:    652
 * Dislikes: 1177
 * Total Accepted:    147.1K
 * Total Submissions: 484K
 * Testcase Example:  '[3,2,1]'
 *
 * Given a non-empty array of integers, return the third maximum number in this
 * array. If it does not exist, return the maximum number. The time complexity
 * must be in O(n).
 *
 * Example 1:
 *
 * Input: [3, 2, 1]
 *
 * Output: 1
 *
 * Explanation: The third maximum is 1.
 *
 *
 *
 * Example 2:
 *
 * Input: [1, 2]
 *
 * Output: 2
 *
 * Explanation: The third maximum does not exist, so the maximum (2) is
 * returned instead.
 *
 *
 *
 * Example 3:
 *
 * Input: [2, 2, 3, 1]
 *
 * Output: 1
 *
 * Explanation: Note that the third maximum here means the third maximum
 * distinct number.
 * Both numbers with value 2 are both considered as second maximum.
 *
 *
 */

// @lc code=start
func thirdMax(nums []int) int {
	return thirdMax2(nums)
}

func thirdMax1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}

	smallestVal := (1 << bits.UintSize) / -2
	// 3个数组以上
	first, second, third := smallestVal, smallestVal, smallestVal
	for i := 0; i < len(nums); i++ {
		if first == nums[i] || second == nums[i] || third == nums[i] {
			continue
		}

		if nums[i] > first {
			third = second
			second = first
			first = nums[i]
		} else if nums[i] > second {
			third = second
			second = nums[i]
		} else if nums[i] > third {
			third = nums[i]
		}
	}
	if third == smallestVal {
		return first
	}

	return third
}

func thirdMax2(nums []int) int {
	var m1, m2, m3 int = math.MinInt64, math.MinInt64, math.MinInt64

	for _, v := range nums {
		if v == m1 || v == m2 || v == m3 {
			continue
		}

		if m1 == math.MinInt64 || v > m1 {
			m3 = m2
			m2 = m1
			m1 = v
		} else if m2 == math.MinInt64 || v > m2 {
			m3 = m2
			m2 = v
		} else if m3 == math.MinInt64 || v > m3 {
			m3 = v
		}
	}

	if m3 == math.MinInt64 {
		return max(m1, m2)
	}

	return m3
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// [1,1,2]\n2
// @lc code=end

