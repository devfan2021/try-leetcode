import "sort"

/*
 * @lc app=leetcode id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 *
 * https://leetcode.com/problems/intersection-of-two-arrays-ii/description/
 *
 * algorithms
 * Easy (51.10%)
 * Likes:    1360
 * Dislikes: 404
 * Total Accepted:    352.3K
 * Total Submissions: 688.4K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * Given two arrays, write a function to compute their intersection.
 *
 * Example 1:
 *
 *
 * Input: nums1 = [1,2,2,1], nums2 = [2,2]
 * Output: [2,2]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * Output: [4,9]
 *
 *
 * Note:
 *
 *
 * Each element in the result should appear as many times as it shows in both
 * arrays.
 * The result can be in any order.
 *
 *
 * Follow up:
 *
 *
 * What if the given array is already sorted? How would you optimize your
 * algorithm?
 * What if nums1's size is small compared to nums2's size? Which algorithm is
 * better?
 * What if elements of nums2 are stored on disk, and the memory is limited such
 * that you cannot load all elements into the memory at once?
 *
 *
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	retValues := make([]int, 0)

	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		val1, val2 := nums1[i], nums2[j]
		if val1 == val2 {

			// 计算nums1连续相等的值
			tmpI := 0
			for i+1 < len(nums1) && nums1[i] == nums1[i+1] {
				i++
				tmpI++
			}

			// 计算nums2连续相等的值
			tmpJ := 0
			for j+1 < len(nums2) && nums2[j] == nums2[j+1] {
				j++
				tmpJ++
			}

			// 计算相同的连续相等值
			count := 1
			if tmpI > 0 && tmpJ > 0 {
				if tmpI > tmpJ {
					count += tmpJ
				} else {
					count += tmpI
				}
			}
			for ; count > 0; count-- {
				retValues = append(retValues, val1)
			}
			i++
			j++
		} else if val1 > val2 {
			j++
		} else {
			i++
		}
	}
	return retValues
}

// @lc code=end