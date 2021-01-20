package main

import (
	"fmt"
)

func main() {
	var q int

	fmt.Scanf("%d\n", &q)
	for ; q > 0; q-- {
		solve()
	}

}

func solve() {
	var n, m, s int
	fmt.Scanf("%d%d\n", &n, &m)

	graph := make([]map[int]bool, n+1)
	for i := 1; i <= n; i++ {
		graph[i] = make(map[int]bool)
	}

	for ; m > 0; m-- {
		var u, v int
		fmt.Scanf("%d%d\n", &u, &v)

		graph[u][v] = true
		graph[v][u] = true
	}

	fmt.Scanf("%d\n", &s)
	qu := []*pair{}

	qu = append(qu, &pair{N: s, L: 0})
	visited := make([]bool, n+1)
	visited[s] = true
	dist := make([]int, n+1)

	for len(qu) > 0 {
		var p *pair
		p, qu = qu[0], qu[1:]
		dist[p.N] = p.L
		for adj := range graph[p.N] {
			if !visited[adj] {
				visited[adj] = true
				qu = append(qu, &pair{N: adj, L: p.L + 1})
			}
		}
	}
	fmt.Printf("%v\n", dist)
	for i := 1; i <= n; i++ {
		if i != s {
			if dist[i] == 0 {
				fmt.Printf("%d ", -1)
			} else {
				fmt.Printf("%d ", dist[i]*6)
			}
		}
	}

	fmt.Println()
}

type pair struct {
	N, L int
}
