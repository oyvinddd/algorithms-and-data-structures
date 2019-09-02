package dfs

import (
	"fmt"

	"github.com/oyvinddd/algorithms/datastructures/graphs"
)

// DFS struct
type DFS struct {
	marked []bool
	count  int
}

// NewDFS constructor
func NewDFS() *DFS {
	return &DFS{count: 0}
}

// DepthFirstSearch executes a recursive depth-first search in the input graph
func (search DFS) DepthFirstSearch(g *graphs.Graph, s int) {
	search.marked = make([]bool, g.NoOfV())
	search.dfs(g, s)
}

// IsConnected checks if input graph is connected
func (search DFS) IsConnected(g *graphs.Graph) bool {
	search.marked = make([]bool, g.NoOfV())
	search.dfs(g, 0)
	return search.count == g.NoOfV()
}

// HasCycle checks if input graph has a cycle or it it is a tree
func HasCycle(g *graphs.Graph) bool {
	// TODO: implement
	return false
}

func (search *DFS) dfs(g *graphs.Graph, v int) {
	fmt.Printf("DFS: visiting node: %v\n", v)
	search.marked[v] = true
	search.count++
	for w := range g.GetAdj(v) {
		if !search.marked[w.(int)] {
			search.dfs(g, w.(int))
		}
	}
}
