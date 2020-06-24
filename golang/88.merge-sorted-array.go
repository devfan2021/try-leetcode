/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 *
 * https://leetcode.com/problems/merge-sorted-array/description/
 *
 * algorithms
 * Easy (39.07%)
 * Likes:    2145
 * Dislikes: 4082
 * Total Accepted:    582.3K
 * Total Submissions: 1.5M
 * Testcase Example:  '[1,2,3,0,0,0]\n3\n[2,5,6]\n3'
 *
 * Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as
 * one sorted array.
 *
 * Note:
 *
 *
 * The number of elements initialized in nums1 and nums2 are m and n
 * respectively.
 * You may assume that nums1 has enough space (size that is greater or equal to
 * m + n) to hold additional elements from nums2.
 *
 *
 * Example:
 *
 *
 * Input:
 * nums1 = [1,2,3,0,0,0], m = 3
 * nums2 = [2,5,6],       n = 3
 *
 * Output: [1,2,2,3,5,6]
 *
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	merge2(nums1, m, nums2, n)
}

// 第一版失败案例, 缺少对临界点的判断处理, nums1为空或者只有一个小值
// [2,0]\n1\n[1]\n1
// [0]\n0\n[1]\n1
func merge2(nums1 []int, m int, nums2 []int, n int) {
	cmpIndex := m - 1
	count := 0
	for i := n - 1; i >= 0; {
		lastIndex := len(nums1) - count - 1
		nVal := nums2[i]
		if m == 0 || cmpIndex < 0 {
			nums1[lastIndex] = nVal
			i--
		} else {
			mVal := nums1[cmpIndex]
			// 大于，直接放后面
			if mVal < nVal {
				nums1[lastIndex] = nVal
				i--
			} else {
				nums1[lastIndex] = mVal
				nums1[cmpIndex] = 0
				cmpIndex--
			}
		}
		count++
	}
}

// 额外加个数组进行处理
func merge1(nums1 []int, m int, nums2 []int, n int) {
	sortArr := make([]int, len(nums1), len(nums1))
	count := len(nums1)

	for i, mIndex, nIndex := 0, 0, 0; i < count; i++ {
		if mIndex == m || nIndex == n {
			if mIndex == m {
				copy(sortArr[m+nIndex:], nums2[nIndex:])
			} else {
				copy(sortArr[n+mIndex:], nums1[mIndex:])
			}
		} else {
			mVal := nums1[mIndex]
			nVal := nums2[nIndex]
			if mVal < nVal {
				sortArr[i] = mVal
				mIndex++
			} else {
				sortArr[i] = nVal
				nIndex++
			}
		}
	}
	copy(nums1[0:], sortArr[0:])
}

// @lc code=end

