package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestDirGraphAddEdge(t *testing.T) {
	dg := NewDirGraph("test")

	// test adding edge with default weight
	dg.AddEdge("x", "y")
	assert.True(t, dg.HasEdge("x", "y"))
	assert.False(t, dg.HasEdge("y", "x"))
	xOutNbrs, _ := dg.GetOutNeighbors("x")
	assert.Contains(t, xOutNbrs, "y")
	yInNbrs, _ := dg.GetInNeighbors("y")
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
	suite.DG.RemoveEdge("a", "b")
	assert.False(suite.T(), suite.DG.HasEdge("a", "b"))

	// test removing edge that does not exist
	suite.DG.RemoveEdge("b", "a")
	assert.False(suite.T(), suite.DG.HasEdge("b", "a"))

	// test removing edge that leavs a node with no neighbors also removes the node
	suite.DG.RemoveEdge("c", "a")
	suite.DG.RemoveEdge("c", "d")
	assert.NotContains(suite.T(), suite.DG.outAdj, "c")
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

type getDirNbrTest struct {
	node    string
	exists  bool
	expNbrs []string
}

var getOutNbrTest = []getDirNbrTest{
	{
		node:    "a",
		exists:  true,
		expNbrs: []string{"b", "c"},
	},
	{
		node:    "b",
		exists:  true,
		expNbrs: []string{"c"},
	},
	{
		node:    "c",
		exists:  true,
		expNbrs: []string{"a", "d"},
	},
	{
		node:   "d",
		exists: false,
	},
	{
		node:   "z",
		exists: false,
	},
}

func (suite *DirGraphTestSuite) TestDirGraphGetOutNeighbors() {
	for _, tt := range getOutNbrTest {
		nbrs, ok := suite.DG.GetOutNeighbors(tt.node)
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

var getInNbrTest = []getDirNbrTest{
	{
		node:    "a",
		exists:  true,
		expNbrs: []string{"c"},
	},
	{
		node:    "b",
		exists:  true,
		expNbrs: []string{"a"},
	},
	{
		node:    "c",
		exists:  true,
		expNbrs: []string{"a", "b"},
	},
	{
		node:    "d",
		exists:  true,
		expNbrs: []string{"c"},
	},
	{
		node:   "z",
		exists: false,
	},
}

func (suite *DirGraphTestSuite) TestDirGraphGetInNeighbors() {
	for _, tt := range getInNbrTest {
		nbrs, ok := suite.DG.GetInNeighbors(tt.node)
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

func (suite *DirGraphTestSuite) TestDirGraphHasEdge() {
	assert.True(suite.T(), suite.DG.HasEdge("a", "c"))
	assert.True(suite.T(), suite.DG.HasEdge("c", "a"))
	assert.False(suite.T(), suite.DG.HasEdge("foo", "bar"))
}

func (suite *DirGraphTestSuite) TestDirGraphGetEdgeWeight() {
	w, ok := suite.DG.GetEdgeWeight("a", "c")
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 2.0, w)

	w, ok = suite.DG.GetEdgeWeight("c", "a")
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 7.0, w)

	w, ok = suite.DG.GetEdgeWeight("foo", "bar")
	assert.False(suite.T(), ok)
}
