/*
 * @lc app=leetcode id=908 lang=golang
 *
 * [908] Smallest Range I
 *
 * https://leetcode.com/problems/smallest-range-i/description/
 *
 * algorithms
 * Easy (66.92%)
 * Likes:    381
 * Dislikes: 1521
 * Total Accepted:    56.3K
 * Total Submissions: 84.1K
 * Testcase Example:  '[1]\n0'
 *
 * You are given an integer array nums and an integer k.
 * 
 * In one operation, you can choose any index i where 0 <= i < nums.length and
 * change nums[i] to nums[i] + x where x is an integer from the range [-k, k].
 * You can apply this operation at most once for each index i.
 * 
 * The score of nums is the difference between the maximum and minimum elements
 * in nums.
 * 
 * Return the minimum score of nums after applying the mentioned operation at
 * most once for each index in it.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [1], k = 0
 * Output: 0
 * Explanation: The score is max(nums) - min(nums) = 1 - 1 = 0.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [0,10], k = 2
 * Output: 6
 * Explanation: Change nums to be [2, 8]. The score is max(nums) - min(nums) =
 * 8 - 2 = 6.
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: nums = [1,3,6], k = 3
 * Output: 0
 * Explanation: Change nums to be [4, 4, 4]. The score is max(nums) - min(nums)
 * = 4 - 4 = 0.
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= nums.length <= 10^4
 * 0 <= nums[i] <= 10^4
 * 0 <= k <= 10^4
 * 
 * 
 */

// @lc code=start
func smallestRangeI(nums []int, k int) int {
	return soluntion1(nums, k)
}

// simllar to optimum solution
func soluntion1(nums []int, k int) int {
  // if len(nums) == 0 {
  //   return 0
  // }

  minVal, maxVal := nums[0], nums[0]

  for _, val := range nums {
    if val < minVal {
      minVal = val
    }
    if val > maxVal {
      maxVal = val
    }
  }

  val := maxVal - minVal - 2 * k
  if val < 0 {
    return 0
  }
  return val
}

// optimum solution
func soluntion2(nums[] int, k int) int {
  min, max := nums[0], nums[0]
  for _, x := range nums {
    if x < min {
      min = x
    }
    if x > max {
      max = x
    }
  }

  min += k
  max -= k
  score := max - min
  if score < 0 {
    return 0
  }
  return score
}

// @lc code=end