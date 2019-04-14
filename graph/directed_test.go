package graph

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testDirGraph struct {
	dg       *DirGraph
	nodes    []string
	invNodes []string
}

func setupTestDirGraph() *testDirGraph {
	nodes := []string{"a", "b", "c", "d"}
	invNodes := []string{"a", "b", "c", "d"}
	edges := []byte("a b 1.5\na c 2\nb c 3.3\nc d 1.1\nd a 7\na d 19\nd d 3.1")
	reader := ioutil.NopCloser(bytes.NewReader(edges))
	dg, _ := NewDirGraph("test", reader)
	return &testDirGraph{
		dg:       dg,
		nodes:    nodes,
		invNodes: invNodes,
	}
}

func TestNewDirGraph(t *testing.T) {
	t.Run("empty graph", func(t *testing.T) {
		dg, err := NewDirGraph("test")
		assert.Nil(t, err)
		assert.Equal(t, "test", dg.Name)
		assert.Empty(t, dg.dirAdj)
		assert.Empty(t, dg.invAdj)
	})
	t.Run("graph from reader", func(t *testing.T) {
		f := []byte("a b\na c 1.5\nc b 2.3")
		nodes := []string{"a", "c"}
		invNodes := []string{"b", "c"}
		edges := []testEdge{
			testEdge{src: "a", tgt: "b", wgt: 1.0},
			testEdge{src: "a", tgt: "c", wgt: 1.5},
			testEdge{src: "c", tgt: "b", wgt: 2.3},
		}

		reader := ioutil.NopCloser(bytes.NewReader(f))
		dg, err := NewDirGraph("test", reader)
		assert.Nil(t, err)
		assert.ElementsMatch(t, nodes, dg.dirAdj.getSrcNodes())
		assert.ElementsMatch(t, invNodes, dg.invAdj.getSrcNodes())

		// test edges
		a := *dg.dirAdj
		invA := *dg.invAdj
		for _, e := range edges {
			assert.True(t, a.HasEdge(e.src, e.tgt))
			assert.True(t, invA.HasEdge(e.tgt, e.src))
		}
	})
}

func TestDirGraphAddEdgeDefaultWeight(t *testing.T) {
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
		"add self loop": {
			src: "a",
			tgt: "a",
			wgt: defaultWgt,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			tg := setupTestDirGraph()
			dg := tg.dg
			dg.AddEdge(test.src, test.tgt)
			assert.True(t, dg.dirAdj.HasEdge(test.src, test.tgt))
			assert.True(t, dg.invAdj.HasEdge(test.tgt, test.src))
		})
	}
}

func TestDirGraphAddEdgeNonDefaultWeight(t *testing.T) {
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
		"add self loop": {
			src: "a",
			tgt: "a",
			wgt: 196.196,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			tg := setupTestDirGraph()
			dg := tg.dg
			dg.AddEdge(test.src, test.tgt, test.wgt)
			assert.True(t, dg.dirAdj.HasEdge(test.src, test.tgt))
			assert.True(t, dg.invAdj.HasEdge(test.tgt, test.src))
		})
	}
}

func TestDirGraphRemoveEdge(t *testing.T) {
	tests := map[string]struct {
		testEdge
		reverseExists bool
	}{
		"remove nonexistent edge": {
			testEdge: testEdge{
				src: "x",
				tgt: "y",
			},
			reverseExists: false,
		},
		"remove nonexistent edge, src exists": {
			testEdge: testEdge{
				src: "a",
				tgt: "y",
			},
			reverseExists: false,
		},
		"remove nonexistent edge, tgt exists": {
			testEdge: testEdge{
				src: "x",
				tgt: "a",
			},
			reverseExists: false,
		},
		"remove nonsymmetric existing edge": {
			testEdge: testEdge{
				src: "a",
				tgt: "b",
			},
			reverseExists: false,
		},
		"remove symmetric existing edge": {
			testEdge: testEdge{
				src: "a",
				tgt: "d",
			},
			reverseExists: true,
		},
		"remove self loop": {
			testEdge: testEdge{
				src: "d",
				tgt: "d",
			},
			// reverse is the same as forward so it
			// does not exist after forward is removed
			reverseExists: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			tg := setupTestDirGraph()
			dg := tg.dg
			dg.RemoveEdge(test.src, test.tgt)
			assert.False(t, dg.dirAdj.HasEdge(test.src, test.tgt))
			assert.False(t, dg.invAdj.HasEdge(test.tgt, test.src))
			// test that only the directed edge was removed
			if test.reverseExists {
				assert.True(t, dg.dirAdj.HasEdge(test.tgt, test.src))
			}
		})
	}
}

