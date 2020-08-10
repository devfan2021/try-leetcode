import (
	"sort"
)

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 *
 * https://leetcode.com/problems/3sum/description/
 *
 * algorithms
 * Medium (26.32%)
 * Likes:    6908
 * Dislikes: 813
 * Total Accepted:    912.8K
 * Total Submissions: 3.5M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * Given an array nums of n integers, are there elements a, b, c in nums such
 * that a + b + c = 0? Find all unique triplets in the array which gives the
 * sum of zero.
 *
 * Note:
 *
 * The solution set must not contain duplicate triplets.
 *
 * Example:
 *
 *
 * Given array nums = [-1, 0, 1, 2, -1, -4],
 *
 * A solution set is:
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 *
 *
 */

// 测试例子： [0,0,0], [0,0,0,0], [-1,0,1,0]
// @lc code=start
func threeSum(nums []int) [][]int {
	return threeSum1(nums)
}

func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	var retArr = make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i-1] == nums[i] { // 剔除第一位重复情况
			continue
		}
		j := i + 1
		last := len(nums) - 1
		for j < last {
			val := nums[i] + nums[j] + nums[last]
			if val > 0 {
				last--
			} else if val < 0 {
				j++
			} else {
				retArr = append(retArr, []int{nums[i], nums[j], nums[last]})
				j++
				// 第二数位有多个值符合，遍历直到下一个不等于的值，类似(-3,1,1,2)
				for nums[j] == nums[j-1] && j < last {
					j++
				}
			}
		}
	}

	return retArr
}

func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	var retArr = make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}
		j := i + 1
		last := len(nums) - 1
		for j < last {
			val := nums[i] + nums[j] + nums[last]
			if val > 0 {
				last--
			} else if val < 0 {
				j++
			} else {
				if j > i+1 && nums[j] == nums[j-1] {
					j++
					continue
				}
				retArr = append(retArr, []int{nums[i], nums[j], nums[last]})
				j++
			}
		}
	}

	return retArr
}

// @lc code=end
