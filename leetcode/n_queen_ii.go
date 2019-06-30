package leetcode

import "fmt"

type DFS struct {
	count int
}

func (d *DFS) dfs(n int, row int, cols int, pie int, na int)  {
	if row >= n {
		d.count += 1
		return
	}
	bits := (^(cols | pie | na) & ((1 << uint(n)) - 1))
	for bits != 0 {
		p := bits & -bits
		d.dfs(n, row + 1, cols | p, (pie | p) << 1, (na | p) >> 1)
		bits = bits & (bits - 1)
	}
}

func totalNQueens(n int) int {
	if n < 1 {
		return 0
	}
	var d DFS
	d.count = 0
	d.dfs(n, 0, 0, 0, 0)
	return d.count;
}





func Test_nQueensII() {
	ret := totalNQueens(8)
	fmt.Println("ret = ", ret)
}