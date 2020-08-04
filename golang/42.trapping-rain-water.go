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
	if len(height) <= 1 {
		return 0
	}

	sum := 0
	cmpVal := height[0]
	for i := 0; i < len(height); {
		calculate := false
		if height[i] > cmpVal {
			cmpVal = height[i]
		}
		if cmpVal > 0 {
			calculate = true
		}

		if calculate {
			nextI := i
			for j := i + 1; j < len(height); j++ {
				if height[j] >= cmpVal {
					nextI = j
					break
				}
			}
			if nextI == i {
				i++
				if i >= len(height) {
					break
				}
				cmpVal = height[i]
				continue
			}

			minVal := cmpVal
			if height[nextI] < minVal {
				minVal = height[nextI]
			}

			for k := i; k <= nextI; k++ {
				minus := minVal - height[k]
				if minus > 0 {
					sum += minus
				}
			}
			// fmt.Println("=======")
			// fmt.Println(sum)

			cmpVal = height[nextI]
			i = nextI
		} else {
			i++
		}
	}
	return sum
}

// @lc code=end