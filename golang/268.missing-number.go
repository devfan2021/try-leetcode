/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 *
 * https://leetcode.com/problems/missing-number/description/
 *
 * algorithms
 * Easy (51.28%)
 * Likes:    1783
 * Dislikes: 2102
 * Total Accepted:    461.3K
 * Total Submissions: 894.9K
 * Testcase Example:  '[3,0,1]'
 *
 * Given an array containing n distinct numbers taken from 0, 1, 2, ..., n,
 * find the one that is missing from the array.
 *
 * Example 1:
 *
 *
 * Input: [3,0,1]
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: [9,6,4,2,3,5,7,0,1]
 * Output: 8
 *
 *
 * Note:
 * Your algorithm should run in linear runtime complexity. Could you implement
 * it using only constant extra space complexity?
 */

// @lc code=start
func missingNumber(nums []int) int {
	return missingNumber3(nums)
}

// use bit, time complexity: O(n), space complexity: O(1)
func missingNumber3(nums []int) int {
	missing := len(nums)
	for i := 0; i < len(nums); i++ {
		missing ^= i ^ nums[i]
	}
	return missing
}

// use math method, time complexity: O(n), space complexity: O(1)
func missingNumber2(nums []int) int {
	n := len(nums)
	total := n * (n + 1) / 2
	sum := 0

	for _, v := range nums {
		sum += v
	}
	return total - sum
}

// use hash, time complexity: O(n), space complexity: O(n)
func missingNumber1(nums []int) int {
	hash := map[int]bool{}
	for _, v := range nums {
		hash[v] = true
	}

	for i := 0; i <= len(nums); i++ {
		if _, ok := hash[i]; !ok {
			return i
		}
	}
	return -1
}

// @lc code=end