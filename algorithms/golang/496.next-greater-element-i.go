/*
 * @lc app=leetcode id=496 lang=golang
 *
 * [496] Next Greater Element I
 *
 * https://leetcode.com/problems/next-greater-element-i/description/
 *
 * algorithms
 * Easy (63.30%)
 * Likes:    1662
 * Dislikes: 2211
 * Total Accepted:    163.7K
 * Total Submissions: 256.4K
 * Testcase Example:  '[4,1,2]\n[1,3,4,2]'
 *
 *
 * You are given two arrays (without duplicates) nums1 and nums2 where nums1’s
 * elements are subset of nums2. Find all the next greater numbers for nums1's
 * elements in the corresponding places of nums2.
 *
 *
 *
 * The Next Greater Number of a number x in nums1 is the first greater number
 * to its right in nums2. If it does not exist, output -1 for this number.
 *
 *
 * Example 1:
 *
 * Input: nums1 = [4,1,2], nums2 = [1,3,4,2].
 * Output: [-1,3,-1]
 * Explanation:
 * ⁠   For number 4 in the first array, you cannot find the next greater number
 * for it in the second array, so output -1.
 * ⁠   For number 1 in the first array, the next greater number for it in the
 * second array is 3.
 * ⁠   For number 2 in the first array, there is no next greater number for it
 * in the second array, so output -1.
 *
 *
 *
 * Example 2:
 *
 * Input: nums1 = [2,4], nums2 = [1,2,3,4].
 * Output: [3,-1]
 * Explanation:
 * ⁠   For number 2 in the first array, the next greater number for it in the
 * second array is 3.
 * ⁠   For number 4 in the first array, there is no next greater number for it
 * in the second array, so output -1.
 *
 *
 *
 *
 * Note:
 *
 * All elements in nums1 and nums2 are unique.
 * The length of both nums1 and nums2 would not exceed 1000.
 *
 *
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	return nextGreaterElement3(nums1, nums2)
}

// using stack and map
func nextGreaterElement3(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	stack, hash := []int{}, map[int]int{}
	for i := 0; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
			hash[stack[len(stack)-1]] = nums2[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}

	for len(stack) > 0 {
		hash[stack[len(stack)-1]] = -1
		stack = stack[:len(stack)-1]
	}

	retVals := make([]int, len(nums1))
	for i := range nums1 {
		retVals[i] = hash[nums1[i]]
	}
	return retVals
}

// similary solution1, concise
func nextGreaterElement2(nums1 []int, nums2 []int) []int {
	hash := map[int]int{}
	for idx, v := range nums2 {
		hash[v] = idx
	}

	retVals := make([]int, len(nums1))
	for i, n := range nums1 {
		idx := hash[n]

		found := false
		for j := idx + 1; j < len(nums2); j++ {
			if nums2[j] > n {
				found = true
				retVals[i] = nums2[j]
				break
			}
		}

		if !found {
			retVals[i] = -1
		}
	}
	return retVals
}

func nextGreaterElement1(nums1 []int, nums2 []int) []int {
	hash := map[int]int{}
	for i, v := range nums2 {
		hash[v] = i
	}

	retVals := make([]int, len(nums1))
	for i, v := range nums1 {
		if index, ok := hash[v]; ok {
			if index == len(nums2)-1 { // last element
				retVals[i] = -1
			} else {
				for j := index + 1; j < len(nums2); j++ {
					if nums2[j] > v {
						retVals[i] = nums2[j]
						break
					}
					if j == len(nums2)-1 { // last element
						retVals[i] = -1
					}
				}
			}
		} else {
			retVals[i] = -1
		}
	}
	return retVals
}

// @lc code=end