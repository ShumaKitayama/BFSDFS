package bfsdfs_test

import (
	"fmt"

	"github.com/shumakitayama/bfsdfs"
)

func Example() {
	g := bfsdfs.Graph[int]{
		1: []int{2, 3},
		2: []int{4, 5},
		3: nil,
		4: nil,
		5: nil,
	}
	fmt.Println("BFS:", bfsdfs.BFS(g, 1))
	fmt.Println("DFS:", bfsdfs.DFS(g, 1))
	// Output:
	// BFS: [1 2 3 4 5]
	// DFS: [1 2 4 5 3]
}