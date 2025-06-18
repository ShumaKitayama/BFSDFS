# bfsdfs: Go ジェネリクス対応グラフ探索ライブラリ

## 概要

`bfsdfs` は、Go のジェネリクス機能を活用したシンプルなグラフ探索ライブラリです。隣接リスト表現のグラフに対して、幅優先探索（BFS）および深さ優先探索（DFS）を簡単に実行できます。

- **BFS (Breadth-First Search):** 最短経路やレベル順探索に有効
- **DFS (Depth-First Search):** 経路探索や連結成分の発見などに有効

## インストール方法

Go モジュールとして利用できます。

```sh
go get <このリポジトリのURL>
```

## 使い方

### 1. グラフの定義

ノード型 `T`（例: int, string など比較可能な型）をキー、隣接ノードのスライスを値とするマップでグラフを表現します。

```go
import "bfsdfs"

graph := bfsdfs.Graph[int]{
    1: {2, 3},
    2: {1, 4},
    3: {1, 4},
    4: {2, 3},
}
```

### 2. 幅優先探索（BFS）の実行

```go
order := bfsdfs.BFS(graph, 1) // 1 を始点とした訪問順
fmt.Println(order) // 例: [1 2 3 4]
```

### 3. 深さ優先探索（DFS）の実行

```go
order := bfsdfs.DFS(graph, 1) // 1 を始点とした訪問順
fmt.Println(order) // 例: [1 3 4 2]
```

## API

### Graph 型

```go
type Graph[T comparable] map[T][]T
```

- 任意の比較可能な型 `T` をノードとして利用可能

### BFS 関数

```go
func BFS[T comparable](g Graph[T], start T) []T
```

- `g`: グラフ
- `start`: 探索開始ノード
- **返り値:** 訪問順のノードスライス（始点が存在しない場合は `nil`）

### DFS 関数

```go
func DFS[T comparable](g Graph[T], start T) []T
```

- `g`: グラフ
- `start`: 探索開始ノード
- **返り値:** 訪問順のノードスライス（始点が存在しない場合は `nil`）

## 注意事項

- 始点ノードがグラフに存在しない場合、`nil` を返します。
- グラフは無向・有向どちらも対応可能です（隣接リストの定義次第）。
- ノード型は `comparable` 制約を満たす必要があります。

## ライセンス

MIT
