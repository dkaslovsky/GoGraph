package graph

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// What assertions can be made here?
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

func TestRemoveEdge(t *testing.T) {
	a := dirAdj{}
	a.addEdge("x", "y", 1)
	a.addEdge("x", "w", 1)
	a.addEdge("x", "z", 1)
	a.addEdge("y", "x", 1)

	// test removing nonexistent edge from existing node
	a.removeEdge("x", "foo")
	xNbrs, xExists := a["x"]
	assert.True(t, xExists, "node x should still exist in adjacency")
	assert.Len(t, xNbrs, 3, "node x should still have 2 neighbors")
	assert.Contains(t, xNbrs, "y", "y should still be neighbor of x")
	assert.Contains(t, xNbrs, "w", "w should still be neighbor of x")
	assert.Contains(t, xNbrs, "z", "z should still be neighbor of x")

	// test removing nonexistent edge from nonexisting node
	aOrig := dirAdj{}
	for outerK, outerV := range a {
		aOrig[outerK] = outerV
	}
	a.removeEdge("foo", "bar")
	isEqual := reflect.DeepEqual(a, aOrig)
	assert.True(t, isEqual, "adjacency should be unchanged")

	// test removing existing edge
	a.removeEdge("x", "w")
	xNbrs = a["x"]
	assert.Len(t, xNbrs, 2, "node x should have 2 remaining neighbors")
	_, wExists := xNbrs["w"]
	assert.False(t, wExists, "node x should no longer have w as a neighbor")

	// test removing edge leaving no neighbors deletes the from node
	xNbrs = a["x"]
	for n := range xNbrs {
		a.removeEdge("x", n)
	}
	_, xExists = a["x"]
	assert.False(t, xExists, "node x should no longer be a key once all neighbors are removed")
	// y should still have an edge to x
	yNbrs, yExists := a["y"]
	assert.True(t, yExists, "node y should still be a key")
	assert.Contains(t, yNbrs, "x", "node x should still be a neighbor of node y")
}

func TestRemoveNode(t *testing.T) {
	a := dirAdj{}
	a.addEdge("x", "y", 1)
	a.addEdge("x", "z", 1)
	a.addEdge("y", "x", 1)
	a.addEdge("y", "z", 1)
	a.addEdge("z", "x", 1)

	a.removeNode("x")
	// x and z should no longer exist
	_, xExists := a["x"]
	assert.False(t, xExists, "node x should no longer exist")
	_, zExists := a["z"]
	assert.False(t, zExists, "node z should no longer exist")
	// y should not have an edge to x
	yNbrs := a["y"]
	assert.NotContains(t, yNbrs, "x", "node x should no longer be a neighbor of node y")

	// test removing nonexistent node
	aOrig := dirAdj{}
	for outerK, outerV := range a {
		aOrig[outerK] = outerV
	}
	a.removeNode("foo")
	isEqual := reflect.DeepEqual(a, aOrig)
	assert.True(t, isEqual, "adjacency should be unchanged")
}

func TestGetNeighbors(t *testing.T) {
	a := dirAdj{}
	a.addEdge("x", "y", 1)
	a.addEdge("x", "z", 1)
	a.addEdge("y", "x", 1)

	// test get neighbors for nonexistent node
	_, vExists := a.getNeighbors("v")
	assert.False(t, vExists, "node v should not exist")

	// test get neighbors for existing node
	xNbrs, xExists := a.getNeighbors("x")
	assert.True(t, xExists, "node x should have neighbors")
	assert.Len(t, xNbrs, 2, "node x should have 2 neighbors")
	assert.Contains(t, xNbrs, "y", "node y should be a neighbor of node x")
	assert.Contains(t, xNbrs, "z", "node z should be a neighbor of node x")
}

func TestHasEdge(t *testing.T) {
	a := dirAdj{}
	a.addEdge("x", "y", 1)
	assert.True(t, a.hasEdge("x", "y"), "edge should exist between x and y")
	assert.False(t, a.hasEdge("x", "z"), "edge should not exist between x and z")
	assert.False(t, a.hasEdge("w", "y"), "edge should not exist between w and y")
}

func TestGetEdgeWeight(t *testing.T) {
	a := dirAdj{}
	a.addEdge("x", "y", 2.1)

	xy, found := a.getEdgeWeight("x", "y")
	assert.True(t, found, "edge should exists between x and y")
	assert.Equal(t, 2.1, xy, "edge between x and y should have weight 2.0")

	_, found = a.getEdgeWeight("x", "z")
	assert.False(t, found, "edge should not exist between x and y")
}
