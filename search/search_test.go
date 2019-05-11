package search

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dkaslovsky/GoGraph/graph"
	n "github.com/dkaslovsky/GoGraph/node"
)

func setupEmptyGraph() *graph.Graph {
	g, _ := graph.NewGraph("empty")
	return g
}

func setupEmptyDirGraph() *graph.DirGraph {
	g, _ := graph.NewDirGraph("empty")
	return g
}

func setupGraph() *graph.Graph {
	g, _ := graph.NewGraph("no loops")
	g.AddEdge("a", "z")
	g.AddEdge("a", "b")
	g.AddEdge("b", "c")
	g.AddEdge("c", "d")
	g.AddEdge("c", "e")
	g.AddEdge("g", "h")
	g.AddEdge("h", "i")
	g.AddEdge("i", "i")
	return g
}

func setupDirGraph() *graph.DirGraph {
	g, _ := graph.NewDirGraph("no loops")
	g.AddEdge("a", "z")
	g.AddEdge("a", "b")
	g.AddEdge("b", "c")
	g.AddEdge("c", "d")
	g.AddEdge("c", "e")
	g.AddEdge("g", "h")
	g.AddEdge("h", "i")
	g.AddEdge("i", "i")
	return g
}

func TestDFS_UndirectedGraph(t *testing.T) {
	tests := map[string]struct {
		g             *graph.Graph
		start         n.Node
		expectedFound []n.Node
	}{
		"empty graph, undirected graph": {
			g:             setupEmptyGraph(),
			start:         "a",
			expectedFound: []n.Node{},
		},
		"starting at root, undirected graph": {
			g:             setupGraph(),
			start:         "a",
			expectedFound: []n.Node{"a", "b", "c", "d", "e", "z"},
		},
		"starting below root, undirected graph": {
			g:             setupGraph(),
			start:         "b",
			expectedFound: []n.Node{"a", "b", "c", "d", "e", "z"},
		},
		"starting from non-existent node, undirected graph": {
			g:             setupGraph(),
			start:         "x",
			expectedFound: []n.Node{},
		},
		"starting from terminal node, undirected graph": {
			g:             setupGraph(),
			start:         "z",
			expectedFound: []n.Node{"a", "b", "c", "d", "e", "z"},
		},
		"starting from node in component with self loop, undirected graph": {
			g:             setupGraph(),
			start:         "h",
			expectedFound: []n.Node{"g", "h", "i"},
		},
		"starting from terminal node with self loop in component, undirected graph": {
			g:             setupGraph(),
			start:         "i",
			expectedFound: []n.Node{"g", "h", "i"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			found := DFS(test.g, test.start)
			assert.ElementsMatch(t, found, test.expectedFound)
		})
	}
}

func TestDFS_DirectedGraph(t *testing.T) {
	tests := map[string]struct {
		g             *graph.DirGraph
		start         n.Node
		expectedFound []n.Node
	}{
		"empty graph, directed graph": {
			g:             setupEmptyDirGraph(),
			start:         "a",
			expectedFound: []n.Node{},
		},
		"starting at root, directed graph": {
			g:             setupDirGraph(),
			start:         "a",
			expectedFound: []n.Node{"a", "b", "c", "d", "e", "z"},
		},
		"starting below root, directed graph": {
			g:             setupDirGraph(),
			start:         "b",
			expectedFound: []n.Node{"b", "c", "d", "e"},
		},
		"starting from non-existent node, directed graph": {
			g:             setupDirGraph(),
			start:         "x",
			expectedFound: []n.Node{},
		},
		"starting from terminal node, directed graph": {
			g:             setupDirGraph(),
			start:         "z",
			expectedFound: []n.Node{"z"},
		},
		"starting from node in component with self loop, directed graph": {
			g:             setupDirGraph(),
			start:         "h",
			expectedFound: []n.Node{"h", "i"},
		},
		"starting from terminal node with self loop in component, directed graph": {
			g:             setupDirGraph(),
			start:         "i",
			expectedFound: []n.Node{"i"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			found := DFS(test.g, test.start)
			assert.ElementsMatch(t, found, test.expectedFound)
		})
	}
}

func TestBFS_UndirectedGraph(t *testing.T) {
	tests := map[string]struct {
		g             *graph.Graph
		start         n.Node
		expectedFound []n.Node
	}{
		"empty graph, undirected graph": {
			g:             setupEmptyGraph(),
			start:         "a",
			expectedFound: []n.Node{},
		},
		"starting at root, undirected graph": {
			g:             setupGraph(),
			start:         "a",
			expectedFound: []n.Node{"a", "b", "c", "d", "e", "z"},
		},
		"starting below root, undirected graph": {
			g:             setupGraph(),
			start:         "b",
			expectedFound: []n.Node{"a", "b", "c", "d", "e", "z"},
		},
		"starting from non-existent node, undirected graph": {
			g:             setupGraph(),
			start:         "x",
			expectedFound: []n.Node{},
		},
		"starting from terminal node, undirected graph": {
			g:             setupGraph(),
			start:         "z",
			expectedFound: []n.Node{"a", "b", "c", "d", "e", "z"},
		},
		"starting from node in component with self loop, undirected graph": {
			g:             setupGraph(),
			start:         "h",
			expectedFound: []n.Node{"g", "h", "i"},
		},
		"starting from terminal node with self loop in component, undirected graph": {
			g:             setupGraph(),
			start:         "i",
			expectedFound: []n.Node{"g", "h", "i"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			found := BFS(test.g, test.start)
			assert.ElementsMatch(t, found, test.expectedFound)
		})
	}
}

func TestBFS_DirectedGraph(t *testing.T) {
	tests := map[string]struct {
		g             *graph.DirGraph
		start         n.Node
		expectedFound []n.Node
	}{
		"empty graph, directed graph": {
			g:             setupEmptyDirGraph(),
			start:         "a",
			expectedFound: []n.Node{},
		},
		"starting at root, directed graph": {
			g:             setupDirGraph(),
			start:         "a",
			expectedFound: []n.Node{"a", "b", "c", "d", "e", "z"},
		},
		"starting below root, directed graph": {
			g:             setupDirGraph(),
			start:         "b",
			expectedFound: []n.Node{"b", "c", "d", "e"},
		},
		"starting from non-existent node, directed graph": {
			g:             setupDirGraph(),
			start:         "x",
			expectedFound: []n.Node{},
		},
		"starting from terminal node, directed graph": {
			g:             setupDirGraph(),
			start:         "z",
			expectedFound: []n.Node{"z"},
		},
		"starting from node in component with self loop, directed graph": {
			g:             setupDirGraph(),
			start:         "h",
			expectedFound: []n.Node{"h", "i"},
		},
		"starting from terminal node with self loop in component, directed graph": {
			g:             setupDirGraph(),
			start:         "i",
			expectedFound: []n.Node{"i"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			found := BFS(test.g, test.start)
			assert.ElementsMatch(t, found, test.expectedFound)
		})
	}
}
