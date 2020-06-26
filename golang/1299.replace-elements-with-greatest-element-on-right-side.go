/*
 * @lc app=leetcode id=1299 lang=golang
 *
 * [1299] Replace Elements with Greatest Element on Right Side
 *
 * https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side/description/
 *
 * algorithms
 * Easy (76.21%)
 * Likes:    327
 * Dislikes: 76
 * Total Accepted:    49.5K
 * Total Submissions: 65K
 * Testcase Example:  '[17,18,5,4,6,1]'
 *
 * Given an array arr, replace every element in that array with the greatest
 * element among the elements to its right, and replace the last element with
 * -1.
 *
 * After doing so, return the array.
 *
 *
 * Example 1:
 * Input: arr = [17,18,5,4,6,1]
 * Output: [18,6,6,6,1,-1]
 *
 *
 * Constraints:
 *
 *
 * 1 <= arr.length <= 10^4
 * 1 <= arr[i] <= 10^5
 *
 */

// @lc code=start
func replaceElements(arr []int) []int {
	return replaceElements1(arr)
}

// 原始写法错了，暂时看最高解决算法
// 巧妙从末尾进行迭代， 因为是找相邻右边最大的值，所以从右边变量比较直观
func replaceElements1(arr []int) []int {
	cmpVal := arr[len(arr)-1]
	arr[len(arr)-1] = -1
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] > cmpVal {
			tmp := arr[i]
			arr[i] = cmpVal
			cmpVal = tmp
		} else {
			arr[i] = cmpVal
		}
	}
	return arr
}

// @lc code=end
