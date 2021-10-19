package main

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

func (t *TrieNode) Insert(word string) {
	node := t
	for _, ch := range word {
		if node.children[ch-'a'] == nil {
			node.children[ch-'a'] = &TrieNode{}
		}
		node = node.children[ch-'a']
	}
	node.isEnd = true
}

type WordDictionary struct {
	trieNode *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{trieNode: &TrieNode{}}
}

func (this *WordDictionary) AddWord(word string) {
	this.trieNode.Insert(word)
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(int, *TrieNode) bool
	dfs = func(index int, node *TrieNode) bool {
		if index == len(word) {
			return node.isEnd
		}

		ch := word[index]
		if ch != '.' {
			child := node.children[ch-'a']
			if child != nil && dfs(index+1, child) {
				return true
			}
		} else {
			for _, child := range node.children {
				if child != nil && dfs(index+1, child) {
					return true
				}
			}
		}
		return false
	}
	return dfs(0, this.trieNode)
}

func main() {

}
