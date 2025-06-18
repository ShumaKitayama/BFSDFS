package main

import (
	"fmt"

	"github.com/shumakitayama/bfsdfs"
)

// Pos は格子上の一点を表す
type Pos struct{ R, C int }

// 迷路定義（#=壁, .=道, S=スタート, G=ゴール）
var maze = []string{
	"###########",
	"#S#...#...#",
	"#.#.#.#.#.#",
	"#.#.#.#.#.#",
	"#...#...#G#",
	"###########",
}

// 迷路 → グラフ＋始点終点 を生成
func buildGraph() (g bfsdfs.Graph[Pos], start, goal Pos) {
	rows, cols := len(maze), len(maze[0])
	g = make(bfsdfs.Graph[Pos])
	dirs := []Pos{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			ch := maze[r][c]
			if ch == '#' {
				continue
			}
			p := Pos{r, c}
			if ch == 'S' {
				start = p
			} else if ch == 'G' {
				goal = p
			}
			// 近傍辺
			for _, d := range dirs {
				nr, nc := r+d.R, c+d.C
				if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
					continue
				}
				if maze[nr][nc] != '#' {
					g[p] = append(g[p], Pos{nr, nc})
				}
			}
		}
	}
	return
}

// BFS 順を利用して距離と親を構築
func shortestPath(g bfsdfs.Graph[Pos], s, t Pos) ([]Pos, bool) {
	order := bfsdfs.BFS(g, s)
	const inf = int(^uint(0) >> 1)
	dist := make(map[Pos]int, len(g))
	prev := make(map[Pos]Pos, len(g))
	for v := range g {
		dist[v] = inf
	}
	dist[s] = 0

	for _, u := range order {
		for _, v := range g[u] {
			if dist[v] == inf {
				dist[v] = dist[u] + 1
				prev[v] = u
			}
		}
	}
	if dist[t] == inf {
		return nil, false
	}

	// 後ろからたどって経路復元
	var path []Pos
	for v := t; v != s; v = prev[v] {
		path = append(path, v)
	}
	path = append(path, s)

	// 逆転
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path, true
}

// DFS を使って連結成分数と最大成分サイズを返す
func components(g bfsdfs.Graph[Pos]) (count, maxSize int) {
	seen := make(map[Pos]bool, len(g))
	for v := range g {
		if seen[v] {
			continue
		}
		comp := bfsdfs.DFS(g, v) // 同じ成分の頂点集合
		for _, x := range comp {
			seen[x] = true
		}
		count++
		if len(comp) > maxSize {
			maxSize = len(comp)
		}
	}
	return
}

func main() {
	g, S, G := buildGraph()

	path, ok := shortestPath(g, S, G)
	if !ok {
		fmt.Println("ゴールへ到達不可")
	} else {
		fmt.Printf("最短手数: %d\n", len(path)-1)
		fmt.Println("経路:")
		for _, p := range path {
			fmt.Printf("(%d,%d) ", p.R, p.C)
		}
		fmt.Println()
	}

	cc, max := components(g)
	fmt.Printf("通行可能領域の連結成分数: %d (最大サイズ %d)\n", cc, max)
}