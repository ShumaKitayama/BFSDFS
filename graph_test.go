package bfsdfs

import "testing"

func TestTraversals(t *testing.T) {
	g := Graph[string]{
		"A": {"B", "C"},
		"B": {"D", "E"},
		"C": {"F"},
		"D": nil,
		"E": nil,
		"F": nil,
	}
	wantBFS := []string{"A", "B", "C", "D", "E", "F"}
	gotBFS := BFS(g, "A")
	if len(gotBFS) != len(wantBFS) {
		t.Fatalf("BFS len mismatch: got %v", gotBFS)
	}
	for i, v := range wantBFS {
		if gotBFS[i] != v {
			t.Fatalf("BFS wrong at %d: got %s want %s", i, gotBFS[i], v)
		}
	}
	wantDFS := []string{"A", "B", "D", "E", "C", "F"}
	gotDFS := DFS(g, "A")
	for i, v := range wantDFS {
		if gotDFS[i] != v {
			t.Fatalf("DFS wrong at %d: got %s want %s", i, gotDFS[i], v)
		}
	}
}