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
	nodes := []string{"a", "b", "c"}
	invNodes := []string{"a", "b", "c", "d"}
	edges := []byte("a b 1.5\na c 2\nb c 3.3\nc d 1.1\nd a 7")
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
			assert.True(t, e.existsIn(a))
			assert.True(t, e.reverseExistsIn(invA))
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
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			tg := setupTestDirGraph()
			dg := tg.dg
			dg.AddEdge(test.src, test.tgt)
			assert.True(t, test.existsIn(*dg.dirAdj))
			assert.True(t, test.reverseExistsIn(*dg.invAdj))
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
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			tg := setupTestDirGraph()
			dg := tg.dg
			dg.AddEdge(test.src, test.tgt, test.wgt)
			assert.True(t, test.existsIn(*dg.dirAdj))
			assert.True(t, test.reverseExistsIn(*dg.invAdj))
		})
	}
}

func TestDirGraphRemoveEdge(t *testing.T) {
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
			tg := setupTestDirGraph()
			dg := tg.dg
			dg.RemoveEdge(test.src, test.tgt)
			assert.False(t, test.existsIn(*dg.dirAdj))
			//assert.False(t, test.reverseExistsIn(*dg.invAdj))
			// add before and after checks, remove an edge that is symmetric as another test case
		})
	}
}

// func (suite *DirGraphTestSuite) TestDirGraphRemoveEdge() {
// 	// test removing edge that does not exist
// 	assert.False(suite.T(), suite.DG.HasEdge("b", "a"))
// 	suite.DG.RemoveEdge("b", "a")
// 	assert.False(suite.T(), suite.DG.HasEdge("b", "a"))

// 	// test removing edges
// 	suite.DG.RemoveEdge("a", "b")
// 	assert.False(suite.T(), suite.DG.HasEdge("a", "b"))

// 	// test removing all edges for a node also removes the node
// 	suite.DG.RemoveEdge("a", "c")
// 	assert.Contains(suite.T(), suite.DG.GetNodes(), "a") // node a still has one edge left
// 	suite.DG.RemoveEdge("c", "a")
// 	assert.NotContains(suite.T(), suite.DG.GetNodes(), "a")
// }

// func (suite *DirGraphTestSuite) TestDirGraphRemoveNode() {
// 	suite.DG.RemoveNode("a")
// 	nodes := suite.DG.GetNodes()
// 	assert.NotContains(suite.T(), nodes, "a")
// 	for _, node := range nodes {
// 		assert.False(suite.T(), suite.DG.HasEdge(node, "a"))
// 		assert.False(suite.T(), suite.DG.HasEdge("a", node))
// 	}
// }

// func (suite *DirGraphTestSuite) TestDirGraphGetNodes() {
// 	nodes := suite.DG.GetNodes()
// 	assert.Len(suite.T(), nodes, len(suite.Nodes))
// 	for _, node := range suite.Nodes {
// 		assert.Contains(suite.T(), nodes, node)
// 	}

// 	// test result on empty graph
// 	dgEmpty, _ := NewDirGraph("testEmpty")
// 	nodes = dgEmpty.GetNodes()
// 	assert.Empty(suite.T(), nodes)
// }

// func (suite *DirGraphTestSuite) TestDirGraphGetInvNeighbors() {
// 	type testCase struct {
// 		node    string
// 		exists  bool
// 		expNbrs []string
// 	}
// 	var table = map[string]testCase{
// 		"get inv neighbors for node a": {
// 			node:    "a",
// 			exists:  true,
// 			expNbrs: []string{"c"},
// 		},
// 		"get inv neighbors for node b": {
// 			node:    "b",
// 			exists:  true,
// 			expNbrs: []string{"a"},
// 		},
// 		"get inv neighbors for node c": {
// 			node:    "c",
// 			exists:  true,
// 			expNbrs: []string{"a", "b"},
// 		},
// 		"get inv neighbors for node d": {
// 			node:    "d",
// 			exists:  true,
// 			expNbrs: []string{"c"},
// 		},
// 		"get inv neighbors for nonexistent node": {
// 			node:   "z",
// 			exists: false,
// 		},
// 	}

// 	for _, tt := range table {
// 		nbrs, ok := suite.DG.GetInvNeighbors(tt.node)
// 		if !tt.exists {
// 			assert.False(suite.T(), ok)
// 			continue
// 		}
// 		assert.Len(suite.T(), nbrs, len(tt.expNbrs))
// 		for _, n := range tt.expNbrs {
// 			assert.Contains(suite.T(), nbrs, n)
// 		}
// 	}
// }

// func (suite *DirGraphTestSuite) TestDirGraphGetTotalDegree() {
// 	d, ok := suite.DG.GetTotalDegree("a")
// 	assert.True(suite.T(), ok)
// 	assert.Equal(suite.T(), 10.5, d)
// 	_, ok = suite.DG.GetTotalDegree("foo")
// 	assert.False(suite.T(), ok)
// }

// func (suite *DirGraphTestSuite) TestDirGraphGetInDegree() {
// 	d, ok := suite.DG.GetInDegree("a")
// 	assert.True(suite.T(), ok)
// 	assert.Equal(suite.T(), 7.0, d)
// 	_, ok = suite.DG.GetInDegree("foo")
// 	assert.False(suite.T(), ok)
// }
