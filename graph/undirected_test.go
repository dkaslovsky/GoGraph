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

func setupGraph() *testGraph {
	nodes := []string{"a", "b", "c", "d"}
	edges := []byte(
		`a b 1.5
		a c 2
		b c 3.3
		c d 1.1
		d a 7`,
	)
	reader := ioutil.NopCloser(bytes.NewReader(edges))
	g, _ := NewGraph("test", reader)

	return &testGraph{
		g:     g,
		nodes: nodes,
	}
}

func TestNewGraph(t *testing.T) {
	t.Run("empty graph", func(t *testing.T) {
		g, err := NewGraph("test")
		assert.Nil(t, err)
		assert.Equal(t, "test", g.Name)
		assert.Empty(t, g.dirAdj)
	})
	t.Run("graph from reader", func(t *testing.T) {
		reader := ioutil.NopCloser(bytes.NewReader([]byte("a b\na c 1.5\nc b 2.3")))
		g, err := NewGraph("test", reader)
		assert.Nil(t, err)

		assert.True(t, g.HasEdge("a", "b"))
		ab, _ := g.GetEdgeWeight("a", "b")
		assert.Equal(t, 1.0, ab)

		assert.True(t, g.HasEdge("a", "c"))
		ac, _ := g.GetEdgeWeight("a", "c")
		assert.Equal(t, 1.5, ac)

		assert.True(t, g.HasEdge("c", "b"))
		cb, _ := g.GetEdgeWeight("c", "b")
		assert.Equal(t, 2.3, cb)
	})
}

func TestGraphAddEdgeDefaultWeight(t *testing.T) {
	tests := map[string]struct {
		src string
		tgt string
	}{
		"add edge to new nodes": {
			src: "x",
			tgt: "y",
		},
		"add edge from new node to existing node": {
			src: "x",
			tgt: "a",
		},
		"add edge from existing node to new node": {
			src: "a",
			tgt: "x",
		},
		"add new edge from existing node to existing node": {
			src: "b",
			tgt: "d",
		},
		"upsert existing edge": {
			src: "a",
			tgt: "b",
		},
	}

	defaultWgt := 1.0
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			tg := setupGraph()
			g := tg.g
			g.AddEdge(test.src, test.tgt)
			assert.True(t, g.dirAdj.HasEdge(test.src, test.tgt))
			assert.True(t, g.dirAdj.HasEdge(test.tgt, test.src))
			wgt, _ := g.dirAdj.GetEdgeWeight(test.src, test.tgt)
			assert.Equal(t, defaultWgt, wgt)
			wgt, _ = g.dirAdj.GetEdgeWeight(test.tgt, test.src)
			assert.Equal(t, defaultWgt, wgt)
		})
	}
}

func TestGraphAddEdgeNonDefaultWeight(t *testing.T) {
	tests := map[string]struct {
		src string
		tgt string
		wgt float64
	}{
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
			tg := setupGraph()
			g := tg.g
			g.AddEdge(test.src, test.tgt, test.wgt)
			assert.True(t, g.dirAdj.HasEdge(test.src, test.tgt))
			assert.True(t, g.dirAdj.HasEdge(test.tgt, test.src))
			wgt, _ := g.dirAdj.GetEdgeWeight(test.src, test.tgt)
			assert.Equal(t, test.wgt, wgt)
			wgt, _ = g.dirAdj.GetEdgeWeight(test.tgt, test.src)
			assert.Equal(t, test.wgt, wgt)
		})
	}
}
