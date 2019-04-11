package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// float64EqualTol is the tolerance at which we consider float64s equal
const float64EqualTol = 1e-9

func setupAdj() dirAdj {
	return dirAdj{
		"x": {"y": 1, "z": 1},
		"y": {"x": 3.2, "z": 9.7},
		"z": {"x": 2.2},
	}
}

func TestAddDirectedEdge(t *testing.T) {
	type edge struct {
		src string
		tgt string
		wgt float64
	}

	tests := map[string]struct {
		a dirAdj
		e edge
	}{
		"add edge with integer weight": {
			dirAdj{},
			edge{src: "a", tgt: "b", wgt: 1},
		},
		"add edge with float weight": {
			dirAdj{},
			edge{src: "a", tgt: "b", wgt: 3.4},
		},
		"upsert edge": {
			dirAdj{"a": {"b": 3.4}},
			edge{src: "a", tgt: "b", wgt: 10.10},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			a := test.a
			e := test.e
			a.addDirectedEdge(e.src, e.tgt, e.wgt)

			// test edge exists
			nbrs, ok := a[e.src]
			assert.True(t, ok)
			assert.Contains(t, nbrs, e.tgt)
			// test weight
			wgt, _ := nbrs[e.tgt]
			assert.Equal(t, float64(e.wgt), wgt)
		})
	}
}

func TestRemoveDirectedEdge(t *testing.T) {
	tests := map[string]struct {
		src           string
		tgts          []string
		tgtsRemaining []string
	}{
		"remove nonexistent edge from existing node": {
			src:           "x",
			tgts:          []string{"foo"},
			tgtsRemaining: []string{"y", "z"},
		},
		"remove nonexistent edge from nonexistent node": {
			src:  "foo",
			tgts: []string{"bar"},
		},
		"remove existing edge": {
			src:           "x",
			tgts:          []string{"y"},
			tgtsRemaining: []string{"z"},
		},
		"remove all edges from node": {
			src:  "y",
			tgts: []string{"x", "z"},
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

			// test that only the specified nodes were removed
			// and the other remain
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
		expectedNodes []string
	}{
		"empty adjacency": {
			a:             dirAdj{},
			expectedNodes: []string{},
		},
		"nonempty adjacency": {
			a:             setupAdj(),
			expectedNodes: []string{"x", "y", "z"},
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
		node         string
		expectedNbrs []string
	}{
		"nonexistent node": {
			node:         "a",
			expectedNbrs: []string{},
		},
		"existing node": {
			node:         "x",
			expectedNbrs: []string{"y", "z"},
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
		node        string
		expectedDeg float64
	}{
		"nonexistent node": {
			node: "a",
		},
		"existing node multiple neighbors": {
			node:        "y",
			expectedDeg: 12.9,
		},
		"existing node single neighbor": {
			node:        "z",
			expectedDeg: 2.2,
		},
	}

	a := setupAdj()
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			d, ok := a.GetOutDegree(test.node)
			if test.expectedDeg == 0 {
				assert.False(t, ok)
				return
			}
			assert.InEpsilon(t, test.expectedDeg, d, float64EqualTol)
		})
	}
}

func TestHasEdge(t *testing.T) {
	tests := map[string]struct {
		src    string
		tgt    string
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
		src    string
		tgt    string
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
	}

	a := setupAdj()
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			weight, ok := a.GetEdgeWeight(test.src, test.tgt)
			if test.weight == 0 {
				assert.False(t, ok)
				return
			}
			assert.True(t, ok)
			assert.Equal(t, test.weight, weight)
		})
	}
}
