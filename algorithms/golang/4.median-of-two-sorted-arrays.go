import "sort"

/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 *
 * https://leetcode.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (29.21%)
 * Likes:    7124
 * Dislikes: 1106
 * Total Accepted:    687K
 * Total Submissions: 2.3M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * There are two sorted arrays nums1 and nums2 of size m and n respectively.
 *
 * Find the median of the two sorted arrays. The overall run time complexity
 * should be O(log (m+n)).
 *
 * You may assume nums1 and nums2 cannot be both empty.
 *
 * Example 1:
 *
 *
 * nums1 = [1, 3]
 * nums2 = [2]
 *
 * The median is 2.0
 *
 *
 * Example 2:
 *
 *
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 *
 * The median is (2 + 3)/2 = 2.5
 *
 *
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return findMedianSortedArrays2(nums1, nums2)
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	lCombined := append(nums1, nums2...)
	sort.Ints(lCombined)
	n := len(lCombined)
	if n%2 == 0 {
		return (float64(lCombined[n/2] + lCombined[(n/2)-1])) / 2
	} else {
		return float64(lCombined[(n-1)/2])
	}
}

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	values := []int{}
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	} else if len(nums1) == 0 {
		values = append(values, nums2...)
	} else if len(nums2) == 0 {
		values = append(values, nums1...)
	} else {
		i, j := 0, 0
		for i < len(nums1) && j < len(nums2) {
			item1, item2 := nums1[i], nums2[j]
			if item1 < item2 {
				values = append(values, item1)
				i++
			} else {
				values = append(values, item2)
				j++
			}
		}

		for i < len(nums1) {
			values = append(values, nums1[i])
			i++
		}

		for j < len(nums2) {
			values = append(values, nums2[j])
			j++
		}
	}
	mid := len(values) / 2
	if len(values)%2 != 0 {
		return float64(values[mid])
	} else {
		return float64((values[mid-1] + values[mid])) / 2
	}
}

// @lc code=end