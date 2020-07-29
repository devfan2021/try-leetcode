/*
 * @lc app=leetcode id=1365 lang=golang
 *
 * [1365] How Many Numbers Are Smaller Than the Current Number
 *
 * https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/description/
 *
 * algorithms
 * Easy (85.18%)
 * Likes:    725
 * Dislikes: 23
 * Total Accepted:    92K
 * Total Submissions: 107.7K
 * Testcase Example:  '[8,1,2,2,3]'
 *
 * Given the array nums, for each nums[i] find out how many numbers in the
 * array are smaller than it. That is, for each nums[i] you have to count the
 * number of valid j's such that j != i and nums[j] < nums[i].
 *
 * Return the answer in an array.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [8,1,2,2,3]
 * Output: [4,0,1,1,3]
 * Explanation:
 * For nums[0]=8 there exist four smaller numbers than it (1, 2, 2 and 3).
 * For nums[1]=1 does not exist any smaller number than it.
 * For nums[2]=2 there exist one smaller number than it (1).
 * For nums[3]=2 there exist one smaller number than it (1).
 * For nums[4]=3 there exist three smaller numbers than it (1, 2 and 2).
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [6,5,4,8]
 * Output: [2,1,0,3]
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [7,7,7,7]
 * Output: [0,0,0,0]
 *
 *
 *
 * Constraints:
 *
 *
 * 2 <= nums.length <= 500
 * 0 <= nums[i] <= 100
 *
 */

// @lc code=start
func smallerNumbersThanCurrent(nums []int) []int {
	return smallerNumbersThanCurrent2(nums)
}

// using range
func smallerNumbersThanCurrent2(nums []int) []int {
	result := make([]int, 0)
	for _, i := range nums {
		count := 0
		for _, j := range nums {
			if i > j {
				count++
			}
		}
		result = append(result, count)
	}
	return result
}

// using for
func smallerNumbersThanCurrent1(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if nums[i] > nums[j] {
				count++
			}
		}
		result[i] = count
	}
	return result
}

// @lc code=end