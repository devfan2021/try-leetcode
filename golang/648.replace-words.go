import "strings"

/*
 * @lc app=leetcode id=648 lang=golang
 *
 * [648] Replace Words
 *
 * https://leetcode.com/problems/replace-words/description/
 *
 * algorithms
 * Medium (56.02%)
 * Likes:    745
 * Dislikes: 131
 * Total Accepted:    56.3K
 * Total Submissions: 99.5K
 * Testcase Example:  '["cat","bat","rat"]\n"the cattle was rattled by the battery"'
 *
 * In English, we have a concept called root, which can be followed by some
 * other words to form another longer word - let's call this word successor.
 * For example, the root an, followed by other, which can form another word
 * another.
 *
 * Now, given a dictionary consisting of many roots and a sentence. You need to
 * replace all the successor in the sentence with the root forming it. If a
 * successor has many roots can form it, replace it with the root with the
 * shortest length.
 *
 * You need to output the sentence after the replacement.
 *
 *
 * Example 1:
 *
 *
 * Input: dict = ["cat","bat","rat"], sentence = "the cattle was rattled by the
 * battery"
 * Output: "the cat was rat by the bat"
 *
 *
 *
 * Constraints:
 *
 *
 * The input will only have lower-case letters.
 * 1 <= dict.length <= 1000
 * 1 <= dict[i].length <= 100
 * 1 <= sentence words number <= 1000
 * 1 <= sentence words length <= 1000
 *
 *
 */

// @lc code=start
func replaceWords(dict []string, sentence string) string {
	return replaceWords2(dict, sentence)
}

// using map, strings.Split, strings.Join
func replaceWords1(dict []string, sentence string) string {
	if len(sentence) == 0 {
		return ""
	}

	hash := map[string]bool{}
	for _, v := range dict {
		hash[v] = true
	}
	words := strings.Split(sentence, " ")
	for i := 0; i < len(words); i++ {
		word := words[i]
		var prefix string
		for j := 1; j <= len(word); j++ {
			if hash[word[:j]] {
				prefix = word[:j]
				break
			}
		}
		if len(prefix) > 0 {
			words[i] = prefix
		}
	}
	return strings.Join(words, " ")
}

type Trie struct {
	IsEnd    bool
	Children map[rune]*Trie
}

func Constructor() Trie {
	return Trie{Children: map[rune]*Trie{}}
}

func (this *Trie) Insert(word string) {
	parent := this
	for _, c := range word {
		if child, ok := parent.Children[c]; ok {
			parent = child
		} else {
			child := &Trie{Children: map[rune]*Trie{}}
			parent.Children[c] = child
			parent = child
		}
	}
	parent.IsEnd = true
}

func (this *Trie) Search(word string) bool {
	parent := this
	for _, c := range word {
		if child, ok := parent.Children[c]; ok {
			parent = child
		} else {
			child := &Trie{Children: map[rune]*Trie{}}
			parent.Children[c] = child
			parent = child
		}
	}
	return parent.IsEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	parent := this
	for _, c := range prefix {
		if _, ok := parent.Children[c]; !ok {
			return false
		}
		parent = parent.Children[c]
	}
	return true
}

// trie: find the smallest root that was a prefix in linear time.
func replaceWords2(dict []string, sentence string) string {
	if len(sentence) == 0 {
		return ""
	}

	trie := Constructor()
	for _, v := range dict {
		trie.Insert(v)
	}
	words := strings.Split(sentence, " ")
	for i := 0; i < len(words); i++ {
		word := words[i]
		var prefix string
		for j := 1; j <= len(word); j++ {
			if trie.Search(word[:j]) {
				prefix = word[:j]
				break
			}
		}
		if len(prefix) > 0 {
			words[i] = prefix
		}
	}
	return strings.Join(words, " ")
}

// @lc code=end