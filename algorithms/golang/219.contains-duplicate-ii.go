/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
 *
 * https://leetcode.com/problems/contains-duplicate-ii/description/
 *
 * algorithms
 * Easy (37.42%)
 * Likes:    837
 * Dislikes: 1003
 * Total Accepted:    268.9K
 * Total Submissions: 717K
 * Testcase Example:  '[1,2,3,1]\n3'
 *
 * Given an array of integers and an integer k, find out whether there are two
 * distinct indices i and j in the array such that nums[i] = nums[j] and the
 * absolute difference between i and j is at most k.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,2,3,1], k = 3
 * Output: true
 *
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1,0,1,1], k = 1
 * Output: true
 *
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [1,2,3,1,2,3], k = 2
 * Output: false
 *
 *
 *
 *
 *
 */
// [99,99]\n2
// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	hash := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		v := nums[i]
		if j, ok := hash[v]; ok {
			if Abs(i-j) <= k {
				return true
			}
		}
		hash[v] = i
	}
	return false
}

func Abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

// @lc code=end