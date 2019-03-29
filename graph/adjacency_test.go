package graph

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestPrint(t *testing.T) {
	a := dirAdj{}
	a.addEdge("x", "y", 1)
	a.addEdge("x", "z", 1)
	a.addEdge("y", "x", 1)
	a.print()
}

func TestAddEdge(t *testing.T) {
	a := dirAdj{}

	// test adding edge with default weight
	a.addEdge("x", "y", 1)
	xNbrs, xExists := a["x"]
	assert.True(t, xExists, "node x should exist in adjacency")
	assert.Len(t, xNbrs, 1, "node x should have 1 neighbor")
	w, yExists := xNbrs["y"]
	assert.True(t, yExists, "node y should x's neighbor")
	assert.Equal(t, 1.0, w, "edge from x to y should have weight 1")

	// test adding (upserting) edge with specified weight
	wgt := 3.7
	a.addEdge("x", "y", wgt)
	w = xNbrs["y"]
	assert.Equal(t, wgt, w, "edge from x to y should have weight 1")
}

type DirAdjTestSuite struct {
	suite.Suite
	A dirAdj
}

func TestDirAdjTestSuite(t *testing.T) {
	suite.Run(t, new(DirAdjTestSuite))
}

func (suite *DirAdjTestSuite) SetupTest() {
	suite.A = dirAdj{}
	suite.A.addEdge("x", "y", 1)
	suite.A.addEdge("x", "z", 1)
	suite.A.addEdge("y", "x", 1)
	suite.A.addEdge("y", "z", 1)
	suite.A.addEdge("z", "x", 1)
}

func (suite *DirAdjTestSuite) TestRemoveEdge() {
	// test removing nonexistent edge from existing node
	suite.A.removeEdge("x", "foo")
	xNbrs, xExists := suite.A["x"]
	assert.True(suite.T(), xExists, "node x should still exist in adjacency")
	assert.Len(suite.T(), xNbrs, 2, "node x should still have 2 neighbors")
	assert.Contains(suite.T(), xNbrs, "y", "node y should still be neighbor of x")
	assert.Contains(suite.T(), xNbrs, "z", "node z should still be neighbor of x")

	// test removing nonexistent edge from nonexisting node
	AOrig := dirAdj{}
	for outerK, outerV := range suite.A {
		AOrig[outerK] = outerV
	}
	suite.A.removeEdge("foo", "bar")
	isEqual := reflect.DeepEqual(suite.A, AOrig)
	assert.True(suite.T(), isEqual, "adjacency should be unchanged")

	// test removing existing edge
	suite.A.removeEdge("x", "w")
	xNbrs = suite.A["x"]
	assert.Len(suite.T(), xNbrs, 2, "node x should have 2 remaining neighbors")
	_, wExists := xNbrs["w"]
	assert.False(suite.T(), wExists, "node x should no longer have w as a neighbor")

	// test removing edge leaving no neighbors deletes the from node
	xNbrs = suite.A["x"]
	for n := range xNbrs {
		suite.A.removeEdge("x", n)
	}
	_, xExists = suite.A["x"]
	assert.False(suite.T(), xExists, "node x should no longer be a key once all neighbors removed")
	// y should still have an edge to x
	yNbrs, yExists := suite.A["y"]
	assert.True(suite.T(), yExists, "node y should still be a key")
	assert.Contains(suite.T(), yNbrs, "x", "node x should still be a neighbor of node y")
}

func (suite *DirAdjTestSuite) TestRemoveNode() {
	suite.A.removeNode("x")
	// x and z should no longer exist
	_, xExists := suite.A["x"]
	assert.False(suite.T(), xExists, "node x should no longer exist")
	_, zExists := suite.A["z"]
	assert.False(suite.T(), zExists, "node z should no longer exist")
	// y should not have an edge to x
	yNbrs := suite.A["y"]
	assert.NotContains(suite.T(), yNbrs, "x", "node x should no longer be a neighbor of node y")

	// test removing nonexistent node
	AOrig := dirAdj{}
	for outerK, outerV := range suite.A {
		AOrig[outerK] = outerV
	}
	suite.A.removeNode("foo")
	isEqual := reflect.DeepEqual(suite.A, AOrig)
	assert.True(suite.T(), isEqual, "adjacency should be unchanged")
}

func (suite *DirAdjTestSuite) TestGetNeighbors() {
	// test get neighbors for nonexistent node
	_, vExists := suite.A.getNeighbors("v")
	assert.False(suite.T(), vExists, "node v should not exist")

	// test get neighbors for existing node
	xNbrs, xExists := suite.A.getNeighbors("x")
	assert.True(suite.T(), xExists, "node x should have neighbors")
	assert.Len(suite.T(), xNbrs, 2, "node x should have 2 neighbors")
	assert.Contains(suite.T(), xNbrs, "y", "node y should be a neighbor of node x")
	assert.Contains(suite.T(), xNbrs, "z", "node z should be a neighbor of node x")
}

func (suite *DirAdjTestSuite) TestHasEdge() {
	assert.True(suite.T(), suite.A.hasEdge("x", "y"), "edge should exist between x and y")
	assert.False(suite.T(), suite.A.hasEdge("w", "y"), "edge should not exist between w and y")
}

func (suite *DirAdjTestSuite) TestGetEdgeWeight() {
	suite.A.addEdge("a", "b", 2.1)

	ab, found := suite.A.getEdgeWeight("a", "b")
	assert.True(suite.T(), found, "edge should exists between a and b")
	assert.Equal(suite.T(), 2.1, ab, "edge between a and b should have weight 2.1")

	_, found = suite.A.getEdgeWeight("b", "a")
	assert.False(suite.T(), found, "edge should not exist between b and a")
}
