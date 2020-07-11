/*
 * @lc app=leetcode id=705 lang=golang
 *
 * [705] Design HashSet
 *
 * https://leetcode.com/problems/design-hashset/description/
 *
 * algorithms
 * Easy (60.14%)
 * Likes:    324
 * Dislikes: 70
 * Total Accepted:    50.6K
 * Total Submissions: 84K
 * Testcase Example:  '["MyHashSet","add","add","contains","contains","add","contains","remove","contains"]\n' +
  '[[],[1],[2],[1],[3],[2],[2],[2],[2]]'
 *
 * Design a HashSet without using any built-in hash table libraries.
 *
 * To be specific, your design should include these functions:
 *
 *
 * add(value): Insert a value into the HashSet.
 * contains(value) : Return whether the value exists in the HashSet or not.
 * remove(value): Remove a value in the HashSet. If the value does not exist in
 * the HashSet, do nothing.
 *
 *
 *
 * Example:
 *
 *
 * MyHashSet hashSet = new MyHashSet();
 * hashSet.add(1);
 * hashSet.add(2);
 * hashSet.contains(1);    // returns true
 * hashSet.contains(3);    // returns false (not found)
 * hashSet.add(2);
 * hashSet.contains(2);    // returns true
 * hashSet.remove(2);
 * hashSet.contains(2);    // returns false (already removed)
 *
 *
 *
 * Note:
 *
 *
 * All values will be in the range of [0, 1000000].
 * The number of operations will be in the range of [1, 10000].
 * Please do not use the built-in HashSet library.
 *
 *
*/

// 1.使用最简单的方式，内部是一个[]int
// 2.设计排序好的[]int ==>如果数据量大，每次都要进行排序，耗时
// 3.设计一个hash方法
// @lc code=start
type MyHashSet struct {
	D []int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{D: []int{-1, 0}}
}

func (this *MyHashSet) Hash(key int) int {
	return key & (len(this.D) - 1)
}

func (this *MyHashSet) Add(key int) {
	k := this.Hash(key)
	for this.D[k] != key && this.Hash(this.D[k]) == k {
		this.D = append(this.D, this.D...)
		k = this.Hash(key)
	}
	this.D[k] = key
}

func (this *MyHashSet) Remove(key int) {
	k := this.Hash(key)
	if this.D[k] == key {
		this.D[k] = k - 1
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	k := this.Hash(key)
	return this.D[k] == key
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
// @lc code=end

