package bfsdfs

// Graph は隣接リスト表現
type Graph[T comparable] map[T][]T

// BFS は幅優先探索の訪問順を返す
func BFS[T comparable](g Graph[T], start T) []T {
	visited := make(map[T]bool, len(g))
	if _, ok := g[start]; !ok {
		return nil // 始点が存在しない場合は空スライス
	}
	queue := []T{start}
	visited[start] = true
	var order []T
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		order = append(order, v)
		for _, n := range g[v] {
			if !visited[n] {
				visited[n] = true
				queue = append(queue, n)
			}
		}
	}
	return order
}

// DFS は深さ優先探索（反復実装）の訪問順を返す
func DFS[T comparable](g Graph[T], start T) []T {
	visited := make(map[T]bool, len(g))
	if _, ok := g[start]; !ok {
		return nil
	}
	stack := []T{start}
	var order []T
	for len(stack) > 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if visited[v] {
			continue
		}
		visited[v] = true
		order = append(order, v)
		neighbors := g[v]
		// 逆順 push で左→右順序を再現
		for i := len(neighbors) - 1; i >= 0; i-- {
			n := neighbors[i]
			if !visited[n] {
				stack = append(stack, n)
			}
		}
	}
	return order
}