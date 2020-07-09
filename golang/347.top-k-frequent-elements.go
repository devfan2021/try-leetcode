import (
	"container/heap"
	"sort"
)

/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 *
 * https://leetcode.com/problems/top-k-frequent-elements/description/
 *
 * algorithms
 * Medium (60.22%)
 * Likes:    3066
 * Dislikes: 212
 * Total Accepted:    389.6K
 * Total Submissions: 645.3K
 * Testcase Example:  '[1,1,1,2,2,3]\n2'
 *
 * Given a non-empty array of integers, return the k most frequent elements.
 *
 * Example 1:
 *
 *
 * Input: nums = [1,1,1,2,2,3], k = 2
 * Output: [1,2]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1], k = 1
 * Output: [1]
 *
 *
 * Note:
 *
 *
 * You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
 * Your algorithm's time complexity must be better than O(n log n), where n is
 * the array's size.
 * It's guaranteed that the answer is unique, in other words the set of the top
 * k frequent elements is unique.
 * You can return the answer in any order.
 *
 *
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	return topKFrequent3(nums, k)
}

type KNode struct {
	key   int
	value int
}

type IntHeap []KNode

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].value > h[j].value }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(KNode))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 采用最小堆的方式来实现
func topKFrequent3(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	if k <= 0 || k > len(nums) {
		return nil
	}

	hash := map[int]int{}
	for i := 0; i < len(nums); i++ {
		hash[nums[i]]++
	}

	hp := &IntHeap{}
	heap.Init(hp)

	for k, v := range hash {
		heap.Push(hp, KNode{key: k, value: v})
	}

	ret := []int{}
	for i := 0; i < k && hp.Len() > 0; i++ {
		ret = append(ret, heap.Pop(hp).(KNode).key)
	}
	return ret
}

// 先hash,再巧妙的用map构造二位数组形式，再进行自定义排序
func topKFrequent2(nums []int, k int) []int {
	if k >= len(nums) {
		return nums
	}
	hash := map[int]int{}
	for i := 0; i < len(nums); i++ {
		hash[nums[i]]++
	}

	tmp := [][]int{}
	for k, v := range hash {
		tmp = append(tmp, []int{k, v})
	}
	sort.Slice(tmp, func(a, b int) bool {
		return tmp[a][1] > tmp[b][1]
	})

	ret := []int{}
	for i := 0; i < k; i++ {
		ret = append(ret, tmp[i][0])
	}
	return ret
}

func topKFrequent1(nums []int, k int) []int {
	if k >= len(nums) {
		return nums
	}
	hash := map[int]int{}
	for i := 0; i < len(nums); i++ {
		hash[nums[i]]++
	}
	// 变成hash排序了(最小堆，快排)
	retVal := make([]int, k)
	for k, _ := range hash {
		retVal = append(retVal, k)
	}
	return retVal
}

// @lc code=end

