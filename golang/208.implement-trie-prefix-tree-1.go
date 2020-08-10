/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 *
 * https://leetcode.com/problems/implement-trie-prefix-tree/description/
 *
 * algorithms
 * Medium (48.18%)
 * Likes:    3458
 * Dislikes: 52
 * Total Accepted:    335.3K
 * Total Submissions: 679.2K
 * Testcase Example:  '["Trie","insert","search","search","startsWith","insert","search"]\n' +
  '[[],["apple"],["apple"],["app"],["app"],["app"],["app"]]'
 *
 * Implement a trie with insert, search, and startsWith methods.
 *
 * Example:
 *
 *
 * Trie trie = new Trie();
 *
 * trie.insert("apple");
 * trie.search("apple");   // returns true
 * trie.search("app");     // returns false
 * trie.startsWith("app"); // returns true
 * trie.insert("app");
 * trie.search("app");     // returns true
 *
 *
 * Note:
 *
 *
 * You may assume that all inputs are consist of lowercase letters a-z.
 * All inputs are guaranteed to be non-empty strings.
 *
 *
*/
// https://leetcode.com/problems/implement-trie-prefix-tree/solution/
// @lc code=start
type Trie struct {
	IsEnd    bool
	Children []*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{IsEnd: false, Children: make([]*Trie, 26)}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	parent := this
	for _, c := range word {
		cInt := int(c - 'a')
		if parent.Children[cInt] != nil {
			parent = parent.Children[cInt]
		} else {
			newChild := &Trie{IsEnd: false, Children: make([]*Trie, 26)}
			parent.Children[cInt] = newChild
			parent = newChild
		}
	}
	parent.IsEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	parent := this
	for _, c := range word {
		cInt := int(c - 'a')
		if parent.Children[cInt] == nil {
			return false
		}
		parent = parent.Children[cInt]
	}
	return parent.IsEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	parent := this
	for _, c := range prefix {
		cInt := int(c - 'a')
		if parent.Children[cInt] != nil {
			return false
		}
		parent = parent.Children[cInt]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end