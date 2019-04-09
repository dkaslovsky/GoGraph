package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type edge struct {
	src string
	tgt string
}

type weightedEdge struct {
	src string
	tgt string
	wgt float64
}

func setupAdj() dirAdj {
	return dirAdj{
		"x": {"y": 1, "z": 1},
		"y": {"x": 1, "z": 1},
		"z": {"x": 2.2},
	}
}

func TestAddDirectedEdge(t *testing.T) {
	tests := map[string]struct {
		a dirAdj
		e weightedEdge
	}{
		"add edge with integer weight": {
			dirAdj{},
			weightedEdge{src: "a", tgt: "b", wgt: 1},
		},
		"add edge with float weight": {
			dirAdj{},
			weightedEdge{src: "a", tgt: "b", wgt: 3.4},
		},
		"upsert edge": {
			dirAdj{"a": {"b": 3.4}},
			weightedEdge{src: "a", tgt: "b", wgt: 10.10},
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
			wgt, ok := nbrs[e.tgt]
			assert.Equal(t, float64(e.wgt), wgt)
		})
	}
}

func TestRemoveDirectedEdge(t *testing.T) {
	tests := map[string]struct {
		src            string
		tgts           []string
		tgtsRemaining  []string
		srcStillExists bool
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
			src := test.src
			for _, tgt := range test.tgts {
				a.removeDirectedEdge(src, tgt)
			}

			nbrs, ok := a[src]
			if len(test.tgtsRemaining) == 0 {
				assert.False(t, ok)
				return
			}
			assert.True(t, ok)
			for _, tgt := range test.tgtsRemaining {
				assert.Contains(t, nbrs, tgt)
			}
		})
	}
}
