/*
 * @lc app=leetcode id=448 lang=golang
 *
 * [448] Find All Numbers Disappeared in an Array
 *
 * https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/description/
 *
 * algorithms
 * Easy (55.69%)
 * Likes:    2756
 * Dislikes: 240
 * Total Accepted:    252.1K
 * Total Submissions: 452.7K
 * Testcase Example:  '[4,3,2,7,8,2,3,1]'
 *
 * Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some
 * elements appear twice and others appear once.
 *
 * Find all the elements of [1, n] inclusive that do not appear in this array.
 *
 * Could you do it without extra space and in O(n) runtime? You may assume the
 * returned list does not count as extra space.
 *
 * Example:
 *
 * Input:
 * [4,3,2,7,8,2,3,1]
 *
 * Output:
 * [5,6]
 *
 *
 */

// @lc code=start
// 官方其他解法感觉有点复杂，时间效率？？
func findDisappearedNumbers(nums []int) []int {
	return findDisappearedNumbers1(nums)
}

// 思路，构造一个比输入数组大1的标记整型数组，默认值为0， 遍历输入数组，
// 在对应位置上标记为1，其他位置上为0标示缺失该数字
func findDisappearedNumbers1(nums []int) []int {
	retVal := make([]int, len(nums)+1, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		retVal[nums[i]] = 1
	}
	j := 0
	for i := 1; i < len(retVal); i++ {
		if retVal[i] == 0 {
			retVal[j] = i
			j++
		}
	}
	return retVal[0:j]
}

// @lc code=end

