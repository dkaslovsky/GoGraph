package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestNewDirGraph(t *testing.T) {
	dg := NewDirGraph("test")
	assert.Equal(t, "test", dg.Name)
	assert.IsType(t, dirAdj{}, dg.invAdj)
}

func TestDirGraphAddEdge(t *testing.T) {
	dg := NewDirGraph("test")

	// test adding edge with default weight
	dg.AddEdge("x", "y")
	assert.True(t, dg.HasEdge("x", "y"))
	assert.False(t, dg.HasEdge("y", "x"))
	xOutNbrs, _ := dg.GetNeighbors("x")
	assert.Contains(t, xOutNbrs, "y")
	yInNbrs, _ := dg.GetInvNeighbors("y")
	assert.Contains(t, yInNbrs, "x")
	// test weight
	wgt, ok := dg.GetEdgeWeight("x", "y")
	assert.True(t, ok)
	assert.Equal(t, 1.0, wgt)

	// test upserting edge, specify weight
	dg.AddEdge("x", "y", 3.67)
	assert.True(t, dg.HasEdge("x", "y"))
	assert.False(t, dg.HasEdge("y", "x"))
	// test weight
	wgt, ok = dg.GetEdgeWeight("x", "y")
	assert.True(t, ok)
	assert.Equal(t, 3.67, wgt)
}

type DirGraphTestSuite struct {
	suite.Suite
	DG    *DirGraph
	Nodes []string
}

func TestDirGraphTestSuite(t *testing.T) {
	suite.Run(t, new(DirGraphTestSuite))
}

func (suite *DirGraphTestSuite) SetupTest() {
	suite.DG = NewDirGraph("test")
	suite.DG.AddEdge("a", "b", 1.5)
	suite.DG.AddEdge("a", "c", 2)
	suite.DG.AddEdge("b", "c", 3.3)
	suite.DG.AddEdge("c", "a", 7)
	suite.DG.AddEdge("c", "d", 1.1)
	suite.Nodes = []string{"a", "b", "c", "d"}
}

func (suite *DirGraphTestSuite) TestDirGraphRemoveEdge() {
	// test removing edge that does not exist
	assert.False(suite.T(), suite.DG.HasEdge("b", "a"))
	suite.DG.RemoveEdge("b", "a")
	assert.False(suite.T(), suite.DG.HasEdge("b", "a"))

	// test removing edges
	suite.DG.RemoveEdge("a", "b")
	assert.False(suite.T(), suite.DG.HasEdge("a", "b"))

	// test removing all edges for a node also removes the node
	suite.DG.RemoveEdge("a", "c")
	assert.Contains(suite.T(), suite.DG.GetNodes(), "a") // node a still has one edge left
	suite.DG.RemoveEdge("c", "a")
	assert.NotContains(suite.T(), suite.DG.GetNodes(), "a")
}

func (suite *DirGraphTestSuite) TestDirGraphRemoveNode() {
	suite.DG.RemoveNode("a")
	nodes := suite.DG.GetNodes()
	assert.NotContains(suite.T(), nodes, "a")
	for _, node := range nodes {
		assert.False(suite.T(), suite.DG.HasEdge(node, "a"))
		assert.False(suite.T(), suite.DG.HasEdge("a", node))
	}
}

func (suite *DirGraphTestSuite) TestDirGraphGetNodes() {
	nodes := suite.DG.GetNodes()
	assert.Len(suite.T(), nodes, len(suite.Nodes))
	for _, node := range suite.Nodes {
		assert.Contains(suite.T(), nodes, node)
	}

	// test result on empty graph
	dgEmpty := NewDirGraph("testEmpty")
	nodes = dgEmpty.GetNodes()
	assert.Empty(suite.T(), nodes)
}

func (suite *DirGraphTestSuite) TestDirGraphGetInvNeighbors() {
	type testCase struct {
		node    string
		exists  bool
		expNbrs []string
	}
	var table = map[string]testCase{
		"get inv neighbors for node a": {
			node:    "a",
			exists:  true,
			expNbrs: []string{"c"},
		},
		"get inv neighbors for node b": {
			node:    "b",
			exists:  true,
			expNbrs: []string{"a"},
		},
		"get inv neighbors for node c": {
			node:    "c",
			exists:  true,
			expNbrs: []string{"a", "b"},
		},
		"get inv neighbors for node d": {
			node:    "d",
			exists:  true,
			expNbrs: []string{"c"},
		},
		"get inv neighbors for nonexistent node": {
			node:   "z",
			exists: false,
		},
	}

	for _, tt := range table {
		nbrs, ok := suite.DG.GetInvNeighbors(tt.node)
		if !tt.exists {
			assert.False(suite.T(), ok)
			continue
		}
		assert.Len(suite.T(), nbrs, len(tt.expNbrs))
		for _, n := range tt.expNbrs {
			assert.Contains(suite.T(), nbrs, n)
		}
	}
}

func (suite *DirGraphTestSuite) TestDirGraphGetTotalDegree() {
	d, ok := suite.DG.GetTotalDegree("a")
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 10.5, d)
	_, ok = suite.DG.GetTotalDegree("foo")
	assert.False(suite.T(), ok)
}

func (suite *DirGraphTestSuite) TestDirGraphGetOutDegree() {
	d, ok := suite.DG.GetOutDegree("a")
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 3.5, d)
	_, ok = suite.DG.GetOutDegree("foo")
	assert.False(suite.T(), ok)
}

func (suite *DirGraphTestSuite) TestDirGraphGetInDegree() {
	d, ok := suite.DG.GetInDegree("a")
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 7.0, d)
	_, ok = suite.DG.GetInDegree("foo")
	assert.False(suite.T(), ok)
}
