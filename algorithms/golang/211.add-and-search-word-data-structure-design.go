/*
 * @lc app=leetcode id=211 lang=golang
 *
 * [211] Add and Search Word - Data structure design
 *
 * https://leetcode.com/problems/add-and-search-word-data-structure-design/description/
 *
 * algorithms
 * Medium (36.20%)
 * Likes:    2182
 * Dislikes: 100
 * Total Accepted:    221.4K
 * Total Submissions: 581.5K
 * Testcase Example:  '["WordDictionary","addWord","addWord","addWord","search","search","search","search"]\n' +
  '[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]'
 *
 * Design a data structure that supports the following two operations:
 *
 *
 * void addWord(word)
 * bool search(word)
 *
 *
 * search(word) can search a literal word or a regular expression string
 * containing only letters a-z or .. A . means it can represent any one
 * letter.
 *
 * Example:
 *
 *
 * addWord("bad")
 * addWord("dad")
 * addWord("mad")
 * search("pad") -> false
 * search("bad") -> true
 * search(".ad") -> true
 * search("b..") -> true
 *
 *
 * Note:
 * You may assume that all words are consist of lowercase letters a-z.
 *
*/

// @lc code=start
type WordDictionary struct {
	IsEnd    bool
	Children [26]*WordDictionary
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	parent := this
	for _, c := range word {
		index := c - 'a'
		if parent.Children[index] == nil {
			parent.Children[index] = &WordDictionary{}
		}
		parent = parent.Children[index]
	}
	parent.IsEnd = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	if len(word) == 0 {
		return this.IsEnd
	}

	if word[0] == '.' {
		// try all possible node
		for j := 0; j < 26; j++ {
			if this.Children[j] != nil && this.Children[j].Search(word[1:]) {
				return true
			}
		}
	} else {
		index := word[0] - 'a'
		// if it has child, keep searching
		if child := this.Children[index]; child != nil {
			return child.Search(word[1:])
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end