import "sort"

/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 *
 * https://leetcode.com/problems/4sum/description/
 *
 * algorithms
 * Medium (33.28%)
 * Likes:    1892
 * Dislikes: 321
 * Total Accepted:    328.8K
 * Total Submissions: 985.2K
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * Given an array nums of n integers and an integer target, are there elements
 * a, b, c, and d in nums such that a + b + c + d = target? Find all unique
 * quadruplets in the array which gives the sum of target.
 *
 * Note:
 *
 * The solution set must not contain duplicate quadruplets.
 *
 * Example:
 *
 *
 * Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.
 *
 * A solution set is:
 * [
 * ⁠ [-1,  0, 0, 1],
 * ⁠ [-2, -1, 1, 2],
 * ⁠ [-2,  0, 0, 2]
 * ]
 *
 *
 */
// [-3,-2,-1,0,0,1,2,3] ==> [[-3,-2,2,3],[-3,-1,1,3],[-3,0,0,3],[-3,0,1,2],[-2,-1,0,3],[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
//                          [[-3,-2,2,3],[-3,-1,1,3],[-3,0,0,3],           [-2,-1,0,3],            [-2,0,0,2],[-1,0,0,1]]
// @lc code=start
func fourSum(nums []int, target int) [][]int {
	return fourSum1(nums, target)
}

// 优化点，某一位上重复值的处理
func fourSum1(nums []int, target int) [][]int {
	retVal := make([][]int, 0)
	sort.Ints(nums)
	for one := 0; one < len(nums)-3; one++ {
		if one != 0 && nums[one] == nums[one-1] {
			continue
		}
		for second := one + 1; second < len(nums)-2; second++ {
			if second > one+1 && nums[second] == nums[second-1] {
				continue
			}
			val := target - nums[one] - nums[second]
			third := second + 1
			four := len(nums) - 1
			for third < four {
				if nums[third]+nums[four] > val {
					four--
				} else if nums[third]+nums[four] < val {
					third++
				} else {
					retVal = append(retVal, []int{nums[one], nums[second], nums[third], nums[four]})
					third++
					for third < four && nums[third] == nums[third-1] {
						third++
					}
					// 加快第三层迭代
					for third < four && nums[four] == nums[four-1] {
						four--
					}
				}
			}
		}
	}
	return retVal
}

func fourSum2(nums []int, target int) [][]int {
	retVal := make([][]int, 0)
	sort.Ints(nums)
	for one := 0; one < len(nums)-3; one++ {
		if one != 0 && nums[one] == nums[one-1] {
			continue
		}
		for second := one + 1; second < len(nums)-2; second++ {
			if second > one+1 && nums[second] == nums[second-1] {
				continue
			}
			val := target - nums[one] - nums[second]
			third := second + 1
			four := len(nums) - 1
			for third < four {
				if nums[third]+nums[four] > val {
					four--
				} else if nums[third]+nums[four] < val {
					third++
				} else {
					retVal = append(retVal, []int{nums[one], nums[second], nums[third], nums[four]})
					third++
					for nums[third] == nums[third-1] && third < four {
						third++
					}
				}
			}
		}
	}
	return retVal
}

// @lc code=end

