/*
 * @lc app=leetcode id=1346 lang=golang
 *
 * [1346] Check If N and Its Double Exist
 *
 * https://leetcode.com/problems/check-if-n-and-its-double-exist/description/
 *
 * algorithms
 * Easy (38.44%)
 * Likes:    150
 * Dislikes: 23
 * Total Accepted:    34.6K
 * Total Submissions: 90.3K
 * Testcase Example:  '[10,2,5,3]'
 *
 * Given an array arr of integers, check if there exists two integers N and M
 * such that N is the double of M ( i.e. N = 2 * M).
 *
 * More formally check if there exists two indices i and j such that :
 *
 *
 * i != j
 * 0 <= i, j < arr.length
 * arr[i] == 2 * arr[j]
 *
 *
 *
 * Example 1:
 *
 *
 * Input: arr = [10,2,5,3]
 * Output: true
 * Explanation: N = 10 is the double of M = 5,that is, 10 = 2 * 5.
 *
 *
 * Example 2:
 *
 *
 * Input: arr = [7,1,14,11]
 * Output: true
 * Explanation: N = 14 is the double of M = 7,that is, 14 = 2 * 7.
 *
 *
 * Example 3:
 *
 *
 * Input: arr = [3,1,7,11]
 * Output: false
 * Explanation: In this case does not exist N and M, such that N = 2 * M.
 *
 *
 *
 * Constraints:
 *
 *
 * 2 <= arr.length <= 500
 * -10^3 <= arr[i] <= 10^3
 *
 *
 */

// @lc code=start
func checkIfExist(arr []int) bool {
	return checkIfExist2(arr)
}

// 空间换时间的策略，增加一个缓存Map，减少一层for循环
func checkIfExist2(arr []int) bool {
	cacheMap := make(map[int]bool)
	for _, val := range arr {
		_, halfValFlag := cacheMap[val/2]
		_, doubleValFlag := cacheMap[val*2]
		// 判断是否除尽的情况
		if doubleValFlag || (halfValFlag && val%2 == 0) {
			return true
		}
		cacheMap[val] = true
	}
	return false
}

func checkIfExist1(arr []int) bool {
	if len(arr) <= 1 {
		return false
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j]*2 || arr[i]*2 == arr[j] {
				return true
			}
		}
	}
	return false
}

// @lc code=end

