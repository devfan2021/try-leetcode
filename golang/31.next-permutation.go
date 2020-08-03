import "sort"

/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 *
 * https://leetcode.com/problems/next-permutation/description/
 *
 * algorithms
 * Medium (32.31%)
 * Likes:    3529
 * Dislikes: 1272
 * Total Accepted:    381.2K
 * Total Submissions: 1.2M
 * Testcase Example:  '[1,2,3]'
 *
 * Implement next permutation, which rearranges numbers into the
 * lexicographically next greater permutation of numbers.
 *
 * If such arrangement is not possible, it must rearrange it as the lowest
 * possible order (ie, sorted in ascending order).
 *
 * The replacement must be in-place and use only constant extra memory.
 *
 * Here are some examples. Inputs are in the left-hand column and its
 * corresponding outputs are in the right-hand column.
 *
 * 1,2,3 → 1,3,2
 * 3,2,1 → 1,2,3
 * 1,1,5 → 1,5,1
 *
 */
// test case
// tmp := []int{1, 2, 3}
// tmp := []int{3, 2, 1}
// tmp := []int{1, 1, 5}
// tmp := []int{2, 3, 1} // [3,1,2]
// tmp := []int{1, 3, 2} // [2,1,3]
// tmp := []int{1, 5, 1} // [5,1,1]

// @lc code=start
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	// find first left less right value index
	exist, cmpIndex := false, len(nums)-1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[cmpIndex] {
			exist = true
			cmpIndex = i
			break
		} else {
			cmpIndex = i
		}
	}

	if exist {
		swapIndex := -1
		for i := cmpIndex + 1; i <= len(nums)-2; i++ {
			if nums[i] > nums[cmpIndex] && nums[cmpIndex] >= nums[i+1] {
				swapIndex = i
				break
			}
		}
		if swapIndex == -1 {
			swapIndex = len(nums) - 1
		}

		// swap two valus
		nums[cmpIndex], nums[swapIndex] = nums[swapIndex], nums[cmpIndex]

		// revrese last next elements
		for i := cmpIndex + 1; i < len(nums); i++ {
			tmp := nums[len(nums)-1]
			for j := len(nums) - 1; j > i; j-- {
				nums[j] = nums[j-1]
			}
			nums[i] = tmp
		}
	} else {
		sort.Ints(nums)
	}
}

// @lc code=end