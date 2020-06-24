/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 *
 * https://leetcode.com/problems/remove-element/description/
 *
 * algorithms
 * Easy (47.85%)
 * Likes:    1405
 * Dislikes: 2654
 * Total Accepted:    608.8K
 * Total Submissions: 1.3M
 * Testcase Example:  '[3,2,2,3]\n3'
 *
 * Given an array nums and a value val, remove all instances of that value
 * in-place and return the new length.
 *
 * Do not allocate extra space for another array, you must do this by modifying
 * the input array in-place with O(1) extra memory.
 *
 * The order of elements can be changed. It doesn't matter what you leave
 * beyond the new length.
 *
 * Example 1:
 *
 *
 * Given nums = [3,2,2,3], val = 3,
 *
 * Your function should return length = 2, with the first two elements of nums
 * being 2.
 *
 * It doesn't matter what you leave beyond the returned length.
 *
 *
 * Example 2:
 *
 *
 * Given nums = [0,1,2,2,3,0,4,2], val = 2,
 *
 * Your function should return length = 5, with the first five elements of nums
 * containing 0, 1, 3, 0, and 4.
 *
 * Note that the order of those five elements can be arbitrary.
 *
 * It doesn't matter what values are set beyond the returned length.
 *
 * Clarification:
 *
 * Confused why the returned value is an integer but your answer is an array?
 *
 * Note that the input array is passed in by reference, which means
 * modification to the input array will be known to the caller as well.
 *
 * Internally you can think of this:
 *
 *
 * // nums is passed in by reference. (i.e., without making a copy)
 * int len = removeElement(nums, val);
 *
 * // any modification to nums in your function would be known by the caller.
 * // using the length returned by your function, it prints the first len
 * elements.
 * for (int i = 0; i < len; i++) {
 * print(nums[i]);
 * }
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	return removeElement2(nums, val)
}

// 利用错位进行简化代码
func removeElement2(nums []int, val int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	nums = nums[0:index]
	return index
}

// 各种判断，逻辑有点太复杂，不是最简便方法
func removeElement1(nums []int, val int) int {

	count := 0
	total := len(nums)
	for i := 0; i < total-count; i++ {
		for total-1-count > i && nums[total-1-count] == val {
			count++
		}

		if count == total {
			break
		}

		if nums[i] == val {
			if i != total-1-count {
				nums[i] = nums[total-1-count]
			}
			count++
		}
	}

	if count == total {
		nums = make([]int, 0, 0)
	} else {
		nums = nums[:total-count]
	}
	return len(nums)
}

// @lc code=end
