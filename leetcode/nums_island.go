package leetcode

import "fmt"

type UnionFinder struct {
	counts int
	parent []int
	rank   []int
}

func (u *UnionFinder) init(grid [][]byte) {
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			u.rank = append(u.rank, 0)
			u.parent = append(u.parent, -1)

			if grid[i][j] == '1' {
				u.counts++
				u.parent[i*n+j] = i*n + j
			}
		}
	}
	m += 1
}

func (u *UnionFinder) find(i int) int {
	if u.parent[i] != i {
		u.parent[i] = u.find(u.parent[i])
	}
	return u.parent[i]
}

func (u *UnionFinder) union(x int, y int) {
	rootx := u.find(x)
	rooty := u.find(y)

	if rootx != rooty {
		if u.rank[rootx] > u.rank[rooty] {
			u.parent[rooty] = rootx
		} else if u.rank[rootx] < u.rank[rooty] {
			u.parent[rootx] = rooty
		} else {
			u.parent[rooty] = rootx
			u.rank[rootx] += 1
		}
		u.counts--
	}
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	uf := &UnionFinder{}
	uf.init(grid)
	ds := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}
			for _, d := range ds {
				dx := i + d[0]
				dy := j + d[1]
				if dx >= 0 && dy >= 0 && dx < m && dy < n && grid[dx][dy] == '1' {
					uf.union(i*n+j, dx*n+dy)
				}
			}
		}
	}
	return uf.counts
}

func Test_NumsIsland() {
	grid2 := [][]byte{{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'}}
	grid1 := [][]byte{{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'}}
	ret1 := numIslands(grid1)
	ret2 := numIslands(grid2)

	fmt.Println("ret1 = ", ret1)
	fmt.Println("ret2 = ", ret2)
}
