package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestNewGraph(t *testing.T) {
	g := NewGraph("test")
	assert.Equal(t, "test", g.Name)
}

func TestGraphAddEdge(t *testing.T) {
	g := NewGraph("test")

	// test adding edge with default weight
	g.AddEdge("x", "y")
	assert.True(t, g.HasEdge("x", "y"))
	assert.True(t, g.HasEdge("y", "x"))
	xNbrs, _ := g.GetNeighbors("x")
	assert.Contains(t, xNbrs, "y")
	yNbrs, _ := g.GetNeighbors("y")
	assert.Contains(t, yNbrs, "x")
	// test weight
	wgt, ok := g.GetEdgeWeight("x", "y")
	assert.True(t, ok)
	assert.Equal(t, 1.0, wgt)
	wgt, ok = g.GetEdgeWeight("y", "x")
	assert.True(t, ok)
	assert.Equal(t, 1.0, wgt)

	// test upserting edge, specify weight
	g.AddEdge("x", "y", 3.67)
	assert.True(t, g.HasEdge("x", "y"))
	assert.True(t, g.HasEdge("y", "x"))
	// test weight
	wgt, ok = g.GetEdgeWeight("x", "y")
	assert.True(t, ok)
	assert.Equal(t, 3.67, wgt)
	wgt, ok = g.GetEdgeWeight("y", "x")
	assert.True(t, ok)
	assert.Equal(t, 3.67, wgt)

	// test upserting edge in reverse order
	g.AddEdge("y", "x", 5.55)
	assert.True(t, g.HasEdge("x", "y"))
	assert.True(t, g.HasEdge("y", "x"))
	// test weight
	wgt, ok = g.GetEdgeWeight("x", "y")
	assert.True(t, ok)
	assert.Equal(t, 5.55, wgt)
	wgt, ok = g.GetEdgeWeight("y", "x")
	assert.True(t, ok)
	assert.Equal(t, 5.55, wgt)
}

type GraphTestSuite struct {
	suite.Suite
	G     *Graph
	Nodes []string
}

func TestGraphTestSuite(t *testing.T) {
	suite.Run(t, new(GraphTestSuite))
}

func (suite *GraphTestSuite) SetupTest() {
	suite.G = NewGraph("test")
	suite.G.AddEdge("a", "b", 1.5)
	suite.G.AddEdge("a", "c", 2)
	suite.G.AddEdge("b", "c", 3.3)
	suite.G.AddEdge("c", "d", 1.1)
	suite.G.AddEdge("d", "a", 7)
	suite.Nodes = []string{"a", "b", "c", "d"}
}

func (suite *GraphTestSuite) TestGraphRemoveEdge() {
	suite.G.RemoveEdge("b", "a")
	assert.False(suite.T(), suite.G.HasEdge("a", "b"))
	assert.False(suite.T(), suite.G.HasEdge("b", "a"))

	// test removing edge that does not exist
	assert.False(suite.T(), suite.G.HasEdge("a", "b"))
	assert.False(suite.T(), suite.G.HasEdge("b", "a"))
	suite.G.RemoveEdge("a", "b")
	assert.False(suite.T(), suite.G.HasEdge("a", "b"))
	assert.False(suite.T(), suite.G.HasEdge("b", "a"))

	// test removing edge that leavs a node with no neighbors also removes the node
	suite.G.RemoveEdge("a", "c")
	suite.G.RemoveEdge("a", "d")
	assert.NotContains(suite.T(), suite.G.GetNodes(), "a")
}

func (suite *GraphTestSuite) TestGraphRemoveNode() {
	suite.G.RemoveNode("d")
	nodes := suite.G.GetNodes()
	assert.NotContains(suite.T(), nodes, "d")
	for _, node := range nodes {
		assert.False(suite.T(), suite.G.HasEdge(node, "d"))
	}
}

func (suite *GraphTestSuite) TestGraphGetNodes() {
	nodes := suite.G.GetNodes()
	assert.Len(suite.T(), nodes, len(suite.Nodes))
	for _, node := range suite.Nodes {
		assert.Contains(suite.T(), nodes, node)
	}

	// test result on empty graph
	gEmpty := NewGraph("testEmpty")
	nodes = gEmpty.GetNodes()
	assert.Empty(suite.T(), nodes)
}

func (suite *GraphTestSuite) TestGraphGetDegree() {
	d, ok := suite.G.GetDegree("a")
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 10.5, d)
	_, ok = suite.G.GetDegree("foo")
	assert.False(suite.T(), ok)
}
