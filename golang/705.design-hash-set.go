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

// @lc code=start
type MyHashSet struct {
	data []int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	val := MyHashSet{}
	return val
}

func (this *MyHashSet) Add(key int) {
	index := this.findByKey(key)
	if index == -1 {
		this.data = append(this.data, key)
	}
}

func (this *MyHashSet) Remove(key int) {
	index := this.findByKey(key)
	if index != -1 {
		length := len(this.data)
		this.data[index] = this.data[length-1]
		this.data[length-1] = 0
		this.data = this.data[:length-1]
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	index := this.findByKey(key)
	if index == -1 {
		return false
	}
	return true
}

func (this *MyHashSet) findByKey(key int) int {
	for index, v := range this.data {
		if v == key {
			return index
		}
	}
	return -1
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
// @lc code=end

