package leetcode

import "fmt"

type TrieN struct {
	value  byte
	isWord bool
	child  [26]*TrieN
}

func trieN0() *TrieN {
	return &TrieN{0, false, [26]*TrieN{}}
	//n := new(TrieN)
	//n.isWord = false
	//n.value = 0
	//return n
}

func trieN1(c byte) *TrieN {
	return &TrieN{c, false, [26]*TrieN{}}
	//n := new(TrieN)
	//n.isWord = false
	//n.value = c
	//return n
}

type SearchII struct {
	dx   [4]int
	dy   [4]int
	root *TrieN
	m    int
	n    int
	res  map[string]string
}

/** Inserts a word into the trie. */
func (this *SearchII) Insert(word string) {
	n := this.root
	for _, v := range word {
		if n.child[v-'a'] == nil {
			n.child[v-'a'] = trieN1(byte(v))
		}
		n = n.child[v-'a']
	}
	n.isWord = true
}

func (this *SearchII) dfs(board [][]byte, i int, j int, cur_word string, cur_dict *TrieN) {
	cur_word += string(board[i][j])
	cur_dict = cur_dict.child[board[i][j]-'a']

	if cur_dict.isWord {
		this.res[cur_word] = cur_word
	}

	tmp := board[i][j]
	board[i][j] = '@'

	for k := 0; k < 4; k++ {
		x := i + this.dx[k]
		y := j + this.dy[k]
		if x >= 0 && x < this.m &&
			y >= 0 && y < this.n &&
			board[x][y] != '@' &&
			cur_dict.child[board[x][y]-'a'] != nil {
			this.dfs(board, x, y, cur_word, cur_dict)
		}
	}

	board[i][j] = tmp
}

func findWords(board [][]byte, words []string) []string {
	fmt.Println("finding words")
	if len(words) == 0 {
		return []string{}
	}
	if len(board) == 0 || len(board[0]) == 0 {
		return []string{}
	}
	var s SearchII
	s.root = trieN0()
	s.dx = [4]int{-1, 1, 0, 0}
	s.dy = [4]int{0, 0, -1, 1}
	s.m = len(board)
	s.n = len(board[0])
	s.res = make(map[string]string)

	for _, word := range words {
		s.Insert(word)
	}

	for i := 0; i < s.m; i++ {
		for j := 0; j < s.n; j++ {
			if s.root.child[board[i][j]-'a'] != nil {
				var cur_word string
				s.dfs(board, i, j, cur_word, s.root)
			}
		}
	}

	var ret []string

	for _, word := range s.res {
		ret = append(ret, word)
	}

	return ret
}

func Test_mySearchII() {
	words := []string{"oath", "pea", "eat", "rain"}
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'}}

	fmt.Println(findWords(board, words))

}
