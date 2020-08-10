/*
 * @lc app=leetcode id=724 lang=golang
 *
 * [724] Find Pivot Index
 *
 * https://leetcode.com/problems/find-pivot-index/description/
 *
 * algorithms
 * Easy (43.66%)
 * Likes:    1063
 * Dislikes: 240
 * Total Accepted:    130.7K
 * Total Submissions: 299.2K
 * Testcase Example:  '[1,7,3,6,5,6]'
 *
 * Given an array of integers nums, write a method that returns the "pivot"
 * index of this array.
 *
 * We define the pivot index as the index where the sum of all the numbers to
 * the left of the index is equal to the sum of all the numbers to the right of
 * the index.
 *
 * If no such index exists, we should return -1. If there are multiple pivot
 * indexes, you should return the left-most pivot index.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,7,3,6,5,6]
 * Output: 3
 * Explanation:
 * The sum of the numbers to the left of index 3 (nums[3] = 6) is equal to the
 * sum of numbers to the right of index 3.
 * Also, 3 is the first index where this occurs.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1,2,3]
 * Output: -1
 * Explanation:
 * There is no index that satisfies the conditions in the problem
 * statement.
 *
 *
 *
 * Constraints:
 *
 *
 * The length of nums will be in the range [0, 10000].
 * Each element nums[i] will be an integer in the range [-1000, 1000].
 *
 *
 */

// @lc code=start
func pivotIndex(nums []int) int {
	return pivotIndex3(nums)
}

// solution by math sum, leftSum: leftSum = sum - leftSum - nums[i]
func pivotIndex3(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	leftSum := 0
	for i := 0; i < len(nums); i++ {
		if leftSum == sum-leftSum-nums[i] {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}

func pivotIndex2(nums []int) int {
	if len(nums) <= 2 {
		return -1
	}
	sum := 0
	for _, val := range nums {
		sum += val
	}

	// handle special case:  total sum minus first element, value is zero
	// test case:  [-1,-1,-1,0,1,1]
	if sum-nums[0] == 0 {
		return 0
	}

	leftSum := 0
	for i := 0; i+2 < len(nums); i++ {
		leftSum += nums[i]
		midVal := nums[i+1]
		if leftSum == (sum-midVal)/2 && (sum-midVal)%2 == 0 {
			return i + 1
		}
	}

	// handle special case:  total sum minus last element, value is zero. this check step should put last, may be some other solution
	// test case:  [-1,-1,1,1,0,0]
	if sum-nums[len(nums)-1] == 0 {
		return len(nums) - 1
	}
	return -1
}

// 分别从两端开始遍历累加，这个方法有问题，如果数组中有负值，就有问题
func pivotIndex1(nums []int) int {
	if len(nums) <= 2 {
		return -1
	}

	leftSum := nums[0]
	rightSum := nums[len(nums)-1]
	for i, j := 0, len(nums)-1; i <= j; {
		if leftSum == rightSum {
			if i+2 == j {
				return i + 1
			} else {
				i++
				j--
				leftSum += nums[i]
				rightSum += nums[j]
			}
		} else if leftSum > rightSum {
			j--
			rightSum += nums[j]
		} else {
			i++
			leftSum += nums[i]
		}
	}
	return -1
}

// [-1,-1,-1,-1,-1,0]
// [-1,-1,-1,-1,-1,-1]
// [-1,-1,-1,0,1,1] --- 开头，结尾特殊案例
// [-1,-1,1,1,0,0] --- 去除最后一个为零，但是中间有其他方案，优先返回其他方案
// @lc code=end