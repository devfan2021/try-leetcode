import "sort"

/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 *
 * https://leetcode.com/problems/intersection-of-two-arrays/description/
 *
 * algorithms
 * Easy (61.67%)
 * Likes:    819
 * Dislikes: 1213
 * Total Accepted:    361.9K
 * Total Submissions: 584.6K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * Given two arrays, write a function to compute their intersection.
 *
 * Example 1:
 *
 *
 * Input: nums1 = [1,2,2,1], nums2 = [2,2]
 * Output: [2]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * Output: [9,4]
 *
 *
 * Note:
 *
 *
 * Each element in the result must be unique.
 * The result can be in any order.
 *
 *
 *
 *
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	return intersection2(nums1, nums2)
}

// 先排序，再一一比对
func intersection2(nums1 []int, nums2 []int) []int {
	retVal := []int{}
	sort.Ints(nums1)
	sort.Ints(nums2)

	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		left, right := nums1[i], nums2[j]
		if left == right {
			retVal = append(retVal, left)
			i++
			// 剔除重复的
			for i < len(nums1) && nums1[i-1] == nums1[i] {
				i++
			}

			j++
			// 剔除重复的
			for j < len(nums2) && nums2[j-1] == nums2[j] {
				j++
			}
		} else if left < right {
			i++
		} else {
			j++
		}
	}
	return retVal
}

func intersection1(nums1 []int, nums2 []int) []int {
	valMap := make(map[int]bool)
	retVal := make([]int, 0)

	for _, val := range nums1 {
		valMap[val] = true
	}

	for _, val := range nums2 {
		if valMap[val] {
			retVal = append(retVal, val)
			valMap[val] = false
			// delete(valMap, val)
		}
	}
	return retVal
}

// @lc code=end
