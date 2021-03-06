import "math/rand"

/*
 * @lc app=leetcode id=380 lang=golang
 *
 * [380] Insert Delete GetRandom O(1)
 *
 * https://leetcode.com/problems/insert-delete-getrandom-o1/description/
 *
 * algorithms
 * Medium (47.15%)
 * Likes:    2473
 * Dislikes: 158
 * Total Accepted:    240.7K
 * Total Submissions: 508.7K
 * Testcase Example:  '["RandomizedSet","insert","remove","insert","getRandom","remove","insert","getRandom"]\n' +
  '[[],[1],[2],[2],[],[1],[2],[]]'
 *
 * Design a data structure that supports all following operations in average
 * O(1) time.
 *
 *
 *
 *
 * insert(val): Inserts an item val to the set if not already present.
 * remove(val): Removes an item val from the set if present.
 * getRandom: Returns a random element from current set of elements (it's
 * guaranteed that at least one element exists when this method is called).
 * Each element must have the same probability of being returned.
 *
 *
 *
 *
 * Example:
 *
 *
 * // Init an empty set.
 * RandomizedSet randomSet = new RandomizedSet();
 *
 * // Inserts 1 to the set. Returns true as 1 was inserted successfully.
 * randomSet.insert(1);
 *
 * // Returns false as 2 does not exist in the set.
 * randomSet.remove(2);
 *
 * // Inserts 2 to the set, returns true. Set now contains [1,2].
 * randomSet.insert(2);
 *
 * // getRandom should return either 1 or 2 randomly.
 * randomSet.getRandom();
 *
 * // Removes 1 from the set, returns true. Set now contains [2].
 * randomSet.remove(1);
 *
 * // 2 was already in the set, so return false.
 * randomSet.insert(2);
 *
 * // Since 2 is the only number in the set, getRandom always return 2.
 * randomSet.getRandom();
 *
 *
*/

// @lc code=start
type RandomizedSet struct {
	values    []int
	vIndexMap map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	m := RandomizedSet{vIndexMap: map[int]int{}}
	return m
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.vIndexMap[val]; ok {
		return false
	}
	this.vIndexMap[val] = len(this.values)
	this.values = append(this.values, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if index, ok := this.vIndexMap[val]; ok {
		// ???????????????????????????????????????????????????????????????????????????
		this.values[index] = this.values[len(this.values)-1]
		this.vIndexMap[this.values[index]] = index
		this.values = this.values[:len(this.values)-1]
		delete(this.vIndexMap, val)
		return true
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	if len(this.values) == 0 {
		return 0
	}
	return this.values[rand.Intn(len(this.values))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end