package leetcode

import (
	"fmt"
)

type TrieNode struct {
	value  byte
	isWord bool
	child  [26]*TrieNode
}

func trieNode0() *TrieNode {
	//return &TrieNode{0, false, nil}
	n := new(TrieNode)
	n.isWord = false
	n.value = 0
	return n
}

func trieNode1(c byte) *TrieNode {
	//return &TrieNode{c, false, nil}
	n := new(TrieNode)
	n.isWord = false
	n.value = c
	return n
}

type Trie struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{trieNode0()}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	n := this.root
	for _, v := range word {
		if n.child[v-'a'] == nil {
			n.child[v-'a'] = trieNode1(byte(v))
		}
		n = n.child[v-'a']
	}
	n.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	n := this.root
	for _, v := range word {
		if n.child[v-'a'] == nil {
			return false
		}
		n = n.child[v-'a']
	}
	return n.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	n := this.root
	for _, v := range prefix {
		if n.child[v-'a'] == nil {
			return false
		}
		n = n.child[v-'a']
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

func Test_myTrie() {
	fmt.Println("test my trie")
	obj := Constructor();
	obj.Insert("apple");
	param_2 := obj.Search("appl");
	param_3 := obj.StartsWith("appl");
	fmt.Println("p2 = ", param_2)
	fmt.Println("p3 = ", param_3)
}
