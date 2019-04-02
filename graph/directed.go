package graph

import (
	"io"
)

// DirGraph is an adjacency map representation of a directed graph
type DirGraph struct {
	Graph
}

// NewDirGraph creates a new directed graph
func NewDirGraph(name string, readers ...io.ReadCloser) (*DirGraph, error) {
	dg := &DirGraph{
		Graph{
			dirAdj: &dirAdj{},
			Name:   name,
			invAdj: &dirAdj{},
		},
	}
	for _, r := range readers {
		err := dg.addFromReader(r)
		if err != nil {
			return dg, err
		}
	}
	return dg, nil
}

// RemoveNode removes a node entirely from a DirGraph such that
// no edges exist between it an any other node
func (dg *DirGraph) RemoveNode(node string) {
	// remove node from dirAdj
	if nbrs, ok := dg.GetNeighbors(node); ok {
		for n := range nbrs {
			dg.RemoveEdge(node, n)
		}
	}
	// also remove node from invAdj
	if nbrs, ok := dg.GetInvNeighbors(node); ok {
		for n := range nbrs {
			dg.RemoveEdge(n, node)
		}
	}
}

// GetNodes gets a slice of all nodes in a DirGraph
func (dg *DirGraph) GetNodes() (nodes []string) {

	nodes = dg.getFromNodes() // guaranteed to be unique

	// maintain map keyed by nodes to avoid adding duplicates from invAdj
	set := map[string]struct{}{}
	for _, node := range nodes {
		set[node] = struct{}{}
	}

	// append invAdj node only if it is not in the set
	invNodes := dg.invAdj.getFromNodes() // guaranteed to be unique
	for _, node := range invNodes {
		if _, ok := set[node]; !ok {
			nodes = append(nodes, node)
		}
	}

	return nodes
}

// GetTotalDegree calculates the sum of weights of all edges from and to a node
func (dg *DirGraph) GetTotalDegree(node string) (deg float64, found bool) {
	outDeg, ok := dg.GetOutDegree(node)
	if !ok {
		return deg, false
	}
	inDeg, ok := dg.GetInDegree(node)
	if !ok {
		return deg, false
	}
	return inDeg + outDeg, true
}
