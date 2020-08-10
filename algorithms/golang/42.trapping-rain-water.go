/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 *
 * https://leetcode.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (48.22%)
 * Likes:    7426
 * Dislikes: 128
 * Total Accepted:    530.5K
 * Total Submissions: 1.1M
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * Given n non-negative integers representing an elevation map where the width
 * of each bar is 1, compute how much water it is able to trap after raining.
 *
 *
 * The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1].
 * In this case, 6 units of rain water (blue section) are being trapped. Thanks
 * Marcos for contributing this image!
 *
 * Example:
 *
 *
 * Input: [0,1,0,2,1,0,1,3,2,1,2,1]
 * Output: 6
 *
 */

// @lc code=start
// [4,2,3]
func trap(height []int) int {
	return trap3(height)
}

// using two pointer, time complexity: O(n), space complexity: O(n)
func trap3(height []int) int {
	left, right, count := 0, len(height)-1, 0
	leftMax, rightMax := 0, 0
	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				count += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				count += rightMax - height[right]
			}
			right--
		}
	}
	return count
}

// this method similary brute method, enhance time complexity, using extra space to store max for each element
// time complexity: O(n), space complexity: O(n)
func trap2(height []int) int {
	if len(height) <= 1 {
		return 0
	}

	leftMaxContainer := make([]int, len(height))
	leftMaxVal := 0
	for i := 0; i < len(height); i++ {
		leftMaxVal = max(leftMaxVal, height[i])
		leftMaxContainer[i] = leftMaxVal
	}

	rightMaxContainer := make([]int, len(height))
	rightMaxVal := 0
	for i := len(height) - 1; i >= 0; i-- {
		rightMaxVal = max(rightMaxVal, height[i])
		rightMaxContainer[i] = rightMaxVal
	}

	count := 0
	for i := 1; i < len(height)-1; i++ {
		leftMax, rightMax := leftMaxContainer[i], rightMaxContainer[i]
		count += min(leftMax, rightMax) - height[i]
	}
	return count
}

// brute method, time complexity: O(n*n), space complexity: O(1)
func trap1(height []int) int {
	count, size := 0, len(height)
	// start at index one and end at len(height) -2, not include first and last element
	for i := 1; i < size-1; i++ {
		maxLeft, maxRight := 0, 0
		for j := i; j >= 0; j-- {
			// search the left part for max bar size
			maxLeft = max(maxLeft, height[j])
		}

		for j := i; j < size; j++ {
			// search the right part for max bar size
			maxRight = max(maxRight, height[j])
		}

		count += min(maxLeft, maxRight) - height[i]
	}
	return count
}

func max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

func min(val1, val2 int) int {
	if val1 > val2 {
		return val2
	}
	return val1
}

// @lc code=end