func TestDirGraphRemoveNode(t *testing.T) {
	tests := map[string]struct {
		node string
	}{
		"remove nonexistent node": {
			node: "x",
		},
		"remove existing node": {
			node: "a",
		},
		"remove existing node with self loop": {
			node: "d",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			tg := setupTestDirGraph()
			dg := tg.dg
			dg.RemoveNode(test.node)

			assert.NotContains(t, *dg.dirAdj, test.node)
			for _, nbrs := range *dg.dirAdj {
				_, ok := nbrs[test.node]
				assert.False(t, ok)
			}
			assert.NotContains(t, *dg.invAdj, test.node)
			for _, nbrs := range *dg.invAdj {
				_, ok := nbrs[test.node]
				assert.False(t, ok)
			}
		})
	}
}

func TestDirGraphGetNodes(t *testing.T) {
	emptyDirGraph, _ := NewDirGraph("")
	tests := map[string]testDirGraph{
		"empty graph": {
			dg:    emptyDirGraph,
			nodes: []string{},
		},
		"nonempty graph": *setupTestDirGraph(),
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			nodes := test.dg.GetNodes()
			assert.ElementsMatch(t, test.nodes, nodes)
		})
	}
}

func TestDirGraphGetInvNeighbors(t *testing.T) {
	tests := map[string]struct {
		node         string
		expectedNbrs []string
	}{
		"nonexistent node": {
			node:         "x",
			expectedNbrs: []string{},
		},
		"existing node": {
			node:         "a",
			expectedNbrs: []string{"b", "c", "d"},
		},
		"existing node with self loop": {
			node:         "d",
			expectedNbrs: []string{"a", "c", "d"},
		},
	}

	tg := setupTestDirGraph()
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			nbrs, ok := tg.dg.GetInvNeighbors(test.node)
			if len(test.expectedNbrs) == 0 {
				assert.False(t, ok)
				return
			}
			assert.True(t, ok)
			for n := range nbrs {
				assert.Contains(t, test.expectedNbrs, n)
			}
		})
	}
}

func TestDirGraphGetTotalDegree(t *testing.T) {
	tests := map[string]struct {
		node        string
		expectedDeg float64
	}{
		"nonexistent node": {
			node: "x",
		},
		"existing node": {
			node:        "a",
			expectedDeg: 29.5,
		},
		"existing node with self loop": {
			node:        "d",
			expectedDeg: 30.2,
		},
	}

	tg := setupTestDirGraph()
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			d, ok := tg.dg.GetTotalDegree(test.node)
			// node does not exist
			if test.expectedDeg == 0 {
				assert.False(t, ok)
				return
			}
			assert.True(t, ok)
			assert.InEpsilon(t, test.expectedDeg, d, float64EqualTol)
		})
	}
}

func TestDirGraphGetInDegree(t *testing.T) {
	tests := map[string]struct {
		node        string
		expectedDeg float64
	}{
		"nonexistent node": {
			node: "x",
		},
		"existing node": {
			node:        "a",
			expectedDeg: 7.0,
		},
		"existing node with self loop": {
			node:        "d",
			expectedDeg: 23.2,
		},
	}

	tg := setupTestDirGraph()
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			d, ok := tg.dg.GetInDegree(test.node)
			// node does not exist
			if test.expectedDeg == 0 {
				assert.False(t, ok)
				return
			}
			assert.True(t, ok)
			assert.InEpsilon(t, test.expectedDeg, d, float64EqualTol)
		})
	}
}
