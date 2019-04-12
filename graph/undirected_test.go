package graph

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testGraph struct {
	g     *Graph
	nodes []string
}

func setupTestGraph() *testGraph {
	nodes := []string{"a", "b", "c", "d"}
	edges := []byte("a b 1.5\na c 2\nb c 3.3\nc d 1.1\nd a 7")
	reader := ioutil.NopCloser(bytes.NewReader(edges))
	g, _ := NewGraph("test", reader)
	return &testGraph{
		g:     g,
		nodes: nodes,
	}
}

func (te testEdge) symmetricExistsIn(a dirAdj) bool {
	// test edge from src to tgt
	n, ok := a[te.src]
	if !ok {
		return false
	}
	w, ok := n[te.tgt]
	if !ok {
		return false
	}
	return w == te.wgt
}

func TestNewGraph(t *testing.T) {
	t.Run("empty graph", func(t *testing.T) {
		g, err := NewGraph("test")
		assert.Nil(t, err)
		assert.Equal(t, "test", g.Name)
		assert.Empty(t, g.dirAdj)
	})
	t.Run("graph from reader", func(t *testing.T) {
		f := []byte("a b\na c 1.5\nc b 2.3")
		nodes := []string{"a", "b", "c"}
		edges := []testEdge{
			testEdge{src: "a", tgt: "b", wgt: 1.0},
			testEdge{src: "a", tgt: "c", wgt: 1.5},
			testEdge{src: "c", tgt: "b", wgt: 2.3},
		}

		reader := ioutil.NopCloser(bytes.NewReader(f))
		g, err := NewGraph("test", reader)
		assert.Nil(t, err)
		assert.ElementsMatch(t, nodes, g.dirAdj.getSrcNodes())
		assert.ElementsMatch(t, nodes, g.invAdj.getSrcNodes())

		// test edges
		a := *g.dirAdj
		invA := *g.invAdj
		for _, e := range edges {
			assert.True(t, e.symmetricExistsIn(a))
			assert.True(t, e.symmetricExistsIn(invA))
		}
	})
}

func TestGraphAddEdgeDefaultWeight(t *testing.T) {
	tests := map[string]testEdge{
		"add edge to new nodes": {
			src: "x",
			tgt: "y",
			wgt: defaultWgt,
		},
		"add edge from new node to existing node": {
			src: "x",
			tgt: "a",
			wgt: defaultWgt,
		},
		"add edge from existing node to new node": {
			src: "a",
			tgt: "x",
			wgt: defaultWgt,
		},
		"add new edge from existing node to existing node": {
			src: "b",
			tgt: "d",
			wgt: defaultWgt,
		},
		"upsert existing edge": {
			src: "a",
			tgt: "b",
			wgt: defaultWgt,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			tg := setupTestGraph()
			g := tg.g
			g.AddEdge(test.src, test.tgt)
			assert.True(t, test.symmetricExistsIn(*g.dirAdj))
			assert.True(t, test.symmetricExistsIn(*g.invAdj))
		})
	}
}

func TestGraphAddEdgeNonDefaultWeight(t *testing.T) {
	tests := map[string]testEdge{
		"add edge to new nodes": {
			src: "x",
			tgt: "y",
			wgt: 3.7,
		},
		"add edge from new node to existing node": {
			src: "x",
			tgt: "a",
			wgt: 4.2,
		},
		"add edge from existing node to new node": {
			src: "a",
			tgt: "x",
			wgt: 9.0,
		},
		"add new edge from existing node to existing node": {
			src: "b",
			tgt: "d",
			wgt: 18.0,
		},
		"upsert existing edge": {
			src: "a",
			tgt: "b",
			wgt: 1.11,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			tg := setupTestGraph()
			g := tg.g
			g.AddEdge(test.src, test.tgt, test.wgt)
			assert.True(t, test.symmetricExistsIn(*g.dirAdj))
			assert.True(t, test.symmetricExistsIn(*g.invAdj))
		})
	}
}

func TestGraphRemoveEdge(t *testing.T) {
	tests := map[string]testEdge{
		"remove nonexistent edge": {
			src: "x",
			tgt: "y",
		},
		"remove nonexistent edge, src exists": {
			src: "a",
			tgt: "y",
		},
		"remove nonexistent edge, tgt exists": {
			src: "x",
			tgt: "a",
		},
		"remove existing edge": {
			src: "a",
			tgt: "b",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			tg := setupTestGraph()
			g := tg.g
			g.RemoveEdge(test.src, test.tgt)
			assert.False(t, test.symmetricExistsIn(*g.dirAdj))
			assert.False(t, test.symmetricExistsIn(*g.invAdj))
		})
	}
}

func TestGraphRemoveNode(t *testing.T) {
	tests := map[string]struct {
		node string
	}{
		"remove nonexistent node": {
			node: "x",
		},
		"remove existing node": {
			node: "a",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			tg := setupTestGraph()
			g := tg.g
			g.RemoveNode(test.node)

			assert.NotContains(t, *g.dirAdj, test.node)
			assert.NotContains(t, *g.invAdj, test.node)
			for node, nbrs := range *g.dirAdj {
				_, ok := nbrs[node]
				assert.False(t, ok)
			}
			for node, nbrs := range *g.invAdj {
				_, ok := nbrs[node]
				assert.False(t, ok)
			}
		})
	}
}

func TestGraphGetNodes(t *testing.T) {
	tests := map[string]testGraph{
		"empty graph": {
			g: &Graph{dirAdj: &dirAdj{}},
		},
		"nonempty graph": *setupTestGraph(),
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			nodes := test.g.GetNodes()
			assert.ElementsMatch(t, test.nodes, nodes)
		})
	}
}

func TestGraphGetInvNeighbors(t *testing.T) {
	t.Run("inverse neighbors equal neighbors", func(t *testing.T) {
		tg := setupTestGraph()
		for _, node := range tg.nodes {
			nbrs, ok := tg.g.GetInvNeighbors(node)
			assert.True(t, ok)
			expectedNbrs, _ := tg.g.GetNeighbors(node)
			assert.Equal(t, expectedNbrs, nbrs)
		}
	})
}

func TestGraphGetInDegree(t *testing.T) {
	t.Run("in degree equals out degree", func(t *testing.T) {
		tg := setupTestGraph()
		for _, node := range tg.nodes {
			d, ok := tg.g.GetInDegree(node)
			assert.True(t, ok)
			expectedDeg, _ := tg.g.GetOutDegree(node)
			assert.Equal(t, expectedDeg, d)
		}
	})
}
