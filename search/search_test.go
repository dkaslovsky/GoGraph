package search

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dkaslovsky/GoGraph/graph"
	n "github.com/dkaslovsky/GoGraph/node"
)

func setupEmptyGraph() *graph.DirGraph {
	g, _ := graph.NewDirGraph("empty")
	return g
}

func setupGraph() *graph.DirGraph {
	g, _ := graph.NewDirGraph("no loops")
	g.AddEdge("a", "z")
	g.AddEdge("a", "b")
	g.AddEdge("b", "c")
	g.AddEdge("c", "d")
	g.AddEdge("c", "e")
	return g
}

func TestDFS(t *testing.T) {
	tests := map[string]struct {
		g             *graph.DirGraph
		start         n.Node
		expectedFound []n.Node
	}{
		"empty graph": {
			g:             setupEmptyGraph(),
			start:         "a",
			expectedFound: []n.Node{},
		},
		"starting at root": {
			g:             setupGraph(),
			start:         "a",
			expectedFound: []n.Node{"a", "b", "c", "d", "e", "z"},
		},
		"starting below root": {
			g:             setupGraph(),
			start:         "b",
			expectedFound: []n.Node{"b", "c", "d", "e"},
		},
		//"starting at root in graph with self loops"
		//"starting below root in graph with self loops"
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			found := DFS(test.g, test.start)
			assert.ElementsMatch(t, found, test.expectedFound)
		})
	}
}
