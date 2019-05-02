package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"

	n "github.com/dkaslovsky/GoGraph/node"
)

// float64EqualTol is the tolerance at which we consider float64s equal
const float64EqualTol = 1e-9

type testEdge struct {
	src n.Node
	tgt n.Node
	wgt float64
}

func setupAdj() dirAdj {
	return dirAdj{
		"x": {"y": 1, "z": 1},
		"y": {"x": 3.2, "z": 9.7},
		"z": {"x": 2.2, "z": 3.4},
	}
}

func TestAddDirectedEdge(t *testing.T) {
	tests := map[string]struct {
		a dirAdj
		testEdge
	}{
		"add edge with integer weight": {
			dirAdj{},
			testEdge{src: "a", tgt: "b", wgt: 1},
		},
		"add edge with float weight": {
			dirAdj{},
			testEdge{src: "a", tgt: "b", wgt: 3.4},
		},
		"upsert edge": {
			dirAdj{"a": {"b": 3.4}},
			testEdge{src: "a", tgt: "b", wgt: 10.10},
		},
		"add self loop": {
			dirAdj{},
			testEdge{src: "a", tgt: "a", wgt: 1.1},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			a := test.a
			a.addDirectedEdge(test.src, test.tgt, test.wgt)

			// test edge exists
			nbrs, ok := a[test.src]
			assert.True(t, ok)
			assert.Contains(t, nbrs, test.tgt)
			// test weight
			wgt, _ := nbrs[test.tgt]
			assert.Equal(t, float64(test.wgt), wgt)
		})
	}
}

func TestRemoveDirectedEdge(t *testing.T) {
	tests := map[string]struct {
		src           n.Node
		tgts          []n.Node
		tgtsRemaining []n.Node
	}{
		"remove nonexistent edge from existing node": {
			src:           "x",
			tgts:          []n.Node{"foo"},
			tgtsRemaining: []n.Node{"y", "z"},
		},
		"remove nonexistent edge from nonexistent node": {
			src:  "foo",
			tgts: []n.Node{"bar"},
		},
		"remove existing edge": {
			src:           "x",
			tgts:          []n.Node{"y"},
			tgtsRemaining: []n.Node{"z"},
		},
		"remove all edges from node": {
			src:  "y",
			tgts: []n.Node{"x", "z"},
		},
		"remove self loop": {
			src:           "z",
			tgts:          []n.Node{"z"},
			tgtsRemaining: []n.Node{"x"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			a := setupAdj()
			for _, tgt := range test.tgts {
				a.removeDirectedEdge(test.src, tgt)
			}

			nbrs, ok := a[test.src]
			// src should be removed if no target nodes remain
			if len(test.tgtsRemaining) == 0 {
				assert.False(t, ok)
				return
			}
			assert.True(t, ok)

			// test that only the specified nodes were
			// removed and the others remain
			for _, tgt := range test.tgts {
				assert.NotContains(t, nbrs, tgt)
			}
			for _, tgt := range test.tgtsRemaining {
				assert.Contains(t, nbrs, tgt)
			}
		})
	}
}

func TestGetSrcNodes(t *testing.T) {
	tests := map[string]struct {
		a             dirAdj
		expectedNodes []n.Node
	}{
		"empty adjacency": {
			a:             dirAdj{},
			expectedNodes: []n.Node{},
		},
		"nonempty adjacency": {
			a:             setupAdj(),
			expectedNodes: []n.Node{"x", "y", "z"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			nodes := test.a.getSrcNodes()
			assert.ElementsMatch(t, nodes, test.expectedNodes)
		})
	}
}

func TestGetNeighbors(t *testing.T) {
	tests := map[string]struct {
		node         n.Node
		expectedNbrs []n.Node
	}{
		"nonexistent node": {
			node:         "a",
			expectedNbrs: []n.Node{},
		},
		"existing node": {
			node:         "x",
			expectedNbrs: []n.Node{"y", "z"},
		},
		"existing node with self loop": {
			node:         "z",
			expectedNbrs: []n.Node{"x", "z"},
		},
	}

	a := setupAdj()
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			nbrs, ok := a.GetNeighbors(test.node)
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

func TestGetOutDegree(t *testing.T) {
	tests := map[string]struct {
		node        n.Node
		expectedDeg float64
	}{
		"nonexistent node": {
			node: "a",
		},
		"existing node": {
			node:        "y",
			expectedDeg: 12.9,
		},
		"existing node with self loop": {
			node:        "z",
			expectedDeg: 5.6,
		},
	}

	a := setupAdj()
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			d, ok := a.GetOutDegree(test.node)
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

func TestHasEdge(t *testing.T) {
	tests := map[string]struct {
		src    n.Node
		tgt    n.Node
		exists bool
	}{
		"nonexistent edge between nonexistent nodes": {
			src:    "a",
			tgt:    "b",
			exists: false,
		},
		"nonexistent edge from existing node to nonexistent node": {
			src:    "x",
			tgt:    "a",
			exists: false,
		},
		"nonexistent edge from nonexistent node to existing node": {
			src:    "a",
			tgt:    "x",
			exists: false,
		},
		"nonexistent edge between existing nodes": {
			src:    "z",
			tgt:    "y",
			exists: false,
		},
		"existing edge": {
			src:    "z",
			tgt:    "x",
			exists: true,
		},
		"existing self loop edge": {
			src:    "z",
			tgt:    "z",
			exists: true,
		},
	}

	a := setupAdj()
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			exists := a.HasEdge(test.src, test.tgt)
			assert.Equal(t, test.exists, exists)
		})
	}
}

func TestGetEdgeWeight(t *testing.T) {
	tests := map[string]struct {
		src    n.Node
		tgt    n.Node
		weight float64
	}{
		"nonexistent edge between nonexistent nodes": {
			src: "a",
			tgt: "b",
		},
		"nonexistent edge from existing node to nonexistent node": {
			src: "x",
			tgt: "a",
		},
		"nonexistent edge from nonexistent node to existing node": {
			src: "a",
			tgt: "x",
		},
		"nonexistent edge between existing nodes": {
			src: "z",
			tgt: "y",
		},
		"existing edge": {
			src:    "x",
			tgt:    "y",
			weight: 1,
		},
		"existing edge reversed direction": {
			src:    "y",
			tgt:    "x",
			weight: 3.2,
		},
		"existing self loop edge": {
			src:    "z",
			tgt:    "z",
			weight: 3.4,
		},
	}

	a := setupAdj()
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			weight, ok := a.GetEdgeWeight(test.src, test.tgt)
			// no edge present
			if test.weight == 0 {
				assert.False(t, ok)
				return
			}
			assert.True(t, ok)
			assert.Equal(t, test.weight, weight)
		})
	}
}
