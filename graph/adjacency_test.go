package graph

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// any way to assert anything here?
func TestPrint(t *testing.T) {
	a := DirAdj{}
	a.AddEdge("x", "y")
	a.AddEdge("x", "z")
	a.AddEdge("y", "x")
	a.Print()
}

func TestAddEdge(t *testing.T) {
	a := DirAdj{}

	// test adding edge with default weight
	a.AddEdge("x", "y")
	xNbrs, xExists := a["x"]
	assert.True(t, xExists, "node x should exist in adjacency")
	assert.Len(t, xNbrs, 1, "node x should have 1 neighbor")
	w, yExists := xNbrs["y"]
	assert.True(t, yExists, "node y should x's neighbor")
	assert.Equal(t, 1.0, w, "edge from x to y should have weight 1")

	// test adding (upserting) edge with specified weight
	wgt := 3.7
	a.AddEdge("x", "y", wgt)
	w = xNbrs["y"]
	assert.Equal(t, wgt, w, "edge from x to y should have weight 1")
}

func TestRemoveEdge(t *testing.T) {
	a := DirAdj{}
	a.AddEdge("x", "y")
	a.AddEdge("x", "z")
	a.AddEdge("y", "x")

	// make a copy of a for testing
	aOrig := DirAdj{}
	for outerK, outerV := range a {
		aOrig[outerK] = outerV
	}

	// test removing nonexistent edge from existing node
	a.RemoveEdge("x", "foo")
	xNbrs, xExists := a["x"]
	assert.True(t, xExists, "node x should still exist in adjacency")
	assert.Len(t, xNbrs, 2, "node x should still have 2 neighbors")
	assert.Contains(t, xNbrs, "y", "y should still be neighbor of x")
	assert.Contains(t, xNbrs, "z", "z should still be neighbor of x")

	// test removing nonexistent edge from nonexisting node
	a.RemoveEdge("foo", "bar")
	isEqual := reflect.DeepEqual(a, aOrig)
	assert.True(t, isEqual, "adjacency should be unchanged")

	// test removing existing edge
	a.AddEdge("foo", "bar")
	a.AddEdge("foo", "baz")
	a.RemoveEdge("foo", "bar")
	fooNbrs := a["foo"]
	assert.Len(t, fooNbrs, 1, "node foo should have 1 remaining neighbor")
	_, barExists := fooNbrs["bar"]
	assert.False(t, barExists, "foo should no longer have bar as a neighbor")
	_, bazExists := fooNbrs["baz"]
	assert.True(t, bazExists, "foo should still have baz as a neighbor")

	// test removing edge leaving no neighbors deletes the from node
	a.RemoveEdge("foo", "baz")
	fooNbrs, fooExists := a["foo"]
	assert.False(t, fooExists, "node foo should no longer exist")
}

func TestGetNeighbors(t *testing.T) {
	a := DirAdj{}
	a.AddEdge("x", "y")
	a.AddEdge("x", "z")
	a.AddEdge("y", "x")

	// test get neighbors for nonexistent node
	_, ok := a.GetNeighbors("foo")
	assert.False(t, ok, "node foo should not exist")

	// test get neighbors for existing node
	xNbrs, ok := a.GetNeighbors("x")
	assert.True(t, ok, "node x should have neighbors")
	assert.Len(t, xNbrs, 2)
	assert.Contains(t, xNbrs, "y")
	assert.Contains(t, xNbrs, "z")
}
